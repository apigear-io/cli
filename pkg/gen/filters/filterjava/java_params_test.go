package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
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
		{"test", "Test3", "opBoolArray", "boolean[] param1"},
		{"test", "Test3", "opIntArray", "int[] param1"},
		{"test", "Test3", "opInt32Array", "int[] param1"},
		{"test", "Test3", "opInt64Array", "long[] param1"},
		{"test", "Test3", "opFloatArray", "float[] param1"},
		{"test", "Test3", "opFloat32Array", "float[] param1"},
		{"test", "Test3", "opFloat64Array", "double[] param1"},
		{"test", "Test3", "opStringArray", "String[] param1"},
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
				r, err := javaParams("", m.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
