package filtergo

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
		{"test", "Test1", "propVoid", ""},
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int32"},
		{"test", "Test1", "propInt32", "int32"},
		{"test", "Test1", "propInt64", "int64"},
		{"test", "Test1", "propFloat", "float32"},
		{"test", "Test1", "propFloat32", "float32"},
		{"test", "Test1", "propFloat64", "float64"},
		{"test", "Test1", "propString", "string"},
		{"test", "Test1", "propBytes", "[]byte"},
		{"test", "Test1", "propAny", "any"},
		{"test", "Test1", "propBoolArray", "[]bool"},
		{"test", "Test1", "propIntArray", "[]int32"},
		{"test", "Test1", "propInt32Array", "[]int32"},
		{"test", "Test1", "propInt64Array", "[]int64"},
		{"test", "Test1", "propFloatArray", "[]float32"},
		{"test", "Test1", "propFloat32Array", "[]float32"},
		{"test", "Test1", "propFloat64Array", "[]float64"},
		{"test", "Test1", "propStringArray", "[]string"},
		{"test", "Test1", "propBytesArray", "[][]byte"},
		{"test", "Test1", "propAnyArray", "[]any"},
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

func TestOperationReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", ""},
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "int32"},
		{"test", "Test3", "opInt32", "int32"},
		{"test", "Test3", "opInt64", "int64"},
		{"test", "Test3", "opFloat", "float32"},
		{"test", "Test3", "opFloat32", "float32"},
		{"test", "Test3", "opFloat64", "float64"},
		{"test", "Test3", "opString", "string"},
		{"test", "Test3", "opBytes", "[]byte"},
		{"test", "Test3", "opAny", "any"},
		{"test", "Test3", "opBoolArray", "[]bool"},
		{"test", "Test3", "opIntArray", "[]int32"},
		{"test", "Test3", "opInt32Array", "[]int32"},
		{"test", "Test3", "opInt64Array", "[]int64"},
		{"test", "Test3", "opFloatArray", "[]float32"},
		{"test", "Test3", "opFloat32Array", "[]float32"},
		{"test", "Test3", "opFloat64Array", "[]float64"},
		{"test", "Test3", "opStringArray", "[]string"},
		{"test", "Test3", "opBytesArray", "[][]byte"},
		{"test", "Test3", "opAnyArray", "[]any"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := goReturn("", op.Return)
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
		{"test", "Test2", "propEnumArray", "[]Enum1"},
		{"test", "Test2", "propStructArray", "[]Struct1"},
		{"test", "Test2", "propInterfaceArray", "[]Interface1"},
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

func TestExternGoReturn(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "prop1", "XType1"},
		{"demo", "Iface1", "prop2", "x.XType2"},
		{"demo", "Iface1", "prop3", "x.XType3A"},
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

func TestImportedExternGoReturn(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface2", "prop1", "x.XType1"},
		{"demo", "Iface2", "prop2", "x.XType2"},
		{"demo", "Iface2", "prop3", "x.XType3A"},
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
