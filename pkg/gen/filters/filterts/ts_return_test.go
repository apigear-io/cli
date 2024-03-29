package filterts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "boolean"},
		{"test", "Test1", "propInt", "number"},
		{"test", "Test1", "propInt32", "number"},
		{"test", "Test1", "propInt64", "number"},
		{"test", "Test1", "propFloat", "number"},
		{"test", "Test1", "propFloat32", "number"},
		{"test", "Test1", "propFloat64", "number"},
		{"test", "Test1", "propString", "string"},
		{"test", "Test1", "propBoolArray", "boolean[]"},
		{"test", "Test1", "propIntArray", "number[]"},
		{"test", "Test1", "propInt32Array", "number[]"},
		{"test", "Test1", "propInt64Array", "number[]"},
		{"test", "Test1", "propFloatArray", "number[]"},
		{"test", "Test1", "propFloat32Array", "number[]"},
		{"test", "Test1", "propFloat64Array", "number[]"},
		{"test", "Test1", "propStringArray", "string[]"},
	}
	for _, sys := range syss {
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
}

func TestOperationReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "boolean"},
		{"test", "Test3", "opInt", "number"},
		{"test", "Test3", "opFloat", "number"},
		{"test", "Test3", "opString", "string"},
		{"test", "Test3", "opBoolArray", "boolean[]"},
		{"test", "Test3", "opIntArray", "number[]"},
		{"test", "Test3", "opFloatArray", "number[]"},
		{"test", "Test3", "opStringArray", "string[]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := tsReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
func TestReturnSymbols(t *testing.T) {
	t.Parallel()
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
		{"test", "Test2", "propEnumArray", "Enum1[]"},
		{"test", "Test2", "propStructArray", "Struct1[]"},
		{"test", "Test2", "propInterfaceArray", "Interface1[]"},
	}
	for _, sys := range syss {
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
}
