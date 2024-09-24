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
	outDir := t.TempDir()
	opts := Options{
		System:       model.NewSystem("test"),
		Force:        false,
		TemplatesDir: "testdata/templates",
		OutputDir:    outDir,
		Features:     []string{"all"},
	}
	g, err := New(opts)
	require.NoError(t, err)
	err = g.ParseTemplatesDir("testdata/templates")
	require.NoError(t, err)
	return g
}

func createMockGenerator(t *testing.T, tplDir string, features []string) (*generator, *MockOutput) {
	out := NewMockOutput()
	opts := Options{
		System:       model.NewSystem("test"),
		Force:        true,
		TemplatesDir: helper.Join(tplDir, "templates"),
		OutputDir:    "testdata/output",
		Features:     features,
		Output:       out,
	}
	g, err := New(opts)
	require.NoError(t, err)
	rules := readRules(t, helper.Join(tplDir, "rules.yaml"))
	err = g.ProcessRules(rules)
	require.NoError(t, err)
	return g, out
}

func TestEmptyRules(t *testing.T) {
	t.Parallel()
	g := createGenerator(t)
	doc, err := ReadRulesDoc("testdata/empty.rules.yaml")
	require.NoError(t, err)
	require.NoError(t, g.ProcessRules(doc))
}

func TestHelloRules(t *testing.T) {
	t.Parallel()
	g := createGenerator(t)
	g.opts.Force = true
	r := readRules(t, "testdata/test.rules.yaml")
	err := g.ProcessRules(r)
	require.NoError(t, err)
	require.Len(t, g.Stats.FilesTouched, 1)
}

func TestForce(t *testing.T) {
	t.Parallel()
	g := createGenerator(t)
	g.opts.Force = true
	r := readRules(t, "testdata/test-preserve.rules.yaml")
	err := g.ProcessRules(r)
	require.NoError(t, err)
	length := len(g.Stats.FilesTouched)
	require.Equal(t, 2, length)
}

func TestDocumentPreserve(t *testing.T) {
	t.Parallel()
	tt := []struct {
		Name           string
		FilesFirstRun  []string
		FilesSecondRun []string
	}{
		{"preserve", []string{"system-preserve.txt", "system.txt"}, []string{}},
	}

	for _, tr := range tt {
		t.Run(tr.Name, func(t *testing.T) {
			g := createGenerator(t)
			r := readRules(t, "testdata/test-preserve.rules.yaml")
			// first run
			err := g.ProcessRules(r)
			require.NoError(t, err)
			require.Len(t, g.Stats.FilesTouched, len(tr.FilesFirstRun))
			for _, file := range tr.FilesFirstRun {
				target := helper.Join(g.opts.OutputDir, file)
				require.Contains(t, g.Stats.FilesTouched, target)
			}
			// second run
			err = g.ProcessRules(r)
			require.NoError(t, err)
			require.Len(t, g.Stats.FilesTouched, len(tr.FilesSecondRun))
			for _, file := range tr.FilesSecondRun {
				target := helper.Join(g.opts.OutputDir, file)
				require.Contains(t, g.Stats.FilesTouched, target)
			}
		})
	}
}
