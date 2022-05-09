package model

type Signal struct {
	NamedNode `json:",inline" yaml:",inline"`
	Inputs    []*TypedNode `json:"inputs" yaml:"inputs"`
}

func NewSignal(name string) *Signal {
	return &Signal{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindSignal,
		},
	}
}

func (s *Signal) ResolveAll() {
	for _, i := range s.Inputs {
		i.ResolveAll()
	}
}

type Method struct {
	NamedNode `json:",inline" yaml:",inline"`
	// maybe inputs and outputs should be a map of name to Parameter
	Inputs []*TypedNode `json:"inputs" yaml:"inputs"`
	Output *TypedNode   `json:"output" yaml:"output"`
}

func NewMethod(name string) *Method {
	return &Method{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindMethod,
		},
	}
}

func (m *Method) ResolveAll() {
	for _, p := range m.Inputs {
		p.ResolveAll()
	}
	if m.Output != nil {
		m.Output.ResolveAll()
	}
}

type Interface struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Properties []*TypedNode `json:"properties" yaml:"properties"`
	Methods    []*Method    `json:"methods" yaml:"methods"`
	Signals    []*Signal    `json:"signals" yaml:"signals"`
}

func NewInterface(name string) *Interface {
	return &Interface{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindInterface,
		},
	}
}

func (i Interface) LookupMethod(name string) *Method {
	for _, m := range i.Methods {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (i Interface) LookupProperty(name string) *TypedNode {
	for _, p := range i.Properties {
		if p.Name == name {
			return p
		}
	}
	return nil
}

func (i Interface) LookupSignal(name string) *Signal {
	for _, s := range i.Signals {
		if s.Name == name {
			return s
		}
	}
	return nil
}

func (i *Interface) ResolveAll() {
	for _, p := range i.Properties {
		p.ResolveAll()
	}
	for _, m := range i.Methods {
		m.ResolveAll()
	}
	for _, s := range i.Signals {
		s.ResolveAll()
	}
}
