package sim

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObjectService(t *testing.T) {
	t.Run("CreateService", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		service := NewObjectService(engine, "test.Service", map[string]any{
			"count": 0,
			"name":  "test",
		})
		defer service.Close()

		assert.Equal(t, "test.Service", service.ObjectId())
		assert.Equal(t, 0, service.GetProperty("count"))
		assert.Equal(t, "test", service.GetProperty("name"))
	})

	t.Run("PropertyGetSet", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		service := NewObjectService(engine, "test.Service", nil)
		defer service.Close()

		// Set and get property
		service.SetProperty("value", 42)
		assert.Equal(t, 42, service.GetProperty("value"))

		// Update property
		service.SetProperty("value", 100)
		assert.Equal(t, 100, service.GetProperty("value"))
	})

	t.Run("PropertyChangeNotification", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		service := NewObjectService(engine, "test.Service", map[string]any{
			"count": 0,
		})
		defer service.Close()

		notified := false
		var notifiedValue any

		service.OnProperty("count", func(value any) {
			notified = true
			notifiedValue = value
		})

		service.SetProperty("count", 10)
		assert.True(t, notified)
		assert.Equal(t, 10, notifiedValue)

		// Same value should not trigger notification
		notified = false
		service.SetProperty("count", 10)
		assert.False(t, notified)
	})

	t.Run("HasProperty", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		service := NewObjectService(engine, "test.Service", map[string]any{
			"exists": true,
		})
		defer service.Close()

		assert.True(t, service.HasProperty("exists"))
		assert.False(t, service.HasProperty("notExists"))
	})

	t.Run("SignalEmission", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		service := NewObjectService(engine, "test.Service", nil)
		defer service.Close()

		received := false
		var receivedArgs []any

		service.OnSignal("testSignal", func(args ...any) {
			received = true
			receivedArgs = args
		})

		service.EmitSignal("testSignal", "arg1", 42, true)
		assert.True(t, received)
		assert.Equal(t, []any{"arg1", 42, true}, receivedArgs)
	})

	t.Run("MethodRegistration", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		// Need to run in event loop to have runtime available
		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Service", nil)
			defer service.Close()

			// Register a method
			methodFn := rt.ToValue(func(call goja.FunctionCall) goja.Value {
				return rt.ToValue("result")
			})
			service.OnMethod("testMethod", methodFn)

			assert.True(t, service.HasMethod("testMethod"))
			assert.False(t, service.HasMethod("nonExistent"))

			// Call the method
			result, err := service.CallMethod("testMethod", "arg1")
			assert.NoError(t, err)
			assert.Equal(t, "result", result.Export())

			done <- true
		})
		<-done
	})

	t.Run("MultiplePropertyListeners", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		service := NewObjectService(engine, "test.Service", map[string]any{
			"value": 0,
		})
		defer service.Close()

		count1 := 0
		count2 := 0

		service.OnProperty("value", func(value any) {
			count1++
		})

		service.OnProperty("value", func(value any) {
			count2++
		})

		service.SetProperty("value", 10)
		assert.Equal(t, 1, count1)
		assert.Equal(t, 1, count2)
	})

	t.Run("SetProperties", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		service := NewObjectService(engine, "test.Service", nil)
		defer service.Close()

		notificationCount := 0
		service.OnProperty("prop1", func(value any) {
			notificationCount++
		})
		service.OnProperty("prop2", func(value any) {
			notificationCount++
		})

		service.SetProperties(map[string]any{
			"prop1": "value1",
			"prop2": "value2",
			"prop3": "value3",
		})

		assert.Equal(t, "value1", service.GetProperty("prop1"))
		assert.Equal(t, "value2", service.GetProperty("prop2"))
		assert.Equal(t, "value3", service.GetProperty("prop3"))
		assert.Equal(t, 2, notificationCount)
	})
}

func TestObjectServiceWithRuntime(t *testing.T) {
	t.Run("MethodWithThisContext", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Counter", map[string]any{
				"count": 5,
			})
			defer service.Close()

			// Create a method that uses 'this'
			incrementMethod := rt.ToValue(func(call goja.FunctionCall) goja.Value {
				this := call.This
				if this == nil || goja.IsUndefined(this) {
					t.Error("'this' is undefined in method")
					return goja.Undefined()
				}

				// Get count property from 'this'
				countVal := this.ToObject(rt).Get("count")
				count := countVal.ToInteger()
				
				// Set new count value
				this.ToObject(rt).Set("count", rt.ToValue(count + 1))
				
				return rt.ToValue(count + 1)
			})

			service.OnMethod("increment", incrementMethod)

			// Test calling with proper context
			proxy := CreateServiceProxy(rt, service)
			
			// Call increment through proxy
			incrementFn := proxy.Get("increment")
			require.NotNil(t, incrementFn)
			
			fn, ok := goja.AssertFunction(incrementFn)
			require.True(t, ok)
			
			result, err := fn(proxy)
			require.NoError(t, err)
			assert.Equal(t, int64(6), result.ToInteger())
			
			// Verify count was updated
			assert.Equal(t, int64(6), service.GetProperty("count"))

			done <- true
		})
		<-done
	})

	t.Run("ChainedMethodCalls", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Builder", map[string]any{
				"value": "",
			})
			defer service.Close()

			// Create methods that return 'this' for chaining
			appendMethod := rt.ToValue(func(call goja.FunctionCall) goja.Value {
				this := call.This
				if len(call.Arguments) > 0 {
					currentVal := this.ToObject(rt).Get("value").String()
					newVal := currentVal + call.Arguments[0].String()
					this.ToObject(rt).Set("value", rt.ToValue(newVal))
				}
				return this // Return 'this' for chaining
			})

			service.OnMethod("append", appendMethod)

			proxy := CreateServiceProxy(rt, service)
			
			// Test method chaining
			appendFn := proxy.Get("append")
			fn, _ := goja.AssertFunction(appendFn)
			
			// First call: append("hello")
			result1, err := fn(proxy, rt.ToValue("hello"))
			require.NoError(t, err)
			assert.Equal(t, proxy, result1.ToObject(rt))
			
			// Second call: append(" world")
			result2, err := fn(proxy, rt.ToValue(" world"))
			require.NoError(t, err)
			assert.Equal(t, proxy, result2.ToObject(rt))
			
			// Verify final value
			assert.Equal(t, "hello world", service.GetProperty("value"))

			done <- true
		})
		<-done
	})
}