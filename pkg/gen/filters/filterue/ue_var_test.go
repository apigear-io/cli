package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestVar(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test1", "propBool", "bPropBool"},
		{"test", "Test1", "propInt", "PropInt"},
		{"test", "Test1", "propFloat", "PropFloat"},
		{"test", "Test1", "propString", "PropString"},
		{"test", "Test1", "propBoolArray", "PropBoolArray"},
		{"test", "Test1", "propIntArray", "PropIntArray"},
		{"test", "Test1", "propFloatArray", "PropFloatArray"},
		{"test", "Test1", "propStringArray", "PropStringArray"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := ueVar("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestVarSymbols(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "PropEnum"},
		{"test", "Test2", "propStruct", "PropStruct"},
		{"test", "Test2", "propInterface", "PropInterface"},
		{"test", "Test2", "propEnumArray", "PropEnumArray"},
		{"test", "Test2", "propStructArray", "PropStructArray"},
		{"test", "Test2", "propInterfaceArray", "PropInterfaceArray"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := ueVar("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}
