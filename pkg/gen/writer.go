package gen

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/helper"
)

type Writer struct {
	sourceDir string
	targetDir string
}

func CompareContentWithFile(sourceBytes []byte, target string) (bool, error) {
	if _, err := os.Stat(target); os.IsNotExist(err) {
		return false, nil
	}
	var targetBytes, err = os.ReadFile(target)
	if err != nil {
		return false, err
	}
	return md5.Sum(sourceBytes) == md5.Sum(targetBytes), nil
}

func (w *Writer) WriteFile(input []byte, target string, force bool) error {
	target = filepath.Join(w.targetDir, target)
	if !force {
		same, err := CompareContentWithFile(input, target)
		if err != nil {
			return fmt.Errorf("error comparing content to file %s: %s", target, err)
		}

		if same {
			log.Infof("skipping file %s", target)
			return nil
		}
	}
	log.Debug("write file ", target)
	dir := filepath.Dir(target)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %s", err)
	}
	return os.WriteFile(target, input, 0644)
}

func (w Writer) CopyFile(source, target string, force bool) error {
	target = filepath.Join(w.targetDir, target)
	source = filepath.Join(w.sourceDir, source)
	return helper.CopyFile(source, target)
}

// NewFileWriter creates a new file writer
func NewFileWriter(sourceDir, targetDir string) IFileWriter {
	return &Writer{
		sourceDir: sourceDir,
		targetDir: targetDir,
	}
}
