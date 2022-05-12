package filtergo

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
		{"test", "Test3", "funcBool", "input1 bool"},
		{"test", "Test3", "funcInt", "input1 int"},
		{"test", "Test3", "funcFloat", "input1 float64"},
		{"test", "Test3", "funcString", "input1 string"},
		{"test", "Test3", "funcBoolArray", "input1 []bool"},
		{"test", "Test3", "funcIntArray", "input1 []int"},
		{"test", "Test3", "funcFloatArray", "input1 []float64"},
		{"test", "Test3", "funcStringArray", "input1 []string"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			meth := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, meth)
			r := goParams(meth.Inputs)
			assert.Equal(t, tt.rt, r)
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
		{"test", "Test4", "funcEnum", "input1 Enum1"},
		{"test", "Test4", "funcStruct", "input1 Struct1"},
		{"test", "Test4", "funcInterface", "input1 *Interface1"},
		{"test", "Test4", "funcEnumArray", "input1 []Enum1"},
		{"test", "Test4", "funcStructArray", "input1 []Struct1"},
		{"test", "Test4", "funcInterfaceArray", "input1 []*Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r := goParams(prop.Inputs)
			assert.Equal(t, tt.rt, r)
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
		{"test", "Test5", "funcBoolBool", "input1 bool, input2 bool"},
		{"test", "Test5", "funcIntInt", "input1 int, input2 int"},
		{"test", "Test5", "funcFloatFloat", "input1 float64, input2 float64"},
		{"test", "Test5", "funcStringString", "input1 string, input2 string"},
		{"test", "Test5", "funcEnumEnum", "input1 Enum1, input2 Enum1"},
		{"test", "Test5", "funcStructStruct", "input1 Struct1, input2 Struct1"},
		{"test", "Test5", "funcInterfaceInterface", "input1 *Interface1, input2 *Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r := goParams(prop.Inputs)
			assert.Equal(t, tt.rt, r)
		})
	}
}
