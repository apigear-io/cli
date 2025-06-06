package sim

import "sync"

// hookEntry is a struct that holds the ID and the hook function
type hookEntry[T any] struct {
	id   string
	hook func(v T)
}

// Hook is a generic type for managing hooks
type Hook[T any] struct {
	mu      sync.RWMutex
	entries []hookEntry[T]
}

func NewHook[T any]() *Hook[T] {
	return &Hook[T]{
		entries: []hookEntry[T]{},
	}
}

// Add a new hook and return a function to unregister it
func (h *Hook[T]) Add(hook func(v T)) func() {
	id := nextId()
	h.mu.Lock()
	defer h.mu.Unlock()
	h.entries = append(h.entries, hookEntry[T]{id: id, hook: hook})
	return func() {
		h.Remove(id)
	}
}

// Emit the hook with the value
func (h *Hook[T]) Emit(v T) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, entry := range h.entries {
		entry.hook(v)
	}
}

// Remove the hook from the list
func (h *Hook[T]) Remove(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i, entry := range h.entries {
		if entry.id == id {
			h.entries = append(h.entries[:i], h.entries[i+1:]...)
			break
		}
	}
}

// Clear all hooks
func (h *Hook[T]) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.entries = []hookEntry[T]{}
}

// Count returns the number of registered hooks
func (h *Hook[T]) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.entries)
}
