package filterjni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJniSignatureParams(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "Z"},
		{"test", "Test3", "opInt", "I"},
		{"test", "Test3", "opInt32", "I"},
		{"test", "Test3", "opInt64", "J"},
		{"test", "Test3", "opFloat", "F"},
		{"test", "Test3", "opFloat32", "F"},
		{"test", "Test3", "opFloat64", "D"},
		{"test", "Test3", "opString", "Ljava/lang/String;"},
		{"test", "Test3", "opBoolArray", "[Z"},
		{"test", "Test3", "opIntArray", "[I"},
		{"test", "Test3", "opInt32Array", "[I"},
		{"test", "Test3", "opInt64Array", "[J"},
		{"test", "Test3", "opFloatArray", "[F"},
		{"test", "Test3", "opFloat32Array", "[F"},
		{"test", "Test3", "opFloat64Array", "[D"},
		{"test", "Test3", "opStringArray", "[Ljava/lang/String;"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := jniJavaSignatureParams(meth.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniSignatureParamsSymbols(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "Ltest/test_api/Enum1;"},
		{"test", "Test4", "opStruct", "Ltest/test_api/Struct1;"},
		{"test", "Test4", "opInterface", "Ltest/test_api/Interface1;"},
		{"test", "Test4", "opEnumArray", "[Ltest/test_api/Enum1;"},
		{"test", "Test4", "opStructArray", "[Ltest/test_api/Struct1;"},
		{"test", "Test4", "opInterfaceArray", "[Ltest/test_api/Interface1;"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniJavaSignatureParams(prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniSignatureParamsMultiple(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "ZZ"},
		{"test", "Test5", "opIntInt", "II"},
		{"test", "Test5", "opFloatFloat", "FF"},
		{"test", "Test5", "opStringString", "Ljava/lang/String;Ljava/lang/String;"},
		{"test", "Test5", "opEnumEnum", "Ltest/test_api/Enum1;Ltest/test_api/Enum1;"},
		{"test", "Test5", "opStructStruct", "Ltest/test_api/Struct1;Ltest/test_api/Struct1;"},
		{"test", "Test5", "opInterfaceInterface", "Ltest/test_api/Interface1;Ltest/test_api/Interface1;"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniJavaSignatureParams(prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
