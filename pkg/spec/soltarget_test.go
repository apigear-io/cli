package spec

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutionTargetGetOutputDir(t *testing.T) {
	t.Run("joins root dir with output path", func(t *testing.T) {
		target := &SolutionTarget{
			Name:   "test-target",
			Output: "output/generated",
		}

		rootDir := "/project"
		outputDir := target.GetOutputDir(rootDir)

		expected := filepath.Join(rootDir, "output/generated")
		assert.Equal(t, expected, outputDir)
	})

	t.Run("handles absolute output path", func(t *testing.T) {
		target := &SolutionTarget{
			Name:   "test-target",
			Output: "/absolute/output",
		}

		rootDir := "/project"
		outputDir := target.GetOutputDir(rootDir)

		// When output is absolute, Join returns the absolute path
		assert.Equal(t, "/absolute/output", outputDir)
	})

	t.Run("handles empty root dir", func(t *testing.T) {
		target := &SolutionTarget{
			Name:   "test-target",
			Output: "output",
		}

		outputDir := target.GetOutputDir("")
		assert.Equal(t, "output", outputDir)
	})

	t.Run("handles nested output paths", func(t *testing.T) {
		target := &SolutionTarget{
			Name:   "test-target",
			Output: "a/b/c/output",
		}

		rootDir := "/project"
		outputDir := target.GetOutputDir(rootDir)

		expected := filepath.Join(rootDir, "a/b/c/output")
		assert.Equal(t, expected, outputDir)
	})
}

func TestSolutionTargetDependencies(t *testing.T) {
	t.Run("returns empty dependencies when not computed", func(t *testing.T) {
		target := &SolutionTarget{
			Name: "test-target",
		}

		// Should return empty slice when not computed
		deps := target.Dependencies()
		assert.Empty(t, deps)
	})

	t.Run("returns dependencies after computation", func(t *testing.T) {
		target := &SolutionTarget{
			Name:     "test-target",
			computed: true,
			dependencies: []string{
				"module1.yaml",
				"module2.yaml",
			},
		}

		deps := target.Dependencies()
		assert.Len(t, deps, 2)
		assert.Contains(t, deps, "module1.yaml")
		assert.Contains(t, deps, "module2.yaml")
	})
}

func TestSolutionTargetExpandedInputs(t *testing.T) {
	t.Run("returns empty expanded inputs when not computed", func(t *testing.T) {
		target := &SolutionTarget{
			Name: "test-target",
		}

		// Should return empty slice when not computed
		inputs := target.ExpandedInputs()
		assert.Empty(t, inputs)
	})

	t.Run("returns expanded inputs after computation", func(t *testing.T) {
		target := &SolutionTarget{
			Name:     "test-target",
			computed: true,
			expandedInputs: []string{
				"expanded1.yaml",
				"expanded2.yaml",
			},
		}

		inputs := target.ExpandedInputs()
		assert.Len(t, inputs, 2)
		assert.Contains(t, inputs, "expanded1.yaml")
		assert.Contains(t, inputs, "expanded2.yaml")
	})
}

func TestSolutionTargetComputeImports(t *testing.T) {
	t.Run("initializes empty maps when imports is nil", func(t *testing.T) {
		target := &SolutionTarget{
			Name: "test-target",
		}

		err := target.computeImports()
		assert.NoError(t, err)
		assert.NotNil(t, target.Imports)
		assert.NotNil(t, target.MetaImports)
		assert.Empty(t, target.Imports)
		assert.Empty(t, target.MetaImports)
	})

	t.Run("reads import files", func(t *testing.T) {
		// Create a temporary import file
		dir := t.TempDir()
		importFile := filepath.Join(dir, "import.json")
		importData := `{"key": "value", "number": 42}`
		err := os.WriteFile(importFile, []byte(importData), 0644)
		assert.NoError(t, err)

		target := &SolutionTarget{
			Name: "test-target",
			Imports: []string{
				importFile,
			},
		}

		err = target.computeImports()
		assert.NoError(t, err)

		// Check that meta imports were populated
		assert.NotNil(t, target.MetaImports)
		assert.Equal(t, "value", target.MetaImports["key"])
		assert.Equal(t, float64(42), target.MetaImports["number"])
	})

	t.Run("handles non-existent import files gracefully", func(t *testing.T) {
		target := &SolutionTarget{
			Name: "test-target",
			Imports: []string{
				"/nonexistent/import.json",
			},
		}

		// Should not error, just log warning
		err := target.computeImports()
		assert.NoError(t, err)
	})

	t.Run("handles multiple import files", func(t *testing.T) {
		dir := t.TempDir()

		// Create first import file
		import1 := filepath.Join(dir, "import1.json")
		err := os.WriteFile(import1, []byte(`{"key1": "value1"}`), 0644)
		assert.NoError(t, err)

		// Create second import file
		import2 := filepath.Join(dir, "import2.json")
		err = os.WriteFile(import2, []byte(`{"key2": "value2"}`), 0644)
		assert.NoError(t, err)

		target := &SolutionTarget{
			Name: "test-target",
			Imports: []string{
				import1,
				import2,
			},
		}

		err = target.computeImports()
		assert.NoError(t, err)

		// Both imports should be merged
		assert.Equal(t, "value1", target.MetaImports["key1"])
		assert.Equal(t, "value2", target.MetaImports["key2"])
	})

	t.Run("later imports override earlier ones", func(t *testing.T) {
		dir := t.TempDir()

		// Create first import file
		import1 := filepath.Join(dir, "import1.json")
		err := os.WriteFile(import1, []byte(`{"shared": "first"}`), 0644)
		assert.NoError(t, err)

		// Create second import file with same key
		import2 := filepath.Join(dir, "import2.json")
		err = os.WriteFile(import2, []byte(`{"shared": "second"}`), 0644)
		assert.NoError(t, err)

		target := &SolutionTarget{
			Name: "test-target",
			Imports: []string{
				import1,
				import2,
			},
		}

		err = target.computeImports()
		assert.NoError(t, err)

		// Second import should override first
		assert.Equal(t, "second", target.MetaImports["shared"])
	})
}

func TestSolutionTargetValidate(t *testing.T) {
	t.Run("fails validation when output is empty", func(t *testing.T) {
		doc := &SolutionDoc{
			Name:    "test-solution",
			RootDir: "/test",
		}

		target := &SolutionTarget{
			Name:     "test-target",
			Output:   "", // Missing output
			Template: "test-template",
		}

		err := target.Validate(doc)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "output is required")
	})

	t.Run("fails validation when template is empty", func(t *testing.T) {
		doc := &SolutionDoc{
			Name:    "test-solution",
			RootDir: "/test",
		}

		target := &SolutionTarget{
			Name:     "test-target",
			Output:   "output",
			Template: "", // Missing template
		}

		err := target.Validate(doc)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "template is required")
	})

	t.Run("initializes nil meta to empty map", func(t *testing.T) {
		doc := &SolutionDoc{
			Name:    "test-solution",
			RootDir: "/test",
		}

		target := &SolutionTarget{
			Name:     "test-target",
			Output:   "output",
			Template: "template",
			Meta:     nil,
		}

		// Will fail during compute phase but Meta should be initialized
		_ = target.Validate(doc)
		assert.NotNil(t, target.Meta)
	})

	t.Run("initializes nil inputs to empty slice", func(t *testing.T) {
		doc := &SolutionDoc{
			Name:    "test-solution",
			RootDir: "/test",
		}

		target := &SolutionTarget{
			Name:     "test-target",
			Output:   "output",
			Template: "template",
			Inputs:   nil,
		}

		// Will fail during compute phase but Inputs should be initialized
		_ = target.Validate(doc)
		assert.NotNil(t, target.Inputs)
		assert.Empty(t, target.Inputs)
	})

	t.Run("initializes nil features to default 'all'", func(t *testing.T) {
		doc := &SolutionDoc{
			Name:    "test-solution",
			RootDir: "/test",
		}

		target := &SolutionTarget{
			Name:     "test-target",
			Output:   "output",
			Template: "template",
			Features: nil,
		}

		// Will fail during compute phase but Features should be initialized
		_ = target.Validate(doc)
		assert.NotNil(t, target.Features)
		assert.Equal(t, []string{"all"}, target.Features)
	})

	t.Run("fails when template dir not found", func(t *testing.T) {
		dir := t.TempDir()

		doc := &SolutionDoc{
			Name:    "test-solution",
			RootDir: dir,
		}

		target := &SolutionTarget{
			Name:     "test-target",
			Output:   "output",
			Template: "nonexistent-template",
		}

		err := target.Validate(doc)
		assert.Error(t, err)
		// Error could be about template dir not found or GetOrInstallTemplate failure
		// Both are acceptable as they indicate missing template
	})
}
