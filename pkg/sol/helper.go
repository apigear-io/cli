package sol

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
)

func GetTemplateDir(rootDir string, template string) (string, error) {
	var templateDir string
	if helper.IsDir(filepath.Join(rootDir, template)) {
		templateDir = filepath.Join(rootDir, template)
	} else if helper.IsDir(filepath.Join(config.GetPackageDir(), template)) {
		templateDir = filepath.Join(config.GetPackageDir(), template)
	} else {
		return "", fmt.Errorf("template %s not found", template)
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
