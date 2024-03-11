package model

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/spec/rkw"
)

type System struct {
	NamedNode `json:",inline" yaml:",inline"`
	Modules   []*Module `json:"modules" yaml:"modules"`
	Checksum  string    `json:"checksum" yaml:"checksum"`
}

// NewSystem creates a new system
func NewSystem(name string) *System {
	return &System{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindSystem,
		},
		Modules: make([]*Module, 0),
	}
}

func (s *System) AddModule(m *Module) {
	s.Modules = append(s.Modules, m)
	m.System = s
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

// LookupInterface looks up an interface by module and interface name
func (s System) LookupInterface(mName string, iName string) *Interface {
	m := s.LookupModule(mName)
	if m == nil {
		return nil
	}
	return m.LookupInterface(mName, iName)
}

// LookupNode looks up a node by module and node name
func (s System) LookupNode(mName string, nName string) *NamedNode {
	m := s.LookupModule(mName)
	if m == nil {
		return nil
	}
	return m.LookupNode(mName, nName)
}

// LookupStruct looks up a struct by module and struct name
func (s System) LookupStruct(mName, sName string) *Struct {
	m := s.LookupModule(mName)
	if m == nil {
		return nil
	}
	return m.LookupStruct(mName, sName)
}

// LookupEnum looks up an enum by module and enum name
func (s System) LookupEnum(mName, eName string) *Enum {
	m := s.LookupModule(mName)
	if m == nil {
		return nil
	}
	return m.LookupEnum(mName, eName)
}

func (s System) LookupProperty(mName, iName, pName string) *TypedNode {
	i := s.LookupInterface(mName, iName)
	if i == nil {
		return nil
	}
	return i.LookupProperty(pName)
}

func (s System) LookupOperation(mName, iName, oName string) *Operation {
	i := s.LookupInterface(mName, iName)
	if i == nil {
		return nil
	}
	return i.LookupOperation(oName)
}

func (s System) LookupSignal(mName, iName, sName string) *Signal {
	i := s.LookupInterface(mName, iName)
	if i == nil {
		return nil
	}
	return i.LookupSignal(sName)
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
	log.Debug().Msgf("system %s: resolved %s", s.Name, s.Checksum)
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

func FQNSplit2(fqn string) (string, string) {
	parts := strings.Split(fqn, ".")
	if len(parts) < 2 {
		return "", ""
	}
	// module name is all parts except the last
	return strings.Join(parts[:len(parts)-1], "."), parts[len(parts)-1]
}

// a.b.C.d
// a.c is the module name
// C is the element name
// d is the member name
func FQNSplit3(fqn string) (string, string, string) {
	parts := strings.Split(fqn, ".")
	if len(parts) < 3 {
		return "", "", ""
	}
	// member name is the last part
	memberName := parts[len(parts)-1]
	// element name is the second to last part
	elementName := parts[len(parts)-2]
	// module name is all parts except the last two
	moduleName := strings.Join(parts[:len(parts)-2], ".")
	return moduleName, elementName, memberName
}
