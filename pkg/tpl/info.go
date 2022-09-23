package tpl

import (
	"strings"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

func GetInfo(name string) (TemplateInfo, error) {
	dir := config.GetPackageDir()
	// get git info for template
	target := helper.Join(dir, name)
	sha1, err := git.LastCommit(target)
	if err != nil {
		log.Warn().Msgf("failed to get git info for template %s", name)
	}
	url, err := git.RemoteUrl(target)
	if err != nil {
		log.Warn().Msgf("failed to get git info for template %s", name)
	}
	return TemplateInfo{
		Name:   strings.TrimSpace(name),
		URL:    strings.TrimSpace(url),
		Commit: strings.TrimSpace(sha1),
		Path:   strings.TrimSpace(target),
	}, err
}
