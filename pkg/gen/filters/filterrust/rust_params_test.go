package filterrust

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
		{"test", "Test3", "opBool", "param1: bool"},
		{"test", "Test3", "opInt", "param1: i32"},
		{"test", "Test3", "opInt32", "param1: i32"},
		{"test", "Test3", "opInt64", "param1: i64"},
		{"test", "Test3", "opFloat", "param1: f32"},
		{"test", "Test3", "opFloat32", "param1: f32"},
		{"test", "Test3", "opFloat64", "param1: f64"},
		{"test", "Test3", "opString", "param1: &String"},
		{"test", "Test3", "opBoolArray", "param1: &[bool]"},
		{"test", "Test3", "opIntArray", "param1: &[i32]"},
		{"test", "Test3", "opInt32Array", "param1: &[i32]"},
		{"test", "Test3", "opInt64Array", "param1: &[i64]"},
		{"test", "Test3", "opFloatArray", "param1: &[f32]"},
		{"test", "Test3", "opFloat32Array", "param1: &[f32]"},
		{"test", "Test3", "opFloat64Array", "param1: &[f64]"},
		{"test", "Test3", "opStringArray", "param1: &[String]"},
		{"test", "Test3", "op_Bool", "param_bool: bool"},
		{"test", "Test3", "op_bool", "param_bool: bool"},
		{"test", "Test3", "op_1", "param_1: bool"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := rustParams("", meth.Params)
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
		{"test", "Test4", "opEnum", "param1: Enum1Enum"},
		{"test", "Test4", "opStruct", "param1: &Struct1"},
		{"test", "Test4", "opInterface", "param1: &Interface1"},
		{"test", "Test4", "opEnumArray", "param1: &[Enum1Enum]"},
		{"test", "Test4", "opStructArray", "param1: &[Struct1]"},
		{"test", "Test4", "opInterfaceArray", "param1: &[&Interface1]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := rustParams("", op.Params)
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
		{"test", "Test5", "opBoolBool", "param1: bool, param2: bool"},
		{"test", "Test5", "opIntInt", "param1: i32, param2: i32"},
		{"test", "Test5", "opFloatFloat", "param1: f32, param2: f32"},
		{"test", "Test5", "opStringString", "param1: &String, param2: &String"},
		{"test", "Test5", "opEnumEnum", "param1: Enum1Enum, param2: Enum1Enum"},
		{"test", "Test5", "opStructStruct", "param1: &Struct1, param2: &Struct1"},
		{"test", "Test5", "opInterfaceInterface", "param1: &Interface1, param2: &Interface1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := rustParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
