package model

import (
	"fmt"
	"strings"

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
	TypeVoid      KindType = "null"
	TypeBool      KindType = "bool"
	TypeInt       KindType = "int"
	TypeFloat     KindType = "float"
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
	Name        string `json:"name" yaml:"name"`
	Kind        Kind   `json:"kind" yaml:"kind"`
	Description string `json:"description" yaml:"description"`
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
	NamedNode `json:",inline" yaml:",inline"`
	Schema    `json:",inline" yaml:",inline"`
}

func NewTypedNode(n string, k Kind) *TypedNode {
	return &TypedNode{
		NamedNode: NamedNode{
			Name: n,
			Kind: k,
		},
		Schema: Schema{
			Type:     "",
			KindType: TypeVoid,
		},
	}
}

func (t TypedNode) GetKind() Kind {
	return t.Kind
}

func (t TypedNode) GetName() string {
	return t.Name
}

func (t TypedNode) GetSchema() *Schema {
	return &t.Schema
}

func (t *TypedNode) ResolveAll(m *Module) error {
	return t.Schema.ResolveAll(m)
}

func (t *TypedNode) NoType() bool {
	return t.Type == ""
}

func (t *TypedNode) IsVoid() bool {
	return t.Type == ""
}

func (t TypedNode) HasType() bool {
	return t.Type != ""
}

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
	Type        string `json:"type" yaml:"type"`
	IsArray     bool   `json:"array" yaml:"array"`
	Module      *Module
	KindType    KindType
	struct_     *Struct
	enum        *Enum
	interface_  *Interface
	IsPrimitive bool
	IsSymbol    bool
	IsResolved  bool
}

func (s Schema) IsEmpty() bool {
	return s.Type == ""
}

// Lookup returns the node with the given name inside the module
func (s Schema) LookupNode(name string) *NamedNode {
	return s.Module.LookupNode(name)
}

func (s *Schema) ResolveAll(m *Module) error {
	if s.IsResolved {
		return nil
	}
	s.Module = m
	switch s.Type {
	case "":
		s.IsPrimitive = false
		s.IsSymbol = false
	case "bool", "int", "float", "string":
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
	s.resolveSymbol()

	return s.enum
}

func (s *Schema) GetStruct() *Struct {
	s.resolveSymbol()
	return s.struct_
}

func (s *Schema) GetInterface() *Interface {
	s.resolveSymbol()
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
