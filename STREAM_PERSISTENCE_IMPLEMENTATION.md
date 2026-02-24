# Stream Persistence Implementation

This document describes the file-based persistence implementation for stream proxies and clients.

## Overview

Previously, stream proxies and clients were only stored in memory and lost on server restart. This implementation adds file-based persistence using YAML/JSON configuration files, so that proxies and clients survive server restarts.

## Architecture

### Components

1. **ConfigPersistence** (`pkg/stream/config/persistence.go`)
   - Thread-safe config file operations with `sync.RWMutex`
   - Read-modify-write pattern with `WithConfig()` helper
   - Atomic operations: AddProxy, UpdateProxy, DeleteProxy, AddClient, UpdateClient, DeleteClient

2. **Config Path Management** (`internal/handler/stream_common.go`)
   - `SetStreamConfigPath(path)` - Sets the config file path
   - `getStreamConfigPath()` - Returns config path (defaults to `./stream.yaml`)
   - `GetStreamServices()` - Exported function to access stream services

3. **Handler Updates**
   - `CreateStreamProxy/Client` - Persists to file first, then adds to memory (rollback on failure)
   - `UpdateStreamProxy/Client` - Upsert pattern (update if exists, add if not)
   - `DeleteStreamProxy/Client` - Removes from memory first, then from file (best effort)

### Persistence Strategy

**Create Operation:**
```go
// 1. Persist to config file first
if err := persistence.AddProxy(name, config); err != nil {
    return error
}

// 2. Add to in-memory manager
if err := services.ProxyManager.AddProxy(name, config); err != nil {
    // Rollback: delete from config file
    _ = persistence.DeleteProxy(name)
    return error
}
```

**Update Operation:**
```go
// 1. Upsert to config file (update or add)
err = persistence.UpdateProxy(name, config)
if err != nil {
    // If doesn't exist in config, add it
    if err := persistence.AddProxy(name, config); err != nil {
        return error
    }
}

// 2. Update in-memory (remove and re-add)
services.ProxyManager.RemoveProxy(name)
services.ProxyManager.AddProxy(name, config)
```

**Delete Operation:**
```go
// 1. Remove from in-memory manager first
if err := services.ProxyManager.RemoveProxy(name); err != nil {
    return error
}

// 2. Best-effort delete from config file
_ = persistence.DeleteProxy(name)
// Note: Ignores errors since item is already removed from memory
```

### Startup Behavior

When the `serve` command starts (`pkg/cmd/serve/serve.go`):

1. Sets config path to `./stream.yaml`
2. Checks if config file exists
3. If exists, loads config and initializes services:
   - Loads proxies from config into ProxyManager
   - Loads clients from config into ClientManager
4. Logs summary of loaded proxies/clients

## File Format

The config file uses the existing `config.Config` structure:

```yaml
proxies:
  my-proxy:
    listen: ws://localhost:8081/ws
    backend: ws://backend:9000/ws
    mode: proxy
    enabled: true

clients:
  my-client:
    url: ws://localhost:8081/ws
    interfaces:
      - demo.Counter
      - demo.Calculator
    enabled: true
    auto_reconnect: true
```

## Thread Safety

- All config file operations are protected by `sync.RWMutex` in `ConfigPersistence`
- Read operations use `RLock()`
- Write operations (add/update/delete) use `Lock()`
- Config path access is protected by `sync.RWMutex` in `stream_common.go`

## Error Handling

### Create Operations
- **Config write fails**: Returns 400 error, nothing added to memory
- **Memory add fails**: Rolls back config file change, returns 400 error

### Update Operations
- **Config update fails (doesn't exist)**: Falls back to Add operation
- **Config add fails**: Returns 400 error
- **Memory update fails**: Returns 400 error
- **Proxy is running**: Returns 400 error (must stop first)

### Delete Operations
- **Memory remove fails**: Returns 404 error
- **Config delete fails**: Ignores error (item already removed from memory)

This approach ensures:
- Consistency: Config file is source of truth
- Resilience: Operations handle cases where config may be out of sync
- Usability: Tests can add items to memory without requiring config files

## Testing

### Test Setup
- Each test creates a temporary directory for config files
- `setupTestStreamServices()` sets up temp config path
- Tests can add items directly to memory (config persistence is optional)
- 34 handler tests verify CRUD operations work correctly

### Test Coverage
- ✅ Create proxy/client with persistence
- ✅ Update proxy/client (upsert behavior)
- ✅ Delete proxy/client (best-effort file cleanup)
- ✅ Validation (missing fields, duplicate names)
- ✅ Error cases (not found, invalid data)

## Usage

### Via Web UI
1. Start server: `apigear serve`
2. Open Web UI: `http://localhost:8080`
3. Navigate to Stream → Proxies or Clients
4. Create/update/delete items via UI
5. All changes persisted to `./stream.yaml`
6. Restart server → items are restored

### Via CLI
```bash
# Start with config file
apigear stream config.yaml

# Or start server with web UI
apigear serve
```

### Via API
```bash
# Create proxy
curl -X POST http://localhost:8080/api/v1/stream/proxies \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-proxy",
    "config": {
      "listen": "ws://localhost:8081/ws",
      "backend": "ws://backend:9000/ws",
      "mode": "proxy"
    }
  }'

# Update proxy
curl -X PUT http://localhost:8080/api/v1/stream/proxies/my-proxy \
  -H "Content-Type: application/json" \
  -d '{
    "listen": "ws://localhost:8082/ws",
    "backend": "ws://backend:9001/ws",
    "mode": "proxy"
  }'

# Delete proxy
curl -X DELETE http://localhost:8080/api/v1/stream/proxies/my-proxy
```

## Files Modified

### New Files
- **None** (reused existing `config.Config` and added methods to `ConfigPersistence`)

### Modified Files

1. **pkg/stream/config/persistence.go**
   - Added `AddProxy`, `UpdateProxy`, `DeleteProxy`
   - Added `AddClient`, `UpdateClient`, `DeleteClient`

2. **internal/handler/stream_common.go**
   - Added `SetStreamConfigPath(path)`
   - Added `getStreamConfigPath()`
   - Added `GetStreamServices()` (exported version)

3. **internal/handler/stream_proxies.go**
   - Updated `CreateStreamProxy` to persist to config
   - Updated `UpdateStreamProxy` to persist changes (upsert)
   - Updated `DeleteStreamProxy` to remove from config (best effort)

4. **internal/handler/stream_clients.go**
   - Updated `CreateStreamClient` to persist to config
   - Updated `UpdateStreamClient` to persist changes (upsert)
   - Updated `DeleteStreamClient` to remove from config (best effort)

5. **pkg/cmd/serve/serve.go**
   - Added config path initialization
   - Added config loading on startup
   - Loads existing proxies/clients into services

6. **internal/handler/stream_proxies_test.go**
   - Updated `setupTestStreamServices()` to use temp config files
   - Added imports for `os` and `path/filepath`

## Benefits

1. **Persistence**: Proxies and clients survive server restarts
2. **Version Control**: Config file can be committed to git
3. **Portability**: Easy to share configurations between environments
4. **Atomic Operations**: Thread-safe config file updates
5. **Backward Compatible**: Existing code continues to work
6. **Test-Friendly**: Tests work with or without config files

## Future Improvements

Potential enhancements:

1. **Config Validation**: Validate entire config on startup
2. **Auto-Backup**: Create backup before modifying config
3. **Watch Mode**: Auto-reload config when file changes
4. **Multiple Files**: Support splitting config across multiple files
5. **Database Backend**: Optional database persistence for larger deployments
6. **Import/Export**: CLI commands for config import/export
7. **Config Diff**: Show what changed in config file

## Migration Guide

### From In-Memory Only

No migration needed! The implementation:
- Works with empty config files (created on first use)
- Handles cases where items exist in memory but not in config
- Uses upsert pattern for updates

### Existing Config Files

If you have an existing `stream.yaml` file:
1. Start server: `apigear serve`
2. Server loads proxies/clients from config
3. Any CRUD operations update the config file
4. No manual migration required

## Troubleshooting

### Config File Not Created
- Check write permissions in current directory
- Verify config path with `getStreamConfigPath()`
- Check server logs for errors

### Items Not Restored on Restart
- Verify `stream.yaml` exists in current directory
- Check file format is valid YAML
- Review server startup logs for loading errors

### Permission Errors
- Ensure server has write access to config directory
- Check file permissions: `chmod 644 stream.yaml`

### Concurrent Access
- All operations are thread-safe via `sync.RWMutex`
- Multiple server instances should use different config files
- Consider database backend for truly distributed setups

## Related Documentation

- [Stream Module Guide](./pkg/stream/README.md)
- [Handler Tests](./STREAM_HANDLER_TESTS.md)
- [UI Improvements](./STREAM_PROXY_UI_IMPROVEMENTS.md)
- [Architecture](./ARCHITECTURE.md)
