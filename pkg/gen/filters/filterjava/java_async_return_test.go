package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsyncReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "CompletableFuture<Void>"},
		{"test", "Test1", "propBool", "CompletableFuture<Boolean>"},
		{"test", "Test1", "propInt", "CompletableFuture<Integer>"},
		{"test", "Test1", "propInt32", "CompletableFuture<Integer>"},
		{"test", "Test1", "propInt64", "CompletableFuture<Long>"},
		{"test", "Test1", "propFloat", "CompletableFuture<Float>"},
		{"test", "Test1", "propFloat32", "CompletableFuture<Float>"},
		{"test", "Test1", "propFloat64", "CompletableFuture<Double>"},
		{"test", "Test1", "propString", "CompletableFuture<String>"},
		{"test", "Test1", "propBoolArray", "CompletableFuture<boolean[]>"},
		{"test", "Test1", "propIntArray", "CompletableFuture<int[]>"},
		{"test", "Test1", "propInt32Array", "CompletableFuture<int[]>"},
		{"test", "Test1", "propInt64Array", "CompletableFuture<long[]>"},
		{"test", "Test1", "propFloatArray", "CompletableFuture<float[]>"},
		{"test", "Test1", "propFloat32Array", "CompletableFuture<float[]>"},
		{"test", "Test1", "propFloat64Array", "CompletableFuture<double[]>"},
		{"test", "Test1", "propStringArray", "CompletableFuture<String[]>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaAsyncReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestOperationAsyncReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "CompletableFuture<Void>"},
		{"test", "Test3", "opBool", "CompletableFuture<Boolean>"},
		{"test", "Test3", "opInt", "CompletableFuture<Integer>"},
		{"test", "Test3", "opInt32", "CompletableFuture<Integer>"},
		{"test", "Test3", "opInt64", "CompletableFuture<Long>"},
		{"test", "Test3", "opFloat", "CompletableFuture<Float>"},
		{"test", "Test3", "opFloat32", "CompletableFuture<Float>"},
		{"test", "Test3", "opFloat64", "CompletableFuture<Double>"},
		{"test", "Test3", "opString", "CompletableFuture<String>"},
		{"test", "Test3", "opBoolArray", "CompletableFuture<boolean[]>"},
		{"test", "Test3", "opIntArray", "CompletableFuture<int[]>"},
		{"test", "Test3", "opInt32Array", "CompletableFuture<int[]>"},
		{"test", "Test3", "opInt64Array", "CompletableFuture<long[]>"},
		{"test", "Test3", "opFloatArray", "CompletableFuture<float[]>"},
		{"test", "Test3", "opFloat32Array", "CompletableFuture<float[]>"},
		{"test", "Test3", "opFloat64Array", "CompletableFuture<double[]>"},
		{"test", "Test3", "opStringArray", "CompletableFuture<String[]>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := javaAsyncReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
