package tpl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
)

func RemoveTemplate(name string) error {
	dir := GetPackageDir()
	log.Infof("remove template %s from %s", name, dir)
	// remove dir from packageDir
	// check if dir exists
	target := filepath.Join(dir, name)
	_, err := os.Stat(target)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("template %s does not exist", name)
	}
	return os.RemoveAll(target)
}
