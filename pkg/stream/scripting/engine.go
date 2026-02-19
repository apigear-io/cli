package scripting

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dop251/goja"
)

// Engine wraps a Goja VM for running JavaScript scripts.
type Engine struct {
	vm *goja.Runtime

	// WebSocket streams (connect() API)
	wsStreams   []*WSStream
	wsStreamsMu sync.Mutex

	// Backend server support (createBackend() API)
	backendServer   BackendServerInterface
	backendServerMu sync.RWMutex
	objects         map[string]*ObjectDefinition
	objectsMu       sync.RWMutex

	// Stats recording (optional, for backend servers)
	stats StatsRecorder

	// Output handling
	outputCh chan OutputEntry

	// Context for cancellation
	ctx    context.Context
	cancel context.CancelFunc

	// Script identification
	id   string
	name string

	// Callback scheduling
	callbackCh chan func()
	wg         sync.WaitGroup

	// Run mode
	onStopCallback func() // Called when engine stops
	stopped        atomic.Bool
	stoppedMu      sync.Mutex // Only protects onStopCallback

	// onMessage handler for raw message handling
	onMessageHandler goja.Callable
	onMessageMu      sync.RWMutex
}

// BackendServerInterface defines methods the engine needs from a backend server.
type BackendServerInterface interface {
	Start() error
	Stop() error
	BroadcastToLinked(objectID string, msg interface{})
	GetLinkedClientCount(objectID string) int
}

// StatsRecorder defines methods for recording connection and message statistics.
// This interface decouples the scripting package from the proxy package.
type StatsRecorder interface {
	ConnectionOpened()
	ConnectionClosed()
	MessageIn(bytes int)
	MessageOut(bytes int)
}

// BackendServerFactory creates backend servers.
type BackendServerFactory func(name, listenAddr string, engine *Engine, stats StatsRecorder) BackendServerInterface

// DefaultBackendServerFactory is set by the backend package during init.
var DefaultBackendServerFactory BackendServerFactory

// OutputEntry represents a console output entry.
type OutputEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

// NewEngine creates a new scripting engine.
func NewEngine(id, name string) *Engine {
	ctx, cancel := context.WithCancel(context.Background())

	e := &Engine{
		vm:         goja.New(),
		objects:    make(map[string]*ObjectDefinition),
		outputCh:   make(chan OutputEntry, 100),
		ctx:        ctx,
		cancel:     cancel,
		id:         id,
		name:       name,
		callbackCh: make(chan func(), 100),
	}

	return e
}

// ID returns the engine's unique identifier.
func (e *Engine) ID() string {
	return e.id
}

// Name returns the script name.
func (e *Engine) Name() string {
	return e.name
}

// Output returns the channel for receiving console output.
func (e *Engine) Output() <-chan OutputEntry {
	return e.outputCh
}

// Run executes a JavaScript script synchronously.
func (e *Engine) Run(script string) error {
	// Set up console
	e.setupConsole(e.vm)

	// Register global functions
	e.registerGlobals(e.vm)

	// Execute the script
	_, err := e.vm.RunString(script)
	return err
}

// RunWithResult executes a JavaScript script and returns the result value.
func (e *Engine) RunWithResult(script string) (goja.Value, error) {
	// Set up console
	e.setupConsole(e.vm)

	// Register global functions
	e.registerGlobals(e.vm)

	// Execute the script and return result
	return e.vm.RunString(script)
}

// RunAsync executes a script and processes callbacks until stopped.
func (e *Engine) RunAsync(script string) error {
	// Set up console
	e.setupConsole(e.vm)

	// Register global functions
	e.registerGlobals(e.vm)

	// Execute the script
	_, err := e.vm.RunString(script)
	if err != nil {
		// Output error to console before returning
		e.writeOutput("error", fmt.Sprintf("Script error: %v", err))
		return err
	}

	// Start callback processor
	e.wg.Add(1)
	go e.processCallbacks()

	return nil
}

// processCallbacks handles scheduled callbacks from timers and message handlers.
func (e *Engine) processCallbacks() {
	defer e.wg.Done()

	for {
		select {
		case <-e.ctx.Done():
			// Drain any remaining callbacks before exiting
			for {
				select {
				case fn := <-e.callbackCh:
					func() {
						defer func() { recover() }()
						fn()
					}()
				default:
					return
				}
			}
		case fn := <-e.callbackCh:
			// Wrap callback execution to catch panics and output to console
			func() {
				defer func() {
					if r := recover(); r != nil {
						e.writeOutput("error", fmt.Sprintf("Callback error: %v", r))
					}
				}()
				fn()
			}()
		}
	}
}

// Stop terminates the script execution.
func (e *Engine) Stop() {
	// Use atomic swap to ensure only one goroutine proceeds
	if e.stopped.Swap(true) {
		return // Already stopped
	}

	// Get callback under lock
	e.stoppedMu.Lock()
	callback := e.onStopCallback
	e.stoppedMu.Unlock()

	e.cancel()
	if e.vm != nil {
		e.vm.Interrupt("stopped")
	}

	// Close all WebSocket streams
	e.wsStreamsMu.Lock()
	for _, ws := range e.wsStreams {
		ws.Close()
	}
	e.wsStreams = nil
	e.wsStreamsMu.Unlock()

	// Stop backend server if running
	e.backendServerMu.RLock()
	server := e.backendServer
	e.backendServerMu.RUnlock()
	if server != nil {
		server.Stop()
	}

	// Wait for callback processor to finish (it exits on ctx.Done())
	e.wg.Wait()

	// Now safe to close output channel - no more callbacks can write to it
	close(e.outputCh)

	if callback != nil {
		callback()
	}
}

// SetOnStopCallback sets a callback to be called when the engine stops.
func (e *Engine) SetOnStopCallback(fn func()) {
	e.stoppedMu.Lock()
	e.onStopCallback = fn
	e.stoppedMu.Unlock()
}

// SetOnMessageHandler sets the raw message handler callback.
func (e *Engine) SetOnMessageHandler(handler goja.Callable) {
	e.onMessageMu.Lock()
	e.onMessageHandler = handler
	e.onMessageMu.Unlock()
}

// GetOnMessageHandler returns the raw message handler callback.
func (e *Engine) GetOnMessageHandler() goja.Callable {
	e.onMessageMu.RLock()
	defer e.onMessageMu.RUnlock()
	return e.onMessageHandler
}

// RegisterObject registers an object definition.
func (e *Engine) RegisterObject(obj *ObjectDefinition) {
	e.objectsMu.Lock()
	defer e.objectsMu.Unlock()
	e.objects[obj.ObjectID] = obj
}

// UnregisterObject removes an object definition.
func (e *Engine) UnregisterObject(objectID string) {
	e.objectsMu.Lock()
	defer e.objectsMu.Unlock()
	delete(e.objects, objectID)
}

// GetObject returns an object definition by ID.
func (e *Engine) GetObject(objectID string) *ObjectDefinition {
	e.objectsMu.RLock()
	defer e.objectsMu.RUnlock()
	return e.objects[objectID]
}

// SetBackendServer sets the backend server reference.
func (e *Engine) SetBackendServer(server BackendServerInterface) {
	e.backendServerMu.Lock()
	e.backendServer = server
	e.backendServerMu.Unlock()
}

// GetBackendServer returns the backend server reference.
func (e *Engine) GetBackendServer() BackendServerInterface {
	e.backendServerMu.RLock()
	defer e.backendServerMu.RUnlock()
	return e.backendServer
}

// SetStats sets the stats recorder for backend servers.
func (e *Engine) SetStats(stats StatsRecorder) {
	e.stats = stats
}

// VM returns the Goja runtime (for backend server to schedule callbacks).
func (e *Engine) VM() *goja.Runtime {
	return e.vm
}

// WriteOutput exposes output writing for backend server.
func (e *Engine) WriteOutput(level, message string) {
	e.writeOutput(level, message)
}

// setupConsole creates a console object for output.
func (e *Engine) setupConsole(vm *goja.Runtime) {
	consoleObj := vm.NewObject()

	// Create console methods
	for _, method := range []string{"log", "info", "warn", "error", "debug"} {
		level := method
		_ = consoleObj.Set(method, func(call goja.FunctionCall) goja.Value {
			msg := formatConsoleArgs(call.Arguments)
			e.writeOutput(level, msg)
			return goja.Undefined()
		})
	}

	_ = vm.Set("console", consoleObj)
}

// formatConsoleArgs formats console arguments to a string.
func formatConsoleArgs(args []goja.Value) string {
	if len(args) == 0 {
		return ""
	}
	result := args[0].String()
	for i := 1; i < len(args); i++ {
		result += " " + args[i].String()
	}
	return result
}

// writeOutput sends output to the output channel.
// Safe to call even after engine is stopped.
func (e *Engine) writeOutput(level, message string) {
	// Check if engine is stopped to avoid sending on closed channel
	if e.stopped.Load() {
		return
	}

	// Use defer/recover to handle race condition where channel closes
	// between the stopped check and the send
	defer func() {
		recover() // Ignore panic from send on closed channel
	}()

	select {
	case e.outputCh <- OutputEntry{Level: level, Message: message}:
	default:
		// Drop if channel is full
	}
}

// registerGlobals registers global JavaScript functions.
func (e *Engine) registerGlobals(vm *goja.Runtime) {
	// Set up faker for random data generation
	e.setupFaker(vm)

	// Set up trace file reader
	e.registerTraceReader(vm)

	// connect(wsUrl) - creates a new WebSocket connection and returns a stream
	// Connection happens asynchronously with automatic retry on failure
	_ = vm.Set("connect", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("connect requires wsUrl argument"))
		}
		wsURL := call.Arguments[0].String()

		ws := NewWSStream(wsURL, e)

		// Track the stream for cleanup
		e.wsStreamsMu.Lock()
		e.wsStreams = append(e.wsStreams, ws)
		e.wsStreamsMu.Unlock()

		return ws.ToValue(vm)
	})

	// createBackend(wsUrl) - creates a backend server and returns a BackendHandle
	_ = vm.Set("createBackend", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("createBackend requires wsUrl argument"))
		}
		wsUrl := call.Arguments[0].String()

		if DefaultBackendServerFactory == nil {
			panic(vm.NewGoError(fmt.Errorf("backend server factory not configured")))
		}

		// Create and start the server
		server := DefaultBackendServerFactory(e.name, wsUrl, e, e.stats)
		e.SetBackendServer(server)

		// Start server in background
		go func() {
			if err := server.Start(); err != nil {
				e.writeOutput("error", "Backend server error: "+err.Error())
			}
		}()

		handle := NewBackendHandle(wsUrl, e)
		return handle.ToValue(vm)
	})

	// after(ms, callback) - executes callback after delay
	_ = vm.Set("after", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("after requires (ms, callback) arguments"))
		}
		ms := call.Arguments[0].ToInteger()
		callback, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(vm.NewTypeError("second argument must be a function"))
		}

		go func() {
			select {
			case <-e.ctx.Done():
				return
			case <-time.After(time.Duration(ms) * time.Millisecond):
				e.ScheduleCallback(func(vm *goja.Runtime) {
					_, err := callback(goja.Undefined())
					if err != nil {
						e.writeOutput("error", fmt.Sprintf("after() callback error: %v", err))
					}
				})
			}
		}()

		return goja.Undefined()
	})

	// every(ms, callback) - executes callback repeatedly at interval, returns stop function
	_ = vm.Set("every", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("every requires (ms, callback) arguments"))
		}
		ms := call.Arguments[0].ToInteger()
		callback, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(vm.NewTypeError("second argument must be a function"))
		}

		stopCh := make(chan struct{})

		go func() {
			ticker := time.NewTicker(time.Duration(ms) * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-e.ctx.Done():
					return
				case <-stopCh:
					return
				case <-ticker.C:
					e.ScheduleCallback(func(vm *goja.Runtime) {
						_, err := callback(goja.Undefined())
						if err != nil {
							e.writeOutput("error", fmt.Sprintf("every() callback error: %v", err))
						}
					})
				}
			}
		}()

		// Return a stop function
		return vm.ToValue(func() {
			close(stopCh)
		})
	})

	// print(args...) - simple output
	_ = vm.Set("print", func(call goja.FunctionCall) goja.Value {
		msg := formatConsoleArgs(call.Arguments)
		e.writeOutput("log", msg)
		return goja.Undefined()
	})

	// exit() - exits the script programmatically
	_ = vm.Set("exit", func(call goja.FunctionCall) goja.Value {
		e.writeOutput("info", "Script exiting...")
		go e.Stop()
		return goja.Undefined()
	})
}

// ScheduleCallback schedules a callback to run on the engine's goroutine.
// Safe to call even after engine is stopped.
func (e *Engine) ScheduleCallback(fn func(*goja.Runtime)) {
	// Quick check if stopped
	if e.stopped.Load() {
		return
	}

	select {
	case <-e.ctx.Done():
		return
	case e.callbackCh <- func() { fn(e.vm) }:
	default:
		// Drop if channel is full
	}
}

// CallHandler invokes a JavaScript callback and reports any errors to the console.
// Use this for event handlers where errors should be visible to the user.
func (e *Engine) CallHandler(handler goja.Callable, handlerName string, args ...goja.Value) {
	_, err := handler(goja.Undefined(), args...)
	if err != nil {
		e.writeOutput("error", fmt.Sprintf("%s error: %v", handlerName, err))
	}
}
