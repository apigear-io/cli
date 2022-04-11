package filtercpp

import (
	"bytes"
	"objectapi/pkg/idl"
	"objectapi/pkg/model"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func loadSystem(t *testing.T) *model.System {
	p := idl.NewIDLParser(model.NewSystem("test"))
	err := p.ParseFile("testdata/test.idl")
	assert.NoError(t, err)
	return p.System
}

func lookupProperty(t *testing.T, sys *model.System, mn string, in string, pn string) model.TypedNode {
	m := sys.ModuleByName(mn)
	assert.NotNil(t, m)
	i := m.InterfaceByName(in)
	assert.NotNil(t, i)
	prop := i.PropertyByName(pn)
	assert.NotNil(t, prop)
	return prop
}
func TestReturn(t *testing.T) {
	sys := loadSystem(t)
	prop := lookupProperty(t, sys, "test", "Test", "prop1")
	tpl := template.New("test")
	tpl.Funcs(template.FuncMap{
		"cpp_return": cppReturn,
	})

	tpl, err := tpl.Parse("{{ cpp_return . }}")
	assert.NoError(t, err)
	buf := bytes.Buffer{}
	err = tpl.Execute(&buf, &prop)
	assert.NoError(t, err)
	assert.Equal(t, "bool", buf.String())
}
