package filterrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestVar(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test1", "propBool", "prop_bool"},
		{"test", "Test1", "propInt", "prop_int"},
		{"test", "Test1", "propInt32", "prop_int32"},
		{"test", "Test1", "propInt64", "prop_int64"},
		{"test", "Test1", "propFloat", "prop_float"},
		{"test", "Test1", "propFloat32", "prop_float32"},
		{"test", "Test1", "propFloat64", "prop_float64"},
		{"test", "Test1", "propString", "prop_string"},
		{"test", "Test1", "propBoolArray", "prop_bool_array"},
		{"test", "Test1", "propIntArray", "prop_int_array"},
		{"test", "Test1", "propFloatArray", "prop_float_array"},
		{"test", "Test1", "propStringArray", "prop_string_array"},
		{"test", "Test1", "prop_Bool", "prop_bool"},
		{"test", "Test1", "prop_bool", "prop_bool"},
		{"test", "Test1", "prop_1", "prop_1"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rsVar("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestVarSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "prop_enum"},
		{"test", "Test2", "propStruct", "prop_struct"},
		{"test", "Test2", "propInterface", "prop_interface"},
		{"test", "Test2", "propEnumArray", "prop_enum_array"},
		{"test", "Test2", "propStructArray", "prop_struct_array"},
		{"test", "Test2", "propInterfaceArray", "prop_interface_array"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rsVar("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestVarSymbolsPrefixVarName(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "_prop_enum"},
		{"test", "Test2", "propStruct", "_prop_struct"},
		{"test", "Test2", "propInterface", "_prop_interface"},
		{"test", "Test2", "propEnumArray", "_prop_enum_array"},
		{"test", "Test2", "propStructArray", "_prop_struct_array"},
		{"test", "Test2", "propInterfaceArray", "_prop_interface_array"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rsVar("_", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
