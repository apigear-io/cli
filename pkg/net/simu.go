package net

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
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

func (s SimuRpcHandler) HandleMessage(r *rpc.Request) error {
	var m rpc.Message
	err := r.AsJSON(&m)
	if err != nil {
		return err
	}
	log.Info().Msgf("handle rpc message: %+v", m)

	symbol, ok := m.Params["symbol"].(string)
	if !ok {
		return fmt.Errorf("no symbol in simu call")
	}

	log.Debug().Msgf("-> %s %+v", m.Method, m.Params)
	switch m.Method {
	case "simu.call":
		iface, op := core.SplitSymbol(symbol)
		data, ok := m.Params["data"].(map[string]any)
		if !ok {
			return fmt.Errorf("simu.call: no or no valid data")
		}
		log.Debug().Msgf("invoke[%d] %s/%s", m.Id, iface, op)
		result, err := s.simu.InvokeOperation(iface, op, data)
		if err != nil {
			return fmt.Errorf("invoke operation: %s", err)
		}
		if result != nil {
			log.Debug().Msgf("reply[%d]  %s/%s: %v", m.Id, iface, op, result)
			err := r.ReplyJSON(rpc.MakeResult(m.Id, result))
			if err != nil {
				log.Error().Msgf("reply: %v", err)
			}
		}
	case "simu.state":
		data, ok := m.Params["data"].(map[string]any)
		if !ok {
			data = map[string]any{}
		}
		if len(data) == 0 {
			// simu.state without data returns the properties of the interface
			log.Debug().Msgf("get %s", symbol)
			props, err := s.simu.GetProperties(symbol)
			if err != nil {
				return fmt.Errorf("get properties: %s", err)
			}
			// send the properties back to the client
			err = r.ReplyJSON(rpc.MakeNotify("simu.state", map[string]any{"data": props, "symbol": symbol}))
			if err != nil {
				log.Error().Msgf("reply: %v", err)
			}
		} else {
			// simu.state with data sets the properties of the interface
			log.Debug().Msgf("set %s", symbol)
			err := s.simu.SetProperties(symbol, data)
			if err != nil {
				return fmt.Errorf("set properties: %s", err)
			}
		}
	}
	return nil
}
