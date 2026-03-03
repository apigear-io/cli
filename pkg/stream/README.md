# ApiGear Stream Module

The Stream module provides WebSocket proxy and ObjectLink client management capabilities for ApiGear CLI, enabling real-time message streaming, protocol debugging, and backend connectivity.

## Features

### WebSocket Proxy
- **Multiple Proxy Modes:**
  - `proxy`: Forward WebSocket messages between frontend and backend
  - `echo`: Echo back all received messages (testing/debugging)
  - `backend`: Act as an ObjectLink backend server
  - `inbound-only`: Accept connections without forwarding

- **Real-time Statistics:**
  - Messages sent/received counts
  - Active connection tracking
  - Bytes transferred
  - Uptime monitoring

- **Message Tracing:**
  - JSONL format trace logs with rotation
  - Message timestamps and directions
  - Best-effort ObjectLink message parsing
  - Protocol-level debugging

### ObjectLink Client Management
- Connect to ObjectLink backends via WebSocket
- Interface linking and management
- Auto-reconnect on connection loss
- Connection status tracking
- Error reporting and diagnostics

### Web UI
- Dashboard with real-time statistics
- Proxy creation and management
- Client configuration and control
- Message monitoring
- Start/stop controls

## Architecture

```
pkg/stream/
├── stream.go           # Package entry point
├── proxy/              # WebSocket proxy implementation
│   ├── proxy.go       # Core proxy with 4 modes
│   ├── manager.go     # Multi-proxy lifecycle management
│   ├── stats.go       # Statistics collection
│   └── echo.go        # Echo server implementation
├── client/             # ObjectLink client management
│   ├── manager.go     # Client lifecycle and connections
│   └── adapter.go     # objectlink-core-go integration
├── relay/              # WebSocket infrastructure
│   ├── connection.go  # Connection abstraction
│   ├── client.go      # WebSocket client
│   └── server.go      # WebSocket server
├── protocol/           # ObjectLink message parsing (best-effort)
│   ├── types.go       # Message type constants
│   └── parser.go      # Message parser for logging/UI
├── config/             # Configuration management
│   ├── config.go      # Config types and validation
│   └── loader.go      # YAML/JSON loading
└── services.go         # Dependency injection container
```

## Usage

### CLI Commands

#### Start Stream Server
```bash
# Start with default config
apigear stream

# Start with custom config
apigear stream --config stream.yaml

# Enable trace logging
apigear stream --trace --trace-dir ./traces

# Watch config for changes
apigear stream --watch
```

#### Proxy Management
```bash
# List all proxies
apigear stream proxy list

# Create a proxy
apigear stream proxy create my-proxy \
  --listen ws://localhost:5550/ws \
  --backend ws://localhost:5560/ws \
  --mode proxy

# Start/stop a proxy
apigear stream proxy start my-proxy
apigear stream proxy stop my-proxy

# View proxy statistics
apigear stream proxy stats my-proxy

# Delete a proxy
apigear stream proxy delete my-proxy
```

#### Client Management
```bash
# List all clients
apigear stream client list

# Create a client
apigear stream client create my-client \
  --url ws://localhost:5560/ws \
  --interfaces demo.Counter,demo.Calculator

# Connect/disconnect
apigear stream client connect my-client
apigear stream client disconnect my-client

# View client status
apigear stream client status my-client

# Delete a client
apigear stream client delete my-client
```

#### Echo Server
```bash
# Quick echo server for testing
apigear stream echo --listen :8888
```

### Configuration File

Create `stream.yaml`:

```yaml
verbose: false
trace: true
traceDir: ./data/traces
logFile: ./logs/stream.log
watch: true

traceConfig:
  maxSizeMB: 10
  maxBackups: 5
  maxAgeDays: 7
  compress: true

web:
  listen: :8080

proxies:
  # Forward proxy
  objectlink:
    listen: ws://localhost:5550/ws
    backend: ws://localhost:5560/ws
    mode: proxy

  # Echo server for testing
  echo:
    listen: ws://localhost:5551/ws
    mode: echo

  # ObjectLink backend
  backend:
    listen: ws://localhost:5552/ws
    mode: backend

clients:
  demo:
    url: ws://localhost:5560/ws
    interfaces:
      - demo.Counter
      - demo.Calculator
    enabled: true
    autoReconnect: true
```

### REST API

The Stream module exposes REST endpoints at `/api/v1/stream`:

#### Dashboard
- `GET /stream/dashboard` - Get overall statistics

#### Proxies
- `GET /stream/proxies` - List all proxies
- `POST /stream/proxies` - Create a proxy
- `GET /stream/proxies/{name}` - Get proxy details
- `PUT /stream/proxies/{name}` - Update proxy config
- `DELETE /stream/proxies/{name}` - Delete a proxy
- `POST /stream/proxies/{name}/start` - Start a proxy
- `POST /stream/proxies/{name}/stop` - Stop a proxy
- `GET /stream/proxies/{name}/stats` - Get proxy statistics

#### Clients
- `GET /stream/clients` - List all clients
- `POST /stream/clients` - Create a client
- `GET /stream/clients/{name}` - Get client details
- `PUT /stream/clients/{name}` - Update client config
- `DELETE /stream/clients/{name}` - Delete a client
- `POST /stream/clients/{name}/connect` - Connect client
- `POST /stream/clients/{name}/disconnect` - Disconnect client

### Web UI

Access the web UI at `http://localhost:8080/stream/dashboard`

**Pages:**
- **Dashboard**: Overview of proxies, clients, and message statistics
- **Proxies**: Manage WebSocket proxies, start/stop, view real-time stats
- **Clients**: Manage ObjectLink clients, connect/disconnect, view status

## Integration with ApiGear

### HTTP Server
Stream routes are registered in `internal/handler/router.go`:
```go
RegisterStreamRoutes(router, streamServices)
```

### Event System
Stream events integrate with ApiGear's monitoring system for unified observability.

### Frontend
React pages use TanStack Query with `useSuspenseQuery` for real-time updates:
```typescript
const { data: proxies } = useProxies(); // Auto-refresh every 3 seconds
```

## Development

### Running Tests
```bash
# All tests
go test ./pkg/stream/...

# Specific package
go test ./pkg/stream/proxy
go test ./pkg/stream/client

# With coverage
go test -cover ./pkg/stream/...
```

### Building
```bash
# Build CLI
task build

# Build with stream module
go build -o apigear ./cmd/apigear
```

### Adding New Proxy Modes
1. Add mode constant in `pkg/stream/proxy/types.go`
2. Implement handler in `pkg/stream/proxy/proxy.go`
3. Update validation in `pkg/stream/config/config.go`
4. Add tests in `pkg/stream/proxy/proxy_test.go`

## Examples

### Example 1: Simple Proxy
```bash
# Create config
cat > stream.yaml <<EOF
proxies:
  my-proxy:
    listen: ws://localhost:8080/ws
    backend: ws://backend:9090/ws
    mode: proxy
EOF

# Start server
apigear stream --config stream.yaml
```

### Example 2: Echo Server for Testing
```bash
# Terminal 1: Start echo server
apigear stream echo --listen :8888

# Terminal 2: Test with wscat
wscat -c ws://localhost:8888
> Hello
< Hello
```

### Example 3: ObjectLink Client
```bash
# Create config
cat > stream.yaml <<EOF
clients:
  demo-client:
    url: ws://localhost:5560/ws
    interfaces:
      - demo.Counter
      - demo.Calculator
    enabled: true
    autoReconnect: true
EOF

# Start and connect
apigear stream --config stream.yaml
apigear stream client connect demo-client
```

### Example 4: Programmatic Usage
```go
import (
    "github.com/apigear-io/cli/pkg/stream/proxy"
    "github.com/apigear-io/cli/pkg/stream/relay"
)

// Create services
services := &stream.Services{
    ProxyManager: proxy.NewManager(),
}

// Create and start a proxy
proxyConfig := &proxy.Config{
    Listen:  "ws://localhost:5550/ws",
    Backend: "ws://localhost:5560/ws",
    Mode:    proxy.ModeProxy,
}

p := proxy.New("my-proxy", proxyConfig, services)
err := p.Start()
if err != nil {
    log.Fatal(err)
}
defer p.Stop()
```

## ObjectLink Protocol

The Stream module uses the [objectlink-core-go](https://github.com/apigear-io/objectlink-core-go) library for ObjectLink protocol implementation in clients and backend mode.

**Message Types:**
- `LINK` (10): Link to interface
- `INIT` (11): Initialize state
- `INVOKE` (30): Invoke method
- `SIGNAL` (40): Emit signal
- `PROPERTY_CHANGE` (50): Property changed
- `UNLINK` (11): Unlink from interface

**Message Parsing:**
The `pkg/stream/protocol/` package provides **best-effort parsing** for logging and UI display only. It does NOT implement protocol semantics - use `objectlink-core-go` for that.

## Performance

- Handles thousands of concurrent WebSocket connections
- Low-latency message forwarding (< 1ms typical)
- Configurable trace log rotation to manage disk space
- Real-time statistics with minimal overhead
- Auto-refresh intervals: 1-5 seconds depending on data type

## Troubleshooting

### Proxy won't start
- Check if listen port is already in use: `lsof -i :<port>`
- Verify backend URL is correct and reachable
- Check logs: `apigear stream --verbose`

### Client connection fails
- Verify backend is running and accessible
- Check WebSocket URL format: must start with `ws://` or `wss://`
- Enable trace logging to see connection details

### High memory usage
- Reduce trace log retention: lower `maxBackups` and `maxAgeDays`
- Decrease refresh intervals in web UI
- Limit number of concurrent connections

## Contributing

When adding new features:
1. Update this README
2. Add tests (aim for > 80% coverage)
3. Update API documentation in handler comments
4. Add E2E tests for new UI features
5. Follow existing patterns (see `CLAUDE.md`)

## License

Same as ApiGear CLI main project.
