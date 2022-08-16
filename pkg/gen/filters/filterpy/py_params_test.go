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
		{"test", "Test3", "opBool", "self, param1: bool"},
		{"test", "Test3", "opInt", "self, param1: int"},
		{"test", "Test3", "opFloat", "self, param1: float"},
		{"test", "Test3", "opString", "self, param1: str"},
		{"test", "Test3", "opBoolArray", "self, param1: list[bool]"},
		{"test", "Test3", "opIntArray", "self, param1: list[int]"},
		{"test", "Test3", "opFloatArray", "self, param1: list[float]"},
		{"test", "Test3", "opStringArray", "self, param1: list[str]"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyParams("", m.Params)
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
		{"test", "Test4", "opEnum", "self, param1: Enum1"},
		{"test", "Test4", "opStruct", "self, param1: Struct1"},
		{"test", "Test4", "opInterface", "self, param1: Interface1"},
		{"test", "Test4", "opEnumArray", "self, param1: list[Enum1]"},
		{"test", "Test4", "opStructArray", "self, param1: list[Struct1]"},
		{"test", "Test4", "opInterfaceArray", "self, param1: list[Interface1]"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyParams("", m.Params)
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
		{"test", "Test5", "opBoolBool", "self, param1: bool, param2: bool"},
		{"test", "Test5", "opIntInt", "self, param1: int, param2: int"},
		{"test", "Test5", "opFloatFloat", "self, param1: float, param2: float"},
		{"test", "Test5", "opStringString", "self, param1: str, param2: str"},
		{"test", "Test5", "opEnumEnum", "self, param1: Enum1, param2: Enum1"},
		{"test", "Test5", "opStructStruct", "self, param1: Struct1, param2: Struct1"},
		{"test", "Test5", "opInterfaceInterface", "self, param1: Interface1, param2: Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyParams("", m.Params)
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
		{"test", "Test5", "opBoolBool", "param1: bool, param2: bool"},
		{"test", "Test5", "opIntInt", "param1: int, param2: int"},
		{"test", "Test5", "opFloatFloat", "param1: float, param2: float"},
		{"test", "Test5", "opStringString", "param1: str, param2: str"},
		{"test", "Test5", "opEnumEnum", "param1: Enum1, param2: Enum1"},
		{"test", "Test5", "opStructStruct", "param1: Struct1, param2: Struct1"},
		{"test", "Test5", "opInterfaceInterface", "param1: Interface1, param2: Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := pyFuncParams("", m.Params)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}
