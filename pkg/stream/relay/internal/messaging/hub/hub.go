package hub

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

// HubOptions configures a Hub instance.
type HubOptions struct {
	// BufferSize is the capacity of the ring buffer for message history.
	// Default: 1000
	BufferSize int

	// PublishBufferSize is the capacity of the async publish channel.
	// Default: 10000
	PublishBufferSize int

	// SubscriberBufferSize is the capacity of each subscriber's channel.
	// Default: 100
	SubscriberBufferSize int
}

// DefaultHubOptions returns the default hub configuration.
func DefaultHubOptions() HubOptions {
	return HubOptions{
		BufferSize:           1000,
		PublishBufferSize:    10000,
		SubscriberBufferSize: 100,
	}
}

// Hub is a generic pub/sub hub that broadcasts items to subscribers.
// It maintains a ring buffer of recent items for history.
// Publishing is async via a buffered channel to avoid blocking.
type Hub[T any] struct {
	buffer      *RingBuffer[T]
	subscribers map[string]chan T
	subBufSize  int
	mu          sync.RWMutex

	// Async publishing
	publishCh chan T
	done      chan struct{}
	wg        sync.WaitGroup
}

// NewHub creates a new hub with the given options.
func NewHub[T any](opts HubOptions) *Hub[T] {
	if opts.BufferSize <= 0 {
		opts.BufferSize = 1000
	}
	if opts.PublishBufferSize <= 0 {
		opts.PublishBufferSize = 10000
	}
	if opts.SubscriberBufferSize <= 0 {
		opts.SubscriberBufferSize = 100
	}

	h := &Hub[T]{
		buffer:      NewRingBuffer[T](opts.BufferSize),
		subscribers: make(map[string]chan T),
		subBufSize:  opts.SubscriberBufferSize,
		publishCh:   make(chan T, opts.PublishBufferSize),
		done:        make(chan struct{}),
	}
	h.start()
	return h
}

// Subscribe creates a new subscription and returns an ID and channel.
// The channel receives published items. Use Unsubscribe to clean up.
func (h *Hub[T]) Subscribe() (string, <-chan T) {
	h.mu.Lock()
	defer h.mu.Unlock()

	id := generateID()
	ch := make(chan T, h.subBufSize)
	h.subscribers[id] = ch
	return id, ch
}

// Unsubscribe removes a subscription by ID and closes its channel.
func (h *Hub[T]) Unsubscribe(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if ch, ok := h.subscribers[id]; ok {
		delete(h.subscribers, id)
		close(ch)
	}
}

// Publish queues an item for async processing.
// Non-blocking: drops the item if the publish channel is full.
func (h *Hub[T]) Publish(item T) {
	select {
	case h.publishCh <- item:
	default:
		// Drop item if publish channel is full (system overloaded)
	}
}

// Entries returns all buffered items in chronological order.
func (h *Hub[T]) Entries() []T {
	return h.buffer.Entries()
}

// Clear removes all items from the buffer.
func (h *Hub[T]) Clear() {
	h.buffer.Clear()
}

// Len returns the number of items currently in the buffer.
func (h *Hub[T]) Len() int {
	return h.buffer.Len()
}

// Stop gracefully stops the hub's background goroutine.
// It drains remaining messages before returning.
func (h *Hub[T]) Stop() {
	close(h.done)
	h.wg.Wait()
}

// start begins the background goroutine that processes published items.
func (h *Hub[T]) start() {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		for {
			select {
			case item := <-h.publishCh:
				h.processItem(item)
			case <-h.done:
				// Drain remaining items before exiting
				for {
					select {
					case item := <-h.publishCh:
						h.processItem(item)
					default:
						return
					}
				}
			}
		}
	}()
}

// processItem handles the actual publishing work (ring buffer + subscribers).
func (h *Hub[T]) processItem(item T) {
	// Add to ring buffer
	h.buffer.Push(item)

	// Send to subscribers
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, ch := range h.subscribers {
		select {
		case ch <- item:
		default:
			// Drop item if channel is full (subscriber is slow)
		}
	}
}

// SubscriberCount returns the number of active subscribers.
func (h *Hub[T]) SubscriberCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.subscribers)
}

// generateID creates a random hex string for subscription IDs.
func generateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
