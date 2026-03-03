package proxy

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/apigear-io/cli/pkg/stream/config"
)

// Manager manages multiple proxy instances.
type Manager struct {
	mu         sync.RWMutex
	proxies    map[string]*Proxy
	stats      *Stats
	messageHub MessageHubPublisher
}

// NewManager creates a new proxy manager.
func NewManager() *Manager {
	return &Manager{
		proxies: make(map[string]*Proxy),
		stats:   NewStats(),
	}
}

// AddProxy adds a new proxy to the manager.
func (m *Manager) AddProxy(name string, cfg config.ProxyConfig) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.proxies[name]; exists {
		return fmt.Errorf("proxy %s already exists", name)
	}

	proxy := NewProxy(name, cfg.Listen, cfg.Backend, cfg)
	proxy.stats = m.stats.GetProxyStats(name)
	proxy.SetMessageHub(m.messageHub)

	m.proxies[name] = proxy

	log.Info().Str("name", name).Msg("proxy added")

	return nil
}

// RemoveProxy removes a proxy from the manager.
func (m *Manager) RemoveProxy(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	proxy, exists := m.proxies[name]
	if !exists {
		return fmt.Errorf("proxy %s not found", name)
	}

	// Stop if running
	if proxy.Status() == StatusRunning {
		if err := proxy.Stop(); err != nil {
			log.Warn().Err(err).Str("name", name).Msg("error stopping proxy")
		}
	}

	delete(m.proxies, name)
	m.stats.RemoveProxyStats(name)

	log.Info().Str("name", name).Msg("proxy removed")

	return nil
}

// GetProxy returns a proxy by name.
func (m *Manager) GetProxy(name string) (*Proxy, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	proxy, exists := m.proxies[name]
	if !exists {
		return nil, fmt.Errorf("proxy %s not found", name)
	}

	return proxy, nil
}

// ListProxies returns information about all proxies.
func (m *Manager) ListProxies() []Info {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]Info, 0, len(m.proxies))
	for _, proxy := range m.proxies {
		result = append(result, proxy.Info())
	}

	return result
}

// StartProxy starts a proxy by name.
func (m *Manager) StartProxy(name string) error {
	proxy, err := m.GetProxy(name)
	if err != nil {
		return err
	}

	return proxy.Start()
}

// StopProxy stops a proxy by name.
func (m *Manager) StopProxy(name string) error {
	proxy, err := m.GetProxy(name)
	if err != nil {
		return err
	}

	return proxy.Stop()
}

// LoadFromConfig loads proxies from configuration.
func (m *Manager) LoadFromConfig(proxies map[string]config.ProxyConfig) error {
	for name, cfg := range proxies {
		if cfg.Disabled {
			log.Info().Str("name", name).Msg("proxy disabled, skipping")
			continue
		}

		if err := m.AddProxy(name, cfg); err != nil {
			log.Warn().Err(err).Str("name", name).Msg("failed to add proxy")
			continue
		}

		// Auto-start proxy
		if err := m.StartProxy(name); err != nil {
			log.Warn().Err(err).Str("name", name).Msg("failed to start proxy")
		}
	}

	return nil
}

// StopAll stops all proxies.
func (m *Manager) StopAll() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for name, proxy := range m.proxies {
		if proxy.Status() == StatusRunning {
			if err := proxy.Stop(); err != nil {
				log.Warn().Err(err).Str("name", name).Msg("error stopping proxy")
			}
		}
	}

	return nil
}

// Close stops and removes all proxies.
func (m *Manager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for name, proxy := range m.proxies {
		if proxy.Status() == StatusRunning {
			if err := proxy.Stop(); err != nil {
				log.Warn().Err(err).Str("name", name).Msg("error stopping proxy")
			}
		}
	}

	m.proxies = make(map[string]*Proxy)
	m.stats = NewStats()

	return nil
}

// Stats returns the global stats collector.
func (m *Manager) Stats() *Stats {
	return m.stats
}

// SetMessageHub sets the message hub for all proxies.
func (m *Manager) SetMessageHub(hub MessageHubPublisher) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.messageHub = hub

	// Update all existing proxies
	for _, proxy := range m.proxies {
		proxy.SetMessageHub(hub)
	}
}
