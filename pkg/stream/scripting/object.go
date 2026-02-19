package scripting

import (
	"sync"

	"github.com/dop251/goja"
)

// ObjectLink message type constants
const (
	MsgLink           = 10
	MsgInit           = 11
	MsgUnlink         = 12
	MsgSetProperty    = 20
	MsgPropertyChange = 50
	MsgInvoke         = 30
	MsgInvokeReply    = 31
	MsgSignal         = 40
	MsgError          = 70
)

// ObjectDefinition defines a registered ObjectLink object.
type ObjectDefinition struct {
	ObjectID   string
	Properties map[string]interface{}
	Methods    map[string]goja.Callable

	// Optional callbacks
	OnLink        goja.Callable
	OnUnlink      goja.Callable
	OnSetProperty goja.Callable

	// Reference to the engine for broadcasting
	engine *Engine

	// Reference to the underlying objectlink object (future integration with objectlink-core-go)
	olObject interface{}

	// Mutex for property access
	mu sync.RWMutex
}

// NewObjectDefinition creates a new object definition.
func NewObjectDefinition(objectID string, engine *Engine) *ObjectDefinition {
	return &ObjectDefinition{
		ObjectID:   objectID,
		Properties: make(map[string]interface{}),
		Methods:    make(map[string]goja.Callable),
		engine:     engine,
	}
}

// GetProperty returns a property value.
func (o *ObjectDefinition) GetProperty(name string) interface{} {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.Properties[name]
}

// SetProperty sets a property value and broadcasts PROPERTY_CHANGE.
// If the new value equals the old value, no write or broadcast occurs.
func (o *ObjectDefinition) SetProperty(name string, value interface{}) {
	o.mu.Lock()
	oldValue, exists := o.Properties[name]
	if exists && valuesEqual(oldValue, value) {
		o.mu.Unlock()
		return // No change, skip write and broadcast
	}
	o.Properties[name] = value
	o.mu.Unlock()

	// Broadcast property change
	o.broadcastPropertyChange(name, value)
}

// valuesEqual compares two values for equality.
// Handles common JSON types: nil, bool, float64, string, []interface{}, map[string]interface{}.
func valuesEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	switch av := a.(type) {
	case bool:
		bv, ok := b.(bool)
		return ok && av == bv
	case float64:
		bv, ok := b.(float64)
		return ok && av == bv
	case int:
		// Handle int comparison with float64 (JSON numbers are float64)
		switch bv := b.(type) {
		case int:
			return av == bv
		case float64:
			return float64(av) == bv
		}
		return false
	case string:
		bv, ok := b.(string)
		return ok && av == bv
	case []interface{}:
		bv, ok := b.([]interface{})
		if !ok || len(av) != len(bv) {
			return false
		}
		for i := range av {
			if !valuesEqual(av[i], bv[i]) {
				return false
			}
		}
		return true
	case map[string]interface{}:
		bv, ok := b.(map[string]interface{})
		if !ok || len(av) != len(bv) {
			return false
		}
		for k, v := range av {
			if bVal, exists := bv[k]; !exists || !valuesEqual(v, bVal) {
				return false
			}
		}
		return true
	default:
		// Fallback: use reflect for other types
		return a == b
	}
}

// GetProperties returns a copy of all properties.
func (o *ObjectDefinition) GetProperties() map[string]interface{} {
	o.mu.RLock()
	defer o.mu.RUnlock()

	props := make(map[string]interface{})
	for k, v := range o.Properties {
		props[k] = v
	}
	return props
}

// broadcastPropertyChange sends PROPERTY_CHANGE to all linked clients.
func (o *ObjectDefinition) broadcastPropertyChange(propName string, value interface{}) {
	propertyID := o.ObjectID + "/" + propName
	msg := []interface{}{MsgPropertyChange, propertyID, value}

	if o.engine != nil {
		if server := o.engine.GetBackendServer(); server != nil {
			server.BroadcastToLinked(o.ObjectID, msg)
		}
	}
}

// Emit sends a SIGNAL to all linked clients.
func (o *ObjectDefinition) Emit(signalName string, args ...interface{}) {
	signalID := o.ObjectID + "/" + signalName
	msg := []interface{}{MsgSignal, signalID, args}

	if o.engine != nil {
		if server := o.engine.GetBackendServer(); server != nil {
			server.BroadcastToLinked(o.ObjectID, msg)
		}
	}
}

// ClientCount returns the number of clients linked to this object.
func (o *ObjectDefinition) ClientCount() int {
	if o.engine != nil {
		if server := o.engine.GetBackendServer(); server != nil {
			return server.GetLinkedClientCount(o.ObjectID)
		}
	}
	return 0
}

// BackendError represents an error from a backend handler.
type BackendError struct {
	Message string
}

// NewBackendError creates a new backend error.
func NewBackendError(message string) *BackendError {
	return &BackendError{Message: message}
}

func (e *BackendError) Error() string {
	return e.Message
}

// ObjectContext provides context for method handlers.
type ObjectContext struct {
	ObjectID string
	object   *ObjectDefinition
	engine   *Engine
}

// NewObjectContext creates a new object context.
func NewObjectContext(obj *ObjectDefinition, engine *Engine) *ObjectContext {
	return &ObjectContext{
		ObjectID: obj.ObjectID,
		object:   obj,
		engine:   engine,
	}
}

// Get returns a property value.
func (c *ObjectContext) Get(propName string) interface{} {
	return c.object.GetProperty(propName)
}

// Set sets a property and broadcasts PROPERTY_CHANGE.
func (c *ObjectContext) Set(propName string, value interface{}) {
	c.object.SetProperty(propName, value)
}

// Properties returns all properties.
func (c *ObjectContext) Properties() map[string]interface{} {
	return c.object.GetProperties()
}

// Emit sends a SIGNAL to all linked clients.
func (c *ObjectContext) Emit(signalName string, args ...interface{}) {
	c.object.Emit(signalName, args...)
}

// ClientCount returns the number of linked clients.
func (c *ObjectContext) ClientCount() int {
	return c.object.ClientCount()
}

// ToValue converts the ObjectContext to a JavaScript object.
func (c *ObjectContext) ToValue(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()

	// objectId property
	_ = obj.Set("objectId", c.ObjectID)

	// get(propName) - Get property value
	_ = obj.Set("get", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		propName := call.Arguments[0].String()
		return vm.ToValue(c.Get(propName))
	})

	// set(propName, value) - Set property and broadcast
	_ = obj.Set("set", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}
		propName := call.Arguments[0].String()
		value := call.Arguments[1].Export()
		c.Set(propName, value)
		return goja.Undefined()
	})

	// properties() - Get all properties
	_ = obj.Set("properties", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(c.Properties())
	})

	// emit(signalName, ...args) - Send signal to linked clients
	_ = obj.Set("emit", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		signalName := call.Arguments[0].String()
		var args []interface{}
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Arguments[i].Export())
		}
		c.Emit(signalName, args...)
		return goja.Undefined()
	})

	// clientCount() - Get number of linked clients
	_ = obj.Set("clientCount", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(c.ClientCount())
	})

	// error(message) - Throw an error (for use in handlers)
	_ = obj.Set("error", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewGoError(NewBackendError("unknown error")))
		}
		panic(vm.NewGoError(NewBackendError(call.Arguments[0].String())))
	})

	return obj
}

// ObjectHandle provides external access to a registered object.
type ObjectHandle struct {
	object *ObjectDefinition
	engine *Engine
}

// NewObjectHandle creates a new object handle.
func NewObjectHandle(obj *ObjectDefinition, engine *Engine) *ObjectHandle {
	return &ObjectHandle{
		object: obj,
		engine: engine,
	}
}

// Get returns a property value.
func (h *ObjectHandle) Get(propName string) interface{} {
	return h.object.GetProperty(propName)
}

// Set sets a property and broadcasts PROPERTY_CHANGE.
func (h *ObjectHandle) Set(propName string, value interface{}) {
	h.object.SetProperty(propName, value)
}

// Emit sends a SIGNAL to all linked clients.
func (h *ObjectHandle) Emit(signalName string, args ...interface{}) {
	h.object.Emit(signalName, args...)
}

// ClientCount returns the number of linked clients.
func (h *ObjectHandle) ClientCount() int {
	return h.object.ClientCount()
}

// ToValue converts the ObjectHandle to a JavaScript object.
func (h *ObjectHandle) ToValue(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()

	// get(propName) - Get property value
	_ = obj.Set("get", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		propName := call.Arguments[0].String()
		return vm.ToValue(h.Get(propName))
	})

	// set(propName, value) - Set property and broadcast
	_ = obj.Set("set", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}
		propName := call.Arguments[0].String()
		value := call.Arguments[1].Export()
		h.Set(propName, value)
		return goja.Undefined()
	})

	// emit(signalName, ...args) - Send signal to linked clients
	_ = obj.Set("emit", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		signalName := call.Arguments[0].String()
		var args []interface{}
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Arguments[i].Export())
		}
		h.Emit(signalName, args...)
		return goja.Undefined()
	})

	// clientCount() - Get number of linked clients
	_ = obj.Set("clientCount", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(h.ClientCount())
	})

	return obj
}

// BackendHandle provides the JavaScript API for registering objects.
type BackendHandle struct {
	url    string
	engine *Engine
}

// NewBackendHandle creates a new backend handle.
func NewBackendHandle(url string, engine *Engine) *BackendHandle {
	return &BackendHandle{
		url:    url,
		engine: engine,
	}
}

// ToValue converts the BackendHandle to a JavaScript object.
func (h *BackendHandle) ToValue(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()

	// register(objectId, config) - Register an object with handlers, returns object handle
	_ = obj.Set("register", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("register requires (objectId, config) arguments"))
		}
		objectID := call.Arguments[0].String()
		config := call.Arguments[1].ToObject(vm)

		objDef := h.registerObject(vm, objectID, config)
		handle := NewObjectHandle(objDef, h.engine)
		return handle.ToValue(vm)
	})

	// unregister(objectId) - Unregister an object
	_ = obj.Set("unregister", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("unregister requires objectId argument"))
		}
		objectID := call.Arguments[0].String()
		h.engine.UnregisterObject(objectID)
		return goja.Undefined()
	})

	// object(objectId) - Get handle to a registered object
	_ = obj.Set("object", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("object requires objectId argument"))
		}
		objectID := call.Arguments[0].String()
		objDef := h.engine.GetObject(objectID)
		if objDef == nil {
			return goja.Undefined()
		}
		handle := NewObjectHandle(objDef, h.engine)
		return handle.ToValue(vm)
	})

	// url property
	_ = obj.Set("url", h.url)

	// onMessage(callback) - Register raw message handler
	// callback(msg, ctx) where ctx has send(msg) function
	_ = obj.Set("onMessage", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onMessage requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onMessage argument must be a function"))
		}
		h.engine.SetOnMessageHandler(callback)
		return goja.Undefined()
	})

	return obj
}

// registerObject creates and registers an object from JavaScript config.
// Returns the registered ObjectDefinition.
func (h *BackendHandle) registerObject(vm *goja.Runtime, objectID string, config *goja.Object) *ObjectDefinition {
	obj := NewObjectDefinition(objectID, h.engine)

	// Extract properties
	if propsVal := config.Get("properties"); propsVal != nil && !goja.IsUndefined(propsVal) {
		if propsObj, ok := propsVal.Export().(map[string]interface{}); ok {
			for k, v := range propsObj {
				obj.Properties[k] = v
			}
		}
	}

	// Extract methods
	if methodsVal := config.Get("methods"); methodsVal != nil && !goja.IsUndefined(methodsVal) {
		methodsObj := methodsVal.ToObject(vm)
		for _, key := range methodsObj.Keys() {
			methodVal := methodsObj.Get(key)
			if fn, ok := goja.AssertFunction(methodVal); ok {
				obj.Methods[key] = fn
			}
		}
	}

	// Extract onLink callback
	if onLinkVal := config.Get("onLink"); onLinkVal != nil && !goja.IsUndefined(onLinkVal) {
		if fn, ok := goja.AssertFunction(onLinkVal); ok {
			obj.OnLink = fn
		}
	}

	// Extract onUnlink callback
	if onUnlinkVal := config.Get("onUnlink"); onUnlinkVal != nil && !goja.IsUndefined(onUnlinkVal) {
		if fn, ok := goja.AssertFunction(onUnlinkVal); ok {
			obj.OnUnlink = fn
		}
	}

	// Extract onSetProperty callback
	if onSetPropVal := config.Get("onSetProperty"); onSetPropVal != nil && !goja.IsUndefined(onSetPropVal) {
		if fn, ok := goja.AssertFunction(onSetPropVal); ok {
			obj.OnSetProperty = fn
		}
	}

	h.engine.RegisterObject(obj)
	return obj
}
