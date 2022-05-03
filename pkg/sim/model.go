package sim

// Scenario is a scenario of a simulation.
type ScenarioDoc struct {
	Name        string           `json:"name" yaml:"name"`
	Version     string           `json:"version" yaml:"version"`
	Description string           `json:"description" yaml:"description"`
	Imports     []string         `json:"imports" yaml:"imports"`
	Interfaces  []InterfaceEntry `json:"interfaces" yaml:"interfaces"`
	Scripts     []ScriptEntry    `json:"scripts" yaml:"scripts"`
}

// InterfaceEntry is a single interface entry in a scenario.
type InterfaceEntry struct {
	Name        string        `json:"name" yaml:"name"`
	Description string        `json:"description" yaml:"description"`
	Properties  []ValueEntry  `json:"properties" yaml:"properties"`
	Methods     []ScriptEntry `json:"methods" yaml:"methods"`
}

// ValueEntry is a named value of an interface method or return entry.
// { "name": "count", "value": 0 }
type ValueEntry struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Value       any    `json:"value" yaml:"value"`
}

// ScriptEntry is a scripted entry
// Incoming method parameters are prefixed with `$params.<name>`
// The current properrties are prefixed with `$props.<name>`
// { "name": "count", "script": "$props.count + 1" }
type ScriptEntry struct {
	Name        string `json:"name" yaml:"name"`
	Source      string `json:"source" yaml:"source"`
	Description string `json:"description" yaml:"description"`
}
