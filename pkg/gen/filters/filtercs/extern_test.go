package filtercs

import (
	"testing"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/stretchr/testify/assert"
)

func makeExtern(name string, meta model.Meta) *model.Extern {
	return &model.Extern{NamedNode: model.NamedNode{Name: name, Meta: meta}}
}

func TestExternParsing(t *testing.T) {
	t.Parallel()
	xe := makeExtern("Foo", model.Meta{
		"cs.namespace": "Bar.Baz",
		"cs.name":      "FooX",
		"cs.default":   "MakeFoo()",
		"cs.using":     "Bar.Baz",
		"cs.package":   "Bar.Pkg",
		"cs.version":   "1.2.3",
	})
	got := csExtern(xe)
	assert.Equal(t, "Bar.Baz", got.Namespace)
	assert.Equal(t, "FooX", got.Name)
	assert.Equal(t, "MakeFoo()", got.Default)
	assert.Equal(t, "Bar.Baz", got.Using)
	assert.Equal(t, "Bar.Pkg", got.Package)
	assert.Equal(t, "1.2.3", got.Version)
}

func TestExternNameFallsBackToExternName(t *testing.T) {
	t.Parallel()
	got := csExtern(makeExtern("Plain", model.Meta{}))
	assert.Equal(t, "Plain", got.Name)
	assert.Equal(t, "", got.Namespace)
}

func TestExterns(t *testing.T) {
	t.Parallel()
	externs := []*model.Extern{
		makeExtern("A", model.Meta{"cs.namespace": "N1"}),
		makeExtern("B", model.Meta{"cs.name": "BB"}),
	}
	got := csExterns(externs)
	assert.Len(t, got, 2)
	assert.Equal(t, "N1", got[0].Namespace)
	assert.Equal(t, "A", got[0].Name)
	assert.Equal(t, "BB", got[1].Name)
}

func TestExternReturn(t *testing.T) {
	syss := loadExternSystems(t)
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "prop1", "XType1"},
		{"demo", "Iface1", "prop2", "Demo.X.XType2"},
		{"demo", "Iface1", "prop3", "Demo.X.XType3A"},
	}
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestExternDefault(t *testing.T) {
	syss := loadExternSystems(t)
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "prop1", "new XType1()"},
		{"demo", "Iface1", "prop2", "new Demo.X.XType2()"},
		{"demo", "Iface1", "prop3", "Demo.X.XTypeFactory.Create()"},
	}
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestExternTestValue(t *testing.T) {
	syss := loadExternSystems(t)
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"demo", "Iface1", "prop1", "new XType1()"},
		{"demo", "Iface1", "prop2", "new Demo.X.XType2()"},
		{"demo", "Iface1", "prop3", "Demo.X.XTypeFactory.Create()"},
	}
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := csTestValue("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
