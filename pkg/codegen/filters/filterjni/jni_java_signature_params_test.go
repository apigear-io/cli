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
		{"test", "Test4", "opInterface", "Ltest/test_api/IInterface1;"},
		{"test", "Test4", "opEnumArray", "[Ltest/test_api/Enum1;"},
		{"test", "Test4", "opStructArray", "[Ltest/test_api/Struct1;"},
		{"test", "Test4", "opInterfaceArray", "[Ltest/test_api/IInterface1;"},
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
		{"test", "Test5", "opInterfaceInterface", "Ltest/test_api/IInterface1;Ltest/test_api/IInterface1;"},
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

func TestImportedExternSignatureParam(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface2", "prop1", "LXType1;"},
		{"demo", "Iface2", "prop2", "Ldemo/x/XType2;"},
		{"demo", "Iface2", "prop3", "Ldemo/x/XType3A;"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniJavaSignatureParam(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
func TestSignatureParamsExternsYaml(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "func1", "LXType1;"},
		{"test_apigear_next", "Iface1", "func3", "Ldemo/x/XType3A;"},
		{"test_apigear_next", "Iface1", "funcList", "[Ldemo/x/XType3A;"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "Ltest/test_api/Enum1;"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "Ltest/test_api/Struct1;"},
	}
	syss := loadExternSystemsYAML(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := jniJavaSignatureParams(op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
