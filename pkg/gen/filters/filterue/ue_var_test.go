package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestVar(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test1", "propBool", "bPropBool"},
		{"test", "Test1", "propInt", "PropInt"},
		{"test", "Test1", "propInt32", "PropInt32"},
		{"test", "Test1", "propInt64", "PropInt64"},
		{"test", "Test1", "propFloat", "PropFloat"},
		{"test", "Test1", "propFloat32", "PropFloat32"},
		{"test", "Test1", "propFloat64", "PropFloat64"},
		{"test", "Test1", "propString", "PropString"},
		{"test", "Test1", "propBoolArray", "PropBoolArray"},
		{"test", "Test1", "propIntArray", "PropIntArray"},
		{"test", "Test1", "propInt32Array", "PropInt32Array"},
		{"test", "Test1", "propInt64Array", "PropInt64Array"},
		{"test", "Test1", "propFloatArray", "PropFloatArray"},
		{"test", "Test1", "propFloat32Array", "PropFloat32Array"},
		{"test", "Test1", "propFloat64Array", "PropFloat64Array"},
		{"test", "Test1", "propStringArray", "PropStringArray"},
		{"test", "Test1", "prop_Bool", "bPropBool"},
		{"test", "Test1", "prop_bool", "bPropBool"},
		{"test", "Test1", "prop_1", "bProp1"},
	}
	for _, sys := range syss {
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
}

func TestVarSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
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
	for _, sys := range syss {
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
}
