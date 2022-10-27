package tpl

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/helper"
)

// RemoveTemplate removes template by name from the cache
func RemoveTemplate(name string) error {
	dir := cfg.TemplateCacheDir()
	log.Info().Msgf("remove template %s from %s", name, dir)
	// remove dir from packageDir
	// check if dir exists
	target := helper.Join(dir, name)
	_, err := os.Stat(target)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("template %s does not exist", name)
	}
	return os.RemoveAll(target)
}
