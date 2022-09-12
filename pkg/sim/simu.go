package sim

import (
	"github.com/apigear-io/cli/pkg/sim/actions"
	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/script"
	"github.com/apigear-io/cli/pkg/spec"
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

func (s *Simulation) Stop() {
	s.aEng.StopAllSequences()
	s.sEng.StopAllSequences()
}

func (s *Simulation) LoadScenario(source string, doc *spec.ScenarioDoc) error {
	log.Infof("simulation load scenario: %s", source)
	return s.aEng.LoadScenario(source, doc)
}

func (s *Simulation) UnloadScenario(source string) error {
	log.Infof("simulation unload scenario: %s", source)
	return s.aEng.UnloadScenario(source)
}

func (s *Simulation) LoadScript(source string, script string) error {
	_, err := s.sEng.LoadScript(source, script)
	if err != nil {
		log.Errorf("sim.loadScript: %s %s", source, err)
	}
	return err
}
