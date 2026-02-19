// Package client provides ObjectLink client management for stream functionality.
//
// This package uses github.com/apigear-io/objectlink-core-go for the actual
// ObjectLink protocol implementation.
package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/apigear-io/objectlink-core-go/log"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/ws"

	"github.com/apigear-io/cli/pkg/stream/config"
)

// Client represents an ObjectLink client connection.
type Client struct {
	name          string
	url           string
	interfaces    []string
	autoReconnect bool
	enabled       bool

	// ObjectLink core components
	registry   *client.Registry
	node       *client.Node
	conn       *ws.Connection
	ctx        context.Context
	cancelFunc context.CancelFunc

	// State tracking
	mu     sync.RWMutex
	status ConnectionStatus
}

// ConnectionStatus represents the connection status of a client.
type ConnectionStatus string

const (
	StatusDisconnected ConnectionStatus = "disconnected"
	StatusConnecting   ConnectionStatus = "connecting"
	StatusConnected    ConnectionStatus = "connected"
	StatusError        ConnectionStatus = "error"
)

// Info returns basic client information.
type Info struct {
	Name          string           `json:"name"`
	URL           string           `json:"url"`
	Interfaces    []string         `json:"interfaces"`
	Status        ConnectionStatus `json:"status"`
	AutoReconnect bool             `json:"autoReconnect"`
	Enabled       bool             `json:"enabled"`
	LastError     string           `json:"lastError,omitempty"`
}

// NewClient creates a new ObjectLink client.
func NewClient(name, url string, interfaces []string, autoReconnect, enabled bool) *Client {
	ctx, cancel := context.WithCancel(context.Background())

	return &Client{
		name:          name,
		url:           url,
		interfaces:    interfaces,
		autoReconnect: autoReconnect,
		enabled:       enabled,
		registry:      client.NewRegistry(),
		ctx:           ctx,
		cancelFunc:    cancel,
		status:        StatusDisconnected,
	}
}

// Connect establishes a connection to the ObjectLink server.
func (c *Client) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return fmt.Errorf("already connected")
	}

	c.status = StatusConnecting

	// Dial WebSocket connection
	conn, err := ws.Dial(c.ctx, c.url)
	if err != nil {
		c.status = StatusError
		return fmt.Errorf("failed to dial: %w", err)
	}

	c.conn = conn
	c.node = client.NewNode(c.registry)
	c.node.SetOutput(conn)
	conn.SetOutput(c.node)

	// Start connection processing
	go c.processConnection()

	// Link interfaces
	for _, iface := range c.interfaces {
		log.Debug().Msgf("client %s: linking interface %s", c.name, iface)
		c.node.LinkRemoteNode(iface)
	}

	c.status = StatusConnected
	log.Info().Msgf("client %s: connected to %s", c.name, c.url)

	return nil
}

// processConnection handles incoming messages from the WebSocket connection.
func (c *Client) processConnection() {
	defer func() {
		c.mu.Lock()
		c.status = StatusDisconnected
		c.conn = nil
		c.mu.Unlock()

		if c.autoReconnect {
			log.Info().Msgf("client %s: auto-reconnecting", c.name)
			if err := c.Connect(); err != nil {
				log.Error().Err(err).Msgf("client %s: reconnection failed", c.name)
			}
		}
	}()

	// Connection will handle message processing through SetOutput
	<-c.ctx.Done()
}

// Disconnect closes the connection.
func (c *Client) Disconnect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return fmt.Errorf("not connected")
	}

	// Unlink interfaces
	for _, iface := range c.interfaces {
		log.Debug().Msgf("client %s: unlinking interface %s", c.name, iface)
		c.node.UnlinkRemoteNode(iface)
	}

	// Close connection
	if err := c.conn.Close(); err != nil {
		log.Warn().Err(err).Msgf("client %s: error closing connection", c.name)
	}

	c.cancelFunc()
	c.conn = nil
	c.node = nil
	c.status = StatusDisconnected

	log.Info().Msgf("client %s: disconnected", c.name)
	return nil
}

// Status returns the current connection status.
func (c *Client) Status() ConnectionStatus {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.status
}

// Info returns client information.
func (c *Client) Info() Info {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return Info{
		Name:          c.name,
		URL:           c.url,
		Interfaces:    c.interfaces,
		Status:        c.status,
		AutoReconnect: c.autoReconnect,
		Enabled:       c.enabled,
	}
}

// Manager manages multiple ObjectLink clients.
type Manager struct {
	mu      sync.RWMutex
	clients map[string]*Client
}

// NewManager creates a new client manager.
func NewManager() *Manager {
	return &Manager{
		clients: make(map[string]*Client),
	}
}

// AddClient adds a new client to the manager.
func (m *Manager) AddClient(name string, cfg config.ClientConfig) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.clients[name]; exists {
		return fmt.Errorf("client %s already exists", name)
	}

	client := NewClient(name, cfg.URL, cfg.Interfaces, cfg.AutoReconnect, cfg.Enabled)
	m.clients[name] = client

	// Auto-connect if enabled
	if cfg.Enabled {
		if err := client.Connect(); err != nil {
			log.Warn().Err(err).Msgf("failed to connect client %s", name)
		}
	}

	return nil
}

// RemoveClient removes a client from the manager.
func (m *Manager) RemoveClient(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	client, exists := m.clients[name]
	if !exists {
		return fmt.Errorf("client %s not found", name)
	}

	// Disconnect if connected
	if client.Status() == StatusConnected {
		if err := client.Disconnect(); err != nil {
			log.Warn().Err(err).Msgf("error disconnecting client %s", name)
		}
	}

	delete(m.clients, name)
	return nil
}

// GetClient returns a client by name.
func (m *Manager) GetClient(name string) (*Client, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	client, exists := m.clients[name]
	if !exists {
		return nil, fmt.Errorf("client %s not found", name)
	}

	return client, nil
}

// ListClients returns information about all clients.
func (m *Manager) ListClients() []Info {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]Info, 0, len(m.clients))
	for _, client := range m.clients {
		result = append(result, client.Info())
	}

	return result
}

// ConnectClient connects a client by name.
func (m *Manager) ConnectClient(name string) error {
	client, err := m.GetClient(name)
	if err != nil {
		return err
	}

	return client.Connect()
}

// DisconnectClient disconnects a client by name.
func (m *Manager) DisconnectClient(name string) error {
	client, err := m.GetClient(name)
	if err != nil {
		return err
	}

	return client.Disconnect()
}

// LoadFromConfig loads clients from configuration.
func (m *Manager) LoadFromConfig(clients map[string]config.ClientConfig) error {
	for name, cfg := range clients {
		if err := m.AddClient(name, cfg); err != nil {
			log.Warn().Err(err).Msgf("failed to add client %s", name)
		}
	}

	return nil
}

// Close disconnects and removes all clients.
func (m *Manager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for name, client := range m.clients {
		if client.Status() == StatusConnected {
			if err := client.Disconnect(); err != nil {
				log.Warn().Err(err).Msgf("error disconnecting client %s", name)
			}
		}
	}

	m.clients = make(map[string]*Client)
	return nil
}
