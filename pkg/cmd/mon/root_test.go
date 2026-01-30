package mon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand(t *testing.T) {
	t.Run("creates root monitor command", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "monitor", cmd.Use)
		assert.Contains(t, cmd.Aliases, "mon")
		assert.Contains(t, cmd.Aliases, "m")
		assert.Contains(t, cmd.Short, "monitor")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Equal(t, []string{"mon", "m"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Contains(t, cmd.Long, "monitor")
		assert.Contains(t, cmd.Long, "API")
	})

	t.Run("adds feed subcommand", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.True(t, cmd.HasSubCommands())

		// Find feed subcommand
		feedCmd, _, err := cmd.Find([]string{"feed"})
		assert.NoError(t, err)
		assert.NotNil(t, feedCmd)
		assert.Equal(t, "feed", feedCmd.Use)
	})

	t.Run("adds run subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find run subcommand
		runCmd, _, err := cmd.Find([]string{"run"})
		assert.NoError(t, err)
		assert.NotNil(t, runCmd)
		assert.Equal(t, "run", runCmd.Use)
	})

	t.Run("has both subcommands", func(t *testing.T) {
		cmd := NewRootCommand()
		subcommands := cmd.Commands()

		assert.Len(t, subcommands, 2)

		// Check that we have feed and run subcommands
		subcommandNames := make([]string, 0, len(subcommands))
		for _, subcmd := range subcommands {
			subcommandNames = append(subcommandNames, subcmd.Use)
		}

		assert.Contains(t, subcommandNames, "feed")
		assert.Contains(t, subcommandNames, "run")
	})

	t.Run("run subcommand has r alias", func(t *testing.T) {
		cmd := NewRootCommand()
		runCmd, _, err := cmd.Find([]string{"r"})
		assert.NoError(t, err)
		assert.NotNil(t, runCmd)
		assert.Equal(t, "run", runCmd.Use)
	})

	t.Run("run subcommand has start alias", func(t *testing.T) {
		cmd := NewRootCommand()
		runCmd, _, err := cmd.Find([]string{"start"})
		assert.NoError(t, err)
		assert.NotNil(t, runCmd)
		assert.Equal(t, "run", runCmd.Use)
	})
}
