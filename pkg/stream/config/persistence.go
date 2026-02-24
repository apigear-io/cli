package config

import (
	"fmt"
	"os"
	"sync"
)

// ConfigPersistence provides thread-safe persistence operations for config files.
type ConfigPersistence struct {
	mu   sync.RWMutex
	path string
}

// NewConfigPersistence creates a new config persistence handler.
func NewConfigPersistence(path string) *ConfigPersistence {
	return &ConfigPersistence{
		path: path,
	}
}

// WithConfig executes a function with the loaded config and saves it back.
// This provides thread-safe read-modify-write operations.
func (cp *ConfigPersistence) WithConfig(fn func(*Config) error) error {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	// Load current config
	cfg, err := cp.loadOrCreate()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Execute the modification function
	if err := fn(cfg); err != nil {
		return err
	}

	// Save the modified config
	if err := SaveConfig(cp.path, cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

// ReadConfig reads the config without locking (for read-only operations).
func (cp *ConfigPersistence) ReadConfig() (*Config, error) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.loadOrCreate()
}

// loadOrCreate loads the config or creates a default one if it doesn't exist.
func (cp *ConfigPersistence) loadOrCreate() (*Config, error) {
	// Check if file exists
	if _, err := os.Stat(cp.path); os.IsNotExist(err) {
		// Create default config
		cfg := DefaultConfig()
		if err := SaveConfig(cp.path, cfg); err != nil {
			return nil, fmt.Errorf("failed to create default config: %w", err)
		}
		return cfg, nil
	}

	// Load existing config
	cfg, err := LoadConfig(cp.path)
	if err != nil {
		return nil, err
	}

	// Initialize maps if nil
	if cfg.Proxies == nil {
		cfg.Proxies = make(map[string]ProxyConfig)
	}
	if cfg.Clients == nil {
		cfg.Clients = make(map[string]ClientConfig)
	}

	return cfg, nil
}

// AddProxy adds a proxy to the config file.
func (cp *ConfigPersistence) AddProxy(name string, proxy ProxyConfig) error {
	return cp.WithConfig(func(cfg *Config) error {
		if _, exists := cfg.Proxies[name]; exists {
			return fmt.Errorf("proxy %s already exists", name)
		}
		cfg.Proxies[name] = proxy
		return nil
	})
}

// UpdateProxy updates a proxy in the config file.
func (cp *ConfigPersistence) UpdateProxy(name string, proxy ProxyConfig) error {
	return cp.WithConfig(func(cfg *Config) error {
		if _, exists := cfg.Proxies[name]; !exists {
			return fmt.Errorf("proxy %s not found", name)
		}
		cfg.Proxies[name] = proxy
		return nil
	})
}

// DeleteProxy removes a proxy from the config file.
func (cp *ConfigPersistence) DeleteProxy(name string) error {
	return cp.WithConfig(func(cfg *Config) error {
		if _, exists := cfg.Proxies[name]; !exists {
			return fmt.Errorf("proxy %s not found", name)
		}
		delete(cfg.Proxies, name)
		return nil
	})
}

// AddClient adds a client to the config file.
func (cp *ConfigPersistence) AddClient(name string, client ClientConfig) error {
	return cp.WithConfig(func(cfg *Config) error {
		if _, exists := cfg.Clients[name]; exists {
			return fmt.Errorf("client %s already exists", name)
		}
		cfg.Clients[name] = client
		return nil
	})
}

// UpdateClient updates a client in the config file.
func (cp *ConfigPersistence) UpdateClient(name string, client ClientConfig) error {
	return cp.WithConfig(func(cfg *Config) error {
		if _, exists := cfg.Clients[name]; !exists {
			return fmt.Errorf("client %s not found", name)
		}
		cfg.Clients[name] = client
		return nil
	})
}

// DeleteClient removes a client from the config file.
func (cp *ConfigPersistence) DeleteClient(name string) error {
	return cp.WithConfig(func(cfg *Config) error {
		if _, exists := cfg.Clients[name]; !exists {
			return fmt.Errorf("client %s not found", name)
		}
		delete(cfg.Clients, name)
		return nil
	})
}
