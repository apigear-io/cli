package gen

import (
	"os"
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
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
	opts := GeneratorOptions{
		System:       model.NewSystem("test"),
		UserForce:    true,
		TemplatesDir: "testdata/templates",
		OutputDir:    "testdata/output",
		UserFeatures: []string{"all"},
	}
	g, err := New(opts)
	assert.NoError(t, err)
	err = g.ParseTemplatesDir("testdata/templates")
	assert.NoError(t, err)
	return g
}

func createMockGenerator(t *testing.T, tplDir string, features []string) (*generator, *MockOutput) {
	out := NewMockOutput()
	opts := GeneratorOptions{
		System:       model.NewSystem("test"),
		UserForce:    true,
		TemplatesDir: helper.Join(tplDir, "templates"),
		OutputDir:    "testdata/output",
		UserFeatures: features,
		Output:       out,
	}
	g, err := New(opts)
	assert.NoError(t, err)
	rules := readRules(t, helper.Join(tplDir, "rules.yaml"))
	err = g.ProcessRules(rules)
	assert.NoError(t, err)
	return g, out
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
