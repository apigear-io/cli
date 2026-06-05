package filtercs

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestPopulateFuncMap(t *testing.T) {
	t.Parallel()
	fm := template.FuncMap{}
	PopulateFuncMap(fm)
	expected := []string{
		"csReturn",
		"csType",
		"csDefault",
		"csParam",
		"csParams",
		"csVar",
		"csVars",
		"csTestValue",
		"csAsyncReturn",
		"csNs",
		"csNsOpen",
		"csNsClose",
		"csExtern",
		"csExterns",
	}
	for _, name := range expected {
		_, ok := fm[name]
		assert.Truef(t, ok, "expected func %q to be registered", name)
	}
}
