# sim

JavaScript simulation engine and ObjectLink runtime.

## Purpose

The `sim` package provides a JavaScript-based simulation environment for creating virtual services and clients. It enables:

- JavaScript execution in a managed event loop
- Virtual service objects with properties, methods, and signals
- WebSocket client connections via ObjectLink protocol
- Bidirectional property/method/signal synchronization

## Key Exports

### Core Components
- `Engine` - JavaScript runtime manager with Goja-based event loop
  - `NewEngine()`, `RunScript()`, `RunFunction()`, `RunOnLoop()`
- `World` - Container for services and channels
  - `CreateService()`, `CreateChannel()`
- `Manager` - High-level orchestrator
  - `ScriptRun()`, `ScriptStop()`, `FunctionRun()`, `Start()`

### Service/Client
- `ObjectService` - Service object in simulation
- `ObjectClient` - Client proxy to remote service
- `Channel` - WebSocket connection wrapper

### ObjectLink Infrastructure
- `OlinkServer` / `IOlinkServer` - ObjectLink protocol server
- `OlinkConnector` / `IOlinkConnector` - ObjectLink WebSocket client

### Utilities
- `Emitter[T]` - Generic event emitter
- `Hook[T]` - Generic hook system

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | Utility functions |
| `log` | Logging |
| `mon` | Monitoring events |
| `net` | HTTP router integration |
