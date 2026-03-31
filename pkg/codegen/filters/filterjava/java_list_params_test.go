package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListParams(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "boolean param1"},
		{"test", "Test3", "opInt", "int param1"},
		{"test", "Test3", "opInt32", "int param1"},
		{"test", "Test3", "opInt64", "long param1"},
		{"test", "Test3", "opFloat", "float param1"},
		{"test", "Test3", "opFloat32", "float param1"},
		{"test", "Test3", "opFloat64", "double param1"},
		{"test", "Test3", "opString", "String param1"},
		{"test", "Test3", "opBoolArray", "List<Boolean> param1"},
		{"test", "Test3", "opIntArray", "List<Integer> param1"},
		{"test", "Test3", "opInt32Array", "List<Integer> param1"},
		{"test", "Test3", "opInt64Array", "List<Long> param1"},
		{"test", "Test3", "opFloatArray", "List<Float> param1"},
		{"test", "Test3", "opFloat32Array", "List<Float> param1"},
		{"test", "Test3", "opFloat64Array", "List<Double> param1"},
		{"test", "Test3", "opStringArray", "List<String> param1"},
		{"test", "Test3", "op_Bool", "boolean param_Bool"},
		{"test", "Test3", "op_bool", "boolean param_bool"},
		{"test", "Test3", "op_1", "boolean param_1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				m := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, m)
				r, err := javaListParams("", m.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
