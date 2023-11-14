package filterrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeRef(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "i32"},
		{"test", "Test1", "propInt32", "i32"},
		{"test", "Test1", "propInt64", "i64"},
		{"test", "Test1", "propFloat", "f32"},
		{"test", "Test1", "propFloat32", "f32"},
		{"test", "Test1", "propFloat64", "f64"},
		{"test", "Test1", "propString", "&String"},
		{"test", "Test1", "propBoolArray", "&Vec<bool>"},
		{"test", "Test1", "propIntArray", "&Vec<i32>"},
		{"test", "Test1", "propInt32Array", "&Vec<i32>"},
		{"test", "Test1", "propInt64Array", "&Vec<i64>"},
		{"test", "Test1", "propFloatArray", "&Vec<f32>"},
		{"test", "Test1", "propFloat32Array", "&Vec<f32>"},
		{"test", "Test1", "propFloat64Array", "&Vec<f64>"},
		{"test", "Test1", "propStringArray", "&Vec<String>"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rsTypeRef("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTypeRefSymbols(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1Enum"},
		{"test", "Test2", "propStruct", "&Struct1"},
		{"test", "Test2", "propInterface", "&Interface1"},
		{"test", "Test2", "propEnumArray", "&Vec<Enum1Enum>"},
		{"test", "Test2", "propStructArray", "&Vec<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "&Vec<&Interface1>"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rsTypeRef("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
