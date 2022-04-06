package spec

type ScenarioProperty struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

type ScenarioInterface struct {
	Name       string             `json:"name" yaml:"name"`
	Properties []ScenarioProperty `json:"properties" yaml:"properties"`
}

type ScenarioDoc struct {
	Schema     string              `json:"schema" yaml:"schema"`
	Interfaces []ScenarioInterface `json:"interfaces" yaml:"interfaces"`
}
