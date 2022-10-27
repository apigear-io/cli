package tpl

import (
	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

// UpgradeTemplates upgrade templates from remote git repo
func UpgradeTemplates(names []string) error {
	log.Info().Msgf("update templates %s", names)
	for _, name := range names {
		// update template
		dst := helper.Join(config.TemplateCacheDir(), name)
		err := git.Pull(dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpgradeAllTemplates upgrade all templates from remote git repo
func UpgradeAllTemplates() error {
	// update all templates
	templates, err := ListTemplates()
	if err != nil {
		return err
	}
	names := []string{}
	for _, t := range templates {
		names = append(names, t.Name)
	}
	return UpgradeTemplates(names)
}
