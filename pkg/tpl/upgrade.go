package tpl

import (
	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

func UpgradeTemplate(name string) error {
	// update template
	target := helper.Join(config.GetPackageDir(), name)
	log.Debug().Msgf("update template %s", name)
	out, err := git.Pull(target)
	if err != nil {
		log.Warn().Msgf("failed to update template %s", name)
	}
	log.Debug().Msgf("%s", out)
	return err
}
