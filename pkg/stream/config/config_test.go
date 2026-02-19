package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadYAMLConfig(t *testing.T) {
	yamlContent := `
verbose: true
trace: true
traceDir: ./traces
proxies:
  test:
    listen: ws://localhost:5550/ws
    backend: ws://localhost:5560/ws
    mode: proxy
clients:
  testclient:
    url: ws://localhost:5560/ws
    interfaces:
      - demo.Counter
    enabled: true
    autoReconnect: true
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	if err := os.WriteFile(configPath, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	config, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if !config.Verbose {
		t.Error("expected Verbose to be true")
	}
	if !config.Trace {
		t.Error("expected Trace to be true")
	}
	if config.TraceDir != "./traces" {
		t.Errorf("expected TraceDir ./traces, got %s", config.TraceDir)
	}

	// Check proxy
	proxy, ok := config.Proxies["test"]
	if !ok {
		t.Fatal("expected proxy 'test' to exist")
	}
	if proxy.Listen != "ws://localhost:5550/ws" {
		t.Errorf("expected Listen ws://localhost:5550/ws, got %s", proxy.Listen)
	}
	if proxy.Backend != "ws://localhost:5560/ws" {
		t.Errorf("expected Backend ws://localhost:5560/ws, got %s", proxy.Backend)
	}
	if proxy.Mode != "proxy" {
		t.Errorf("expected Mode proxy, got %s", proxy.Mode)
	}

	// Check client
	client, ok := config.Clients["testclient"]
	if !ok {
		t.Fatal("expected client 'testclient' to exist")
	}
	if client.URL != "ws://localhost:5560/ws" {
		t.Errorf("expected URL ws://localhost:5560/ws, got %s", client.URL)
	}
	if len(client.Interfaces) != 1 || client.Interfaces[0] != "demo.Counter" {
		t.Errorf("expected Interfaces [demo.Counter], got %v", client.Interfaces)
	}
	if !client.Enabled {
		t.Error("expected Enabled to be true")
	}
	if !client.AutoReconnect {
		t.Error("expected AutoReconnect to be true")
	}
}

func TestSaveConfig(t *testing.T) {
	config := &Config{
		Verbose:  true,
		TraceDir: "./traces",
		Proxies: map[string]ProxyConfig{
			"test": {
				Listen:  "ws://localhost:5550/ws",
				Backend: "ws://localhost:5560/ws",
				Mode:    "proxy",
			},
		},
	}

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	if err := SaveConfig(configPath, config); err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	// Load it back
	loaded, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if !loaded.Verbose {
		t.Error("expected Verbose to be true")
	}
	if loaded.TraceDir != "./traces" {
		t.Errorf("expected TraceDir ./traces, got %s", loaded.TraceDir)
	}

	proxy, ok := loaded.Proxies["test"]
	if !ok {
		t.Fatal("expected proxy 'test' to exist")
	}
	if proxy.Listen != "ws://localhost:5550/ws" {
		t.Errorf("expected Listen ws://localhost:5550/ws, got %s", proxy.Listen)
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid proxy config",
			config: &Config{
				Proxies: map[string]ProxyConfig{
					"test": {
						Listen:  "ws://localhost:5550/ws",
						Backend: "ws://localhost:5560/ws",
						Mode:    "proxy",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing listen address",
			config: &Config{
				Proxies: map[string]ProxyConfig{
					"test": {
						Backend: "ws://localhost:5560/ws",
						Mode:    "proxy",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing mode",
			config: &Config{
				Proxies: map[string]ProxyConfig{
					"test": {
						Listen:  "ws://localhost:5550/ws",
						Backend: "ws://localhost:5560/ws",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid mode",
			config: &Config{
				Proxies: map[string]ProxyConfig{
					"test": {
						Listen:  "ws://localhost:5550/ws",
						Backend: "ws://localhost:5560/ws",
						Mode:    "invalid",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "proxy mode without backend",
			config: &Config{
				Proxies: map[string]ProxyConfig{
					"test": {
						Listen: "ws://localhost:5550/ws",
						Mode:   "proxy",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "echo mode without backend",
			config: &Config{
				Proxies: map[string]ProxyConfig{
					"test": {
						Listen: "ws://localhost:5550/ws",
						Mode:   "echo",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "client without URL",
			config: &Config{
				Clients: map[string]ClientConfig{
					"test": {
						Interfaces: []string{"demo.Counter"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid client",
			config: &Config{
				Clients: map[string]ClientConfig{
					"test": {
						URL:        "ws://localhost:5560/ws",
						Interfaces: []string{"demo.Counter"},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadOrCreateConfig(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	// First call should create the file
	config, created, err := LoadOrCreateConfig(configPath)
	if err != nil {
		t.Fatalf("LoadOrCreateConfig failed: %v", err)
	}
	if !created {
		t.Error("expected config to be created")
	}
	if config == nil {
		t.Fatal("expected config to be returned")
	}

	// Second call should load existing file
	config2, created2, err := LoadOrCreateConfig(configPath)
	if err != nil {
		t.Fatalf("LoadOrCreateConfig failed: %v", err)
	}
	if created2 {
		t.Error("expected config not to be created")
	}
	if config2 == nil {
		t.Fatal("expected config to be returned")
	}

	// Verify file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("expected config file to exist")
	}
}

func TestDefaultTraceConfig(t *testing.T) {
	tc := DefaultTraceConfig()

	if tc.MaxSizeMB != 10 {
		t.Errorf("expected MaxSizeMB 10, got %d", tc.MaxSizeMB)
	}
	if tc.MaxBackups != 5 {
		t.Errorf("expected MaxBackups 5, got %d", tc.MaxBackups)
	}
	if tc.MaxAgeDays != 7 {
		t.Errorf("expected MaxAgeDays 7, got %d", tc.MaxAgeDays)
	}
	if !tc.Compress {
		t.Error("expected Compress to be true")
	}
}
