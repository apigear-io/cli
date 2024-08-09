package model

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec/rkw"
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

func (s *Signal) Validate(m *Module) error {
	if s.Params == nil {
		s.Params = make([]*TypedNode, 0)
	}
	for _, i := range s.Params {
		err := i.Validate(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Signal) CheckReservedWords(langs []rkw.Lang) {
	s.Name = rkw.CheckAndEscapeName(s.Name, "signal")
	for _, p := range s.Params {
		p.CheckReservedWords(langs)
	}
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

func (m *Operation) Validate(mod *Module) error {
	if m.Return == nil {
		m.Return = NewTypedNode("", KindReturn)
	}
	if m.Params == nil {
		m.Params = make([]*TypedNode, 0)
	}
	for _, p := range m.Params {
		err := p.Validate(mod)
		if err != nil {
			return err
		}
	}
	if m.Return != nil {
		err := m.Return.Validate(mod)
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

func (m *Operation) CheckReservedWords(langs []rkw.Lang) {
	m.Name = rkw.CheckAndEscapeName(m.Name, "operation")
	for _, p := range m.Params {
		p.CheckReservedWords(langs)
	}
	if m.Return != nil {
		m.Return.CheckReservedWords(langs)
	}
}

type Extends struct {
	Name   string `json:"name" yaml:"name"`
	Import string `json:"import" yaml:"import"`
}

type Interface struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Module     *Module      `json:"-" yaml:"-"`
	Extends    Extends      `json:"extends" yaml:"extends"`
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
		Extends:    Extends{},
		Properties: make([]*TypedNode, 0),
		Operations: make([]*Operation, 0),
		Signals:    make([]*Signal, 0),
	}
}

func (i *Interface) HasExtends() bool {
	return i.Extends.Name != ""
}

func (i Interface) LookupMember(name string) *NamedNode {
	for _, p := range i.Properties {
		if p.Name == name {
			return &p.NamedNode
		}
	}
	for _, m := range i.Operations {
		if m.Name == name {
			return &m.NamedNode
		}
	}
	for _, s := range i.Signals {
		if s.Name == name {
			return &s.NamedNode
		}
	}
	return nil
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

func (i *Interface) Validate(mod *Module) error {
	// check if any names are duplicated
	names := make(map[string]bool)
	if i.HasExtends() {
		if i.Extends.Name == i.Name && i.Extends.Import == "" {
			log.Warn().Msgf("%s: interface extends itself", i.Name)
		}
		if i.Module.LookupInterface(i.Extends.Import, i.Extends.Name) == nil {
			log.Warn().Msgf("%s: interface extends unknown interface: %s", i.Name, i.Extends.Name)
		}
	}
	for _, p := range i.Properties {
		err := p.Validate(mod)
		if err != nil {
			return err
		}
		if names[p.Name] {
			return fmt.Errorf("%s: duplicate name: %s", i.Name, p.Name)
		}
		names[p.Name] = true
	}
	for _, op := range i.Operations {
		if err := op.Validate(mod); err != nil {
			return err
		}
		if names[op.Name] {
			return fmt.Errorf("%s: duplicate name: %s", i.Name, op.Name)
		}
		names[op.Name] = true
	}
	for _, s := range i.Signals {
		if err := s.Validate(mod); err != nil {
			return err
		}
		if names[s.Name] {
			return fmt.Errorf("%s: duplicate name: %s", i.Name, s.Name)
		}
		names[s.Name] = true
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

func (i *Interface) CheckReservedWords(langs []rkw.Lang) {
	i.Name = rkw.CheckAndEscapeName(i.Name, "interface")
	for _, p := range i.Properties {
		p.CheckReservedWords(langs)
	}
	for _, o := range i.Operations {
		o.CheckReservedWords(langs)
	}
	for _, s := range i.Signals {
		s.CheckReservedWords(langs)
	}
}
