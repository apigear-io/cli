package idl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManyParamsFuncs(t *testing.T) {
	s, err := LoadIdlFromFiles("advanced", []string{"./testdata/advanced.idl"})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	table := []struct {
		iName  string
		fName  string
		pCount int
	}{
		{"ManyParamInterface", "func0", 0},
		{"ManyParamInterface", "func1", 1},
		{"ManyParamInterface", "func2", 2},
		{"ManyParamInterface", "func3", 3},
		{"ManyParamInterface", "func4", 4},
		{"NestedStruct1Interface", "func1", 1},
		{"NestedStruct2Interface", "func1", 1},
		{"NestedStruct2Interface", "func2", 2},
		{"NestedStruct3Interface", "func1", 1},
		{"NestedStruct3Interface", "func2", 2},
		{"NestedStruct3Interface", "func3", 3},
	}
	for _, tr := range table {
		t.Run(tr.iName+"."+tr.fName, func(t *testing.T) {
			i := s.LookupInterface("tb.adv", tr.iName)
			assert.NotNil(t, i)
			f := i.LookupOperation(tr.fName)
			assert.NotNil(t, f)
			assert.Equal(t, tr.pCount, len(f.Params))
		})
	}
}

func TestManyParamsSigs(t *testing.T) {
	s, err := LoadIdlFromFiles("advanced", []string{"./testdata/advanced.idl"})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	table := []struct {
		iName    string
		sName    string
		pCount   int
		isStruct bool
	}{
		{"ManyParamInterface", "sig0", 0, false},
		{"ManyParamInterface", "sig1", 1, false},
		{"ManyParamInterface", "sig2", 2, false},
		{"ManyParamInterface", "sig3", 3, false},
		{"ManyParamInterface", "sig4", 4, false},
		{"NestedStruct1Interface", "sig1", 1, true},
		{"NestedStruct2Interface", "sig1", 1, true},
		{"NestedStruct2Interface", "sig2", 2, true},
		{"NestedStruct3Interface", "sig1", 1, true},
		{"NestedStruct3Interface", "sig2", 2, true},
		{"NestedStruct3Interface", "sig3", 3, true},
	}
	for _, tr := range table {
		t.Run(tr.iName+"."+tr.sName, func(t *testing.T) {
			i := s.LookupInterface("tb.adv", tr.iName)
			assert.NotNil(t, i)
			f := i.LookupSignal(tr.sName)
			assert.NotNil(t, f)
			assert.Equal(t, tr.pCount, len(f.Params))
			if tr.pCount > 0 {
				assert.Equal(t, tr.isStruct, f.Params[0].IsStruct())
			}
		})
	}
}
