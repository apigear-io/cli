package relay

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// MessageHandler handles protocol-specific message processing.
type MessageHandler interface {
	// OnConnected is called after successful connection.
	// Can be used to send initial messages (e.g., LINK in ObjectLink).
	OnConnected(client *GenericClient)

	// OnDisconnected is called when connection is lost.
	OnDisconnected(client *GenericClient)

	// OnMessage is called for each received message.
	// Return error to close the connection.
	OnMessage(client *GenericClient, messageType int, data []byte) error
}

// GenericClientConfig holds configuration for a generic WebSocket client.
type GenericClientConfig struct {
	// Name is the unique identifier for this client
	Name string

	// URL is the WebSocket URL to connect to
	URL string

	// AutoReconnect enables automatic reconnection on connection loss
	AutoReconnect bool

	// Enabled controls whether the client should start
	Enabled bool

	// Handler processes protocol-specific messages
	Handler MessageHandler

	// Context for cancellation (optional, will create one if nil)
	Context context.Context
}

// GenericClient is a generic WebSocket client with auto-reconnect.
type GenericClient struct {
	config GenericClientConfig

	// Connection state
	state       atomic.Value // State
	conn        *websocket.Conn
	connMu      sync.Mutex
	statusMu    sync.RWMutex // Protects retryCount, lastError, connectedAt
	retryCount  int
	lastError   string
	connectedAt *int64

	// Control
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewGenericClient creates a new generic WebSocket client.
func NewGenericClient(config GenericClientConfig) *GenericClient {
	c := &GenericClient{
		config: config,
	}
	c.state.Store(StateDisconnected)

	return c
}

// Name returns the client name.
func (c *GenericClient) Name() string {
	return c.config.Name
}

// URL returns the WebSocket URL.
func (c *GenericClient) URL() string {
	return c.config.URL
}

// State returns the current connection state.
func (c *GenericClient) State() State {
	return c.state.Load().(State)
}

// Start begins the client lifecycle.
func (c *GenericClient) Start() error {
	if !c.config.Enabled {
		log.Info().Str("client", c.config.Name).Msg("Client is disabled, not starting")
		return nil
	}

	if c.config.Context != nil {
		c.ctx, c.cancel = context.WithCancel(c.config.Context)
	} else {
		c.ctx, c.cancel = context.WithCancel(context.Background())
	}

	c.wg.Add(1)
	go c.connectLoop()

	log.Info().Str("client", c.config.Name).Str("url", c.config.URL).Msg("Client started")
	return nil
}

// Stop gracefully shuts down the client.
func (c *GenericClient) Stop() error {
	log.Debug().Str("client", c.config.Name).Msg("Stop() called")

	// Cancel context first
	if c.cancel != nil {
		c.cancel()
		log.Debug().Str("client", c.config.Name).Msg("Context cancelled")
	}

	// Close connection (unblocks readLoop)
	c.connMu.Lock()
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
		log.Debug().Str("client", c.config.Name).Msg("Connection closed")
	}
	c.connMu.Unlock()

	// Wait for goroutines
	log.Debug().Str("client", c.config.Name).Msg("Waiting for goroutines...")
	c.wg.Wait()
	log.Debug().Str("client", c.config.Name).Msg("Goroutines finished")

	// Reset context so Start() can create a new one
	c.ctx = nil
	c.cancel = nil

	c.updateState(StateDisconnected)
	log.Info().Str("client", c.config.Name).Msg("Client stopped")
	return nil
}

// Connect attempts a single connection.
func (c *GenericClient) Connect() error {
	c.connMu.Lock()
	if c.conn != nil {
		c.connMu.Unlock()
		return nil // Already connected
	}
	c.connMu.Unlock()

	c.updateState(StateConnecting)

	conn, _, err := websocket.DefaultDialer.DialContext(c.ctx, c.config.URL, nil)
	if err != nil {
		c.statusMu.Lock()
		c.lastError = err.Error()
		c.statusMu.Unlock()
		c.updateState(StateDisconnected)
		return fmt.Errorf("failed to connect: %w", err)
	}

	c.connMu.Lock()
	c.conn = conn
	c.connMu.Unlock()

	now := time.Now().UnixMilli()
	c.statusMu.Lock()
	c.connectedAt = &now
	c.statusMu.Unlock()

	c.updateState(StateConnected)
	log.Info().Str("client", c.config.Name).Str("url", c.config.URL).Msg("Connected")

	// Notify handler
	if c.config.Handler != nil {
		c.config.Handler.OnConnected(c)
	}

	return nil
}

// Disconnect closes the current connection.
func (c *GenericClient) Disconnect() {
	c.connMu.Lock()
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	c.connMu.Unlock()

	c.updateState(StateDisconnected)
	log.Info().Str("client", c.config.Name).Msg("Disconnected")

	// Notify handler
	if c.config.Handler != nil {
		c.config.Handler.OnDisconnected(c)
	}
}

// SendRaw sends a raw WebSocket message.
func (c *GenericClient) SendRaw(messageType int, data []byte) error {
	c.connMu.Lock()
	conn := c.conn
	c.connMu.Unlock()

	if conn == nil {
		return fmt.Errorf("not connected")
	}

	return conn.WriteMessage(messageType, data)
}

// GetStatus returns the current client status.
func (c *GenericClient) GetStatus() Status {
	c.statusMu.RLock()
	defer c.statusMu.RUnlock()

	return Status{
		Name:        c.config.Name,
		URL:         c.config.URL,
		State:       c.State(),
		RetryCount:  c.retryCount,
		LastError:   c.lastError,
		ConnectedAt: c.connectedAt,
	}
}

// connectLoop manages connection lifecycle with auto-reconnection.
func (c *GenericClient) connectLoop() {
	defer c.wg.Done()

	baseDelay := 500 * time.Millisecond
	maxDelay := 4 * time.Second

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		if err := c.Connect(); err != nil {
			c.statusMu.Lock()
			c.retryCount++
			retryCount := c.retryCount
			c.statusMu.Unlock()
			log.Warn().Err(err).Str("client", c.config.Name).Int("retry", retryCount).Msg("Connection failed, will retry")

			if !c.config.AutoReconnect {
				return
			}

			c.updateState(StateRetrying)

			// Exponential backoff
			delay := baseDelay * time.Duration(1<<min(retryCount-1, 6))
			if delay > maxDelay {
				delay = maxDelay
			}

			select {
			case <-time.After(delay):
			case <-c.ctx.Done():
				return
			}
			continue
		}

		// Reset retry count on successful connection
		c.statusMu.Lock()
		c.retryCount = 0
		c.statusMu.Unlock()

		// Start read loop (blocks until disconnected)
		c.readLoop()

		// Connection closed
		c.connMu.Lock()
		c.conn = nil
		c.connMu.Unlock()

		c.statusMu.Lock()
		c.connectedAt = nil
		c.statusMu.Unlock()

		// Notify handler
		if c.config.Handler != nil {
			c.config.Handler.OnDisconnected(c)
		}

		if !c.config.AutoReconnect {
			c.updateState(StateDisconnected)
			return
		}

		log.Info().Str("client", c.config.Name).Msg("Connection lost, will reconnect")
	}
}

// readLoop reads messages from the WebSocket.
func (c *GenericClient) readLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		c.connMu.Lock()
		conn := c.conn
		c.connMu.Unlock()

		if conn == nil {
			return
		}

		messageType, data, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
				c.statusMu.Lock()
				c.lastError = err.Error()
				c.statusMu.Unlock()
				log.Error().Err(err).Str("client", c.config.Name).Msg("WebSocket read error")
			}
			return
		}

		// Pass to handler
		if c.config.Handler != nil {
			if err := c.config.Handler.OnMessage(c, messageType, data); err != nil {
				log.Error().Err(err).Str("client", c.config.Name).Msg("Message handler error")
				return
			}
		}
	}
}

// updateState updates the connection state.
func (c *GenericClient) updateState(state State) {
	c.state.Store(state)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
