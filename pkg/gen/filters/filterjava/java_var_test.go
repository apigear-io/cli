package filterjava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVar(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "propBool"},
		{"test", "Test1", "propInt", "propInt"},
		{"test", "Test1", "propInt32", "propInt32"},
		{"test", "Test1", "propInt64", "propInt64"},
		{"test", "Test1", "propFloat", "propFloat"},
		{"test", "Test1", "propFloat32", "propFloat32"},
		{"test", "Test1", "propFloat64", "propFloat64"},
		{"test", "Test1", "propString", "propString"},
		{"test", "Test1", "propBoolArray", "propBoolArray"},
		{"test", "Test1", "propIntArray", "propIntArray"},
		{"test", "Test1", "propInt32Array", "propInt32Array"},
		{"test", "Test1", "propInt64Array", "propInt64Array"},
		{"test", "Test1", "propFloatArray", "propFloatArray"},
		{"test", "Test1", "propFloat32Array", "propFloat32Array"},
		{"test", "Test1", "propFloat64Array", "propFloat64Array"},
		{"test", "Test1", "propStringArray", "propStringArray"},
		{"test", "Test1", "prop_Bool", "prop_Bool"},
		{"test", "Test1", "prop_bool", "prop_bool"},
		{"test", "Test1", "prop_1", "prop_1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := javaVar(prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
