package gen

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMust(t *testing.T) {
	t.Run("does nothing when error is nil", func(t *testing.T) {
		// Should not panic
		assert.NotPanics(t, func() {
			Must(nil)
		})
	})

	// Note: Cannot test the error case as it calls logging.Fatal which exits the process
}

func TestMakeSolution(t *testing.T) {
	t.Run("creates solution doc from options", func(t *testing.T) {
		options := &ExpertOptions{
			Inputs:      []string{"input1.yaml", "input2.yaml"},
			OutputDir:   "output",
			Features:    []string{"feature1", "feature2"},
			Force:       true,
			TemplateDir: "templates",
		}

		doc := MakeSolution(options)

		require.NotNil(t, doc)
		assert.NotEmpty(t, doc.RootDir) // Should be set to current working directory
		assert.Len(t, doc.Targets, 1)

		target := doc.Targets[0]
		assert.Equal(t, options.Inputs, target.Inputs)
		assert.Equal(t, options.OutputDir, target.Output)
		assert.Equal(t, options.TemplateDir, target.Template)
		assert.Equal(t, options.Features, target.Features)
		assert.Equal(t, options.Force, target.Force)
	})

	t.Run("creates solution with single input", func(t *testing.T) {
		options := &ExpertOptions{
			Inputs:      []string{"single.yaml"},
			OutputDir:   "out",
			Features:    []string{"all"},
			Force:       false,
			TemplateDir: "tpl",
		}

		doc := MakeSolution(options)

		require.NotNil(t, doc)
		require.Len(t, doc.Targets, 1)
		assert.Equal(t, []string{"single.yaml"}, doc.Targets[0].Inputs)
	})

	t.Run("handles empty features", func(t *testing.T) {
		options := &ExpertOptions{
			Inputs:      []string{"input.yaml"},
			OutputDir:   "output",
			Features:    []string{},
			TemplateDir: "templates",
		}

		doc := MakeSolution(options)

		require.NotNil(t, doc)
		require.Len(t, doc.Targets, 1)
		assert.Empty(t, doc.Targets[0].Features)
	})

	t.Run("sets force flag correctly", func(t *testing.T) {
		optionsTrue := &ExpertOptions{
			Inputs:      []string{"input.yaml"},
			OutputDir:   "output",
			Features:    []string{"all"},
			Force:       true,
			TemplateDir: "templates",
		}

		docTrue := MakeSolution(optionsTrue)
		assert.True(t, docTrue.Targets[0].Force)

		optionsFalse := &ExpertOptions{
			Inputs:      []string{"input.yaml"},
			OutputDir:   "output",
			Features:    []string{"all"},
			Force:       false,
			TemplateDir: "templates",
		}

		docFalse := MakeSolution(optionsFalse)
		assert.False(t, docFalse.Targets[0].Force)
	})
}

func TestNewExpertCommand(t *testing.T) {
	t.Run("creates expert command", func(t *testing.T) {
		cmd := NewExpertCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "expert", cmd.Use)
		assert.Contains(t, cmd.Aliases, "x")
		assert.Contains(t, cmd.Short, "expert mode")
	})

	t.Run("has x alias", func(t *testing.T) {
		cmd := NewExpertCommand()
		assert.Equal(t, []string{"x"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewExpertCommand()
		assert.Contains(t, cmd.Long, "expert mode")
		assert.Contains(t, cmd.Long, "solution document")
	})

	t.Run("has template flag", func(t *testing.T) {
		cmd := NewExpertCommand()
		flag := cmd.Flags().Lookup("template")
		assert.NotNil(t, flag)
		assert.Equal(t, "t", flag.Shorthand)
		assert.Equal(t, "tpl", flag.DefValue)
		assert.True(t, isRequired(cmd, "template"))
	})

	t.Run("has input flag", func(t *testing.T) {
		cmd := NewExpertCommand()
		flag := cmd.Flags().Lookup("input")
		assert.NotNil(t, flag)
		assert.Equal(t, "i", flag.Shorthand)
		assert.True(t, isRequired(cmd, "input"))
	})

	t.Run("has output flag", func(t *testing.T) {
		cmd := NewExpertCommand()
		flag := cmd.Flags().Lookup("output")
		assert.NotNil(t, flag)
		assert.Equal(t, "o", flag.Shorthand)
		assert.Equal(t, "out", flag.DefValue)
		assert.True(t, isRequired(cmd, "output"))
	})

	t.Run("has features flag", func(t *testing.T) {
		cmd := NewExpertCommand()
		flag := cmd.Flags().Lookup("features")
		assert.NotNil(t, flag)
		assert.Equal(t, "f", flag.Shorthand)
	})

	t.Run("has force flag", func(t *testing.T) {
		cmd := NewExpertCommand()
		flag := cmd.Flags().Lookup("force")
		assert.NotNil(t, flag)
		assert.Equal(t, "false", flag.DefValue)
	})

	t.Run("has watch flag", func(t *testing.T) {
		cmd := NewExpertCommand()
		flag := cmd.Flags().Lookup("watch")
		assert.NotNil(t, flag)
		assert.Equal(t, "false", flag.DefValue)
	})

	t.Run("template flag is required", func(t *testing.T) {
		cmd := NewExpertCommand()
		assert.True(t, isRequired(cmd, "template"))
	})

	t.Run("input flag is required", func(t *testing.T) {
		cmd := NewExpertCommand()
		assert.True(t, isRequired(cmd, "input"))
	})

	t.Run("output flag is required", func(t *testing.T) {
		cmd := NewExpertCommand()
		assert.True(t, isRequired(cmd, "output"))
	})

	t.Run("features flag is optional", func(t *testing.T) {
		cmd := NewExpertCommand()
		assert.False(t, isRequired(cmd, "features"))
	})

	t.Run("accepts all flags", func(t *testing.T) {
		cmd := NewExpertCommand()
		err := cmd.ParseFlags([]string{
			"--template", "my-template",
			"--input", "input1.yaml",
			"--input", "input2.yaml",
			"--output", "my-output",
			"--features", "feature1",
			"--features", "feature2",
			"--force",
			"--watch",
		})
		assert.NoError(t, err)

		template, _ := cmd.Flags().GetString("template")
		assert.Equal(t, "my-template", template)

		inputs, _ := cmd.Flags().GetStringSlice("input")
		assert.Contains(t, inputs, "input1.yaml")
		assert.Contains(t, inputs, "input2.yaml")

		output, _ := cmd.Flags().GetString("output")
		assert.Equal(t, "my-output", output)

		features, _ := cmd.Flags().GetStringSlice("features")
		assert.Contains(t, features, "feature1")
		assert.Contains(t, features, "feature2")

		force, _ := cmd.Flags().GetBool("force")
		assert.True(t, force)

		watch, _ := cmd.Flags().GetBool("watch")
		assert.True(t, watch)
	})
}

// Helper function to check if a flag is required
func isRequired(cmd *cobra.Command, flagName string) bool {
	flag := cmd.Flags().Lookup(flagName)
	if flag == nil {
		return false
	}
	// Check annotations for required flags
	annotations := flag.Annotations
	if annotations != nil {
		if _, ok := annotations[cobra.BashCompOneRequiredFlag]; ok {
			return true
		}
	}
	// Cobra marks required flags differently, let's check the Required field
	// Unfortunately, the Required field is not exported, so we check if validation fails
	return false // We can't fully test this without executing
}
