package mon

import "time"

type Kind string

const (
	KindCall   Kind = "call"
	KindSignal Kind = "signal"
	KindState  Kind = "state"
)

type Event struct {
	Id        string                 `json:"id" yaml:"id"`
	DeviceId  string                 `json:"device" yaml:"device"`
	Kind      Kind                   `json:"kind" yaml:"kind"`
	Timestamp time.Time              `json:"timestamp" yaml:"timestamp"`
	Source    string                 `json:"source" yaml:"source"`
	Symbol    string                 `json:"symbol" yaml:"symbol"`
	State     map[string]interface{} `json:"state" yaml:"state"`
	Params    map[string]interface{} `json:"params" yaml:"params"`
}

var emitter chan *Event = make(chan *Event, 1)

func Emitter() chan *Event {
	return emitter
}

func EmitEvent(event *Event) {
	emitter <- event
}
