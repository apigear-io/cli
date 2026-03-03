// Package core provides low-level WebSocket connection infrastructure.
//
// This package includes thread-safe connection wrappers, connection pooling,
// and lifecycle management with auto-reconnect capabilities.
package core

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Connection represents a thread-safe WebSocket connection.
// All implementations must be safe for concurrent use.
type Connection interface {
	// ReadMessage reads the next message from the connection.
	ReadMessage() (messageType int, data []byte, err error)

	// WriteMessage writes a message to the connection with mutex protection.
	// Returns ErrCloseSent if the connection is already closed.
	WriteMessage(messageType int, data []byte) error

	// Close closes the connection exactly once.
	// Subsequent calls return nil.
	Close() error

	// Done returns a channel that is closed when the connection closes.
	// This can be used with select to wait for connection closure.
	Done() <-chan struct{}

	// ID returns a unique identifier for this connection.
	ID() string
}

// websocketConnection provides thread-safe operations on a WebSocket connection.
// It ensures connections are closed exactly once and protects writes from races.
type websocketConnection struct {
	conn      *websocket.Conn
	id        string
	writeMu   sync.Mutex
	closeOnce sync.Once
	closed    chan struct{}
}

// NewConnection creates a new thread-safe WebSocket connection wrapper.
// The id parameter should be a unique identifier for this connection.
func NewConnection(conn *websocket.Conn, id string) Connection {
	return &websocketConnection{
		conn:   conn,
		id:     id,
		closed: make(chan struct{}),
	}
}

// ID returns the unique identifier for this connection.
func (c *websocketConnection) ID() string {
	return c.id
}

// Close closes the connection exactly once, signaling all waiters on Done().
// Subsequent calls to Close return nil.
func (c *websocketConnection) Close() error {
	var err error
	c.closeOnce.Do(func() {
		close(c.closed)
		err = c.conn.Close()
	})
	return err
}

// WriteMessage writes a message with mutex protection.
// Returns ErrCloseSent if the connection is already closed.
func (c *websocketConnection) WriteMessage(messageType int, data []byte) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	select {
	case <-c.closed:
		return websocket.ErrCloseSent
	default:
	}
	return c.conn.WriteMessage(messageType, data)
}

// ReadMessage reads a message from the connection.
// This method is not protected by a mutex as WebSocket connections
// are safe for concurrent reads and writes from different goroutines.
func (c *websocketConnection) ReadMessage() (int, []byte, error) {
	return c.conn.ReadMessage()
}

// Done returns a channel that is closed when the connection is closed.
// This can be used with select to wait for connection closure:
//
//	select {
//	case <-conn.Done():
//	    // connection closed
//	case msg := <-msgChan:
//	    // process message
//	}
func (c *websocketConnection) Done() <-chan struct{} {
	return c.closed
}
