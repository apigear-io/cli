package idl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumIdl(t *testing.T) {
	s, err := LoadIdlFromFiles("enum", []string{"./testdata/enum.idl"})
	assert.NoError(t, err)
	table := []struct {
		eName string
		vName string
		value int
	}{
		{"Enum0", "Value0", 0},
		{"Enum0", "Value1", 1},
		{"Enum0", "Value2", 2},
		{"Enum1", "Value1", 1},
		{"Enum1", "Value2", 2},
		{"Enum1", "Value3", 3},
		{"Enum2", "Value2", 2},
		{"Enum2", "Value1", 1},
		{"Enum2", "Value0", 0},
		{"Enum3", "Value3", 3},
		{"Enum3", "Value2", 2},
		{"Enum3", "Value1", 1},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.eName, tr.vName), func(t *testing.T) {
			e := s.LookupEnum("tb.enum", tr.eName)
			assert.NotNil(t, e)
			v := e.LookupMember(tr.vName)
			assert.NotNil(t, v)
			assert.Equal(t, tr.value, v.Value)
		})
	}
}

func TestEnumProps(t *testing.T) {
	s, err := LoadIdlFromFiles("enum", []string{"./testdata/enum.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		pName string
		tName string
	}{
		{"EnumInterface", "prop0", "Enum0"},
		{"EnumInterface", "prop1", "Enum1"},
		{"EnumInterface", "prop2", "Enum2"},
		{"EnumInterface", "prop3", "Enum3"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.pName), func(t *testing.T) {
			m := s.LookupModule("tb.enum")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.enum", tr.iName)

			assert.NotNil(t, i)
			p := i.LookupProperty(tr.pName)
			assert.NotNil(t, p)
			ref := m.LookupEnum(tr.tName)
			assert.NotNil(t, ref)
			assert.Equal(t, tr.tName, p.Type)
			assert.Equal(t, ref, p.GetEnum())
		})
	}
}

func TestEnumFuncs(t *testing.T) {
	s, err := LoadIdlFromFiles("enum", []string{"./testdata/enum.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		fName string
		pName string
		tName string
	}{
		{"EnumInterface", "func0", "param1", "Enum0"},
		{"EnumInterface", "func1", "param1", "Enum1"},
		{"EnumInterface", "func2", "param1", "Enum2"},
		{"EnumInterface", "func3", "param1", "Enum3"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.fName), func(t *testing.T) {
			m := s.LookupModule("tb.enum")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.enum", tr.iName)

			assert.NotNil(t, i)
			f := i.LookupOperation(tr.fName)
			assert.NotNil(t, f)
			ref := m.LookupEnum(tr.tName)
			assert.NotNil(t, ref)
			assert.Equal(t, tr.tName, f.Return.Type)
			assert.Equal(t, ref, f.Return.GetEnum())
			assert.Equal(t, 1, len(f.Params))
			assert.Equal(t, tr.tName, f.Params[0].Type)
			assert.Equal(t, ref, f.Params[0].GetEnum())
			assert.Equal(t, tr.pName, f.Params[0].Name)
		})
	}
}

func TestEnumSignals(t *testing.T) {
	s, err := LoadIdlFromFiles("enum", []string{"./testdata/enum.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		sName string
		pName string
		tName string
	}{
		{"EnumInterface", "sig0", "param1", "Enum0"},
		{"EnumInterface", "sig1", "param1", "Enum1"},
		{"EnumInterface", "sig2", "param1", "Enum2"},
		{"EnumInterface", "sig3", "param1", "Enum3"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.sName), func(t *testing.T) {
			m := s.LookupModule("tb.enum")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.enum", tr.iName)

			assert.NotNil(t, i)
			s := i.LookupSignal(tr.sName)
			assert.NotNil(t, s)
			ref := m.LookupEnum(tr.tName)
			assert.NotNil(t, ref)
			assert.Equal(t, 1, len(s.Params))
			assert.Equal(t, tr.tName, s.Params[0].Type)
			assert.Equal(t, ref, s.Params[0].GetEnum())
			assert.Equal(t, tr.pName, s.Params[0].Name)
		})
	}
}
