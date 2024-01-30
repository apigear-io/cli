package gen

import (
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/helper"
)

type OutputWriter interface {
	Write(input []byte, target string) error
	Copy(source, target string) error
	Compare(input []byte, target string) (bool, error)
}

type fsWriter struct {
}

var _ OutputWriter = (*fsWriter)(nil)

func NewOutputWriter() *fsWriter {
	return &fsWriter{}
}

func (f *fsWriter) Write(input []byte, target string) error {

	// write document to file system
	dir := filepath.Dir(target)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(target, input, 0644)
}

func (f *fsWriter) Copy(source, target string) error {
	return helper.CopyFile(source, target)
}

func (f *fsWriter) Compare(input []byte, target string) (bool, error) {
	return CompareContentWithFile(input, target)
}

type MockOutput struct {
	Writes   map[string]string
	Copies   map[string]string
	Compares map[string]bool
}

var _ OutputWriter = (*MockOutput)(nil)

func NewMockOutput() *MockOutput {
	return &MockOutput{
		Writes:   make(map[string]string),
		Copies:   make(map[string]string),
		Compares: make(map[string]bool),
	}
}

func (m *MockOutput) Write(input []byte, target string) error {
	m.Writes[target] = string(input)
	return nil
}

func (m *MockOutput) Copy(source, target string) error {
	m.Copies[source] = target
	return nil
}

func (m *MockOutput) Compare(input []byte, target string) (bool, error) {
	return m.Compares[target], nil
}
