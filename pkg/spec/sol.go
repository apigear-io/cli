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

func (l *SolutionLayer) Resolve() error {
	if l.Inputs == nil {
		l.Inputs = make([]string, 0)
	}
	if l.Features == nil {
		l.Features = make([]string, 0)
	}
	return nil
}

type SolutionDoc struct {
	Schema      string          `json:"schema" yaml:"schema"`
	Version     string          `json:"version" yaml:"version"`
	Name        string          `json:"name" yaml:"name"`
	Description string          `json:"description" yaml:"description"`
	RootDir     string          `json:"rootDir" yaml:"rootDir"`
	Layers      []SolutionLayer `json:"layers" yaml:"layers"`
}

func (s *SolutionDoc) Resolve() error {
	if s.Layers == nil {
		s.Layers = make([]SolutionLayer, 0)
	}
	for _, l := range s.Layers {
		l.Resolve()
	}
	return nil
}
