package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "boolean propBool"},
		{"test", "Test1", "propInt", "int propInt"},
		{"test", "Test1", "propInt32", "int propInt32"},
		{"test", "Test1", "propInt64", "long propInt64"},
		{"test", "Test1", "propFloat", "float propFloat"},
		{"test", "Test1", "propFloat32", "float propFloat32"},
		{"test", "Test1", "propFloat64", "double propFloat64"},
		{"test", "Test1", "propString", "String propString"},
		{"test", "Test1", "propBoolArray", "boolean[] propBoolArray"},
		{"test", "Test1", "propIntArray", "int[] propIntArray"},
		{"test", "Test1", "propInt32Array", "int[] propInt32Array"},
		{"test", "Test1", "propInt64Array", "long[] propInt64Array"},
		{"test", "Test1", "propFloatArray", "float[] propFloatArray"},
		{"test", "Test1", "propFloat32Array", "float[] propFloat32Array"},
		{"test", "Test1", "propFloat64Array", "double[] propFloat64Array"},
		{"test", "Test1", "propStringArray", "String[] propStringArray"},
		{"test", "Test1", "prop_Bool", "boolean prop_Bool"},
		{"test", "Test1", "prop_bool", "boolean prop_bool"},
		{"test", "Test1", "prop_1", "boolean prop_1"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1 propEnum"},
		{"test", "Test2", "propStruct", "Struct1 propStruct"},
		{"test", "Test2", "propInterface", "Interface1 propInterface"},
		{"test", "Test2", "propEnumArray", "Enum1[] propEnumArray"},
		{"test", "Test2", "propStructArray", "Struct1[] propStructArray"},
		{"test", "Test2", "propInterfaceArray", "Interface1[] propInterfaceArray"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamWithErrors(t *testing.T) {
	s, err := javaParam("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}

func TestExternParam(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "func1", "XType1 arg1"},
		{"demo", "Iface1", "func2", "XType2 arg1"},
		{"demo", "Iface1", "func3", "XType3A arg1"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := javaParam("", op.Params[0])
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
