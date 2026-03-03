package core

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// ConnectOptions holds options for establishing a WebSocket connection.
type ConnectOptions struct {
	// Header specifies additional HTTP headers to send in the handshake request
	Header http.Header

	// Dialer is the websocket.Dialer to use. If nil, a default dialer is used.
	Dialer *websocket.Dialer
}

// LifecycleManager handles WebSocket connection lifecycle including auto-reconnect.
type LifecycleManager struct {
	// BaseDelay is the initial delay before the first reconnection attempt.
	// Default: 500ms
	BaseDelay time.Duration

	// MaxDelay is the maximum delay between reconnection attempts.
	// Default: 4s
	MaxDelay time.Duration

	// MaxRetries is the maximum number of reconnection attempts. 0 means unlimited.
	// Default: 0 (unlimited)
	MaxRetries int

	// AutoReconnect enables automatic reconnection on connection loss.
	// Default: false
	AutoReconnect bool
}

// DefaultLifecycleManager returns a LifecycleManager with sensible defaults.
func DefaultLifecycleManager() *LifecycleManager {
	return &LifecycleManager{
		BaseDelay:     500 * time.Millisecond,
		MaxDelay:      4 * time.Second,
		MaxRetries:    0, // unlimited
		AutoReconnect: false,
	}
}

// Connect establishes a WebSocket connection with retry logic.
// If AutoReconnect is enabled, it will retry on failure with exponential backoff.
// Returns the established connection or an error if connection fails.
func (m *LifecycleManager) Connect(ctx context.Context, url string, opts ConnectOptions) (Connection, error) {
	if !m.AutoReconnect {
		// Single connection attempt
		return m.connectOnce(ctx, url, opts)
	}

	// Retry loop with exponential backoff
	var retryCount int
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		conn, err := m.connectOnce(ctx, url, opts)
		if err == nil {
			return conn, nil
		}

		retryCount++
		if m.MaxRetries > 0 && retryCount >= m.MaxRetries {
			return nil, fmt.Errorf("max retries (%d) exceeded: %w", m.MaxRetries, err)
		}

		log.Warn().
			Err(err).
			Str("url", url).
			Int("retry", retryCount).
			Msg("Connection failed, will retry")

		// Calculate exponential backoff with cap
		delay := m.BaseDelay * time.Duration(1<<min(retryCount-1, 6))
		if delay > m.MaxDelay {
			delay = m.MaxDelay
		}

		select {
		case <-time.After(delay):
			// Continue to next retry
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// connectOnce attempts a single connection without retry.
func (m *LifecycleManager) connectOnce(ctx context.Context, url string, opts ConnectOptions) (Connection, error) {
	dialer := opts.Dialer
	if dialer == nil {
		dialer = websocket.DefaultDialer
	}

	conn, _, err := dialer.DialContext(ctx, url, opts.Header)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", url, err)
	}

	// Generate a unique ID for this connection
	id := fmt.Sprintf("%s-%d", url, time.Now().UnixNano())
	return NewConnection(conn, id), nil
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
