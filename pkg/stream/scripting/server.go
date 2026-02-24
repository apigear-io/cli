package scripting

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
	"github.com/dop251/goja"
)

// BackendServer wraps an ObjectLink server and provides JavaScript integration.
type BackendServer struct {
	name       string
	listenAddr string
	engine     *Engine
	stats      StatsRecorder

	// ObjectLink components
	registry *remote.Registry
	hub      *ws.Hub
	server   *http.Server
	listener net.Listener

	// Context for lifecycle management
	ctx        context.Context
	cancelFunc context.CancelFunc

	// State tracking
	mu sync.RWMutex
}

// NewBackendServer creates a new backend server.
func NewBackendServer(name, listenAddr string, engine *Engine, stats StatsRecorder) *BackendServer {
	ctx, cancel := context.WithCancel(context.Background())

	registry := remote.NewRegistry()
	hub := ws.NewHub(ctx, registry)

	s := &BackendServer{
		name:       name,
		listenAddr: listenAddr,
		engine:     engine,
		stats:      stats,
		registry:   registry,
		hub:        hub,
		ctx:        ctx,
		cancelFunc: cancel,
	}

	// Set server reference in engine
	engine.SetBackendServer(s)

	return s
}

// Name returns the backend server name.
func (s *BackendServer) Name() string {
	return s.name
}

// Start begins listening for WebSocket connections.
func (s *BackendServer) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.server != nil {
		return fmt.Errorf("server already running")
	}

	// Parse the URL to get the address
	u, err := url.Parse(s.listenAddr)
	if err != nil {
		return fmt.Errorf("invalid listen address: %w", err)
	}

	// Create HTTP server
	mux := http.NewServeMux()
	mux.Handle(u.Path, s.hub)

	s.server = &http.Server{
		Handler: mux,
	}

	// Start listening
	listener, err := net.Listen("tcp", u.Host)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", u.Host, err)
	}
	s.listener = listener

	s.engine.writeOutput("info", fmt.Sprintf("Backend server '%s' listening on %s", s.name, s.listenAddr))

	// Start serving in background
	go func() {
		if err := s.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			s.engine.writeOutput("error", fmt.Sprintf("Backend server error: %v", err))
		}
	}()

	return nil
}

// Stop gracefully shuts down the server.
func (s *BackendServer) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.server == nil {
		return nil // Already stopped
	}

	s.engine.writeOutput("info", fmt.Sprintf("Backend server '%s' stopping", s.name))

	// Shutdown HTTP server
	if err := s.server.Shutdown(context.Background()); err != nil {
		s.engine.writeOutput("warn", fmt.Sprintf("Error shutting down server: %v", err))
	}

	// Close hub
	s.hub.Close()

	// Cancel context
	s.cancelFunc()

	s.server = nil
	s.listener = nil

	return nil
}

// RegisterObject registers an ObjectLink object.
func (s *BackendServer) RegisterObject(obj *ObjectDefinition) error {
	// Create object source wrapper
	source := &objectSource{
		objectDef: obj,
		engine:    s.engine,
	}

	// Add to registry
	if err := s.registry.AddObjectSource(source); err != nil {
		return fmt.Errorf("failed to register object: %w", err)
	}

	s.engine.writeOutput("info", fmt.Sprintf("Registered object: %s", obj.ObjectID))
	return nil
}

// UnregisterObject removes an object from the server.
func (s *BackendServer) UnregisterObject(objectID string) error {
	source := s.registry.GetObjectSource(objectID)
	if source != nil {
		s.registry.RemoveObjectSource(source)
		s.engine.writeOutput("info", fmt.Sprintf("Unregistered object: %s", objectID))
	}
	return nil
}

// GetObject returns a registered object by ID.
func (s *BackendServer) GetObject(objectID string) *ObjectDefinition {
	return s.engine.GetObject(objectID)
}

// BroadcastToLinked sends a message to all clients linked to an object.
func (s *BackendServer) BroadcastToLinked(objectID string, msg interface{}) {
	// This is called by ObjectDefinition methods like SetProperty and Emit
	// The objectlink-core-go handles broadcasting automatically through
	// registry.NotifyPropertyChange and registry.NotifySignal
	// So we don't need to do anything here
	_ = objectID
	_ = msg
}

// GetLinkedClientCount returns the number of clients linked to an object.
func (s *BackendServer) GetLinkedClientCount(objectID string) int {
	nodes := s.registry.GetRemoteNodes(objectID)
	return len(nodes)
}

// objectSource implements remote.IObjectSource to wrap ObjectDefinition.
type objectSource struct {
	objectDef *ObjectDefinition
	engine    *Engine
}

// ObjectId returns the object identifier.
func (o *objectSource) ObjectId() string {
	return o.objectDef.ObjectID
}

// Invoke calls a method on the object.
func (o *objectSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	// Extract method name from methodId (format: "objectId/methodName")
	var methodName string
	for i := len(methodId) - 1; i >= 0; i-- {
		if methodId[i] == '/' {
			methodName = methodId[i+1:]
			break
		}
	}

	if methodName == "" {
		return nil, fmt.Errorf("invalid method ID: %s", methodId)
	}

	// Get the method handler
	method, exists := o.objectDef.Methods[methodName]
	if !exists {
		return nil, fmt.Errorf("method not found: %s", methodName)
	}

	// Create context for the handler
	ctx := NewObjectContext(o.objectDef, o.engine)

	// Schedule callback to run in engine's goroutine
	resultCh := make(chan core.Any, 1)
	errorCh := make(chan error, 1)

	o.engine.ScheduleCallback(func(vm *goja.Runtime) {
		// Convert args to JavaScript object
		var params interface{}
		if len(args) > 0 {
			params = args[0]
		} else {
			params = make(map[string]interface{})
		}

		// Call the method
		result, err := method(goja.Undefined(), vm.ToValue(params), ctx.ToValue(vm))
		if err != nil {
			errorCh <- err
			return
		}

		// Export result
		if result != nil && !goja.IsUndefined(result) {
			resultCh <- result.Export()
		} else {
			resultCh <- nil
		}
	})

	// Wait for result
	select {
	case err := <-errorCh:
		return nil, err
	case result := <-resultCh:
		return result, nil
	case <-o.engine.ctx.Done():
		return nil, fmt.Errorf("engine stopped")
	}
}

// SetProperty sets a property value on the object.
func (o *objectSource) SetProperty(propertyId string, value core.Any) error {
	// Extract property name from propertyId (format: "objectId/propName")
	var propName string
	for i := len(propertyId) - 1; i >= 0; i-- {
		if propertyId[i] == '/' {
			propName = propertyId[i+1:]
			break
		}
	}

	if propName == "" {
		return fmt.Errorf("invalid property ID: %s", propertyId)
	}

	// Check if there's an onSetProperty callback
	if o.objectDef.OnSetProperty != nil {
		// Create context for the handler
		ctx := NewObjectContext(o.objectDef, o.engine)

		// Schedule callback to run in engine's goroutine
		errorCh := make(chan error, 1)

		o.engine.ScheduleCallback(func(vm *goja.Runtime) {
			_, err := o.objectDef.OnSetProperty(
				goja.Undefined(),
				vm.ToValue(propName),
				vm.ToValue(value),
				ctx.ToValue(vm),
			)
			if err != nil {
				errorCh <- err
			} else {
				errorCh <- nil
			}
		})

		// Wait for callback to complete
		select {
		case err := <-errorCh:
			if err != nil {
				return err
			}
		case <-o.engine.ctx.Done():
			return fmt.Errorf("engine stopped")
		}
	} else {
		// No callback, just set the property directly
		o.objectDef.SetProperty(propName, value)
	}

	return nil
}

// Linked is called when a client links to this object.
func (o *objectSource) Linked(objectId string, node *remote.Node) error {
	if o.objectDef.OnLink != nil {
		// Create context for the handler
		ctx := NewObjectContext(o.objectDef, o.engine)

		// Schedule callback to run in engine's goroutine
		o.engine.ScheduleCallback(func(vm *goja.Runtime) {
			o.engine.CallHandler(o.objectDef.OnLink, "onLink", ctx.ToValue(vm))
		})
	}

	return nil
}

// CollectProperties returns all current property values.
func (o *objectSource) CollectProperties() (core.KWArgs, error) {
	props := o.objectDef.GetProperties()

	// Convert to core.KWArgs (map[string]interface{})
	result := make(core.KWArgs)
	for k, v := range props {
		result[k] = v
	}

	return result, nil
}

// init function to register the backend server factory.
func init() {
	DefaultBackendServerFactory = func(name, listenAddr string, engine *Engine, stats StatsRecorder) BackendServerInterface {
		return NewBackendServer(name, listenAddr, engine, stats)
	}
}
