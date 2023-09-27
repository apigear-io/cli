package spec

import "fmt"

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
	err := s.Compute()
	if err != nil {
		return err
	}
	if s.Layers != nil {
		return fmt.Errorf("sol-doc %s: layers are deprecated, use targets instead", s.Name)
	}
	if s.Targets == nil {
		s.Targets = make([]*SolutionTarget, 0)
	}
	for _, l := range s.Targets {
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
	if s.computed {
		return nil
	}
	s.computed = true
	for _, c := range compute {
		err := c(s)
		if err != nil {
			return err
		}
	}
	if len(s.Layers) > 0 {
		log.Warn().Msg("layers inside solutions are deprecated, use targets instead")
		s.Targets = append(s.Targets, s.Layers...)
		s.Layers = nil
	}
	for _, t := range s.Targets {
		err := t.Compute(s)
		if err != nil {
			return err
		}
	}
	return s.Validate()
}

// AggregateDependencies computes the dependencies of each layer.
func (s *SolutionDoc) AggregateDependencies() []string {
	deps := make([]string, 0)
	for _, l := range s.Targets {
		deps = append(deps, l.Dependencies()...)
	}
	return deps
}
