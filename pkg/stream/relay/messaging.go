package relay

import (
	"github.com/apigear-io/cli/pkg/stream/relay/internal/messaging/forward"
	"github.com/apigear-io/cli/pkg/stream/relay/internal/messaging/hub"
)

// Hub is a generic pub/sub message hub with ring buffer support.
//
// Hubs broadcast messages to multiple subscribers while maintaining
// a ring buffer of recent messages for history. Publishing is async
// and non-blocking.
//
// Example usage:
//
//	hub := wsrelay.NewHub[string](wsrelay.DefaultHubOptions())
//	defer hub.Stop()
//
//	// Subscribe
//	subID, ch := hub.Subscribe()
//	defer hub.Unsubscribe(subID)
//
//	// Publish (non-blocking)
//	hub.Publish("Hello, World!")
//
//	// Receive
//	msg := <-ch
//
//	// Get history
//	history := hub.Entries()
type Hub[T any] = hub.Hub[T]

// HubOptions holds configuration options for a Hub.
//
//	opts := wsrelay.HubOptions{
//	    BufferSize:           100,   // Ring buffer size
//	    PublishBufferSize:    10000, // Async publish channel
//	    SubscriberBufferSize: 50,    // Per-subscriber buffer
//	}
type HubOptions = hub.HubOptions

// RingBuffer is a thread-safe circular buffer for storing messages.
//
// RingBuffers maintain the N most recent items, overwriting the oldest
// when full. All operations are thread-safe.
//
// Example usage:
//
//	buffer := wsrelay.NewRingBuffer[int](3)
//	buffer.Push(1)
//	buffer.Push(2)
//	buffer.Push(3)
//	buffer.Push(4) // Overwrites 1
//	entries := buffer.Entries() // [2, 3, 4]
type RingBuffer[T any] = hub.RingBuffer[T]

// Forwarder defines the interface for message forwarding strategies.
//
// Forwarders read messages from a source connection and write them to
// a destination connection, with optional delays or throttling.
//
// Three strategies are available:
//   - Direct: No delay, immediate forwarding
//   - Delayed: Fixed delay for all messages
//   - Throttled: Speed scaling (e.g., 0.5 = half speed)
//
// Example usage:
//
//	// Direct forwarding
//	forwarder := wsrelay.NewForwarder(wsrelay.ForwarderOptions{})
//	err := forwarder.Forward(sourceConn, destConn)
//
//	// With 100ms delay
//	forwarder := wsrelay.NewForwarder(wsrelay.ForwarderOptions{
//	    Delay: 100 * time.Millisecond,
//	})
//
//	// Half speed (doubles message gaps)
//	forwarder := wsrelay.NewForwarder(wsrelay.ForwarderOptions{
//	    Speed: 0.5,
//	})
type Forwarder = forward.Forwarder

// ForwarderOptions configures a forwarder.
//
//	opts := wsrelay.ForwarderOptions{
//	    Delay:      100 * time.Millisecond, // Fixed delay (optional)
//	    Speed:      0.5,                     // Speed factor (optional, overrides delay)
//	    BufferSize: 1000,                    // Message queue size
//	    OnMessage: func(msg wsrelay.Message) {
//	        log.Printf("Forwarded: %d bytes", len(msg.Data))
//	    },
//	}
type ForwarderOptions = forward.Options

// Message represents a WebSocket message to be forwarded.
//
//	msg := wsrelay.Message{
//	    Type: websocket.TextMessage,  // or BinaryMessage
//	    Data: []byte("Hello"),
//	}
type Message = forward.Message

// NewHub creates a new Hub with the given options.
//
// Use DefaultHubOptions() for sensible defaults.
func NewHub[T any](opts HubOptions) *Hub[T] {
	return hub.NewHub[T](opts)
}

// NewRingBuffer creates a new ring buffer with the specified capacity.
//
// The buffer will store up to capacity items. When full, the oldest
// item is overwritten.
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return hub.NewRingBuffer[T](capacity)
}

// NewForwarder creates the appropriate forwarder based on options.
//
// If Speed is set (0 < Speed < 1), creates a throttled forwarder.
// If Delay is set, creates a delayed forwarder.
// Otherwise, creates a direct forwarder with no delay.
func NewForwarder(opts ForwarderOptions) Forwarder {
	return forward.NewForwarder(opts)
}

// DefaultHubOptions returns HubOptions with sensible defaults.
//
//	opts := wsrelay.DefaultHubOptions()
//	// BufferSize: 1000, PublishBufferSize: 10000, SubscriberBufferSize: 100
func DefaultHubOptions() HubOptions {
	return hub.DefaultHubOptions()
}
