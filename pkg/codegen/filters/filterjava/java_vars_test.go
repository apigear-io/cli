package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVars(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opBool", "param1"},
		{"test", "Test3", "opInt", "param1"},
		{"test", "Test3", "opInt32", "param1"},
		{"test", "Test3", "opInt64", "param1"},
		{"test", "Test3", "opFloat", "param1"},
		{"test", "Test3", "opFloat32", "param1"},
		{"test", "Test3", "opFloat64", "param1"},
		{"test", "Test3", "opString", "param1"},
		{"test", "Test3", "opBoolArray", "param1"},
		{"test", "Test3", "opIntArray", "param1"},
		{"test", "Test3", "opInt32Array", "param1"},
		{"test", "Test3", "opInt64Array", "param1"},
		{"test", "Test3", "opFloatArray", "param1"},
		{"test", "Test3", "opFloat32Array", "param1"},
		{"test", "Test3", "opFloat64Array", "param1"},
		{"test", "Test3", "opStringArray", "param1"},
		{"test", "Test3", "op_Bool", "param_Bool"},
		{"test", "Test3", "op_bool", "param_bool"},
		{"test", "Test3", "op_1", "param_1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := javaVars(op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
