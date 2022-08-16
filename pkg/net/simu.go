package net

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/net/rpc"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/core"
)

type SimuRpcHandler struct {
	simu *sim.Simulation
}

func NewSimuRpcHandler(simu *sim.Simulation) *SimuRpcHandler {
	r := SimuRpcHandler{
		simu: simu,
	}
	return &r
}

func (s SimuRpcHandler) HandleMessage(r rpc.RpcRequest) error {
	log.Debugf("simu rpc handler: %v", r.Msg)
	m := r.Msg

	switch m.Method {
	case "simu.call":
		var symbol string
		if m.Params == nil {
			return fmt.Errorf("simu.call: no params")
		}
		symbol, ok := m.Params["symbol"].(string)
		if !ok {
			return fmt.Errorf("simu.call: no or no valid symbol")
		}
		iface, op := core.SplitSymbol(symbol)
		data, ok := m.Params["data"].(map[string]any)
		if !ok {
			return fmt.Errorf("simu.call: no or no valid data")
		}
		_, err := s.simu.InvokeOperation(iface, op, data)
		if err != nil {
			return fmt.Errorf("invoke operation: %s", err)
		}
	case "simu.state":
		symbol, ok := m.Params["symbol"].(string)
		if !ok {
			return fmt.Errorf("simu.state: no or no valid symbol: %v", m.Params)
		}
		data, ok := m.Params["data"].(map[string]any)
		if !ok {
			return fmt.Errorf("simu.state: no or no valid data: %v", m.Params)
		}
		if data != nil {
			// simu.state with data sets the properties of the interface
			err := s.simu.SetProperties(symbol, data)
			if err != nil {
				return fmt.Errorf("set properties: %s", err)
			}
		} else {
			// simu.state without data returns the properties of the interface
			props, err := s.simu.GetProperties(symbol)
			if err != nil {
				return fmt.Errorf("get properties: %s", err)
			}
			// send the properties back to the client
			r.Reply(rpc.MakeNotify("simu.state", rpc.RpcParams{"data": props}))
		}
	}
	return nil
}
