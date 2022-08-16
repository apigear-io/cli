package filtercpp

import (
	"reflect"
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
		{"test", "Test3", "opFloat", "double param1"},
		{"test", "Test3", "opString", "const std::string &param1"},
		{"test", "Test3", "opBoolArray", "const std::vector<bool> &param1"},
		{"test", "Test3", "opIntArray", "const std::vector<int> &param1"},
		{"test", "Test3", "opFloatArray", "const std::vector<double> &param1"},
		{"test", "Test3", "opStringArray", "const std::vector<std::string> &param1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, meth)
			r, err := cppParams(reflect.ValueOf(meth))
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r.String())
		})
	}
}

func TestParamsSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "Enum1 param1"},
		{"test", "Test4", "opStruct", "const Struct1 &param1"},
		{"test", "Test4", "opInterface", "Interface1 *param1"},
		{"test", "Test4", "opEnumArray", "const std::vector<Enum1> &param1"},
		{"test", "Test4", "opStructArray", "const std::vector<Struct1> &param1"},
		{"test", "Test4", "opInterfaceArray", "const std::vector<Interface1*> &param1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := cppParams(reflect.ValueOf(prop))
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r.String())
		})
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
		{"test", "Test5", "opFloatFloat", "double param1, double param2"},
		{"test", "Test5", "opStringString", "const std::string &param1, const std::string &param2"},
		{"test", "Test5", "opEnumEnum", "Enum1 param1, Enum1 param2"},
		{"test", "Test5", "opStructStruct", "const Struct1 &param1, const Struct1 &param2"},
		{"test", "Test5", "opInterfaceInterface", "Interface1 *param1, Interface1 *param2"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := cppParams(reflect.ValueOf(prop))
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r.String())
		})
	}
}
