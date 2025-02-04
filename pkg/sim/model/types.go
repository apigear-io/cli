package model

import "time"

type Script struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

// IsEmpty returns true if the script is empty
func (s Script) IsEmpty() bool {
	return s.Name == "" && s.Source == ""
}

type WorldStatus struct {
	Name       string    `json:"name"`
	IsActive   bool      `json:"active"`
	ActorCount int       `json:"actorCount"`
	LastUpdate time.Time `json:"lastUpdate"`
}
