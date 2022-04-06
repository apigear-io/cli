package model

// Kind is an enumeration of the kinds of nodes.
type Kind string

const (
	Kind_System    Kind = "system"
	Kind_Module    Kind = "module"
	Kind_Import    Kind = "import"
	Kind_Interface Kind = "interface"
	Kind_Property  Kind = "property"
	Kind_Method    Kind = "method"
	Kind_Input     Kind = "input"
	Kind_Signal    Kind = "signal"
	Kind_Struct    Kind = "struct"
	Kind_Field     Kind = "field"
	Kind_Enum      Kind = "enum"
	Kind_Member    Kind = "member"
)

type TypeKind string

// NamedNode is a base node with a name and a kind.
type NamedNode struct {
	Name string `json:"name" yaml:"name"`
	Kind Kind   `json:"kind" yaml:"kind"`
}

func (n *NamedNode) String() string {
	return n.Name
}

// TypedNode is a base node with a schema type.
type TypedNode struct {
	NamedNode `json:",inline" yaml:",inline"`
	Schema    *Schema `json:"schema" yaml:"schema"`
}

// Schema is a type definition.
type Schema struct {
	Type   string `json:"type" yaml:"type"`
	Items  string `json:"items" yaml:"items"`
	Format string `json:"format" yaml:"format"`
}

func NewSchema() *Schema {
	return &Schema{}
}

func IsPrimitive(s string) bool {
	switch s {
	case "bool", "int", "float", "string":
		return true
	default:
		return false
	}
}
