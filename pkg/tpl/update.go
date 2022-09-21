package tpl

import (
	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

func UpdateTemplate(name string) error {
	// update template
	target := helper.Join(config.GetPackageDir(), name)
	log.Debugf("update template %s", name)
	out, err := git.Pull(target)
	if err != nil {
		log.Warnf("failed to update template %s", name)
	}
	log.Debugf("%s", out)
	return err
}
