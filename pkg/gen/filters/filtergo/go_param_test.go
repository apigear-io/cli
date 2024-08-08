package filtergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "propBool bool"},
		{"test", "Test1", "propInt", "propInt int32"},
		{"test", "Test1", "propInt32", "propInt32 int32"},
		{"test", "Test1", "propInt64", "propInt64 int64"},
		{"test", "Test1", "propFloat", "propFloat float32"},
		{"test", "Test1", "propString", "propString string"},
		{"test", "Test1", "propBytes", "propBytes []byte"},
		{"test", "Test1", "propBoolArray", "propBoolArray []bool"},
		{"test", "Test1", "propIntArray", "propIntArray []int32"},
		{"test", "Test1", "propFloatArray", "propFloatArray []float32"},
		{"test", "Test1", "propStringArray", "propStringArray []string"},
		{"test", "Test1", "propBytesArray", "propBytesArray [][]byte"},
		{"test", "Test1", "prop_Bool", "prop_Bool bool"},
		{"test", "Test1", "prop_bool", "prop_bool bool"},
		{"test", "Test1", "prop_1", "prop_1 bool"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "propEnum Enum1"},
		{"test", "Test2", "propStruct", "propStruct Struct1"},
		{"test", "Test2", "propInterface", "propInterface Interface1"},
		{"test", "Test2", "propEnumArray", "propEnumArray []Enum1"},
		{"test", "Test2", "propStructArray", "propStructArray []Struct1"},
		{"test", "Test2", "propInterfaceArray", "propInterfaceArray []Interface1"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamWithErrors(t *testing.T) {
	s, err := goParam("", nil)
	assert.Error(t, err)
	assert.Equal(t, "xxx", s)
}

func TestExternParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "func1", "arg1 XType1"},
		{"demo", "Iface1", "func2", "arg1 x.XType2"},
		{"demo", "Iface1", "func3", "arg1 x.XType3A"},
	}
	syss := loadExternSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := goParams("", op.Params)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
