package gen

import (
	"os"
	"testing"
	"text/template"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type MockFileWriter struct {
	Writes map[string]string
}

func (m *MockFileWriter) WriteFile(input []byte, target string, force bool) error {
	m.Writes[target] = string(input)
	return nil
}

func (w *MockFileWriter) CopyFile(source, target string, force bool) error {
	w.Writes[target] = source
	return nil
}

func NewMockFileWriter() *MockFileWriter {
	return &MockFileWriter{
		Writes: make(map[string]string),
	}
}

func readRules(t *testing.T, filename string) spec.RulesDoc {
	content, err := os.ReadFile(filename)
	assert.NoError(t, err)
	var file spec.RulesDoc
	err = yaml.Unmarshal(content, &file)
	assert.NoError(t, err)
	return file
}

func createGenerator(t *testing.T, w IFileWriter) *generator {
	var g = &generator{
		Writer:       w,
		Template:     template.New(""),
		System:       model.NewSystem("test"),
		UserForce:    false,
		TemplatesDir: "testdata/templates",
		OutputDir:    "testdata/output",
	}
	err := g.ParseTemplatesDir("testdata/templates")
	assert.NoError(t, err)
	return g
}
func TestEmptyRules(t *testing.T) {
	w := NewMockFileWriter()
	g := createGenerator(t, w)
	r := readRules(t, "testdata/empty.rules.yaml")
	assert.NoError(t, g.ProcessRulesDoc(r))
}

func TestHelloRules(t *testing.T) {
	w := NewMockFileWriter()
	g := createGenerator(t, w)
	r := readRules(t, "testdata/test.rules.yaml")
	err := g.ProcessRulesDoc(r)
	assert.NoError(t, err)
	assert.Contains(t, w.Writes, "system.txt")
}

func TestModules(t *testing.T) {
	w := NewMockFileWriter()
	g := createGenerator(t, w)
	r := readRules(t, "testdata/test.rules.yaml")
	err := g.ProcessRulesDoc(r)
	assert.NoError(t, err)
	assert.Contains(t, w.Writes, "system.txt")
}
