package client

import (
	"fmt"
	"sync"
)

// Registry manages a collection of WebSocket clients with thread-safe operations.
type Registry struct {
	clients map[string]Client
	mu      sync.RWMutex
}

// NewRegistry creates a new client registry.
func NewRegistry() *Registry {
	return &Registry{
		clients: make(map[string]Client),
	}
}

// Add adds a client to the registry.
// Returns ErrClientAlreadyExists if a client with the same name already exists.
func (r *Registry) Add(client Client) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := client.Name()
	if _, exists := r.clients[name]; exists {
		return fmt.Errorf("%w: %s", ErrClientAlreadyExists, name)
	}

	r.clients[name] = client
	return nil
}

// Get retrieves a client by name.
// Returns ErrClientNotFound if the client doesn't exist.
func (r *Registry) Get(name string) (Client, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	client, ok := r.clients[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrClientNotFound, name)
	}

	return client, nil
}

// Remove removes a client from the registry without stopping it.
// Returns ErrClientNotFound if the client doesn't exist.
func (r *Registry) Remove(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.clients[name]; !exists {
		return fmt.Errorf("%w: %s", ErrClientNotFound, name)
	}

	delete(r.clients, name)
	return nil
}

// List returns a slice of all registered clients.
func (r *Registry) List() []Client {
	r.mu.RLock()
	defer r.mu.RUnlock()

	clients := make([]Client, 0, len(r.clients))
	for _, client := range r.clients {
		clients = append(clients, client)
	}
	return clients
}

// Names returns a slice of all registered client names.
func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.clients))
	for name := range r.clients {
		names = append(names, name)
	}
	return names
}

// Has checks if a client with the given name exists.
func (r *Registry) Has(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.clients[name]
	return exists
}

// Size returns the number of registered clients.
func (r *Registry) Size() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.clients)
}

// Clear removes all clients from the registry without stopping them.
func (r *Registry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clients = make(map[string]Client)
}

// StopAll stops all registered clients and clears the registry.
// Returns the last error encountered, if any.
func (r *Registry) StopAll() error {
	// Get all clients without holding the lock during Stop()
	r.mu.Lock()
	clients := make([]Client, 0, len(r.clients))
	for _, client := range r.clients {
		clients = append(clients, client)
	}
	r.mu.Unlock()

	// Stop all clients
	var lastErr error
	for _, client := range clients {
		if err := client.Stop(); err != nil {
			lastErr = err
		}
	}

	// Clear the registry
	r.Clear()

	return lastErr
}
