package sim

import "sync"

type emitterEntry[T any] struct {
	id      string
	handler func(args T)
}

type hookPair[T any] struct {
	key   string
	value T
}

type Emitter[T any] struct {
	mu      sync.RWMutex
	entries map[string][]emitterEntry[T]
	hook    Hook[hookPair[T]]
}

func NewEmitter[T any]() *Emitter[T] {
	return &Emitter[T]{
		entries: make(map[string][]emitterEntry[T]),
	}
}

// Add adds a handler for the given event.
// It returns a function that can be called to remove the handler.
func (e *Emitter[T]) Add(event string, handler func(value T)) func() {
	id := nextId()
	e.mu.Lock()
	defer e.mu.Unlock()
	e.entries[event] = append(e.entries[event], emitterEntry[T]{id: id, handler: handler})
	return func() {
		e.Remove(event, id)
	}
}

func (e *Emitter[T]) Any(handler func(key string, value T)) {
	e.hook.Add(func(h hookPair[T]) {
		handler(h.key, h.value)
	})

}

// Remove removes the handler for the given event.
// If the event has no handlers, it does nothing.
func (e *Emitter[T]) Remove(event string, id string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if handlers, ok := e.entries[event]; ok {
		for i, handler := range handlers {
			if handler.id == id {
				e.entries[event] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
}

// Emit triggers the handlers for the given event.
// It returns an error if any of the handlers return an error.
func (e *Emitter[T]) Emit(event string, value T) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if handlers, ok := e.entries[event]; ok {
		for _, handler := range handlers {
			handler.handler(value)
		}
	}
	e.hook.Emit(hookPair[T]{
		key:   event,
		value: value,
	})
}

// Clear clears all handlers for the given event.
// It returns the number of handlers removed.
func (e *Emitter[T]) Clear(event string) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	if handlers, ok := e.entries[event]; ok {
		count := len(handlers)
		delete(e.entries, event)
		return count
	}
	return 0
}

// ClearAll clears all handlers for all events.
// It returns the number of handlers removed.
func (e *Emitter[T]) ClearAll() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	count := 0
	for event := range e.entries {
		count += len(e.entries[event])
		delete(e.entries, event)
	}
	e.hook.Clear()
	return count
}

// Has checks if there are any handlers for the given event.
// It returns true if there are handlers, false otherwise.
func (e *Emitter[T]) Has(event string) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	_, ok := e.entries[event]
	return ok
}

// Count returns the number of handlers for the given event.
// It returns 0 if there are no handlers.
func (e *Emitter[T]) Count(event string) int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if handlers, ok := e.entries[event]; ok {
		return len(handlers)
	}
	return 0
}

// CountAll returns the total number of handlers for all events.
// It returns 0 if there are no handlers.
func (e *Emitter[T]) CountAll() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	count := 0
	for _, handlers := range e.entries {
		count += len(handlers)
	}
	return count
}
