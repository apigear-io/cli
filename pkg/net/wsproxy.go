package net

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// MessageMode defines whether a proxy route operates on text or binary frames.
type MessageMode int

const (
	// MessageModeText indicates that proxy messages are plain text frames.
	MessageModeText MessageMode = iota
	// MessageModeBinary indicates that proxy messages are binary frames.
	MessageModeBinary
)

// MessageDirection describes the movement of a proxied message.
type MessageDirection int

const (
	// DirectionClientToUpstream covers messages flowing from connected clients to the upstream service.
	DirectionClientToUpstream MessageDirection = iota
	// DirectionUpstreamToClient covers messages flowing from the upstream service back to the client.
	DirectionUpstreamToClient
)

// RouteConfig configures how requests for a specific path should be proxied.
type RouteConfig struct {
	// Path describes the HTTP route to match, supporting colon-style parameters (e.g. /ws/:id).
	Path string
	// Param identifies the named parameter that maps to Targets keys.
	Param string
	// Targets maps parameter values to upstream WebSocket URLs.
	Targets map[string]string
	// Mode controls whether the route expects text or binary messages.
	Mode MessageMode
}

// MiddlewareFunc allows consumers to inspect and mutate proxied WebSocket messages.
type MiddlewareFunc func(ctx context.Context, msg *ProxyMessage) error

// ProxyMessage carries metadata and payload for middleware inspection.
type ProxyMessage struct {
	Connection *ConnectionInfo
	Direction  MessageDirection
	Type       MessageMode
	Data       []byte
	Drop       bool
}

// ConnectionInfo holds details about a proxied WebSocket session.
type ConnectionInfo struct {
	ID        string
	Route     *RouteConfig
	TargetID  string
	TargetURL string
	Request   *http.Request
}

// ProxyOptions encapsulates WSProxy configuration.
type ProxyOptions struct {
	BasePath          string
	Routes            []RouteConfig
	Dialer            *websocket.Dialer
	Upgrader          *websocket.Upgrader
	ReconnectAttempts int
	ReconnectBackoff  time.Duration
	Middlewares       []MiddlewareFunc
	OnConnect         func(ctx context.Context, info *ConnectionInfo) error
	OnDisconnect      func(ctx context.Context, info *ConnectionInfo, err error)
}

// WSProxy upgrades incoming HTTP requests to WebSockets and bridges them to upstream targets.
type WSProxy struct {
	opts   ProxyOptions
	router chi.Router

	mu sync.RWMutex
	// routes are tracked in opts.Routes
}

// ErrTargetNotConfigured indicates the target value is missing for a given parameter.
var ErrTargetNotConfigured = errors.New("wsproxy: target not configured")

// ErrUnexpectedMessageType indicates a frame type that does not align with the configured MessageMode.
var ErrUnexpectedMessageType = errors.New("wsproxy: unexpected websocket message type")

// NewWSProxy validates the provided options and returns a ready-to-use proxy.
func NewWSProxy(opts ProxyOptions) (*WSProxy, error) {
	if opts.Upgrader == nil {
		opts.Upgrader = &websocket.Upgrader{
			CheckOrigin: func(*http.Request) bool { return true },
		}
	}
	if opts.Dialer == nil {
		opts.Dialer = websocket.DefaultDialer
	}
	if opts.ReconnectAttempts < 1 {
		opts.ReconnectAttempts = 1
	}
	if opts.ReconnectBackoff <= 0 {
		opts.ReconnectBackoff = 500 * time.Millisecond
	}
	opts.BasePath = normalizeBasePath(opts.BasePath)

	existingRoutes := opts.Routes
	existingMiddleware := opts.Middlewares
	opts.Routes = nil
	opts.Middlewares = nil

	proxy := &WSProxy{
		opts:   opts,
		router: chi.NewRouter(),
	}

	for idx := range existingRoutes {
		if err := proxy.AddRoute(existingRoutes[idx]); err != nil {
			return nil, fmt.Errorf("wsproxy: route %d invalid: %w", idx, err)
		}
	}

	for _, mw := range existingMiddleware {
		proxy.Use(mw)
	}

	return proxy, nil
}

// Use appends middleware handlers that can inspect or drop proxied messages.
func (p *WSProxy) Use(mw MiddlewareFunc) {
	if mw == nil {
		return
	}
	p.mu.Lock()
	p.opts.Middlewares = append(p.opts.Middlewares, mw)
	p.mu.Unlock()
}

// AddRoute registers an additional proxy route at runtime.
func (p *WSProxy) AddRoute(route RouteConfig) error {
	if strings.TrimSpace(route.Path) == "" && strings.TrimSpace(p.opts.BasePath) == "" {
		return errors.New("wsproxy: route path cannot be empty when base path is empty")
	}
	if len(route.Targets) == 0 {
		return fmt.Errorf("wsproxy: route %s must define at least one target", route.Path)
	}
	if route.Param == "" && strings.Contains(route.Path, ":") {
		return fmt.Errorf("wsproxy: route %s requires Param to select target", route.Path)
	}
	if route.Mode != MessageModeText && route.Mode != MessageModeBinary {
		route.Mode = MessageModeText
	}
	routeCopy := route
	p.mu.Lock()
	p.opts.Routes = append(p.opts.Routes, routeCopy)
	cfg := &p.opts.Routes[len(p.opts.Routes)-1]
	path := buildRoutePath(p.opts.BasePath, cfg.Path)
	p.router.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.serveRoute(w, r, cfg)
	}))
	p.mu.Unlock()
	return nil
}

// ServeHTTP routes the request to the configured WebSocket route handlers.
func (p *WSProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *WSProxy) serveRoute(w http.ResponseWriter, r *http.Request, route *RouteConfig) {
	targetID := ""
	if route.Param != "" {
		targetID = chi.URLParam(r, route.Param)
	}

	targetID, targetURL, err := resolveTarget(route, targetID)
	if err != nil {
		http.Error(w, ErrTargetNotConfigured.Error(), http.StatusNotFound)
		return
	}

	conn, err := p.opts.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn().Err(err).Msg("websocket upgrade failed")
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Warn().Err(err).Msg("failed to close client websocket")
		}
	}()

	upstream, err := p.dialUpstream(r.Context(), targetURL)
	if err != nil {
		log.Warn().Err(err).Str("target", targetURL).Msg("websocket upstream dial failed")
		_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseTryAgainLater, "upstream unavailable"))
		return
	}
	defer func() {
		if err := upstream.Close(); err != nil {
			log.Warn().Err(err).Str("target", targetURL).Msg("failed to close upstream websocket")
		}
	}()

	connectionID := uuid.NewString()
	info := &ConnectionInfo{
		ID:        connectionID,
		Route:     route,
		TargetID:  targetID,
		TargetURL: targetURL,
		Request:   r,
	}

	if p.opts.OnConnect != nil {
		if err := p.opts.OnConnect(r.Context(), info); err != nil {
			log.Warn().Err(err).Str("connection", connectionID).Msg("wsproxy connect hook rejected client")
			_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "connection rejected"))
			return
		}
	}

	sessionCtx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var (
		wg        sync.WaitGroup
		resultErr error
		errOnce   sync.Once
	)
	recordError := func(err error) {
		if err == nil {
			return
		}
		errOnce.Do(func() {
			resultErr = err
			cancel()
		})
	}

	switch route.Mode {
	case MessageModeText:
		wg.Add(2)
		go forward(sessionCtx, &wg, recordError, p.middlewareRunner, info, conn, upstream, DirectionClientToUpstream, websocket.TextMessage)
		go forward(sessionCtx, &wg, recordError, p.middlewareRunner, info, upstream, conn, DirectionUpstreamToClient, websocket.TextMessage)
	case MessageModeBinary:
		wg.Add(2)
		go forward(sessionCtx, &wg, recordError, p.middlewareRunner, info, conn, upstream, DirectionClientToUpstream, websocket.BinaryMessage)
		go forward(sessionCtx, &wg, recordError, p.middlewareRunner, info, upstream, conn, DirectionUpstreamToClient, websocket.BinaryMessage)
	default:
		recordError(fmt.Errorf("wsproxy: unsupported message mode %d", route.Mode))
	}

	wg.Wait()

	if p.opts.OnDisconnect != nil {
		p.opts.OnDisconnect(r.Context(), info, resultErr)
	}
}

type middlewareRunner func(ctx context.Context, msg *ProxyMessage) error

func (p *WSProxy) middlewareRunner(ctx context.Context, msg *ProxyMessage) error {
	p.mu.RLock()
	middlewares := append([]MiddlewareFunc(nil), p.opts.Middlewares...)
	p.mu.RUnlock()

	for _, mw := range middlewares {
		if mw == nil {
			continue
		}
		if err := mw(ctx, msg); err != nil {
			return err
		}
		if msg.Drop {
			return nil
		}
	}
	return nil
}

func forward(ctx context.Context, wg *sync.WaitGroup, recordErr func(error), run middlewareRunner, info *ConnectionInfo, reader *websocket.Conn, writer *websocket.Conn, direction MessageDirection, expectedType int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		frameType, payload, err := reader.ReadMessage()
		if err != nil {
			recordErr(err)
			return
		}

		if frameType != expectedType {
			switch frameType {
			case websocket.TextMessage, websocket.BinaryMessage:
				recordErr(fmt.Errorf("%w: got %d expected %d", ErrUnexpectedMessageType, frameType, expectedType))
			default:
				// Ignore control frames; gorilla handles ping/pong automatically.
			}
			continue
		}

		msg := &ProxyMessage{
			Connection: info,
			Direction:  direction,
			Type:       modeFromFrame(frameType),
			Data:       payload,
		}

		if run != nil {
			if err := run(ctx, msg); err != nil {
				recordErr(err)
				return
			}
			if msg.Drop {
				continue
			}
		}

		if err := writer.WriteMessage(frameType, msg.Data); err != nil {
			recordErr(err)
			return
		}
	}
}

func (p *WSProxy) dialUpstream(ctx context.Context, target string) (*websocket.Conn, error) {
	var lastErr error
	for attempt := 0; attempt < p.opts.ReconnectAttempts; attempt++ {
		conn, _, err := p.opts.Dialer.DialContext(ctx, target, nil)
		if err == nil {
			return conn, nil
		}
		lastErr = err
		log.Warn().Err(err).Str("target", target).Int("attempt", attempt+1).Msg("wsproxy upstream dial failed")
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(p.opts.ReconnectBackoff):
		}
	}
	return nil, lastErr
}

func modeFromFrame(frameType int) MessageMode {
	switch frameType {
	case websocket.BinaryMessage:
		return MessageModeBinary
	default:
		return MessageModeText
	}
}

func resolveTarget(route *RouteConfig, requestedID string) (string, string, error) {
	if route.Param != "" {
		if requestedID == "" {
			return "", "", ErrTargetNotConfigured
		}
		if url, ok := route.Targets[requestedID]; ok {
			return requestedID, url, nil
		}
		return "", "", ErrTargetNotConfigured
	}

	if requestedID != "" {
		if url, ok := route.Targets[requestedID]; ok {
			return requestedID, url, nil
		}
	}

	if url, ok := route.Targets[""]; ok {
		return "", url, nil
	}

	if len(route.Targets) == 1 {
		for id, url := range route.Targets {
			return id, url, nil
		}
	}

	return "", "", ErrTargetNotConfigured
}

func buildRoutePath(base, path string) string {
	base = normalizeBasePath(base)
	path = strings.TrimSpace(path)

	switch {
	case base == "" && (path == "" || path == "/"):
		return convertColonParams("/")
	case base == "":
		return convertColonParams("/" + strings.TrimPrefix(path, "/"))
	case path == "" || path == "/":
		if base == "" {
			return convertColonParams("/")
		}
		return convertColonParams(base)
	default:
		return convertColonParams(strings.TrimRight(base, "/") + "/" + strings.TrimPrefix(path, "/"))
	}
}

func normalizeBasePath(base string) string {
	base = strings.TrimSpace(base)
	if base == "" {
		return ""
	}
	if !strings.HasPrefix(base, "/") {
		base = "/" + base
	}
	if len(base) > 1 {
		base = strings.TrimRight(base, "/")
		if base == "" {
			base = "/"
		}
	}
	return base
}

func convertColonParams(path string) string {
	var b strings.Builder
	b.Grow(len(path) + 4)
	for i := 0; i < len(path); i++ {
		if path[i] == ':' {
			j := i + 1
			for j < len(path) && path[j] != '/' {
				j++
			}
			if j > i+1 {
				b.WriteByte('{')
				b.WriteString(path[i+1 : j])
				b.WriteByte('}')
				i = j - 1
				continue
			}
		}
		b.WriteByte(path[i])
	}
	result := b.String()
	if result == "" {
		return "/"
	}
	return result
}
