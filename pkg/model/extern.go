package model

type Extern struct {
	NamedNode `json:",inline" yaml:",inline"`
}

func NewExtern(name string) *Extern {
	return &Extern{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindExtern,
		},
	}
}

func (e *Extern) Validate(m *Module) error {
	return nil
}

func (e *Extern) IsEmpty() bool {
	return e.Name == ""
}
