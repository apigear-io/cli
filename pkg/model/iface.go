package model

type Signal struct {
	NamedNode `json:",inline" yaml:",inline"`
	Inputs    []TypedNode `json:"inputs" yaml:"inputs"`
}

func NewSignal(name string) Signal {
	return Signal{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindSignal,
		},
	}
}

type Method struct {
	NamedNode `json:",inline" yaml:",inline"`
	// maybe inputs and outputs should be a map of name to Parameter
	Inputs []TypedNode `json:"inputs" yaml:"inputs"`
	Output TypedNode   `json:"output" yaml:"output"`
}

func InitMethod(name string) Method {
	return Method{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindMethod,
		},
	}
}

type Interface struct {
	NamedNode  `json:",inline" yaml:",inline"`
	Methods    []Method    `json:"methods" yaml:"methods"`
	Properties []TypedNode `json:"properties" yaml:"properties"`
	Signals    []Signal    `json:"signals" yaml:"signals"`
}

func InitInterface(name string) Interface {
	return Interface{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindInterface,
		},
	}
}

func (i Interface) MethodByName(name string) Method {
	for _, m := range i.Methods {
		if m.Name == name {
			return m
		}
	}
	return Method{}
}

func (i Interface) PropertyByName(name string) TypedNode {
	for _, p := range i.Properties {
		if p.Name == name {
			return p
		}
	}
	return TypedNode{}
}

func (i Interface) SignalByName(name string) Signal {
	for _, s := range i.Signals {
		if s.Name == name {
			return s
		}
	}
	return Signal{}
}
