package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand(t *testing.T) {
	t.Run("creates root config command", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "config", cmd.Use)
		assert.Contains(t, cmd.Aliases, "cfg")
		assert.Contains(t, cmd.Aliases, "c")
		assert.Contains(t, cmd.Short, "config")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Equal(t, []string{"cfg", "c"}, cmd.Aliases)
	})

	t.Run("adds info subcommand", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.True(t, cmd.HasSubCommands())

		// Find info subcommand
		infoCmd, _, err := cmd.Find([]string{"info"})
		assert.NoError(t, err)
		assert.NotNil(t, infoCmd)
		assert.Equal(t, "info", infoCmd.Use)
	})

	t.Run("adds get subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find get subcommand
		getCmd, _, err := cmd.Find([]string{"get"})
		assert.NoError(t, err)
		assert.NotNil(t, getCmd)
		assert.Equal(t, "get", getCmd.Use)
	})

	t.Run("adds env subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find env subcommand
		envCmd, _, err := cmd.Find([]string{"env"})
		assert.NoError(t, err)
		assert.NotNil(t, envCmd)
		assert.Equal(t, "env", envCmd.Use)
	})

	t.Run("has all three subcommands", func(t *testing.T) {
		cmd := NewRootCommand()
		subcommands := cmd.Commands()

		assert.Len(t, subcommands, 3)

		// Check that we have info, get, and env
		subcommandNames := make([]string, 0, len(subcommands))
		for _, subcmd := range subcommands {
			subcommandNames = append(subcommandNames, subcmd.Use)
		}

		assert.Contains(t, subcommandNames, "info")
		assert.Contains(t, subcommandNames, "get")
		assert.Contains(t, subcommandNames, "env")
	})

	t.Run("can execute info subcommand via root", func(t *testing.T) {
		cmd := NewRootCommand()
		out := ExecuteCmd(t, cmd, "info")

		assert.Contains(t, out, "info:")
		assert.Contains(t, out, "config file:")
	})

	t.Run("can execute get subcommand via root", func(t *testing.T) {
		cmd := NewRootCommand()
		out := ExecuteCmd(t, cmd, "get")

		assert.Contains(t, out, "settings")
	})

	t.Run("can execute env subcommand via root", func(t *testing.T) {
		cmd := NewRootCommand()
		out := ExecuteCmd(t, cmd, "env")

		assert.Contains(t, out, "APIGEAR_")
	})
}
