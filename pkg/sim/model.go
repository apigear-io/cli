package sim

// Scenario is a scenario of a simulation.
type ScenarioDoc struct {
	Name        string         `json:"name" yaml:"name"`
	Version     string         `json:"version" yaml:"version"`
	Description string         `json:"description" yaml:"description"`
	Services    []ServiceEntry `json:"services" yaml:"services"`
}

// ServiceEntry is a single interface entry in a scenario.
type ServiceEntry struct {
	Name        string          `json:"name" yaml:"name"`
	Description string          `json:"description" yaml:"description"`
	Properties  []PropertyEntry `json:"properties" yaml:"properties"`
	Methods     []MethodEntry   `json:"methods" yaml:"methods"`
}

// PropertyEntry is a named value of an interface method or return entry.
// { "name": "count", "value": 0 }
type PropertyEntry struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Value       any    `json:"value" yaml:"value"`
}

type MethodEntry struct {
	Name        string        `json:"name" yaml:"name"`
	Description string        `json:"description" yaml:"description"`
	Actions     []ActionEntry `json:"actions" yaml:"actions"`
}

type ActionEntry struct {
	Name        string             `json:"name" yaml:"name"`
	Description string             `json:"description" yaml:"description"`
	Params      []ActionParamEntry `json:"params" yaml:"params"`
}

type ActionParamEntry struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Value       any    `json:"value" yaml:"value"`
}
