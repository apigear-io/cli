package tpl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
)

type TemplateInfo struct {
	Name   string
	URL    string
	Commit string
	Path   string
}

func ListTemplates() ([]TemplateInfo, error) {
	// list all dirs in packageDir
	dir := config.GetPackageDir()
	// walk package dir to find a dir that contains a .git dir
	var infos []TemplateInfo
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk template dir: %s", err)
		}
		if info.IsDir() && info.Name() != "." && info.Name() != ".." {
			if _, err := os.Stat(helper.Join(path, ".git")); err == nil {
				name, err := filepath.Rel(dir, path)
				if err != nil {
					return fmt.Errorf("failed to get relative path for %s", path)
				}
				infos = append(infos, TemplateInfo{
					Name: name,
					Path: path,
				})
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list templates: %s", err)
	}
	return infos, nil
}
