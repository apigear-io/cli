package filtercpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeRef(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propInt32", "int32_t"},
		{"test", "Test1", "propInt64", "int64_t"},
		{"test", "Test1", "propFloat", "float"},
		{"test", "Test1", "propFloat32", "float"},
		{"test", "Test1", "propFloat64", "double"},
		{"test", "Test1", "propString", "const std::string&"},
		{"test", "Test1", "propBoolArray", "const std::list<bool>&"},
		{"test", "Test1", "propIntArray", "const std::list<int>&"},
		{"test", "Test1", "propInt32Array", "const std::list<int32_t>&"},
		{"test", "Test1", "propInt64Array", "const std::list<int64_t>&"},
		{"test", "Test1", "propFloatArray", "const std::list<float>&"},
		{"test", "Test1", "propFloat32Array", "const std::list<float>&"},
		{"test", "Test1", "propFloat64Array", "const std::list<double>&"},
		{"test", "Test1", "propStringArray", "const std::list<std::string>&"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppTypeRef("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTypeRefSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1Enum"},
		{"test", "Test2", "propStruct", "const Struct1&"},
		{"test", "Test2", "propInterface", "Interface1*"},
		{"test", "Test2", "propEnumArray", "const std::list<Enum1Enum>&"},
		{"test", "Test2", "propStructArray", "const std::list<Struct1>&"},
		{"test", "Test2", "propInterfaceArray", "const std::list<Interface1*>&"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppTypeRef("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
