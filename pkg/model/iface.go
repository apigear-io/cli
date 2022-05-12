package model

type InputsProvider interface {
	GetInputs() []*TypedNode
}

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

func (s *Signal) GetInputs() []*TypedNode {
	return s.Inputs
}

func (s *Signal) ResolveAll(m *Module) error {
	for _, i := range s.Inputs {
		err := i.ResolveAll(m)
		if err != nil {
			return err
		}
	}
	return nil
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
		Inputs: make([]*TypedNode, 0),
		Output: NewTypeNode("", KindOutput),
	}
}

func (m *Method) GetInputs() []*TypedNode {
	return m.Inputs
}

func (m *Method) GetName() string {
	return m.Name
}
func (m *Method) GetKind() Kind {
	return KindMethod
}
func (m *Method) GetSchema() *Schema {
	return &m.Output.Schema
}

func (m *Method) ResolveAll(mod *Module) error {
	if m.Output == nil {
		m.Output = NewTypeNode("", KindOutput)
	}
	if m.Inputs == nil {
		m.Inputs = make([]*TypedNode, 0)
	}
	for _, p := range m.Inputs {
		err := p.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	if m.Output != nil {
		err := m.Output.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	return nil
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

func (i *Interface) ResolveAll(mod *Module) error {
	for _, p := range i.Properties {
		err := p.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	for _, meth := range i.Methods {
		err := meth.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	for _, s := range i.Signals {
		err := s.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	return nil
}
