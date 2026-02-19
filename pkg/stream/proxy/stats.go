package proxy

import (
	"sync"
	"sync/atomic"
	"time"
)

// Stats tracks proxy statistics.
type Stats struct {
	mu                sync.RWMutex
	proxies           map[string]*ProxyStats
	startTime         time.Time
	messagesReceived  atomic.Int64
	messagesSent      atomic.Int64
	bytesReceived     atomic.Int64
	bytesSent         atomic.Int64
}

// ProxyStats tracks statistics for a single proxy.
type ProxyStats struct {
	Name              string
	MessagesReceived  atomic.Int64
	MessagesSent      atomic.Int64
	BytesReceived     atomic.Int64
	BytesSent         atomic.Int64
	ActiveConnections atomic.Int32
	StartTime         time.Time
}

// NewStats creates a new stats collector.
func NewStats() *Stats {
	return &Stats{
		proxies:   make(map[string]*ProxyStats),
		startTime: time.Now(),
	}
}

// GetProxyStats returns stats for a specific proxy, creating if needed.
func (s *Stats) GetProxyStats(name string) *ProxyStats {
	s.mu.Lock()
	defer s.mu.Unlock()

	if stats, exists := s.proxies[name]; exists {
		return stats
	}

	stats := &ProxyStats{
		Name:      name,
		StartTime: time.Now(),
	}
	s.proxies[name] = stats
	return stats
}

// RemoveProxyStats removes stats for a proxy.
func (s *Stats) RemoveProxyStats(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.proxies, name)
}

// RecordMessageReceived records a received message.
func (ps *ProxyStats) RecordMessageReceived(size int) {
	ps.MessagesReceived.Add(1)
	ps.BytesReceived.Add(int64(size))
}

// RecordMessageSent records a sent message.
func (ps *ProxyStats) RecordMessageSent(size int) {
	ps.MessagesSent.Add(1)
	ps.BytesSent.Add(int64(size))
}

// RecordConnectionOpened records a new connection.
func (ps *ProxyStats) RecordConnectionOpened() {
	ps.ActiveConnections.Add(1)
}

// RecordConnectionClosed records a closed connection.
func (ps *ProxyStats) RecordConnectionClosed() {
	ps.ActiveConnections.Add(-1)
}

// GetInfo returns current statistics as Info.
func (ps *ProxyStats) GetInfo() Info {
	return Info{
		Name:              ps.Name,
		MessagesReceived:  ps.MessagesReceived.Load(),
		MessagesSent:      ps.MessagesSent.Load(),
		BytesReceived:     ps.BytesReceived.Load(),
		BytesSent:         ps.BytesSent.Load(),
		ActiveConnections: int(ps.ActiveConnections.Load()),
		Uptime:            int64(time.Since(ps.StartTime).Seconds()),
	}
}

// AllProxyStats returns stats for all proxies.
func (s *Stats) AllProxyStats() []Info {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Info, 0, len(s.proxies))
	for _, stats := range s.proxies {
		result = append(result, stats.GetInfo())
	}
	return result
}

// GlobalStats returns global statistics across all proxies.
func (s *Stats) GlobalStats() struct {
	TotalProxies      int
	MessagesReceived  int64
	MessagesSent      int64
	BytesReceived     int64
	BytesSent         int64
	Uptime            int64
} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var messagesReceived, messagesSent, bytesReceived, bytesSent int64
	for _, stats := range s.proxies {
		messagesReceived += stats.MessagesReceived.Load()
		messagesSent += stats.MessagesSent.Load()
		bytesReceived += stats.BytesReceived.Load()
		bytesSent += stats.BytesSent.Load()
	}

	return struct {
		TotalProxies      int
		MessagesReceived  int64
		MessagesSent      int64
		BytesReceived     int64
		BytesSent         int64
		Uptime            int64
	}{
		TotalProxies:     len(s.proxies),
		MessagesReceived: messagesReceived,
		MessagesSent:     messagesSent,
		BytesReceived:    bytesReceived,
		BytesSent:        bytesSent,
		Uptime:           int64(time.Since(s.startTime).Seconds()),
	}
}
