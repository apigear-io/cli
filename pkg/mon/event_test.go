package mon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeCall(t *testing.T) {
	f := NewEventFactory("123", "test")
	// make a call object and validate content
	call := f.MakeCall("test", "a", "b")
	assert.Equal(t, "test", call.Symbol)
	assert.Equal(t, []any{"a", "b"}, call.Params)
}

func TestMakeSignal(t *testing.T) {
	f := NewEventFactory("123", "test")
	// make a signal object and validate content
	signal := f.MakeSignal("test", "a", "b")
	assert.Equal(t, "test", signal.Symbol)
	assert.Equal(t, []any{"a", "b"}, signal.Params)
}

func TestMakeState(t *testing.T) {
	f := NewEventFactory("123", "test")
	// make a state object and validate content
	state := f.MakeState("test", map[string]any{"a": "b"})
	assert.Equal(t, "test", state.Symbol)
	assert.Equal(t, map[string]any{"a": "b"}, state.Props)
}
