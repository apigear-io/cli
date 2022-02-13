package model

type Signal struct {
	NamedNode `json:",inline" yaml:",inline"`
	Inputs    []*Input `json:"params" yaml:"params"`
}

func NewSignal(name string) *Signal {
	return &Signal{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Signal,
		},
	}
}

type Property struct {
	TypedNode `json:",inline" yaml:",inline"`
}

func NewProperty(name string) *Property {
	return &Property{
		TypedNode: TypedNode{
			NamedNode: NamedNode{
				Name: name,
				Kind: Kind_Property,
			},
		},
	}
}

type Input struct {
	TypedNode `json:",inline" yaml:",inline"`
}

type Output struct {
	Schema *Schema `json:"schema" yaml:"schema"`
}

func NewMethodInput(name string) *Input {
	return &Input{
		TypedNode: TypedNode{
			NamedNode: NamedNode{
				Name: name,
				Kind: Kind_Input,
			},
		},
	}
}

func NewMethodOutput(s *Schema) *Output {
	return &Output{Schema: s}
}

type Method struct {
	NamedNode `json:",inline" yaml:",inline"`
	// maybe inputs and outputs should be a map of name to Parameter
	Inputs []*Input `json:"inputs" yaml:"inputs"`
	Output *Output  `json:"output" yaml:"output"`
}

func NewMethod(name string) *Method {
	return &Method{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Method,
		},
	}
}

type Interface struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Methods    []*Method   `json:"methods" yaml:"methods"`
	Properties []*Property `json:"properties" yaml:"properties"`
	Signals    []*Signal   `json:"signals" yaml:"signals"`
}

func NewInterface(name string) *Interface {
	return &Interface{
		NamedNode: NamedNode{
			Name: name,
			Kind: Kind_Interface,
		},
	}
}
