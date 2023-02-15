package filtercpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operations params, operation return, signal params, struct fields
func TestDefaultFromIdl(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propInt32", "0"},
		{"test", "Test1", "propInt64", "0"},
		{"test", "Test1", "propFloat", "0.0f"},
		{"test", "Test1", "propFloat32", "0.0f"},
		{"test", "Test1", "propFloat64", "0.0"},
		{"test", "Test1", "propString", "std::string()"},
		{"test", "Test1", "propBoolArray", "std::list<bool>()"},
		{"test", "Test1", "propIntArray", "std::list<int>()"},
		{"test", "Test1", "propInt32Array", "std::list<int32_t>()"},
		{"test", "Test1", "propInt64Array", "std::list<int64_t>()"},
		{"test", "Test1", "propFloatArray", "std::list<float>()"},
		{"test", "Test1", "propFloat32Array", "std::list<float>()"},
		{"test", "Test1", "propFloat64Array", "std::list<double>()"},
		{"test", "Test1", "propStringArray", "std::list<std::string>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1Enum::Default"},
		{"test", "Test2", "propStruct", "Struct1()"},
		{"test", "Test2", "propInterface", "nullptr"},
		{"test", "Test2", "propEnumArray", "std::list<Enum1Enum>()"},
		{"test", "Test2", "propStructArray", "std::list<Struct1>()"},
		{"test", "Test2", "propInterfaceArray", "std::list<Interface1*>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
