package model

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/spec/rkw"
)

type System struct {
	NamedNode `json:",inline" yaml:",inline"`
	Modules   []*Module `json:"modules" yaml:"modules"`
	Checksum  string    `json:"checksum" yaml:"checksum"`
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
	s.computeIdentifier()
	err := s.computeChecksum()
	if err != nil {
		return err
	}
	log.Info().Msgf("system %s resolved: %s", s.Name, s.Checksum)
	return nil
}

// TODO: clean up this code
func (s *System) computeIdentifier() {
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

// TODO: clean up this code
func (s *System) computeChecksum() error {
	var buffer bytes.Buffer
	for _, m := range s.Modules {
		if len(m.Checksum) == 0 {
			return fmt.Errorf("module %s checksum not computed", m.Name)
		}
		buffer.WriteString(m.Checksum)
	}
	sum := md5.Sum(buffer.Bytes())
	s.Checksum = hex.EncodeToString(sum[:])
	return nil
}
