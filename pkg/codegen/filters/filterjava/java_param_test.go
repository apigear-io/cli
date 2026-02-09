package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "boolean propBool"},
		{"test", "Test1", "propInt", "int propInt"},
		{"test", "Test1", "propInt32", "int propInt32"},
		{"test", "Test1", "propInt64", "long propInt64"},
		{"test", "Test1", "propFloat", "float propFloat"},
		{"test", "Test1", "propFloat32", "float propFloat32"},
		{"test", "Test1", "propFloat64", "double propFloat64"},
		{"test", "Test1", "propString", "String propString"},
		{"test", "Test1", "propBoolArray", "boolean[] propBoolArray"},
		{"test", "Test1", "propIntArray", "int[] propIntArray"},
		{"test", "Test1", "propInt32Array", "int[] propInt32Array"},
		{"test", "Test1", "propInt64Array", "long[] propInt64Array"},
		{"test", "Test1", "propFloatArray", "float[] propFloatArray"},
		{"test", "Test1", "propFloat32Array", "float[] propFloat32Array"},
		{"test", "Test1", "propFloat64Array", "double[] propFloat64Array"},
		{"test", "Test1", "propStringArray", "String[] propStringArray"},
		{"test", "Test1", "prop_Bool", "boolean prop_Bool"},
		{"test", "Test1", "prop_bool", "boolean prop_bool"},
		{"test", "Test1", "prop_1", "boolean prop_1"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1 propEnum"},
		{"test", "Test2", "propStruct", "Struct1 propStruct"},
		{"test", "Test2", "propInterface", "IInterface1 propInterface"},
		{"test", "Test2", "propEnumArray", "Enum1[] propEnumArray"},
		{"test", "Test2", "propStructArray", "Struct1[] propStructArray"},
		{"test", "Test2", "propInterfaceArray", "IInterface1[] propInterfaceArray"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamWithErrors(t *testing.T) {
	s, err := javaParam("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}

func TestExternParam(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "func1", "XType1 arg1"},
		{"demo", "Iface1", "func2", "demo.x.XType2 arg1"},
		{"demo", "Iface1", "func3", "demo.x.XType3A arg1"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := javaParam("", op.Params[0])
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
		{"test_apigear_next", "Iface1", "func1", "XType1 arg1"},
		{"test_apigear_next", "Iface1", "func3", "demo.x.XType3A arg1"},
		{"test_apigear_next", "Iface1", "funcList", "demo.x.XType3A[] arg1"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "test.test_api.Enum1 arg1"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "test.test_api.Struct1 arg1"},
	}
	syss := loadExternSystemsYAML(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := javaParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
