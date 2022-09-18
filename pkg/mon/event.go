package mon

import (
	"time"

	"github.com/google/uuid"
)

// EventType is the type of event.
type EventType string

func (e EventType) String() string {
	return string(e)
}

func ParseEventType(s string) EventType {
	return EventType(s)
}

type Payload map[string]any

const (
	TypeCall   EventType = "call"
	TypeSignal EventType = "signal"
	TypeState  EventType = "state"
)

// Event represents an API event.
type Event struct {
	Id        string    `json:"id" yaml:"id" csv:"id"`
	Source    string    `json:"source" yaml:"source" csv:"source"`
	Type      EventType `json:"type" yaml:"type" csv:"type"`
	Timestamp time.Time `json:"timestamp" yaml:"timestamp" csv:"timestamp"`
	Symbol    string    `json:"symbol" yaml:"symbol" csv:"symbol"`
	Data      Payload   `json:"data" yaml:"data" csv:"data"`
}

// EventFactory is used to create events.
// Factory associates device ids and sources with events.
type EventFactory struct {
	Source string
}

// NewEventFactory creates a new event factory.
func NewEventFactory(source string) *EventFactory {
	return &EventFactory{
		Source: source,
	}
}

// MakeEvent creates an event with the given kind, symbol and params.
func (f EventFactory) MakeEvent(kind EventType, symbol string, data Payload) *Event {
	return &Event{
		Id:        uuid.New().String(),
		Type:      kind,
		Timestamp: time.Now(),
		Source:    f.Source,
		Symbol:    symbol,
		Data:      data,
	}
}

// MakeCall creates a call event with the given symbol and params.
func (f EventFactory) MakeCall(symbol string, data Payload) *Event {
	return f.MakeEvent(TypeCall, symbol, data)
}

// MakeSignal creates a signal event with the given symbol and params.
func (f EventFactory) MakeSignal(symbol string, data Payload) *Event {
	return f.MakeEvent(TypeSignal, symbol, data)
}

// MakeState creates a state event with the given symbol and props.
func (f EventFactory) MakeState(symbol string, data Payload) *Event {
	return f.MakeEvent(TypeState, symbol, data)
}

// Sanitize ensures events are valid and fills in missing fields.
func (f EventFactory) Sanitize(event *Event) *Event {
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
