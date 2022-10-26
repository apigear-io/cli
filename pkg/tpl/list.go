package tpl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
)

// ListTemplates lists all templates in the cache
func ListTemplates() ([]TemplateInfo, error) {
	// list all dirs in packageDir
	dir := config.TemplatesDir()
	// walk package dir to find a dir that contains a .git dir
	var infos []TemplateInfo
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk template dir: %s", err)
		}
		if info.IsDir() && info.Name() != "." && info.Name() != ".." {
			if helper.IsDir(helper.Join(path, ".git")) {
				name, err := filepath.Rel(dir, path)
				if err != nil {
					return fmt.Errorf("get relative path for %s", path)
				}
				infos = append(infos, TemplateInfo{
					Name:    name,
					Path:    path,
					InCache: true,
				})
				// no need to traverse into this dir
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list templates: %s", err)
	}
	return infos, nil
}
