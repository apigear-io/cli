package js

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
)

func setupTestActor(t *testing.T) (*Actor, *World) {
	vm := goja.New()
	world := NewWorld("test", vm)
	state := vm.NewObject()
	actor, err := NewActor("test", state, world)
	assert.NoError(t, err)
	assert.NotNil(t, actor)
	return actor, world
}

// Properties Tests
func TestActorProperties(t *testing.T) {
	t.Run("SetProperty and GetProperty", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		value := actor.vm.ToValue("test-value")
		actor.SetProperty("testProp", value)
		result := actor.GetProperty("testProp")
		assert.Equal(t, "test-value", result.String())
	})

	t.Run("HasProperty", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		value := actor.vm.ToValue("test-value")
		actor.SetProperty("testProp", value)
		assert.True(t, actor.HasProperty("testProp"))
		assert.False(t, actor.HasProperty("nonExistentProp"))
	})

	t.Run("OnProperty", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		called := false
		handler := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			called = true
			return goja.Undefined(), nil
		}
		actor.OnProperty("testProp", handler)
		actor.SetProperty("testProp", actor.vm.ToValue("new-value"))
		assert.True(t, called)
	})

	t.Run("SetState and GetState", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		newState := actor.vm.NewObject()
		err := newState.Set("prop1", actor.vm.ToValue("value1"))
		assert.NoError(t, err)
		err = newState.Set("prop2", actor.vm.ToValue("value2"))
		assert.NoError(t, err)
		actor.SetState(newState)
		state := actor.GetState()
		assert.Equal(t, "value1", state.Get("prop1").String())
		assert.Equal(t, "value2", state.Get("prop2").String())
	})
}

// Methods Tests
func TestActorMethods(t *testing.T) {
	t.Run("SetMethod and HasMethod", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		method := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			return actor.vm.ToValue("result"), nil
		}
		actor.SetMethod("testMethod", method)
		assert.True(t, actor.HasMethod("testMethod"))
		assert.False(t, actor.HasMethod("nonExistentMethod"))
	})

	t.Run("CallMethod", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		method := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			return actor.vm.ToValue("result"), nil
		}
		actor.SetMethod("testMethod", method)
		result, err := actor.CallMethod("testMethod")
		assert.NoError(t, err)
		assert.Equal(t, "result", result.String())
	})

	t.Run("CallMethod with arguments", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		method := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			return args[0], nil
		}
		actor.SetMethod("testMethod", method)
		arg := actor.vm.ToValue("test-arg")
		result, err := actor.CallMethod("testMethod", arg)
		assert.NoError(t, err)
		assert.Equal(t, "test-arg", result.String())
	})

	t.Run("OnMethod", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		called := false
		method := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			return actor.vm.ToValue("result"), nil
		}
		handler := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			called = true
			return goja.Undefined(), nil
		}
		actor.SetMethod("testMethod", method)
		actor.OnMethod("testMethod", handler)
		_, err := actor.CallMethod("testMethod")
		assert.NoError(t, err)
		assert.True(t, called)
	})
}

// Signals Tests
func TestActorSignals(t *testing.T) {
	t.Run("EmitSignal and OnSignal", func(t *testing.T) {
		actor, _ := setupTestActor(t)
		receivedArgs := make([]goja.Value, 0)
		handler := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			receivedArgs = args
			return goja.Undefined(), nil
		}
		actor.OnSignal("testSignal", handler)
		arg := actor.vm.ToValue("test-arg")
		actor.EmitSignal("testSignal", arg)
		assert.Equal(t, 1, len(receivedArgs))
		assert.Equal(t, "test-arg", receivedArgs[0].String())
	})
}

func TestActorCreation(t *testing.T) {
	t.Run("NewActor with valid parameters", func(t *testing.T) {
		vm := goja.New()
		world := NewWorld("test", vm)
		state := vm.NewObject()
		actor, err := NewActor("test", state, world)
		assert.NoError(t, err)
		assert.NotNil(t, actor)
		assert.Equal(t, "test", actor.Id())
	})

	t.Run("NewActor with empty id", func(t *testing.T) {
		vm := goja.New()
		world := NewWorld("test", vm)
		state := vm.NewObject()
		actor, err := NewActor("", state, world)
		assert.Error(t, err)
		assert.Nil(t, actor)
	})

	t.Run("NewActor with nil world", func(t *testing.T) {
		actor, err := NewActor("test", nil, nil)
		assert.Error(t, err)
		assert.Nil(t, actor)
	})

	t.Run("NewActor with nil state", func(t *testing.T) {
		vm := goja.New()
		world := NewWorld("test", vm)
		actor, err := NewActor("test", nil, world)
		assert.NoError(t, err)
		assert.NotNil(t, actor)
	})
}
