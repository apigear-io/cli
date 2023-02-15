package sol

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/helper"
)

func resolveTemplateDir(rootDir string, template string) (string, error) {
	var templateDir string
	if helper.IsDir(helper.Join(rootDir, template)) {
		templateDir = helper.Join(rootDir, template)
	} else if helper.IsDir(helper.Join(cfg.TemplateCacheDir(), template)) {
		templateDir = helper.Join(cfg.TemplateCacheDir(), template)
	} else {
		return "", fmt.Errorf("template dir %s not found", template)
	}
	return templateDir, nil
}

func hasExtension(file string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}
	return false
}
