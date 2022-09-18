package cmd

import "fmt"

type choice struct {
	Allowed []string
	Value   string
}

func NewChoice(allowed []string, d string) *choice {
	return &choice{
		Allowed: allowed,
		Value:   d,
	}
}

func (c choice) String() string {
	return c.Value
}

// Set sets the value of the choice.
func (c *choice) Set(value string) error {
	if !c.IsValid(value) {
		return fmt.Errorf("invalid value %q", value)
	}
	c.Value = value
	return nil
}

// IsValid returns true if the choice is valid.
func (c *choice) IsValid(value string) bool {
	for _, v := range c.Allowed {
		if v == value {
			return true
		}
	}
	return false
}

// Type returns the type of the choice.
func (c *choice) Type() string {
	return "string"
}
