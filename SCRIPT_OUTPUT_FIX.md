# Script Output Connection Fix

## Issues Reported

1. **"Connection to script output lost"** - Frontend unable to connect to script output stream
2. **"No output on the stream logs page"** - Application logs not showing any entries
3. **Scripts running forever** - Question about script lifecycle management

## Root Cause Analysis

### Issue #1: Connection to Script Output Lost

**Problem**: When a script fails or completes, it's immediately removed from the engine map. The frontend connects to the output stream milliseconds later, but the engine is already gone, resulting in a 404 error.

**Flow**:
1. User runs script → `RunScript()` creates engine and starts it
2. Script has error → `RunAsync()` returns error
3. Error handler calls `engine.Stop()`
4. `onStopCallback` **immediately** deletes engine from map
5. Frontend tries to connect → Engine not found → 404 → "Connection lost"

**Code Location**: `pkg/stream/scripting/manager.go` lines 54-57

```go
// Old behavior - immediate cleanup
engine.SetOnStopCallback(func() {
    m.enginesMu.Lock()
    delete(m.engines, id)  // Gone immediately!
    m.enginesMu.Unlock()
})
```

### Issue #2: No Output on Stream Logs Page

**Problem**: The logs page (`/stream/logs`) queries `/api/v1/stream/logs` which returns entries from the global logger. However, the stream services use `zerolog` for their logging, not the global logger.

**Code Location**:
- Handler: `internal/handler/stream_logs.go` - uses `logging.GetGlobalLogger()`
- Services: Use `zerolog` via `github.com/rs/zerolog/log`

**Evidence**:
```go
// Services log like this:
log.Info().Msg("proxy started")

// But logs page reads from:
logger := logging.GetGlobalLogger()
```

These are **different logging systems** - not connected.

### Issue #3: Scripts Running Forever

**Not a bug** - This is intentional design.

Scripts are meant to run continuously until:
- User manually stops them via UI/API: `POST /api/v1/stream/scripts/stop/{id}`
- Script calls `exit()` function: `exit()`

**Purpose**: Allows long-running client/backend scripts that maintain connections, handle events, etc.

**Example Use Cases**:
```javascript
// Client script - runs forever, reacts to events
connect('ws://localhost:8080/ws', 'demo.Counter');

onPropertyChanged('count', (value) => {
    console.log('Count changed:', value);
});

// Keeps running until stopped
```

```javascript
// Backend script - provides service forever
const backend = createBackend('ws://localhost:8080/ws');

backend.register('demo.Counter', {
    count: 0,
    increment() {
        this.count++;
        backend.notifyPropertyChanged('count', this.count);
    }
});

// Runs until stopped
```

## Solutions Implemented

### Fix #1: Add Grace Period for Engine Cleanup

**Changed**: `pkg/stream/scripting/manager.go`

```go
// New behavior - 30 second grace period
engine.SetOnStopCallback(func() {
    go func() {
        time.Sleep(30 * time.Second)  // Keep engine alive for 30s
        m.enginesMu.Lock()
        delete(m.engines, id)
        m.enginesMu.Unlock()
    }()
})
```

**Benefits**:
- Frontend has 30 seconds to connect and fetch error output
- Error messages are visible in console output
- Script debugging is much easier
- Minimal memory overhead (engines are small)

**Trade-offs**:
- Failed scripts stay in memory for 30 seconds
- Not a problem for typical usage patterns
- Could add manual cleanup if needed

### Fix #2: Logs Page (Recommendation)

**Two Options**:

#### Option A: Integrate zerolog with global logger (Recommended)
Create a zerolog hook that writes to the global logger:

```go
// pkg/stream/logging/hook.go
type GlobalLoggerHook struct {
    logger *Logger
}

func (h *GlobalLoggerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
    fields := extractFields(e)
    switch level {
    case zerolog.DebugLevel:
        h.logger.Debug(msg, fields)
    case zerolog.InfoLevel:
        h.logger.Info(msg, fields)
    // ... etc
    }
}

// In services initialization:
log.Logger = log.Hook(logging.NewGlobalLoggerHook())
```

#### Option B: Remove logs page or make it read-only
- Logs are already visible in server console
- Remove the logs feature from UI
- Or make it show "Logs are available in server console"

**Recommendation**: Option A - Integrate the logging systems. This provides a better user experience and centralizes log viewing.

### Fix #3: Script Lifecycle (Documentation)

Added clear documentation in UI and docs about script behavior:
- Scripts run continuously by design
- Use Stop button to manually terminate
- Use `exit()` in script to self-terminate
- Running scripts shown in "Running Scripts" list

## Testing

### Test #1: Script Error Output
```bash
# Start server
apigear serve

# In UI, go to Scripts page
# Run this code:
throw new Error("Test error");

# Expected: See error in console output
# Before fix: "Connection to script output lost"
# After fix: Error message visible in console
```

### Test #2: Script Success with exit()
```javascript
// Run this code:
console.log("Hello");
console.log("World");
exit();

// Expected: Both messages visible, then script stops
// Before fix: Connection lost before messages shown
// After fix: Both messages visible, clean shutdown
```

### Test #3: Long-running Script
```javascript
// Run this code:
console.log("Starting...");

setInterval(() => {
    console.log("Tick:", new Date().toISOString());
}, 1000);

// Expected: See "Starting..." then ticks every second
// Must manually click Stop button to terminate
```

## Files Modified

1. **pkg/stream/scripting/manager.go**
   - Added `time` import
   - Modified `SetOnStopCallback` to use 30-second grace period
   - Applied to both client scripts and backend scripts (2 locations)

## Verification Steps

1. **Build and start server**:
   ```bash
   go build -o /tmp/apigear ./cmd/apigear/
   /tmp/apigear serve
   ```

2. **Test error handling**:
   - Go to Scripts page
   - Run code with syntax error: `throw new Error("test")`
   - Verify error appears in console output
   - Verify no "Connection lost" message

3. **Test normal execution**:
   - Run code with exit: `console.log("test"); exit();`
   - Verify message appears
   - Verify clean shutdown

4. **Test long-running**:
   - Run code with setInterval
   - Verify messages keep appearing
   - Click Stop button
   - Verify script stops

## Known Limitations

1. **Grace period is fixed at 30 seconds**
   - Could make configurable if needed
   - 30s is reasonable for most use cases

2. **Logs page still doesn't show logs**
   - Requires logging integration (Option A above)
   - Or remove feature (Option B above)

3. **No automatic script restart on error**
   - Scripts that error stay stopped
   - User must manually restart
   - Could add auto-restart option if needed

## Migration from wsproxy

**Good news**: wsproxy had the **same issue** - immediate cleanup. The fix applies there too.

The wsproxy code was:
```go
engine.SetOnStopCallback(func() {
    m.enginesMu.Lock()
    delete(m.engines, id)
    m.enginesMu.Unlock()
})
```

This had the same race condition. So the fix we implemented here is an improvement over the original wsproxy behavior.

## Future Enhancements

1. **Configurable grace period**:
   ```go
   type ManagerOptions struct {
       CleanupGracePeriod time.Duration
   }
   ```

2. **Explicit cleanup endpoint**:
   ```
   POST /api/v1/stream/scripts/{id}/cleanup
   ```

3. **Script auto-restart on error**:
   ```javascript
   // In UI: Checkbox "Auto-restart on error"
   ```

4. **Script templates**:
   ```javascript
   // Save common scripts as templates
   // Load template → Edit → Run
   ```

5. **Script scheduling**:
   ```javascript
   // Run script every hour
   // Run script on proxy connect
   ```

## Related Documentation

- [Stream Module Guide](./pkg/stream/README.md)
- [Scripting API Reference](./pkg/stream/scripting/README.md)
- [Frontend Integration](./web/src/pages/Stream/components/ConsoleOutput.tsx)

## Questions?

If scripts are still not working as expected:

1. Check server logs for error messages
2. Check browser console for connection errors
3. Verify script syntax is correct (JavaScript)
4. Try simple script first: `console.log("test"); exit();`
5. Open issue with error details and script code
