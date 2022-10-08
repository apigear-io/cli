package filterts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestDefaultFromIdl(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propFloat", "0.0"},
		{"test", "Test1", "propString", "\"\""},
		{"test", "Test1", "propBoolArray", "[]"},
		{"test", "Test1", "propIntArray", "[]"},
		{"test", "Test1", "propFloatArray", "[]"},
		{"test", "Test1", "propStringArray", "[]"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := tsDefault("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1.Default"},
		{"test", "Test2", "propStruct", "new Struct1()"},
		{"test", "Test2", "propInterface", "null"},
		{"test", "Test2", "propEnumArray", "[]"},
		{"test", "Test2", "propStructArray", "[]"},
		{"test", "Test2", "propInterfaceArray", "[]"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := tsDefault("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestDefaultWithErrors(t *testing.T) {
	s, err := tsDefault("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
