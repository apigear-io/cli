package model

import "fmt"

// Enum is an enumeration.
type Enum struct {
	NamedNode `json:",inline" yaml:",inline"`
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

// ResolveAll resolves all references in the enum.
func (e *Enum) ResolveAll(mod *Module) error {
	names := make(map[string]bool)
	autoValue := true
	for _, mem := range e.Members {
		if names[mem.Name] {
			return fmt.Errorf("%s: duplicate name: %s", e.Name, mem.Name)
		}
		names[mem.Name] = true
		err := mem.ResolveAll(mod)
		if err != nil {
			return err
		}
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

// ResolveAll resolves all references in the enum member.
func (e *EnumMember) ResolveAll(m *Module) error {
	return nil
}
