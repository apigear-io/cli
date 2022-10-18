package filterpy

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
		{"test", "Test3", "opBool", "param1"},
		{"test", "Test3", "opInt", "param1"},
		{"test", "Test3", "opFloat", "param1"},
		{"test", "Test3", "opString", "param1"},
		{"test", "Test3", "opBoolArray", "param1"},
		{"test", "Test3", "opIntArray", "param1"},
		{"test", "Test3", "opFloatArray", "param1"},
		{"test", "Test3", "opStringArray", "param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := pyVars(meth.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestVarsSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test4", "opEnum", "param1"},
		{"test", "Test4", "opStruct", "param1"},
		{"test", "Test4", "opInterface", "param1"},
		{"test", "Test4", "opEnumArray", "param1"},
		{"test", "Test4", "opStructArray", "param1"},
		{"test", "Test4", "opInterfaceArray", "param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyVars(prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestVarsMultiple(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "param1, param2"},
		{"test", "Test5", "opIntInt", "param1, param2"},
		{"test", "Test5", "opFloatFloat", "param1, param2"},
		{"test", "Test5", "opStringString", "param1, param2"},
		{"test", "Test5", "opEnumEnum", "param1, param2"},
		{"test", "Test5", "opStructStruct", "param1, param2"},
		{"test", "Test5", "opInterfaceInterface", "param1, param2"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyVars(prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
