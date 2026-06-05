package filtercs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestValue(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", ""},
		{"test", "Test1", "propBool", "true"},
		{"test", "Test1", "propInt", "1"},
		{"test", "Test1", "propInt32", "1"},
		{"test", "Test1", "propInt64", "1L"},
		{"test", "Test1", "propFloat", "1.0f"},
		{"test", "Test1", "propFloat32", "1.0f"},
		{"test", "Test1", "propFloat64", "1.0"},
		{"test", "Test1", "propString", "\"xyz\""},
		{"test", "Test1", "propBytes", "new byte[] { 1 }"},
		{"test", "Test1", "propAny", "1"},
		// array types intentionally return the inner element value
		{"test", "Test1", "propBoolArray", "true"},
		{"test", "Test1", "propIntArray", "1"},
		{"test", "Test1", "propStringArray", "\"xyz\""},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1.NotDefault"},
		{"test", "InterfaceNamesCheck", "lowerEnumProp", "EnumLowerNames.SecondValue"},
		{"test", "Test2", "propStruct", "new Struct1()"},
		{"test", "Test2", "propInterface", "new Interface1()"},
		{"test", "Test2", "propEnumArray", "Enum1.NotDefault"},
		{"test", "Test2", "propStructArray", "new Struct1()"},
		{"test", "Test2", "propInterfaceArray", "new Interface1()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueSymbolsWithPrefix(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "MyPrefix.Enum1.NotDefault"},
		{"test", "Test2", "propStruct", "new MyPrefix.Struct1()"},
		{"test", "Test2", "propInterface", "new MyPrefix.Interface1()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csTestValue("MyPrefix.", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueWithErrors(t *testing.T) {
	t.Parallel()
	s, err := csTestValue("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
