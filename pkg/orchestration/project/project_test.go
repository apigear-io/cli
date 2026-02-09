package project

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeDocumentName(t *testing.T) {
	tests := []struct {
		name     string
		docType  string
		docName  string
		expected string
	}{
		{"module document", "module", "demo", "demo.module.yaml"},
		{"solution document", "solution", "demo", "demo.solution.yaml"},
		{"scenario document", "scenario", "demo", "demo.scenario.yaml"},
		{"invalid type", "invalid", "demo", ""},
		{"module with hyphen", "module", "my-module", "my-module.module.yaml"},
		{"solution with underscore", "solution", "my_solution", "my_solution.solution.yaml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MakeDocumentName(tt.docType, tt.docName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInitProject(t *testing.T) {
	t.Run("creates new project with apigear directory", func(t *testing.T) {
		dir := t.TempDir()
		projectDir := filepath.Join(dir, "test-project")

		info, err := InitProject(projectDir)
		require.NoError(t, err)
		require.NotNil(t, info)

		// Verify apigear directory exists
		apigearDir := filepath.Join(projectDir, "apigear")
		assert.True(t, foundation.IsDir(apigearDir))

		// Verify project info
		assert.Equal(t, "test-project", info.Name)
		assert.Equal(t, projectDir, info.Path)
		assert.NotEmpty(t, info.Documents)
	})

	t.Run("creates demo files", func(t *testing.T) {
		dir := t.TempDir()
		projectDir := filepath.Join(dir, "demo-project")

		info, err := InitProject(projectDir)
		require.NoError(t, err)

		apigearDir := filepath.Join(projectDir, "apigear")

		// Check demo files exist
		demoModule := filepath.Join(apigearDir, "demo.module.yaml")
		assert.True(t, foundation.IsFile(demoModule))

		demoIdl := filepath.Join(apigearDir, "demo.module.idl")
		assert.True(t, foundation.IsFile(demoIdl))

		demoSolution := filepath.Join(apigearDir, "demo.solution.yaml")
		assert.True(t, foundation.IsFile(demoSolution))

		demoSim := filepath.Join(apigearDir, "demo.sim.js")
		assert.True(t, foundation.IsFile(demoSim))

		// Verify documents are listed
		assert.Len(t, info.Documents, 4)
	})

	t.Run("initializes in existing directory", func(t *testing.T) {
		dir := t.TempDir()

		info, err := InitProject(dir)
		require.NoError(t, err)
		assert.NotNil(t, info)

		apigearDir := filepath.Join(dir, "apigear")
		assert.True(t, foundation.IsDir(apigearDir))
	})

	t.Run("handles existing apigear directory", func(t *testing.T) {
		dir := t.TempDir()

		// Create apigear directory first
		apigearDir := filepath.Join(dir, "apigear")
		err := os.Mkdir(apigearDir, 0755)
		require.NoError(t, err)

		// Should not fail
		info, err := InitProject(dir)
		require.NoError(t, err)
		assert.NotNil(t, info)
	})
}

func TestOpenProject(t *testing.T) {
	t.Run("opens existing project", func(t *testing.T) {
		dir := t.TempDir()

		// First init a project
		_, err := InitProject(dir)
		require.NoError(t, err)

		// Now open it
		info, err := OpenProject(dir)
		require.NoError(t, err)
		require.NotNil(t, info)

		assert.Equal(t, filepath.Base(dir), info.Name)
		assert.Equal(t, dir, info.Path)
		assert.NotEmpty(t, info.Documents)
	})

	t.Run("fails for non-existent directory", func(t *testing.T) {
		info, err := OpenProject("/nonexistent/path")
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("fails for directory without apigear", func(t *testing.T) {
		dir := t.TempDir()

		info, err := OpenProject(dir)
		assert.Error(t, err)
		assert.Nil(t, info)
	})
}

func TestReadProject(t *testing.T) {
	t.Run("reads project with documents", func(t *testing.T) {
		dir := t.TempDir()

		// Init project first
		_, err := InitProject(dir)
		require.NoError(t, err)

		// Read the project
		info, err := ReadProject(dir)
		require.NoError(t, err)
		require.NotNil(t, info)

		assert.Equal(t, filepath.Base(dir), info.Name)
		assert.Equal(t, dir, info.Path)
		assert.Len(t, info.Documents, 4) // demo files

		// Verify document types
		for _, doc := range info.Documents {
			assert.NotEmpty(t, doc.Name)
			assert.NotEmpty(t, doc.Path)
			assert.Contains(t, []string{"module", "simulation", "solution"}, doc.Type)
		}
	})

	t.Run("reads project with custom documents", func(t *testing.T) {
		dir := t.TempDir()
		apigearDir := filepath.Join(dir, "apigear")
		err := os.MkdirAll(apigearDir, 0755)
		require.NoError(t, err)

		// Create custom documents
		customModule := filepath.Join(apigearDir, "custom.module.yaml")
		err = os.WriteFile(customModule, []byte("# custom module"), 0644)
		require.NoError(t, err)

		customSolution := filepath.Join(apigearDir, "custom.solution.yaml")
		err = os.WriteFile(customSolution, []byte("# custom solution"), 0644)
		require.NoError(t, err)

		// Read project
		info, err := ReadProject(dir)
		require.NoError(t, err)

		assert.Len(t, info.Documents, 2)
	})

	t.Run("sets current project", func(t *testing.T) {
		dir := t.TempDir()
		_, err := InitProject(dir)
		require.NoError(t, err)

		_, err = ReadProject(dir)
		require.NoError(t, err)

		// Check current project is set
		current := CurrentProject()
		assert.NotNil(t, current)
		assert.Equal(t, dir, current.Path)
	})

	t.Run("fails for non-existent directory", func(t *testing.T) {
		info, err := ReadProject("/nonexistent/path")
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("fails for directory without apigear", func(t *testing.T) {
		dir := t.TempDir()

		info, err := ReadProject(dir)
		assert.Error(t, err)
		assert.Nil(t, info)
	})
}

func TestGetProjectInfo(t *testing.T) {
	t.Run("gets project info", func(t *testing.T) {
		dir := t.TempDir()

		// Init project first
		_, err := InitProject(dir)
		require.NoError(t, err)

		// Get project info
		info, err := GetProjectInfo(dir)
		require.NoError(t, err)
		require.NotNil(t, info)

		assert.Equal(t, filepath.Base(dir), info.Name)
		assert.Equal(t, dir, info.Path)
	})
}

func TestAddDocument(t *testing.T) {
	t.Run("adds module document", func(t *testing.T) {
		dir := t.TempDir()

		// Init project
		_, err := InitProject(dir)
		require.NoError(t, err)

		// Add module document
		docPath, err := AddDocument(dir, "module", "custom")
		require.NoError(t, err)

		expectedPath := filepath.Join(dir, "apigear", "custom.module.yaml")
		assert.Equal(t, expectedPath, docPath)
		assert.True(t, foundation.IsFile(docPath))
	})

	t.Run("adds solution document", func(t *testing.T) {
		dir := t.TempDir()

		// Init project
		_, err := InitProject(dir)
		require.NoError(t, err)

		// Add solution document
		docPath, err := AddDocument(dir, "solution", "custom")
		require.NoError(t, err)

		expectedPath := filepath.Join(dir, "apigear", "custom.solution.yaml")
		assert.Equal(t, expectedPath, docPath)
		assert.True(t, foundation.IsFile(docPath))
	})

	t.Run("simulation type not supported by MakeDocumentName", func(t *testing.T) {
		dir := t.TempDir()

		// Init project
		_, err := InitProject(dir)
		require.NoError(t, err)

		// AddDocument with "simulation" type doesn't work because
		// MakeDocumentName returns empty string for "simulation"
		// This is a limitation in the current implementation
		docPath, err := AddDocument(dir, "simulation", "custom")
		assert.Error(t, err)
		assert.Empty(t, docPath)
	})

	t.Run("fails for invalid document type", func(t *testing.T) {
		dir := t.TempDir()

		_, err := InitProject(dir)
		require.NoError(t, err)

		docPath, err := AddDocument(dir, "invalid", "custom")
		assert.Error(t, err)
		assert.Empty(t, docPath)
		assert.Contains(t, err.Error(), "invalid document type")
	})

	t.Run("fails if document already exists", func(t *testing.T) {
		dir := t.TempDir()

		_, err := InitProject(dir)
		require.NoError(t, err)

		// Add document first time
		_, err = AddDocument(dir, "module", "test")
		require.NoError(t, err)

		// Try to add again
		_, err = AddDocument(dir, "module", "test")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "already exists")
	})
}

func TestCurrentProject(t *testing.T) {
	t.Run("returns current project after read", func(t *testing.T) {
		dir := t.TempDir()

		_, err := InitProject(dir)
		require.NoError(t, err)

		_, err = ReadProject(dir)
		require.NoError(t, err)

		current := CurrentProject()
		assert.NotNil(t, current)
		assert.Equal(t, dir, current.Path)
	})

	t.Run("returns nil before any project is opened", func(t *testing.T) {
		// Reset current project
		currentProject = nil

		current := CurrentProject()
		assert.Nil(t, current)
	})
}

// Note: Tests for ImportProject, PackProject, OpenEditor, OpenStudio, and RecentProjectInfos
// are excluded as they require external dependencies (git, zip, exec commands, config)
// that should be mocked or tested in integration tests.
