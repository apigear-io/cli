package sim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngineCreate(t *testing.T) {
	engine := NewEngine(EngineOptions{})
	defer engine.Close()
	assert.NotNil(t, engine)
}

func TestEngineCreateService(t *testing.T) {

	// TODO: avoid ws hub is created, pass in an interface
	server := &MockEngineServer{}
	engine := NewEngine(EngineOptions{Server: server})
	service, err := engine.world.CreateService("test", nil)
	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.Len(t, server.sources, 1)
	defer engine.Close()
	assert.NotNil(t, engine)
}
