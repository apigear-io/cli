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
		{"test", "Test1", "propFloat", "double propFloat"},
		{"test", "Test1", "propString", "const std::string &propString"},
		{"test", "Test1", "propBoolArray", "const std::vector<bool> &propBoolArray"},
		{"test", "Test1", "propIntArray", "const std::vector<int> &propIntArray"},
		{"test", "Test1", "propFloatArray", "const std::vector<double> &propFloatArray"},
		{"test", "Test1", "propStringArray", "const std::vector<std::string> &propStringArray"},
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
		{"test", "Test2", "propEnum", "Enum1 propEnum"},
		{"test", "Test2", "propStruct", "const Struct1 &propStruct"},
		{"test", "Test2", "propInterface", "Interface1 *propInterface"},
		{"test", "Test2", "propEnumArray", "const std::vector<Enum1> &propEnumArray"},
		{"test", "Test2", "propStructArray", "const std::vector<Struct1> &propStructArray"},
		{"test", "Test2", "propInterfaceArray", "const std::vector<Interface1*> &propInterfaceArray"},
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
