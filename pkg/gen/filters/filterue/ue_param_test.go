package filterue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "bool bPropBool"},
		{"test", "Test1", "propInt", "int32 PropInt"},
		{"test", "Test1", "propInt32", "int32 PropInt32"},
		{"test", "Test1", "propInt64", "int64 PropInt64"},
		{"test", "Test1", "propFloat", "float PropFloat"},
		{"test", "Test1", "propFloat32", "float PropFloat32"},
		{"test", "Test1", "propFloat64", "double PropFloat64"},
		{"test", "Test1", "propString", "const FString& PropString"},
		{"test", "Test1", "propBoolArray", "const TArray<bool>& PropBoolArray"},
		{"test", "Test1", "propIntArray", "const TArray<int32>& PropIntArray"},
		{"test", "Test1", "propInt32Array", "const TArray<int32>& PropInt32Array"},
		{"test", "Test1", "propInt64Array", "const TArray<int64>& PropInt64Array"},
		{"test", "Test1", "propFloatArray", "const TArray<float>& PropFloatArray"},
		{"test", "Test1", "propFloat32Array", "const TArray<float>& PropFloat32Array"},
		{"test", "Test1", "propFloat64Array", "const TArray<double>& PropFloat64Array"},
		{"test", "Test1", "propStringArray", "const TArray<FString>& PropStringArray"},
		{"test", "Test1", "prop_Bool", "bool bPropBool"},
		{"test", "Test1", "prop_bool", "bool bPropBool"},
		{"test", "Test1", "prop_1", "bool bProp1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "ETestEnum1 PropEnum"},
		{"test", "Test2", "propStruct", "const FTestStruct1& PropStruct"},
		{"test", "Test2", "propInterface", "FTestInterface1* PropInterface"},
		{"test", "Test2", "propEnumArray", "const TArray<ETestEnum1>& PropEnumArray"},
		{"test", "Test2", "propStructArray", "const TArray<FTestStruct1>& PropStructArray"},
		{"test", "Test2", "propInterfaceArray", "const TArray<FTestInterface1*>& PropInterfaceArray"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := ueParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamWithErrors(t *testing.T) {
	t.Parallel()
	s, err := ueParam("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}
