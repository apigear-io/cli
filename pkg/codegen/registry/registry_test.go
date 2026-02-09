package registry

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/apigear-io/cli/pkg/foundation/git"
	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRegistry(t *testing.T) {
	t.Run("creates new registry", func(t *testing.T) {
		dir := t.TempDir()
		url := "https://example.com/registry.git"

		r := NewRegistry(dir, url)

		assert.NotNil(t, r)
		assert.Equal(t, dir, r.RegistryDir)
		assert.Equal(t, url, r.RegistryURL)
	})
}

func TestRegistryLoadSave(t *testing.T) {
	dir := t.TempDir()
	r := NewRegistry(dir, "https://example.com/registry.git")

	t.Run("save and load registry", func(t *testing.T) {
		// Create a test registry
		registry := &TemplateRegistry{
			Name:        "Test Registry",
			Description: "A test registry",
			Entries: []*git.RepoInfo{
				{
					Name:        "template1@1.0.0",
					Description: "Test template 1",
					Git:         "https://example.com/template1.git",
				},
				{
					Name:        "template2@2.0.0",
					Description: "Test template 2",
					Git:         "https://example.com/template2.git",
				},
			},
		}

		r.Registry = registry

		// Save
		err := r.Save()
		require.NoError(t, err)

		// Verify file exists
		registryFile := filepath.Join(dir, "registry.json")
		assert.True(t, foundation.IsFile(registryFile))

		// Load
		r2 := NewRegistry(dir, "https://example.com/registry.git")
		err = r2.Load()
		require.NoError(t, err)

		// Verify loaded data
		assert.Equal(t, registry.Name, r2.Registry.Name)
		assert.Equal(t, registry.Description, r2.Registry.Description)
		assert.Len(t, r2.Registry.Entries, 2)
	})

	t.Run("load non-existent registry returns error", func(t *testing.T) {
		emptyDir := t.TempDir()
		r := NewRegistry(emptyDir, "https://example.com/registry.git")

		err := r.Load()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "registry file not found")
	})

	t.Run("load handles windows paths", func(t *testing.T) {
		// Create registry with windows-style paths
		registry := &TemplateRegistry{
			Name: "Test",
			Entries: []*git.RepoInfo{
				{
					Name: "path\\with\\backslashes",
					Git:  "https://example.com/test.git",
				},
			},
		}

		// Save
		data, err := json.Marshal(registry)
		require.NoError(t, err)

		registryFile := filepath.Join(dir, "registry.json")
		err = os.WriteFile(registryFile, data, 0644)
		require.NoError(t, err)

		// Load
		r := NewRegistry(dir, "https://example.com/registry.git")
		err = r.Load()
		require.NoError(t, err)

		// Verify backslashes are converted to forward slashes
		assert.Equal(t, "path/with/backslashes", r.Registry.Entries[0].Name)
	})
}

func TestRegistryGet(t *testing.T) {
	dir := t.TempDir()
	r := NewRegistry(dir, "https://example.com/registry.git")

	// Setup test registry
	registry := &TemplateRegistry{
		Name: "Test",
		Entries: []*git.RepoInfo{
			{
				Name: "template1",
				Git:  "https://example.com/template1.git",
			},
			{
				Name: "template2",
				Git:  "https://example.com/template2.git",
			},
		},
	}

	// Save registry
	r.Registry = registry
	err := r.Save()
	require.NoError(t, err)

	t.Run("get existing template", func(t *testing.T) {
		// Reset registry to force load
		r.Registry = nil

		info, err := r.Get("template1")
		require.NoError(t, err)
		assert.Equal(t, "template1", info.Name)
	})

	t.Run("get template with version suffix", func(t *testing.T) {
		r.Registry = registry

		info, err := r.Get("template1@1.0.0")
		require.NoError(t, err)
		assert.Equal(t, "template1", info.Name)
	})

	t.Run("get non-existent template", func(t *testing.T) {
		r.Registry = registry

		info, err := r.Get("nonexistent")
		assert.Error(t, err)
		assert.Nil(t, info)
		assert.Contains(t, err.Error(), "not found")
	})
}

func TestRegistryList(t *testing.T) {
	dir := t.TempDir()
	r := NewRegistry(dir, "https://example.com/registry.git")

	t.Run("list templates", func(t *testing.T) {
		registry := &TemplateRegistry{
			Name: "Test",
			Entries: []*git.RepoInfo{
				{Name: "template1"},
				{Name: "template2"},
				{Name: "template3"},
			},
		}

		r.Registry = registry
		err := r.Save()
		require.NoError(t, err)

		// Reset to force load
		r.Registry = nil

		entries, err := r.List()
		require.NoError(t, err)
		assert.Len(t, entries, 3)
	})
}

func TestRegistrySearch(t *testing.T) {
	dir := t.TempDir()
	r := NewRegistry(dir, "https://example.com/registry.git")

	registry := &TemplateRegistry{
		Name: "Test",
		Entries: []*git.RepoInfo{
			{Name: "cpp-template"},
			{Name: "python-template"},
			{Name: "go-template"},
		},
	}

	r.Registry = registry
	err := r.Save()
	require.NoError(t, err)

	t.Run("search with pattern", func(t *testing.T) {
		r.Registry = registry

		results, err := r.Search("python")
		require.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "python-template", results[0].Name)
	})

	t.Run("search with empty pattern returns all", func(t *testing.T) {
		r.Registry = registry

		results, err := r.Search("")
		require.NoError(t, err)
		assert.Len(t, results, 3)
	})

	t.Run("search with no match returns empty", func(t *testing.T) {
		r.Registry = registry

		results, err := r.Search("nonexistent")
		require.NoError(t, err)
		assert.Nil(t, results)
	})
}

func TestRegistryFixRepoId(t *testing.T) {
	dir := t.TempDir()
	r := NewRegistry(dir, "https://example.com/registry.git")

	registry := &TemplateRegistry{
		Name: "Test",
		Entries: []*git.RepoInfo{
			{
				Name: "template1",
				Latest: git.VersionInfo{
					Name: "v1.2.3",
				},
			},
		},
	}

	r.Registry = registry
	err := r.Save()
	require.NoError(t, err)

	t.Run("fix repo id without version", func(t *testing.T) {
		r.Registry = registry

		fixed, err := r.FixRepoId("template1")
		require.NoError(t, err)
		assert.Equal(t, "template1@v1.2.3", fixed)
	})

	t.Run("fix repo id with latest", func(t *testing.T) {
		r.Registry = registry

		fixed, err := r.FixRepoId("template1@latest")
		require.NoError(t, err)
		assert.Equal(t, "template1@v1.2.3", fixed)
	})

	t.Run("keep specific version", func(t *testing.T) {
		r.Registry = registry

		fixed, err := r.FixRepoId("template1@v1.0.0")
		require.NoError(t, err)
		assert.Equal(t, "template1@v1.0.0", fixed)
	})

	t.Run("non-existent template returns error", func(t *testing.T) {
		r.Registry = registry

		fixed, err := r.FixRepoId("nonexistent")
		assert.Error(t, err)
		assert.Empty(t, fixed)
	})
}

// Note: Tests for Update, Reset, and ensureRegistry require git operations
// and would need mocking or integration test setup.
