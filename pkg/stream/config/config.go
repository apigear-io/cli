// Package config provides configuration management for stream functionality.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

// TraceConfig defines trace file rotation settings.
type TraceConfig struct {
	MaxSizeMB  int  `json:"maxSizeMB" yaml:"maxSizeMB"`    // Maximum file size in MB before rotation
	MaxBackups int  `json:"maxBackups" yaml:"maxBackups"`  // Maximum number of old files to keep
	MaxAgeDays int  `json:"maxAgeDays" yaml:"maxAgeDays"`  // Maximum age in days before deletion
	Compress   bool `json:"compress" yaml:"compress"`      // Compress rotated files
}

// DefaultTraceConfig returns default trace configuration.
func DefaultTraceConfig() TraceConfig {
	return TraceConfig{
		MaxSizeMB:  10,
		MaxBackups: 5,
		MaxAgeDays: 7,
		Compress:   true,
	}
}

// WebConfig defines the web UI server configuration.
type WebConfig struct {
	Listen string `json:"listen,omitempty" yaml:"listen,omitempty"`
}

// ProxyConfig defines a single proxy configuration.
type ProxyConfig struct {
	Listen  string `json:"listen" yaml:"listen"`    // Listen address (e.g., "ws://localhost:5550/ws")
	Backend string `json:"backend" yaml:"backend"`  // Backend URL (e.g., "ws://localhost:5560/ws")
	Mode    string `json:"mode" yaml:"mode"`        // Mode: "proxy", "echo", "backend", "inbound-only"
	Disabled bool  `json:"disabled,omitempty" yaml:"disabled,omitempty"` // If true, proxy is not started
}

// ClientConfig defines an ObjectLink client configuration.
type ClientConfig struct {
	URL           string   `json:"url" yaml:"url"`                           // WebSocket URL
	Interfaces    []string `json:"interfaces" yaml:"interfaces"`             // ObjectLink interfaces to link
	Enabled       bool     `json:"enabled" yaml:"enabled"`                   // Whether client is enabled
	AutoReconnect bool     `json:"autoReconnect" yaml:"autoReconnect"`       // Auto-reconnect on disconnect
}

// Config is the top-level stream configuration.
type Config struct {
	Verbose     bool                   `json:"verbose,omitempty" yaml:"verbose,omitempty"`
	Trace       bool                   `json:"trace,omitempty" yaml:"trace,omitempty"`
	TraceConfig TraceConfig            `json:"traceConfig,omitempty" yaml:"traceConfig,omitempty"`
	TraceDir    string                 `json:"traceDir,omitempty" yaml:"traceDir,omitempty"`
	LogFile     string                 `json:"logFile,omitempty" yaml:"logFile,omitempty"`
	Watch       bool                   `json:"watch,omitempty" yaml:"watch,omitempty"`
	Web         WebConfig              `json:"web,omitempty" yaml:"web,omitempty"`
	Proxies     map[string]ProxyConfig `json:"proxies" yaml:"proxies"`
	Clients     map[string]ClientConfig `json:"clients,omitempty" yaml:"clients,omitempty"`
}

// LoadConfig loads configuration from a YAML or JSON file.
func LoadConfig(path string) (*Config, error) {
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return loadYAMLConfig(path)
	}
	return loadJSONConfig(path)
}

func loadJSONConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func loadYAMLConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves the configuration to a file.
func SaveConfig(path string, config *Config) error {
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return saveYAMLConfig(path, config)
	}
	return saveJSONConfig(path, config)
}

func saveJSONConfig(path string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func saveYAMLConfig(path string, config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// ReadConfigRaw reads the raw content of a config file.
func ReadConfigRaw(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteConfigRaw writes raw content to a config file.
func WriteConfigRaw(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// ValidateConfigYAML validates that the content is valid YAML config.
func ValidateConfigYAML(content string) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal([]byte(content), &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// DefaultConfig returns a sample configuration.
func DefaultConfig() *Config {
	return &Config{
		TraceConfig: DefaultTraceConfig(),
		Proxies:     map[string]ProxyConfig{},
		Clients:     map[string]ClientConfig{},
	}
}

// DefaultConfigYAML returns a sample YAML configuration with comments.
const DefaultConfigYAML = `# ApiGear Stream configuration
# Documentation: https://github.com/apigear-io/cli

# Proxy definitions - forward WebSocket connections to backends
# proxies:
#   example:
#     listen: ws://localhost:5550/ws
#     backend: ws://localhost:5551/ws
#     mode: proxy  # proxy, echo, backend, inbound-only
#     disabled: false

# Client definitions - connect to ObjectLink backends
# clients:
#   example:
#     url: ws://localhost:5560/ws
#     interfaces:
#       - Module.Interface
#     enabled: true
#     autoReconnect: true

# Web UI settings
# web:
#   listen: ":8080"

# Trace file settings
# trace: false
# traceDir: "./data/traces"
# traceConfig:
#   maxSizeMB: 10
#   maxBackups: 5
#   maxAgeDays: 7
#   compress: true

# Application log file
# logFile: "./data/logs/stream.log"
`

// LoadOrCreateConfig loads the config file, or creates a default one if it doesn't exist.
// Returns the config, whether it was created, and any error.
func LoadOrCreateConfig(path string) (*Config, bool, error) {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create default config
		if err := os.WriteFile(path, []byte(DefaultConfigYAML), 0644); err != nil {
			return nil, false, fmt.Errorf("failed to create config file: %w", err)
		}
		// Load the created config
		config, err := LoadConfig(path)
		if err != nil {
			return nil, true, err
		}
		return config, true, nil
	}

	// File exists, load it
	config, err := LoadConfig(path)
	if err != nil {
		return nil, false, err
	}
	return config, false, nil
}

// Validate validates the configuration and returns any errors.
func (c *Config) Validate() error {
	// Validate proxies
	for name, proxy := range c.Proxies {
		if proxy.Listen == "" {
			return fmt.Errorf("proxy %s: listen address is required", name)
		}
		if proxy.Mode == "" {
			return fmt.Errorf("proxy %s: mode is required", name)
		}
		validModes := map[string]bool{
			"proxy":         true,
			"echo":          true,
			"backend":       true,
			"inbound-only":  true,
		}
		if !validModes[proxy.Mode] {
			return fmt.Errorf("proxy %s: invalid mode %s (must be proxy, echo, backend, or inbound-only)", name, proxy.Mode)
		}
		if proxy.Mode == "proxy" && proxy.Backend == "" {
			return fmt.Errorf("proxy %s: backend URL is required for proxy mode", name)
		}
	}

	// Validate clients
	for name, client := range c.Clients {
		if client.URL == "" {
			return fmt.Errorf("client %s: URL is required", name)
		}
	}

	return nil
}

// GetTraceConfig returns trace config with defaults applied.
func (c *Config) GetTraceConfig() TraceConfig {
	if c.TraceConfig.MaxSizeMB == 0 {
		c.TraceConfig = DefaultTraceConfig()
	}
	return c.TraceConfig
}
