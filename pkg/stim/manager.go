package stim

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/stim/js"
	"github.com/apigear-io/cli/pkg/stim/model"
)

type Stimulation struct {
	rts map[string]*js.Runtime
}

func New() *Stimulation {
	return &Stimulation{
		rts: make(map[string]*js.Runtime),
	}
}

func (m *Stimulation) RunScript(id string, script model.Script) (any, error) {
	log.Info().Str("id", id).Str("script", script.Name).Msg("running script")
	rt, ok := m.rts[id]
	if ok { // reset the runtime
		rt.Interrupt()
		delete(m.rts, id)
	}
	// create a new runtime
	rt = js.NewRuntime(id)
	m.rts[id] = rt
	return rt.RunScript(script)
}
