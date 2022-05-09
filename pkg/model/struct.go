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

func (s *Struct) ResolveAll() error {
	for _, f := range s.Fields {
		err := f.ResolveAll()
		if err != nil {
			return err
		}
	}
	return nil
}
