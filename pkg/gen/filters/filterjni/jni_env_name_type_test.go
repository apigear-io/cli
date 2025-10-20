package filterjni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestJniToEnvNameType(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "Boolean"},
		{"test", "Test3", "opInt", "Int"},
		{"test", "Test3", "opInt32", "Int"},
		{"test", "Test3", "opInt64", "Long"},
		{"test", "Test3", "opFloat", "Float"},
		{"test", "Test3", "opFloat32", "Float"},
		{"test", "Test3", "opFloat64", "Double"},
		{"test", "Test3", "opString", "Object"},
		{"test", "Test3", "opBoolArray", "Boolean"},
		{"test", "Test3", "opIntArray", "Int"},
		{"test", "Test3", "opInt32Array", "Int"},
		{"test", "Test3", "opInt64Array", "Long"},
		{"test", "Test3", "opFloatArray", "Float"},
		{"test", "Test3", "opFloat32Array", "Float"},
		{"test", "Test3", "opFloat64Array", "Double"},
		{"test", "Test3", "opStringArray", "Object"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := jniToEnvNameType(op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJnEnvNameTypeOperationReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "Boolean"},
		{"test", "Test3", "opInt", "Int"},
		{"test", "Test3", "opInt32", "Int"},
		{"test", "Test3", "opInt64", "Long"},
		{"test", "Test3", "opFloat", "Float"},
		{"test", "Test3", "opFloat32", "Float"},
		{"test", "Test3", "opFloat64", "Double"},
		{"test", "Test3", "opString", "Object"},
		{"test", "Test3", "opBoolArray", "Boolean"},
		{"test", "Test3", "opIntArray", "Int"},
		{"test", "Test3", "opInt32Array", "Int"},
		{"test", "Test3", "opInt64Array", "Long"},
		{"test", "Test3", "opFloatArray", "Float"},
		{"test", "Test3", "opFloat32Array", "Float"},
		{"test", "Test3", "opFloat64Array", "Double"},
		{"test", "Test3", "opStringArray", "Object"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := jniToEnvNameType(op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniToEnvNameTypeSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Object"},
		{"test", "Test2", "propStruct", "Object"},
		{"test", "Test2", "propInterface", "Object"},
		{"test", "Test2", "propEnumArray", "Object"},
		{"test", "Test2", "propStructArray", "Object"},
		{"test", "Test2", "propInterfaceArray", "Object"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniToEnvNameType(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestImportedExternJniToEnvNameType(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface2", "prop1", "Object"},
		{"demo", "Iface2", "prop2", "Object"},
		{"demo", "Iface2", "prop3", "Object"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniToEnvNameType(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
func TestJniToEnvNameTypeExternsYaml(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "func1", "Object"},
		{"test_apigear_next", "Iface1", "func3", "Object"},
		{"test_apigear_next", "Iface1", "funcList", "Object"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "Object"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "Object"},
	}
	syss := loadExternSystemsYAML(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := jniToEnvNameType(op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
