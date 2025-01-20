package js

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewSimulation tests the NewSimulation function
func TestNewSimulation(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{"valid simulation", "test-simulation", false},
		{"empty id", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := NewRuntime(tt.id)
			if tt.wantErr {
				assert.Nil(t, si)
			} else {
				assert.NotNil(t, si)
				assert.NotNil(t, si.GetWorld())
			}
		})
	}
}

// TestCheckWorld tests $world variable
func TestCheckWorld(t *testing.T) {
	si := NewRuntime("test-simulation")
	assert.NotNil(t, si)
	assert.NotNil(t, si.GetWorld())
	script := `
		if ($world === undefined) {
			throw new Error("world not found");
		}
	`
	_, err := si.vm.RunString(script)
	assert.NoError(t, err)
}

// TestGetActor tests GetActor function
func TestGetActor(t *testing.T) {
	si := NewRuntime("test-simulation")
	assert.NotNil(t, si)
	assert.NotNil(t, si.GetWorld())
	script := `
		const actor = $world.createActor("actor1", {});
		actor;
	`
	_, err := si.vm.RunString(script)
	assert.NoError(t, err)
	actor := si.GetActor("actor1")
	assert.NotNil(t, actor)
	assert.Equal(t, "actor1", actor.Id())
}

// TestInterrupt tests the Interupt function
func TestInterrupt(t *testing.T) {
	si := NewRuntime("test-simulation")
	assert.NotNil(t, si)
	assert.NotNil(t, si.GetWorld())
	script := `
		const result = 42;
		while (true) {
			// do nothing
		}
		result;		
	`
	wg := sync.WaitGroup{}
	wg.Add(1)
	var value any
	go func() {
		defer wg.Done()
		result, err := si.vm.RunString(script)
		assert.Error(t, err)
		assert.Nil(t, result)
		value = result
	}()
	si.Interupt()
	wg.Wait()
	assert.Nil(t, value)

}
