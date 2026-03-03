package relay

import "github.com/gorilla/websocket"

// Message types for WebSocket messages.
// These wrap the gorilla/websocket constants to avoid exposing
// the dependency to consuming packages.
const (
	// TextMessage denotes a text data message.
	TextMessage = websocket.TextMessage

	// BinaryMessage denotes a binary data message.
	BinaryMessage = websocket.BinaryMessage

	// CloseMessage denotes a close control message.
	CloseMessage = websocket.CloseMessage

	// PingMessage denotes a ping control message.
	PingMessage = websocket.PingMessage

	// PongMessage denotes a pong control message.
	PongMessage = websocket.PongMessage
)

// Close codes for WebSocket close messages.
const (
	CloseNormalClosure           = websocket.CloseNormalClosure
	CloseGoingAway              = websocket.CloseGoingAway
	CloseProtocolError          = websocket.CloseProtocolError
	CloseUnsupportedData        = websocket.CloseUnsupportedData
	CloseNoStatusReceived       = websocket.CloseNoStatusReceived
	CloseAbnormalClosure        = websocket.CloseAbnormalClosure
	CloseInvalidFramePayloadData = websocket.CloseInvalidFramePayloadData
	ClosePolicyViolation        = websocket.ClosePolicyViolation
	CloseMessageTooBig          = websocket.CloseMessageTooBig
	CloseMandatoryExtension     = websocket.CloseMandatoryExtension
	CloseInternalServerErr      = websocket.CloseInternalServerErr
	CloseServiceRestart         = websocket.CloseServiceRestart
	CloseTryAgainLater          = websocket.CloseTryAgainLater
	CloseTLSHandshake           = websocket.CloseTLSHandshake
)
