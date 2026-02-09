package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUeIsStdSimpleType(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt bool
	}{
		{"test", "Test3", "opBool", true},
		{"test", "Test3", "opInt", true},
		{"test", "Test3", "opInt32", true},
		{"test", "Test3", "opInt64", true},
		{"test", "Test3", "opFloat", true},
		{"test", "Test3", "opFloat32", true},
		{"test", "Test3", "opFloat64", true},
		{"test", "Test3", "opString", false},
		{"test", "Test3", "opBoolArray", false},
		{"test", "Test3", "opIntArray", false},
		{"test", "Test3", "opInt32Array", false},
		{"test", "Test3", "opInt64Array", false},
		{"test", "Test3", "opFloatArray", false},
		{"test", "Test3", "opFloat32Array", false},
		{"test", "Test3", "opFloat64Array", false},
		{"test", "Test3", "opStringArray", false},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := ueIsStdSimpleType(op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestUeIsStdSimpleTypeSymbols(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt bool
	}{
		{"test", "Test2", "propEnum", true},
		{"test", "Test2", "propStruct", false},
		{"test", "Test2", "propInterface", false},
		{"test", "Test2", "propEnumArray", false},
		{"test", "Test2", "propStructArray", false},
		{"test", "Test2", "propInterfaceArray", false},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueIsStdSimpleType(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
