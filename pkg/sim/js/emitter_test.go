package js

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
)

func TestEmitter(t *testing.T) {
	vm := goja.New()
	e := NewEmitter(vm)

	t.Run("On and Emit", func(t *testing.T) {
		called := false
		fn := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			called = true
			return nil, nil
		}
		e.On("test", fn)
		e.Emit("test")
		assert.True(t, called, "handler should be called")
	})

	t.Run("Off", func(t *testing.T) {
		called := false
		fn := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			called = true
			return nil, nil
		}
		e.On("test", fn)
		e.Emit("test")
		assert.True(t, called, "handler should be called")
		called = false
		e.Off("test")
		assert.False(t, called, "handler should not be called after Off")
	})

	t.Run("Clear", func(t *testing.T) {
		called := false
		fn := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			called = true
			return nil, nil
		}
		e.On("test", fn)
		e.Emit("test")
		assert.True(t, called, "handler should be called")
		called = false
		e.Clear()
		e.Emit("test")
		assert.False(t, called, "handler should not be called after Clear")
	})

	t.Run("All events", func(t *testing.T) {
		fn := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			return nil, nil
		}
		e.Clear()
		e.On("event1", fn)
		e.On("event2", fn)
		events := e.All()
		assert.Len(t, events, 2, "should have 2 events")
		assert.Contains(t, events, "event1")
		assert.Contains(t, events, "event2")
	})

	t.Run("OnAny", func(t *testing.T) {
		var receivedEvent string
		var receivedArgs []goja.Value
		fn := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			receivedEvent = args[0].String()
			receivedArgs = args[1:]
			return goja.Undefined(), nil
		}
		e.Clear()
		e.OnAny(fn)
		testArg := vm.ToValue("testArg")
		e.Emit("testEvent", testArg)
		assert.Equal(t, "testEvent", receivedEvent)
		assert.Equal(t, testArg, receivedArgs[0])
	})

	t.Run("Unsubscribe with returned function", func(t *testing.T) {
		called := false
		fn := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
			called = true
			return goja.Undefined(), nil
		}
		unsubscribe := e.On("test", fn)
		_, _ = unsubscribe(nil)
		e.Emit("test")
		assert.False(t, called, "handler should not be called after unsubscribe")
	})
}
