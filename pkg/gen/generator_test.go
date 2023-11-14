package gen

import (
	"os"
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func readRules(t *testing.T, filename string) *spec.RulesDoc {
	content, err := os.ReadFile(filename)
	require.NoError(t, err)
	var doc spec.RulesDoc
	err = yaml.Unmarshal(content, &doc)
	require.NoError(t, err)
	return &doc
}

func createGenerator(t *testing.T) *generator {
	opts := GeneratorOptions{
		System:         model.NewSystem("test"),
		TargetForce:    true,
		TemplatesDir:   "testdata/templates",
		OutputDir:      "testdata/output",
		TargetFeatures: []string{"all"},
	}
	g, err := New(opts)
	require.NoError(t, err)
	err = g.ParseTemplatesDir("testdata/templates")
	require.NoError(t, err)
	return g
}

func createMockGenerator(t *testing.T, tplDir string, features []string) (*generator, *MockOutput) {
	out := NewMockOutput()
	opts := GeneratorOptions{
		System:         model.NewSystem("test"),
		TargetForce:    true,
		TemplatesDir:   helper.Join(tplDir, "templates"),
		OutputDir:      "testdata/output",
		TargetFeatures: features,
		Output:         out,
	}
	g, err := New(opts)
	require.NoError(t, err)
	rules := readRules(t, helper.Join(tplDir, "rules.yaml"))
	err = g.ProcessRules(rules)
	require.NoError(t, err)
	return g, out
}

func TestEmptyRules(t *testing.T) {
	g := createGenerator(t)
	doc, err := ReadRulesDoc("testdata/empty.rules.yaml")
	require.NoError(t, err)
	require.NoError(t, g.ProcessRules(doc))
}

func TestHelloRules(t *testing.T) {
	g := createGenerator(t)
	g.TargetForce = true
	r := readRules(t, "testdata/test.rules.yaml")
	err := g.ProcessRules(r)
	require.NoError(t, err)
	require.Len(t, g.Stats.FilesTouched, 0)
}

func TestHelloForcedRules(t *testing.T) {
	g := createGenerator(t)
	g.TargetForce = true
	r := readRules(t, "testdata/test-force.rules.yaml")
	err := g.ProcessRules(r)
	require.NoError(t, err)
	length := len(g.Stats.FilesTouched)
	require.Equal(t, 1, length)
	require.Contains(t, g.Stats.FilesTouched[0], "system-force.txt")
}

func TestForce(t *testing.T) {
	tt := []struct {
		Name         string
		TargetForce  bool
		FilesTouched []string
	}{
		{"force", true, []string{"testdata/output/system-force.txt"}},
		{"no-force", false, []string{}},
	}

	for _, tr := range tt {
		t.Run(tr.Name, func(t *testing.T) {
			g := createGenerator(t)
			g.TargetForce = tr.TargetForce
			r := readRules(t, "testdata/test-force.rules.yaml")
			err := g.ProcessRules(r)
			require.NoError(t, err)
			length := len(g.Stats.FilesTouched)
			require.Len(t, tr.FilesTouched, length)
			for _, file := range tr.FilesTouched {
				require.Contains(t, g.Stats.FilesTouched, file)
			}
		})
	}
}
