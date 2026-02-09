package events

import (
	"sync"

	"github.com/rs/zerolog/log"
)

// STUB: NATS Removed - Re-integration Point
//
// This is a stub implementation of IEventBus that replaced the NATS-based
// NatsEventBus. All methods are no-ops that log warnings on first use.
//
// Original Functionality Removed:
// - Publish/subscribe messaging via NATS
// - Request/response pattern with 10s timeout
// - Distributed event routing across processes
// - Event handler registration and middleware support
// - JetStream persistent storage
//
// To re-integrate NATS:
// 1. Add NATS dependencies to go.mod:
//    - github.com/nats-io/nats.go
//    - github.com/nats-io/nats-server/v2 (if embedded server needed)
// 2. Restore pkg/evt/nats.go from git history:
//    git show HEAD~N:pkg/evt/nats.go > pkg/evt/nats.go
// 3. Restore pkg/net/nats.server.go if embedded NATS server is needed:
//    git show HEAD~N:pkg/net/nats.server.go > pkg/net/nats.server.go
// 4. Update NetworkManager in pkg/net/manager.go:
//    - Add NATS configuration options back to Options struct
//    - Add natsServer and nc fields back to NetworkManager
//    - Restore StartNATS(), StopNATS(), and related methods
// 5. Update pkg/net/http.monitor.go to accept NATS connection and publish events
// 6. Replace NewStubEventBus() calls with NewNatsEventBus() calls
// 7. Run tests: go test ./pkg/evt/... ./pkg/net/...
//
// Current Behavior:
// - Publish(): Logs warning, does nothing
// - Request(): Logs warning, returns error event
// - Register(): Logs warning, does nothing
// - Use(): Logs warning, does nothing
// - Close(): Silent no-op

// StubEventBus is a no-op implementation of IEventBus
type StubEventBus struct {
	mu          sync.Mutex
	warnedOnce  map[string]bool
	handlers    map[string]HandlerFunc
	middleware  []HandlerFunc
}

// NewStubEventBus creates a new stub event bus
func NewStubEventBus() *StubEventBus {
	return &StubEventBus{
		warnedOnce: make(map[string]bool),
		handlers:   make(map[string]HandlerFunc),
	}
}

// Publish is a no-op that logs a warning on first use
func (s *StubEventBus) Publish(e *Event) error {
	s.logOnce("publish", "Event bus Publish called but NATS is disabled (stub implementation)")
	log.Debug().
		Str("kind", e.Kind).
		Msg("Event publish (no-op, NATS disabled)")
	return nil
}

// Request returns an error event with a warning on first use
func (s *StubEventBus) Request(e *Event) (*Event, error) {
	s.logOnce("request", "Event bus Request called but NATS is disabled (stub implementation)")
	log.Debug().
		Str("kind", e.Kind).
		Msg("Event request (no-op, NATS disabled)")
	return NewErrorEvent(e.Kind, "event bus disabled: NATS not available"), nil
}

// Register is a no-op that logs a warning on first use
func (s *StubEventBus) Register(kind string, fn HandlerFunc) {
	// Log before acquiring lock to avoid deadlock
	s.logOnce("register", "Event bus Register called but NATS is disabled (stub implementation)")

	s.mu.Lock()
	defer s.mu.Unlock()

	log.Debug().
		Str("kind", kind).
		Msg("Event handler registration (no-op, NATS disabled)")

	// Store handler anyway for interface compliance, though it won't be called
	s.handlers[kind] = fn
}

// Use is a no-op that logs a warning on first use
func (s *StubEventBus) Use(mw ...HandlerFunc) {
	// Log before acquiring lock to avoid deadlock
	s.logOnce("use", "Event bus Use (middleware) called but NATS is disabled (stub implementation)")

	s.mu.Lock()
	defer s.mu.Unlock()

	log.Debug().
		Int("count", len(mw)).
		Msg("Event middleware registration (no-op, NATS disabled)")

	// Store middleware anyway for interface compliance, though it won't be called
	s.middleware = append(s.middleware, mw...)
}

// Close is a silent no-op
func (s *StubEventBus) Close() error {
	return nil
}

// logOnce logs a warning message only once per operation type
func (s *StubEventBus) logOnce(operation, message string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.warnedOnce[operation] {
		log.Warn().Msg(message)
		s.warnedOnce[operation] = true
	}
}
