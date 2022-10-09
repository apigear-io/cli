package filterpy

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

		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propFloat", "float"},
		{"test", "Test1", "propString", "str"},
		{"test", "Test1", "propBoolArray", "list[bool]"},
		{"test", "Test1", "propIntArray", "list[int]"},
		{"test", "Test1", "propFloatArray", "list[float]"},
		{"test", "Test1", "propStringArray", "list[str]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyReturn("", prop)
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
		{"test", "Test2", "propInterface", "Interface1"},
		{"test", "Test2", "propEnumArray", "list[Enum1]"},
		{"test", "Test2", "propStructArray", "list[Struct1]"},
		{"test", "Test2", "propInterfaceArray", "list[Interface1]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
