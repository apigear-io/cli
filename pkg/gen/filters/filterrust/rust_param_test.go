package filterrust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "propBool: bool"},
		{"test", "Test1", "propInt", "propInt: i32"},
		{"test", "Test1", "propInt32", "propInt32: i32"},
		{"test", "Test1", "propInt64", "propInt64: i64"},
		{"test", "Test1", "propFloat", "propFloat: f32"},
		{"test", "Test1", "propFloat32", "propFloat32: f32"},
		{"test", "Test1", "propFloat64", "propFloat64: f64"},
		{"test", "Test1", "propString", "propString: &String"},
		{"test", "Test1", "propBoolArray", "propBoolArray: &[bool]"},
		{"test", "Test1", "propIntArray", "propIntArray: &[i32]"},
		{"test", "Test1", "propFloatArray", "propFloatArray: &[f32]"},
		{"test", "Test1", "propStringArray", "propStringArray: &[String]"},
		{"test", "Test1", "prop_Bool", "prop_Bool: bool"},
		{"test", "Test1", "prop_bool", "prop_bool: bool"},
		{"test", "Test1", "prop_1", "prop_1: bool"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rustParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "propEnum: Enum1Enum"},
		{"test", "Test2", "propStruct", "propStruct: &Struct1"},
		{"test", "Test2", "propInterface", "propInterface: &Interface1"},
		{"test", "Test2", "propEnumArray", "propEnumArray: &[Enum1Enum]"},
		{"test", "Test2", "propStructArray", "propStructArray: &[Struct1]"},
		{"test", "Test2", "propInterfaceArray", "propInterfaceArray: &[&Interface1]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rustParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
