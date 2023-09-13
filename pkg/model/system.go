package model

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec/rkw"
)

type System struct {
	NamedNode `json:",inline" yaml:",inline"`
	Modules   []*Module `json:"modules" yaml:"modules"`
}

func NewSystem(name string) *System {
	return &System{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindSystem,
		},
		Modules: make([]*Module, 0),
	}
}

// LookupModule looks up a module by name
func (s System) LookupModule(name string) *Module {
	for _, m := range s.Modules {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (s System) LookupInterface(moduleName string, ifaceName string) *Interface {
	m := s.LookupModule(moduleName)
	if m == nil {
		return nil
	}
	return m.LookupInterface(ifaceName)
}

func (s System) LookupStruct(moduleName string, structName string) *Struct {
	m := s.LookupModule(moduleName)
	if m == nil {
		return nil
	}
	return m.LookupStruct(structName)
}

func (s System) LookupEnum(moduleName string, enumName string) *Enum {
	m := s.LookupModule(moduleName)
	if m == nil {
		return nil
	}
	return m.LookupEnum(enumName)
}

func (s System) LookupProperty(moduleName string, ifaceName string, propName string) *TypedNode {
	i := s.LookupInterface(moduleName, ifaceName)
	if i == nil {
		return nil
	}
	return i.LookupProperty(propName)
}

func (s System) LookupOperation(moduleName string, ifaceName string, operationName string) *Operation {
	i := s.LookupInterface(moduleName, ifaceName)
	if i == nil {
		return nil
	}
	return i.LookupOperation(operationName)
}

func (s System) LookupSignal(moduleName string, ifaceName string, eventName string) *Signal {
	i := s.LookupInterface(moduleName, ifaceName)
	if i == nil {
		return nil
	}
	return i.LookupSignal(eventName)
}

func (s *System) Validate() error {
	rkw.CheckName(s.Name, "system")
	for _, m := range s.Modules {
		err := m.Validate()
		if err != nil {
			return err
		}
	}
	// check if there are duplicate module names
	names := make(map[string]bool)
	for _, m := range s.Modules {
		if names[m.Name] {
			return fmt.Errorf("%s: duplicate name: %s", s.Name, m.Name)
		}
		names[m.Name] = true
	}
	s.ComputeIdentifier()
	return nil
}

func (s *System) ComputeIdentifier() {
	var idx uint = 1
	for _, m := range s.Modules {
		m.Id = idx
		idx++
		for _, i := range m.Interfaces {
			i.Id = idx
			idx++
			for _, o := range i.Operations {
				o.Id = idx
				idx++
			}
			for _, p := range i.Properties {
				p.Id = idx
				idx++
			}
			for _, s := range i.Signals {
				s.Id = idx
				idx++
			}
		}
		for _, s := range m.Structs {
			s.Id = idx
			idx++
			for _, p := range s.Fields {
				p.Id = idx
				idx++
			}
		}
		for _, e := range m.Enums {
			e.Id = idx
			idx++
			for _, p := range e.Members {
				p.Id = idx
				idx++
			}
		}
	}
}
