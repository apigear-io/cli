package filtergo

import (
	"objectapi/pkg/idl"
	"objectapi/pkg/model"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadSystem(t *testing.T) *model.System {
	system := model.NewSystem("test")
	p := idl.NewParser(system)
	err := p.ParseFile("testdata/test.idl")
	assert.NoError(t, err)
	err = system.ResolveAll()
	assert.NoError(t, err)
	return system
}

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
		{"test", "Test1", "propFloat", "float64"},
		{"test", "Test1", "propString", "string"},
		{"test", "Test1", "propBoolArray", "[]bool"},
		{"test", "Test1", "propIntArray", "[]int"},
		{"test", "Test1", "propFloatArray", "[]float64"},
		{"test", "Test1", "propStringArray", "[]string"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := goReturn(reflect.ValueOf(prop))
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
		{"test", "Test2", "propInterface", "*Interface1"},
		{"test", "Test2", "propEnumArray", "[]Enum1"},
		{"test", "Test2", "propStructArray", "[]Struct1"},
		{"test", "Test2", "propInterfaceArray", "[]*Interface1"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := goReturn(reflect.ValueOf(prop))
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r.String())
		})
	}
}
