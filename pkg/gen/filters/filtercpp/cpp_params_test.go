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
		{"test", "Test3", "funcBool", "bool input1"},
		{"test", "Test3", "funcInt", "int input1"},
		{"test", "Test3", "funcFloat", "double input1"},
		{"test", "Test3", "funcString", "const std::string &input1"},
		{"test", "Test3", "funcBoolArray", "const std::vector<bool> &input1"},
		{"test", "Test3", "funcIntArray", "const std::vector<int> &input1"},
		{"test", "Test3", "funcFloatArray", "const std::vector<double> &input1"},
		{"test", "Test3", "funcStringArray", "const std::vector<std::string> &input1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			meth := sys.LookupMethod(tt.mn, tt.in, tt.pn)
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
		{"test", "Test4", "funcEnum", "Enum1 input1"},
		{"test", "Test4", "funcStruct", "const Struct1 &input1"},
		{"test", "Test4", "funcInterface", "Interface1 *input1"},
		{"test", "Test4", "funcEnumArray", "const std::vector<Enum1> &input1"},
		{"test", "Test4", "funcStructArray", "const std::vector<Struct1> &input1"},
		{"test", "Test4", "funcInterfaceArray", "const std::vector<Interface1*> &input1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupMethod(tt.mn, tt.in, tt.pn)
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
		{"test", "Test5", "funcBoolBool", "bool input1, bool input2"},
		{"test", "Test5", "funcIntInt", "int input1, int input2"},
		{"test", "Test5", "funcFloatFloat", "double input1, double input2"},
		{"test", "Test5", "funcStringString", "const std::string &input1, const std::string &input2"},
		{"test", "Test5", "funcEnumEnum", "Enum1 input1, Enum1 input2"},
		{"test", "Test5", "funcStructStruct", "const Struct1 &input1, const Struct1 &input2"},
		{"test", "Test5", "funcInterfaceInterface", "Interface1 *input1, Interface1 *input2"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := cppParams(reflect.ValueOf(prop))
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r.String())
		})
	}
}
