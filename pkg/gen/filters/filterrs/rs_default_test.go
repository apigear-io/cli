package filterrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operations params, operation return, signal params, struct fields
func TestDefaultFromIdl(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "Default::default()"},
		{"test", "Test1", "propBool", "Default::default()"},
		{"test", "Test1", "propInt", "Default::default()"},
		{"test", "Test1", "propInt32", "Default::default()"},
		{"test", "Test1", "propInt64", "Default::default()"},
		{"test", "Test1", "propFloat", "Default::default()"},
		{"test", "Test1", "propFloat32", "Default::default()"},
		{"test", "Test1", "propFloat64", "Default::default()"},
		{"test", "Test1", "propString", "Default::default()"},
		{"test", "Test1", "propBoolArray", "Default::default()"},
		{"test", "Test1", "propIntArray", "Default::default()"},
		{"test", "Test1", "propInt32Array", "Default::default()"},
		{"test", "Test1", "propInt64Array", "Default::default()"},
		{"test", "Test1", "propFloatArray", "Default::default()"},
		{"test", "Test1", "propFloat32Array", "Default::default()"},
		{"test", "Test1", "propFloat64Array", "Default::default()"},
		{"test", "Test1", "propStringArray", "Default::default()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rsDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Default::default()"},
		{"test", "Test2", "propStruct", "Default::default()"},
		{"test", "Test2", "propInterface", "Default::default()"},
		{"test", "Test2", "propEnumArray", "Default::default()"},
		{"test", "Test2", "propStructArray", "Default::default()"},
		{"test", "Test2", "propInterfaceArray", "Default::default()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := rsDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
