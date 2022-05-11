package filtercpp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, method inputs, method outputs, signal inputs, struct fields
func TestReturn(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propFloat", "double"},
		{"test", "Test1", "propString", "std::string"},
		{"test", "Test1", "propBoolArray", "std::vector<bool>"},
		{"test", "Test1", "propIntArray", "std::vector<int>"},
		{"test", "Test1", "propFloatArray", "std::vector<double>"},
		{"test", "Test1", "propStringArray", "std::vector<std::string>"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := cppReturn(reflect.ValueOf(prop))
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r.String())
		})
	}
}

func TestReturnSymbols(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1"},
		{"test", "Test2", "propStruct", "Struct1"},
		{"test", "Test2", "propInterface", "Interface1*"},
		{"test", "Test2", "propEnumArray", "std::vector<Enum1>"},
		{"test", "Test2", "propStructArray", "std::vector<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "std::vector<Interface1*>"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := cppReturn(reflect.ValueOf(prop))
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r.String())
		})
	}
}
