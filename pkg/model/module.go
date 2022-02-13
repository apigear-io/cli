package model

type Import struct {
	NamedNode `json:",inline" yaml:",inline"`
	Version   string `json:"version" yaml:"version"`
}

func NewImport(name string, version string) *Import {
	return &Import{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Import,
		},
		Version: version,
	}
}

type Module struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Version    string       `json:"version" yaml:"version"`
	Imports    []*Import    `json:"imports" yaml:"imports"`
	Interfaces []*Interface `json:"interfaces" yaml:"interfaces"`
	Structs    []*Struct    `json:"structs" yaml:"structs"`
	Enums      []*Enum      `json:"enums" yaml:"enums"`
}

func NewModule(name string, version string) *Module {
	return &Module{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Module,
		},
		Version:    version,
		Imports:    make([]*Import, 0),
		Interfaces: make([]*Interface, 0),
		Structs:    make([]*Struct, 0),
		Enums:      make([]*Enum, 0),
	}
}
