package proxy

import (
	"testing"

	"github.com/apigear-io/cli/pkg/stream/config"
)

func TestNewManager(t *testing.T) {
	manager := NewManager()

	if manager == nil {
		t.Fatal("expected manager to be created")
	}

	proxies := manager.ListProxies()
	if len(proxies) != 0 {
		t.Errorf("expected no proxies, got %d", len(proxies))
	}
}

func TestManagerAddProxy(t *testing.T) {
	manager := NewManager()

	cfg := config.ProxyConfig{
		Listen:  "ws://localhost:5550/ws",
		Backend: "ws://localhost:5560/ws",
		Mode:    "proxy",
	}

	err := manager.AddProxy("test", cfg)
	if err != nil {
		t.Fatalf("AddProxy failed: %v", err)
	}

	// Try to add the same proxy again
	err = manager.AddProxy("test", cfg)
	if err == nil {
		t.Error("expected error when adding duplicate proxy")
	}

	// Verify proxy exists
	proxy, err := manager.GetProxy("test")
	if err != nil {
		t.Fatalf("GetProxy failed: %v", err)
	}

	if proxy.name != "test" {
		t.Errorf("expected name test, got %s", proxy.name)
	}

	// Verify list
	proxies := manager.ListProxies()
	if len(proxies) != 1 {
		t.Errorf("expected 1 proxy, got %d", len(proxies))
	}
	if proxies[0].Name != "test" {
		t.Errorf("expected proxy name test, got %s", proxies[0].Name)
	}
}

func TestManagerRemoveProxy(t *testing.T) {
	manager := NewManager()

	cfg := config.ProxyConfig{
		Listen:  "ws://localhost:5550/ws",
		Backend: "ws://localhost:5560/ws",
		Mode:    "proxy",
	}

	err := manager.AddProxy("test", cfg)
	if err != nil {
		t.Fatalf("AddProxy failed: %v", err)
	}

	// Remove proxy
	err = manager.RemoveProxy("test")
	if err != nil {
		t.Fatalf("RemoveProxy failed: %v", err)
	}

	// Verify proxy doesn't exist
	_, err = manager.GetProxy("test")
	if err == nil {
		t.Error("expected error when getting removed proxy")
	}

	// Try to remove non-existent proxy
	err = manager.RemoveProxy("test")
	if err == nil {
		t.Error("expected error when removing non-existent proxy")
	}
}

func TestManagerLoadFromConfig(t *testing.T) {
	manager := NewManager()

	proxies := map[string]config.ProxyConfig{
		"proxy1": {
			Listen:  "ws://localhost:5550/ws",
			Backend: "ws://localhost:5560/ws",
			Mode:    "proxy",
		},
		"proxy2": {
			Listen: "ws://localhost:5551/ws",
			Mode:   "echo",
		},
		"proxy3": {
			Listen:   "ws://localhost:5552/ws",
			Backend:  "ws://localhost:5562/ws",
			Mode:     "proxy",
			Disabled: true,
		},
	}

	err := manager.LoadFromConfig(proxies)
	if err != nil {
		t.Fatalf("LoadFromConfig failed: %v", err)
	}

	// Verify proxies loaded (excluding disabled)
	proxyList := manager.ListProxies()
	if len(proxyList) != 2 {
		t.Errorf("expected 2 proxies, got %d", len(proxyList))
	}

	// Verify proxy1
	proxy1, err := manager.GetProxy("proxy1")
	if err != nil {
		t.Fatalf("GetProxy failed: %v", err)
	}
	if proxy1.backend != "ws://localhost:5560/ws" {
		t.Errorf("expected proxy1 backend ws://localhost:5560/ws, got %s", proxy1.backend)
	}

	// Verify proxy2
	proxy2, err := manager.GetProxy("proxy2")
	if err != nil {
		t.Fatalf("GetProxy failed: %v", err)
	}
	if proxy2.mode != ModeEcho {
		t.Errorf("expected proxy2 mode echo, got %s", proxy2.mode.String())
	}

	// Verify proxy3 was not loaded
	_, err = manager.GetProxy("proxy3")
	if err == nil {
		t.Error("expected proxy3 to not be loaded (disabled)")
	}

	// Cleanup - stop all proxies
	manager.StopAll()
}

func TestManagerClose(t *testing.T) {
	manager := NewManager()

	cfg := config.ProxyConfig{
		Listen:  "ws://localhost:5550/ws",
		Backend: "ws://localhost:5560/ws",
		Mode:    "proxy",
	}

	err := manager.AddProxy("test", cfg)
	if err != nil {
		t.Fatalf("AddProxy failed: %v", err)
	}

	// Close manager
	err = manager.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	// Verify all proxies removed
	proxies := manager.ListProxies()
	if len(proxies) != 0 {
		t.Errorf("expected no proxies after close, got %d", len(proxies))
	}
}

func TestParseMode(t *testing.T) {
	tests := []struct {
		input    string
		expected Mode
	}{
		{"proxy", ModeProxy},
		{"echo", ModeEcho},
		{"backend", ModeBackend},
		{"inbound", ModeInbound},
		{"inbound-only", ModeInbound},
		{"", ModeProxy},       // default
		{"unknown", ModeProxy}, // default
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ParseMode(tt.input)
			if result != tt.expected {
				t.Errorf("ParseMode(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestModeString(t *testing.T) {
	tests := []struct {
		mode     Mode
		expected string
	}{
		{ModeProxy, "proxy"},
		{ModeEcho, "echo"},
		{ModeBackend, "backend"},
		{ModeInbound, "inbound-only"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.mode.String()
			if result != tt.expected {
				t.Errorf("Mode.String() = %q, want %q", result, tt.expected)
			}
		})
	}
}
