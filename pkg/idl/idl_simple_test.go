package idl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleProps(t *testing.T) {
	s, err := loadIdl("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		pName string
		tName string
	}{
		{"SimpleInterface", "propBool", "bool"},
		{"SimpleInterface", "propInt", "int"},
		{"SimpleInterface", "propFloat", "float"},
		{"SimpleInterface", "propString", "string"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.pName), func(t *testing.T) {
			m := s.LookupModule("tb.simple")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.simple", tr.iName)
			assert.NotNil(t, i)
			p := i.LookupProperty(tr.pName)
			assert.NotNil(t, p)
			assert.Equal(t, tr.tName, p.Type)
			assert.Equal(t, tr.pName, p.Name)
		})
	}
}

func TestSimpleFuncs(t *testing.T) {
	s, err := loadIdl("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		fName string
		pName string
		tName string
	}{
		{"SimpleInterface", "funcBool", "paramBool", "bool"},
		{"SimpleInterface", "funcInt", "paramInt", "int"},
		{"SimpleInterface", "funcFloat", "paramFloat", "float"},
		{"SimpleInterface", "funcString", "paramString", "string"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.fName), func(t *testing.T) {
			m := s.LookupModule("tb.simple")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.simple", tr.iName)
			assert.NotNil(t, i)
			f := i.LookupOperation(tr.fName)
			assert.NotNil(t, f)
			assert.Equal(t, tr.fName, f.Name)
			assert.Equal(t, 1, len(f.Params))
			p := f.Params[0]
			assert.Equal(t, tr.pName, p.Name)
			assert.Equal(t, tr.tName, p.Type)
			assert.Equal(t, tr.tName, f.Return.Type)
		})
	}
}

func TestSimpleSignals(t *testing.T) {
	s, err := loadIdl("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		sName string
		pName string
		tName string
	}{
		{"SimpleInterface", "sigBool", "paramBool", "bool"},
		{"SimpleInterface", "sigInt", "paramInt", "int"},
		{"SimpleInterface", "sigFloat", "paramFloat", "float"},
		{"SimpleInterface", "sigString", "paramString", "string"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.sName), func(t *testing.T) {
			m := s.LookupModule("tb.simple")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.simple", tr.iName)
			assert.NotNil(t, i)
			s := i.LookupSignal(tr.sName)
			assert.NotNil(t, s)
			assert.Equal(t, tr.sName, s.Name)
			assert.Equal(t, 1, len(s.Params))
			p := s.Params[0]
			assert.Equal(t, tr.pName, p.Name)
			assert.Equal(t, tr.tName, p.Type)
		})
	}
}

func TestSimpleArrayProps(t *testing.T) {
	s, err := loadIdl("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		pName string
		tName string
	}{
		{"SimpleArrayInterface", "propBool", "bool"},
		{"SimpleArrayInterface", "propInt", "int"},
		{"SimpleArrayInterface", "propFloat", "float"},
		{"SimpleArrayInterface", "propString", "string"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.pName), func(t *testing.T) {
			m := s.LookupModule("tb.simple")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.simple", tr.iName)
			assert.NotNil(t, i)
			p := i.LookupProperty(tr.pName)
			assert.NotNil(t, p)
			assert.Equal(t, tr.tName, p.Type)
			assert.Equal(t, tr.pName, p.Name)
			assert.True(t, p.IsArray)
		})
	}
}

func TestSimpleArrayFuncs(t *testing.T) {
	s, err := loadIdl("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		fName string
		pName string
		tName string
	}{
		{"SimpleArrayInterface", "funcBool", "paramBool", "bool"},
		{"SimpleArrayInterface", "funcInt", "paramInt", "int"},
		{"SimpleArrayInterface", "funcFloat", "paramFloat", "float"},
		{"SimpleArrayInterface", "funcString", "paramString", "string"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.fName), func(t *testing.T) {
			m := s.LookupModule("tb.simple")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.simple", tr.iName)
			assert.NotNil(t, i)
			f := i.LookupOperation(tr.fName)
			assert.NotNil(t, f)
			assert.Equal(t, tr.fName, f.Name)
			assert.Equal(t, 1, len(f.Params))
			p := f.Params[0]
			assert.Equal(t, tr.pName, p.Name)
			assert.Equal(t, tr.tName, p.Type)
			assert.Equal(t, tr.tName, f.Return.Type)
			assert.True(t, p.IsArray)
		})
	}
}

func TestSimpleArraySignals(t *testing.T) {
	s, err := loadIdl("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	table := []struct {
		iName string
		sName string
		pName string
		tName string
	}{
		{"SimpleArrayInterface", "sigBool", "paramBool", "bool"},
		{"SimpleArrayInterface", "sigInt", "paramInt", "int"},
		{"SimpleArrayInterface", "sigFloat", "paramFloat", "float"},
		{"SimpleArrayInterface", "sigString", "paramString", "string"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.iName, tr.sName), func(t *testing.T) {
			m := s.LookupModule("tb.simple")
			assert.NotNil(t, m)
			i := s.LookupInterface("tb.simple", tr.iName)
			assert.NotNil(t, i)
			s := i.LookupSignal(tr.sName)
			assert.NotNil(t, s)
			assert.Equal(t, tr.sName, s.Name)
			assert.Equal(t, 1, len(s.Params))
			p := s.Params[0]
			assert.Equal(t, tr.pName, p.Name)
			assert.Equal(t, tr.tName, p.Type)
			assert.True(t, p.IsArray)
		})
	}
}
