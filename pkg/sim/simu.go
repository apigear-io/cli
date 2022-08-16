package sim

import (
	"github.com/apigear-io/cli/pkg/sim/actions"
	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/script"
)

// Simulate runs one or more simulation scenarios.
// To unload a scenario call the Stop operation with the scenario name.
// To stop all scenarios call the StopAll operation.
// A scenario can simulation properties, operations and signals using static values or scripted behavior.
// The scripted behavior can be either triggered by a call or by a script entry
// All results are send out via the event channel.
type Simulation struct {
	aEng *actions.Engine
	sEng *script.Engine
	*core.MultiEngine
}

func NewSimulation() *Simulation {
	aEng := actions.NewEngine()
	sEng := script.NewEngine()
	s := &Simulation{
		MultiEngine: core.NewMultiEngine(aEng, sEng),
		aEng:        aEng,
		sEng:        sEng,
	}
	return s
}

func (s *Simulation) LoadScenario(doc *actions.ScenarioDoc) error {
	s.aEng.LoadScenario(doc)
	return nil
}

func (s *Simulation) LoadScript(name string, script string) error {
	_, err := s.sEng.LoadScript(name, script)
	if err != nil {
		log.Errorf("sim.loadScript: %s %s", name, err)
	}
	return err
}
