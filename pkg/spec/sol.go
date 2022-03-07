package spec

type Layer struct {
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	Output      string   `json:"output" yaml:"output"`
	Inputs      []string `json:"inputs" yaml:"inputs"`
	Template    string   `json:"template" yaml:"template"`
	Features    []string `json:"features" yaml:"features"`
}

type SolutionDoc struct {
	Schema      string  `json:"schema" yaml:"schema"`
	Version     string  `json:"version" yaml:"version"`
	Name        string  `json:"name" yaml:"name"`
	Description string  `json:"description" yaml:"description"`
	RootDir     string  `json:"rootDir" yaml:"rootDir"`
	Layers      []Layer `json:"layers" yaml:"layers"`
}
