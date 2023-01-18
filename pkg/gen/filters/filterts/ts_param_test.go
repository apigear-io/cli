package filterts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "propBool: boolean"},
		{"test", "Test1", "propInt", "propInt: number"},
		{"test", "Test1", "propInt32", "propInt32: number"},
		{"test", "Test1", "propInt64", "propInt64: number"},
		{"test", "Test1", "propFloat", "propFloat: number"},
		{"test", "Test1", "propFloat32", "propFloat32: number"},
		{"test", "Test1", "propFloat64", "propFloat64: number"},
		{"test", "Test1", "propString", "propString: string"},
		{"test", "Test1", "propBoolArray", "propBoolArray: boolean[]"},
		{"test", "Test1", "propIntArray", "propIntArray: number[]"},
		{"test", "Test1", "propInt32Array", "propInt32Array: number[]"},
		{"test", "Test1", "propInt64Array", "propInt64Array: number[]"},
		{"test", "Test1", "propFloatArray", "propFloatArray: number[]"},
		{"test", "Test1", "propFloat32Array", "propFloat32Array: number[]"},
		{"test", "Test1", "propFloat64Array", "propFloat64Array: number[]"},
		{"test", "Test1", "propStringArray", "propStringArray: string[]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := tsParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "propEnum: Enum1"},
		{"test", "Test2", "propStruct", "propStruct: Struct1"},
		{"test", "Test2", "propInterface", "propInterface: Interface1"},
		{"test", "Test2", "propEnumArray", "propEnumArray: Enum1[]"},
		{"test", "Test2", "propStructArray", "propStructArray: Struct1[]"},
		{"test", "Test2", "propInterfaceArray", "propInterfaceArray: Interface1[]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := tsParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
