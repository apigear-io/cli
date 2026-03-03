package relay

import "github.com/apigear-io/cli/pkg/stream/relay/internal/client"

// State represents the connection state of a client.
//
// Clients transition through these states:
//
//	Disconnected → Connecting → Connected
//	      ↑            ↓
//	      ←─── Retrying ←───
//
// Example usage:
//
//	if client.State() == wsrelay.StateConnected {
//	    client.SendRaw(websocket.TextMessage, data)
//	}
type State = client.State

// Status represents the status of a client connection.
//
// Status includes connection state, retry count, and error information.
//
// Example usage:
//
//	status := wsrelay.Status{
//	    Name:       "my-client",
//	    URL:        "ws://localhost:8080/ws",
//	    State:      wsrelay.StateConnected,
//	    RetryCount: 0,
//	}
//
// Fields:
//   - Name: Client identifier
//   - URL: WebSocket URL
//   - State: Current connection state
//   - RetryCount: Number of reconnection attempts
//   - LastError: Most recent error message
//   - ConnectedAt: Unix timestamp of connection (if connected)
type Status = client.Status

// Client connection states.
const (
	StateDisconnected State = client.StateDisconnected // Not connected
	StateConnecting   State = client.StateConnecting   // Connection in progress
	StateConnected    State = client.StateConnected    // Successfully connected
	StateRetrying     State = client.StateRetrying     // Reconnecting after failure
)
