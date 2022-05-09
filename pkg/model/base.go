package model

import (
	"fmt"
	"objectapi/pkg/log"
)

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
	KindOutput    Kind = "output"
	KindSignal    Kind = "signal"
	KindStruct    Kind = "struct"
	KindField     Kind = "field"
	KindEnum      Kind = "enum"
	KindMember    Kind = "member"
)

type KindType string

const (
	TypeBool      = "bool"
	TypeInt       = "int"
	TypeFloat     = "float"
	TypeString    = "string"
	TypeEnum      = "enum"
	TypeStruct    = "struct"
	TypeInterface = "interface"
	TypeUnknown   = "unknown"
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
	Schema    *Schema `json:"schema" yaml:"schema"`
}

func NewTypeNode(n string, k Kind) *TypedNode {
	return &TypedNode{
		NamedNode: NamedNode{
			Name: n,
			Kind: k,
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
	return t.Schema
}

func (t *TypedNode) ResolveAll() error {
	err := t.Schema.ResolveSymbol()
	if err != nil {
		return err
	}
	return t.Schema.ResolveType()
}

// TypeNode is a node with type information.
// { type: array, items: { type: string } }
type Schema struct {
	Type        string `json:"type" yaml:"type"`
	Module      *Module
	KindType    KindType
	struct_     *Struct
	enum        *Enum
	interface_  *Interface
	IsArray     bool
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

func (s *Schema) ResolveSymbol() error {
	if s.IsResolved {
		return nil
	}
	s.IsResolved = true
	if s.IsSymbol {
		le := s.Module.LookupEnum(s.Type)
		if le == nil {
			return fmt.Errorf("unknown symbol: %s", s.Type)
		} else {
			s.enum = le
			s.KindType = TypeEnum
			return nil
		}
		ls := s.Module.LookupStruct(s.Type)
		if ls == nil {
			return fmt.Errorf("unknown symbol %s", s.Type)
		} else {
			s.struct_ = ls
			s.KindType = TypeStruct
			return nil
		}
		li := s.Module.LookupInterface(s.Type)
		if li == nil {
			return fmt.Errorf("unknown symbol %s", s.Type)
		} else {
			s.interface_ = li
			s.KindType = TypeInterface
			return nil
		}
	}
	return nil
}

func (s *Schema) ResolveType() error {
	if !s.IsResolved {
		s.ResolveSymbol()
	}
	kind := ""
	if s.IsPrimitive {
		kind = s.Type
	} else if s.IsSymbol {
		if s.IsInterface() {
			kind = "interface"
		} else if s.IsStruct() {
			kind = "struct"
		} else if s.IsEnum() {
			kind = "enum"
		}
	}
	if kind == "" {
		log.Warnf("resolved type %s to %s", s.Type, kind)
		kind = "unknown"
		return fmt.Errorf("unknown type %s", s.Type)
	}
	s.KindType = KindType(kind)
	return nil
}

func (s *Schema) GetEnum() *Enum {
	s.ResolveSymbol()
	return s.enum
}

func (s *Schema) GetStruct() *Struct {
	s.ResolveSymbol()
	return s.struct_
}

func (s *Schema) GetInterface() *Interface {
	s.ResolveSymbol()
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
