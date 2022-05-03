package net

import (
	"fmt"
	"objectapi/pkg/log"
	"objectapi/pkg/net/rpc"
	"objectapi/pkg/sim"
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
		if len(m.Params) != 2 {
			return fmt.Errorf("invalid params for simu.call: %v", m.Params)
		}
		symbol := m.Params[0].(string)
		args := m.Params[1].([]interface{})
		s.simu.CallMethod(symbol, args)
	}
	return nil
}
