package sim

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/nats-io/nats.go"
)

const (
	CmdScriptStart    = "script.start"
	CmdScriptStop     = "script.stop"
	CmdFunctionRun    = "function.run"
	ControllerSubject = "sim.controller"
)

type RpcRequest struct {
	Action       string `json:"action"`
	World        string `json:"world,omitempty"`
	Script       Script `json:"script,omitempty"`
	Function     string `json:"function,omitempty"`
	FunctionArgs []any  `json:"function_args,omitempty"`
}

type RpcResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   []byte `json:"data,omitempty"`
}

type Controller struct {
	nc  *nats.Conn
	sub *nats.Subscription
	m   *Manager
}

func NewController(nc *nats.Conn, m *Manager) (*Controller, error) {
	c := &Controller{
		nc: nc,
		m:  m,
	}
	sub, err := nc.Subscribe(ControllerSubject, c.handleMsg)
	if err != nil {
		return nil, err
	}
	c.sub = sub
	return c, nil
}

func (c *Controller) Close() error {
	if c.sub != nil {
		c.nc.Drain()
		return c.sub.Unsubscribe()
	}
	return nil
}

func (c *Controller) handleMsg(msg *nats.Msg) {
	var req RpcRequest
	err := json.Unmarshal(msg.Data, &req)
	if err != nil {
		c.replyError(msg, "invalid request")
		return
	}
	switch req.Action {
	case CmdScriptStart:
		resp := c.handleStart(req)
		c.respond(msg, resp)
	case CmdScriptStop:
		resp := c.handleStop(req)
		c.respond(msg, resp)
	case CmdFunctionRun:
		resp := c.handleRunFunction(req)
		c.respond(msg, resp)
	default:
		c.replyError(msg, "unknown action")
	}

}

func (c *Controller) replyError(msg *nats.Msg, errMsg string) {
	reply := msg.Reply
	if reply == "" {
		return
	}
	resp := RpcResponse{
		Status: "error",
		Error:  errMsg,
	}
	data, _ := json.Marshal(resp)
	msg.Respond(data)
}

func (c *Controller) handleStart(req RpcRequest) RpcResponse {
	c.m.ScriptRun(req.Script)
	// Implement start logic here
	return RpcResponse{Status: "started"}
}

func (c *Controller) handleStop(req RpcRequest) RpcResponse {
	c.m.ScriptStop(req.World)
	// Implement stop logic here
	return RpcResponse{Status: "stopped"}
}

func (c *Controller) handleRunFunction(req RpcRequest) RpcResponse {
	c.m.FunctionRun(req.Function, req.FunctionArgs)
	// Implement function run logic here
	return RpcResponse{Status: "function run"}
}

func (c *Controller) respond(msg *nats.Msg, resp RpcResponse) {
	if msg.Reply == "" {
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		c.replyError(msg, "failed to marshal response")
		return
	}
	msg.Respond(data)
}

type Client struct {
	nc *nats.Conn
}

func NewClient(nc *nats.Conn) *Client {
	return &Client{nc: nc}
}

func (c *Client) SendCommand(req RpcRequest) (RpcResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return RpcResponse{}, err
	}
	msg, err := c.nc.Request(ControllerSubject, data, nats.DefaultTimeout)
	if err != nil {
		return RpcResponse{}, err
	}
	var resp RpcResponse
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return RpcResponse{}, err
	}
	return resp, nil
}

func (c *Client) RunScript(fname string) (RpcResponse, error) {
	absName, error := filepath.Abs(fname)
	if error != nil {
		return RpcResponse{}, error
	}
	content, err := os.ReadFile(absName)
	if err != nil {
		return RpcResponse{}, err
	}

	script := NewScript(absName, string(content))
	req := RpcRequest{
		Action: CmdScriptStart,
		Script: script,
	}
	return c.SendCommand(req)
}
func (c *Client) StopScript(world string) (RpcResponse, error) {
	req := RpcRequest{
		Action: CmdScriptStop,
		World:  world,
	}
	return c.SendCommand(req)
}
func (c *Client) RunFunction(function string, args []any) (RpcResponse, error) {
	req := RpcRequest{
		Action:       CmdFunctionRun,
		Function:     function,
		FunctionArgs: args,
	}
	return c.SendCommand(req)
}

func WithClient(ctx context.Context, natsServer string, action func(ctx context.Context, client *Client) error) error {
	nc, err := nats.Connect(natsServer)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to nats server")
		return err
	}
	defer func() {
		nc.Drain()
		nc.Close()
	}()
	client := NewClient(nc)
	return action(ctx, client)
}
