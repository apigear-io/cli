package spec

import "fmt"

// ScenarioDoc is a scenario document as part of a simulation.
type ScenarioDoc struct {
	Schema      string            `json:"schema" yaml:"schema"`
	Name        string            `json:"name" yaml:"name"`
	Source      string            `json:"source" yaml:"source"`
	Description string            `json:"description" yaml:"description"`
	Version     string            `json:"version" yaml:"version"`
	Interfaces  []*InterfaceEntry `json:"interfaces" yaml:"interfaces"`
	Sequences   []*SequenceEntry  `json:"sequences" yaml:"sequences"`
}

func (d *ScenarioDoc) Validate() error {
	if d.Interfaces == nil {
		d.Interfaces = make([]*InterfaceEntry, 0)
	}
	if d.Sequences == nil {
		d.Sequences = make([]*SequenceEntry, 0)
	}
	for _, iface := range d.Interfaces {
		if err := iface.Validate(); err != nil {
			return err
		}
	}
	for _, sequence := range d.Sequences {
		if err := sequence.Validate(); err != nil {
			return err
		}
	}
	return nil
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

func (e *InterfaceEntry) Validate() error {
	if e.Properties == nil {
		e.Properties = make(map[string]any)
	}
	if e.Operations == nil {
		e.Operations = make([]*ActionListEntry, 0)
	}
	return nil
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
	// Name is the name of the sequence.
	Name string `json:"name" yaml:"name"`
	// Description is the description of the sequence.
	Description string `json:"description" yaml:"description"`
	// Interface is the name of the default interface used.
	Interface string `json:"interface" yaml:"interface"`
	// Interval is the interval in milliseconds between runs.
	Interval int `json:"interval" yaml:"interval"`
	// Loops is the number of times the sequence should be run.
	Loops int `json:"loops" yaml:"loops"`
	// Forever is true if the sequence should be run forever.
	Forever bool `json:"forever" yaml:"forever"`
	// Steps is the list of steps in the sequence.
	Steps []*ActionListEntry `json:"steps" yaml:"steps"`
}

func (e *SequenceEntry) Validate() error {
	if e.Interface == "" {
		return fmt.Errorf("sequence %s: interface is required", e.Name)
	}
	if e.Steps == nil {
		e.Steps = make([]*ActionListEntry, 0)
	}
	return nil
}

// ActionListEntry represents a list of actions
type ActionListEntry struct {
	// Name is the name of the action list.
	Name string `json:"name" yaml:"name"`
	// Description is the description of the action list.
	Description string `json:"description" yaml:"description"`
	// Actions is the list of actions in the action list.
	Actions []ActionEntry `json:"actions" yaml:"actions"`
}

// ActionEntry represents an action in an operation or sequence.
type ActionEntry map[string]map[string]any
