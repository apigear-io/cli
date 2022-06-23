package gen

import (
	"apigear/pkg/log"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Writer struct {
	outputDir string
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

func (w *Writer) WriteFile(file string, bytes []byte, force bool) error {
	target := path.Join(w.outputDir, file)
	if !force {
		same, err := CompareContentWithFile(bytes, target)
		if err != nil {
			return fmt.Errorf("error comparing content to file %s: %s", target, err)
		}

		if same {
			log.Infof("skipping file %s", target)
			return nil
		}
	}
	log.Debug("write file ", target)
	dir := path.Dir(target)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %s", err)
	}
	return ioutil.WriteFile(target, bytes, 0644)
}

// NewFileWriter creates a new file writer
func NewFileWriter(outputDir string) IFileWriter {
	return &Writer{outputDir: outputDir}
}
