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
	TypeBool      KindType = "bool"
	TypeInt       KindType = "int"
	TypeInt32     KindType = "int32"
	TypeInt64     KindType = "int64"
	TypeFloat     KindType = "float"
	TypeFloat32   KindType = "float32"
	TypeFloat64   KindType = "float64"
	TypeString    KindType = "string"
	TypeEnum      KindType = "enum"
	TypeStruct    KindType = "struct"
	TypeInterface KindType = "interface"
)

type ITypeProvider interface {
	GetName() string
	GetKind() Kind
	GetSchema() *Schema
}

type IModuleProvider interface {
	GetModule() *Module
}

// NamedNode is a base node with a name and a kind.
// { "name": "foo", "kind": "interface" }
type NamedNode struct {
	Name        string                 `json:"name" yaml:"name"`
	Kind        Kind                   `json:"kind" yaml:"kind"`
	Description string                 `json:"description" yaml:"description"`
	Meta        map[string]interface{} `json:"meta" yaml:"meta"`
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

// GetKind returns the kind of the node
func (t TypedNode) GetKind() Kind {
	return t.Kind
}

// GetName returns the name of the node
func (t TypedNode) GetName() string {
	return t.Name
}

// GetSchema returns the schema of the node
func (t TypedNode) GetSchema() *Schema {
	return &t.Schema
}

// Validate resolves all the types in the schema
func (t *TypedNode) Validate(m *Module) error {
	rkw.CheckName(t.Name, "type")
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

// TypeNode is a node with type information.
// { type: array, items: { type: string } }
type Schema struct {
	Type        string     `json:"type" yaml:"type"`
	IsArray     bool       `json:"array" yaml:"array"`
	Module      *Module    `json:"-" yaml:"-"`
	KindType    KindType   `json:"kindType" yaml:"kindType"`
	struct_     *Struct    `json:"-" yaml:"-"`
	enum        *Enum      `json:"-" yaml:"-"`
	interface_  *Interface `json:"-" yaml:"-"`
	IsPrimitive bool       `json:"isPrimitive" yaml:"isPrimitive"`
	IsSymbol    bool       `json:"isSymbol" yaml:"isSymbol"`
	IsResolved  bool       `json:"-" yaml:"-"`
}

// IsEmpty returns true if the schema is empty
func (s Schema) IsEmpty() bool {
	return s.Type == ""
}

// Lookup returns the node with the given name inside the module
func (s Schema) LookupNode(name string) *NamedNode {
	return s.Module.LookupNode(name)
}

// Validate resolves all the types in the schema
func (s *Schema) Validate(m *Module) error {
	if s.IsResolved {
		return nil
	}
	if s.Type == "" {
		s.Type = "void"
	}
	s.Module = m
	switch s.Type {
	case "void":
		s.IsPrimitive = false
		s.IsSymbol = false
	case "bool", "int", "float", "string", "int32", "int64", "float32", "float64":
		s.IsPrimitive = true
		s.IsSymbol = false
	default:
		s.IsPrimitive = false
		s.IsSymbol = true
	}
	err := s.resolveSymbol()
	if err != nil {
		return err
	}
	s.resolveType()
	s.IsResolved = true
	return nil
}

func (s *Schema) resolveSymbol() error {
	if s.IsResolved {
		return nil
	}
	if s.IsSymbol {
		le := s.Module.LookupEnum(s.Type)
		if le != nil {
			s.enum = le
			s.KindType = TypeEnum
			return nil
		}
		ls := s.Module.LookupStruct(s.Type)
		if ls != nil {
			s.struct_ = ls
			s.KindType = TypeStruct
			return nil
		}
		li := s.Module.LookupInterface(s.Type)
		if li != nil {
			s.interface_ = li
			s.KindType = TypeInterface
			return nil
		}
		return fmt.Errorf("symbol %s not found", s.Type)
	}
	return nil
}

func (s *Schema) resolveType() {
	if s.IsResolved {
		return
	}
	var kind KindType
	if s.IsPrimitive {
		kind = KindType(s.Type)
	} else if s.IsSymbol {
		if s.IsInterface() {
			kind = TypeInterface
		} else if s.IsStruct() {
			kind = TypeStruct
		} else if s.IsEnum() {
			kind = TypeEnum
		}
	} else {
		kind = TypeVoid
	}
	s.KindType = KindType(kind)
}

func (s *Schema) GetEnum() *Enum {
	err := s.resolveSymbol()
	if err != nil {
		return nil
	}

	return s.enum
}

func (s *Schema) GetStruct() *Struct {
	err := s.resolveSymbol()
	if err != nil {
		return nil
	}
	return s.struct_
}

func (s *Schema) GetInterface() *Interface {
	err := s.resolveSymbol()
	if err != nil {
		return nil
	}
	return s.interface_
}

func (s *Schema) IsEnum() bool {
	return s.GetEnum() != nil
}

func (s *Schema) IsStruct() bool {
	return s.GetStruct() != nil
}

func (s *Schema) IsInterface() bool {
	return s.GetInterface() != nil
}
