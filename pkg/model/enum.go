package model

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec/rkw"
)

// Enum is an enumeration.
type Enum struct {
	NamedNode `json:",inline" yaml:",inline"`
	Module    *Module       `json:"-" yaml:"-"`
	Members   []*EnumMember `json:"members" yaml:"members"`
}

// InitEnum creates a new Enum.
func NewEnum(name string) *Enum {
	return &Enum{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindEnum,
		},
	}
}

// Validate resolves all references in the enum.
func (e *Enum) Validate(mod *Module) error {
	names := make(map[string]bool)
	autoValue := true
	for _, mem := range e.Members {
		err := mem.Validate(mod)
		if err != nil {
			return err
		}
		if names[mem.Name] {
			return fmt.Errorf("%s: duplicate name: %s", e.Name, mem.Name)
		}
		names[mem.Name] = true
		if mem.Value != 0 {
			autoValue = false
		}
	}
	if autoValue {
		for i, mem := range e.Members {
			mem.Value = i
		}
	}
	return nil
}

// LookupMember returns the member with the given name, or nil if no such member exists.
func (e *Enum) LookupMember(name string) *EnumMember {
	for _, mem := range e.Members {
		if mem.Name == name {
			return mem
		}
	}
	return nil
}

// Default returns the default member of the enum, which is the first member
func (e *Enum) Default() *EnumMember {
	if len(e.Members) > 0 {
		return e.Members[0]
	}
	return &EnumMember{}
}

// NoMembers returns true if the enum has no members.
func (e *Enum) NoMembers() bool {
	return len(e.Members) == 0
}

// CheckReservedWords checks the names of the enum.
func (e *Enum) CheckReservedWords(langs []rkw.Lang) {
	rkw.CheckIsReserved(langs, e.Name, "enum")
	for _, mem := range e.Members {
		mem.CheckReservedWords(langs)
	}
}

// EnumMember is a member of an enumeration.
type EnumMember struct {
	NamedNode `json:",inline" yaml:",inline"`
	Value     int `json:"value" yaml:"value"`
}

// NewEnumMember creates a new EnumMember.
func NewEnumMember(name string, value int) *EnumMember {
	return &EnumMember{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindMember,
		},
		Value: value,
	}
}

// Validate resolves all references in the enum member.
func (e *EnumMember) Validate(m *Module) error {
	return nil
}

// CheckReservedWords checks the names of the enum member.
func (e *EnumMember) CheckReservedWords(langs []rkw.Lang) {
	rkw.CheckIsReserved(langs, e.Name, "enum member")
}
