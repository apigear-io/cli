package relay

import "github.com/apigear-io/cli/pkg/stream/relay/internal/core"

// LifecycleManager handles WebSocket connection lifecycle including auto-reconnect.
//
// LifecycleManagers establish connections with automatic retry logic and
// exponential backoff. Configure retry behavior via struct fields.
//
// Example usage:
//
//	manager := wsrelay.DefaultLifecycleManager()
//	manager.AutoReconnect = true
//	manager.MaxRetries = 5
//
//	ctx := context.Background()
//	conn, err := manager.Connect(ctx, "ws://localhost:8080/ws", wsrelay.ConnectOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer conn.Close()
//
// Configuration:
//
//	type LifecycleManager struct {
//	    BaseDelay     time.Duration  // Initial retry delay (default: 500ms)
//	    MaxDelay      time.Duration  // Max retry delay (default: 4s)
//	    MaxRetries    int            // Max attempts, 0=unlimited (default: 0)
//	    AutoReconnect bool           // Enable retry (default: false)
//	}
type LifecycleManager = core.LifecycleManager

// ConnectOptions holds options for establishing a WebSocket connection.
//
// Example usage:
//
//	opts := wsrelay.ConnectOptions{
//	    Header: http.Header{
//	        "Authorization": []string{"Bearer token"},
//	    },
//	    Dialer: customDialer,
//	}
type ConnectOptions = core.ConnectOptions

// DefaultLifecycleManager returns a LifecycleManager with sensible defaults.
//
// Default configuration:
//   - BaseDelay: 500ms
//   - MaxDelay: 4s
//   - MaxRetries: 0 (unlimited)
//   - AutoReconnect: false
func DefaultLifecycleManager() *LifecycleManager {
	return core.DefaultLifecycleManager()
}
