package filtercpp

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
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "true"},
		{"test", "Test1", "propInt", "1"},
		{"test", "Test1", "propInt32", "1"},
		{"test", "Test1", "propInt64", "1LL"},
		{"test", "Test1", "propFloat", "1.1f"},
		{"test", "Test1", "propFloat32", "1.1f"},
		{"test", "Test1", "propFloat64", "1.1"},
		{"test", "Test1", "propString", "std::string(\"xyz\")"},
		{"test", "Test1", "propBoolArray", "true"}, // all the array types return value intentionally, it may be put into empty array
		{"test", "Test1", "propIntArray", "1"},
		{"test", "Test1", "propInt32Array", "1"},
		{"test", "Test1", "propInt64Array", "1LL"},
		{"test", "Test1", "propFloatArray", "1.1f"},
		{"test", "Test1", "propFloat32Array", "1.1f"},
		{"test", "Test1", "propFloat64Array", "1.1"},
		{"test", "Test1", "propStringArray", "std::string(\"xyz\")"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppTestValue("", prop)
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
		{"test", "Test2", "propEnum", "Enum1Enum::NotDefault"},
		{"test", "InterfaceNamesCheck", "lowerEnumProp", "EnumLowerNamesEnum::secondValue"},
		{"test", "Test2", "propStruct", "Struct1()"},
		{"test", "Test2", "propInterface", "Interface1()"},
		{"test", "Test2", "propEnumArray", "Enum1Enum::NotDefault"},
		{"test", "Test2", "propStructArray", "Struct1()"},
		{"test", "Test2", "propInterfaceArray", "Interface1()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueWithErrors(t *testing.T) {
	t.Parallel()
	s, err := cppTestValue("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}

func TestTestValueFromIdlWithPrefix_makesNoDifference(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "true"},
		{"test", "Test1", "propInt", "1"},
		{"test", "Test1", "propInt32", "1"},
		{"test", "Test1", "propInt64", "1LL"},
		{"test", "Test1", "propFloat", "1.1f"},
		{"test", "Test1", "propFloat32", "1.1f"},
		{"test", "Test1", "propFloat64", "1.1"},
		{"test", "Test1", "propString", "std::string(\"xyz\")"},
		{"test", "Test1", "propBoolArray", "true"}, // all the array types return value intentionally, it may be put into empty array
		{"test", "Test1", "propIntArray", "1"},
		{"test", "Test1", "propInt32Array", "1"},
		{"test", "Test1", "propInt64Array", "1LL"},
		{"test", "Test1", "propFloatArray", "1.1f"},
		{"test", "Test1", "propFloat32Array", "1.1f"},
		{"test", "Test1", "propFloat64Array", "1.1"},
		{"test", "Test1", "propStringArray", "std::string(\"xyz\")"},
	}
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppTestValue(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueSymbolsFromIdlWithPrefix(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "my_prefix::Enum1Enum::NotDefault"},
		{"test", "Test2", "propStruct", "my_prefix::Struct1()"},
		{"test", "Test2", "propInterface", "my_prefix::Interface1()"},
		{"test", "Test2", "propEnumArray", "my_prefix::Enum1Enum::NotDefault"},
		{"test", "Test2", "propStructArray", "my_prefix::Struct1()"},
		{"test", "Test2", "propInterfaceArray", "my_prefix::Interface1()"},
	}
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppTestValue(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueExterns(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "prop1", "XType1()"},
		{"test_apigear_next", "Iface1", "prop2", "demo::x::XType2()"},
		{"test_apigear_next", "Iface1", "prop3", "demo::x::XtypeFactory::create()"},
		{"test_apigear_next", "Iface1", "propList", "demo::x::XtypeFactory::create()"},
		{"test_apigear_next", "Iface1", "propImportedEnum", "Test::Enum1Enum::NotDefault"},
		{"test_apigear_next", "Iface1", "propImportedStruct", "Test::Struct1()"},
	}
	syss := loadExternSystems(t)
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				prop := sys.LookupProperty(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, prop)
				r, err := cppTestValue(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
