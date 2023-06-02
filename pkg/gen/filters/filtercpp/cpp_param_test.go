package filtercpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "bool propBool"},
		{"test", "Test1", "propInt", "int propInt"},
		{"test", "Test1", "propInt32", "int32_t propInt32"},
		{"test", "Test1", "propInt64", "int64_t propInt64"},
		{"test", "Test1", "propFloat", "float propFloat"},
		{"test", "Test1", "propFloat32", "float propFloat32"},
		{"test", "Test1", "propFloat64", "double propFloat64"},
		{"test", "Test1", "propString", "const std::string& propString"},
		{"test", "Test1", "propBoolArray", "const std::list<bool>& propBoolArray"},
		{"test", "Test1", "propIntArray", "const std::list<int>& propIntArray"},
		{"test", "Test1", "propFloatArray", "const std::list<float>& propFloatArray"},
		{"test", "Test1", "propStringArray", "const std::list<std::string>& propStringArray"},
		{"test", "Test1", "prop_Bool", "bool prop_Bool"},
		{"test", "Test1", "prop_bool", "bool prop_bool"},
		{"test", "Test1", "prop_1", "bool prop_1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1Enum propEnum"},
		{"test", "Test2", "propStruct", "const Struct1& propStruct"},
		{"test", "Test2", "propInterface", "Interface1* propInterface"},
		{"test", "Test2", "propEnumArray", "const std::list<Enum1Enum>& propEnumArray"},
		{"test", "Test2", "propStructArray", "const std::list<Struct1>& propStructArray"},
		{"test", "Test2", "propInterfaceArray", "const std::list<Interface1*>& propInterfaceArray"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
