package model

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec/rkw"
)

type Struct struct {
	NamedNode `json:",inline" yaml:",inline"`
	Module    *Module      `json:"-" yaml:"-"`
	Fields    []*TypedNode `json:"fields" yaml:"fields"`
}

func NewStruct(name string) *Struct {
	return &Struct{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindStruct,
		},
		Fields: make([]*TypedNode, 0),
	}
}

func (s *Struct) Validate(m *Module) error {
	// check for duplicate fields
	rkw.CheckName(s.Name, "struct")
	names := make(map[string]bool)
	if s.Fields == nil {
		s.Fields = make([]*TypedNode, 0)
	}
	for _, f := range s.Fields {
		err := f.Validate(m)
		if err != nil {
			return err
		}
		if names[f.Name] {
			return fmt.Errorf("%s: duplicate name: %s", s.Name, f.Name)
		}
		names[f.Name] = true
	}
	return nil
}

func (s *Struct) LookupField(name string) *TypedNode {
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (s *Struct) NoFields() bool {
	return len(s.Fields) == 0
}
