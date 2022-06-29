package sim

import (
	"apigear/pkg/log"
	"fmt"
)

// Simulate runs one or more simulation scenarios.
// To unload a scenario call the Stop method with the scenario name.
// To stop all scenarios call the StopAll method.
// A scenario can simulation properties, methods and signals using static values or scripted behavior.
// The scripted behavior can be either triggered by a call or by a script entry
// All results are send out via the event channel.
type Simulation struct {
	Scenarios []*ScenarioDoc
}

func NewSimulation() *Simulation {
	return &Simulation{
		Scenarios: []*ScenarioDoc{},
	}
}

func (s *Simulation) AddScenario(scenario *ScenarioDoc) {
	if scenario != nil {
		fmt.Println("Simulation.start", scenario.Name)
		s.Scenarios = append(s.Scenarios, scenario)
	}
}

func (s *Simulation) RemoveScenario(name string) {
	fmt.Println("Simulation.stop", name)
}

func (s *Simulation) RemoveAll() {
	fmt.Println("Simulation.stopAll")
}

func (s Simulation) LookupMethod(ifaceName string, methodName string) *MethodEntry {
	log.Debugf("sim.lookupMethod: %s %s", ifaceName, methodName)
	iface := s.LookupInterface(ifaceName)
	if iface == nil {
		return nil
	}
	for _, m := range iface.Methods {
		if m.Name == methodName {
			return &m
		}
	}
	return nil
}

func (s Simulation) LookupInterface(ifaceName string) *InterfaceEntry {
	log.Debugf("sim.lookupInterface: %s", ifaceName)
	for _, scenario := range s.Scenarios {
		for _, iface := range scenario.Interfaces {
			if iface.Name == ifaceName {
				return &iface
			}
		}
	}
	return nil
}

func (s Simulation) CallMethod(service string, method string, params map[string]any) error {
	log.Debugf("sim.call: %s#%s %v", service, method, params)
	iface := s.LookupInterface(service)
	if iface == nil {
		return fmt.Errorf("interface %s not found", service)
	}
	for _, m := range iface.Methods {
		if m.Name == method {
			log.Debugf("TODO: call method %s.%s", service, method)
		}
	}
	return nil
}
