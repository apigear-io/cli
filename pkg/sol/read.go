package sol

import (
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/spec"

	"gopkg.in/yaml.v2"
)

func ReadSolutionDoc(file string) (*spec.SolutionDoc, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	doc := &spec.SolutionDoc{}
	err = yaml.Unmarshal(data, doc)
	if err != nil {
		return nil, err
	}
	// set the root dir
	if doc.RootDir == "" {
		doc.RootDir = filepath.Dir(file)
	}
	doc.RootDir, err = filepath.Abs(doc.RootDir)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
