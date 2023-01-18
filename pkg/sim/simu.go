package sim

import (
	"github.com/apigear-io/cli/pkg/sim/actions"
	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/ostore"
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
	*core.MultiEngine
	store ostore.IObjectStore
	aEng  *actions.Engine
	sEng  *script.Engine
}

func NewSimulation() *Simulation {
	aStore := ostore.NewMemoryStore()
	aEng := actions.NewEngine(aStore)
	sEng := script.NewEngine(aStore)
	s := &Simulation{
		MultiEngine: core.NewMultiEngine(aEng, sEng),
		store:       aStore,
		aEng:        aEng,
		sEng:        sEng,
	}
	s.init()
	return s
}

func (s *Simulation) init() {
	s.store.OnEvent(func(evt ostore.StoreEvent) {
		switch evt.Type {
		case ostore.EventTypeCreate, ostore.EventTypeUpdate:
			for prop, val := range evt.Value {
				s.EmitPropertyChanged(evt.Id, prop, val)
			}
		}
	})
}

func (s *Simulation) Stop() {
	s.aEng.StopAllSequences()
	s.sEng.StopAllSequences()
}

func (s *Simulation) LoadScenario(source string, doc *spec.ScenarioDoc) error {
	log.Debug().Msgf("simulation load scenario: %s", source)
	return s.aEng.LoadScenario(source, doc)
}

func (s *Simulation) UnloadScenario(source string) error {
	log.Debug().Msgf("simulation unload scenario: %s", source)
	return s.aEng.UnloadScenario(source)
}

func (s *Simulation) LoadScript(source string, script string) error {
	_, err := s.sEng.LoadScript(source, script)
	if err != nil {
		log.Error().Msgf("sim.loadScript: %s %s", source, err)
	}
	return err
}
