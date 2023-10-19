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
	isValid     bool              `json:"-" yaml:"-"` // is valid
	computed    bool              `json:"-" yaml:"-"` // is computed
}

func (s *SolutionDoc) Accept(visitor SolutionVisitor) error {
	err := visitor.VisitSolutionDoc(s)
	if err != nil {
		return err
	}
	for _, l := range s.Targets {
		err := l.Accept(s, visitor)
		if err != nil {
			return err
		}
	}
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

func (s *SolutionDoc) GetMeta(key string) any {
	if s.Meta == nil {
		return nil
	}
	return s.Meta[key]
}

func (s *SolutionDoc) SetMeta(key string, value any) {
	if s.Meta == nil {
		s.Meta = make(map[string]any)
	}
	s.Meta[key] = value
}

// AddTarget adds a target to the solution.
func (s *SolutionDoc) AddTarget(name string) *SolutionTarget {
	t := &SolutionTarget{
		Name: name,
	}
	s.Targets = append(s.Targets, t)
	return t
}
