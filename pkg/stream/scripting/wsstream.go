package scripting

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/dop251/goja"
)

// TODO: Full ObjectLink integration with objectlink-core-go
// This is a working stub implementation that provides the JavaScript API.
// The full implementation will use objectlink-core-go for:
// - Client connection management
// - Message routing (LINK, INIT, INVOKE, SIGNAL, PROPERTY_CHANGE, etc.)
// - Auto-reconnect behavior
// - Status notifications

// pendingInvoke tracks a pending invoke request.
type pendingInvoke struct {
	methodID string
	promise  *Promise
}

// WSStream represents a WebSocket-connected stream for ObjectLink communication.
type WSStream struct {
	url    string
	engine *Engine

	// Connection state
	connected atomic.Bool
	ctx       context.Context
	cancel    context.CancelFunc

	// Request ID counter for invoke calls
	requestID atomic.Int64

	// Pending invoke requests: requestID -> promise resolver
	pendingInvokes   map[int64]*pendingInvoke
	pendingInvokesMu sync.RWMutex

	// Event handlers
	onInitHandlers           []goja.Callable
	onPropertyChangeHandlers []goja.Callable
	onSignalHandlers         []goja.Callable
	onErrorHandlers          []goja.Callable
	onConnectHandlers        []goja.Callable
	onDisconnectHandlers     []goja.Callable
	onMessageHandler         goja.Callable // Raw message handler
	handlersMu               sync.RWMutex

	// Interfaces created from this stream
	interfaces   map[string]*InterfaceHandle
	interfacesMu sync.RWMutex

	wg sync.WaitGroup
}

// Counter for generating unique client names
var streamCounter int64

// NewWSStream creates a new WebSocket stream and starts connecting to the given URL.
// Connection happens asynchronously with automatic retry on failure.
func NewWSStream(url string, engine *Engine) *WSStream {
	ctx, cancel := context.WithCancel(engine.ctx)

	ws := &WSStream{
		url:            url,
		engine:         engine,
		ctx:            ctx,
		cancel:         cancel,
		pendingInvokes: make(map[int64]*pendingInvoke),
		interfaces:     make(map[string]*InterfaceHandle),
	}

	// TODO: Initialize objectlink-core-go client here
	// - Create client with unique name
	// - Subscribe to messages and status updates
	// - Start connection with auto-reconnect

	// For now, mark as disconnected
	ws.connected.Store(false)

	return ws
}

// sendMessage sends a message over the WebSocket connection.
func (ws *WSStream) sendMessage(msg interface{}) error {
	// TODO: Send via objectlink-core-go client
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_ = data
	return fmt.Errorf("not implemented - needs objectlink-core-go integration")
}

// Link sends a LINK message.
func (ws *WSStream) Link(objectID string) {
	// TODO: Send via objectlink-core-go client
	_ = objectID
}

// Unlink sends an UNLINK message.
func (ws *WSStream) Unlink(objectID string) {
	// TODO: Send via objectlink-core-go client
	_ = objectID
}

// SetProperty sends a SET_PROPERTY message.
func (ws *WSStream) SetProperty(propertyID string, value interface{}) {
	// TODO: Send via objectlink-core-go client
	_ = propertyID
	_ = value
}

// Invoke sends an INVOKE message and returns a Promise.
func (ws *WSStream) Invoke(vm *goja.Runtime, methodID string, args []interface{}) goja.Value {
	reqID := ws.requestID.Add(1)
	promise := NewPromise(ws.engine)

	ws.pendingInvokesMu.Lock()
	ws.pendingInvokes[reqID] = &pendingInvoke{
		methodID: methodID,
		promise:  promise,
	}
	ws.pendingInvokesMu.Unlock()

	// TODO: Send via objectlink-core-go client
	// For now, reject the promise immediately
	go func() {
		promise.Reject(vm.ToValue("not implemented - needs objectlink-core-go integration"))
	}()

	return promise.ToValue(vm)
}

// Interface returns or creates an InterfaceHandle for the given object ID.
func (ws *WSStream) Interface(objectID string) *InterfaceHandle {
	ws.interfacesMu.Lock()
	defer ws.interfacesMu.Unlock()

	if iface, ok := ws.interfaces[objectID]; ok {
		return iface
	}

	iface := NewInterfaceHandle(objectID, ws)
	ws.interfaces[objectID] = iface
	return iface
}

// Close closes the WebSocket connection.
func (ws *WSStream) Close() {
	// TODO: Stop objectlink-core-go client
	ws.cancel()
	ws.wg.Wait()
}

// ScheduleCallback schedules a callback to run on the engine's goroutine.
// This satisfies the StreamSender interface.
func (ws *WSStream) ScheduleCallback(fn func(*goja.Runtime)) {
	ws.engine.ScheduleCallback(fn)
}

// OnInit registers a handler for INIT messages.
func (ws *WSStream) OnInit(handler goja.Callable) {
	ws.handlersMu.Lock()
	ws.onInitHandlers = append(ws.onInitHandlers, handler)
	ws.handlersMu.Unlock()
}

// OnPropertyChange registers a handler for PROPERTY_CHANGE messages.
func (ws *WSStream) OnPropertyChange(handler goja.Callable) {
	ws.handlersMu.Lock()
	ws.onPropertyChangeHandlers = append(ws.onPropertyChangeHandlers, handler)
	ws.handlersMu.Unlock()
}

// OnSignal registers a handler for SIGNAL messages.
func (ws *WSStream) OnSignal(handler goja.Callable) {
	ws.handlersMu.Lock()
	ws.onSignalHandlers = append(ws.onSignalHandlers, handler)
	ws.handlersMu.Unlock()
}

// OnError registers a handler for ERROR messages.
func (ws *WSStream) OnError(handler goja.Callable) {
	ws.handlersMu.Lock()
	ws.onErrorHandlers = append(ws.onErrorHandlers, handler)
	ws.handlersMu.Unlock()
}

// OnConnect registers a handler for connection events.
func (ws *WSStream) OnConnect(handler goja.Callable) {
	ws.handlersMu.Lock()
	ws.onConnectHandlers = append(ws.onConnectHandlers, handler)
	ws.handlersMu.Unlock()
}

// OnDisconnect registers a handler for disconnection events.
func (ws *WSStream) OnDisconnect(handler goja.Callable) {
	ws.handlersMu.Lock()
	ws.onDisconnectHandlers = append(ws.onDisconnectHandlers, handler)
	ws.handlersMu.Unlock()
}

// ToValue converts the WSStream to a JavaScript object.
func (ws *WSStream) ToValue(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()

	// link(objectId) - Send LINK message
	_ = obj.Set("link", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("link requires objectId argument"))
		}
		objectID := call.Arguments[0].String()
		ws.Link(objectID)
		return goja.Undefined()
	})

	// unlink(objectId) - Send UNLINK message
	_ = obj.Set("unlink", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("unlink requires objectId argument"))
		}
		objectID := call.Arguments[0].String()
		ws.Unlink(objectID)
		return goja.Undefined()
	})

	// setProperty(propertyId, value) - Send SET_PROPERTY message
	_ = obj.Set("setProperty", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("setProperty requires (propertyId, value) arguments"))
		}
		propertyID := call.Arguments[0].String()
		value := call.Arguments[1].Export()
		ws.SetProperty(propertyID, value)
		return goja.Undefined()
	})

	// invoke(methodId, ...args) - Send INVOKE message, returns Promise
	_ = obj.Set("invoke", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("invoke requires methodId argument"))
		}
		methodID := call.Arguments[0].String()
		args := make([]interface{}, len(call.Arguments)-1)
		for i := 1; i < len(call.Arguments); i++ {
			args[i-1] = call.Arguments[i].Export()
		}
		return ws.Invoke(vm, methodID, args)
	})

	// onInit(callback) - Register INIT handler
	_ = obj.Set("onInit", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onInit requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onInit argument must be a function"))
		}
		ws.OnInit(callback)
		return goja.Undefined()
	})

	// onPropertyChange(callback) - Register PROPERTY_CHANGE handler
	_ = obj.Set("onPropertyChange", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onPropertyChange requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onPropertyChange argument must be a function"))
		}
		ws.OnPropertyChange(callback)
		return goja.Undefined()
	})

	// onSignal(callback) - Register SIGNAL handler
	_ = obj.Set("onSignal", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onSignal requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onSignal argument must be a function"))
		}
		ws.OnSignal(callback)
		return goja.Undefined()
	})

	// onError(callback) - Register ERROR handler
	_ = obj.Set("onError", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onError requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onError argument must be a function"))
		}
		ws.OnError(callback)
		return goja.Undefined()
	})

	// onConnect(callback) - Register connect handler
	_ = obj.Set("onConnect", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onConnect requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onConnect argument must be a function"))
		}
		ws.OnConnect(callback)
		return goja.Undefined()
	})

	// onDisconnect(callback) - Register disconnect handler
	_ = obj.Set("onDisconnect", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onDisconnect requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onDisconnect argument must be a function"))
		}
		ws.OnDisconnect(callback)
		return goja.Undefined()
	})

	// interface(objectId) - Create InterfaceHandle wrapper
	_ = obj.Set("interface", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("interface requires objectId argument"))
		}
		objectID := call.Arguments[0].String()
		iface := ws.Interface(objectID)
		return iface.ToValue(vm)
	})

	// close() - Close the connection
	_ = obj.Set("close", func(call goja.FunctionCall) goja.Value {
		ws.Close()
		return goja.Undefined()
	})

	// send(msg) - Send raw message
	_ = obj.Set("send", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("send requires message argument"))
		}
		msg := call.Arguments[0].Export()
		if err := ws.sendMessage(msg); err != nil {
			panic(vm.NewGoError(err))
		}
		return goja.Undefined()
	})

	// onMessage(callback) - Register raw message handler
	// When set, bypasses ObjectLink processing
	_ = obj.Set("onMessage", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("onMessage requires callback argument"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(vm.NewTypeError("onMessage argument must be a function"))
		}
		ws.handlersMu.Lock()
		ws.onMessageHandler = callback
		ws.handlersMu.Unlock()
		return goja.Undefined()
	})

	// url property
	_ = obj.Set("url", ws.url)

	// connected property (dynamic getter)
	_ = obj.DefineAccessorProperty("connected", vm.ToValue(func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(ws.connected.Load())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)

	return obj
}

// splitID splits "objectId/name" into (objectId, name).
func splitID(id string) (string, string) {
	idx := strings.LastIndexByte(id, '/')
	if idx >= 0 {
		return id[:idx], id[idx+1:]
	}
	return id, ""
}
