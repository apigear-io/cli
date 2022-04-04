package gen

import (
	"crypto/md5"
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

func CompareContentWithFile(content string, target string) (bool, error) {
	var bytes, err = ioutil.ReadFile(target)
	if err != nil {
		return false, err
	}
	return md5.Sum([]byte(content)) == md5.Sum(bytes), nil
}

func (w *Writer) WriteFile(file string, content string) error {
	target := path.Join(w.outputDir, file)
	same, err := CompareContentWithFile(content, target)
	if err != nil {
		return fmt.Errorf("error comparing content to file %s: %s", target, err)
	}
	if same {
		log.Infof("skipping file %s", target)
		return nil
	}
	log.Info("write file ", target)
	dir := path.Dir(target)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %s", err)
	}
	return ioutil.WriteFile(target, []byte(content), 0644)
}

func NewFileWriter(outputDir string) FileWriter {
	return &Writer{outputDir: outputDir}
}
