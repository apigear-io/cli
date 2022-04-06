package spec

type SolutionLayer struct {
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	Inputs      []string `json:"inputs" yaml:"inputs"`
	Output      string   `json:"output" yaml:"output"`
	Template    string   `json:"template" yaml:"template"`
	Features    []string `json:"features" yaml:"features"`
	Force       bool     `json:"force" yaml:"force"`
}

type SolutionDoc struct {
	Schema      string          `json:"schema" yaml:"schema"`
	Version     string          `json:"version" yaml:"version"`
	Name        string          `json:"name" yaml:"name"`
	Description string          `json:"description" yaml:"description"`
	RootDir     string          `json:"rootDir" yaml:"rootDir"`
	Layers      []SolutionLayer `json:"layers" yaml:"layers"`
}
