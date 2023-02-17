package spec

type SolutionDoc struct {
	Schema      string           `json:"schema" yaml:"schema"`
	Version     string           `json:"version" yaml:"version"`
	Name        string           `json:"name" yaml:"name"`
	Description string           `json:"description" yaml:"description"`
	RootDir     string           `json:"rootDir" yaml:"rootDir"`
	Layers      []*SolutionLayer `json:"layers" yaml:"layers"`
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

// ComputeDependencies computes the dependencies of a solution.
// A solution is a list of layers. Each layer has a list of inputs.
// The dependencies of a solution are:
// - the inputs of all layers
// - the solution file
// - the template dir and rules file
// The inputs are relative paths to the root directory of the solution.
func (s *SolutionDoc) ComputeDependencies() []string {
	deps := make([]string, 0)
	for _, l := range s.Layers {
		l.ComputeDependencies(s.RootDir)
		deps = append(deps, l.dependencies...)
	}
	return deps
}

// Compute computes the solution.
// It computes the dependencies and expanded inputs of each layer.
func (s *SolutionDoc) Compute() error {
	for _, l := range s.Layers {
		err := l.Compute(s.RootDir)
		if err != nil {
			return err
		}
	}
	return nil
}
