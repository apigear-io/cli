package filtercpp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, method inputs, method outputs, signal inputs, struct fields
func TestDefaultFromIdl(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "prop1", "false"},
		{"test", "Test1", "prop2", "0"},
		{"test", "Test1", "prop3", "0.0"},
		{"test", "Test1", "prop4", "std::string()"},
		{"test", "Test1", "prop5", "std::vector<bool>()"},
		{"test", "Test1", "prop6", "std::vector<int>()"},
		{"test", "Test1", "prop7", "std::vector<double>()"},
		{"test", "Test1", "prop8", "std::vector<std::string>()"},
	}
	for _, tt := range propTests {
		prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
		assert.NotNil(t, prop)
		r, err := cppDefault(reflect.ValueOf(prop))
		assert.NoError(t, err)
		assert.Equal(t, tt.rt, r.String())
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "prop1", "Enum1::Default"},
		{"test", "Test2", "prop2", "Struct1()"},
		{"test", "Test2", "prop3", "nullptr"},
		{"test", "Test2", "prop4", "std::vector<Enum1>()"},
		{"test", "Test2", "prop5", "std::vector<Struct1>()"},
		{"test", "Test2", "prop6", "std::vector<Interface1*>()"},
	}
	for _, tt := range propTests {
		prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
		assert.NotNil(t, prop)
		r, err := cppDefault(reflect.ValueOf(prop))
		assert.NoError(t, err)
		assert.Equal(t, tt.rt, r.String())
	}
}
