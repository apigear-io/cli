package stream

import (
	"encoding/json"

	"github.com/apigear-io/cli/pkg/runtime/monitoring"
	"github.com/apigear-io/cli/pkg/stream/protocol"
)

// EventAdapter connects stream events to ApiGear's monitoring system.
// It converts proxy and client events into monitoring.Event format
// and emits them through monitoring.Emitter for unified observability.
type EventAdapter struct {
	factory *monitoring.EventFactory
	enabled bool
}

// NewEventAdapter creates a new event adapter.
func NewEventAdapter(source string) *EventAdapter {
	return &EventAdapter{
		factory: monitoring.NewEventFactory(source),
		enabled: true,
	}
}

// SetEnabled enables or disables event emission.
func (a *EventAdapter) SetEnabled(enabled bool) {
	a.enabled = enabled
}

// ProxyMessage emits an event for a proxy message.
// Converts the raw message to a monitoring event with appropriate type.
func (a *EventAdapter) ProxyMessage(proxyName string, direction string, msgData []byte) {
	if !a.enabled {
		return
	}

	// Parse message to extract details
	parsed := protocol.ParseMessage(msgData)

	// Use symbol as event symbol, fall back to proxy name
	symbol := parsed.Symbol
	if symbol == "" {
		symbol = "stream." + proxyName
	}

	// Build event data
	data := monitoring.Payload{
		"proxy":     proxyName,
		"direction": direction,
		"msgType":   parsed.MsgTypeName,
		"timestamp": parsed.Timestamp,
	}

	// Add message-specific fields
	if parsed.Symbol != "" {
		data["objectId"] = parsed.Symbol
	}
	if parsed.RequestID != nil && *parsed.RequestID > 0 {
		data["requestId"] = *parsed.RequestID
	}
	if parsed.Args != nil {
		data["args"] = parsed.Args
	}

	// Determine event type based on message type and direction
	var event *monitoring.Event
	switch direction {
	case "SEND":
		// Outgoing messages are signals (responses, broadcasts)
		event = a.factory.MakeSignal(symbol, data)
	case "RECV":
		// Incoming messages are calls (requests)
		event = a.factory.MakeCall(symbol, data)
	default:
		// Unknown direction, treat as signal
		event = a.factory.MakeSignal(symbol, data)
	}

	monitoring.Emitter.FireHook(event)
}

// ProxyStateChange emits an event for proxy state changes.
func (a *EventAdapter) ProxyStateChange(proxyName, state string, details monitoring.Payload) {
	if !a.enabled {
		return
	}

	symbol := "stream.proxy." + proxyName

	data := monitoring.Payload{
		"proxy": proxyName,
		"state": state,
	}

	// Merge additional details
	for k, v := range details {
		data[k] = v
	}

	event := a.factory.MakeState(symbol, data)
	monitoring.Emitter.FireHook(event)
}

// ClientStateChange emits an event for client state changes.
func (a *EventAdapter) ClientStateChange(clientName, state string, details monitoring.Payload) {
	if !a.enabled {
		return
	}

	symbol := "stream.client." + clientName

	data := monitoring.Payload{
		"client": clientName,
		"state":  state,
	}

	// Merge additional details
	for k, v := range details {
		data[k] = v
	}

	event := a.factory.MakeState(symbol, data)
	monitoring.Emitter.FireHook(event)
}

// ProxyStats emits an event for proxy statistics.
func (a *EventAdapter) ProxyStats(proxyName string, stats ProxyStats) {
	if !a.enabled {
		return
	}

	symbol := "stream.proxy." + proxyName

	data := monitoring.Payload{
		"proxy":              proxyName,
		"messagesReceived":   stats.MessagesReceived,
		"messagesSent":       stats.MessagesSent,
		"activeConnections":  stats.ActiveConnections,
		"bytesReceived":      stats.BytesReceived,
		"bytesSent":          stats.BytesSent,
		"uptimeSeconds":      stats.UptimeSeconds,
	}

	event := a.factory.MakeState(symbol, data)
	monitoring.Emitter.FireHook(event)
}

// ProxyStats represents proxy statistics for event emission.
type ProxyStats struct {
	MessagesReceived   int64
	MessagesSent       int64
	ActiveConnections  int
	BytesReceived      int64
	BytesSent          int64
	UptimeSeconds      int64
}

// ScriptOutput emits an event for script console output.
func (a *EventAdapter) ScriptOutput(scriptName, level, message string) {
	if !a.enabled {
		return
	}

	symbol := "stream.script." + scriptName

	data := monitoring.Payload{
		"script":  scriptName,
		"level":   level,
		"message": message,
	}

	event := a.factory.MakeSignal(symbol, data)
	monitoring.Emitter.FireHook(event)
}

// TraceEvent emits an event for trace file operations.
func (a *EventAdapter) TraceEvent(operation, filename string, details monitoring.Payload) {
	if !a.enabled {
		return
	}

	symbol := "stream.trace." + operation

	data := monitoring.Payload{
		"operation": operation,
		"filename":  filename,
	}

	// Merge additional details
	for k, v := range details {
		data[k] = v
	}

	event := a.factory.MakeSignal(symbol, data)
	monitoring.Emitter.FireHook(event)
}

// ParsedMessageEvent represents a parsed message event for SSE.
type ParsedMessageEvent struct {
	Type      string          `json:"type"`
	Proxy     string          `json:"proxy"`
	Direction string          `json:"direction"`
	Timestamp int64           `json:"timestamp"`
	Message   json.RawMessage `json:"message"`
	Parsed    *ParsedMessage  `json:"parsed,omitempty"`
}

// ParsedMessage contains parsed message details.
type ParsedMessage struct {
	MsgType     int         `json:"msgType"`
	MsgTypeName string      `json:"msgTypeName"`
	Symbol      string      `json:"symbol,omitempty"`
	ObjectID    string      `json:"objectId,omitempty"`
	RequestID   int64       `json:"requestId,omitempty"`
	Args        interface{} `json:"args,omitempty"`
}

// ConvertToParsedMessageEvent converts a protocol.ParsedMessage to an event.
func ConvertToParsedMessageEvent(proxyName string, direction string, msgData []byte) ParsedMessageEvent {
	parsed := protocol.ParseMessage(msgData)

	event := ParsedMessageEvent{
		Type:      "message",
		Proxy:     proxyName,
		Direction: direction,
		Timestamp: parsed.Timestamp,
		Message:   msgData,
	}

	if parsed.MsgType > 0 {
		pm := &ParsedMessage{
			MsgType:     parsed.MsgType,
			MsgTypeName: parsed.MsgTypeName,
			Symbol:      parsed.Symbol,
			ObjectID:    parsed.Symbol, // Use Symbol as ObjectID
			Args:        parsed.Args,
		}
		if parsed.RequestID != nil {
			pm.RequestID = int64(*parsed.RequestID)
		}
		event.Parsed = pm
	}

	return event
}
