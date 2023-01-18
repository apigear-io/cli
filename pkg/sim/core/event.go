package core

import "time"

type SimuEventType string

const (
	EventSimuStart       SimuEventType = "start"
	EventSimuStop        SimuEventType = "stop"
	EventCall            SimuEventType = "call"
	EventReply           SimuEventType = "response"
	EventSignal          SimuEventType = "signal"
	EventPropertySet     SimuEventType = "set"
	EventPropertyChanged SimuEventType = "changed"
	EventError           SimuEventType = "error"
)

type SimuEvent struct {
	Type      SimuEventType  `json:"type"`
	Timestamp time.Time      `json:"timestamp"`
	Symbol    string         `json:"symbol"`
	Name      string         `json:"name"`
	Args      []any          `json:"args"`
	KWArgs    map[string]any `json:"kwargs"`
	Error     string         `json:"error"`
}
