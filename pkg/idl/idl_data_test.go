package idl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataProps(t *testing.T) {
	s, err := loadIdl("data", []string{"./testdata/data.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName   string
		pName   string
		tName   string
		isArray bool
	}{
		{"StructInterface", "propBool", "StructBool", false},
		{"StructInterface", "propInt", "StructInt", false},
		{"StructInterface", "propFloat", "StructFloat", false},
		{"StructInterface", "propString", "StructString", false},
		{"StructArrayInterface", "propBool", "StructBool", true},
		{"StructArrayInterface", "propInt", "StructInt", true},
		{"StructArrayInterface", "propFloat", "StructFloat", true},
		{"StructArrayInterface", "propString", "StructString", true},
	}
	for _, tr := range table {
		t.Run(tr.iName+"."+tr.pName, func(t *testing.T) {
			m := s.LookupModule("tb.data")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.data", tr.iName)
			assert.NotNil(t, i)
			p := i.LookupProperty(tr.pName)
			assert.NotNil(t, p)
			assert.Equal(t, tr.tName, p.Type)
			assert.Equal(t, tr.pName, p.Name)
			assert.True(t, p.IsStruct())
			assert.Equal(t, tr.isArray, p.IsArray)
			ref := p.GetStruct()
			assert.NotNil(t, ref)
			assert.Equal(t, tr.tName, ref.Name)
		})
	}
}

func TestDataFuncs(t *testing.T) {
	s, err := loadIdl("data", []string{"./testdata/data.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName   string
		fName   string
		pName   string
		tName   string
		isArray bool
	}{
		{"StructInterface", "funcBool", "paramBool", "StructBool", false},
		{"StructInterface", "funcInt", "paramInt", "StructInt", false},
		{"StructInterface", "funcFloat", "paramFloat", "StructFloat", false},
		{"StructInterface", "funcString", "paramString", "StructString", false},
		{"StructArrayInterface", "funcBool", "paramBool", "StructBool", true},
		{"StructArrayInterface", "funcInt", "paramInt", "StructInt", true},
		{"StructArrayInterface", "funcFloat", "paramFloat", "StructFloat", true},
		{"StructArrayInterface", "funcString", "paramString", "StructString", true},
	}
	for _, tr := range table {
		t.Run(tr.iName+"."+tr.fName, func(t *testing.T) {
			m := s.LookupModule("tb.data")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.data", tr.iName)
			assert.NotNil(t, i)
			f := i.LookupOperation(tr.fName)
			assert.NotNil(t, f)
			assert.Equal(t, tr.fName, f.Name)
			assert.Equal(t, 1, len(f.Params))
			assert.Equal(t, tr.pName, f.Params[0].Name)
			assert.Equal(t, tr.tName, f.Params[0].Type)
			assert.Equal(t, tr.isArray, f.Params[0].IsArray)
			assert.True(t, f.Params[0].IsStruct())
			ref := f.Params[0].GetStruct()
			assert.NotNil(t, ref)
			assert.Equal(t, tr.tName, ref.Name)
		})
	}
}

func TestDataSignals(t *testing.T) {
	s, err := loadIdl("data", []string{"./testdata/data.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName   string
		sName   string
		pName   string
		tName   string
		isArray bool
	}{
		{"StructInterface", "sigBool", "paramBool", "StructBool", false},
		{"StructInterface", "sigInt", "paramInt", "StructInt", false},
		{"StructInterface", "sigFloat", "paramFloat", "StructFloat", false},
		{"StructInterface", "sigString", "paramString", "StructString", false},
		{"StructArrayInterface", "sigBool", "paramBool", "StructBool", true},
		{"StructArrayInterface", "sigInt", "paramInt", "StructInt", true},
		{"StructArrayInterface", "sigFloat", "paramFloat", "StructFloat", true},
		{"StructArrayInterface", "sigString", "paramString", "StructString", true},
	}
	for _, tr := range table {
		t.Run(tr.iName+"."+tr.sName, func(t *testing.T) {
			m := s.LookupModule("tb.data")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.data", tr.iName)
			assert.NotNil(t, i)
			f := i.LookupSignal(tr.sName)
			assert.NotNil(t, f)
			assert.Equal(t, tr.sName, f.Name)
			assert.Equal(t, 1, len(f.Params))
			assert.Equal(t, tr.pName, f.Params[0].Name)
			assert.Equal(t, tr.tName, f.Params[0].Type)
			assert.True(t, f.Params[0].IsStruct())
			ref := f.Params[0].GetStruct()
			assert.NotNil(t, ref)
			assert.Equal(t, tr.tName, ref.Name)
			assert.Equal(t, tr.isArray, f.Params[0].IsArray)
		})
	}
}

func TestStructs(t *testing.T) {
	s, err := loadIdl("structs", []string{"./testdata/data.idl"})
	assert.NoError(t, err)
	table := []struct {
		sName string
		fName string
		tName string
	}{
		{"StructBool", "fieldBool", "bool"},
		{"StructInt", "fieldInt", "int"},
		{"StructFloat", "fieldFloat", "float"},
		{"StructString", "fieldString", "string"},
	}
	for _, tr := range table {
		t.Run(tr.sName+"."+tr.fName, func(t *testing.T) {
			m := s.LookupModule("tb.data")
			assert.NotNil(t, m)
			s := m.LookupStruct(tr.sName)
			assert.NotNil(t, s)
			f := s.LookupField(tr.fName)
			assert.NotNil(t, f)
			assert.Equal(t, tr.fName, f.Name)
			assert.Equal(t, tr.tName, f.Type)
		})
	}
}
