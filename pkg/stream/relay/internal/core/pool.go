package core

import (
	"sync"

	"github.com/google/uuid"
)

// ConnectionPool manages a collection of WebSocket connections with thread-safe operations.
type ConnectionPool interface {
	// Add adds a connection to the pool with an auto-generated ID.
	// Returns the generated connection ID.
	Add(conn Connection) string

	// AddWithID adds a connection to the pool with a specific ID.
	// Returns ErrDuplicateConnection if the ID already exists.
	AddWithID(id string, conn Connection) error

	// Get retrieves a connection by ID.
	// Returns ErrConnectionNotFound if the connection doesn't exist.
	Get(id string) (Connection, error)

	// Remove removes a connection from the pool by ID.
	// The connection is NOT closed by this operation.
	// Returns ErrConnectionNotFound if the connection doesn't exist.
	Remove(id string) error

	// List returns a slice of all connection IDs currently in the pool.
	List() []string

	// Close closes all connections in the pool and clears the pool.
	// After calling Close, all other operations will return ErrPoolClosed.
	Close() error

	// Size returns the current number of connections in the pool.
	Size() int
}

// connectionPool is a thread-safe implementation of ConnectionPool.
type connectionPool struct {
	mu          sync.RWMutex
	connections map[string]Connection
	closed      bool
}

// NewConnectionPool creates a new connection pool.
func NewConnectionPool() ConnectionPool {
	return &connectionPool{
		connections: make(map[string]Connection),
	}
}

// Add adds a connection to the pool with an auto-generated UUID.
func (p *connectionPool) Add(conn Connection) string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ""
	}

	id := uuid.New().String()
	p.connections[id] = conn
	return id
}

// AddWithID adds a connection to the pool with a specific ID.
func (p *connectionPool) AddWithID(id string, conn Connection) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPoolClosed
	}

	if _, exists := p.connections[id]; exists {
		return ErrDuplicateConnection
	}

	p.connections[id] = conn
	return nil
}

// Get retrieves a connection by ID.
func (p *connectionPool) Get(id string) (Connection, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return nil, ErrPoolClosed
	}

	conn, exists := p.connections[id]
	if !exists {
		return nil, ErrConnectionNotFound
	}

	return conn, nil
}

// Remove removes a connection from the pool without closing it.
func (p *connectionPool) Remove(id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPoolClosed
	}

	if _, exists := p.connections[id]; !exists {
		return ErrConnectionNotFound
	}

	delete(p.connections, id)
	return nil
}

// List returns a slice of all connection IDs.
func (p *connectionPool) List() []string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	ids := make([]string, 0, len(p.connections))
	for id := range p.connections {
		ids = append(ids, id)
	}
	return ids
}

// Close closes all connections and clears the pool.
func (p *connectionPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPoolClosed
	}

	var lastErr error
	for id, conn := range p.connections {
		if err := conn.Close(); err != nil {
			lastErr = err
		}
		delete(p.connections, id)
	}

	p.closed = true
	return lastErr
}

// Size returns the current number of connections.
func (p *connectionPool) Size() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.connections)
}
