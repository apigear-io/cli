package spec

type SolutionDoc struct {
	Schema      string            `json:"schema" yaml:"schema"`
	Version     string            `json:"version" yaml:"version"`
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	RootDir     string            `json:"rootDir" yaml:"rootDir"`
	Meta        map[string]any    `json:"meta" yaml:"meta"`
	Layers      []*SolutionTarget `json:"layers" yaml:"layers"`
	Targets     []*SolutionTarget `json:"targets" yaml:"targets"`
	// computed fields
	computed bool `json:"-" yaml:"-"`
}

func (s *SolutionDoc) Validate() error {
	// basic validation
	if s.Targets == nil {
		s.Targets = make([]*SolutionTarget, 0)
	}
	if err := s.compute(); err != nil {
		return err
	}
	for _, t := range s.Targets {
		err := t.Validate(s)
		if err != nil {
			return err
		}
	}
	return nil
}

// compute computes derived fields.
func (s *SolutionDoc) compute() error {
	if s.computed {
		return nil
	}
	if len(s.Layers) > 0 {
		log.Warn().Msg("layers inside solutions are deprecated, use targets instead")
		s.Targets = append(s.Targets, s.Layers...)
		s.Layers = nil
	}
	s.computed = true
	return nil
}

// AggregateDependencies computes the dependencies of each layer.
func (s *SolutionDoc) AggregateDependencies() []string {
	deps := make([]string, 0)
	for _, l := range s.Targets {
		deps = append(deps, l.Dependencies()...)
	}
	return deps
}
