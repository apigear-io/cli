package tpl

import (
	"apigear/pkg/log"
	"errors"
	"fmt"
	"os"
	"path"
)

func RemoveTemplate(name string) error {
	dir := GetPackageDir()
	log.Infof("remove template %s from %s", name, dir)
	// remove dir from packageDir
	// check if dir exists
	target := path.Join(dir, name)
	_, err := os.Stat(target)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("template %s does not exist", name)
	}
	return os.RemoveAll(target)
}
