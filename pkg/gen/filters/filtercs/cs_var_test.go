package filtercs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVar(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	for _, sys := range syss {
		op := sys.LookupOperation("test", "Test3", "opString")
		assert.NotNil(t, op)
		r, err := csVar(op.Params[0])
		assert.NoError(t, err)
		assert.Equal(t, "param1", r)
	}
}

func TestVars(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test5", "opBoolBool", "param1, param2"},
		{"test", "Test5", "opStringString", "param1, param2"},
		{"test", "Test3", "op_Bool", "param_Bool"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := csVars(op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
