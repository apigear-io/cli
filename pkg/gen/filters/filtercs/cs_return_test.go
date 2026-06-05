package filtercs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the primitive and array types
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
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propInt32", "int"},
		{"test", "Test1", "propInt64", "long"},
		{"test", "Test1", "propFloat", "float"},
		{"test", "Test1", "propFloat32", "float"},
		{"test", "Test1", "propFloat64", "double"},
		{"test", "Test1", "propString", "string"},
		{"test", "Test1", "propBytes", "byte[]"},
		{"test", "Test1", "propAny", "object"},
		{"test", "Test1", "propBoolArray", "List<bool>"},
		{"test", "Test1", "propIntArray", "List<int>"},
		{"test", "Test1", "propInt32Array", "List<int>"},
		{"test", "Test1", "propInt64Array", "List<long>"},
		{"test", "Test1", "propFloatArray", "List<float>"},
		{"test", "Test1", "propFloat32Array", "List<float>"},
		{"test", "Test1", "propFloat64Array", "List<double>"},
		{"test", "Test1", "propStringArray", "List<string>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csReturn("", prop)
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
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "int"},
		{"test", "Test3", "opInt32", "int"},
		{"test", "Test3", "opInt64", "long"},
		{"test", "Test3", "opFloat", "float"},
		{"test", "Test3", "opFloat32", "float"},
		{"test", "Test3", "opFloat64", "double"},
		{"test", "Test3", "opString", "string"},
		{"test", "Test3", "opBoolArray", "List<bool>"},
		{"test", "Test3", "opIntArray", "List<int>"},
		{"test", "Test3", "opFloatArray", "List<float>"},
		{"test", "Test3", "opStringArray", "List<string>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := csReturn("", op.Return)
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
		{"test", "Test2", "propInterface", "IInterface1"},
		{"test", "Test2", "propEnumArray", "List<Enum1>"},
		{"test", "Test2", "propStructArray", "List<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "List<IInterface1>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestPrefixedReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propString", "string"},
		{"test", "Test1", "propIntArray", "List<int>"},
		{"test", "Test2", "propEnum", "MyPrefix.Enum1"},
		{"test", "Test2", "propStruct", "MyPrefix.Struct1"},
		{"test", "Test2", "propInterface", "MyPrefix.IInterface1"},
		{"test", "Test2", "propEnumArray", "List<MyPrefix.Enum1>"},
		{"test", "Test2", "propStructArray", "List<MyPrefix.Struct1>"},
		{"test", "Test2", "propInterfaceArray", "List<MyPrefix.IInterface1>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csReturn("MyPrefix.", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestType(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	for _, sys := range syss {
		prop := sys.LookupProperty("test", "Test2", "propStruct")
		assert.NotNil(t, prop)
		r, err := csType("", prop)
		assert.NoError(t, err)
		assert.Equal(t, "Struct1", r)
	}
}
