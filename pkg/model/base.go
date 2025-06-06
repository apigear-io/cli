package model

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/spec/rkw"
	"github.com/ettle/strcase"
)

// Kind is an enumeration of the kinds of nodes.
type Kind string

const (
	KindSystem    Kind = "system"
	KindModule    Kind = "module"
	KindImport    Kind = "import"
	KindExtern    Kind = "extern"
	KindInterface Kind = "interface"
	KindProperty  Kind = "property"
	KindOperation Kind = "operation"
	KindParam     Kind = "param"
	KindReturn    Kind = "return"
	KindSignal    Kind = "signal"
	KindStruct    Kind = "struct"
	KindField     Kind = "field"
	KindEnum      Kind = "enum"
	KindMember    Kind = "member"
)

type KindType string

const (
	TypeVoid      KindType = "void"
	TypeAny       KindType = "any"
	TypeBool      KindType = "bool"
	TypeInt       KindType = "int"
	TypeInt32     KindType = "int32"
	TypeInt64     KindType = "int64"
	TypeFloat     KindType = "float"
	TypeFloat32   KindType = "float32"
	TypeFloat64   KindType = "float64"
	TypeString    KindType = "string"
	TypeBytes     KindType = "bytes"
	TypeExtern    KindType = "extern"
	TypeEnum      KindType = "enum"
	TypeStruct    KindType = "struct"
	TypeInterface KindType = "interface"
)

// Meta is a map of string to interface
// It is used to store additional information in a node
type Meta map[string]interface{}

// Has returns true if the key exists in the meta
func (m Meta) Has(key string) bool {
	_, ok := m[key]
	return ok
}

// Get returns a value from the meta
func (m Meta) Get(key string) interface{} {
	return m[key]
}

// Set sets a key value in the meta
func (m Meta) Set(key string, value interface{}) {
	m[key] = value
}

// GetString returns a string value from the meta
func (m Meta) GetString(key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

// GetBool returns a bool value from the meta
func (m Meta) GetBool(key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

// GetInt returns an int value from the meta
func (m Meta) GetInt(key string) int {
	if v, ok := m[key].(int); ok {
		return v
	}
	return 0
}

// GetFloat returns a float value from the meta
func (m Meta) GetFloat(key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	return 0
}

// AssertKey asserts that a key exists in the meta
func (m Meta) AssertKey(key string) (string, error) {
	if !m.Has(key) {
		return "", fmt.Errorf("missing key %s in meta", key)

	}
	return "", nil
}

// NamedNode is a base node with a name and a kind.
// { "name": "foo", "kind": "interface" }
type NamedNode struct {
	Id          uint   `json:"-" yaml:"-"` // internal id
	Name        string `json:"name" yaml:"name"`
	Kind        Kind   `json:"kind" yaml:"kind"`
	Description string `json:"description" yaml:"description"`
	Meta        Meta   `json:"meta" yaml:"meta"`
}

func (n *NamedNode) String() string {
	return n.Name
}

func (n *NamedNode) ShortName() string {
	words := strings.Split(n.Name, ".")
	return words[len(words)-1]
}

func (n *NamedNode) AsPath() string {
	return strcase.ToCase(n.Name, strcase.LowerCase, '/')
}

func (n NamedNode) IsEmpty() bool {
	return n.Name == ""
}

// TypedNode is a base node with a schema type.
// { name: "foo", kind: "property", type: "string" }
type TypedNode struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Schema     `json:",inline" yaml:",inline"`
	IsReadOnly bool `json:"readonly" yaml:"readonly"`
}

// NewTypedNode creates a new typed node
func NewTypedNode(n string, k Kind) *TypedNode {
	return &TypedNode{
		NamedNode: NamedNode{
			Name: n,
			Kind: k,
		},
		Schema: Schema{
			Type:     "void",
			KindType: TypeVoid,
		},
	}
}

// Validate resolves all the types in the schema
func (t *TypedNode) Validate(m *Module) error {
	return t.Schema.Validate(m)
}

// IsVoid returns true if the schema is void
func (t *TypedNode) IsVoid() bool {
	return t.Type == "void"
}

// IsNotVoid returns true if the schema is not void
func (s Schema) IsNotVoid() bool {
	return s.Type != "void"
}

func (t *TypedNode) IsAny() bool {
	return t.Type == "any"
}

func (t *TypedNode) IsInt() bool {
	return t.Type == "int" || t.Type == "int32" || t.Type == "int64"
}

// IsFloat returns true if the schema is a float (e.g. float, float32, float64)
func (t *TypedNode) IsFloat() bool {
	return t.Type == "float" || t.Type == "float32" || t.Type == "float64"
}

// TypeName returns the name of the type, e.g. String, Int32, Int32Array, InterfaceFoo, StructBar, EnumBaz
// Can be used to call conversion functions based on type name
func (t TypedNode) TypeName() string {
	// if IsArray prefix with Array
	// is is isSymbol prefix with Interface, Struct, Enum
	// if is primitive append primitive name
	s := ""
	if t.IsSymbol {
		s += strcase.ToCamel(t.Type)
	} else if t.IsPrimitive {
		s += strcase.ToCamel(t.Type)
	}
	if t.IsArray {
		s += "Array"
	}
	return s
}

func (t *TypedNode) CheckReservedWords(langs []rkw.Lang) {
	rkw.CheckIsReserved(langs, t.Name, "type")
}
