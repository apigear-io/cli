package filtercs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	for _, sys := range syss {
		op := sys.LookupOperation("test", "Test3", "opString")
		assert.NotNil(t, op)
		r, err := csParam("", op.Params[0])
		assert.NoError(t, err)
		assert.Equal(t, "string param1", r)
	}
}

func TestParams(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "bool param1"},
		{"test", "Test3", "opInt", "int param1"},
		{"test", "Test3", "opInt32", "int param1"},
		{"test", "Test3", "opInt64", "long param1"},
		{"test", "Test3", "opFloat", "float param1"},
		{"test", "Test3", "opFloat32", "float param1"},
		{"test", "Test3", "opFloat64", "double param1"},
		{"test", "Test3", "opString", "string param1"},
		{"test", "Test3", "opBoolArray", "List<bool> param1"},
		{"test", "Test3", "opIntArray", "List<int> param1"},
		{"test", "Test3", "opStringArray", "List<string> param1"},
		{"test", "Test3", "op_Bool", "bool param_Bool"},
		{"test", "Test3", "op_bool", "bool param_bool"},
		{"test", "Test3", "op_1", "bool param_1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := csParams("", meth.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsSymbols(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "Enum1 param1"},
		{"test", "Test4", "opStruct", "Struct1 param1"},
		{"test", "Test4", "opInterface", "IInterface1 param1"},
		{"test", "Test4", "opEnumArray", "List<Enum1> param1"},
		{"test", "Test4", "opStructArray", "List<Struct1> param1"},
		{"test", "Test4", "opInterfaceArray", "List<IInterface1> param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := csParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsMultiple(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "bool param1, bool param2"},
		{"test", "Test5", "opIntInt", "int param1, int param2"},
		{"test", "Test5", "opStringString", "string param1, string param2"},
		{"test", "Test5", "opEnumEnum", "Enum1 param1, Enum1 param2"},
		{"test", "Test5", "opStructStruct", "Struct1 param1, Struct1 param2"},
		{"test", "Test5", "opInterfaceInterface", "IInterface1 param1, IInterface1 param2"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := csParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
