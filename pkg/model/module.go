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

// ShortName returns the last part of the name, i.e. the module name
func (m *Module) ShortName() string {
	parts := strings.Split(m.Name, ".")
	return parts[len(parts)-1]
}

func (m *Module) CheckImport(mName string) bool {
	for _, i := range m.Imports {
		if i.Name == mName {
			return true
		}
	}
	log.Warn().Msgf("module %s does not have import %s. Make sure to import all your dependencies.", m.Name, mName)
	return false
}

// LookupNode looks up a named node by name
func (m Module) LookupNode(mName, nName string) *NamedNode {
	i := m.LookupInterface(mName, nName)
	if i != nil {
		return &i.NamedNode
	}
	s := m.LookupStruct(mName, nName)
	if s != nil {
		return &s.NamedNode
	}
	e := m.LookupEnum(mName, nName)
	if e != nil {
		return &e.NamedNode
	}
	return nil
}

// LookupLocalInterface looks up an interface by name
func (m Module) LookupLocalInterface(iName string) *Interface {
	for _, i := range m.Interfaces {
		if i.Name == iName {
			return i
		}
	}
	return nil
}

// LookupInterface looks up an interface by name
func (m Module) LookupInterface(mName, iName string) *Interface {
	if mName == "" || mName == m.Name {
		return m.LookupLocalInterface(iName)
	}
	if m.System != nil {
		m.CheckImport(mName)
		return m.System.LookupInterface(mName, iName)
	}
	return nil
}

// LookupLocalStruct looks up a struct by name
func (m Module) LookupLocalStruct(sName string) *Struct {
	for _, s := range m.Structs {
		if s.Name == sName {
			return s
		}
	}
	return nil
}

// LookupStruct looks up a struct by name
func (m Module) LookupStruct(mName, sName string) *Struct {
	if mName == "" || mName == m.Name {
		return m.LookupLocalStruct(sName)
	}
	if m.System != nil {
		m.CheckImport(mName)
		return m.System.LookupStruct(mName, sName)
	}
	return nil
}

func (m Module) LookupLocalEnum(eName string) *Enum {
	for _, e := range m.Enums {
		if e.Name == eName {
			return e
		}
	}
	return nil
}

// LookupEnum looks up an enum by name
func (m Module) LookupEnum(mName, eName string) *Enum {
	if mName == "" || mName == m.Name {
		return m.LookupLocalEnum(eName)
	}
	if m.System != nil {
		m.CheckImport(mName)
		return m.System.LookupEnum(mName, eName)
	}
	return nil
}

// LookupLocalDefaultEnumMember looks up the enum by name and returns the default member
func (m Module) LookupLocalDefaultEnumMember(eName string) *EnumMember {
	e := m.LookupLocalEnum(eName)
	if e != nil {
		return e.Default()
	}
	return nil
}

// LookupDefaultEnumMember looks up the enum by name and returns the default member
func (m Module) LookupDefaultEnumMember(mName, eName string) *EnumMember {
	e := m.LookupEnum(mName, eName)
	if e != nil {
		return e.Default()
	}
	return nil
}

func (m *Module) Validate() error {
	m.compute()
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

func (m *Module) compute() {
	for _, i := range m.Interfaces {
		i.Module = m
	}
	for _, s := range m.Structs {
		s.Module = m
	}
	for _, e := range m.Enums {
		e.Module = m
	}
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

func (m *Module) CheckReservedWords(langs []rkw.Lang) {
	rkw.CheckIsReserved(langs, m.Name, "module")
	for _, i := range m.Interfaces {
		i.CheckReservedWords(langs)
	}
	for _, s := range m.Structs {
		s.CheckReservedWords(langs)
	}
	for _, e := range m.Enums {
		e.CheckReservedWords(langs)
	}
}
