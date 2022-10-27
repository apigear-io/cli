package tpl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

// ListTemplates lists all templates in the cache
func ListTemplates() ([]*git.RepoInfo, error) {
	set := make(map[string]*git.RepoInfo)
	registry, err := SearchRegistry("")
	if err != nil {
		return []*git.RepoInfo{}, err
	}
	cached, err := ListCachedRepos()
	if err != nil {
		return []*git.RepoInfo{}, err
	}
	// merge
	for _, info := range registry {
		info.InCache = false
		info.InRegistry = true
		set[info.Name] = info
	}
	for _, info := range cached {
		if _, ok := set[info.Name]; ok {
			set[info.Name].InCache = true
		} else {
			info.InCache = true
			info.InRegistry = false
			set[info.Name] = info
		}
	}
	// convert to list
	var result []*git.RepoInfo
	for _, info := range set {
		result = append(result, info)
	}
	return result, nil
}

// ListTemplates lists all templates in the cache
func ListCachedRepos() ([]*git.RepoInfo, error) {
	// list all dirs in packageDir
	dir := cfg.TemplateCacheDir()
	// walk package dir to find a dir that contains a .git dir
	var infos []*git.RepoInfo
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
				infos = append(infos, &git.RepoInfo{
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
