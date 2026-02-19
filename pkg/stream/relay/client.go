package relay

import "github.com/apigear-io/cli/pkg/stream/relay/internal/client"

// Client is the generic WebSocket client interface.
//
// Implementations handle connection lifecycle, reconnection, and message sending.
// Clients support automatic reconnection with retry logic.
//
// The Client interface defines:
//
//	type Client interface {
//	    Name() string                    // Unique identifier
//	    URL() string                     // WebSocket URL
//	    State() State                    // Current state
//	    Start() error                    // Begin lifecycle
//	    Stop() error                     // Graceful shutdown
//	    Connect() error                  // Single connection attempt
//	    Disconnect()                     // Close current connection
//	    SendRaw(int, []byte) error       // Send WebSocket message
//	}
type Client = client.Client

// ClientRegistry manages a collection of WebSocket clients with thread-safe operations.
//
// Registries track multiple clients and provide lifecycle management.
// All operations are thread-safe.
//
// Example usage:
//
//	registry := wsrelay.NewClientRegistry()
//
//	// Add clients (implement Client interface)
//	err := registry.Add(myClient)
//
//	// Retrieve and manage
//	client, err := registry.Get("client-name")
//	clients := registry.List()
//	names := registry.Names()
//
//	// Stop all clients
//	registry.StopAll()
type ClientRegistry = client.Registry

// EventHub manages status and message broadcasting for clients.
//
// EventHubs track client connection status and broadcast messages to
// subscribers. They maintain a ring buffer of recent messages and support
// multiple concurrent subscribers.
//
// Example usage:
//
//	hub := wsrelay.NewEventHub[string](1000)
//
//	// Track status
//	hub.UpdateStatus(wsrelay.Status{
//	    Name: "client-1",
//	    State: wsrelay.StateConnected,
//	})
//	status := hub.GetStatus("client-1")
//
//	// Subscribe to events
//	statusCh := hub.SubscribeStatus()
//	msgCh := hub.SubscribeMessages()
//
//	// Publish messages
//	hub.PublishMessage("Hello")
type EventHub[M any] = client.EventHub[M]

// NewClientRegistry creates a new client registry.
//
// The registry is empty initially. Add clients with Add().
func NewClientRegistry() *ClientRegistry {
	return client.NewRegistry()
}

// NewEventHub creates a new event hub with the specified message buffer size.
//
// If messageBufferSize is 0, uses the default (1000).
// The buffer stores recent messages for new subscribers.
func NewEventHub[M any](messageBufferSize int) *EventHub[M] {
	return client.NewEventHub[M](messageBufferSize)
}
