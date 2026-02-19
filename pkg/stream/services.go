package stream

import (
	"github.com/apigear-io/cli/pkg/stream/client"
	"github.com/apigear-io/cli/pkg/stream/proxy"
)

// Services is a dependency injection container for all stream components.
type Services struct {
	// Proxy management
	ProxyManager *proxy.Manager

	// Client management
	ClientManager *client.Manager

	// Statistics
	Stats *proxy.Stats

	// TODO: Add more services as we implement them
	// - TraceManager *tracing.Manager
	// - ScriptManager *scripting.Manager
	// - MessageHub *relay.Hub
	// - EventFactory *monitoring.EventFactory
}

// NewServices creates a new services container with all dependencies initialized.
func NewServices() *Services {
	return &Services{
		ProxyManager:  proxy.NewManager(),
		ClientManager: client.NewManager(),
		Stats:         proxy.NewStats(),
	}
}

// Close cleanly shuts down all services.
func (s *Services) Close() error {
	// Stop all proxies
	if s.ProxyManager != nil {
		if err := s.ProxyManager.Close(); err != nil {
			return err
		}
	}

	// Stop all clients
	if s.ClientManager != nil {
		if err := s.ClientManager.Close(); err != nil {
			return err
		}
	}

	return nil
}
