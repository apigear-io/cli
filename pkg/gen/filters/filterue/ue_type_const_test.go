package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestConstType(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int32"},
		{"test", "Test1", "propInt32", "int32"},
		{"test", "Test1", "propInt64", "int64"},
		{"test", "Test1", "propFloat", "float"},
		{"test", "Test1", "propFloat32", "float"},
		{"test", "Test1", "propFloat64", "double"},
		{"test", "Test1", "propString", "const FString&"},
		{"test", "Test1", "propBoolArray", "const TArray<bool>&"},
		{"test", "Test1", "propIntArray", "const TArray<int32>&"},
		{"test", "Test1", "propInt32Array", "const TArray<int32>&"},
		{"test", "Test1", "propInt64Array", "const TArray<int64>&"},
		{"test", "Test1", "propFloatArray", "const TArray<float>&"},
		{"test", "Test1", "propFloat32Array", "const TArray<float>&"},
		{"test", "Test1", "propFloat64Array", "const TArray<double>&"},
		{"test", "Test1", "propStringArray", "const TArray<FString>&"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueConstType("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestConstTypeSymbols(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "const ETestEnum1&"},
		{"test", "Test2", "propStruct", "const FTestStruct1&"},
		{"test", "Test2", "propInterface", "FTestInterface1*"},
		{"test", "Test2", "propEnumArray", "const TArray<ETestEnum1>&"},
		{"test", "Test2", "propStructArray", "const TArray<FTestStruct1>&"},
		{"test", "Test2", "propInterfaceArray", "const TArray<FTestInterface1*>&"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueConstType("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
