package filterrust

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
		{"test", "Test1", "propVoid", "()"},
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "i32"},
		{"test", "Test1", "propInt32", "i32"},
		{"test", "Test1", "propInt64", "i64"},
		{"test", "Test1", "propFloat", "f32"},
		{"test", "Test1", "propFloat32", "f32"},
		{"test", "Test1", "propFloat64", "f64"},
		{"test", "Test1", "propString", "String"},
		{"test", "Test1", "propBoolArray", "Vec<bool>"},
		{"test", "Test1", "propIntArray", "Vec<i32>"},
		{"test", "Test1", "propInt32Array", "Vec<i32>"},
		{"test", "Test1", "propInt64Array", "Vec<i64>"},
		{"test", "Test1", "propFloatArray", "Vec<f32>"},
		{"test", "Test1", "propFloat32Array", "Vec<f32>"},
		{"test", "Test1", "propFloat64Array", "Vec<f64>"},
		{"test", "Test1", "propStringArray", "Vec<String>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rustReturn("", prop)
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
		{"test", "Test3", "opVoid", "()"},
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "i32"},
		{"test", "Test3", "opInt32", "i32"},
		{"test", "Test3", "opInt64", "i64"},
		{"test", "Test3", "opFloat", "f32"},
		{"test", "Test3", "opFloat32", "f32"},
		{"test", "Test3", "opFloat64", "f64"},
		{"test", "Test3", "opString", "String"},
		{"test", "Test3", "opBoolArray", "Vec<bool>"},
		{"test", "Test3", "opIntArray", "Vec<i32>"},
		{"test", "Test3", "opFloatArray", "Vec<f32>"},
		{"test", "Test3", "opStringArray", "Vec<String>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := rustReturn("", op.Return)
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
		{"test", "Test2", "propEnum", "Enum1Enum"},
		{"test", "Test2", "propStruct", "Struct1"},
		{"test", "Test2", "propInterface", "&Interface1"},
		{"test", "Test2", "propEnumArray", "Vec<Enum1Enum>"},
		{"test", "Test2", "propStructArray", "Vec<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "Vec<&Interface1>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rustReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
