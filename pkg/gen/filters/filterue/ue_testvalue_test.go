package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestTestValueFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "true"},
		{"test", "Test1", "propInt", "1"},
		{"test", "Test1", "propInt32", "1"},
		{"test", "Test1", "propInt64", "1LL"},
		{"test", "Test1", "propFloat", "1.0f"},
		{"test", "Test1", "propFloat32", "1.0f"},
		{"test", "Test1", "propFloat64", "1.0"},
		{"test", "Test1", "propString", "FString(\"xyz\")"},
		{"test", "Test1", "propBoolArray", "true"},
		{"test", "Test1", "propIntArray", "1"},
		{"test", "Test1", "propInt32Array", "1"},
		{"test", "Test1", "propInt64Array", "1LL"},
		{"test", "Test1", "propFloatArray", "1.0f"},
		{"test", "Test1", "propFloat32Array", "1.0f"},
		{"test", "Test1", "propFloat64Array", "1.0"},
		{"test", "Test1", "propStringArray", "FString(\"xyz\")"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestTestValueSymbolsFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn  string
		in  string
		pn  string
		val string
	}{
		// EnumValues: {"test", "Test2", "propEnum", "ETestEnum1::Default"},
		{"test", "Test2", "propEnum", "ETestEnum1::TE1_NOTDEFAULT"},
		{"test", "Test2", "propStruct", "FTestStruct1()"},
		{"test", "Test2", "propInterface", "FTestInterface1()"},
		{"test", "Test2", "propEnumArray", "ETestEnum1::TE1_NOTDEFAULT"},
		{"test", "Test2", "propStructArray", "FTestStruct1()"},
		{"test", "Test2", "propInterfaceArray", "FTestInterface1()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.val, r)
			})
		}
	}
}

func TestTestValueWithErrors(t *testing.T) {
	t.Parallel()
	s, err := ueTestValue("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
