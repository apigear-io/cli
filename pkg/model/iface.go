package model

type Signal struct {
	NamedNode `json:",inline" yaml:",inline"`
	Params    []*TypedNode `json:"params" yaml:"params"`
}

func NewSignal(name string) *Signal {
	return &Signal{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindSignal,
		},
	}
}

func (s *Signal) GetParams() []*TypedNode {
	return s.Params
}

func (s *Signal) ResolveAll(m *Module) error {
	for _, i := range s.Params {
		err := i.ResolveAll(m)
		if err != nil {
			return err
		}
	}
	return nil
}

type Operation struct {
	NamedNode `json:",inline" yaml:",inline"`
	Params    []*TypedNode `json:"params" yaml:"params"`
	Return    *TypedNode   `json:"return" yaml:"return"`
}

func NewOperation(name string) *Operation {
	return &Operation{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindOperation,
		},
		Params: make([]*TypedNode, 0),
		Return: NewTypedNode("", KindReturn),
	}
}

func (m *Operation) GetParams() []*TypedNode {
	return m.Params
}

func (m *Operation) GetName() string {
	return m.Name
}
func (m *Operation) GetKind() Kind {
	return KindOperation
}
func (m *Operation) GetSchema() *Schema {
	return &m.Return.Schema
}

func (m *Operation) ResolveAll(mod *Module) error {
	if m.Return == nil {
		m.Return = NewTypedNode("", KindReturn)
	}
	if m.Params == nil {
		m.Params = make([]*TypedNode, 0)
	}
	for _, p := range m.Params {
		err := p.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	if m.Return != nil {
		err := m.Return.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	return nil
}

type Interface struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Properties []*TypedNode `json:"properties" yaml:"properties"`
	Operations []*Operation `json:"operations" yaml:"operations"`
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

func (i Interface) LookupOperation(name string) *Operation {
	for _, m := range i.Operations {
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
	for _, op := range i.Operations {
		err := op.ResolveAll(mod)
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

func (i Interface) NoProperties() bool {
	return len(i.Properties) == 0
}

func (i Interface) NoOperations() bool {
	return len(i.Operations) == 0
}

func (i Interface) NoSignals() bool {
	return len(i.Signals) == 0
}

func (i Interface) NoMembers() bool {
	return i.NoProperties() && i.NoOperations() && i.NoSignals()
}
