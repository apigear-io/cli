package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaImport(t *testing.T) {
	s := readSystem(t)
	assert.NotNil(t, s)
	m := s.LookupModule("b")
	assert.NotNil(t, m)
	assert.Equal(t, 1, len(m.Imports))
	assert.Equal(t, "a", m.Imports[0].Name)

	i := m.LookupInterface("b", "B")
	assert.NotNil(t, i)

	p := i.LookupProperty("value")
	assert.NotNil(t, p)
	assert.Equal(t, "A", p.Type)
	assert.Equal(t, "a", p.Import)

	s2 := m.LookupStruct("a", "A")
	assert.NotNil(t, s2)

}
