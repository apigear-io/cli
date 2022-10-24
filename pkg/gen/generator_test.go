package gen

import (
	"os"
	"testing"
	"text/template"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func readRules(t *testing.T, filename string) *spec.RulesDoc {
	content, err := os.ReadFile(filename)
	assert.NoError(t, err)
	var doc spec.RulesDoc
	err = yaml.Unmarshal(content, &doc)
	assert.NoError(t, err)
	return &doc
}

func createGenerator(t *testing.T) *generator {
	var g = &generator{
		Template:     template.New(""),
		System:       model.NewSystem("test"),
		UserForce:    true,
		TemplatesDir: "testdata/templates",
		OutputDir:    "testdata/output",
		Features:     []string{"all"},
	}
	err := g.ParseTemplatesDir("testdata/templates")
	assert.NoError(t, err)
	return g
}
func TestEmptyRules(t *testing.T) {
	g := createGenerator(t)
	doc, err := ReadRulesDoc("testdata/empty.rules.yaml")
	assert.NoError(t, err)
	assert.NoError(t, g.ProcessRules(doc))
}

func TestHelloRules(t *testing.T) {
	g := createGenerator(t)
	r := readRules(t, "testdata/test.rules.yaml")
	err := g.ProcessRules(r)
	assert.NoError(t, err)
	assert.Contains(t, g.Stats.FilesTouched[0], "system.txt")
}

func TestModules(t *testing.T) {
	g := createGenerator(t)
	r := readRules(t, "testdata/test.rules.yaml")
	err := g.ProcessRules(r)
	assert.NoError(t, err)
	assert.Contains(t, g.Stats.FilesTouched[0], "system.txt")
}
