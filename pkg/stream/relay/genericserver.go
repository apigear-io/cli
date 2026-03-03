package relay

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var defaultUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

// ServerConnection represents a connected WebSocket client.
type ServerConnection struct {
	id   string
	conn *websocket.Conn
	mu   sync.Mutex
}

// NewServerConnection creates a new server connection wrapper.
func NewServerConnection(id string, conn *websocket.Conn) *ServerConnection {
	return &ServerConnection{
		id:   id,
		conn: conn,
	}
}

// ID returns the connection ID.
func (c *ServerConnection) ID() string {
	return c.id
}

// SendRaw sends a raw message to the client.
func (c *ServerConnection) SendRaw(messageType int, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.WriteMessage(messageType, data)
}

// Close closes the connection.
func (c *ServerConnection) Close() error {
	return c.conn.Close()
}

// ConnectionHandler handles protocol-specific connection logic.
type ConnectionHandler interface {
	// OnConnected is called when a client connects.
	OnConnected(conn *ServerConnection)

	// OnDisconnected is called when a client disconnects.
	OnDisconnected(conn *ServerConnection)

	// OnMessage is called for each received message.
	// Return error to close the connection.
	OnMessage(conn *ServerConnection, messageType int, data []byte) error
}

// GenericServerConfig holds configuration for a generic WebSocket server.
type GenericServerConfig struct {
	// Name is the server identifier
	Name string

	// ListenAddr is the address to listen on (port, host:port, or full ws:// URL)
	ListenAddr string

	// Handler processes protocol-specific messages
	Handler ConnectionHandler

	// Upgrader is the WebSocket upgrader (optional, uses default if nil)
	Upgrader *websocket.Upgrader
}

// GenericServer is a generic WebSocket server.
type GenericServer struct {
	config GenericServerConfig

	// HTTP server
	server   *http.Server
	serverMu sync.RWMutex

	// Connections
	connections   map[string]*ServerConnection
	connectionsMu sync.RWMutex

	// Connection ID counter
	connID atomic.Int64
}

// NewGenericServer creates a new generic WebSocket server.
func NewGenericServer(config GenericServerConfig) *GenericServer {
	if config.Upgrader == nil {
		config.Upgrader = &defaultUpgrader
	}

	return &GenericServer{
		config:      config,
		connections: make(map[string]*ServerConnection),
	}
}

// Name returns the server name.
func (s *GenericServer) Name() string {
	return s.config.Name
}

// Start begins listening for WebSocket connections.
func (s *GenericServer) Start() error {
	serverAddr, wsPath := parseListenAddr(s.config.ListenAddr)

	mux := http.NewServeMux()
	mux.HandleFunc(wsPath, s.handleWebSocket)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	s.serverMu.Lock()
	s.server = server
	s.serverMu.Unlock()

	log.Info().
		Str("addr", serverAddr).
		Str("path", wsPath).
		Str("server", s.config.Name).
		Msg("WebSocket server starting")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop gracefully shuts down the server.
func (s *GenericServer) Stop() error {
	s.serverMu.RLock()
	server := s.server
	s.serverMu.RUnlock()

	if server == nil {
		return nil
	}

	log.Info().Str("server", s.config.Name).Msg("WebSocket server stopping")

	// Close all connections
	s.connectionsMu.Lock()
	for _, conn := range s.connections {
		_ = conn.Close()
	}
	s.connections = make(map[string]*ServerConnection)
	s.connectionsMu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}

// GetConnection returns a connection by ID.
func (s *GenericServer) GetConnection(id string) *ServerConnection {
	s.connectionsMu.RLock()
	defer s.connectionsMu.RUnlock()
	return s.connections[id]
}

// GetAllConnections returns all active connections.
func (s *GenericServer) GetAllConnections() []*ServerConnection {
	s.connectionsMu.RLock()
	defer s.connectionsMu.RUnlock()

	conns := make([]*ServerConnection, 0, len(s.connections))
	for _, conn := range s.connections {
		conns = append(conns, conn)
	}
	return conns
}

// ConnectionCount returns the number of active connections.
func (s *GenericServer) ConnectionCount() int {
	s.connectionsMu.RLock()
	defer s.connectionsMu.RUnlock()
	return len(s.connections)
}

// Broadcast sends a message to all connected clients.
func (s *GenericServer) Broadcast(messageType int, data []byte) {
	s.connectionsMu.RLock()
	defer s.connectionsMu.RUnlock()

	for _, conn := range s.connections {
		if err := conn.SendRaw(messageType, data); err != nil {
			log.Error().
				Err(err).
				Str("conn", conn.ID()).
				Str("server", s.config.Name).
				Msg("Failed to broadcast message")
		}
	}
}

// handleWebSocket handles incoming WebSocket connections.
func (s *GenericServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.config.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Str("server", s.config.Name).Msg("Failed to upgrade connection")
		return
	}

	connID := fmt.Sprintf("conn-%d", s.connID.Add(1))
	serverConn := NewServerConnection(connID, conn)

	s.connectionsMu.Lock()
	s.connections[connID] = serverConn
	s.connectionsMu.Unlock()

	log.Debug().Str("conn", connID).Str("server", s.config.Name).Msg("Client connected")

	// Notify handler
	if s.config.Handler != nil {
		s.config.Handler.OnConnected(serverConn)
	}

	// Handle messages
	s.handleConnection(serverConn)

	// Cleanup
	s.connectionsMu.Lock()
	delete(s.connections, connID)
	s.connectionsMu.Unlock()

	// Notify handler
	if s.config.Handler != nil {
		s.config.Handler.OnDisconnected(serverConn)
	}

	log.Debug().Str("conn", connID).Str("server", s.config.Name).Msg("Client disconnected")
}

// handleConnection processes messages from a client.
func (s *GenericServer) handleConnection(conn *ServerConnection) {
	for {
		messageType, data, err := conn.conn.ReadMessage()
		if err != nil {
			// Only log truly unexpected close errors
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
				log.Error().
					Err(err).
					Str("conn", conn.ID()).
					Str("server", s.config.Name).
					Msg("Read error")
			}
			return
		}

		// Pass to handler
		if s.config.Handler != nil {
			if err := s.config.Handler.OnMessage(conn, messageType, data); err != nil {
				log.Error().
					Err(err).
					Str("conn", conn.ID()).
					Str("server", s.config.Name).
					Msg("Message handler error")
				return
			}
		}
	}
}

// parseListenAddr parses a listen address which can be:
// - Simple port: ":5560" or "5560"
// - Host:port: "localhost:5560"
// - Full WebSocket URL: "ws://localhost:5560/ws"
// Returns the address for http.Server and the path for the WebSocket handler.
func parseListenAddr(addr string) (serverAddr, wsPath string) {
	wsPath = "/ws" // Default path

	// Check if it's a full URL
	if strings.HasPrefix(addr, "ws://") || strings.HasPrefix(addr, "wss://") {
		if u, err := url.Parse(addr); err == nil {
			serverAddr = u.Host
			if u.Path != "" {
				wsPath = u.Path
			}
			return
		}
	}

	// Simple address format
	serverAddr = addr
	return
}
