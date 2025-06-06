package evt

import (
	"github.com/mitchellh/mapstructure"
)

type Event struct {
	Kind  string `json:"kind" mapstructure:"kind"`
	Value any    `json:"data" mapstructure:"value"`
	Error string `json:"error" mapstructure:"error"`
	Meta  map[string]any
}

func NewEvent(kind string, value any) *Event {
	return &Event{
		Kind:  kind,
		Value: value,
		Meta:  make(map[string]any),
	}
}

func NewErrorEvent(kind, err string) *Event {
	return &Event{
		Kind:  kind,
		Error: err,
		Meta:  make(map[string]any),
	}
}

func (e Event) Export(v any) error {
	return mapstructure.Decode(e.Value, v)
}

func (e *Event) Set(key string, value any) {
	e.Meta[key] = value
}

func (e *Event) Get(key string) any {
	return e.Meta[key]
}
