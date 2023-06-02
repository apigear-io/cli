package filterpy

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
		{"test", "Test1", "propInt", "prop_int: int"},
		{"test", "Test1", "propInt32", "prop_int32: int32"},
		{"test", "Test1", "propInt64", "prop_int64: int64"},
		{"test", "Test1", "propFloat", "prop_float: float"},
		{"test", "Test1", "propFloat32", "prop_float32: float32"},
		{"test", "Test1", "propFloat64", "prop_float64: float64"},
		{"test", "Test1", "propString", "prop_string: str"},
		{"test", "Test1", "propBoolArray", "prop_bool_array: list[bool]"},
		{"test", "Test1", "propIntArray", "prop_int_array: list[int]"},
		{"test", "Test1", "propInt32Array", "prop_int32_array: list[int32]"},
		{"test", "Test1", "propInt64Array", "prop_int64_array: list[int64]"},
		{"test", "Test1", "propFloatArray", "prop_float_array: list[float]"},
		{"test", "Test1", "propFloat32Array", "prop_float32_array: list[float32]"},
		{"test", "Test1", "propFloat64Array", "prop_float64_array: list[float64]"},
		{"test", "Test1", "propStringArray", "prop_string_array: list[str]"},
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
				r, err := pyParam("", prop)
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
		{"test", "Test2", "propEnum", "prop_enum: Enum1"},
		{"test", "Test2", "propStruct", "prop_struct: Struct1"},
		{"test", "Test2", "propInterface", "prop_interface: Interface1"},
		{"test", "Test2", "propEnumArray", "prop_enum_array: list[Enum1]"},
		{"test", "Test2", "propStructArray", "prop_struct_array: list[Struct1]"},
		{"test", "Test2", "propInterfaceArray", "prop_interface_array: list[Interface1]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
