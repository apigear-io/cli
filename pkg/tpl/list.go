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
	for _, remoteInfo := range registry {
		remoteInfo.InCache = false
		remoteInfo.InRegistry = true
		set[remoteInfo.Name] = remoteInfo
	}
	for _, cachedInfo := range cached {
		remoteInfo, ok := set[cachedInfo.Name]
		if ok {
			remoteInfo.InCache = true
			remoteInfo.InRegistry = true
			remoteInfo.Tag = cachedInfo.Tag
			remoteInfo.Commit = cachedInfo.Commit
		} else {
			cachedInfo.InCache = true
			cachedInfo.InRegistry = false
			set[cachedInfo.Name] = cachedInfo
		}
	}
	// convert to list
	var result []*git.RepoInfo
	for _, info := range set {
		result = append(result, info)
	}
	// sort by name
	git.SortRepoInfo(result)
	return result, nil
}

func SearchTemplates(pattern string) ([]*git.RepoInfo, error) {
	result, err := ListTemplates()
	if err != nil {
		return []*git.RepoInfo{}, err
	}
	var filtered []*git.RepoInfo
	for _, info := range result {
		if helper.Contains(info.Name, pattern) {
			filtered = append(filtered, info)
		}
	}
	return filtered, nil
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
		if git.IsLocalGitRepo(path) {
			info, err := git.LocalRepoInfo(dir, path)
			if err != nil {
				return fmt.Errorf("get local repo info: %s", err)
			}
			infos = append(infos, info)
			// no need to traverse into this dir
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list templates: %s", err)
	}
	git.SortRepoInfo(infos)
	return infos, nil
}
