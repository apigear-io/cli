package spec

type SolutionDoc struct {
	Schema      string           `json:"schema" yaml:"schema"`
	Version     string           `json:"version" yaml:"version"`
	Name        string           `json:"name" yaml:"name"`
	Description string           `json:"description" yaml:"description"`
	RootDir     string           `json:"rootDir" yaml:"rootDir"`
	Meta        map[string]any   `json:"meta" yaml:"meta"`
	Layers      []*SolutionLayer `json:"layers" yaml:"layers"`
	// computed fields
	computed bool `json:"-" yaml:"-"`
}

func (s *SolutionDoc) Validate() error {
	if s.Layers == nil {
		s.Layers = make([]*SolutionLayer, 0)
	}
	for _, l := range s.Layers {
		err := l.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// Compute computes the solution.
// It computes the dependencies and expanded inputs of each layer.
func (s *SolutionDoc) Compute(compute ...func(doc *SolutionDoc) error) error {
	s.computed = true
	for _, c := range compute {
		err := c(s)
		if err != nil {
			return err
		}
	}
	for _, l := range s.Layers {
		err := l.Compute(s)
		if err != nil {
			return err
		}
	}
	return s.Validate()
}

// AggregateDependencies computes the dependencies of each layer.
func (s *SolutionDoc) AggregateDependencies() []string {
	deps := make([]string, 0)
	for _, l := range s.Layers {
		deps = append(deps, l.Dependencies()...)
	}
	return deps
}
