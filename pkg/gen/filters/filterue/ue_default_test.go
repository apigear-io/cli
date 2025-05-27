package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestDefaultFromIdl(t *testing.T) {
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
		{"test", "Test1", "propInt64", "0LL"},
		{"test", "Test1", "propFloat", "0.0f"},
		{"test", "Test1", "propFloat32", "0.0f"},
		{"test", "Test1", "propFloat64", "0.0"},
		{"test", "Test1", "propString", "FString()"},
		{"test", "Test1", "propBoolArray", "TArray<bool>()"},
		{"test", "Test1", "propIntArray", "TArray<int32>()"},
		{"test", "Test1", "propInt32Array", "TArray<int32>()"},
		{"test", "Test1", "propInt64Array", "TArray<int64>()"},
		{"test", "Test1", "propFloatArray", "TArray<float>()"},
		{"test", "Test1", "propFloat32Array", "TArray<float>()"},
		{"test", "Test1", "propFloat64Array", "TArray<double>()"},
		{"test", "Test1", "propStringArray", "TArray<FString>()"},
	}
	for _, sys := range syss {
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
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn  string
		in  string
		pn  string
		val string
	}{
		// EnumValues: {"test", "Test2", "propEnum", "ETestEnum1::Default"},
		{"test", "Test2", "propEnum", "ETestEnum1::TE1_Default"},
		{"test", "Test2", "propStruct", "FTestStruct1()"},
		{"test", "Test2", "propInterface", "FTestInterface1()"},
		{"test", "Test2", "propEnumArray", "TArray<ETestEnum1>()"},
		{"test", "Test2", "propStructArray", "TArray<FTestStruct1>()"},
		{"test", "Test2", "propInterfaceArray", "TArray<FTestInterface1*>()"},
	}
	for _, sys := range syss {
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
}

func TestDefaultWithErrors(t *testing.T) {
	t.Parallel()
	s, err := ueDefault("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
