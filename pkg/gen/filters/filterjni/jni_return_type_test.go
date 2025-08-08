package filterjni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestJniReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "jboolean"},
		{"test", "Test3", "opInt", "jint"},
		{"test", "Test3", "opInt32", "jint"},
		{"test", "Test3", "opInt64", "jlong"},
		{"test", "Test3", "opFloat", "jfloat"},
		{"test", "Test3", "opFloat32", "jfloat"},
		{"test", "Test3", "opFloat64", "jdouble"},
		{"test", "Test3", "opString", "jstring"},
		{"test", "Test3", "opBoolArray", "jbooleanArray"},
		{"test", "Test3", "opIntArray", "jintArray"},
		{"test", "Test3", "opInt32Array", "jintArray"},
		{"test", "Test3", "opInt64Array", "jlongArray"},
		{"test", "Test3", "opFloatArray", "jfloatArray"},
		{"test", "Test3", "opFloat32Array", "jfloatArray"},
		{"test", "Test3", "opFloat64Array", "jdoubleArray"},
		{"test", "Test3", "opStringArray", "jobjectArray"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := jniToReturnType(op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniOperationReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "jboolean"},
		{"test", "Test3", "opInt", "jint"},
		{"test", "Test3", "opInt32", "jint"},
		{"test", "Test3", "opInt64", "jlong"},
		{"test", "Test3", "opFloat", "jfloat"},
		{"test", "Test3", "opFloat32", "jfloat"},
		{"test", "Test3", "opFloat64", "jdouble"},
		{"test", "Test3", "opString", "jstring"},
		{"test", "Test3", "opBoolArray", "jbooleanArray"},
		{"test", "Test3", "opIntArray", "jintArray"},
		{"test", "Test3", "opInt32Array", "jintArray"},
		{"test", "Test3", "opInt64Array", "jlongArray"},
		{"test", "Test3", "opFloatArray", "jfloatArray"},
		{"test", "Test3", "opFloat32Array", "jfloatArray"},
		{"test", "Test3", "opFloat64Array", "jdoubleArray"},
		{"test", "Test3", "opStringArray", "jobjectArray"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := jniToReturnType(op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestJniReturnSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "jobject"},
		{"test", "Test2", "propStruct", "jobject"},
		{"test", "Test2", "propInterface", "jobject"},
		{"test", "Test2", "propEnumArray", "jobjectArray"},
		{"test", "Test2", "propStructArray", "jobjectArray"},
		{"test", "Test2", "propInterfaceArray", "jobjectArray"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := jniToReturnType(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
