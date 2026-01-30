package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand(t *testing.T) {
	t.Run("creates root spec command", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "spec", cmd.Use)
		assert.Contains(t, cmd.Aliases, "s")
		assert.Contains(t, cmd.Short, "validate")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Equal(t, []string{"s"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Contains(t, cmd.Long, "file formats")
		assert.Contains(t, cmd.Long, "apigear")
	})

	t.Run("adds check subcommand", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.True(t, cmd.HasSubCommands())

		// Find check subcommand
		checkCmd, _, err := cmd.Find([]string{"check"})
		assert.NoError(t, err)
		assert.NotNil(t, checkCmd)
		assert.Equal(t, "check", checkCmd.Use)
	})

	t.Run("adds show subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find show subcommand (Use is "schema" but has "show" alias)
		showCmd, _, err := cmd.Find([]string{"schema"})
		assert.NoError(t, err)
		assert.NotNil(t, showCmd)
		assert.Equal(t, "schema", showCmd.Use)
	})

	t.Run("has both subcommands", func(t *testing.T) {
		cmd := NewRootCommand()
		subcommands := cmd.Commands()

		assert.Len(t, subcommands, 2)

		// Check that we have check and schema subcommands
		subcommandNames := make([]string, 0, len(subcommands))
		for _, subcmd := range subcommands {
			subcommandNames = append(subcommandNames, subcmd.Use)
		}

		assert.Contains(t, subcommandNames, "check")
		assert.Contains(t, subcommandNames, "schema")
	})

	t.Run("check subcommand has c alias", func(t *testing.T) {
		cmd := NewRootCommand()
		checkCmd, _, err := cmd.Find([]string{"c"})
		assert.NoError(t, err)
		assert.NotNil(t, checkCmd)
		assert.Equal(t, "check", checkCmd.Use)
	})

	t.Run("check subcommand has lint alias", func(t *testing.T) {
		cmd := NewRootCommand()
		checkCmd, _, err := cmd.Find([]string{"lint"})
		assert.NoError(t, err)
		assert.NotNil(t, checkCmd)
		assert.Equal(t, "check", checkCmd.Use)
	})

	t.Run("show subcommand has s alias", func(t *testing.T) {
		cmd := NewRootCommand()
		showCmd, _, err := cmd.Find([]string{"s"})
		assert.NoError(t, err)
		assert.NotNil(t, showCmd)
		assert.Equal(t, "schema", showCmd.Use)
	})

	t.Run("show subcommand has show alias", func(t *testing.T) {
		cmd := NewRootCommand()
		showCmd, _, err := cmd.Find([]string{"show"})
		assert.NoError(t, err)
		assert.NotNil(t, showCmd)
		assert.Equal(t, "schema", showCmd.Use)
	})

	t.Run("show subcommand has view alias", func(t *testing.T) {
		cmd := NewRootCommand()
		showCmd, _, err := cmd.Find([]string{"view"})
		assert.NoError(t, err)
		assert.NotNil(t, showCmd)
		assert.Equal(t, "schema", showCmd.Use)
	})
}
