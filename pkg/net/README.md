# net

Network management layer for HTTP infrastructure (NATS removed).

## Current Status

**NATS dependencies have been removed.** The network manager now only handles HTTP services. Monitor events are received but not broadcast.

## Purpose

The `net` package provides a central orchestrator for network services:

- ✅ **HTTP Server**: REST API endpoints and WebSocket connections via chi router
- ✅ **Monitor Integration**: HTTP endpoint receives events, fires local hooks
- ❌ **NATS Server**: Removed - no embedded pub/sub messaging
- ❌ **Event Broadcasting**: Removed - no distributed event routing

## What Still Works

**Network Manager:**
- `NetworkManager` - Orchestrates HTTP server
- `NewManager()` - Create new manager
- `Start()`, `Stop()`, `Wait()` - Lifecycle management
- `EnableMonitor()` - Activate monitoring HTTP endpoint
- `MonitorEmitter()` - Access local event hook emitter (still functional)
- `GetMonitorAddress()` - Returns HTTP monitor endpoint URL
- `HttpServer()` - Access HTTP server instance

**HTTP Server:**
- `HTTPServer` - HTTP server wrapper with chi router
- `NewHTTPServer()` - Create HTTP server
- `Router()` - Access chi router for adding handlers
- Full HTTP/WebSocket functionality

**Monitor Handler:**
- `MonitorRequestHandler()` - Receives events via HTTP POST
- Events logged with details (source, type, id, subject)
- Local hooks fired via `mon.Emitter.FireHook()` (still works)
- **Does not broadcast** events to remote subscribers

**Utilities:**
- `NDJSONScanner` - NDJSON stream processor

## What No Longer Works

- ❌ NATS server (embedded or external)
- ❌ Event broadcasting via NATS pub/sub
- ❌ Distributed event routing across processes
- ❌ Monitor event subscriptions from other processes
- ❌ `OnMonitorEvent()` method for subscribing to events
- ❌ NATS configuration options (NatsHost, NatsPort, etc.)

## Configuration Changes

**Options struct simplified:**
```go
type Options struct {
    HttpAddr          string  // HTTP server address (default: "localhost:5555")
    HttpDisabled      bool    // Disable HTTP server
    MonitorDisabled   bool    // Disable monitor endpoint
    ObjectAPIDisabled bool    // Disable object API
    Logging           bool    // Enable logging
}
```

**Removed configuration:**
- `NatsHost`, `NatsPort` - No longer needed
- `NatsDisabled`, `NatsListen` - No longer needed
- `NatsLeafURL`, `NatsCredentials` - No longer needed

## Monitor Functionality

The monitor endpoint continues to work with degraded functionality:

**Endpoint:** `POST /monitor/{source}`

**What happens:**
1. ✅ HTTP endpoint receives events
2. ✅ Events validated and processed
3. ✅ Event details logged (source, type, id, subject)
4. ✅ Local hooks fired (`mon.Emitter.FireHook()`)
5. ❌ Events **not broadcast** to remote subscribers
6. ✅ Returns HTTP 200 OK

**Example:**
```bash
# This still works - events are received and logged locally
curl -X POST http://localhost:5555/monitor/my-source \
  -H "Content-Type: application/json" \
  -d '[{"type":"test.event","data":{"foo":"bar"}}]'
```

## Re-integrating NATS

To restore NATS functionality:

1. **Add dependencies to go.mod:**
   ```bash
   go get github.com/nats-io/nats.go
   go get github.com/nats-io/nats-server/v2
   ```

2. **Restore NATS server from git history:**
   ```bash
   git log --oneline --all --full-history -- pkg/net/nats.server.go
   git show COMMIT_HASH:pkg/net/nats.server.go > pkg/net/nats.server.go
   ```

3. **Update NetworkManager (manager.go):**
   - Add import: `github.com/nats-io/nats.go`
   - Add NATS config options to `Options` struct
   - Add `natsServer *NatsServer` and `nc *nats.Conn` to `NetworkManager`
   - Restore methods: `StartNATS()`, `StopNATS()`, `NatsConnection()`, `NatsClientURL()`, `OnMonitorEvent()`
   - Update `Start()` to launch NATS server conditionally
   - Update `Stop()` to stop NATS server

4. **Update monitor handler (http.monitor.go):**
   - Add import: `github.com/nats-io/nats.go`
   - Add `nc *nats.Conn` parameter to `MonitorRequestHandler()`
   - Restore NATS publishing code in event loop
   - Update `EnableMonitor()` to pass NATS connection

5. **Restore event bus:**
   - Follow steps in `pkg/evt/README.md`

6. **Test:**
   ```bash
   go test ./pkg/net/...
   go build ./cmd/apigear
   ```

## Dependencies

| Package | Purpose |
|---------|---------|
| `helper` | Hook event system |
| `log` | Logging |
| `mon` | Monitor event types and emitter |
