package rpc

import "encoding/json"

type RequestHandler interface {
	HandleRequest(m Request) error
}

// Request represents a request from a client.
type Request struct {
	data []byte
	conn *Connection
	hub  *Hub
	err  error
}

// Error returns the error that occurred while processing the request.
func (r Request) Error() error {
	return r.err
}

// Reply writes a response to the client.
func (r Request) Reply(data []byte) {
	r.conn.Write(data)
}

// ReplyJSON writes a JSON response to the client.
func (r Request) ReplyJSON(data interface{}) error {
	return r.conn.WriteJSON(data)
}

// AsData returns the request data as a string.
func (r Request) AsData() []byte {
	return r.data
}

// AsJSON decodes the request data as JSON.
func (r Request) AsJSON(v interface{}) error {
	return json.Unmarshal(r.data, v)
}

// Broadcast sends data to all connections
func (r Request) Broadcast(data []byte) {
	r.hub.Broadcast(data)
}

// BroadcastJSON encodes v as JSON and broadcasts it to all connections.
func (r Request) BroadcastJSON(v interface{}) error {
	return r.hub.BroadcastJSON(v)
}
