package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListReturn(t *testing.T) {
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
		{"test", "Test1", "propBoolArray", "List<Boolean>"},
		{"test", "Test1", "propIntArray", "List<Integer>"},
		{"test", "Test1", "propInt32Array", "List<Integer>"},
		{"test", "Test1", "propInt64Array", "List<Long>"},
		{"test", "Test1", "propFloatArray", "List<Float>"},
		{"test", "Test1", "propFloat32Array", "List<Float>"},
		{"test", "Test1", "propFloat64Array", "List<Double>"},
		{"test", "Test1", "propStringArray", "List<String>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaListReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestListOperationReturn(t *testing.T) {
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
		{"test", "Test3", "opBoolArray", "List<Boolean>"},
		{"test", "Test3", "opIntArray", "List<Integer>"},
		{"test", "Test3", "opInt32Array", "List<Integer>"},
		{"test", "Test3", "opInt64Array", "List<Long>"},
		{"test", "Test3", "opFloatArray", "List<Float>"},
		{"test", "Test3", "opFloat32Array", "List<Float>"},
		{"test", "Test3", "opFloat64Array", "List<Double>"},
		{"test", "Test3", "opStringArray", "List<String>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := javaListReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestListReturnSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1"},
		{"test", "Test2", "propStruct", "Struct1"},
		{"test", "Test2", "propInterface", "IInterface1"},
		{"test", "Test2", "propEnumArray", "List<Enum1>"},
		{"test", "Test2", "propStructArray", "List<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "List<IInterface1>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaListReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestListReturnExternsYaml(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "func1", "XType1"},
		{"test_apigear_next", "Iface1", "func3", "demo.x.XType3A"},
		{"test_apigear_next", "Iface1", "funcList", "List<demo.x.XType3A>"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "test.test_api.Enum1"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "test.test_api.Struct1"},
	}
	syss := loadExternSystemsYAML(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := javaListReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
