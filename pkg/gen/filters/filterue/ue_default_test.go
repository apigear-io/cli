package filterue

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
		{"test", "Test1", "propFloat", "0.0f"},
		{"test", "Test1", "propString", "FString()"},
		{"test", "Test1", "propBoolArray", "TArray<bool>()"},
		{"test", "Test1", "propIntArray", "TArray<int32>()"},
		{"test", "Test1", "propFloatArray", "TArray<float>()"},
		{"test", "Test1", "propStringArray", "TArray<FString>()"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := ueDefault("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	sys := loadSystem(t)
	var propTests = []struct {
		mn  string
		in  string
		pn  string
		val string
	}{
		// EnumValues: {"test", "Test2", "propEnum", "ETestEnum1::Default"},
		{"test", "Test2", "propEnum", "ETestEnum1::DEFAULT"},
		{"test", "Test2", "propStruct", "FTestStruct1()"},
		{"test", "Test2", "propInterface", "FTestInterface1()"},
		{"test", "Test2", "propEnumArray", "TArray<ETestEnum1>()"},
		{"test", "Test2", "propStructArray", "TArray<FTestStruct1>()"},
		{"test", "Test2", "propInterfaceArray", "TArray<FTestInterface1>()"},
	}
	for _, tt := range propTests {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := ueDefault("", prop)
			assert.NoError(t, err)
			assert.Equal(t, tt.val, r)
		})
	}
}

func TestDefaultWithErrors(t *testing.T) {
	s, err := ueDefault("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
