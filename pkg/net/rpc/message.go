package rpc

var (
	idSeq uint64 = 0
)

func NextId() uint64 {
	idSeq++
	return idSeq
}

type RpcRequestHandler interface {
	HandleMessage(m RpcRequest) error
}

type RpcMessageHandler interface {
	HandleMessage(m RpcMessage) error
}

type RpcRequest struct {
	Msg  RpcMessage
	Conn *Connection
}

type RpcResponse struct {
	Msg RpcMessage
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RpcMessage struct {
	Version string   `json:"version"`
	Method  string   `json:"method"`
	Result  any      `json:"result"`
	Id      uint64   `json:"id"`
	Params  []any    `json:"params"`
	Error   RpcError `json:"error,omitempty"`
}

func MakeError(code int, msg string) RpcMessage {
	return RpcMessage{
		Version: "2.0",
		Error:   RpcError{Code: code, Message: msg},
	}
}

func MakeCall(method string, params []any) RpcMessage {
	return RpcMessage{
		Version: "2.0",
		Method:  method,
		Id:      NextId(),
		Params:  params,
	}
}

func MakeNotify(method string, params []any) RpcMessage {
	return RpcMessage{
		Version: "2.0",
		Method:  method,
		Params:  params,
	}
}

func MakeResult(id uint64, result any) RpcMessage {
	return RpcMessage{
		Version: "2.0",
		Id:      id,
		Result:  result,
	}
}