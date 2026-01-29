package cfg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	t.Run("creates config with defaults", func(t *testing.T) {
		dir := t.TempDir()

		cfg, err := NewConfig(dir)
		require.NoError(t, err)
		require.NotNil(t, cfg)

		// Check default values
		assert.Equal(t, 4333, cfg.GetInt(KeyServerPort))
		assert.Equal(t, "code", cfg.GetString(KeyEditorCommand))
		assert.Equal(t, "stable", cfg.GetString(KeyUpdateChannel))
		assert.Equal(t, registryUrl, cfg.GetString(KeyRegistryUrl))
	})

	t.Run("creates cache directory", func(t *testing.T) {
		dir := t.TempDir()

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		cacheDir := cfg.GetString(KeyCacheDir)
		assert.True(t, helper.IsDir(cacheDir))
	})

	t.Run("creates registry directory", func(t *testing.T) {
		dir := t.TempDir()

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		registryDir := cfg.GetString(KeyRegistryDir)
		assert.True(t, helper.IsDir(registryDir))
	})

	t.Run("creates config file if not exists", func(t *testing.T) {
		dir := t.TempDir()

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		cfgFile := filepath.Join(dir, "config.json")
		assert.True(t, helper.IsFile(cfgFile))
		assert.NotEmpty(t, cfg.ConfigFileUsed())
	})

	t.Run("reads existing config file", func(t *testing.T) {
		dir := t.TempDir()

		// Create config file
		cfgFile := filepath.Join(dir, "config.json")
		configData := `{"server_port": 8080, "editor_command": "vim"}`
		err := os.WriteFile(cfgFile, []byte(configData), 0644)
		require.NoError(t, err)

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		assert.Equal(t, 8080, cfg.GetInt(KeyServerPort))
		assert.Equal(t, "vim", cfg.GetString(KeyEditorCommand))
	})

	t.Run("respects APIGEAR_CACHE_DIR environment variable", func(t *testing.T) {
		dir := t.TempDir()
		customCacheDir := filepath.Join(dir, "custom-cache")

		os.Setenv("APIGEAR_CACHE_DIR", customCacheDir)
		defer os.Unsetenv("APIGEAR_CACHE_DIR")

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		assert.Equal(t, customCacheDir, cfg.GetString(KeyCacheDir))
		assert.True(t, helper.IsDir(customCacheDir))
	})

	t.Run("respects APIGEAR_REGISTRY_DIR environment variable", func(t *testing.T) {
		dir := t.TempDir()
		customRegistryDir := filepath.Join(dir, "custom-registry")

		os.Setenv("APIGEAR_REGISTRY_DIR", customRegistryDir)
		defer os.Unsetenv("APIGEAR_REGISTRY_DIR")

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		assert.Equal(t, customRegistryDir, cfg.GetString(KeyRegistryDir))
		assert.True(t, helper.IsDir(customRegistryDir))
	})

	t.Run("sets all default values", func(t *testing.T) {
		dir := t.TempDir()

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		// Check all defaults
		assert.Equal(t, 4333, cfg.GetInt(KeyServerPort))
		assert.Equal(t, "code", cfg.GetString(KeyEditorCommand))
		assert.Equal(t, "stable", cfg.GetString(KeyUpdateChannel))
		assert.Equal(t, "0.0.0", cfg.GetString(KeyVersion))
		assert.Equal(t, "none", cfg.GetString(KeyCommit))
		assert.Equal(t, "unknown", cfg.GetString(KeyDate))
		assert.Equal(t, 960, cfg.GetInt(KeyWindowWidth))
		assert.Equal(t, 720, cfg.GetInt(KeyWindowHeight))
	})
}

func TestSetConfig(t *testing.T) {
	t.Run("sets global config", func(t *testing.T) {
		dir := t.TempDir()

		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		// Backup original config
		originalCfg := v
		defer func() {
			SetConfig(originalCfg)
		}()

		// Set new config
		SetConfig(cfg)

		// Verify it was set
		assert.NotNil(t, v)
	})
}

func TestConfigThreadSafety(t *testing.T) {
	t.Run("concurrent reads and writes", func(t *testing.T) {
		dir := t.TempDir()
		cfg, err := NewConfig(dir)
		require.NoError(t, err)

		// Backup original config
		originalCfg := v
		defer func() {
			SetConfig(originalCfg)
		}()

		SetConfig(cfg)

		// Concurrent writes
		done := make(chan bool)
		for i := 0; i < 10; i++ {
			go func(val int) {
				Set("test_key", val)
				done <- true
			}(i)
		}

		// Concurrent reads
		for i := 0; i < 10; i++ {
			go func() {
				_ = Get("test_key")
				done <- true
			}()
		}

		// Wait for all goroutines
		for i := 0; i < 20; i++ {
			<-done
		}

		// Test passed if no race conditions occurred
	})
}
