package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "bool bParam1"},
		{"test", "Test3", "opInt", "int32 Param1"},
		{"test", "Test3", "opFloat", "float Param1"},
		{"test", "Test3", "opString", "const FString& Param1"},
		{"test", "Test3", "opBoolArray", "const TArray<bool>& Param1"},
		{"test", "Test3", "opIntArray", "const TArray<int32>& Param1"},
		{"test", "Test3", "opFloatArray", "const TArray<float>& Param1"},
		{"test", "Test3", "opStringArray", "const TArray<FString>& Param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := ueParam("", meth.Params[0])
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "const ETestEnum1& Param1"},
		{"test", "Test4", "opStruct", "const FTestStruct1& Param1"},
		{"test", "Test4", "opInterface", "FTestInterface1* Param1"},
		{"test", "Test4", "opEnumArray", "const TArray<ETestEnum1>& Param1"},
		{"test", "Test4", "opStructArray", "const TArray<FTestStruct1>& Param1"},
		{"test", "Test4", "opInterfaceArray", "const TArray<FTestInterface1*>& Param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := ueParam("", meth.Params[0])
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamWithErrors(t *testing.T) {
	s, err := ueParam("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
