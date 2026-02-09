package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestConfig creates a test config in a temporary directory
func setupTestConfig(t *testing.T) {
	dir := t.TempDir()
	cfg, err := NewConfig(dir)
	require.NoError(t, err)
	SetConfig(cfg)
}

func TestConfigDir(t *testing.T) {
	setupTestConfig(t)

	dir := ConfigDir()
	assert.NotEmpty(t, dir)
}

func TestRecentEntries(t *testing.T) {
	setupTestConfig(t)

	t.Run("empty recent entries", func(t *testing.T) {
		entries := RecentEntries()
		assert.NotNil(t, entries)
	})

	t.Run("append recent entry", func(t *testing.T) {
		err := AppendRecentEntry("/path/to/project1")
		require.NoError(t, err)

		entries := RecentEntries()
		assert.Contains(t, entries, "/path/to/project1")
		assert.Equal(t, "/path/to/project1", entries[0])
	})

	t.Run("append multiple entries", func(t *testing.T) {
		setupTestConfig(t)

		err := AppendRecentEntry("/path/1")
		require.NoError(t, err)
		err = AppendRecentEntry("/path/2")
		require.NoError(t, err)
		err = AppendRecentEntry("/path/3")
		require.NoError(t, err)

		entries := RecentEntries()
		assert.Len(t, entries, 3)
		assert.Equal(t, "/path/3", entries[0])
		assert.Equal(t, "/path/2", entries[1])
		assert.Equal(t, "/path/1", entries[2])
	})

	t.Run("duplicate entry moves to front", func(t *testing.T) {
		setupTestConfig(t)

		err := AppendRecentEntry("/path/1")
		require.NoError(t, err)
		err = AppendRecentEntry("/path/2")
		require.NoError(t, err)
		err = AppendRecentEntry("/path/3")
		require.NoError(t, err)

		// Add duplicate
		err = AppendRecentEntry("/path/2")
		require.NoError(t, err)

		entries := RecentEntries()
		assert.Len(t, entries, 3)
		assert.Equal(t, "/path/2", entries[0])
		assert.Equal(t, "/path/3", entries[1])
		assert.Equal(t, "/path/1", entries[2])
	})

	t.Run("limits to 5 entries", func(t *testing.T) {
		setupTestConfig(t)

		for i := 1; i <= 7; i++ {
			err := AppendRecentEntry("/path/" + string(rune('0'+i)))
			require.NoError(t, err)
		}

		entries := RecentEntries()
		assert.Len(t, entries, 5)
		// Most recent should be first
		assert.Equal(t, "/path/7", entries[0])
	})

	t.Run("remove recent entry", func(t *testing.T) {
		setupTestConfig(t)

		err := AppendRecentEntry("/path/1")
		require.NoError(t, err)
		err = AppendRecentEntry("/path/2")
		require.NoError(t, err)
		err = AppendRecentEntry("/path/3")
		require.NoError(t, err)

		err = RemoveRecentEntry("/path/2")
		require.NoError(t, err)

		entries := RecentEntries()
		assert.Len(t, entries, 2)
		assert.NotContains(t, entries, "/path/2")
		assert.Contains(t, entries, "/path/1")
		assert.Contains(t, entries, "/path/3")
	})

	t.Run("remove non-existent entry", func(t *testing.T) {
		setupTestConfig(t)

		err := AppendRecentEntry("/path/1")
		require.NoError(t, err)

		err = RemoveRecentEntry("/path/nonexistent")
		require.NoError(t, err)

		entries := RecentEntries()
		assert.Len(t, entries, 1)
	})
}

func TestBuildInfo(t *testing.T) {
	setupTestConfig(t)

	t.Run("set and get build info", func(t *testing.T) {
		info := BuildInfo{
			Version: "1.0.0",
			Commit:  "abc123",
			Date:    "2024-01-01",
		}

		SetBuildInfo("cli", info)
		result := GetBuildInfo("cli")

		assert.Equal(t, info.Version, result.Version)
		assert.Equal(t, info.Commit, result.Commit)
		assert.Equal(t, info.Date, result.Date)
	})

	t.Run("get non-existent build info", func(t *testing.T) {
		result := GetBuildInfo("nonexistent")

		// Should return zero value
		assert.Empty(t, result.Version)
		assert.Empty(t, result.Commit)
		assert.Empty(t, result.Date)
	})

	t.Run("multiple build infos", func(t *testing.T) {
		info1 := BuildInfo{Version: "1.0.0", Commit: "abc", Date: "2024-01-01"}
		info2 := BuildInfo{Version: "2.0.0", Commit: "def", Date: "2024-02-01"}

		SetBuildInfo("cli", info1)
		SetBuildInfo("studio", info2)

		result1 := GetBuildInfo("cli")
		result2 := GetBuildInfo("studio")

		assert.Equal(t, "1.0.0", result1.Version)
		assert.Equal(t, "2.0.0", result2.Version)
	})
}

func TestConfigGetSet(t *testing.T) {
	setupTestConfig(t)

	t.Run("IsSet", func(t *testing.T) {
		Set("test_key", "test_value")
		assert.True(t, IsSet("test_key"))
		assert.False(t, IsSet("nonexistent_key"))
	})

	t.Run("Get and Set", func(t *testing.T) {
		Set("string_key", "value")
		assert.Equal(t, "value", Get("string_key"))

		Set("int_key", 42)
		assert.Equal(t, 42, Get("int_key"))
	})

	t.Run("GetInt", func(t *testing.T) {
		Set("int_key", 100)
		assert.Equal(t, 100, GetInt("int_key"))

		// Non-existent key returns 0
		assert.Equal(t, 0, GetInt("nonexistent"))
	})

	t.Run("GetBool and SetBool", func(t *testing.T) {
		SetBool("bool_key", true)
		assert.True(t, GetBool("bool_key"))

		SetBool("bool_key", false)
		assert.False(t, GetBool("bool_key"))
	})

	t.Run("GetString", func(t *testing.T) {
		Set("string_key", "hello")
		assert.Equal(t, "hello", GetString("string_key"))

		// Non-existent key returns empty string
		assert.Equal(t, "", GetString("nonexistent"))
	})
}

func TestConfigHelpers(t *testing.T) {
	setupTestConfig(t)

	t.Run("EditorCommand", func(t *testing.T) {
		cmd := EditorCommand()
		assert.NotEmpty(t, cmd)
		assert.Equal(t, "code", cmd)
	})

	t.Run("ServerPort", func(t *testing.T) {
		port := ServerPort()
		assert.NotEmpty(t, port)
	})

	t.Run("UpdateChannel", func(t *testing.T) {
		channel := UpdateChannel()
		assert.NotEmpty(t, channel)
		assert.Equal(t, "stable", channel)
	})

	t.Run("RegistryDir", func(t *testing.T) {
		dir := RegistryDir()
		assert.NotEmpty(t, dir)
	})

	t.Run("RegistryCachePath", func(t *testing.T) {
		path := RegistryCachePath()
		assert.NotEmpty(t, path)
		assert.Contains(t, path, "registry.json")
	})

	t.Run("CacheDir", func(t *testing.T) {
		dir := CacheDir()
		assert.NotEmpty(t, dir)
	})

	t.Run("RegistryUrl", func(t *testing.T) {
		url := RegistryUrl()
		assert.NotEmpty(t, url)
		assert.Equal(t, registryUrl, url)
	})

	t.Run("AllSettings", func(t *testing.T) {
		settings := AllSettings()
		assert.NotEmpty(t, settings)

		// Check that default settings are present
		assert.Contains(t, settings, "server_port")
		assert.Contains(t, settings, "editor_command")
	})

	t.Run("ConfigFileUsed", func(t *testing.T) {
		file := ConfigFileUsed()
		assert.NotEmpty(t, file)
		assert.Contains(t, file, "config.json")
	})
}

func TestWriteConfig(t *testing.T) {
	setupTestConfig(t)

	t.Run("write config", func(t *testing.T) {
		Set("test_key", "test_value")

		err := WriteConfig()
		assert.NoError(t, err)

		// Verify the value persists
		assert.Equal(t, "test_value", GetString("test_key"))
	})
}
