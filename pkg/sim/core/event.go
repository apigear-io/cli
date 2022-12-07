package core

import "time"

type APIEventType string

const (
	EventSimuStart       APIEventType = "start"
	EventSimuStop        APIEventType = "stop"
	EventCall            APIEventType = "call"
	EventReply           APIEventType = "response"
	EventSignal          APIEventType = "signal"
	EventPropertySet     APIEventType = "set"
	EventPropertyChanged APIEventType = "changed"
)

type APIEvent struct {
	Timestamp time.Time      `json:"timestamp"`
	Type      APIEventType   `json:"type"`
	Symbol    string         `json:"symbol"`
	Name      string         `json:"member"`
	Args      []any          `json:"args"`
	KWArgs    map[string]any `json:"kwargs"`
	Error     error          `json:"error"`
}
