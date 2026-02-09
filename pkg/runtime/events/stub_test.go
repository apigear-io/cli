package events

import (
	"testing"
)

// TestStubEventBusImplementsInterface verifies that StubEventBus implements IEventBus
func TestStubEventBusImplementsInterface(t *testing.T) {
	var _ IEventBus = (*StubEventBus)(nil)
}

// TestStubEventBusPublish verifies Publish is a no-op
func TestStubEventBusPublish(t *testing.T) {
	bus := NewStubEventBus()
	defer bus.Close()

	e := NewEvent("test.event", map[string]any{"key": "value"})
	err := bus.Publish(e)

	if err != nil {
		t.Errorf("Expected Publish to return nil error, got: %v", err)
	}
}

// TestStubEventBusRequest verifies Request returns an error event
func TestStubEventBusRequest(t *testing.T) {
	bus := NewStubEventBus()
	defer bus.Close()

	e := NewEvent("test.request", map[string]any{"key": "value"})
	resp, err := bus.Request(e)

	if err != nil {
		t.Errorf("Expected Request to return nil error, got: %v", err)
	}

	if resp == nil {
		t.Fatal("Expected Request to return a response event, got nil")
	}

	if resp.Error == "" {
		t.Error("Expected response event to have an error message")
	}

	if resp.Kind != "test.request" {
		t.Errorf("Expected response kind to be 'test.request', got: %s", resp.Kind)
	}
}

// TestStubEventBusRegister verifies Register is a no-op
func TestStubEventBusRegister(t *testing.T) {
	bus := NewStubEventBus()
	defer bus.Close()

	called := false
	handler := func(e *Event) (*Event, error) {
		called = true
		return e, nil
	}

	// Register should not panic
	bus.Register("test.event", handler)

	// Handler won't be called in stub implementation
	if called {
		t.Error("Handler should not be called in stub implementation")
	}
}

// TestStubEventBusUse verifies Use is a no-op
func TestStubEventBusUse(t *testing.T) {
	bus := NewStubEventBus()
	defer bus.Close()

	called := false
	middleware := func(e *Event) (*Event, error) {
		called = true
		return e, nil
	}

	// Use should not panic
	bus.Use(middleware)

	// Middleware won't be called in stub implementation
	if called {
		t.Error("Middleware should not be called in stub implementation")
	}
}

// TestStubEventBusClose verifies Close is a no-op
func TestStubEventBusClose(t *testing.T) {
	bus := NewStubEventBus()

	err := bus.Close()
	if err != nil {
		t.Errorf("Expected Close to return nil error, got: %v", err)
	}

	// Should be safe to close multiple times
	err = bus.Close()
	if err != nil {
		t.Errorf("Expected second Close to return nil error, got: %v", err)
	}
}

// TestStubEventBusConcurrency verifies thread safety
func TestStubEventBusConcurrency(t *testing.T) {
	bus := NewStubEventBus()
	defer bus.Close()

	// Run operations concurrently to check for race conditions
	done := make(chan bool, 3)

	go func() {
		for i := 0; i < 100; i++ {
			bus.Publish(NewEvent("test", nil))
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			bus.Register("test", func(e *Event) (*Event, error) { return e, nil })
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			bus.Use(func(e *Event) (*Event, error) { return e, nil })
		}
		done <- true
	}()

	// Wait for all goroutines to complete
	<-done
	<-done
	<-done
}
