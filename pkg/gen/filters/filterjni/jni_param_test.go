package filterjni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJniParam(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "jboolean propBool"},
		{"test", "Test1", "propInt", "jint propInt"},
		{"test", "Test1", "propInt32", "jint propInt32"},
		{"test", "Test1", "propInt64", "jlong propInt64"},
		{"test", "Test1", "propFloat", "jfloat propFloat"},
		{"test", "Test1", "propFloat32", "jfloat propFloat32"},
		{"test", "Test1", "propFloat64", "jdouble propFloat64"},
		{"test", "Test1", "propString", "jstring propString"},
		{"test", "Test1", "propBoolArray", "jbooleanArray propBoolArray"},
		{"test", "Test1", "propIntArray", "jintArray propIntArray"},
		{"test", "Test1", "propInt32Array", "jintArray propInt32Array"},
		{"test", "Test1", "propInt64Array", "jlongArray propInt64Array"},
		{"test", "Test1", "propFloatArray", "jfloatArray propFloatArray"},
		{"test", "Test1", "propFloat32Array", "jfloatArray propFloat32Array"},
		{"test", "Test1", "propFloat64Array", "jdoubleArray propFloat64Array"},
		{"test", "Test1", "propStringArray", "jobjectArray propStringArray"},
		{"test", "Test1", "prop_Bool", "jboolean prop_Bool"},
		{"test", "Test1", "prop_bool", "jboolean prop_bool"},
		{"test", "Test1", "prop_1", "jboolean prop_1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniJavaParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniParamSymbols(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "jobject propEnum"},
		{"test", "Test2", "propStruct", "jobject propStruct"},
		{"test", "Test2", "propInterface", "jobject propInterface"},
		{"test", "Test2", "propEnumArray", "jobjectArray propEnumArray"},
		{"test", "Test2", "propStructArray", "jobjectArray propStructArray"},
		{"test", "Test2", "propInterfaceArray", "jobjectArray propInterfaceArray"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniJavaParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
