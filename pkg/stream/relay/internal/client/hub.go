package client

import (
	"sync"

	"github.com/apigear-io/cli/pkg/stream/relay/internal/messaging/hub"
)

const (
	defaultStatusBufferSize  = 100
	defaultMessageBufferSize = 1000
)

// EventHub manages status and message broadcasting for clients.
// It uses generics to support any message type M.
type EventHub[M any] struct {
	// Status pub/sub
	statusSubscribers map[chan Status]struct{}
	statuses          map[string]*Status
	statusMu          sync.RWMutex

	// Message pub/sub with ring buffer
	messageSubscribers map[chan M]struct{}
	messageBuffer      *hub.RingBuffer[M]
	messageMu          sync.RWMutex
}

// NewEventHub creates a new event hub with the specified message buffer size.
// If messageBufferSize is 0, uses the default (1000).
func NewEventHub[M any](messageBufferSize int) *EventHub[M] {
	if messageBufferSize == 0 {
		messageBufferSize = defaultMessageBufferSize
	}

	return &EventHub[M]{
		statusSubscribers:  make(map[chan Status]struct{}),
		statuses:           make(map[string]*Status),
		messageSubscribers: make(map[chan M]struct{}),
		messageBuffer:      hub.NewRingBuffer[M](messageBufferSize),
	}
}

// UpdateStatus updates and broadcasts a client status.
func (h *EventHub[M]) UpdateStatus(status Status) {
	h.statusMu.Lock()
	h.statuses[status.Name] = &status

	// Collect subscribers while holding lock
	subscribers := make([]chan Status, 0, len(h.statusSubscribers))
	for ch := range h.statusSubscribers {
		subscribers = append(subscribers, ch)
	}
	h.statusMu.Unlock()

	// Non-blocking send to subscribers
	for _, ch := range subscribers {
		select {
		case ch <- status:
		default:
			// Drop if subscriber is slow
		}
	}
}

// GetStatus returns the current status of a client.
func (h *EventHub[M]) GetStatus(name string) *Status {
	h.statusMu.RLock()
	defer h.statusMu.RUnlock()
	return h.statuses[name]
}

// GetAllStatuses returns all client statuses.
func (h *EventHub[M]) GetAllStatuses() []Status {
	h.statusMu.RLock()
	defer h.statusMu.RUnlock()

	statuses := make([]Status, 0, len(h.statuses))
	for _, s := range h.statuses {
		statuses = append(statuses, *s)
	}
	return statuses
}

// RemoveStatus removes a client status.
func (h *EventHub[M]) RemoveStatus(name string) {
	h.statusMu.Lock()
	delete(h.statuses, name)
	h.statusMu.Unlock()
}

// SubscribeStatus subscribes to client status updates.
// Returns a channel that will receive status updates.
// Remember to call UnsubscribeStatus when done.
func (h *EventHub[M]) SubscribeStatus() chan Status {
	ch := make(chan Status, defaultStatusBufferSize)
	h.statusMu.Lock()
	h.statusSubscribers[ch] = struct{}{}
	h.statusMu.Unlock()
	return ch
}

// UnsubscribeStatus unsubscribes from client status updates and closes the channel.
func (h *EventHub[M]) UnsubscribeStatus(ch chan Status) {
	h.statusMu.Lock()
	delete(h.statusSubscribers, ch)
	h.statusMu.Unlock()
	close(ch)
}

// PublishMessage broadcasts a client message and adds it to the buffer.
func (h *EventHub[M]) PublishMessage(msg M) {
	// Add to ring buffer
	h.messageBuffer.Push(msg)

	// Collect subscribers
	h.messageMu.RLock()
	subscribers := make([]chan M, 0, len(h.messageSubscribers))
	for ch := range h.messageSubscribers {
		subscribers = append(subscribers, ch)
	}
	h.messageMu.RUnlock()

	// Non-blocking send to subscribers
	for _, ch := range subscribers {
		select {
		case ch <- msg:
		default:
			// Drop if subscriber is slow
		}
	}
}

// GetMessageBuffer returns all buffered messages in chronological order.
func (h *EventHub[M]) GetMessageBuffer() []M {
	return h.messageBuffer.Entries()
}

// ClearMessageBuffer clears the message buffer.
func (h *EventHub[M]) ClearMessageBuffer() {
	h.messageBuffer.Clear()
}

// SubscribeMessages subscribes to client messages.
// Returns a channel that will receive messages.
// Remember to call UnsubscribeMessages when done.
func (h *EventHub[M]) SubscribeMessages() chan M {
	ch := make(chan M, defaultMessageBufferSize)
	h.messageMu.Lock()
	h.messageSubscribers[ch] = struct{}{}
	h.messageMu.Unlock()
	return ch
}

// UnsubscribeMessages unsubscribes from client messages and closes the channel.
func (h *EventHub[M]) UnsubscribeMessages(ch chan M) {
	h.messageMu.Lock()
	delete(h.messageSubscribers, ch)
	h.messageMu.Unlock()
	close(ch)
}
