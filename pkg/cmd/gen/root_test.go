package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand(t *testing.T) {
	t.Run("creates root generate command", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "generate", cmd.Use)
		assert.Contains(t, cmd.Aliases, "gen")
		assert.Contains(t, cmd.Aliases, "g")
		assert.Contains(t, cmd.Short, "Generate code")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Equal(t, []string{"gen", "g"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Contains(t, cmd.Long, "generate")
		assert.Contains(t, cmd.Long, "API")
		assert.Contains(t, cmd.Long, "templates")
	})

	t.Run("adds expert subcommand", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.True(t, cmd.HasSubCommands())

		// Find expert subcommand
		expertCmd, _, err := cmd.Find([]string{"expert"})
		assert.NoError(t, err)
		assert.NotNil(t, expertCmd)
		assert.Equal(t, "expert", expertCmd.Use)
	})

	t.Run("adds solution subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find solution subcommand
		solCmd, _, err := cmd.Find([]string{"solution"})
		assert.NoError(t, err)
		assert.NotNil(t, solCmd)
		assert.Contains(t, solCmd.Use, "solution")
	})

	t.Run("has both subcommands", func(t *testing.T) {
		cmd := NewRootCommand()
		subcommands := cmd.Commands()

		assert.Len(t, subcommands, 2)

		// Check that we have expert and solution
		subcommandNames := make([]string, 0, len(subcommands))
		for _, subcmd := range subcommands {
			subcommandNames = append(subcommandNames, subcmd.Use)
		}

		assert.Contains(t, subcommandNames, "expert")
		assert.True(t, containsSolution(subcommandNames))
	})

	t.Run("expert subcommand has x alias", func(t *testing.T) {
		cmd := NewRootCommand()
		expertCmd, _, err := cmd.Find([]string{"x"})
		assert.NoError(t, err)
		assert.NotNil(t, expertCmd)
		assert.Equal(t, "expert", expertCmd.Use)
	})

	t.Run("solution subcommand has sol alias", func(t *testing.T) {
		cmd := NewRootCommand()
		solCmd, _, err := cmd.Find([]string{"sol"})
		assert.NoError(t, err)
		assert.NotNil(t, solCmd)
		assert.Contains(t, solCmd.Use, "solution")
	})
}

// Helper function to check if any use string contains "solution"
func containsSolution(uses []string) bool {
	for _, use := range uses {
		if len(use) >= 8 && use[:8] == "solution" {
			return true
		}
	}
	return false
}
