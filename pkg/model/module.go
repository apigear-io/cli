package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/apigear-io/cli/pkg/spec/rkw"
)

type Version string

func (v Version) String() string {
	return string(v)
}

func (v Version) parts() []int {
	parts := strings.Split(v.String(), ".")
	result := make([]int, len(parts))
	for idx, p := range parts {
		result[idx], _ = strconv.Atoi(p)
	}
	return result
}
func (v Version) Major() int {
	parts := v.parts()
	if len(parts) < 1 {
		return 0
	}
	return parts[0]
}

func (v Version) Minor() int {
	parts := v.parts()
	if len(parts) < 2 {
		return 0
	}
	return parts[1]
}

func (v Version) Patch() int {
	parts := v.parts()
	if len(parts) < 3 {
		return 0
	}
	return parts[2]
}

type Import struct {
	NamedNode `json:",inline" yaml:",inline"`
}

func NewImport(name string, version string) *Import {
	return &Import{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindImport,
		},
	}
}

type Module struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Version    Version      `json:"version" yaml:"version"`
	Imports    []*Import    `json:"imports" yaml:"imports"`
	Interfaces []*Interface `json:"interfaces" yaml:"interfaces"`
	Structs    []*Struct    `json:"structs" yaml:"structs"`
	Enums      []*Enum      `json:"enums" yaml:"enums"`
	Checksum   string       `json:"checksum" yaml:"checksum"`
	System     *System      `json:"-"` // reference to the parent system
}

func NewModule(n string, v string) *Module {
	return &Module{
		NamedNode: NamedNode{
			Name: n,
			Kind: KindModule,
		},
		Version:    Version(v),
		Imports:    make([]*Import, 0),
		Interfaces: make([]*Interface, 0),
		Structs:    make([]*Struct, 0),
		Enums:      make([]*Enum, 0),
	}
}

func (m Module) LookupNode(name string) *NamedNode {
	for _, i := range m.Interfaces {
		if i.Name == name {
			return &i.NamedNode
		}
	}
	for _, s := range m.Structs {
		if s.Name == name {
			return &s.NamedNode
		}
	}
	for _, e := range m.Enums {
		if e.Name == name {
			return &e.NamedNode
		}
	}
	return nil
}

func (m Module) LookupInterface(name string) *Interface {
	for _, i := range m.Interfaces {
		if i.Name == name {
			return i
		}
	}
	return nil
}

func (m Module) LookupStruct(name string) *Struct {
	for _, s := range m.Structs {
		if s.Name == name {
			return s
		}
	}
	return nil
}

func (m Module) LookupEnum(name string) *Enum {
	for _, e := range m.Enums {
		if e.Name == name {
			return e
		}
	}
	return nil
}

func (m Module) LookupDefaultEnumMember(name string) *EnumMember {
	e := m.LookupEnum(name)
	if e != nil {
		return e.Members[0]
	}
	return nil
}

func (m *Module) Validate() error {
	rkw.CheckName(m.Name, "module")
	// check for duplicate names
	names := make(map[string]bool)
	for _, i := range m.Interfaces {
		err := i.Validate(m)
		if err != nil {
			return err
		}
		if names[i.Name] {
			return fmt.Errorf("%s: duplicate name %s", m.Name, i.Name)
		}
		names[i.Name] = true
	}
	for _, s := range m.Structs {
		err := s.Validate(m)
		if err != nil {
			return err
		}
		if names[s.Name] {
			return fmt.Errorf("%s: duplicate name %s", m.Name, s.Name)
		}
		names[s.Name] = true
	}
	for _, e := range m.Enums {
		err := e.Validate(m)
		if err != nil {
			return err
		}
		if names[e.Name] {
			return fmt.Errorf("%s: duplicate name %s", m.Name, e.Name)
		}
		names[e.Name] = true
	}
	m.computeChecksum()
	return nil
}

// TODO: clean up this code
func (m *Module) computeChecksum() {
	var buffer strings.Builder
	buffer.WriteString(m.Name)
	for _, i := range m.Interfaces {
		buffer.WriteString(i.Name)
		for _, o := range i.Operations {
			buffer.WriteString(o.Name)
			for _, p := range o.Params {
				buffer.WriteString(p.Name)
				buffer.WriteString(p.Type)
			}
			buffer.WriteString(o.Return.Type)
		}
		for _, p := range i.Properties {
			buffer.WriteString(p.Name)
			buffer.WriteString(p.Type)
		}
		for _, s := range i.Signals {
			buffer.WriteString(s.Name)
			for _, p := range s.Params {
				buffer.WriteString(p.Name)
				buffer.WriteString(p.Type)
			}
		}
	}
	for _, s := range m.Structs {
		buffer.WriteString(s.Name)
		for _, p := range s.Fields {
			buffer.WriteString(p.Name)
			buffer.WriteString(p.Type)
		}
	}
	for _, e := range m.Enums {
		buffer.WriteString(e.Name)
		for _, p := range e.Members {
			buffer.WriteString(p.Name)
			buffer.WriteString(strconv.Itoa(p.Value))
		}
	}
	sum := md5.Sum([]byte(buffer.String()))
	m.Checksum = hex.EncodeToString(sum[:])
}
