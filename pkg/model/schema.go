package model

import "strings"

// TypeNode is a node with type information.
// { type: array, items: { type: string } }
type Schema struct {
	Type        string     `json:"type" yaml:"type"`
	Import      string     `json:"import" yaml:"import"`
	IsArray     bool       `json:"array" yaml:"array"`
	Module      *Module    `json:"-" yaml:"-"`
	KindType    KindType   `json:"kindType" yaml:"kindType"`
	IsPrimitive bool       `json:"isPrimitive" yaml:"isPrimitive"`
	IsSymbol    bool       `json:"isSymbol" yaml:"isSymbol"`
	struct_     *Struct    `json:"-" yaml:"-"`
	enum_       *Enum      `json:"-" yaml:"-"`
	interface_  *Interface `json:"-" yaml:"-"`
	isResolved  bool       `json:"-" yaml:"-"`
	isComputed  bool       `json:"-" yaml:"-"`
	isValid     bool       `json:"-" yaml:"-"`
}

func (s *Schema) IsImported() bool {
	return s.Import != ""
}

func (s *Schema) ShortImportName() string {
	parts := strings.Split(s.Import, ".")
	return parts[len(parts)-1]
}

// IsEmpty returns true if the schema is empty
func (s Schema) IsEmpty() bool {
	return s.Type == ""
}

// Lookup returns the node with the given name inside the module
func (s Schema) LookupNode(mName, nName string) *NamedNode {
	if s.Module == nil {
		return nil
	}
	return s.Module.LookupNode(mName, nName)
}

func (s *Schema) compute() {
	if s.isComputed {
		return
	}
	s.isComputed = true
	if s.Type == "" {
		s.Type = "void"
	}
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
}

// Validate resolves all the types in the schema
func (s *Schema) Validate(m *Module) error {
	if s.isValid {
		return nil
	}
	s.isValid = true
	s.Module = m
	s.compute()
	s.resolve()
	return nil
}

func (s *Schema) resolve() {
	if s.isResolved {
		return
	}
	s.isResolved = true
	switch {
	case s.IsPrimitive:
		s.KindType = KindType(s.Type)
	case s.IsSymbol:
		le := s.Module.LookupEnum(s.Import, s.Type)
		if le != nil {
			s.enum_ = le
			s.KindType = TypeEnum
			break
		}
		ls := s.Module.LookupStruct(s.Import, s.Type)
		if ls != nil {
			s.struct_ = ls
			s.KindType = TypeStruct
			break
		}
		li := s.Module.LookupInterface(s.Import, s.Type)
		if li != nil {
			s.interface_ = li
			s.KindType = TypeInterface
			break
		}
	default:
		s.KindType = TypeVoid
	}
}

func (s *Schema) GetEnum() *Enum {
	s.resolve()
	return s.enum_
}

func (s *Schema) GetStruct() *Struct {
	s.resolve()
	return s.struct_
}

func (s *Schema) GetInterface() *Interface {
	s.resolve()
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

func (s Schema) InnerSchema() Schema {
	s.IsArray = false
	return s
}

func (s Schema) System() *System {
	if s.Module == nil {
		return nil
	}
	return s.Module.System
}
