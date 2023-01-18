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
		{"SimpleInterface", "propInt32", "int32"},
		{"SimpleInterface", "propInt64", "int64"},
		{"SimpleInterface", "propFloat", "float"},
		{"SimpleInterface", "propFloat32", "float32"},
		{"SimpleInterface", "propFloat64", "float64"},
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
		{"SimpleInterface", "funcInt32", "paramInt32", "int32"},
		{"SimpleInterface", "funcInt64", "paramInt64", "int64"},
		{"SimpleInterface", "funcFloat", "paramFloat", "float"},
		{"SimpleInterface", "funcFloat32", "paramFloat32", "float32"},
		{"SimpleInterface", "funcFloat64", "paramFloat64", "float64"},
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
		{"SimpleInterface", "sigInt32", "paramInt32", "int32"},
		{"SimpleInterface", "sigInt64", "paramInt64", "int64"},
		{"SimpleInterface", "sigFloat", "paramFloat", "float"},
		{"SimpleInterface", "sigFloat32", "paramFloat32", "float32"},
		{"SimpleInterface", "sigFloat64", "paramFloat64", "float64"},
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
		{"SimpleArrayInterface", "propInt32", "int32"},
		{"SimpleArrayInterface", "propInt64", "int64"},
		{"SimpleArrayInterface", "propFloat", "float"},
		{"SimpleArrayInterface", "propFloat32", "float32"},
		{"SimpleArrayInterface", "propFloat64", "float64"},
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
		{"SimpleArrayInterface", "funcInt32", "paramInt32", "int32"},
		{"SimpleArrayInterface", "funcInt64", "paramInt64", "int64"},
		{"SimpleArrayInterface", "funcFloat", "paramFloat", "float"},
		{"SimpleArrayInterface", "funcFloat32", "paramFloat32", "float32"},
		{"SimpleArrayInterface", "funcFloat64", "paramFloat64", "float64"},
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
		{"SimpleArrayInterface", "sigInt32", "paramInt32", "int32"},
		{"SimpleArrayInterface", "sigInt64", "paramInt64", "int64"},
		{"SimpleArrayInterface", "sigFloat", "paramFloat", "float"},
		{"SimpleArrayInterface", "sigFloat32", "paramFloat32", "float32"},
		{"SimpleArrayInterface", "sigFloat64", "paramFloat64", "float64"},
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
