package filterts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "param1: boolean"},
		{"test", "Test3", "opInt", "param1: number"},
		{"test", "Test3", "opInt32", "param1: number"},
		{"test", "Test3", "opInt64", "param1: number"},
		{"test", "Test3", "opFloat", "param1: number"},
		{"test", "Test3", "opFloat32", "param1: number"},
		{"test", "Test3", "opFloat64", "param1: number"},
		{"test", "Test3", "opString", "param1: string"},
		{"test", "Test3", "opBoolArray", "param1: boolean[]"},
		{"test", "Test3", "opIntArray", "param1: number[]"},
		{"test", "Test3", "opInt32Array", "param1: number[]"},
		{"test", "Test3", "opInt64Array", "param1: number[]"},
		{"test", "Test3", "opFloatArray", "param1: number[]"},
		{"test", "Test3", "opFloat32Array", "param1: number[]"},
		{"test", "Test3", "opFloat64Array", "param1: number[]"},
		{"test", "Test3", "opStringArray", "param1: string[]"},
		{"test", "Test3", "op_Bool", "param_Bool: boolean"},
		{"test", "Test3", "op_bool", "param_bool: boolean"},
		{"test", "Test3", "op_1", "param_1: boolean"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, m)
				r, err := tsParams("", m.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "param1: Enum1"},
		{"test", "Test4", "opStruct", "param1: Struct1"},
		{"test", "Test4", "opInterface", "param1: Interface1"},
		{"test", "Test4", "opEnumArray", "param1: Enum1[]"},
		{"test", "Test4", "opStructArray", "param1: Struct1[]"},
		{"test", "Test4", "opInterfaceArray", "param1: Interface1[]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, m)
				r, err := tsParams("", m.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsMultiple(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "param1: boolean, param2: boolean"},
		{"test", "Test5", "opIntInt", "param1: number, param2: number"},
		{"test", "Test5", "opFloatFloat", "param1: number, param2: number"},
		{"test", "Test5", "opStringString", "param1: string, param2: string"},
		{"test", "Test5", "opEnumEnum", "param1: Enum1, param2: Enum1"},
		{"test", "Test5", "opStructStruct", "param1: Struct1, param2: Struct1"},
		{"test", "Test5", "opInterfaceInterface", "param1: Interface1, param2: Interface1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, m)
				r, err := tsParams("", m.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
