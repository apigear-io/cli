package protocol

import "encoding/json"

// ParsedMessage represents a parsed ObjectLink message with extracted fields.
// This is useful for monitoring, logging, and displaying messages in UIs.
type ParsedMessage struct {
	Timestamp   int64       `json:"ts"`          // Unix timestamp (milliseconds)
	Direction   string      `json:"dir"`         // "SEND" or "RECV"
	Proxy       string      `json:"proxy"`       // Proxy name (for multi-proxy setups)
	MsgType     int         `json:"msgType"`     // Message type code
	MsgTypeName string      `json:"msgTypeName"` // Human-readable message type
	Symbol      string      `json:"symbol"`      // ObjectID, propertyID, methodID, or signalID
	RequestID   *int        `json:"requestId"`   // Request ID for INVOKE/INVOKE_REPLY messages
	Args        any         `json:"args"`        // Arguments or result payload
}

// ParseMessage parses an ObjectLink message from JSON and extracts its fields.
// This performs best-effort parsing - if individual fields fail to parse,
// they are left empty rather than returning an error.
//
// The function handles all ObjectLink message types:
//   - LINK, UNLINK: Extracts objectID
//   - INIT: Extracts objectID and properties
//   - SET_PROPERTY, PROPERTY_CHANGE: Extracts propertyID and value
//   - INVOKE: Extracts requestID, methodID, and args
//   - INVOKE_REPLY: Extracts requestID, methodID, and result
//   - SIGNAL: Extracts signalID and args
//   - ERROR: Extracts requestID and error string
func ParseMessage(raw json.RawMessage) ParsedMessage {
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil || len(arr) == 0 {
		return ParsedMessage{MsgTypeName: "UNKNOWN"}
	}

	var msgType int
	if err := json.Unmarshal(arr[0], &msgType); err != nil {
		return ParsedMessage{MsgTypeName: "UNKNOWN"}
	}

	parsed := ParsedMessage{
		MsgType:     msgType,
		MsgTypeName: MsgTypeName(msgType),
	}

	// Best-effort parsing: ignore unmarshal errors for individual fields
	switch msgType {
	case MsgLink, MsgUnlink:
		// [10, "module.Object"]
		if len(arr) > 1 {
			_ = json.Unmarshal(arr[1], &parsed.Symbol)
		}

	case MsgInit:
		// [11, "module.Object", {properties}]
		if len(arr) > 1 {
			_ = json.Unmarshal(arr[1], &parsed.Symbol)
		}
		if len(arr) > 2 {
			var props any
			_ = json.Unmarshal(arr[2], &props)
			parsed.Args = props
		}

	case MsgSetProperty, MsgPropertyChange:
		// [20, "module.Object/property", value]
		if len(arr) > 1 {
			_ = json.Unmarshal(arr[1], &parsed.Symbol)
		}
		if len(arr) > 2 {
			var value any
			_ = json.Unmarshal(arr[2], &value)
			parsed.Args = value
		}

	case MsgInvoke:
		// [30, requestId, "module.Object/method", args]
		if len(arr) > 1 {
			var reqID int
			_ = json.Unmarshal(arr[1], &reqID)
			parsed.RequestID = &reqID
		}
		if len(arr) > 2 {
			_ = json.Unmarshal(arr[2], &parsed.Symbol)
		}
		if len(arr) > 3 {
			var args any
			_ = json.Unmarshal(arr[3], &args)
			parsed.Args = args
		}

	case MsgInvokeReply:
		// [31, requestId, "module.Object/method", result]
		if len(arr) > 1 {
			var reqID int
			_ = json.Unmarshal(arr[1], &reqID)
			parsed.RequestID = &reqID
		}
		if len(arr) > 2 {
			_ = json.Unmarshal(arr[2], &parsed.Symbol)
		}
		if len(arr) > 3 {
			var result any
			_ = json.Unmarshal(arr[3], &result)
			parsed.Args = result
		}

	case MsgSignal:
		// [40, "module.Object/signal", args]
		if len(arr) > 1 {
			_ = json.Unmarshal(arr[1], &parsed.Symbol)
		}
		if len(arr) > 2 {
			var args interface{}
			_ = json.Unmarshal(arr[2], &args)
			parsed.Args = args
		}

	case MsgError:
		// [90, origMsgType, requestId, errorString]
		if len(arr) > 2 {
			var reqID int
			_ = json.Unmarshal(arr[2], &reqID)
			parsed.RequestID = &reqID
		}
		if len(arr) > 3 {
			var errMsg string
			_ = json.Unmarshal(arr[3], &errMsg)
			parsed.Args = errMsg
		}
	}

	return parsed
}
