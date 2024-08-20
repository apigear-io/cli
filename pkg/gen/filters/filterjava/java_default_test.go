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
		{"test", "Test2", "propInterface", "new Interface1()"},
		{"test", "Test2", "propEnumArray", "new Enum1[]{}"},
		{"test", "Test2", "propStructArray", "new Struct1[]{}"},
		{"test", "Test2", "propInterfaceArray", "new Interface1[]{}"},
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
		{"demo", "Iface1", "prop2", "new XType2()"},
		{"demo", "Iface1", "prop3", "new XType3A()"},
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
