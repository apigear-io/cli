package net

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/wsrpc/rpc"
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

func (s SimuRpcHandler) HandleMessage(r *rpc.Request) error {
	var m rpc.Message
	err := r.AsJSON(&m)
	if err != nil {
		return err
	}
	symbol, ok := m.Params["symbol"].(string)
	if !ok {
		return fmt.Errorf("invalid symbol: %v", m.Params["symbol"])
	}
	log.Debugf("-> %s %s", m.Method, m.Params["symbol"])
	switch m.Method {
	case "simu.call":
		iface, op := core.SplitSymbol(symbol)
		data, ok := m.Params["data"].(map[string]any)
		if !ok {
			return fmt.Errorf("simu.call: no or no valid data")
		}
		log.Infof("invoke[%d] %s/%s", m.Id, iface, op)
		result, err := s.simu.InvokeOperation(iface, op, data)
		if err != nil {
			return fmt.Errorf("invoke operation: %s", err)
		}
		if result != nil {
			log.Infof("reply[%d]  %s/%s: %v", m.Id, iface, op, result)
			err := r.ReplyJSON(rpc.MakeResult(m.Id, result))
			if err != nil {
				log.Errorf("failed to reply: %v", err)
			}
		}
	case "simu.state":
		data, ok := m.Params["data"].(map[string]any)
		if !ok {
			data = map[string]any{}
		}
		if len(data) == 0 {
			// simu.state without data returns the properties of the interface
			log.Infof("get %s", symbol)
			props, err := s.simu.GetProperties(symbol)
			if err != nil {
				return fmt.Errorf("get properties: %s", err)
			}
			// send the properties back to the client
			err = r.ReplyJSON(rpc.MakeNotify("simu.state", map[string]any{"data": props, "symbol": symbol}))
			if err != nil {
				log.Errorf("failed to reply: %v", err)
			}
		} else {
			// simu.state with data sets the properties of the interface
			log.Infof("set %s", symbol)
			err := s.simu.SetProperties(symbol, data)
			if err != nil {
				return fmt.Errorf("set properties: %s", err)
			}
		}
	}
	return nil
}
