package js

import (
	"sync"
	"sync/atomic"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/dop251/goja"
)

type emitterEntry struct {
	id uint64
	fn goja.Callable
}

type KWArgs map[string]any

type Emitter struct {
	vm          *goja.Runtime
	handlers    map[string][]emitterEntry
	anyHandlers []emitterEntry
	mu          sync.RWMutex
	nextID      atomic.Uint64
}

func NewEmitter(vm *goja.Runtime) *Emitter {
	return &Emitter{
		vm:          vm,
		handlers:    make(map[string][]emitterEntry),
		anyHandlers: []emitterEntry{},
	}
}

func (e *Emitter) nextHandlerID() uint64 {
	return e.nextID.Add(1)
}

func (e *Emitter) On(event string, fn goja.Callable) goja.Callable {
	id := e.nextHandlerID()

	e.mu.Lock()
	defer e.mu.Unlock()
	handler := emitterEntry{id: id, fn: fn}
	e.handlers[event] = append(e.handlers[event], handler)

	return func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		e.off(event, id)
		return nil, nil
	}
}

func (e *Emitter) off(event string, id uint64) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for i, handler := range e.handlers[event] {
		if handler.id == id {
			e.handlers[event] = append(e.handlers[event][:i], e.handlers[event][i+1:]...)
			break
		}
	}
}

func (e *Emitter) Emit(event string, args ...goja.Value) {
	e.mu.RLock()
	handlers := e.handlers[event]
	e.mu.RUnlock()

	for _, handler := range handlers {
		_, err := handler.fn(nil, args...)
		if err != nil {
			log.Error().Err(err).Msg("failed to call event handler")
		}
	}

	for _, handler := range e.anyHandlers {
		anyArgs := make([]goja.Value, 0, len(args)+1)
		anyArgs = append(anyArgs, e.vm.ToValue(event))
		anyArgs = append(anyArgs, args...)
		_, err := handler.fn(nil, anyArgs...)
		if err != nil {
			log.Error().Err(err).Msg("failed to call event handler")
		}
	}
}

// Off unregisters all handlers for the event
func (e *Emitter) Off(event string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.handlers, event)
}

// Clear clears all handlers
func (e *Emitter) Clear() {
	log.Info().Msg("clearing event handlers")
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers = make(map[string][]emitterEntry)
	e.anyHandlers = []emitterEntry{}
}

// All returns the list of events
func (e *Emitter) All() []string {
	log.Info().Msg("getting all events")
	e.mu.RLock()
	defer e.mu.RUnlock()
	events := make([]string, 0, len(e.handlers))
	for event := range e.handlers {
		events = append(events, event)
	}
	return events
}

// All listens to all events
func (e *Emitter) OnAny(fn goja.Callable) goja.Callable {
	e.mu.Lock()
	defer e.mu.Unlock()
	id := e.nextHandlerID()
	e.anyHandlers = append(e.anyHandlers, emitterEntry{id: id, fn: fn})
	return func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		e.offAny(id)
		return nil, nil
	}
}

func (e *Emitter) offAny(id uint64) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for i, handler := range e.anyHandlers {
		if handler.id == id {
			e.anyHandlers = append(e.anyHandlers[:i], e.anyHandlers[i+1:]...)
			break
		}
	}
}
