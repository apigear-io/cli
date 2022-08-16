package actions

import "github.com/apigear-io/cli/pkg/sim/core"

// ScenarioDoc is a scenario document as part of a simulation.
type ScenarioDoc struct {
	Schema      string            `json:"schema" yaml:"schema"`
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	Version     string            `json:"version" yaml:"version"`
	Interfaces  []*InterfaceEntry `json:"interfaces" yaml:"interfaces"`
	Sequences   []*SequenceEntry  `json:"sequences" yaml:"sequences"`
}

// GetInterface returns the interface entry with the given name.
func (s *ScenarioDoc) GetInterface(name string) *InterfaceEntry {
	for _, iface := range s.Interfaces {
		if iface.Name == name {
			return iface
		}
	}
	return nil
}

// GetSequence returns the sequence entry with the given name.
func (s *ScenarioDoc) GetSequence(name string) *SequenceEntry {
	for _, sequence := range s.Sequences {
		if sequence.Name == name {
			return sequence
		}
	}
	return nil
}

// InterfaceEntry represents an interface in a scenario.
type InterfaceEntry struct {
	Name        string             `json:"name" yaml:"name"`
	Description string             `json:"description" yaml:"description"`
	Properties  map[string]any     `json:"properties" yaml:"properties"`
	Operations  []*ActionListEntry `json:"operations" yaml:"operations"`
}

// GetOperation returns the operation entry with the given name.
func (e InterfaceEntry) GetOperation(name string) *ActionListEntry {
	for _, o := range e.Operations {
		if o.Name == name {
			return o
		}
	}
	return nil
}

// SequenceEntry represents a sequence in a scenario.
type SequenceEntry struct {
	Name             string             `json:"name" yaml:"name"`
	DefaultInterface string             `json:"interface" yaml:"interface"`
	Description      string             `json:"description" yaml:"description"`
	Interval         int                `json:"interval" yaml:"interval"`
	Repeat           int                `json:"repeat" yaml:"repeat"`
	Steps            []*ActionListEntry `json:"steps" yaml:"steps"`
}

// ActionListEntry represents a list of actions
type ActionListEntry struct {
	Name        string        `json:"name" yaml:"name"`
	Description string        `json:"description" yaml:"description"`
	Actions     []ActionEntry `json:"actions" yaml:"actions"`
}

// ActionEntry represents an action in an operation or sequence.
type ActionEntry map[string]core.KWArgs
