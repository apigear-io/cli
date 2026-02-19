package stream

import (
	"github.com/apigear-io/cli/pkg/stream/client"
	"github.com/apigear-io/cli/pkg/stream/proxy"
	"github.com/apigear-io/cli/pkg/stream/scripting"
)

// Services is a dependency injection container for all stream components.
type Services struct {
	// Proxy management
	ProxyManager *proxy.Manager

	// Client management
	ClientManager *client.Manager

	// Statistics
	Stats *proxy.Stats

	// Script management
	ScriptManager *scripting.Manager

	// Message hub for real-time message streaming (optional)
	MessageHub *MessageHub

	// Event adapter for monitoring integration
	EventAdapter *EventAdapter
}

// MessageHub is a placeholder for message pub/sub functionality.
// TODO: Implement message hub for real-time message streaming.
type MessageHub struct {
	// Implementation pending
}

// Subscribe subscribes to messages for a specific proxy.
func (h *MessageHub) Subscribe(proxyName string) chan ProxyMessage {
	// Stub implementation
	return make(chan ProxyMessage)
}

// Unsubscribe unsubscribes from messages.
func (h *MessageHub) Unsubscribe(ch chan ProxyMessage) {
	// Stub implementation
	close(ch)
}

// ProxyMessage represents a proxy message event.
type ProxyMessage struct {
	Direction string `json:"direction"` // "SEND" or "RECV"
	Data      []byte `json:"data"`      // Raw message data
}

// NewServices creates a new services container with all dependencies initialized.
func NewServices() *Services {
	eventAdapter := NewEventAdapter("stream")

	return &Services{
		ProxyManager:  proxy.NewManager(),
		ClientManager: client.NewManager(),
		Stats:         proxy.NewStats(),
		ScriptManager: scripting.NewManager("./data/scripts", nil),
		MessageHub:    nil, // Optional, can be set later
		EventAdapter:  eventAdapter,
	}
}

// GetDashboardStats returns aggregated statistics for the dashboard.
func (s *Services) GetDashboardStats() DashboardStats {
	// Count proxies by status
	proxies := s.ProxyManager.ListProxies()
	var running, stopped int
	for _, p := range proxies {
		if p.Status == "running" {
			running++
		} else {
			stopped++
		}
	}

	// Count clients by status
	clients := s.ClientManager.ListClients()
	var connected, disconnected int
	for _, c := range clients {
		if c.Status == "connected" {
			connected++
		} else {
			disconnected++
		}
	}

	return DashboardStats{
		Proxies: DashboardProxyStats{
			Total:   len(proxies),
			Running: running,
			Stopped: stopped,
		},
		Clients: DashboardClientStats{
			Total:        len(clients),
			Connected:    connected,
			Disconnected: disconnected,
		},
		Messages: DashboardMessageStats{
			Total: 0, // TODO: implement Stats.GetTotalMessages()
			Rate:  0, // TODO: implement Stats.GetMessageRate()
		},
	}
}

// DashboardStats contains aggregated statistics for the dashboard.
type DashboardStats struct {
	Proxies  DashboardProxyStats   `json:"proxies"`
	Clients  DashboardClientStats  `json:"clients"`
	Messages DashboardMessageStats `json:"messages"`
}

// DashboardProxyStats contains proxy statistics for dashboard.
type DashboardProxyStats struct {
	Total   int `json:"total"`
	Running int `json:"running"`
	Stopped int `json:"stopped"`
}

// DashboardClientStats contains client statistics for dashboard.
type DashboardClientStats struct {
	Total        int `json:"total"`
	Connected    int `json:"connected"`
	Disconnected int `json:"disconnected"`
}

// DashboardMessageStats contains message statistics for dashboard.
type DashboardMessageStats struct {
	Total int64   `json:"total"`
	Rate  float64 `json:"rate"`
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
