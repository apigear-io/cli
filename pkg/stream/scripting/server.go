package scripting

import (
	"fmt"
)

// TODO: Full ObjectLink backend server integration with objectlink-core-go
// This is a working stub implementation that provides the backend server API.
// The full implementation will use objectlink-core-go for:
// - Server lifecycle management (Start/Stop)
// - Client connection handling
// - Message routing (LINK, INIT, INVOKE, SET_PROPERTY, etc.)
// - Broadcasting to linked clients
// - Statistics recording

// BackendServer wraps an ObjectLink server and provides JavaScript integration.
type BackendServer struct {
	name   string
	engine *Engine
	stats  StatsRecorder
}

// NewBackendServer creates a new backend server.
func NewBackendServer(name, listenAddr string, engine *Engine, stats StatsRecorder) *BackendServer {
	s := &BackendServer{
		name:   name,
		engine: engine,
		stats:  stats,
	}

	// Set server reference in engine
	engine.SetBackendServer(s)

	// TODO: Create objectlink-core-go server
	_ = listenAddr

	return s
}

// Name returns the backend server name.
func (s *BackendServer) Name() string {
	return s.name
}

// Start begins listening for WebSocket connections.
func (s *BackendServer) Start() error {
	s.engine.writeOutput("info", fmt.Sprintf("Backend server '%s' starting (stub)", s.name))
	// TODO: Start objectlink-core-go server
	return fmt.Errorf("not implemented - needs objectlink-core-go integration")
}

// Stop gracefully shuts down the server.
func (s *BackendServer) Stop() error {
	s.engine.writeOutput("info", fmt.Sprintf("Backend server '%s' stopping (stub)", s.name))
	// TODO: Stop objectlink-core-go server
	return nil
}

// RegisterObject registers an ObjectLink object.
func (s *BackendServer) RegisterObject(obj *ObjectDefinition) error {
	// TODO: Convert ObjectDefinition to objectlink-core-go object
	// - Copy properties
	// - Wrap goja.Callable methods
	// - Wrap callbacks (OnLink, OnUnlink, OnSetProperty)
	// - Register with server
	_ = obj
	return fmt.Errorf("not implemented - needs objectlink-core-go integration")
}

// UnregisterObject removes an object from the server.
func (s *BackendServer) UnregisterObject(objectID string) error {
	// TODO: Unregister from objectlink-core-go server
	_ = objectID
	return nil
}

// GetObject returns a registered object by ID.
func (s *BackendServer) GetObject(objectID string) *ObjectDefinition {
	return s.engine.GetObject(objectID)
}

// BroadcastToLinked sends a message to all clients linked to an object.
func (s *BackendServer) BroadcastToLinked(objectID string, msg interface{}) {
	// TODO: Broadcast via objectlink-core-go server
	_ = objectID
	_ = msg
}

// GetLinkedClientCount returns the number of clients linked to an object.
func (s *BackendServer) GetLinkedClientCount(objectID string) int {
	// TODO: Query objectlink-core-go server
	_ = objectID
	return 0
}

// init function to register the backend server factory.
func init() {
	DefaultBackendServerFactory = func(name, listenAddr string, engine *Engine, stats StatsRecorder) BackendServerInterface {
		return NewBackendServer(name, listenAddr, engine, stats)
	}
}
