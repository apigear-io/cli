package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestReturn(t *testing.T) {
	syss := loadTestSystems(t)
	var retTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "int32"},
		{"test", "Test3", "opFloat", "float"},
		{"test", "Test3", "opString", "FString"},
		{"test", "Test3", "opBoolArray", "TArray<bool>"},
		{"test", "Test3", "opIntArray", "TArray<int32>"},
		{"test", "Test3", "opFloatArray", "TArray<float>"},
		{"test", "Test3", "opStringArray", "TArray<FString>"},
	}
	for _, sys := range syss {
		for _, tt := range retTests {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := ueReturn("", meth.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestReturnSymbols(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "ETestEnum1"},
		{"test", "Test4", "opStruct", "FTestStruct1"},
		{"test", "Test4", "opInterface", "FTestInterface1*"},
		{"test", "Test4", "opEnumArray", "TArray<ETestEnum1>"},
		{"test", "Test4", "opStructArray", "TArray<FTestStruct1>"},
		{"test", "Test4", "opInterfaceArray", "TArray<FTestInterface1*>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := ueReturn("", meth.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
