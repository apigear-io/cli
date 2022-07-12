package gen

import (
	"io/ioutil"
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

func (m *MockFileWriter) WriteFile(fn string, buf []byte, force bool) error {
	m.Writes[fn] = string(buf)
	return nil
}

func NewMockFileWriter() *MockFileWriter {
	return &MockFileWriter{
		Writes: make(map[string]string),
	}
}

func readRules(t *testing.T, filename string) spec.RulesDoc {
	content, err := ioutil.ReadFile(filename)
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
