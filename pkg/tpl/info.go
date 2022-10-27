package tpl

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

// GetLocalTemplateInfo returns information about a template
// either from an installed of from a template registry
func GetLocalTemplateInfo(name string) (*git.RepoInfo, error) {
	dir := cfg.TemplateCacheDir()
	// get git info for template
	target := helper.Join(dir, name)
	if !helper.IsDir(target) {
		return nil, fmt.Errorf("template %s not found", name)
	}
	sha1, err := git.RepoLastCommit(target)
	if err != nil {
		log.Warn().Msgf("get git info for template %s", name)
	}
	url, err := git.RepoRemoteUrl(target)
	if err != nil {
		log.Warn().Msgf("get git info for template %s", name)
	}
	return &git.RepoInfo{
		Name:   strings.TrimSpace(name),
		Git:    strings.TrimSpace(url),
		Commit: strings.TrimSpace(sha1),
		Path:   strings.TrimSpace(target),
	}, err
}
