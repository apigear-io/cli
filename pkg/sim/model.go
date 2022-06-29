package sim

// Scenario is a scenario of a simulation.
type ScenarioDoc struct {
	Name        string                    `json:"name" yaml:"name"`
	Version     string                    `json:"version" yaml:"version"`
	Description string                    `json:"description" yaml:"description"`
	Interfaces  map[string]InterfaceEntry `json:"interfaces" yaml:"interfaces"`
}

// InterfaceEntry is a single interface entry in a scenario.
type InterfaceEntry struct {
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
	Properties  []PropertyEntry        `json:"properties" yaml:"properties"`
	Methods     map[string]MethodEntry `json:"methods" yaml:"methods"`
}

// PropertyEntry is a named value of an interface method or return entry.
// { "name": "count", "value": 0 }
type PropertyEntry struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Value       any    `json:"value" yaml:"value"`
}

type MethodEntry struct {
	Name    string   `json:"name" yaml:"name"`
	Actions []string `json:"actions" yaml:"actions"`
}
