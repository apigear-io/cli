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
	err := system.ResolveAll()
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
		{"test", "Test1", "prop1", "bool"},
		{"test", "Test1", "prop2", "int"},
		{"test", "Test1", "prop3", "float64"},
		{"test", "Test1", "prop4", "string"},
		{"test", "Test1", "prop5", "[]bool"},
		{"test", "Test1", "prop6", "[]int"},
		{"test", "Test1", "prop7", "[]float64"},
		{"test", "Test1", "prop8", "[]string"},
	}
	for _, tt := range propTests {
		prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
		assert.NotNil(t, prop)
		r, err := goReturn(reflect.ValueOf(prop))
		assert.NoError(t, err)
		assert.Equal(t, tt.rt, r.String())
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
		{"test", "Test2", "prop1", "Enum1"},
		{"test", "Test2", "prop2", "Struct1"},
		{"test", "Test2", "prop3", "Interface1"},
		{"test", "Test2", "prop4", "[]Enum1"},
		{"test", "Test2", "prop5", "[]Struct1"},
		{"test", "Test2", "prop6", "[]Interface1"},
	}
	for _, tt := range propTests {
		prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
		assert.NotNil(t, prop)
		r, err := goReturn(reflect.ValueOf(prop))
		assert.NoError(t, err)
		assert.Equal(t, tt.rt, r.String())
	}
}
