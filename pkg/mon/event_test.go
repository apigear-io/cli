package mon

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
