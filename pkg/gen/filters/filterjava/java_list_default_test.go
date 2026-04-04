package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListDefault(t *testing.T) {
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
		{"test", "Test1", "propString", "new String()"},
		{"test", "Test1", "propBoolArray", "new ArrayList<>()"},
		{"test", "Test1", "propIntArray", "new ArrayList<>()"},
		{"test", "Test1", "propInt32Array", "new ArrayList<>()"},
		{"test", "Test1", "propInt64Array", "new ArrayList<>()"},
		{"test", "Test1", "propFloatArray", "new ArrayList<>()"},
		{"test", "Test1", "propFloat32Array", "new ArrayList<>()"},
		{"test", "Test1", "propFloat64Array", "new ArrayList<>()"},
		{"test", "Test1", "propStringArray", "new ArrayList<>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaListDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestListDefaultSymbols(t *testing.T) {
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
		{"test", "Test2", "propEnumArray", "new ArrayList<>()"},
		{"test", "Test2", "propStructArray", "new ArrayList<>()"},
		{"test", "Test2", "propInterfaceArray", "new ArrayList<>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaListDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
