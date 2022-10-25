package filtergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestReturn(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", ""},
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int64"},
		{"test", "Test1", "propFloat", "float64"},
		{"test", "Test1", "propString", "string"},
		{"test", "Test1", "propBoolArray", "[]bool"},
		{"test", "Test1", "propIntArray", "[]int64"},
		{"test", "Test1", "propFloatArray", "[]float64"},
		{"test", "Test1", "propStringArray", "[]string"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestReturnSymbols(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1"},
		{"test", "Test2", "propStruct", "Struct1"},
		{"test", "Test2", "propInterface", "*Interface1"},
		{"test", "Test2", "propEnumArray", "[]Enum1"},
		{"test", "Test2", "propStructArray", "[]Struct1"},
		{"test", "Test2", "propInterfaceArray", "[]*Interface1"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestReturnWithErrors(t *testing.T) {
	s, err := goReturn("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
