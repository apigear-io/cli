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
		{"test", "Test3", "funcBool", "input1: boolean"},
		{"test", "Test3", "funcInt", "input1: number"},
		{"test", "Test3", "funcFloat", "input1: number"},
		{"test", "Test3", "funcString", "input1: string"},
		{"test", "Test3", "funcBoolArray", "input1: boolean[]"},
		{"test", "Test3", "funcIntArray", "input1: number[]"},
		{"test", "Test3", "funcFloatArray", "input1: number[]"},
		{"test", "Test3", "funcStringArray", "input1: string[]"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := tsParams(m.Inputs)
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
		{"test", "Test4", "funcEnum", "input1: Enum1"},
		{"test", "Test4", "funcStruct", "input1: Struct1"},
		{"test", "Test4", "funcInterface", "input1: Interface1"},
		{"test", "Test4", "funcEnumArray", "input1: Enum1[]"},
		{"test", "Test4", "funcStructArray", "input1: Struct1[]"},
		{"test", "Test4", "funcInterfaceArray", "input1: Interface1[]"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := tsParams(m.Inputs)
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
		{"test", "Test5", "funcBoolBool", "input1: boolean, input2: boolean"},
		{"test", "Test5", "funcIntInt", "input1: number, input2: number"},
		{"test", "Test5", "funcFloatFloat", "input1: number, input2: number"},
		{"test", "Test5", "funcStringString", "input1: string, input2: string"},
		{"test", "Test5", "funcEnumEnum", "input1: Enum1, input2: Enum1"},
		{"test", "Test5", "funcStructStruct", "input1: Struct1, input2: Struct1"},
		{"test", "Test5", "funcInterfaceInterface", "input1: Interface1, input2: Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			m := sys.LookupMethod(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, m)
			r, err := tsParams(m.Inputs)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}
