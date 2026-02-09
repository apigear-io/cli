package mon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServerCommand(t *testing.T) {
	t.Run("creates run command", func(t *testing.T) {
		cmd := NewServerCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "run", cmd.Use)
		assert.Contains(t, cmd.Aliases, "r")
		assert.Contains(t, cmd.Aliases, "start")
		assert.Contains(t, cmd.Short, "monitor server")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewServerCommand()
		assert.Equal(t, []string{"r", "start"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewServerCommand()
		assert.Contains(t, cmd.Long, "monitor server")
		assert.Contains(t, cmd.Long, "HTTP")
		assert.Contains(t, cmd.Long, "API calls")
	})

	t.Run("has addr flag", func(t *testing.T) {
		cmd := NewServerCommand()
		flag := cmd.Flags().Lookup("addr")
		assert.NotNil(t, flag)
		assert.Equal(t, "a", flag.Shorthand)
		assert.Equal(t, "127.0.0.1:5555", flag.DefValue)
		assert.Contains(t, flag.Usage, "address")
	})

	t.Run("addr flag defaults to 127.0.0.1:5555", func(t *testing.T) {
		cmd := NewServerCommand()
		addr, err := cmd.Flags().GetString("addr")
		assert.NoError(t, err)
		assert.Equal(t, "127.0.0.1:5555", addr)
	})

	t.Run("accepts addr flag", func(t *testing.T) {
		cmd := NewServerCommand()
		err := cmd.ParseFlags([]string{"--addr", "0.0.0.0:8080"})
		assert.NoError(t, err)

		addr, err := cmd.Flags().GetString("addr")
		assert.NoError(t, err)
		assert.Equal(t, "0.0.0.0:8080", addr)
	})

	t.Run("accepts short addr flag", func(t *testing.T) {
		cmd := NewServerCommand()
		err := cmd.ParseFlags([]string{"-a", "localhost:9999"})
		assert.NoError(t, err)

		addr, err := cmd.Flags().GetString("addr")
		assert.NoError(t, err)
		assert.Equal(t, "localhost:9999", addr)
	})

	t.Run("has RunE function", func(t *testing.T) {
		cmd := NewServerCommand()
		assert.NotNil(t, cmd.RunE)
	})
}
