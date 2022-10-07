package sol

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
)

func GetTemplateDir(rootDir string, template string) (string, error) {
	var templateDir string
	if helper.IsDir(helper.Join(rootDir, template)) {
		templateDir = helper.Join(rootDir, template)
	} else if helper.IsDir(helper.Join(config.TemplatesDir(), template)) {
		templateDir = helper.Join(config.TemplatesDir(), template)
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
