package gen

import (
	"io/ioutil"
	"objectapi/pkg/model"
	"objectapi/pkg/spec"
	"testing"
	"text/template"

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

func createGenerator(w IFileWriter) *Generator {
	var processor = &Generator{
		Writer:    w,
		Template:  template.New(""),
		System:    model.NewSystem("test"),
		UserForce: false,
	}
	return processor
}
func TestEmptyRules(t *testing.T) {
	w := NewMockFileWriter()
	g := createGenerator(w)
	r := readRules(t, "testdata/empty.rules.yaml")
	g.ProcessRulesDoc(r)
}

func TestHelloRules(t *testing.T) {
	w := NewMockFileWriter()
	g := createGenerator(w)
	r := readRules(t, "testdata/hello.rules.yaml")
	g.ProcessRulesDoc(r)
	assert.Equal(t, "hello.txt", w.Writes["hello.txt"])
}
