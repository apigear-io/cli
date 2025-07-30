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

// AcceptModelVisitor implements the AcceptModelVisitor interface for Struct
func (s *Struct) AcceptModelVisitor(v ModelVisitor) error {
	err := v.VisitStruct(s)
	if err != nil {
		return err
	}
	for _, f := range s.Fields {
		err = f.AcceptModelVisitor(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Struct) Validate(m *Module) error {
	// check for duplicate fields
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

func (s *Struct) CheckReservedWords(langs []rkw.Lang) {
	rkw.CheckIsReserved(langs, s.Name, "struct")
	for _, f := range s.Fields {
		f.CheckReservedWords(langs)
	}
}
