package mon

import (
	"time"

	"github.com/google/uuid"
)

// Kind is the type of event.
type Kind string

const (
	KindCall   Kind = "call"
	KindSignal Kind = "signal"
	KindState  Kind = "state"
)

// Event represents an API event.
type Event struct {
	Id        string                 `json:"id" yaml:"id"`
	DeviceId  string                 `json:"device" yaml:"device"`
	Kind      Kind                   `json:"kind" yaml:"kind"`
	Timestamp time.Time              `json:"timestamp" yaml:"timestamp"`
	Source    string                 `json:"source" yaml:"source"`
	Symbol    string                 `json:"symbol" yaml:"symbol"`
	Props     map[string]interface{} `json:"state" yaml:"state"`
	Params    []any                  `json:"params" yaml:"params"`
}

// EventFactory is used to create events.
// Factory associates device ids and sources with events.
type EventFactory struct {
	DeviceId string
	Source   string
}

// NewEventFactory creates a new event factory.
func NewEventFactory(deviceId string, source string) *EventFactory {
	return &EventFactory{
		DeviceId: deviceId,
		Source:   source,
	}
}

// MakeEvent creates an event with the given kind, symbol and params.
func (f EventFactory) MakeEvent(kind Kind, symbol string, params []any, props map[string]any) *Event {
	return &Event{
		Id:        uuid.New().String(),
		DeviceId:  f.DeviceId,
		Kind:      kind,
		Timestamp: time.Now(),
		Source:    f.Source,
		Symbol:    symbol,
		Params:    params,
		Props:     props,
	}
}

// MakeCall creates a call event with the given symbol and params.
func (f EventFactory) MakeCall(symbol string, params ...any) *Event {
	return f.MakeEvent(KindCall, symbol, params, nil)
}

// MakeSignal creates a signal event with the given symbol and params.
func (f EventFactory) MakeSignal(symbol string, params ...any) *Event {
	return f.MakeEvent(KindSignal, symbol, params, nil)
}

// MakeState creates a state event with the given symbol and props.
func (f EventFactory) MakeState(symbol string, props map[string]interface{}) *Event {
	return f.MakeEvent(KindState, symbol, nil, props)
}

// Sanitize ensures events are valid and fills in missing fields.
func (f EventFactory) Sanitize(event *Event) *Event {
	if event.DeviceId == "" {
		event.DeviceId = f.DeviceId
	}
	if event.Source == "" {
		event.Source = f.Source
	}
	if event.Id == "" {
		event.Id = uuid.New().String()
	}
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}
	return event
}

var emitter = make(chan *Event)

// Emitter returns the emitter channel.
func Emitter() chan *Event {
	return emitter
}

// EmitEvents writes events to the emitter channel.
func EmitEvent(event *Event) {
	emitter <- event
}
