package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFromIdl(t *testing.T) {
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
		{"test", "Test1", "propInt64", "0L"},
		{"test", "Test1", "propFloat", "0.0f"},
		{"test", "Test1", "propFloat32", "0.0f"},
		{"test", "Test1", "propFloat64", "0.0"},
		{"test", "Test1", "propString", "new String()"},
		{"test", "Test1", "propBoolArray", "new boolean[]{}"},
		{"test", "Test1", "propIntArray", "new int[]{}"},
		{"test", "Test1", "propInt32Array", "new int[]{}"},
		{"test", "Test1", "propInt64Array", "new long[]{}"},
		{"test", "Test1", "propFloatArray", "new float[]{}"},
		{"test", "Test1", "propFloat32Array", "new float[]{}"},
		{"test", "Test1", "propFloat64Array", "new double[]{}"},
		{"test", "Test1", "propStringArray", "new String[]{}"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaDefault("", prop)
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
		mn  string
		in  string
		pn  string
		val string
	}{
		// EnumValues: {"test", "Test2", "propEnum", "ETestEnum1::Default"},
		{"test", "Test2", "propEnum", "Enum1.Default"},
		{"test", "Test2", "propStruct", "new Struct1()"},
		{"test", "Test2", "propInterface", "null"},
		{"test", "Test2", "propEnumArray", "new Enum1[]{}"},
		{"test", "Test2", "propStructArray", "new Struct1[]{}"},
		{"test", "Test2", "propInterfaceArray", "new IInterface1[]{}"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.val, r)
			})
		}
	}
}

func TestDefaultWithErrors(t *testing.T) {
	t.Parallel()
	s, err := javaDefault("", nil)
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
		{"demo", "Iface1", "prop1", "new XType1()"},
		{"demo", "Iface1", "prop2", "new demo.x.XType2()"},
		{"demo", "Iface1", "prop3", "new demo.x.XType3A()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaDefault("", prop)
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
		{"test_apigear_next", "Iface1", "prop1", "new XType1()"},
		{"test_apigear_next", "Iface1", "prop2", "new demo.x.XType2A()"},
		{"test_apigear_next", "Iface1", "prop3", "someCtorXType3A()"},
		{"test_apigear_next", "Iface1", "propList", "new demo.x.XType3A[]{}"},
		{"test_apigear_next", "Iface1", "propImportedEnum", "test.test_api.Enum1.Default"},
		{"test_apigear_next", "Iface1", "propImportedStruct", "new test.test_api.Struct1()"},
	}
	syss := loadExternSystemsYAML(t)
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				prop := sys.LookupProperty(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, prop)
				r, err := javaDefault(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
