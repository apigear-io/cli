package filterrust

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
		{"test", "Test1", "propBool", "prop_bool: bool"},
		{"test", "Test1", "propInt", "prop_int: i32"},
		{"test", "Test1", "propInt32", "prop_int32: i32"},
		{"test", "Test1", "propInt64", "prop_int64: i64"},
		{"test", "Test1", "propFloat", "prop_float: f32"},
		{"test", "Test1", "propFloat32", "prop_float32: f32"},
		{"test", "Test1", "propFloat64", "prop_float64: f64"},
		{"test", "Test1", "propString", "prop_string: &str"},
		{"test", "Test1", "propBoolArray", "prop_bool_array: &Vec<bool>"},
		{"test", "Test1", "propIntArray", "prop_int_array: &Vec<i32>"},
		{"test", "Test1", "propFloatArray", "prop_float_array: &Vec<f32>"},
		{"test", "Test1", "propStringArray", "prop_string_array: &Vec<String>"},
		{"test", "Test1", "prop_Bool", "prop_bool: bool"},
		{"test", "Test1", "prop_bool", "prop_bool: bool"},
		{"test", "Test1", "prop_1", "prop_1: bool"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rustParam("", "", prop)
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
		{"test", "Test2", "propEnum", "prop_enum: Enum1Enum"},
		{"test", "Test2", "propStruct", "prop_struct: &Struct1"},
		{"test", "Test2", "propInterface", "prop_interface: &Interface1"},
		{"test", "Test2", "propEnumArray", "prop_enum_array: &Vec<Enum1Enum>"},
		{"test", "Test2", "propStructArray", "prop_struct_array: &Vec<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "prop_interface_array: &Vec<&Interface1>"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rustParam("", "", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
