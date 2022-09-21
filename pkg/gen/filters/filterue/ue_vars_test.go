package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVars(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "bParam1"},
		{"test", "Test3", "opInt", "Param1"},
		{"test", "Test3", "opFloat", "Param1"},
		{"test", "Test3", "opString", "Param1"},
		{"test", "Test3", "opBoolArray", "Param1"},
		{"test", "Test3", "opIntArray", "Param1"},
		{"test", "Test3", "opFloatArray", "Param1"},
		{"test", "Test3", "opStringArray", "Param1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, meth)
			r, err := ueVars("", meth.Params)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestVarsSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test4", "opEnum", "Param1"},
		{"test", "Test4", "opStruct", "Param1"},
		{"test", "Test4", "opInterface", "Param1"},
		{"test", "Test4", "opEnumArray", "Param1"},
		{"test", "Test4", "opStructArray", "Param1"},
		{"test", "Test4", "opInterfaceArray", "Param1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := ueVars("", prop.Params)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestVarsMultiple(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "bParam1, bParam2"},
		{"test", "Test5", "opIntInt", "Param1, Param2"},
		{"test", "Test5", "opFloatFloat", "Param1, Param2"},
		{"test", "Test5", "opStringString", "Param1, Param2"},
		{"test", "Test5", "opEnumEnum", "Param1, Param2"},
		{"test", "Test5", "opStructStruct", "Param1, Param2"},
		{"test", "Test5", "opInterfaceInterface", "Param1, Param2"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := ueVars("", prop.Params)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}
