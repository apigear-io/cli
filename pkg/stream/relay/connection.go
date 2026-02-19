package relay

import (
	"github.com/apigear-io/cli/pkg/stream/relay/internal/core"
	"github.com/gorilla/websocket"
)

// Connection represents a thread-safe WebSocket connection.
// All implementations must be safe for concurrent use.
//
// Connections are safe for concurrent writes from multiple goroutines
// and can be used with a separate reader goroutine. The Done() channel
// signals when the connection is closed.
//
// Example usage:
//
//	conn := wsrelay.NewConnection(websocketConn, "client-id")
//
//	// Send messages (thread-safe)
//	conn.WriteMessage(websocket.TextMessage, []byte("hello"))
//
//	// Read messages
//	messageType, data, err := conn.ReadMessage()
//
//	// Detect closure
//	select {
//	case <-conn.Done():
//	    log.Println("Connection closed")
//	}
//
//	conn.Close() // Safe to call multiple times
type Connection = core.Connection

// ConnectionPool manages a collection of WebSocket connections with thread-safe operations.
//
// Pools are useful for tracking active connections and broadcasting messages.
// All operations are thread-safe and can be called from multiple goroutines.
//
// Example usage:
//
//	pool := wsrelay.NewConnectionPool()
//
//	// Add connections
//	id := pool.Add(conn)                    // Auto-generated ID
//	pool.AddWithID("custom-id", conn)       // Custom ID
//
//	// Retrieve and manage
//	conn, err := pool.Get("custom-id")
//	ids := pool.List()
//	size := pool.Size()
//
//	// Cleanup
//	pool.Close() // Closes all connections
type ConnectionPool = core.ConnectionPool

// NewConnection creates a new thread-safe WebSocket connection wrapper.
// The id parameter should be a unique identifier for this connection.
//
// The returned Connection is safe for concurrent use:
//   - Multiple goroutines can call WriteMessage simultaneously
//   - One goroutine should handle ReadMessage
//   - Close can be called from any goroutine
func NewConnection(conn *websocket.Conn, id string) Connection {
	return core.NewConnection(conn, id)
}

// NewConnectionPool creates a new connection pool.
//
// The pool is empty initially. Add connections with Add() or AddWithID().
func NewConnectionPool() ConnectionPool {
	return core.NewConnectionPool()
}
