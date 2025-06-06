package helper

import (
	"slices"
	"sync"

	"github.com/google/uuid"
)

type EventHandler[T any] func(data T)

type emitterEntry[T any] struct {
	handler EventHandler[T]
	id      string
}

type Emitter[T any] struct {
	sync.RWMutex
	entries map[string][]emitterEntry[T]
}

func NewEmitter[T any]() *Emitter[T] {
	return &Emitter[T]{
		entries: make(map[string][]emitterEntry[T]),
	}
}

func (e *Emitter[T]) Add(name string, handler EventHandler[T]) func() {
	e.Lock()
	defer e.Unlock()
	id := uuid.New().String()
	e.entries[name] = append(e.entries[name], emitterEntry[T]{
		handler: handler,
		id:      id,
	})
	return func() {
		e.Remove(name, id)
	}
}

func (e *Emitter[T]) Emit(name string, data T) {
	e.RLock()
	defer e.RUnlock()
	entry, ok := e.entries[name]
	if !ok {
		return
	}
	for _, entry := range entry {
		entry.handler(data)
	}
}

func (e *Emitter[T]) Remove(name string, id string) {
	e.Lock()
	defer e.Unlock()
	entry, ok := e.entries[name]
	if !ok {
		return
	}
	e.entries[name] = slices.DeleteFunc(entry, func(e emitterEntry[T]) bool {
		return e.id == id
	})
}
