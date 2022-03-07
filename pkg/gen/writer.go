package gen

import (
	"fmt"
	"io/ioutil"
	"objectapi/pkg/logger"
	"os"
	"path"
)

type Writer struct {
	outputDir string
}

var log = logger.Get()

func (w *Writer) WriteFile(file string, content string) error {
	target := path.Join(w.outputDir, file)
	log.Info("write file ", target)
	dir := path.Dir(target)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %s", err)
	}
	return ioutil.WriteFile(target, []byte(content), 0644)
}

func NewFileWriter(outputDir string) FileWriter {
	return &Writer{outputDir: outputDir}
}
