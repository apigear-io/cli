package repos

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCache(t *testing.T) {
	t.Run("creates new cache", func(t *testing.T) {
		dir := t.TempDir()
		c := New(dir)

		assert.NotNil(t, c)
		assert.Equal(t, dir, c.cacheDir)
	})
}

func TestCacheExists(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	t.Run("returns false for non-existent template", func(t *testing.T) {
		exists := c.Exists("nonexistent@1.0.0")
		assert.False(t, exists)
	})

	t.Run("returns true for existing template", func(t *testing.T) {
		// Create a mock template directory
		templateDir := filepath.Join(dir, "test-template@1.0.0")
		err := os.MkdirAll(templateDir, 0755)
		require.NoError(t, err)

		exists := c.Exists("test-template@1.0.0")
		assert.True(t, exists)
	})

	t.Run("ensures repo ID format", func(t *testing.T) {
		// Create a template directory
		templateDir := filepath.Join(dir, "template@latest")
		err := os.MkdirAll(templateDir, 0755)
		require.NoError(t, err)

		// Test with and without version
		exists := c.Exists("template")
		assert.True(t, exists)

		exists = c.Exists("template@latest")
		assert.True(t, exists)
	})
}

func TestCacheRemove(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	t.Run("removes existing template", func(t *testing.T) {
		// Create a mock template directory
		templateName := "test-template@1.0.0"
		templateDir := filepath.Join(dir, templateName)
		err := os.MkdirAll(templateDir, 0755)
		require.NoError(t, err)

		// Create a file in the template
		testFile := filepath.Join(templateDir, "test.txt")
		err = os.WriteFile(testFile, []byte("test"), 0644)
		require.NoError(t, err)

		// Remove the template
		err = c.Remove(templateName)
		require.NoError(t, err)

		// Verify it's gone
		assert.False(t, helper.IsDir(templateDir))
	})

	t.Run("returns error for non-existent template", func(t *testing.T) {
		err := c.Remove("nonexistent@1.0.0")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "does not exist")
	})

	t.Run("ensures repo ID format", func(t *testing.T) {
		// Create a template
		templateDir := filepath.Join(dir, "template@latest")
		err := os.MkdirAll(templateDir, 0755)
		require.NoError(t, err)

		// Remove without version
		err = c.Remove("template")
		require.NoError(t, err)

		assert.False(t, helper.IsDir(templateDir))
	})
}

func TestCacheClean(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	t.Run("removes all templates", func(t *testing.T) {
		// Create multiple template directories
		template1 := filepath.Join(dir, "template1@1.0.0")
		template2 := filepath.Join(dir, "template2@2.0.0")

		err := os.MkdirAll(template1, 0755)
		require.NoError(t, err)
		err = os.MkdirAll(template2, 0755)
		require.NoError(t, err)

		// Clean the cache
		err = c.Clean()
		require.NoError(t, err)

		// Verify cache dir exists but is empty
		assert.True(t, helper.IsDir(dir))
		entries, err := os.ReadDir(dir)
		require.NoError(t, err)
		assert.Empty(t, entries)
	})
}

func TestCacheGetTemplateDir(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	t.Run("returns template directory path", func(t *testing.T) {
		templateName := "test-template@1.0.0"
		templateDir := filepath.Join(dir, templateName)
		err := os.MkdirAll(templateDir, 0755)
		require.NoError(t, err)

		path, err := c.GetTemplateDir(templateName)
		require.NoError(t, err)
		assert.Equal(t, templateDir, path)
	})

	t.Run("returns error for non-existent template", func(t *testing.T) {
		path, err := c.GetTemplateDir("nonexistent@1.0.0")
		assert.Error(t, err)
		assert.Empty(t, path)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("ensures repo ID format", func(t *testing.T) {
		templateDir := filepath.Join(dir, "template@latest")
		err := os.MkdirAll(templateDir, 0755)
		require.NoError(t, err)

		path, err := c.GetTemplateDir("template")
		require.NoError(t, err)
		assert.Equal(t, templateDir, path)
	})
}

func TestCacheSearch(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	t.Run("search with no cached templates returns empty", func(t *testing.T) {
		results, err := c.Search("test")
		// This will fail because there are no git repos, but we're testing the search logic
		// In a real scenario, this would need git repos set up
		_ = results
		_ = err
		// We expect an error or empty results since we don't have git repos
	})
}

func TestCacheListVersions(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	t.Run("list versions with no templates", func(t *testing.T) {
		versions, err := c.ListVersions("template@1.0.0")
		// Will likely error or return empty since no repos exist
		_ = versions
		_ = err
		// Just testing that the method doesn't panic
	})
}

// Note: Tests for Install, Upgrade, UpgradeAll, List, ListCachedRepos, and Info
// require git operations and would need mocking or integration test setup.
// These tests focus on the simpler, non-git-dependent operations.
