package idl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleIdl(t *testing.T) {
	s, err := LoadIdlFromFiles("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "simple", s.Name)
	assert.Equal(t, 1, len(s.Modules))
	assert.Equal(t, "tb.simple", s.Modules[0].Name)
}
