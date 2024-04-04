package filtergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExternPropsReturn(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "prop1", "XType1"},
		{"demo", "Iface1", "prop2", "x.XType2"},
		{"demo", "Iface1", "prop3", "x.XType3A"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := goReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestExternFuncArgsReturn(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		fa string
		rt string
	}{
		{"demo", "Iface1", "func1", "arg1 XType1", "XType1"},
		{"demo", "Iface1", "func2", "arg1 x.XType2", "x.XType2"},
		{"demo", "Iface1", "func3", "arg1 x.XType3A", "x.XType3A"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				assert.Len(t, op.Params, 1)
				r, err := goParam("", op.Params[0])
				assert.NoError(t, err)
				assert.Equal(t, tt.fa, r)
				r, err = goReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestExternStructField(t *testing.T) {
	syss := loadExternSystems(t)
	var propTests = []struct {
		mn string
		in string
		fn string
		ft string
	}{
		{"demo", "Struct1", "field1", "XType1"},
		{"demo", "Struct1", "field2", "x.XType2"},
		{"demo", "Struct1", "field3", "x.XType3A"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.fn, func(t *testing.T) {
				prop := sys.LookupField(tt.mn, tt.in, tt.fn)
				assert.NotNil(t, prop)
				r, err := goReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.ft, r)
			})
		}
	}
}
