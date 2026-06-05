package filtercs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propInt32", "0"},
		{"test", "Test1", "propInt64", "0L"},
		{"test", "Test1", "propFloat", "0.0f"},
		{"test", "Test1", "propFloat32", "0.0f"},
		{"test", "Test1", "propFloat64", "0.0"},
		{"test", "Test1", "propString", "string.Empty"},
		{"test", "Test1", "propBytes", "System.Array.Empty<byte>()"},
		{"test", "Test1", "propAny", "null"},
		{"test", "Test1", "propBoolArray", "new List<bool>()"},
		{"test", "Test1", "propIntArray", "new List<int>()"},
		{"test", "Test1", "propInt64Array", "new List<long>()"},
		{"test", "Test1", "propStringArray", "new List<string>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1.Default"},
		{"test", "Test2", "propStruct", "new Struct1()"},
		{"test", "Test2", "propInterface", "null"},
		{"test", "Test2", "propEnumArray", "new List<Enum1>()"},
		{"test", "Test2", "propStructArray", "new List<Struct1>()"},
		{"test", "Test2", "propInterfaceArray", "new List<IInterface1>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
