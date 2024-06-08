package filterpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "prop_bool: bool"},
		{"test", "Test1", "propInt", "prop_int: int"},
		{"test", "Test1", "propInt32", "prop_int32: int"},
		{"test", "Test1", "propInt64", "prop_int64: int"},
		{"test", "Test1", "propFloat", "prop_float: float"},
		{"test", "Test1", "propFloat32", "prop_float32: float"},
		{"test", "Test1", "propFloat64", "prop_float64: float"},
		{"test", "Test1", "propString", "prop_string: str"},
		{"test", "Test1", "propBoolArray", "prop_bool_array: list[bool]"},
		{"test", "Test1", "propIntArray", "prop_int_array: list[int]"},
		{"test", "Test1", "propInt32Array", "prop_int32_array: list[int]"},
		{"test", "Test1", "propInt64Array", "prop_int64_array: list[int]"},
		{"test", "Test1", "propFloatArray", "prop_float_array: list[float]"},
		{"test", "Test1", "propFloat32Array", "prop_float32_array: list[float]"},
		{"test", "Test1", "propFloat64Array", "prop_float64_array: list[float]"},
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
	t.Parallel()
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

func TestExternParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "func1", "self, arg1: XType1"},
		{"demo", "Iface1", "func2", "self, arg1: XType2"},
		{"demo", "Iface1", "func3", "self, arg1: XType3A"},
	}
	syss := loadExternSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := pyParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsExterns(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "func1", "self, arg1: XType1"},
		{"test_apigear_next", "Iface1", "func3", "self, arg1: demo.x.XType3A"},
		{"test_apigear_next", "Iface1", "funcList", "self, arg1: list[demo.x.XType3A]"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "self, arg1: test.api.Enum1"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "self, arg1: test.api.Struct1"},
	}
	syss := loadExternSystemsYAML(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := pyParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
