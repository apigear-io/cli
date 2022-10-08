package model

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

func (e *Enum) ResolveAll(mod *Module) error {
	autoValue := true
	for _, mem := range e.Members {
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

func (e *Enum) LookupMember(name string) *EnumMember {
	for _, mem := range e.Members {
		if mem.Name == name {
			return mem
		}
	}
	return nil
}

func (e *Enum) Default() *EnumMember {
	if len(e.Members) > 0 {
		return e.Members[0]
	}
	return &EnumMember{}
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

func (e *EnumMember) ResolveAll(m *Module) error {
	return nil
}
