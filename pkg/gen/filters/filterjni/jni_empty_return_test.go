package filterjni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TesJniEmptyReturnFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propInt32", "0"},
		{"test", "Test1", "propInt64", "0"},
		{"test", "Test1", "propFloat", "0"},
		{"test", "Test1", "propFloat32", "0"},
		{"test", "Test1", "propFloat64", "0"},
		{"test", "Test1", "propString", "nullptr"},
		{"test", "Test1", "propBoolArray", "nullptr"},
		{"test", "Test1", "propIntArray", "nullptr"},
		{"test", "Test1", "propInt32Array", "nullptr"},
		{"test", "Test1", "propInt64Array", "nullptr"},
		{"test", "Test1", "propFloatArray", "nullptr"},
		{"test", "Test1", "propFloat32Array", "nullptr"},
		{"test", "Test1", "propFloat64Array", "nullptr"},
		{"test", "Test1", "propStringArray", "nullptr"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniEmptyReturn(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniEmptyReturnSymbolsFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn  string
		in  string
		pn  string
		val string
	}{
		// EnumValues: {"test", "Test2", "propEnum", "ETestEnum1::Default"},
		{"test", "Test2", "propEnum", "nullptr"},
		{"test", "Test2", "propStruct", "nullptr"},
		{"test", "Test2", "propInterface", "nullptr"},
		{"test", "Test2", "propEnumArray", "nullptr"},
		{"test", "Test2", "propStructArray", "nullptr"},
		{"test", "Test2", "propInterfaceArray", "nullptr"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniEmptyReturn(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.val, r)
			})
		}
	}
}

func TestImportedExternEmptyReturn(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface2", "prop1", "nullptr"},
		{"demo", "Iface2", "prop2", "nullptr"},
		{"demo", "Iface2", "prop3", "nullptr"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniEmptyReturn(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
func TestEmptyReturnExternsYaml(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "func1", "nullptr"},
		{"test_apigear_next", "Iface1", "func3", "nullptr"},
		{"test_apigear_next", "Iface1", "funcList", "nullptr"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "nullptr"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "nullptr"},
	}
	syss := loadExternSystemsYAML(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := jniEmptyReturn(op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
