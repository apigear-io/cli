package filterpy

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
		{"test", "Test1", "propVoid", "None"},
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propInt32", "int32"},
		{"test", "Test1", "propInt64", "int64"},
		{"test", "Test1", "propFloat", "float"},
		{"test", "Test1", "propFloat32", "float32"},
		{"test", "Test1", "propFloat64", "float64"},
		{"test", "Test1", "propString", "str"},
		{"test", "Test1", "propBoolArray", "list[bool]"},
		{"test", "Test1", "propIntArray", "list[int]"},
		{"test", "Test1", "propInt32Array", "list[int32]"},
		{"test", "Test1", "propInt64Array", "list[int64]"},
		{"test", "Test1", "propFloatArray", "list[float]"},
		{"test", "Test1", "propFloat32Array", "list[float32]"},
		{"test", "Test1", "propFloat64Array", "list[float64]"},
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

func TestOperationReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "None"},
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "int"},
		{"test", "Test3", "opInt32", "int32"},
		{"test", "Test3", "opInt64", "int64"},
		{"test", "Test3", "opFloat", "float"},
		{"test", "Test3", "opFloat32", "float32"},
		{"test", "Test3", "opFloat64", "float64"},
		{"test", "Test3", "opString", "str"},
		{"test", "Test3", "opBoolArray", "list[bool]"},
		{"test", "Test3", "opIntArray", "list[int]"},
		{"test", "Test3", "opInt32Array", "list[int32]"},
		{"test", "Test3", "opInt64Array", "list[int64]"},
		{"test", "Test3", "opFloatArray", "list[float]"},
		{"test", "Test3", "opFloat32Array", "list[float32]"},
		{"test", "Test3", "opFloat64Array", "list[float64]"},
		{"test", "Test3", "opStringArray", "list[str]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := pyReturn("", op.Return)
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

func TestImportedExternReturn(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface2", "prop1", "XType1"},
		{"demo", "Iface2", "prop2", "XType2"},
		{"demo", "Iface2", "prop3", "XType3A"},
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
