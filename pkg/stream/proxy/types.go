// Package proxy provides WebSocket proxy functionality with ObjectLink protocol support.
package proxy

import (
	"context"

	"github.com/apigear-io/cli/pkg/stream/relay"
)

// Mode defines the operational mode of a proxy.
type Mode int

const (
	ModeProxy   Mode = iota // Forward messages to backend
	ModeEcho                // Internal echo server
	ModeBackend             // Backend script mode
	ModeInbound             // Inbound-only (no backend)
)

// String returns the string representation of the mode.
func (m Mode) String() string {
	switch m {
	case ModeProxy:
		return "proxy"
	case ModeEcho:
		return "echo"
	case ModeBackend:
		return "backend"
	case ModeInbound:
		return "inbound-only"
	default:
		return "unknown"
	}
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) Mode {
	switch s {
	case "echo":
		return ModeEcho
	case "backend":
		return ModeBackend
	case "inbound", "inbound-only":
		return ModeInbound
	case "proxy":
		fallthrough
	default:
		return ModeProxy
	}
}

// Status represents the proxy status.
type Status string

const (
	StatusStopped Status = "stopped"
	StatusRunning Status = "running"
	StatusError   Status = "error"
)

// Info contains proxy information and statistics.
type Info struct {
	Name              string `json:"name"`
	Listen            string `json:"listen"`
	Backend           string `json:"backend"`
	Mode              string `json:"mode"`
	Status            Status `json:"status"`
	MessagesReceived  int64  `json:"messagesReceived"`
	MessagesSent      int64  `json:"messagesSent"`
	ActiveConnections int    `json:"activeConnections"`
	BytesReceived     int64  `json:"bytesReceived"`
	BytesSent         int64  `json:"bytesSent"`
	Uptime            int64  `json:"uptime"` // seconds
}

// Direction indicates message flow direction.
type Direction int

const (
	DirectionSend Direction = iota // Client to backend
	DirectionRecv                  // Backend to client
)

// String returns the string representation of the direction.
func (d Direction) String() string {
	switch d {
	case DirectionSend:
		return "SEND"
	case DirectionRecv:
		return "RECV"
	default:
		return "UNKNOWN"
	}
}

// Forwarder handles message forwarding between connections.
type Forwarder interface {
	// Forward sets up bidirectional forwarding between source and destination.
	Forward(ctx context.Context, src, dst relay.Connection) error
	// Close releases any resources.
	Close() error
}

// BackendConnector manages backend connection lifecycle.
type BackendConnector interface {
	// Connect establishes a connection to the backend.
	Connect(ctx context.Context) (relay.Connection, error)
	// Close releases resources.
	Close() error
}
