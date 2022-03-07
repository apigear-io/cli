package spec

type Document struct {
	Source string `json:"source" yaml:"source"`
	Target string `json:"target" yaml:"target"`
	Force  bool   `json:"force" yaml:"force"`
}

type Scope struct {
	Name      string     `json:"name" yaml:"name"`
	Match     string     `json:"match" yaml:"match"`
	Documents []Document `json:"documents" yaml:"documents"`
}

type Feature struct {
	Name   string  `json:"name" yaml:"name"`
	Scopes []Scope `json:"scopes" yaml:"scopes"`
}

type RulesDoc struct {
	Schema   string    `json:"schema" yaml:"schema"`
	Features []Feature `json:"features" yaml:"features"`
}
