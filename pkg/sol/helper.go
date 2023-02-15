package sol

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/helper"
)

func GetTemplateDir(rootDir string, template string) (string, error) {
	return FallbackDir(template, rootDir, cfg.TemplateCacheDir())
}

// FallbackDir returns the first dir that exists.
func FallbackDir(name string, dirs ...string) (string, error) {
	for _, dir := range dirs {
		if helper.IsDir(helper.Join(dir, name)) {
			return helper.Join(dir, name), nil
		}
	}
	return "", fmt.Errorf("dir %s not found", name)
}
