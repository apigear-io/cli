package tpl

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
)

// GetLocalTemplateInfo returns information about a template
// either from an installed of from a template registry
func GetLocalTemplateInfo(name string) (TemplateInfo, error) {
	dir := config.TemplatesDir()
	// get git info for template
	target := helper.Join(dir, name)
	if !helper.IsDir(target) {
		return TemplateInfo{}, fmt.Errorf("template %s not found", name)
	}
	sha1, err := git.RepoLastCommit(target)
	if err != nil {
		log.Warn().Msgf("failed to get git info for template %s", name)
	}
	url, err := git.RepoRemoteUrl(target)
	if err != nil {
		log.Warn().Msgf("failed to get git info for template %s", name)
	}
	return TemplateInfo{
		Name:   strings.TrimSpace(name),
		Git:    strings.TrimSpace(url),
		Commit: strings.TrimSpace(sha1),
		Path:   strings.TrimSpace(target),
	}, err
}
