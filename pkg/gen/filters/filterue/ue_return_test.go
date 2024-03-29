package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "int32"},
		{"test", "Test3", "opInt32", "int32"},
		{"test", "Test3", "opInt64", "int64"},
		{"test", "Test3", "opFloat", "float"},
		{"test", "Test3", "opFloat32", "float"},
		{"test", "Test3", "opFloat64", "double"},
		{"test", "Test3", "opString", "FString"},
		{"test", "Test3", "opBoolArray", "TArray<bool>"},
		{"test", "Test3", "opIntArray", "TArray<int32>"},
		{"test", "Test3", "opInt32Array", "TArray<int32>"},
		{"test", "Test3", "opInt64Array", "TArray<int64>"},
		{"test", "Test3", "opFloatArray", "TArray<float>"},
		{"test", "Test3", "opFloat32Array", "TArray<float>"},
		{"test", "Test3", "opFloat64Array", "TArray<double>"},
		{"test", "Test3", "opStringArray", "TArray<FString>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := ueReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestOperationReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "int32"},
		{"test", "Test3", "opInt32", "int32"},
		{"test", "Test3", "opInt64", "int64"},
		{"test", "Test3", "opFloat", "float"},
		{"test", "Test3", "opFloat32", "float"},
		{"test", "Test3", "opFloat64", "double"},
		{"test", "Test3", "opString", "FString"},
		{"test", "Test3", "opBoolArray", "TArray<bool>"},
		{"test", "Test3", "opIntArray", "TArray<int32>"},
		{"test", "Test3", "opInt32Array", "TArray<int32>"},
		{"test", "Test3", "opInt64Array", "TArray<int64>"},
		{"test", "Test3", "opFloatArray", "TArray<float>"},
		{"test", "Test3", "opFloat32Array", "TArray<float>"},
		{"test", "Test3", "opFloat64Array", "TArray<double>"},
		{"test", "Test3", "opStringArray", "TArray<FString>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := ueReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestReturnSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "ETestEnum1"},
		{"test", "Test2", "propStruct", "FTestStruct1"},
		{"test", "Test2", "propInterface", "FTestInterface1*"},
		{"test", "Test2", "propEnumArray", "TArray<ETestEnum1>"},
		{"test", "Test2", "propStructArray", "TArray<FTestStruct1>"},
		{"test", "Test2", "propInterfaceArray", "TArray<FTestInterface1*>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
