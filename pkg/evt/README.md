# evt

Event bus abstraction with stub implementation (NATS removed).

## Current Status

**NATS dependencies have been removed.** The event bus is now a stub implementation that provides interface compatibility but no actual message distribution.

## Purpose

The `evt` package provides an event bus abstraction for publish/subscribe and request/response patterns. Previously built on NATS, it now uses a stub implementation that:

- ✅ Maintains API compatibility via `IEventBus` interface
- ✅ Logs warnings when event bus methods are called
- ❌ Does not distribute events across processes
- ❌ Does not provide request/response functionality
- ❌ Does not execute registered handlers or middleware

## Current Functionality

**Event Types:**
- `Event` - Message struct with Kind, Value, Error, and Meta fields
- `NewEvent()`, `NewErrorEvent()` - Event constructors

**Stub Event Bus:**
- `NewStubEventBus()` - Creates a no-op event bus that implements `IEventBus`
- `Publish()` - Logs warning, does nothing
- `Request()` - Logs warning, returns error event
- `Register()` - Logs warning, stores handler but never calls it
- `Use()` - Logs warning, stores middleware but never calls it
- `Close()` - Silent no-op

## What No Longer Works

- ❌ Distributed event routing via NATS
- ❌ Event publishing to remote subscribers
- ❌ Request/response pattern with timeouts
- ❌ Handler execution for registered event types
- ❌ Middleware processing
- ❌ JetStream persistent storage

## Re-integrating NATS

To restore NATS functionality:

1. **Add dependencies to go.mod:**
   ```bash
   go get github.com/nats-io/nats.go
   go get github.com/nats-io/nats-server/v2
   ```

2. **Restore implementation files from git history:**
   ```bash
   # Find the commit where NATS was removed
   git log --oneline --all --full-history -- pkg/evt/nats.go

   # Restore the file (replace COMMIT_HASH)
   git show COMMIT_HASH:pkg/evt/nats.go > pkg/evt/nats.go
   git show COMMIT_HASH:pkg/evt/nats_test.go > pkg/evt/nats_test.go
   ```

3. **Restore NATS server in pkg/net:**
   ```bash
   git show COMMIT_HASH:pkg/net/nats.server.go > pkg/net/nats.server.go
   ```

4. **Update NetworkManager (pkg/net/manager.go):**
   - Add NATS configuration options to `Options` struct
   - Add `natsServer` and `nc` fields to `NetworkManager`
   - Restore `StartNATS()`, `StopNATS()`, `NatsConnection()` methods
   - Update `Start()` to launch NATS server
   - Update `EnableMonitor()` to pass NATS connection

5. **Update monitor handler (pkg/net/http.monitor.go):**
   - Add `*nats.Conn` parameter to `MonitorRequestHandler()`
   - Restore NATS publishing code

6. **Replace stub usage:**
   - Find code using `NewStubEventBus()` and replace with `NewNatsEventBus()`

7. **Test:**
   ```bash
   go test ./pkg/evt/...
   go test ./pkg/net/...
   go build ./cmd/apigear
   ```

## Dependencies

This package has no dependencies on other `pkg/` packages.
