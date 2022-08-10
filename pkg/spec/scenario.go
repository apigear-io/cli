package spec

type ScenarioAction map[string][]string

type ScenarioOperation struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Actions     []ScenarioAction `json:"actions"`
}
type ScenarioInterface struct {
	Name       string         `json:"name" yaml:"name"`
	Properties map[string]any `json:"properties" yaml:"properties"`
}

type ScenarioDoc struct {
	Schema     string              `json:"schema" yaml:"schema"`
	Interfaces []ScenarioInterface `json:"interfaces" yaml:"interfaces"`
}
