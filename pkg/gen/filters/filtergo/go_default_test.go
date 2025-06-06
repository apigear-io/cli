package filtergo

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
		{"test", "Test1", "propVoid", ""},
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "int32(0)"},
		{"test", "Test1", "propInt32", "int32(0)"},
		{"test", "Test1", "propInt64", "int64(0)"},
		{"test", "Test1", "propFloat", "float32(0.0)"},
		{"test", "Test1", "propFloat32", "float32(0.0)"},
		{"test", "Test1", "propFloat64", "float64(0.0)"},
		{"test", "Test1", "propString", "\"\""},
		{"test", "Test1", "propBytes", "[]byte{}"},
		{"test", "Test1", "propAny", "nil"},
		{"test", "Test1", "propBoolArray", "[]bool{}"},
		{"test", "Test1", "propIntArray", "[]int32{}"},
		{"test", "Test1", "propInt32Array", "[]int32{}"},
		{"test", "Test1", "propInt64Array", "[]int64{}"},
		{"test", "Test1", "propFloatArray", "[]float32{}"},
		{"test", "Test1", "propFloat32Array", "[]float32{}"},
		{"test", "Test1", "propFloat64Array", "[]float64{}"},
		{"test", "Test1", "propStringArray", "[]string{}"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn  string
		in  string
		pn  string
		val string
	}{
		{"test", "Test2", "propEnum", "Enum1Default"},
		{"test", "Test2", "propStruct", "Struct1{}"},
		{"test", "Test2", "propInterface", "nil"},
		{"test", "Test2", "propEnumArray", "[]Enum1{}"},
		{"test", "Test2", "propStructArray", "[]Struct1{}"},
		{"test", "Test2", "propInterfaceArray", "[]Interface1{}"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.val, r)
			})
		}
	}
}

func TestDefaultWithErrors(t *testing.T) {
	s, err := goDefault("", nil)
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
		{"demo", "Iface1", "prop1", "XType1{}"},
		{"demo", "Iface1", "prop2", "x.XType2{}"},
		{"demo", "Iface1", "prop3", "x.XType3A{}"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
