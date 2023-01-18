package filtercpp

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
		{"test", "Test3", "opBool", "bool param1"},
		{"test", "Test3", "opInt", "int param1"},
		{"test", "Test3", "opInt32", "int32_t param1"},
		{"test", "Test3", "opInt64", "int64_t param1"},
		{"test", "Test3", "opFloat", "float param1"},
		{"test", "Test3", "opFloat32", "float param1"},
		{"test", "Test3", "opFloat64", "double param1"},
		{"test", "Test3", "opString", "const std::string& param1"},
		{"test", "Test3", "opBoolArray", "const std::list<bool>& param1"},
		{"test", "Test3", "opIntArray", "const std::list<int>& param1"},
		{"test", "Test3", "opInt32Array", "const std::list<int32_t>& param1"},
		{"test", "Test3", "opInt64Array", "const std::list<int64_t>& param1"},
		{"test", "Test3", "opFloatArray", "const std::list<float>& param1"},
		{"test", "Test3", "opFloat32Array", "const std::list<float>& param1"},
		{"test", "Test3", "opFloat64Array", "const std::list<double>& param1"},
		{"test", "Test3", "opStringArray", "const std::list<std::string>& param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := cppParams("", meth.Params)
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
		{"test", "Test4", "opEnum", "Enum1Enum param1"},
		{"test", "Test4", "opStruct", "const Struct1& param1"},
		{"test", "Test4", "opInterface", "Interface1* param1"},
		{"test", "Test4", "opEnumArray", "const std::list<Enum1Enum>& param1"},
		{"test", "Test4", "opStructArray", "const std::list<Struct1>& param1"},
		{"test", "Test4", "opInterfaceArray", "const std::list<Interface1*>& param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := cppParams("", op.Params)
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
		{"test", "Test5", "opBoolBool", "bool param1, bool param2"},
		{"test", "Test5", "opIntInt", "int param1, int param2"},
		{"test", "Test5", "opFloatFloat", "float param1, float param2"},
		{"test", "Test5", "opStringString", "const std::string& param1, const std::string& param2"},
		{"test", "Test5", "opEnumEnum", "Enum1Enum param1, Enum1Enum param2"},
		{"test", "Test5", "opStructStruct", "const Struct1& param1, const Struct1& param2"},
		{"test", "Test5", "opInterfaceInterface", "Interface1* param1, Interface1* param2"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := cppParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
