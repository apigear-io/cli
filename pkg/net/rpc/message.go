package rpc

import (
	"encoding/json"
)

var (
	idSeq uint64 = 0
)

func NextId() uint64 {
	idSeq++
	return idSeq
}

type MessageHandler interface {
	HandleMessage(m Message) error
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	Method string         `json:"method"`
	Id     uint64         `json:"id"`
	Params map[string]any `json:"params"`
	Result any            `json:"result"`
	Error  RpcError       `json:"error,omitempty"`
}

func MessageFromJson(data []byte, m *Message) error {
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	if m.Id == 0 {
		m.Id = NextId()
	}
	return nil
}

func MakeError(code int, msg string) Message {
	return Message{
		Error: RpcError{Code: code, Message: msg},
	}
}

func MakeCall(method string, id uint64, params map[string]any) Message {
	if id == 0 {
		id = NextId()
	}
	return Message{
		Method: method,
		Id:     id,
		Params: params,
	}
}

func MakeNotify(method string, params map[string]any) Message {
	return Message{
		Method: method,
		Params: params,
	}
}

func MakeResult(id uint64, result any) Message {
	return Message{
		Id:     id,
		Result: result,
	}
}
