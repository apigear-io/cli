package idl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManyModules(t *testing.T) {
	s, err := LoadIdlFromFiles("many", []string{"./testdata/simple.idl", "./testdata/data.idl", "./testdata/enum.idl"})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "many", s.Name)
	assert.Equal(t, 3, len(s.Modules))
	assert.Equal(t, "tb.simple", s.Modules[0].Name)
	assert.Equal(t, "tb.data", s.Modules[1].Name)
	assert.Equal(t, "tb.enum", s.Modules[2].Name)
	assert.Equal(t, 2, len(s.Modules[0].Interfaces))
	assert.Equal(t, 2, len(s.Modules[1].Interfaces))
	assert.Equal(t, 1, len(s.Modules[2].Interfaces))
}
