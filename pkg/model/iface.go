package model

import (
	"fmt"
	"sort"
)

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
		Params: make([]*TypedNode, 0),
	}
}

func (s *Signal) ResolveAll(m *Module) error {
	if s.Params == nil {
		s.Params = make([]*TypedNode, 0)
	}
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

func (m *Operation) ParamNames() []string {
	names := make([]string, 0)
	for _, p := range m.Params {
		names = append(names, p.Name)
	}
	return names
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
		Properties: make([]*TypedNode, 0),
		Operations: make([]*Operation, 0),
		Signals:    make([]*Signal, 0),
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
	// check if any names are duplicated
	names := make(map[string]bool)
	for _, p := range i.Properties {
		if names[p.Name] {
			return fmt.Errorf("%s: duplicate name: %s", i.Name, p.Name)
		}
		names[p.Name] = true
		err := p.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	for _, op := range i.Operations {
		if names[op.Name] {
			return fmt.Errorf("%s: duplicate name: %s", i.Name, op.Name)
		}
		names[op.Name] = true
		err := op.ResolveAll(mod)
		if err != nil {
			return err
		}
	}
	for _, s := range i.Signals {
		if names[s.Name] {
			return fmt.Errorf("%s: duplicate name: %s", i.Name, s.Name)
		}
		names[s.Name] = true
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

func (i Interface) SortedProperties() []*TypedNode {
	out := make([]*TypedNode, len(i.Properties))
	copy(out, i.Properties)
	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return out
}

func (i Interface) SortedOperations() []*Operation {
	out := make([]*Operation, len(i.Operations))
	copy(out, i.Operations)
	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return out
}

func (i Interface) SortedSignals() []*Signal {
	out := make([]*Signal, len(i.Signals))
	copy(out, i.Signals)
	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return out
}
