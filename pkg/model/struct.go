package model

type Struct struct {
	NamedNode `json:",inline" yaml:",inline"`
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

func (s *Struct) ResolveAll(m *Module) error {
	if s.Fields == nil {
		s.Fields = make([]*TypedNode, 0)
	}
	for _, f := range s.Fields {
		err := f.ResolveAll(m)
		if err != nil {
			return err
		}
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
