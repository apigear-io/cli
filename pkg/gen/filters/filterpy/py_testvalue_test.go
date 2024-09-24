package filterpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestTestValueFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "None"},
		{"test", "Test1", "propBool", "True"},
		{"test", "Test1", "propInt", "1"},
		{"test", "Test1", "propInt32", "1"},
		{"test", "Test1", "propInt64", "1"},
		{"test", "Test1", "propFloat", "1.1"},
		{"test", "Test1", "propFloat32", "1.1"},
		{"test", "Test1", "propFloat64", "1.1"},
		{"test", "Test1", "propString", "\"xyz\""},
		{"test", "Test1", "propBoolArray", "True"}, // all the array types return value intentionally, it may be put into empty array
		{"test", "Test1", "propIntArray", "1"},
		{"test", "Test1", "propInt32Array", "1"},
		{"test", "Test1", "propInt64Array", "1"},
		{"test", "Test1", "propFloatArray", "1.1"},
		{"test", "Test1", "propFloat32Array", "1.1"},
		{"test", "Test1", "propFloat64Array", "1.1"},
		{"test", "Test1", "propStringArray", "\"xyz\""},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueSymbolsFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1.NOT_DEFAULT"},
		{"test", "Test2", "propStruct", "Struct1()"},
		{"test", "Test2", "propInterface", "Interface1()"},
		{"test", "Test2", "propEnumArray", "Enum1.NOT_DEFAULT"},
		{"test", "Test2", "propStructArray", "Struct1()"},
		{"test", "Test2", "propInterfaceArray", "Interface1()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueWithErrors(t *testing.T) {
	t.Parallel()
	s, err := pyTestValue("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}

func TestTestValueReturnsExternDefault(t *testing.T) {
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
				r, err := pyTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueReturnsDefaultExterns(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "prop1", "XType1()"},
		{"test_apigear_next", "Iface1", "prop2", "demo.x.XType2()"},
		{"test_apigear_next", "Iface1", "prop3", "demo.x.createXType3A()"},
		{"test_apigear_next", "Iface1", "propList", "demo.x.createXType3A()"},
		{"test_apigear_next", "Iface1", "propImportedEnum", "test.api.Enum1.NOT_DEFAULT"},
		{"test_apigear_next", "Iface1", "propImportedStruct", "test.api.Struct1()"},
	}
	syss := loadExternSystemsYAML(t)
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				prop := sys.LookupProperty(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, prop)
				r, err := pyTestValue(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
