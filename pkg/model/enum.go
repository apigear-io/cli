package model

// Enum is an enumeration.
type Enum struct {
	NamedNode `json:",inline" yaml:",inline"`
	Members   []*EnumMember `json:"members" yaml:"members"`
}

// NewEnum creates a new Enum.
func NewEnum(name string) *Enum {
	return &Enum{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Enum,
		},
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
			Kind: Kind_Member,
		},
		Value: value,
	}
}
