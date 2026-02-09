package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "boolean"},
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propInt32", "int"},
		{"test", "Test1", "propInt64", "long"},
		{"test", "Test1", "propFloat", "float"},
		{"test", "Test1", "propFloat32", "float"},
		{"test", "Test1", "propFloat64", "double"},
		{"test", "Test1", "propString", "String"},
		{"test", "Test1", "propBoolArray", "boolean[]"},
		{"test", "Test1", "propIntArray", "int[]"},
		{"test", "Test1", "propInt32Array", "int[]"},
		{"test", "Test1", "propInt64Array", "long[]"},
		{"test", "Test1", "propFloatArray", "float[]"},
		{"test", "Test1", "propFloat32Array", "float[]"},
		{"test", "Test1", "propFloat64Array", "double[]"},
		{"test", "Test1", "propStringArray", "String[]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestOperationReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "boolean"},
		{"test", "Test3", "opInt", "int"},
		{"test", "Test3", "opInt32", "int"},
		{"test", "Test3", "opInt64", "long"},
		{"test", "Test3", "opFloat", "float"},
		{"test", "Test3", "opFloat32", "float"},
		{"test", "Test3", "opFloat64", "double"},
		{"test", "Test3", "opString", "String"},
		{"test", "Test3", "opBoolArray", "boolean[]"},
		{"test", "Test3", "opIntArray", "int[]"},
		{"test", "Test3", "opInt32Array", "int[]"},
		{"test", "Test3", "opInt64Array", "long[]"},
		{"test", "Test3", "opFloatArray", "float[]"},
		{"test", "Test3", "opFloat32Array", "float[]"},
		{"test", "Test3", "opFloat64Array", "double[]"},
		{"test", "Test3", "opStringArray", "String[]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := javaReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestImportedExternReturn(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface2", "prop1", "XType1"},
		{"demo", "Iface2", "prop2", "demo.x.XType2"},
		{"demo", "Iface2", "prop3", "demo.x.XType3A"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
func TestReturnExternsYaml(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "func1", "XType1"},
		{"test_apigear_next", "Iface1", "func3", "demo.x.XType3A"},
		{"test_apigear_next", "Iface1", "funcList", "demo.x.XType3A[]"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "test.test_api.Enum1"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "test.test_api.Struct1"},
	}
	syss := loadExternSystemsYAML(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := javaReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
