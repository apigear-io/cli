package sim

import (
	"testing"
	"time"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceProxy(t *testing.T) {
	t.Run("PropertyAccess", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Service", map[string]any{
				"name":  "test",
				"count": 42,
			})
			defer service.Close()

			proxy := CreateServiceProxy(rt, service)

			// Test property get
			nameVal := proxy.Get("name")
			assert.Equal(t, "test", nameVal.Export())

			countVal := proxy.Get("count")
			assert.Equal(t, int64(42), countVal.Export())

			// Test property set
			proxy.Set("count", rt.ToValue(100))
			assert.Equal(t, int64(100), service.GetProperty("count"))

			done <- true
		})
		<-done
	})

	t.Run("MethodAssignmentAndCall", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Service", nil)
			defer service.Close()

			proxy := CreateServiceProxy(rt, service)

			// Assign a method through the proxy
			methodCalled := false
			method := rt.ToValue(func(call goja.FunctionCall) goja.Value {
				methodCalled = true
				return rt.ToValue("success")
			})

			proxy.Set("testMethod", method)
			assert.True(t, service.HasMethod("testMethod"))

			// Call the method through the proxy
			methodVal := proxy.Get("testMethod")
			fn, ok := goja.AssertFunction(methodVal)
			require.True(t, ok)

			result, err := fn(proxy)
			require.NoError(t, err)
			assert.True(t, methodCalled)
			assert.Equal(t, "success", result.Export())

			done <- true
		})
		<-done
	})

	t.Run("ThisBindingInMethods", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		script := `
		const service = $createService("test.Counter", {
			count: 10,
			max: 100
		});

		// Test 1: Method with 'this' accessing properties
		service.increment = function() {
			// 'this' should be the proxy, not undefined
			if (!this) {
				throw new Error("'this' is undefined");
			}
			if (this.count === undefined) {
				throw new Error("Cannot access count property");
			}
			this.count = this.count + 1;
			return this.count;
		};

		// Test 2: Method calling another method
		service.doubleIncrement = function() {
			this.increment();
			return this.increment();
		};

		// Test 3: Method returning 'this' for chaining
		service.setCount = function(value) {
			this.count = value;
			return this; // Should return the proxy
		};

		// Run tests
		const result1 = service.increment();
		if (result1 !== 11) {
			throw new Error("increment failed: expected 11, got " + result1);
		}

		const result2 = service.doubleIncrement();
		if (result2 !== 13) {
			throw new Error("doubleIncrement failed: expected 13, got " + result2);
		}

		const chainResult = service.setCount(20).increment();
		if (chainResult !== 21) {
			throw new Error("chaining failed: expected 21, got " + chainResult);
		}

		if (service.count !== 21) {
			throw new Error("final count wrong: expected 21, got " + service.count);
		}

		"success";
		`

		engine.RunScript("test_this_binding.js", script)
		time.Sleep(50 * time.Millisecond)
	})

	t.Run("OnMethodForEvents", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		script := `
		const service = $createService("test.EventEmitter", {
			value: 0
		});

		let propertyChangeCount = 0;
		let lastPropertyValue = null;

		// Test property change listener
		service.on('value', function(newValue) {
			propertyChangeCount++;
			lastPropertyValue = newValue;
		});

		service.value = 10;
		service.value = 20;

		if (propertyChangeCount !== 2) {
			throw new Error("Property listener not called correctly: " + propertyChangeCount);
		}

		if (lastPropertyValue !== 20) {
			throw new Error("Property value incorrect: " + lastPropertyValue);
		}

		// Test signal listener
		let signalReceived = false;
		let signalArgs = null;

		service.on('customSignal', function(arg1, arg2, arg3) {
			signalReceived = true;
			signalArgs = [arg1, arg2, arg3];
		});

		service.emit('customSignal', 'a', 'b', 'c');

		if (!signalReceived) {
			throw new Error("Signal not received");
		}

		if (signalArgs[0] !== 'a' || signalArgs[1] !== 'b' || signalArgs[2] !== 'c') {
			throw new Error("Signal args incorrect");
		}

		"success";
		`

		engine.RunScript("test_events.js", script)
		time.Sleep(50 * time.Millisecond)
	})

	t.Run("RawServiceAccess", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		script := `
		const service = $createService("test.Service", {
			prop: "value"
		});

		// Access raw service through $
		const raw = service.$;
		if (!raw) {
			throw new Error("Cannot access raw service");
		}

		// The raw service should be the target object
		if (typeof raw !== 'object') {
			throw new Error("Raw service is not an object");
		}

		"success";
		`

		engine.RunScript("test_raw_access.js", script)
		time.Sleep(50 * time.Millisecond)
	})

	t.Run("ProxyTraps", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Service", map[string]any{
				"prop1": "value1",
				"prop2": "value2",
			})
			defer service.Close()

			// Register a method
			service.OnMethod("method1", rt.ToValue(func(call goja.FunctionCall) goja.Value {
				return rt.ToValue("method1Result")
			}))

			proxy := CreateServiceProxy(rt, service)

			// Test Has trap
			script := `
			function testHas(obj) {
				return {
					hasProp1: 'prop1' in obj,
					hasProp2: 'prop2' in obj,
					hasMethod1: 'method1' in obj,
					hasNonExistent: 'nonExistent' in obj,
					hasOn: 'on' in obj,
					hasEmit: 'emit' in obj,
					hasDollar: '$' in obj
				};
			}
			`
			_, err := rt.RunString(script)
			require.NoError(t, err)

			testHas, ok := goja.AssertFunction(rt.Get("testHas"))
			require.True(t, ok)

			result, err := testHas(goja.Undefined(), proxy)
			require.NoError(t, err)

			resultObj := result.ToObject(rt)
			assert.True(t, resultObj.Get("hasProp1").ToBoolean())
			assert.True(t, resultObj.Get("hasProp2").ToBoolean())
			assert.True(t, resultObj.Get("hasMethod1").ToBoolean())
			assert.False(t, resultObj.Get("hasNonExistent").ToBoolean())
			assert.True(t, resultObj.Get("hasOn").ToBoolean())
			assert.True(t, resultObj.Get("hasEmit").ToBoolean())
			assert.True(t, resultObj.Get("hasDollar").ToBoolean())

			// Test OwnKeys trap
			ownKeysScript := `
			function getOwnKeys(obj) {
				return Object.keys(obj);
			}
			`
			_, err = rt.RunString(ownKeysScript)
			require.NoError(t, err)

			getOwnKeys, ok := goja.AssertFunction(rt.Get("getOwnKeys"))
			require.True(t, ok)

			keysResult, err := getOwnKeys(goja.Undefined(), proxy)
			require.NoError(t, err)

			keys := keysResult.Export().([]interface{})
			keyMap := make(map[string]bool)
			for _, k := range keys {
				keyMap[k.(string)] = true
			}

			assert.True(t, keyMap["prop1"])
			assert.True(t, keyMap["prop2"])
			assert.True(t, keyMap["method1"])
			assert.True(t, keyMap["$"])
			assert.True(t, keyMap["on"])
			assert.True(t, keyMap["emit"])

			done <- true
		})
		<-done
	})

	t.Run("ComplexScenario", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		script := `
		// Create a calculator service
		const calc = $createService("test.Calculator", {
			result: 0,
			history: []
		});

		// Add methods that use 'this' extensively
		calc.add = function(value) {
			this.result = this.result + value;
			// Note: arrays need special handling in Go/JS bridge
			return this;
		};

		calc.subtract = function(value) {
			this.result = this.result - value;
			return this;
		};

		calc.multiply = function(value) {
			this.result = this.result * value;
			return this;
		};

		calc.clear = function() {
			this.result = 0;
			this.emit('cleared');
			return this;
		};

		// Track events
		let clearedCount = 0;
		calc.on('cleared', function() {
			clearedCount++;
		});

		let resultChanges = 0;
		calc.on('result', function(newValue) {
			resultChanges++;
		});

		// Test method chaining with 'this' binding
		calc.add(5).multiply(3).subtract(7);
		
		if (calc.result !== 8) {
			throw new Error("Calculation wrong: expected 8, got " + calc.result);
		}

		// Clear and verify event
		calc.clear();
		if (calc.result !== 0) {
			throw new Error("Clear failed");
		}
		if (clearedCount !== 1) {
			throw new Error("Clear event not emitted");
		}

		// Verify property change notifications
		if (resultChanges < 4) {
			throw new Error("Property changes not tracked correctly: " + resultChanges);
		}

		// Test accessing method through variable
		const addMethod = calc.add;
		// This should still work because we bind 'this' in the proxy
		addMethod.call(calc, 10);
		if (calc.result !== 10) {
			throw new Error("Method call with explicit context failed");
		}

		"success";
		`

		engine.RunScript("test_complex.js", script)
		time.Sleep(50 * time.Millisecond)
	})
}

func TestProxyEdgeCases(t *testing.T) {
	t.Run("UndefinedPropertyWarning", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Service", map[string]any{
				"exists": true,
			})
			defer service.Close()

			proxy := CreateServiceProxy(rt, service)

			// Access undefined property - should return undefined
			val := proxy.Get("nonExistent")
			assert.True(t, goja.IsUndefined(val))

			done <- true
		})
		<-done
	})

	t.Run("PropertyDescriptor", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Service", map[string]any{
				"prop": "value",
			})
			defer service.Close()

			proxy := CreateServiceProxy(rt, service)

			script := `
			function getDescriptor(obj, prop) {
				const desc = Object.getOwnPropertyDescriptor(obj, prop);
				return desc ? {
					configurable: desc.configurable,
					enumerable: desc.enumerable,
					writable: desc.writable,
					hasValue: desc.value !== undefined
				} : null;
			}
			`
			_, err := rt.RunString(script)
			require.NoError(t, err)

			getDescriptor, ok := goja.AssertFunction(rt.Get("getDescriptor"))
			require.True(t, ok)

			result, err := getDescriptor(goja.Undefined(), proxy, rt.ToValue("prop"))
			require.NoError(t, err)

			if !goja.IsNull(result) {
				resultObj := result.ToObject(rt)
				// Check that descriptor flags are set correctly
				assert.True(t, resultObj.Get("configurable").ToBoolean())
				assert.True(t, resultObj.Get("enumerable").ToBoolean())
				assert.True(t, resultObj.Get("writable").ToBoolean())
			}

			done <- true
		})
		<-done
	})

	t.Run("DefineProperty", func(t *testing.T) {
		engine := NewEngine(EngineOptions{})
		defer engine.Close()

		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			service := NewObjectService(engine, "test.Service", nil)
			defer service.Close()

			proxy := CreateServiceProxy(rt, service)

			script := `
			function defineNewProperty(obj) {
				Object.defineProperty(obj, 'newProp', {
					value: 'newValue',
					writable: true,
					enumerable: true,
					configurable: true
				});
				
				// Define a method
				Object.defineProperty(obj, 'newMethod', {
					value: function() { return 'methodResult'; },
					writable: true,
					enumerable: true,
					configurable: true
				});
				
				return {
					propValue: obj.newProp,
					methodExists: typeof obj.newMethod === 'function',
					methodResult: obj.newMethod()
				};
			}
			`
			_, err := rt.RunString(script)
			require.NoError(t, err)

			defineNewProperty, ok := goja.AssertFunction(rt.Get("defineNewProperty"))
			require.True(t, ok)

			result, err := defineNewProperty(goja.Undefined(), proxy)
			require.NoError(t, err)

			resultObj := result.ToObject(rt)
			assert.Equal(t, "newValue", resultObj.Get("propValue").Export())
			assert.True(t, resultObj.Get("methodExists").ToBoolean())
			assert.Equal(t, "methodResult", resultObj.Get("methodResult").Export())

			// Verify in service
			assert.Equal(t, "newValue", service.GetProperty("newProp"))
			assert.True(t, service.HasMethod("newMethod"))

			done <- true
		})
		<-done
	})
}