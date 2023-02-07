package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "bool bParam1"},
		{"test", "Test3", "opInt", "int32 Param1"},
		{"test", "Test3", "opInt32", "int32 Param1"},
		{"test", "Test3", "opInt64", "int64 Param1"},
		{"test", "Test3", "opFloat", "float Param1"},
		{"test", "Test3", "opFloat32", "float Param1"},
		{"test", "Test3", "opFloat64", "double Param1"},
		{"test", "Test3", "opString", "const FString& Param1"},
		{"test", "Test3", "opBoolArray", "const TArray<bool>& Param1"},
		{"test", "Test3", "opIntArray", "const TArray<int32>& Param1"},
		{"test", "Test3", "opInt32Array", "const TArray<int32>& Param1"},
		{"test", "Test3", "opInt64Array", "const TArray<int64>& Param1"},
		{"test", "Test3", "opFloatArray", "const TArray<float>& Param1"},
		{"test", "Test3", "opFloat32Array", "const TArray<float>& Param1"},
		{"test", "Test3", "opFloat64Array", "const TArray<double>& Param1"},
		{"test", "Test3", "opStringArray", "const TArray<FString>& Param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := ueParams("", meth.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test4", "opEnum", "ETestEnum1 Param1"},
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
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueParams("", prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsMultiple(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "bool bParam1, bool bParam2"},
		{"test", "Test5", "opIntInt", "int32 Param1, int32 Param2"},
		{"test", "Test5", "opFloatFloat", "float Param1, float Param2"},
		{"test", "Test5", "opStringString", "const FString& Param1, const FString& Param2"},
		{"test", "Test5", "opEnumEnum", "ETestEnum1 Param1, ETestEnum1 Param2"},
		{"test", "Test5", "opStructStruct", "const FTestStruct1& Param1, const FTestStruct1& Param2"},
		{"test", "Test5", "opInterfaceInterface", "FTestInterface1* Param1, FTestInterface1* Param2"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueParams("", prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamsWithErrors(t *testing.T) {
	s, err := ueParams("", nil)
	assert.Error(t, err)
	assert.Equal(t, "", s)
}
