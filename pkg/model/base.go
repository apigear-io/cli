package model

// Kind is an enumeration of the kinds of nodes.
type Kind string

const (
	KindSystem    Kind = "system"
	KindModule    Kind = "module"
	KindImport    Kind = "import"
	KindInterface Kind = "interface"
	KindProperty  Kind = "property"
	KindMethod    Kind = "method"
	KindInput     Kind = "input"
	KindSignal    Kind = "signal"
	KindStruct    Kind = "struct"
	KindField     Kind = "field"
	KindEnum      Kind = "enum"
	KindMember    Kind = "member"
)

type ISchemaProvider interface {
	GetSchema() Schema
}

// NamedNode is a base node with a name and a kind.
// { "name": "foo", "kind": "interface" }
type NamedNode struct {
	Name string `json:"name" yaml:"name"`
	Kind Kind   `json:"kind" yaml:"kind"`
}

func (n *NamedNode) String() string {
	return n.Name
}

func (n NamedNode) IsEmpty() bool {
	return n.Name == ""
}

// TypedNode is a base node with a schema type.
// { name: "foo", kind: "property", type: "string" }
type TypedNode struct {
	NamedNode `json:",inline" yaml:",inline"`
	Schema    Schema `json:"schema" yaml:"schema"`
}

func InitTypeNode(n string, k Kind) TypedNode {
	return TypedNode{
		NamedNode: NamedNode{
			Name: n,
			Kind: k,
		},
	}
}

func (t TypedNode) GetSchema() Schema {
	return t.Schema
}

// TypeNode is a node with type information.
// { type: array, items: { type: string } }
type Schema struct {
	Type string `json:"type" yaml:"type"`
}

func (s Schema) IsEmpty() bool {
	return s.Type == ""
}
