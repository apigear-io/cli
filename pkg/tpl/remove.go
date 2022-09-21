package tpl

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
)

func RemoveTemplate(name string) error {
	dir := config.GetPackageDir()
	log.Infof("remove template %s from %s", name, dir)
	// remove dir from packageDir
	// check if dir exists
	target := helper.Join(dir, name)
	_, err := os.Stat(target)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("template %s does not exist", name)
	}
	return os.RemoveAll(target)
}
