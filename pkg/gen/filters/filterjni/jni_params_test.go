package filterjni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJniParams(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "jboolean param1"},
		{"test", "Test3", "opInt", "jint param1"},
		{"test", "Test3", "opInt32", "jint param1"},
		{"test", "Test3", "opInt64", "jlong param1"},
		{"test", "Test3", "opFloat", "jfloat param1"},
		{"test", "Test3", "opFloat32", "jfloat param1"},
		{"test", "Test3", "opFloat64", "jdouble param1"},
		{"test", "Test3", "opString", "jstring param1"},
		{"test", "Test3", "opBoolArray", "jbooleanArray param1"},
		{"test", "Test3", "opIntArray", "jintArray param1"},
		{"test", "Test3", "opInt32Array", "jintArray param1"},
		{"test", "Test3", "opInt64Array", "jlongArray param1"},
		{"test", "Test3", "opFloatArray", "jfloatArray param1"},
		{"test", "Test3", "opFloat32Array", "jfloatArray param1"},
		{"test", "Test3", "opFloat64Array", "jdoubleArray param1"},
		{"test", "Test3", "opStringArray", "jobjectArray param1"},
		{"test", "Test3", "op_Bool", "jboolean param_Bool"},
		{"test", "Test3", "op_bool", "jboolean param_bool"},
		{"test", "Test3", "op_1", "jboolean param_1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				meth := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, meth)
				r, err := jniJavaParams("", meth.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniParamsSymbols(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{

		{"test", "Test4", "opEnum", "jobject param1"},
		{"test", "Test4", "opStruct", "jobject param1"},
		{"test", "Test4", "opInterface", "jobject param1"},
		{"test", "Test4", "opEnumArray", "jobjectArray param1"},
		{"test", "Test4", "opStructArray", "jobjectArray param1"},
		{"test", "Test4", "opInterfaceArray", "jobjectArray param1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniJavaParams("", prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniParamsMultiple(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "jboolean param1, jboolean param2"},
		{"test", "Test5", "opIntInt", "jint param1, jint param2"},
		{"test", "Test5", "opFloatFloat", "jfloat param1, jfloat param2"},
		{"test", "Test5", "opStringString", "jstring param1, jstring param2"},
		{"test", "Test5", "opEnumEnum", "jobject param1, jobject param2"},
		{"test", "Test5", "opStructStruct", "jobject param1, jobject param2"},
		{"test", "Test5", "opInterfaceInterface", "jobject param1, jobject param2"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniJavaParams("", prop.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
