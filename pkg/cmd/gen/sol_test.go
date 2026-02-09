package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSolutionCommand(t *testing.T) {
	t.Run("creates solution command", func(t *testing.T) {
		cmd := NewSolutionCommand()
		assert.NotNil(t, cmd)
		assert.Contains(t, cmd.Use, "solution")
		assert.Contains(t, cmd.Aliases, "sol")
		assert.Contains(t, cmd.Aliases, "s")
		assert.Contains(t, cmd.Short, "solution document")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewSolutionCommand()
		assert.Equal(t, []string{"sol", "s"}, cmd.Aliases)
	})

	t.Run("requires exactly one argument", func(t *testing.T) {
		cmd := NewSolutionCommand()
		assert.NotNil(t, cmd.Args)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewSolutionCommand()
		assert.Contains(t, cmd.Long, "solution")
		assert.Contains(t, cmd.Long, "yaml")
		assert.Contains(t, cmd.Long, "layer")
	})

	t.Run("has watch flag", func(t *testing.T) {
		cmd := NewSolutionCommand()
		flag := cmd.Flags().Lookup("watch")
		assert.NotNil(t, flag)
		assert.Equal(t, "false", flag.DefValue)
		assert.Contains(t, flag.Usage, "watch")
	})

	t.Run("has force flag", func(t *testing.T) {
		cmd := NewSolutionCommand()
		flag := cmd.Flags().Lookup("force")
		assert.NotNil(t, flag)
		assert.Equal(t, "false", flag.DefValue)
		assert.Contains(t, flag.Usage, "force")
	})

	t.Run("watch flag defaults to false", func(t *testing.T) {
		cmd := NewSolutionCommand()
		watch, err := cmd.Flags().GetBool("watch")
		assert.NoError(t, err)
		assert.False(t, watch)
	})

	t.Run("force flag defaults to false", func(t *testing.T) {
		cmd := NewSolutionCommand()
		force, err := cmd.Flags().GetBool("force")
		assert.NoError(t, err)
		assert.False(t, force)
	})

	t.Run("accepts solution file argument", func(t *testing.T) {
		cmd := NewSolutionCommand()
		cmd.SetArgs([]string{"test.solution.yaml"})
		err := cmd.ParseFlags([]string{})
		assert.NoError(t, err)
	})

	t.Run("accepts watch flag", func(t *testing.T) {
		cmd := NewSolutionCommand()
		cmd.SetArgs([]string{"--watch", "test.solution.yaml"})
		err := cmd.ParseFlags([]string{"--watch"})
		assert.NoError(t, err)

		watch, err := cmd.Flags().GetBool("watch")
		assert.NoError(t, err)
		assert.True(t, watch)
	})

	t.Run("accepts force flag", func(t *testing.T) {
		cmd := NewSolutionCommand()
		cmd.SetArgs([]string{"--force", "test.solution.yaml"})
		err := cmd.ParseFlags([]string{"--force"})
		assert.NoError(t, err)

		force, err := cmd.Flags().GetBool("force")
		assert.NoError(t, err)
		assert.True(t, force)
	})

	t.Run("accepts both flags", func(t *testing.T) {
		cmd := NewSolutionCommand()
		err := cmd.ParseFlags([]string{"--watch", "--force"})
		assert.NoError(t, err)

		watch, err := cmd.Flags().GetBool("watch")
		assert.NoError(t, err)
		assert.True(t, watch)

		force, err := cmd.Flags().GetBool("force")
		assert.NoError(t, err)
		assert.True(t, force)
	})
}
