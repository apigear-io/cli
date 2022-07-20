package tpl

import (
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/log"
)

func UpdateTemplate(name string) error {
	// update template
	target := filepath.Join(config.GetPackageDir(), name)
	log.Debugf("update template %s", name)
	out, err := git.Pull(target)
	if err != nil {
		log.Warnf("failed to update template %s", name)
	}
	log.Debugf("%s", out)
	return err
}
