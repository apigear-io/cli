package handler

import (
	"sync"

	"github.com/apigear-io/cli/pkg/stream"
)

// streamServices holds the stream services singleton
var (
	streamServices     *stream.Services
	streamServicesOnce sync.Once
	streamServicesMu   sync.RWMutex
	streamConfigPath   string
	streamConfigMu     sync.RWMutex
)

// getStreamServices returns the stream services singleton, initializing if needed
func getStreamServices() *stream.Services {
	streamServicesOnce.Do(func() {
		streamServices = stream.NewServices()
	})
	return streamServices
}

// GetStreamServices returns the stream services singleton (exported version for external packages)
func GetStreamServices() *stream.Services {
	return getStreamServices()
}

// setStreamServices sets a custom stream services instance (for testing)
func setStreamServices(services *stream.Services) {
	streamServicesMu.Lock()
	defer streamServicesMu.Unlock()
	streamServices = services
}

// SetStreamConfigPath sets the config file path for persistence
func SetStreamConfigPath(path string) {
	streamConfigMu.Lock()
	defer streamConfigMu.Unlock()
	streamConfigPath = path
}

// getStreamConfigPath returns the config file path
func getStreamConfigPath() string {
	streamConfigMu.RLock()
	defer streamConfigMu.RUnlock()
	if streamConfigPath == "" {
		return "./stream.yaml" // Default path
	}
	return streamConfigPath
}

// StreamDashboardStats represents dashboard statistics
type StreamDashboardStats struct {
	Proxies struct {
		Total   int `json:"total"`
		Running int `json:"running"`
		Stopped int `json:"stopped"`
	} `json:"proxies"`
	Clients struct {
		Total        int `json:"total"`
		Connected    int `json:"connected"`
		Disconnected int `json:"disconnected"`
	} `json:"clients"`
	Messages struct {
		Total int64   `json:"total"`
		Rate  float64 `json:"rate"` // Messages per second
	} `json:"messages"`
}
