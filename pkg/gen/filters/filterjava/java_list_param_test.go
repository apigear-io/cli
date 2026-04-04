package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListParam(t *testing.T) {
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
		{"test", "Test1", "propBoolArray", "List<Boolean> propBoolArray"},
		{"test", "Test1", "propIntArray", "List<Integer> propIntArray"},
		{"test", "Test1", "propInt32Array", "List<Integer> propInt32Array"},
		{"test", "Test1", "propInt64Array", "List<Long> propInt64Array"},
		{"test", "Test1", "propFloatArray", "List<Float> propFloatArray"},
		{"test", "Test1", "propFloat32Array", "List<Float> propFloat32Array"},
		{"test", "Test1", "propFloat64Array", "List<Double> propFloat64Array"},
		{"test", "Test1", "propStringArray", "List<String> propStringArray"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaListParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestListParamSymbols(t *testing.T) {
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
		{"test", "Test2", "propInterface", "IInterface1 propInterface"},
		{"test", "Test2", "propEnumArray", "List<Enum1> propEnumArray"},
		{"test", "Test2", "propStructArray", "List<Struct1> propStructArray"},
		{"test", "Test2", "propInterfaceArray", "List<IInterface1> propInterfaceArray"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaListParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
