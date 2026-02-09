package monitoring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SOURCE = "123"
	CALL   = "demo/Counter#increment"
	SIGNAL = "demo/Counter#shutdown"
	STATE  = "demo/Counter"
)

var PAYLOAD = Payload{"a": 1, "b": 2}

func TestMakeCall(t *testing.T) {
	f := NewEventFactory(SOURCE)
	// make a call object and validate content
	call := f.MakeCall(CALL, PAYLOAD)
	assert.Equal(t, CALL, call.Symbol)
	assert.Equal(t, PAYLOAD, call.Data)
}

func TestMakeSignal(t *testing.T) {
	f := NewEventFactory(SOURCE)
	// make a signal object and validate content
	signal := f.MakeSignal(SIGNAL, PAYLOAD)
	assert.Equal(t, SIGNAL, signal.Symbol)
	assert.Equal(t, PAYLOAD, signal.Data)
}

func TestMakeState(t *testing.T) {
	f := NewEventFactory(SOURCE)
	// make a state object and validate content
	state := f.MakeState(STATE, PAYLOAD)
	assert.Equal(t, STATE, state.Symbol)
	assert.Equal(t, PAYLOAD, state.Data)
}

func TestEventTypeString(t *testing.T) {
	t.Run("converts TypeCall to string", func(t *testing.T) {
		et := TypeCall
		assert.Equal(t, "call", et.String())
	})

	t.Run("converts TypeSignal to string", func(t *testing.T) {
		et := TypeSignal
		assert.Equal(t, "signal", et.String())
	})

	t.Run("converts TypeState to string", func(t *testing.T) {
		et := TypeState
		assert.Equal(t, "state", et.String())
	})

	t.Run("converts custom EventType to string", func(t *testing.T) {
		et := EventType("custom")
		assert.Equal(t, "custom", et.String())
	})
}

func TestEventSubject(t *testing.T) {
	t.Run("returns monitoring.source format", func(t *testing.T) {
		event := &Event{
			Source: "device123",
		}
		assert.Equal(t, "monitoring.device123", event.Subject())
	})

	t.Run("handles empty source", func(t *testing.T) {
		event := &Event{
			Source: "",
		}
		assert.Equal(t, "monitoring.", event.Subject())
	})

	t.Run("handles source with special characters", func(t *testing.T) {
		event := &Event{
			Source: "device-123_test",
		}
		assert.Equal(t, "monitoring.device-123_test", event.Subject())
	})
}

func TestEventFactorySanitize(t *testing.T) {
	f := NewEventFactory("default-source")

	t.Run("fills in missing source", func(t *testing.T) {
		event := &Event{
			Type:   TypeCall,
			Symbol: "test",
		}
		sanitized := f.Sanitize(event)
		assert.Equal(t, "default-source", sanitized.Source)
	})

	t.Run("preserves existing source", func(t *testing.T) {
		event := &Event{
			Type:   TypeCall,
			Symbol: "test",
			Source: "existing-source",
		}
		sanitized := f.Sanitize(event)
		assert.Equal(t, "existing-source", sanitized.Source)
	})

	t.Run("fills in missing id", func(t *testing.T) {
		event := &Event{
			Type:   TypeCall,
			Symbol: "test",
		}
		sanitized := f.Sanitize(event)
		assert.NotEmpty(t, sanitized.Id)
		// Should be a valid UUID format
		assert.Len(t, sanitized.Id, 36) // UUID length with dashes
	})

	t.Run("preserves existing id", func(t *testing.T) {
		event := &Event{
			Type:   TypeCall,
			Symbol: "test",
			Id:     "existing-id",
		}
		sanitized := f.Sanitize(event)
		assert.Equal(t, "existing-id", sanitized.Id)
	})

	t.Run("fills in missing timestamp", func(t *testing.T) {
		event := &Event{
			Type:   TypeCall,
			Symbol: "test",
		}
		sanitized := f.Sanitize(event)
		assert.False(t, sanitized.Timestamp.IsZero())
	})

	t.Run("sanitizes all missing fields at once", func(t *testing.T) {
		event := &Event{
			Type:   TypeCall,
			Symbol: "test",
			Data:   Payload{"key": "value"},
		}
		sanitized := f.Sanitize(event)
		assert.Equal(t, "default-source", sanitized.Source)
		assert.NotEmpty(t, sanitized.Id)
		assert.False(t, sanitized.Timestamp.IsZero())
		// Original fields should be preserved
		assert.Equal(t, TypeCall, sanitized.Type)
		assert.Equal(t, "test", sanitized.Symbol)
		assert.Equal(t, Payload{"key": "value"}, sanitized.Data)
	})
}
