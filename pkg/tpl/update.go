package tpl

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

func UpdateRegistry() error {
	log.Info().Msgf("updating template registry")
	_, err := git.CloneOrPullRegistry()
	if err != nil {
		return err
	}
	if !helper.IsDir(config.GetTemplateRegistryDir()) {
		return fmt.Errorf("template registry not found")
	}
	fn := config.GetTempleRegistryFile()
	if !helper.IsFile(fn) {
		return fmt.Errorf("registry file %s not found", fn)
	}
	reg, err := git.ReadRegistry(fn)
	if err != nil {
		return err
	}
	log.Info().Msgf("found %d templates in registry", len(reg.Templates))
	for _, t := range reg.Templates {
		log.Info().Msgf("updating template %s", t.Name)
		err := git.UpdateRegistry(t)
		if err != nil {
			return err
		}
	}
	return nil
}
