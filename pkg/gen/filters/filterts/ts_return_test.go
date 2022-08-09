package filterts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, method inputs, method outputs, signal inputs, struct fields
func TestReturn(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test1", "propBool", "boolean"},
		{"test", "Test1", "propInt", "number"},
		{"test", "Test1", "propFloat", "number"},
		{"test", "Test1", "propString", "string"},
		{"test", "Test1", "propBoolArray", "boolean[]"},
		{"test", "Test1", "propIntArray", "number[]"},
		{"test", "Test1", "propFloatArray", "number[]"},
		{"test", "Test1", "propStringArray", "string[]"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := tsReturn("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestReturnSymbols(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1"},
		{"test", "Test2", "propStruct", "Struct1"},
		{"test", "Test2", "propInterface", "Interface1"},
		{"test", "Test2", "propEnumArray", "Enum1[]"},
		{"test", "Test2", "propStructArray", "Struct1[]"},
		{"test", "Test2", "propInterfaceArray", "Interface1[]"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := tsReturn("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}
