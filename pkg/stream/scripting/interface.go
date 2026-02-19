package scripting

import (
	"sync"

	"github.com/dop251/goja"
)

// StreamSender is an interface for sending ObjectLink messages.
type StreamSender interface {
	SetProperty(propertyID string, value interface{})
	Invoke(vm *goja.Runtime, methodID string, args []interface{}) goja.Value
	ScheduleCallback(fn func(*goja.Runtime))
}

// InterfaceHandle provides a high-level wrapper for a specific ObjectLink interface.
// It simplifies method calls and property access by auto-completing symbol paths.
type InterfaceHandle struct {
	objectID string
	sender   StreamSender // Sender interface (WSStream)
	engine   *Engine      // Engine for error reporting

	// Event handlers specific to this interface
	onPropertyChangeHandlers map[string][]goja.Callable // propName -> handlers
	onSignalHandlers         map[string][]goja.Callable // signalName -> handlers
	onInitHandlers           []goja.Callable
	handlersMu               sync.RWMutex

	// Cached properties from INIT message
	properties   map[string]interface{}
	propertiesMu sync.RWMutex
}

// NewInterfaceHandle creates a new interface handle for a WSStream.
func NewInterfaceHandle(objectID string, ws *WSStream) *InterfaceHandle {
	return &InterfaceHandle{
		objectID:                 objectID,
		sender:                   ws,
		engine:                   ws.engine,
		onPropertyChangeHandlers: make(map[string][]goja.Callable),
		onSignalHandlers:         make(map[string][]goja.Callable),
		properties:               make(map[string]interface{}),
	}
}

// ToValue converts the InterfaceHandle to a JavaScript object.
func (i *InterfaceHandle) ToValue(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()

	// invoke(methodName, ...args) - Invoke method with auto-completed path, returns Promise
	_ = obj.Set("invoke", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("invoke requires methodName argument"))
		}
		methodName := call.Arguments[0].String()

		// Collect remaining arguments as spread args
		var args []interface{}
		for j := 1; j < len(call.Arguments); j++ {
			args = append(args, call.Arguments[j].Export())
		}

		return i.invoke(vm, methodName, args)
	})

	// setProperty(propName, value) - Set property with auto-completed path
	_ = obj.Set("setProperty", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("setProperty requires propName and value arguments"))
		}
		propName := call.Arguments[0].String()
		value := call.Arguments[1].Export()
		i.SetProperty(propName, value)
		return goja.Undefined()
	})

	// getProperty(propName) - Get cached property value
	_ = obj.Set("getProperty", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("getProperty requires propName argument"))
		}
		propName := call.Arguments[0].String()
		value := i.GetProperty(propName)
		return vm.ToValue(value)
	})

	// onPropertyChange(propName, callback) or onPropertyChange(callback)
	_ = obj.Set("onPropertyChange", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onPropertyChange requires at least one argument"))
		}

		// Check if first arg is a string (property name) or function (callback)
		if callback, ok := goja.AssertFunction(call.Arguments[0]); ok {
			// onPropertyChange(callback) - listen to all properties
			i.OnPropertyChange("", callback)
		} else if len(call.Arguments) >= 2 {
			// onPropertyChange(propName, callback)
			propName := call.Arguments[0].String()
			callback, ok := goja.AssertFunction(call.Arguments[1])
			if !ok {
				panic(vm.NewTypeError("second argument must be a function"))
			}
			i.OnPropertyChange(propName, callback)
		} else {
			panic(vm.NewTypeError("onPropertyChange requires (callback) or (propName, callback)"))
		}
		return goja.Undefined()
	})

	// onSignal(signalName, callback)
	_ = obj.Set("onSignal", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("onSignal requires signalName and callback arguments"))
		}
		signalName := call.Arguments[0].String()
		callback, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(vm.NewTypeError("second argument must be a function"))
		}
		i.OnSignal(signalName, callback)
		return goja.Undefined()
	})

	// onInit(callback) - Called when INIT is received for this interface
	_ = obj.Set("onInit", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onInit requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("argument must be a function"))
		}
		i.OnInit(callback)
		return goja.Undefined()
	})

	// properties - Get all cached properties
	_ = obj.Set("properties", vm.ToValue(func() map[string]interface{} {
		i.propertiesMu.RLock()
		defer i.propertiesMu.RUnlock()
		result := make(map[string]interface{}, len(i.properties))
		for k, v := range i.properties {
			result[k] = v
		}
		return result
	}))

	// objectId property
	_ = obj.Set("objectId", i.objectID)

	return obj
}

// invoke calls a method on this interface and returns a Promise.
func (i *InterfaceHandle) invoke(vm *goja.Runtime, methodName string, args []interface{}) goja.Value {
	methodID := i.objectID + "/" + methodName
	return i.sender.Invoke(vm, methodID, args)
}

// SetProperty sets a property on this interface.
func (i *InterfaceHandle) SetProperty(propName string, value interface{}) {
	propertyID := i.objectID + "/" + propName
	i.sender.SetProperty(propertyID, value)
}

// SetProperties updates all cached properties.
func (i *InterfaceHandle) SetProperties(props map[string]interface{}) {
	i.propertiesMu.Lock()
	for k, v := range props {
		i.properties[k] = v
	}
	i.propertiesMu.Unlock()
}

// getScheduleCallback returns a function to schedule callbacks.
func (i *InterfaceHandle) getScheduleCallback() func(func(*goja.Runtime)) {
	return i.sender.ScheduleCallback
}

// GetProperty returns a cached property value.
func (i *InterfaceHandle) GetProperty(propName string) interface{} {
	i.propertiesMu.RLock()
	defer i.propertiesMu.RUnlock()
	return i.properties[propName]
}

// OnPropertyChange registers a property change handler.
// If propName is empty, the callback is called for all property changes.
func (i *InterfaceHandle) OnPropertyChange(propName string, callback goja.Callable) {
	i.handlersMu.Lock()
	defer i.handlersMu.Unlock()
	i.onPropertyChangeHandlers[propName] = append(i.onPropertyChangeHandlers[propName], callback)
}

// OnSignal registers a signal handler.
func (i *InterfaceHandle) OnSignal(signalName string, callback goja.Callable) {
	i.handlersMu.Lock()
	defer i.handlersMu.Unlock()
	i.onSignalHandlers[signalName] = append(i.onSignalHandlers[signalName], callback)
}

// OnInit registers an init handler.
func (i *InterfaceHandle) OnInit(callback goja.Callable) {
	i.handlersMu.Lock()
	defer i.handlersMu.Unlock()
	i.onInitHandlers = append(i.onInitHandlers, callback)
}

// HandleInit is called when an INIT message is received for this interface.
func (i *InterfaceHandle) HandleInit(vm *goja.Runtime, properties interface{}) {
	// Update cached properties
	if props, ok := properties.(map[string]interface{}); ok {
		i.propertiesMu.Lock()
		for k, v := range props {
			i.properties[k] = v
		}
		i.propertiesMu.Unlock()
	}

	// Call handlers
	i.handlersMu.RLock()
	handlers := make([]goja.Callable, len(i.onInitHandlers))
	copy(handlers, i.onInitHandlers)
	i.handlersMu.RUnlock()

	for _, handler := range handlers {
		i.engine.CallHandler(handler, "interface.onInit", vm.ToValue(properties))
	}
}

// HandlePropertyChange is the exported version for WSStream.
func (i *InterfaceHandle) HandlePropertyChange(vm *goja.Runtime, propName string, value interface{}) {
	// Update cached property
	i.propertiesMu.Lock()
	i.properties[propName] = value
	i.propertiesMu.Unlock()

	// Get handlers
	i.handlersMu.RLock()
	// Specific handlers for this property
	specificHandlers := make([]goja.Callable, len(i.onPropertyChangeHandlers[propName]))
	copy(specificHandlers, i.onPropertyChangeHandlers[propName])
	// Generic handlers (empty string key)
	genericHandlers := make([]goja.Callable, len(i.onPropertyChangeHandlers[""]))
	copy(genericHandlers, i.onPropertyChangeHandlers[""])
	i.handlersMu.RUnlock()

	// Call specific handlers with just value
	for _, handler := range specificHandlers {
		i.engine.CallHandler(handler, "interface.onPropertyChange", vm.ToValue(value))
	}
	// Call generic handlers with propName and value
	for _, handler := range genericHandlers {
		i.engine.CallHandler(handler, "interface.onPropertyChange", vm.ToValue(propName), vm.ToValue(value))
	}
}

// HandleSignal is the exported version for WSStream.
func (i *InterfaceHandle) HandleSignal(vm *goja.Runtime, signalName string, args interface{}) {
	// Get handlers
	i.handlersMu.RLock()
	handlers := make([]goja.Callable, len(i.onSignalHandlers[signalName]))
	copy(handlers, i.onSignalHandlers[signalName])
	i.handlersMu.RUnlock()

	// Convert args array to individual arguments
	argsSlice, ok := args.([]interface{})
	if !ok {
		argsSlice = []interface{}{args}
	}

	jsArgs := make([]goja.Value, len(argsSlice))
	for idx, arg := range argsSlice {
		jsArgs[idx] = vm.ToValue(arg)
	}

	for _, handler := range handlers {
		i.engine.CallHandler(handler, "interface.onSignal", jsArgs...)
	}
}
