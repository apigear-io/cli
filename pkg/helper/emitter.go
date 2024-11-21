package helper

import (
	"slices"
	"sync"

	"github.com/google/uuid"
)

type EventHandler func(args ...any)

type emitterEntry struct {
	handler EventHandler
	id      string
}

type Emitter struct {
	sync.RWMutex
	entries map[string][]emitterEntry
}

func NewEmitter() *Emitter {
	return &Emitter{
		entries: make(map[string][]emitterEntry),
	}
}

func (e *Emitter) Add(name string, handler EventHandler) func() {
	e.Lock()
	defer e.Unlock()
	id := uuid.New().String()
	e.entries[name] = append(e.entries[name], emitterEntry{
		handler: handler,
		id:      id,
	})
	return func() {
		e.Remove(name, id)
	}
}

func (e *Emitter) Emit(name string, args ...any) {
	e.RLock()
	defer e.RUnlock()
	entry, ok := e.entries[name]
	if !ok {
		return
	}
	for _, entry := range entry {
		entry.handler(args...)
	}
}

func (e *Emitter) Remove(name string, id string) {
	e.Lock()
	defer e.Unlock()
	entry, ok := e.entries[name]
	if !ok {
		return
	}
	e.entries[name] = slices.DeleteFunc(entry, func(e emitterEntry) bool {
		return e.id == id
	})
}
