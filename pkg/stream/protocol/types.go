// Package protocol provides best-effort ObjectLink message parsing for logging and UI display.
//
// This package does NOT implement the ObjectLink protocol semantics - it only parses
// messages to extract fields for monitoring, logging, and visualization purposes.
// For actual protocol implementation, use github.com/apigear-io/objectlink-core-go.
package protocol

// ObjectLink protocol message types
const (
	MsgLink           = 10 // [10, objectId]
	MsgInit           = 11 // [11, objectId, properties]
	MsgUnlink         = 12 // [12, objectId]
	MsgSetProperty    = 20 // [20, propertyId, value]
	MsgPropertyChange = 21 // [21, propertyId, value]
	MsgInvoke         = 30 // [30, requestId, methodId, args]
	MsgInvokeReply    = 31 // [31, requestId, methodId, value]
	MsgSignal         = 40 // [40, signalId, args]
	MsgError          = 90 // [90, origMsgType, requestId, error]
)

// MsgTypeName returns a human-readable name for a message type.
func MsgTypeName(msgType int) string {
	switch msgType {
	case MsgLink:
		return "LINK"
	case MsgInit:
		return "INIT"
	case MsgUnlink:
		return "UNLINK"
	case MsgSetProperty:
		return "SET_PROPERTY"
	case MsgPropertyChange:
		return "PROPERTY_CHANGE"
	case MsgInvoke:
		return "INVOKE"
	case MsgInvokeReply:
		return "INVOKE_REPLY"
	case MsgSignal:
		return "SIGNAL"
	case MsgError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
