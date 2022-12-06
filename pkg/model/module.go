package model

import (
	"strconv"
	"strings"
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
	Version   string `json:"version" yaml:"version"`
}

func NewImport(name string, version string) *Import {
	return &Import{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindImport,
		},
		Version: version,
	}
}

type Module struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Version    Version      `json:"version" yaml:"version"`
	Imports    []*Import    `json:"imports" yaml:"imports"`
	Interfaces []*Interface `json:"interfaces" yaml:"interfaces"`
	Structs    []*Struct    `json:"structs" yaml:"structs"`
	Enums      []*Enum      `json:"enums" yaml:"enums"`
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

func (m *Module) ResolveAll() error {
	for _, i := range m.Interfaces {
		err := i.ResolveAll(m)
		if err != nil {
			return err
		}
	}
	for _, s := range m.Structs {
		err := s.ResolveAll(m)
		if err != nil {
			return err
		}
	}
	for _, e := range m.Enums {
		err := e.ResolveAll(m)
		if err != nil {
			return err
		}
	}
	return nil
}
