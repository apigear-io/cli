package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoValue(t *testing.T) {
	e := NewEnum("foo")
	e.Members = []*EnumMember{
		NewEnumMember("a", 0),
		NewEnumMember("b", 0),
		NewEnumMember("c", 0),
	}
	err := e.Validate(nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, e.Members[0].Value)
	assert.Equal(t, 1, e.Members[1].Value)
	assert.Equal(t, 2, e.Members[2].Value)
}

func TestNoAutoValue(t *testing.T) {
	e := NewEnum("foo")
	e.Members = []*EnumMember{
		NewEnumMember("a", 0),
		NewEnumMember("b", 2),
		NewEnumMember("c", 1),
	}
	err := e.Validate(nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, e.Members[0].Value)
	assert.Equal(t, 2, e.Members[1].Value)
	assert.Equal(t, 1, e.Members[2].Value)
}
