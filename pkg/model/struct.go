package model

type Struct struct {
	NamedNode `json:",inline" yaml:",inline"`
	Fields    []*StructField `json:"fields" yaml:"fields"`
}

func NewStruct(name string) *Struct {
	return &Struct{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Struct,
		},
		Fields: make([]*StructField, 0),
	}
}

type StructField struct {
	NamedNode `json:",inline" yaml:",inline"`
	Schema    *Schema `json:"schema" yaml:"schema"`
}

func NewStructField(name string) *StructField {
	return &StructField{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Field,
		},
	}
}
