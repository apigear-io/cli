// Package client provides generic WebSocket client abstractions.
//
// This package defines protocol-agnostic client interfaces and implementations
// that can be used for any WebSocket-based protocol.
package client

import (
	"context"
	"errors"
)

// State represents the connection state of a client.
type State string

const (
	StateDisconnected State = "disconnected"
	StateConnecting   State = "connecting"
	StateConnected    State = "connected"
	StateRetrying     State = "retrying"
)

// Status represents the status of a client connection.
// This is a generic status type that can be extended with additional fields.
type Status struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	State       State  `json:"state"`
	RetryCount  int    `json:"retryCount"`
	LastError   string `json:"lastError,omitempty"`
	ConnectedAt *int64 `json:"connectedAt,omitempty"`
}

// Client is the generic WebSocket client interface.
// Implementations handle connection lifecycle, reconnection, and message sending.
type Client interface {
	// Name returns the unique name of this client.
	Name() string

	// URL returns the WebSocket URL this client connects to.
	URL() string

	// State returns the current connection state.
	State() State

	// Start begins the client lifecycle (connection + reconnection loop).
	// Returns error if already started or if initial setup fails.
	Start() error

	// Stop gracefully shuts down the client and all goroutines.
	// Safe to call multiple times.
	Stop() error

	// Connect establishes a WebSocket connection (single attempt).
	// Returns error if connection fails.
	Connect() error

	// Disconnect closes the current connection without stopping the client.
	Disconnect()

	// SendRaw sends a raw WebSocket message.
	SendRaw(messageType int, data []byte) error
}

// Common errors
var (
	// ErrClientNotFound is returned when a client is not found in the registry
	ErrClientNotFound = errors.New("client not found")

	// ErrClientAlreadyExists is returned when trying to add a client that already exists
	ErrClientAlreadyExists = errors.New("client already exists")

	// ErrNotConnected is returned when attempting operations that require an active connection
	ErrNotConnected = errors.New("client not connected")

	// ErrAlreadyStarted is returned when trying to start a client that's already running
	ErrAlreadyStarted = errors.New("client already started")
)

// ConnectOptions holds options for establishing a WebSocket connection.
type ConnectOptions struct {
	// AutoReconnect enables automatic reconnection on connection loss
	AutoReconnect bool

	// Context for cancellation
	Context context.Context
}
