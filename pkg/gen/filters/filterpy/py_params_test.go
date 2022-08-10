package filterpy

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
		{"test", "Test3", "funcBool", "self, input1: bool"},
		{"test", "Test3", "funcInt", "self, input1: int"},
		{"test", "Test3", "funcFloat", "self, input1: float"},
		{"test", "Test3", "funcString", "self, input1: str"},
		{"test", "Test3", "funcBoolArray", "self, input1: list[bool]"},
		{"test", "Test3", "funcIntArray", "self, input1: list[int]"},
		{"test", "Test3", "funcFloatArray", "self, input1: list[float]"},
		{"test", "Test3", "funcStringArray", "self, input1: list[str]"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyParams("", m.Inputs)
			assert.NoError(t, err)
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
		{"test", "Test4", "funcEnum", "self, input1: Enum1"},
		{"test", "Test4", "funcStruct", "self, input1: Struct1"},
		{"test", "Test4", "funcInterface", "self, input1: Interface1"},
		{"test", "Test4", "funcEnumArray", "self, input1: list[Enum1]"},
		{"test", "Test4", "funcStructArray", "self, input1: list[Struct1]"},
		{"test", "Test4", "funcInterfaceArray", "self, input1: list[Interface1]"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyParams("", m.Inputs)
			assert.NoError(t, err)
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
		{"test", "Test5", "funcBoolBool", "self, input1: bool, input2: bool"},
		{"test", "Test5", "funcIntInt", "self, input1: int, input2: int"},
		{"test", "Test5", "funcFloatFloat", "self, input1: float, input2: float"},
		{"test", "Test5", "funcStringString", "self, input1: str, input2: str"},
		{"test", "Test5", "funcEnumEnum", "self, input1: Enum1, input2: Enum1"},
		{"test", "Test5", "funcStructStruct", "self, input1: Struct1, input2: Struct1"},
		{"test", "Test5", "funcInterfaceInterface", "self, input1: Interface1, input2: Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyParams("", m.Inputs)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestFuncParamsMultiple(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "funcBoolBool", "input1: bool, input2: bool"},
		{"test", "Test5", "funcIntInt", "input1: int, input2: int"},
		{"test", "Test5", "funcFloatFloat", "input1: float, input2: float"},
		{"test", "Test5", "funcStringString", "input1: str, input2: str"},
		{"test", "Test5", "funcEnumEnum", "input1: Enum1, input2: Enum1"},
		{"test", "Test5", "funcStructStruct", "input1: Struct1, input2: Struct1"},
		{"test", "Test5", "funcInterfaceInterface", "input1: Interface1, input2: Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyFuncParams("", m.Inputs)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}