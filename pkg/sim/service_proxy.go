package sim

import (
	"github.com/dop251/goja"
)

func CreateServiceProxy(vm *goja.Runtime, service *ObjectService) *goja.Object {
	// Create target object
	target := vm.NewObject()

	// Store the service reference
	target.Set("__service", service)

	// Create proxy with all trap handlers
	proxyConfig := &goja.ProxyTrapConfig{
		Get: func(target *goja.Object, property string, receiver goja.Value) (value goja.Value) {
			// Access to raw service object - just return the service
			// Goja will automatically handle the method name conversion
			if property == "$" {
				return vm.ToValue(service)
			}

			// Method access - return bound method with proxy as context
			if service.HasMethod(property) {
				method := service.GetMethod(property)
				// Create a wrapper function that will be called with the proxy as context
				return vm.ToValue(func(call goja.FunctionCall) goja.Value {
					// Call the original method with the proxy as 'this'
					result, err := method(receiver, call.Arguments...)
					if err != nil {
						panic(err)
					}
					return result
				})
			}

			// Property access
			if service.HasProperty(property) {
				val := service.GetProperty(property)
				// Return undefined for nil values (more JavaScript-idiomatic)
				if val == nil {
					return goja.Undefined()
				}
				return vm.ToValue(val)
			}

			// Convenience method: on
			if property == "on" {
				return vm.ToValue(func(call goja.FunctionCall) goja.Value {
					if len(call.Arguments) < 2 {
						panic(vm.NewTypeError("on requires at least 2 arguments"))
					}
					event := call.Argument(0).String()
					callback := call.Argument(1)

					if fn, ok := goja.AssertFunction(callback); ok {
						if service.HasProperty(event) {
							service.OnProperty(event, func(value any) {
								fn(goja.Undefined(), vm.ToValue(value))
							})
						} else {
							service.OnSignal(event, func(args ...any) {
								jsArgs := make([]goja.Value, len(args))
								for i, arg := range args {
									jsArgs[i] = vm.ToValue(arg)
								}
								fn(goja.Undefined(), jsArgs...)
							})
						}
					}
					return goja.Undefined()
				})
			}

			// Convenience method: emit
			if property == "emit" {
				return vm.ToValue(func(call goja.FunctionCall) goja.Value {
					if len(call.Arguments) < 1 {
						panic(vm.NewTypeError("emit requires at least 1 argument"))
					}
					signal := call.Argument(0).String()
					args := make([]any, len(call.Arguments)-1)
					for i := 1; i < len(call.Arguments); i++ {
						args[i-1] = call.Arguments[i].Export()
					}
					service.EmitSignal(signal, args...)
					return goja.Undefined()
				})
			}

			// Built-in service methods and properties
			if val := target.Get(property); val != nil && !goja.IsUndefined(val) {
				if fn, ok := goja.AssertFunction(val); ok {
					log.Debug().Str("property", property).Msg("Returning built-in function")
					return vm.ToValue(fn)
				}
				return val
			}

			// Direct service method access (objectId, hasMethod, hasProperty, etc.)
			serviceVal := vm.ToValue(service)
			if serviceObj := serviceVal.ToObject(vm); serviceObj != nil {
				if method := serviceObj.Get(property); method != nil && !goja.IsUndefined(method) {
					return method
				}
			}

			// Undefined property - provide helpful error
			if property != "" && property[0] != '_' {
				keys := make([]string, 0, len(service.properties))
				for k := range service.properties {
					keys = append(keys, k)
				}
				log.Warn().Str("property", property).Str("objectId", service.ObjectId()).Strs("available", keys).Msg("Property not found on service")
			}

			return goja.Undefined()
		},

		Set: func(target *goja.Object, property string, value goja.Value, receiver goja.Value) bool {
			// Don't intercept internal properties except __proto__
			if property == "$" || (len(property) > 0 && property[0] == '_' && property != "__proto__") {
				target.Set(property, value)
				return true
			}

			// Function assignment = method registration
			if _, ok := goja.AssertFunction(value); ok {
				log.Debug().Str("property", property).Msg("Registering method")
				service.OnMethod(property, value)
				// If this was previously a property, remove it
				service.RemoveProperty(property)
				return true
			}

			// Property assignment (including when overwriting a method)
			service.SetProperty(property, value.Export())
			// If this was previously a method, remove it
			service.RemoveMethod(property)
			return true
		},

		Has: func(target *goja.Object, property string) bool {
			// Check for special properties
			if property == "$" || property == "on" || property == "emit" {
				return true
			}

			// Check service properties and methods
			return service.HasProperty(property) || service.HasMethod(property)
		},

		OwnKeys: func(target *goja.Object) *goja.Object {
			// Collect all property and method names
			keys := make([]any, 0)

			// Add properties
			for k := range service.properties {
				keys = append(keys, k)
			}

			// Add methods
			for k := range service.methods {
				keys = append(keys, k)
			}

			// Add special properties
			keys = append(keys, "$", "on", "emit")

			return vm.ToValue(keys).ToObject(vm)
		},

		GetOwnPropertyDescriptor: func(target *goja.Object, property string) goja.PropertyDescriptor {
			// Check if property exists
			if property == "$" || property == "on" || property == "emit" ||
				service.HasProperty(property) || service.HasMethod(property) {
				return goja.PropertyDescriptor{
					Configurable: goja.FLAG_TRUE,
					Enumerable:   goja.FLAG_TRUE,
					Writable:     goja.FLAG_TRUE,
				}
			}
			return goja.PropertyDescriptor{}
		},

		DefineProperty: func(target *goja.Object, property string, descriptor goja.PropertyDescriptor) bool {
			// Allow property definition
			if descriptor.Value != nil {
				if _, ok := goja.AssertFunction(descriptor.Value); ok {
					service.OnMethod(property, descriptor.Value)
				} else {
					service.SetProperty(property, descriptor.Value.Export())
				}
				return true
			}
			return false
		},

		DeleteProperty: func(target *goja.Object, property string) bool {
			// For now, don't allow deletion of properties
			// This could be enhanced to support property removal if needed
			return false
		},
	}

	proxy := vm.NewProxy(target, proxyConfig)
	return vm.ToValue(proxy).ToObject(vm)
}

// CreateService creates a new service with proxy wrapper
func CreateService(engine *Engine, objectId string, properties map[string]any) *goja.Object {
	service := NewObjectService(engine, objectId, properties)
	return CreateServiceProxy(engine.rt, service)
}
