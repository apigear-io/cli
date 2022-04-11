package model

type Import struct {
	NamedNode `json:",inline" yaml:",inline"`
	Version   string `json:"version" yaml:"version"`
}

func InitImport(name string, version string) Import {
	return Import{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindImport,
		},
		Version: version,
	}
}

type Module struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Version    string      `json:"version" yaml:"version"`
	Imports    []Import    `json:"imports" yaml:"imports"`
	Interfaces []Interface `json:"interfaces" yaml:"interfaces"`
	Structs    []Struct    `json:"structs" yaml:"structs"`
	Enums      []Enum      `json:"enums" yaml:"enums"`
}

func InitModule(n string, v string) Module {
	return Module{
		NamedNode: NamedNode{
			Name: n,
			Kind: KindModule,
		},
		Version: v,
	}
}

func (m Module) InterfaceByName(name string) Interface {
	for _, i := range m.Interfaces {
		if i.Name == name {
			return i
		}
	}
	return Interface{}
}

func (m Module) StructByName(name string) Struct {
	for _, s := range m.Structs {
		if s.Name == name {
			return s
		}
	}
	return Struct{}
}

func (m Module) EnumByName(name string) Enum {
	for _, e := range m.Enums {
		if e.Name == name {
			return e
		}
	}
	return Enum{}
}
