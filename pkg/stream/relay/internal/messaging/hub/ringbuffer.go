package hub

import "sync"

// RingBuffer is a thread-safe circular buffer with fixed capacity.
// When the buffer is full, new items overwrite the oldest ones.
type RingBuffer[T any] struct {
	items    []T
	capacity int
	head     int // next write position
	size     int // current number of items
	mu       sync.RWMutex
}

// NewRingBuffer creates a new ring buffer with the given capacity.
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		capacity = 1
	}
	return &RingBuffer[T]{
		items:    make([]T, capacity),
		capacity: capacity,
	}
}

// Push adds an item to the buffer. If the buffer is full,
// the oldest item is overwritten.
func (r *RingBuffer[T]) Push(item T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items[r.head] = item
	r.head = (r.head + 1) % r.capacity

	if r.size < r.capacity {
		r.size++
	}
}

// Entries returns a copy of all items in the buffer, ordered from oldest to newest.
func (r *RingBuffer[T]) Entries() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.size == 0 {
		return nil
	}

	result := make([]T, r.size)
	if r.size < r.capacity {
		// Buffer not yet full, items start at index 0
		copy(result, r.items[:r.size])
	} else {
		// Buffer is full, oldest item is at head
		// Copy from head to end, then from 0 to head
		firstPart := r.capacity - r.head
		copy(result[:firstPart], r.items[r.head:])
		copy(result[firstPart:], r.items[:r.head])
	}

	return result
}

// Len returns the current number of items in the buffer.
func (r *RingBuffer[T]) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.size
}

// Cap returns the capacity of the buffer.
func (r *RingBuffer[T]) Cap() int {
	return r.capacity
}

// Clear removes all items from the buffer.
func (r *RingBuffer[T]) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	var zero T
	for i := range r.items {
		r.items[i] = zero
	}
	r.head = 0
	r.size = 0
}
