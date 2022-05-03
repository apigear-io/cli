package sim

type APIEventType string
type ProgramEventType string

const (
	EventSimulationStart ProgramEventType = "simulation.start"
	EventSimulationStop  ProgramEventType = "simulation.stop"
	EventCall            APIEventType     = "call"
	EventResponse        APIEventType     = "response"
	EventSignal          APIEventType     = "signal"
	EventPropertySet     APIEventType     = "property.set"
	EventPropertyChanged APIEventType     = "property.changed"
)

type APIEvent struct {
	Type   APIEventType   `json:"type"`
	Symbol string         `json:"symbol"`
	Params []any          `json:"params"`
	Props  map[string]any `json:"props"`
	Error  error          `json:"error"`
}

type ErrorEvent struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
}

type ProgramEvent struct {
	Type    APIEventType `json:"type"`
	Message string       `json:"message"`
}

var ApiEventEmitter = make(chan APIEvent)
var ProgramEventEmitter = make(chan ProgramEvent)
var ErrorEventEmitter = make(chan ErrorEvent)
