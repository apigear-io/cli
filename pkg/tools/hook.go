package tools

import (
	"fmt"
	"slices"
	"sync"

	"github.com/google/uuid"
)

func makeHookId() string {
	return uuid.NewString()
}

type entry[T any] struct {
	h  func(*T)
	id string
}

type Hook[T any] struct {
	rw       sync.RWMutex
	handlers []entry[T]
}

// NewHook creates a new Hook.
func NewHook[T any]() *Hook[T] {
	return &Hook[T]{}
}

// Add adds the handler to the end of the list
func (h *Hook[T]) Add(fn func(*T)) func() {
	h.rw.Lock()
	defer h.rw.Unlock()

	id := makeHookId()
	h.handlers = slices.Insert(h.handlers, len(h.handlers), entry[T]{
		h:  fn,
		id: id,
	})
	return func() {
		h.remove(id)
	}
}

// PreAdd adds the handler to the beginning of the list
func (h *Hook[T]) PreAdd(fn func(*T)) func() {
	h.rw.Lock()
	defer h.rw.Unlock()

	id := makeHookId()
	h.handlers = slices.Insert(h.handlers, 0, entry[T]{
		h:  fn,
		id: id,
	})
	return func() {
		h.remove(id)
	}
}

// remove removes the handler with the specified id
func (h *Hook[T]) remove(id string) {
	h.rw.Lock()
	defer h.rw.Unlock()

	for i, entry := range h.handlers {
		if entry.id == id {
			h.handlers = slices.Delete(h.handlers, i, i+1)
			break
		}
	}
}

// Fire calls all handlers in the list
func (h *Hook[T]) Fire(event *T, oneOf ...func(*T)) {
	h.rw.RLock()

	// make new entries to avoid concurrent modification
	entries := make([]entry[T], 0, len(h.handlers)+len(oneOf))
	entries = append(entries, h.handlers...)

	for i, h := range oneOf {
		entries = append(entries, entry[T]{
			h:  h,
			id: fmt.Sprintf("@%d", i),
		})
	}

	h.rw.RUnlock()

	// call handlers
	for _, entry := range entries {
		entry.h(event)
	}
}

// Connect connects the hook to another hook
func (h *Hook[T]) Connect(other *Hook[T]) func() {
	return other.Add(func(event *T) {
		h.Fire(event)
	})
}

// Clear clears the list of handlers
func (h *Hook[T]) Clear() {
	h.rw.Lock()
	defer h.rw.Unlock()

	h.handlers = nil
}

// Len returns the number of handlers
func (h *Hook[T]) Len() int {
	h.rw.RLock()
	defer h.rw.RUnlock()

	return len(h.handlers)
}
