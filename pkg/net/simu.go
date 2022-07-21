package net

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/net/rpc"
	"github.com/apigear-io/cli/pkg/sim"
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
		service := m.Params["service"].(string)
		operation := m.Params["operation"].(string)
		params := m.Params["params"].(map[string]any)
		err := s.simu.CallMethod(service, operation, params)
		if err != nil {
			return fmt.Errorf("simu.call: %s", err)
		}
	}
	return nil
}
