package helper

import "sync"

// EventHandler callback for an event.
type EventHandler[T any] func(evt T)

// EventEmitter is a typed event emitter.
type EventEmitter[T any] struct {
	sync.RWMutex
	handlers []EventHandler[T]
}

// NewEventEmitter creates a new event emitter.
func NewEventEmitter[T any]() *EventEmitter[T] {
	return &EventEmitter[T]{}
}

// Emit emits an event to all handlers.
func (e *EventEmitter[T]) Emit(evt T) {
	e.RLock()
	defer e.RUnlock()
	for _, handler := range e.handlers {
		handler(evt)
	}
}

// On registers a new event handler.
func (e *EventEmitter[T]) On(handler EventHandler[T]) {
	e.Lock()
	defer e.Unlock()
	e.handlers = append(e.handlers, handler)
}

// Clear removes all handlers.
func (e *EventEmitter[T]) Clear() {
	e.Lock()
	defer e.Unlock()
	e.handlers = nil
}
