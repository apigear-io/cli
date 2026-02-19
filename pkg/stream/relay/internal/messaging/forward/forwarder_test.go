package forward

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// mockConnection is a test double for Connection interface.
type mockConnection struct {
	messages  []Message
	writeErr  error
	readErr   error
	readIndex int
	written   []Message
	done      chan struct{}
	mu        sync.Mutex
}

func newMockConnection(messages []Message) *mockConnection {
	return &mockConnection{
		messages: messages,
		done:     make(chan struct{}),
		written:  make([]Message, 0),
	}
}

func (m *mockConnection) ReadMessage() (messageType int, p []byte, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.readErr != nil {
		return 0, nil, m.readErr
	}

	if m.readIndex >= len(m.messages) {
		return 0, nil, errors.New("no more messages")
	}

	msg := m.messages[m.readIndex]
	m.readIndex++
	return msg.Type, msg.Data, nil
}

func (m *mockConnection) WriteMessage(messageType int, data []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.writeErr != nil {
		return m.writeErr
	}

	m.written = append(m.written, Message{Type: messageType, Data: data})
	return nil
}

func (m *mockConnection) Done() <-chan struct{} {
	return m.done
}

func (m *mockConnection) Close() error {
	close(m.done)
	return nil
}

func (m *mockConnection) ID() string {
	return "mock-connection"
}

func (m *mockConnection) getWritten() []Message {
	m.mu.Lock()
	defer m.mu.Unlock()
	result := make([]Message, len(m.written))
	copy(result, m.written)
	return result
}

func TestDirectForwarder(t *testing.T) {
	messages := []Message{
		{Type: 1, Data: []byte("hello")},
		{Type: 1, Data: []byte("world")},
	}

	src := newMockConnection(messages)
	dst := newMockConnection(nil)

	var received []Message
	f := NewForwarder(Options{
		OnMessage: func(msg Message) {
			received = append(received, msg)
		},
	})

	// Forward will return error when src runs out of messages
	_ = f.Forward(src, dst)

	written := dst.getWritten()
	if len(written) != 2 {
		t.Fatalf("expected 2 messages written, got %d", len(written))
	}

	if string(written[0].Data) != "hello" {
		t.Errorf("expected 'hello', got %s", string(written[0].Data))
	}

	if string(written[1].Data) != "world" {
		t.Errorf("expected 'world', got %s", string(written[1].Data))
	}

	if len(received) != 2 {
		t.Errorf("expected 2 OnMessage calls, got %d", len(received))
	}
}

func TestDirectForwarder_WriteError(t *testing.T) {
	messages := []Message{
		{Type: 1, Data: []byte("hello")},
	}

	src := newMockConnection(messages)
	dst := newMockConnection(nil)
	dst.writeErr = errors.New("write failed")

	f := NewForwarder(Options{})
	err := f.Forward(src, dst)

	if err == nil || err.Error() != "write failed" {
		t.Errorf("expected write error, got %v", err)
	}
}

func TestDelayedForwarder(t *testing.T) {
	messages := []Message{
		{Type: 1, Data: []byte("delayed")},
	}

	src := newMockConnection(messages)
	dst := newMockConnection(nil)

	f := NewForwarder(Options{
		Delay:  50 * time.Millisecond,
	})

	start := time.Now()
	_ = f.Forward(src, dst)
	elapsed := time.Since(start)

	// Should have delayed by approximately 50ms
	if elapsed < 40*time.Millisecond {
		t.Errorf("expected delay of ~50ms, got %v", elapsed)
	}

	written := dst.getWritten()
	if len(written) != 1 {
		t.Fatalf("expected 1 message written, got %d", len(written))
	}

	if string(written[0].Data) != "delayed" {
		t.Errorf("expected 'delayed', got %s", string(written[0].Data))
	}
}

func TestThrottledForwarder(t *testing.T) {
	// Test basic throttled forwarding
	messages := []Message{
		{Type: 1, Data: []byte("throttled")},
	}

	src := newMockConnection(messages)
	dst := newMockConnection(nil)

	f := NewForwarder(Options{
		Speed:  0.5, // Half speed
	})

	_ = f.Forward(src, dst)

	written := dst.getWritten()
	if len(written) != 1 {
		t.Fatalf("expected 1 message written, got %d", len(written))
	}

	if string(written[0].Data) != "throttled" {
		t.Errorf("expected 'throttled', got %s", string(written[0].Data))
	}
}

func TestNewForwarder_SelectsCorrectType(t *testing.T) {
	tests := []struct {
		name     string
		opts     Options
		expected string
	}{
		{
			name:     "direct with no delay",
			opts:     Options{},
			expected: "*forwarder.DirectForwarder",
		},
		{
			name:     "delayed with delay set",
			opts:     Options{Delay: time.Second},
			expected: "*forwarder.DelayedForwarder",
		},
		{
			name:     "throttled with speed set",
			opts:     Options{Speed: 0.5},
			expected: "*forwarder.ThrottledForwarder",
		},
		{
			name:     "throttled takes precedence over delay",
			opts:     Options{Delay: time.Second, Speed: 0.5},
			expected: "*forwarder.ThrottledForwarder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewForwarder(tt.opts)
			typeName := getTypeName(f)
			if typeName != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, typeName)
			}
		})
	}
}

func getTypeName(v interface{}) string {
	switch v.(type) {
	case *DirectForwarder:
		return "*forwarder.DirectForwarder"
	case *DelayedForwarder:
		return "*forwarder.DelayedForwarder"
	case *ThrottledForwarder:
		return "*forwarder.ThrottledForwarder"
	default:
		return "unknown"
	}
}

func TestDefaultBufferSize(t *testing.T) {
	f := NewForwarder(Options{})
	direct, ok := f.(*DirectForwarder)
	if !ok {
		t.Fatal("expected DirectForwarder")
	}
	if direct.opts.BufferSize != 1000 {
		t.Errorf("expected default buffer size 1000, got %d", direct.opts.BufferSize)
	}
}
