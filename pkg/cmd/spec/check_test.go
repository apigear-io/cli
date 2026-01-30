package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCheckCommand(t *testing.T) {
	t.Run("creates check command", func(t *testing.T) {
		cmd := NewCheckCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "check", cmd.Use)
		assert.Contains(t, cmd.Aliases, "c")
		assert.Contains(t, cmd.Aliases, "lint")
		assert.Contains(t, cmd.Short, "Check")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewCheckCommand()
		assert.Equal(t, []string{"c", "lint"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewCheckCommand()
		assert.Contains(t, cmd.Long, "Check")
		assert.Contains(t, cmd.Long, "documents")
		assert.Contains(t, cmd.Long, "errors")
	})

	t.Run("requires exactly one argument", func(t *testing.T) {
		cmd := NewCheckCommand()
		assert.NotNil(t, cmd.Args)

		// Test with no arguments
		err := cmd.Args(cmd, []string{})
		assert.Error(t, err)

		// Test with one argument (should pass)
		err = cmd.Args(cmd, []string{"file.yaml"})
		assert.NoError(t, err)

		// Test with two arguments
		err = cmd.Args(cmd, []string{"file1.yaml", "file2.yaml"})
		assert.Error(t, err)
	})

	t.Run("has RunE function", func(t *testing.T) {
		cmd := NewCheckCommand()
		assert.NotNil(t, cmd.RunE)
	})

	t.Run("accepts single file argument", func(t *testing.T) {
		cmd := NewCheckCommand()
		cmd.SetArgs([]string{"test.yaml"})
		err := cmd.ParseFlags([]string{})
		assert.NoError(t, err)
	})
}
