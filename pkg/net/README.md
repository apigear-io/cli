# net

Unified network management layer for HTTP and NATS infrastructure.

## Purpose

The `net` package provides a central orchestrator for network services, enabling:

- **HTTP Server**: REST API endpoints and WebSocket connections via chi router
- **NATS Server**: Embedded pub/sub messaging server
- **Monitor Integration**: Event broadcasting and subscription

## Key Exports

### Network Manager
- `NetworkManager` - Central orchestrator for all network services
- `NewManager()` - Create new manager
- `Start()`, `Stop()`, `Wait()` - Lifecycle management
- `EnableMonitor()` - Activate monitoring endpoint
- `MonitorEmitter()` - Access event hook emitter

### HTTP Server
- `HTTPServer` - HTTP server wrapper with chi router
- `NewHTTPServer()` - Create HTTP server
- `Router()` - Access chi router for adding handlers

### NATS Server
- `NatsServer` - Embedded NATS server wrapper
- `NewNatsServer()` - Create embedded server
- `ClientURL()`, `Connection()` - Client connectivity

### Utilities
- `NDJSONScanner` - NDJSON stream processor
- `MonitorRequestHandler()` - HTTP handler for monitor events

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Config directory for NATS data |
| `helper` | Hook event system |
| `log` | Logging |
| `mon` | Monitor event types and emitter |
