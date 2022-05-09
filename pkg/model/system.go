package model

import "objectapi/pkg/log"

type System struct {
	NamedNode `json:",inline" yaml:",inline"`
	Modules   []*Module `json:"modules" yaml:"modules"`
}

func NewSystem(name string) *System {
	return &System{
		NamedNode: NamedNode{
			Name: name,
			Kind: KindSystem,
		},
		Modules: make([]*Module, 0),
	}
}

func (s System) LookupModule(name string) *Module {
	for _, m := range s.Modules {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (s System) LookupInterface(moduleName string, ifaceName string) *Interface {
	m := s.LookupModule(moduleName)
	if m == nil {
		return nil
	}
	return m.LookupInterface(ifaceName)
}

func (s System) LookupStruct(moduleName string, structName string) *Struct {
	m := s.LookupModule(moduleName)
	if m == nil {
		return nil
	}
	return m.LookupStruct(structName)
}

func (s System) LookupEnum(moduleName string, enumName string) *Enum {
	m := s.LookupModule(moduleName)
	if m == nil {
		return nil
	}
	return m.LookupEnum(enumName)
}

func (s System) LookupProperty(moduleName string, ifaceName string, propName string) *TypedNode {
	i := s.LookupInterface(moduleName, ifaceName)
	if i == nil {
		return nil
	}
	return i.LookupProperty(propName)
}

func (s System) LookupMethod(moduleName string, ifaceName string, methodName string) *Method {
	i := s.LookupInterface(moduleName, ifaceName)
	if i == nil {
		return nil
	}
	return i.LookupMethod(methodName)
}

func (s System) LookupSignal(moduleName string, ifaceName string, eventName string) *Signal {
	i := s.LookupInterface(moduleName, ifaceName)
	if i == nil {
		return nil
	}
	return i.LookupSignal(eventName)
}

func (s *System) ResolveAll() error {
	log.Infof("Resolving system %s", s.Name)
	for _, m := range s.Modules {
		err := m.ResolveAll()
		if err != nil {
			return err
		}
	}
	return nil
}
