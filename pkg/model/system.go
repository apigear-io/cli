package model

type System struct {
	NamedNode `json:",inline" yaml:",inline"`
	Modules   []Module `json:"modules" yaml:"modules"`
}

func NewSystem(name string) *System {
	return &System{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindSystem,
		},
		Modules: make([]Module, 0),
	}
}

func (s System) ModuleByName(name string) Module {
	for _, m := range s.Modules {
		if m.Name == name {
			return m
		}
	}
	return Module{}
}
