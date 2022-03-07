package spec

type Property struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

type Interface struct {
	Name       string     `json:"name" yaml:"name"`
	Properties []Property `json:"properties" yaml:"properties"`
}

type ScenarioDoc struct {
	Schema     string      `json:"schema" yaml:"schema"`
	Interfaces []Interface `json:"interfaces" yaml:"interfaces"`
}
