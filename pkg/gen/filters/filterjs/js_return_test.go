package filterjs

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
		{"test", "Test1", "propBool", ""},
		{"test", "Test1", "propInt", ""},
		{"test", "Test1", "propInt32", ""},
		{"test", "Test1", "propInt64", ""},
		{"test", "Test1", "propFloat", ""},
		{"test", "Test1", "propFloat32", ""},
		{"test", "Test1", "propFloat64", ""},
		{"test", "Test1", "propString", ""},
		{"test", "Test1", "propBoolArray", ""},
		{"test", "Test1", "propIntArray", ""},
		{"test", "Test1", "propInt32Array", ""},
		{"test", "Test1", "propInt64Array", ""},
		{"test", "Test1", "propFloatArray", ""},
		{"test", "Test1", "propFloat32Array", ""},
		{"test", "Test1", "propFloat64Array", ""},
		{"test", "Test1", "propStringArray", ""},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jsReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestOperationReturn(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", ""},
		{"test", "Test3", "opBool", ""},
		{"test", "Test3", "opInt", ""},
		{"test", "Test3", "opFloat", ""},
		{"test", "Test3", "opString", ""},
		{"test", "Test3", "opBoolArray", ""},
		{"test", "Test3", "opIntArray", ""},
		{"test", "Test3", "opFloatArray", ""},
		{"test", "Test3", "opStringArray", ""},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := jsReturn("", op.Return)
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
		{"test", "Test2", "propEnum", ""},
		{"test", "Test2", "propStruct", ""},
		{"test", "Test2", "propInterface", ""},
		{"test", "Test2", "propEnumArray", ""},
		{"test", "Test2", "propStructArray", ""},
		{"test", "Test2", "propInterfaceArray", ""},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jsReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
