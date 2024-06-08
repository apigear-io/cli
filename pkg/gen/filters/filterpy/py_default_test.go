package filterpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestDefaultFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "None"},
		{"test", "Test1", "propBool", "False"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propInt32", "0"},
		{"test", "Test1", "propInt64", "0"},
		{"test", "Test1", "propFloat", "0.0"},
		{"test", "Test1", "propFloat32", "0.0"},
		{"test", "Test1", "propFloat64", "0.0"},
		{"test", "Test1", "propString", "\"\""},
		{"test", "Test1", "propBoolArray", "[]"},
		{"test", "Test1", "propIntArray", "[]"},
		{"test", "Test1", "propInt32Array", "[]"},
		{"test", "Test1", "propInt64Array", "[]"},
		{"test", "Test1", "propFloatArray", "[]"},
		{"test", "Test1", "propFloat32Array", "[]"},
		{"test", "Test1", "propFloat64Array", "[]"},
		{"test", "Test1", "propStringArray", "[]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1.DEFAULT"},
		{"test", "Test2", "propStruct", "Struct1()"},
		{"test", "Test2", "propInterface", "None"},
		{"test", "Test2", "propEnumArray", "[]"},
		{"test", "Test2", "propStructArray", "[]"},
		{"test", "Test2", "propInterfaceArray", "[]"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultWithErrors(t *testing.T) {
	t.Parallel()
	s, err := pyDefault("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}

func TestExternDefault(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "prop1", "XType1()"},
		{"demo", "Iface1", "prop2", "XType2()"},
		{"demo", "Iface1", "prop3", "XType3A()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultExterns(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "prop1", "XType1()"},
		{"test_apigear_next", "Iface1", "prop3", "demo.x.XType3A()"},
		{"test_apigear_next", "Iface1", "propList", "[]"},
		{"test_apigear_next", "Iface1", "propImportedEnum", "test.api.Enum1.DEFAULT"},
		{"test_apigear_next", "Iface1", "propImportedStruct", "test.api.Struct1()"},
	}
	syss := loadExternSystemsYAML(t)
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				prop := sys.LookupProperty(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, prop)
				r, err := pyDefault(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
