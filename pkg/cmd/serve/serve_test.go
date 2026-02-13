package serve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServeCommand(t *testing.T) {
	cmd := NewServeCommand()

	assert.Equal(t, "serve", cmd.Use)
	assert.Contains(t, cmd.Aliases, "server")
	assert.Contains(t, cmd.Aliases, "s")
	assert.NotNil(t, cmd.RunE)

	// Verify flags
	addrFlag := cmd.Flags().Lookup("addr")
	assert.NotNil(t, addrFlag)

	hostFlag := cmd.Flags().Lookup("host")
	assert.NotNil(t, hostFlag)
	assert.Equal(t, "localhost", hostFlag.DefValue)

	portFlag := cmd.Flags().Lookup("port")
	assert.NotNil(t, portFlag)
	assert.Equal(t, "8080", portFlag.DefValue)
}
