package tests

import (
	"apigear/pkg/log"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

var TEST_DATA string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	TEST_DATA = filepath.Join(cwd, "testdata")

}

func exists(args []string, inputFile string) ([]byte, error) {
	buf := bytes.NewBufferString("")
	for _, arg := range args {
		_, err := os.Stat(arg)
		if err != nil {
			log.Errorf("%s", err)
			return nil, err
		}
		fmt.Fprintf(buf, "exists %s\n", arg)
	}
	return buf.Bytes(), nil
}

func copyTestData(srcDir, dstDir string, entries ...string) error {
	for _, entry := range entries {
		src := filepath.Join(srcDir, entry)
		dst := filepath.Join(dstDir, entry)
		buf, err := os.ReadFile(src)
		if err != nil {
			return err
		}
		err = os.MkdirAll(filepath.Dir(dst), 0755)
		if err != nil {
			return err
		}
		err = os.WriteFile(dst, buf, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
