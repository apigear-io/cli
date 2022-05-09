package filtercpp

import (
	"reflect"
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
		{"test", "Test1", "prop1", "bool prop1"},
		{"test", "Test1", "prop2", "int prop2"},
		{"test", "Test1", "prop3", "double prop3"},
		{"test", "Test1", "prop4", "const std::string &prop4"},
		{"test", "Test1", "prop5", "const std::vector<bool> &prop5"},
		{"test", "Test1", "prop6", "const std::vector<int> &prop6"},
		{"test", "Test1", "prop7", "const std::vector<double> &prop7"},
		{"test", "Test1", "prop8", "const std::vector<std::string> &prop8"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
		assert.NotNil(t, prop)
		r, err := cppParam(reflect.ValueOf(prop))
		assert.NoError(t, err)
		assert.Equal(t, tt.rt, r.String())
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "prop1", "Enum1 prop1"},
		{"test", "Test2", "prop2", "const Struct1 &prop2"},
		{"test", "Test2", "prop3", "Interface1 *prop3"},
		{"test", "Test2", "prop4", "const std::vector<Enum1> &prop4"},
		{"test", "Test2", "prop5", "const std::vector<Struct1> &prop5"},
		{"test", "Test2", "prop6", "const std::vector<Interface1*> &prop6"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
		assert.NotNil(t, prop)
		r, err := cppParam(reflect.ValueOf(prop))
		assert.NoError(t, err)
		assert.Equal(t, tt.rt, r.String())
	}
}
