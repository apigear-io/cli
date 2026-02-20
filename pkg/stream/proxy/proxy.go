package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/apigear-io/cli/pkg/stream/config"
	"github.com/apigear-io/cli/pkg/stream/protocol"
	"github.com/apigear-io/cli/pkg/stream/relay"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TraceEntry represents a single trace log entry.
type TraceEntry struct {
	Timestamp int64           `json:"ts"`
	Direction string          `json:"dir"`
	Proxy     string          `json:"proxy"`
	Message   json.RawMessage `json:"msg"`
}

// MessageHubPublisher is an interface for publishing messages to a message hub.
// This avoids import cycles between proxy and stream packages.
type MessageHubPublisher interface {
	Publish(proxyName string, direction string, data []byte, timestamp int64)
}

// Proxy represents a WebSocket proxy instance.
type Proxy struct {
	name       string
	listenAddr string
	backend    string
	mode       Mode
	verbose    bool
	trace      bool
	traceConfig config.TraceConfig

	// HTTP server
	server     *http.Server
	listener   net.Listener
	actualAddr string
	serverMu   sync.Mutex

	// Trace logging
	traceWriter *lumberjack.Logger
	traceMu     sync.Mutex

	// Echo server (for echo mode)
	echoServer *EchoServer

	// Statistics
	stats *ProxyStats

	// Message hub for real-time streaming (optional)
	messageHub MessageHubPublisher

	// Context for lifecycle management
	ctx        context.Context
	cancelFunc context.CancelFunc

	// Status tracking
	statusMu sync.RWMutex
	status   Status
	startTime time.Time

	// Active connections
	activeConns   map[uint64]*activeConnection
	activeConnsMu sync.RWMutex
	connIDCounter atomic.Uint64
}

// activeConnection tracks an active proxy connection.
type activeConnection struct {
	id      uint64
	client  relay.Connection
	backend relay.Connection
}

// NewProxy creates a new proxy instance.
func NewProxy(name, listenAddr, backend string, cfg config.ProxyConfig) *Proxy {
	ctx, cancel := context.WithCancel(context.Background())

	mode := ParseMode(cfg.Mode)

	// Create a standalone stats collector for this proxy
	stats := NewStats()
	proxyStats := stats.GetProxyStats(name)

	return &Proxy{
		name:        name,
		listenAddr:  listenAddr,
		backend:     backend,
		mode:        mode,
		traceConfig: config.DefaultTraceConfig(),
		ctx:         ctx,
		cancelFunc:  cancel,
		status:      StatusStopped,
		stats:       proxyStats,
		activeConns: make(map[uint64]*activeConnection),
	}
}

// Start starts the proxy server.
func (p *Proxy) Start() error {
	p.serverMu.Lock()
	defer p.serverMu.Unlock()

	if p.server != nil {
		return fmt.Errorf("proxy already running")
	}

	p.statusMu.Lock()
	p.status = StatusRunning
	p.startTime = time.Now()
	p.statusMu.Unlock()

	// Initialize trace logging if enabled
	if p.trace {
		p.initTraceLogging()
	}

	// Initialize echo server for echo mode
	if p.mode == ModeEcho {
		p.echoServer = NewEchoServer(p.name, p.stats)
	}

	// Parse listen address
	u, err := url.Parse(p.listenAddr)
	if err != nil {
		return fmt.Errorf("invalid listen address: %w", err)
	}

	// Create listener first to get actual port
	listener, err := net.Listen("tcp", u.Host)
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}
	p.listener = listener
	p.actualAddr = listener.Addr().String()

	// Create HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc(u.Path, p.handleWebSocket)

	p.server = &http.Server{
		Handler: mux,
	}

	log.Info().
		Str("proxy", p.name).
		Str("listen", p.actualAddr).
		Str("backend", p.backend).
		Str("mode", p.mode.String()).
		Msg("proxy started")

	// Start server in background
	go func() {
		if err := p.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Str("proxy", p.name).Msg("proxy server error")
			p.statusMu.Lock()
			p.status = StatusError
			p.statusMu.Unlock()
		}
	}()

	return nil
}

// Stop stops the proxy server.
func (p *Proxy) Stop() error {
	p.serverMu.Lock()
	defer p.serverMu.Unlock()

	if p.server == nil {
		return fmt.Errorf("proxy not running")
	}

	// Cancel context
	p.cancelFunc()

	// Shutdown HTTP server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.server.Shutdown(shutdownCtx); err != nil {
		log.Warn().Err(err).Str("proxy", p.name).Msg("error shutting down proxy")
	}

	// Close listener
	if p.listener != nil {
		p.listener.Close()
		p.listener = nil
	}

	p.server = nil
	p.actualAddr = ""

	// Close trace writer
	if p.traceWriter != nil {
		p.traceMu.Lock()
		if err := p.traceWriter.Close(); err != nil {
			log.Warn().Err(err).Str("proxy", p.name).Msg("error closing trace writer")
		}
		p.traceWriter = nil
		p.traceMu.Unlock()
	}

	p.statusMu.Lock()
	p.status = StatusStopped
	p.statusMu.Unlock()

	log.Info().Str("proxy", p.name).Msg("proxy stopped")

	return nil
}

// handleWebSocket handles incoming WebSocket connections.
func (p *Proxy) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Str("proxy", p.name).Msg("failed to upgrade connection")
		return
	}

	// Wrap in relay.Connection
	connID := p.connIDCounter.Add(1)
	clientConn := relay.NewConnection(conn, fmt.Sprintf("%s-client-%d", p.name, connID))
	defer clientConn.Close()

	// Track connection
	p.stats.RecordConnectionOpened()
	defer p.stats.RecordConnectionClosed()

	log.Debug().
		Str("proxy", p.name).
		Uint64("connID", connID).
		Msg("client connected")

	// Handle based on mode
	switch p.mode {
	case ModeEcho:
		p.handleEcho(clientConn)
	case ModeProxy:
		p.handleProxy(clientConn)
	case ModeInbound:
		p.handleInbound(clientConn)
	case ModeBackend:
		p.handleBackend(clientConn)
	default:
		log.Error().Str("proxy", p.name).Str("mode", p.mode.String()).Msg("unknown proxy mode")
	}

	log.Debug().
		Str("proxy", p.name).
		Uint64("connID", connID).
		Msg("client disconnected")
}

// handleEcho handles echo mode - sends messages back to client.
func (p *Proxy) handleEcho(clientConn relay.Connection) {
	if p.echoServer == nil {
		log.Error().Str("proxy", p.name).Msg("echo server not initialized")
		return
	}

	if err := p.echoServer.Handle(p.ctx, clientConn); err != nil {
		if err != context.Canceled {
			log.Debug().Err(err).Str("proxy", p.name).Msg("echo handler error")
		}
	}
}

// handleProxy handles proxy mode - forwards messages to backend.
func (p *Proxy) handleProxy(clientConn relay.Connection) {
	// Connect to backend
	backendWS, _, err := websocket.DefaultDialer.Dial(p.backend, nil)
	if err != nil {
		log.Error().Err(err).Str("proxy", p.name).Msg("failed to connect to backend")
		return
	}

	backendConn := relay.NewConnection(backendWS, fmt.Sprintf("%s-backend", p.name))
	defer backendConn.Close()

	// Track active connection
	connID := p.connIDCounter.Load()
	p.activeConnsMu.Lock()
	p.activeConns[connID] = &activeConnection{
		id:      connID,
		client:  clientConn,
		backend: backendConn,
	}
	p.activeConnsMu.Unlock()

	defer func() {
		p.activeConnsMu.Lock()
		delete(p.activeConns, connID)
		p.activeConnsMu.Unlock()
	}()

	// Forward messages bidirectionally
	errChan := make(chan error, 2)

	// Client -> Backend
	go func() {
		errChan <- p.forwardMessages(clientConn, backendConn, DirectionSend)
	}()

	// Backend -> Client
	go func() {
		errChan <- p.forwardMessages(backendConn, clientConn, DirectionRecv)
	}()

	// Wait for either direction to error
	err = <-errChan

	if err != nil && err != context.Canceled {
		log.Debug().Err(err).Str("proxy", p.name).Msg("forwarding error")
	}
}

// handleInbound handles inbound-only mode - logs and discards messages.
func (p *Proxy) handleInbound(clientConn relay.Connection) {
	for {
		select {
		case <-p.ctx.Done():
			return
		case <-clientConn.Done():
			return
		default:
		}

		_, data, err := clientConn.ReadMessage()
		if err != nil {
			return
		}

		// Log message
		p.logMessage(DirectionSend, data)
		p.stats.RecordMessageReceived(len(data))
	}
}

// handleBackend handles backend mode - JavaScript backend (to be implemented).
func (p *Proxy) handleBackend(clientConn relay.Connection) {
	log.Warn().Str("proxy", p.name).Msg("backend mode not yet implemented")
	// TODO: Implement JavaScript backend integration
}

// forwardMessages forwards messages from src to dst.
func (p *Proxy) forwardMessages(src, dst relay.Connection, direction Direction) error {
	for {
		select {
		case <-p.ctx.Done():
			return context.Canceled
		case <-src.Done():
			return nil
		case <-dst.Done():
			return nil
		default:
		}

		// Read message from source
		msgType, data, err := src.ReadMessage()
		if err != nil {
			return err
		}

		// Log and record stats
		p.logMessage(direction, data)
		if direction == DirectionSend {
			p.stats.RecordMessageReceived(len(data))
		} else {
			p.stats.RecordMessageSent(len(data))
		}

		// Write to destination
		if err := dst.WriteMessage(msgType, data); err != nil {
			return err
		}
	}
}

// initTraceLogging initializes trace file logging.
func (p *Proxy) initTraceLogging() {
	p.traceMu.Lock()
	defer p.traceMu.Unlock()

	filename := fmt.Sprintf("traces/%s.jsonl", p.name)

	p.traceWriter = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    p.traceConfig.MaxSizeMB,
		MaxBackups: p.traceConfig.MaxBackups,
		MaxAge:     p.traceConfig.MaxAgeDays,
		Compress:   p.traceConfig.Compress,
	}

	log.Info().Str("proxy", p.name).Str("file", filename).Msg("trace logging enabled")
}

// logMessage logs a message to trace file and console.
func (p *Proxy) logMessage(direction Direction, msg []byte) {
	timestamp := time.Now().UnixMilli()

	if p.verbose {
		// Parse message for display
		parsed := protocol.ParseMessage(json.RawMessage(msg))
		log.Debug().
			Str("proxy", p.name).
			Str("dir", direction.String()).
			Str("type", parsed.MsgTypeName).
			Str("symbol", parsed.Symbol).
			Msg("message")
	}

	if p.trace && p.traceWriter != nil {
		entry := TraceEntry{
			Timestamp: timestamp,
			Direction: direction.String(),
			Proxy:     p.name,
			Message:   json.RawMessage(msg),
		}

		data, err := json.Marshal(entry)
		if err != nil {
			log.Warn().Err(err).Msg("failed to marshal trace entry")
			return
		}

		p.traceMu.Lock()
		if p.traceWriter != nil {
			data = append(data, '\n')
			if _, err := p.traceWriter.Write(data); err != nil {
				log.Warn().Err(err).Msg("failed to write trace entry")
			}
		}
		p.traceMu.Unlock()
	}

	// Publish to message hub if available
	if p.messageHub != nil {
		p.messageHub.Publish(p.name, direction.String(), msg, timestamp)
	}
}

// Info returns proxy information and statistics.
func (p *Proxy) Info() Info {
	p.statusMu.RLock()
	status := p.status
	startTime := p.startTime
	p.statusMu.RUnlock()

	info := p.stats.GetInfo()
	info.Name = p.name
	info.Listen = p.listenAddr
	info.Backend = p.backend
	info.Mode = p.mode.String()
	info.Status = status

	if !startTime.IsZero() {
		info.Uptime = int64(time.Since(startTime).Seconds())
	}

	return info
}

// Status returns the current proxy status.
func (p *Proxy) Status() Status {
	p.statusMu.RLock()
	defer p.statusMu.RUnlock()
	return p.status
}

// SetVerbose enables or disables verbose logging.
func (p *Proxy) SetVerbose(enabled bool) {
	p.verbose = enabled
}

// SetTrace enables or disables trace logging.
func (p *Proxy) SetTrace(enabled bool) {
	if enabled && !p.trace {
		p.trace = true
		p.initTraceLogging()
	} else if !enabled && p.trace {
		p.trace = false
		p.traceMu.Lock()
		if p.traceWriter != nil {
			p.traceWriter.Close()
			p.traceWriter = nil
		}
		p.traceMu.Unlock()
	}
}

// GetListenAddr returns the actual listen address without protocol (host:port).
// This is useful when using port 0 to get the actual assigned port.
func (p *Proxy) GetListenAddr() string {
	p.serverMu.Lock()
	defer p.serverMu.Unlock()

	return p.actualAddr
}

// SetMessageHub sets the message hub for real-time message streaming.
func (p *Proxy) SetMessageHub(hub MessageHubPublisher) {
	p.messageHub = hub
}
