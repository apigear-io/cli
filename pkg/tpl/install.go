package tpl

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

// InstallTemplate clones a template using git from an url into a local directory.
func InstallTemplate(name string, repo string) error {
	// check if repo is a local dir or an url
	dir := helper.Join(config.GetPackageDir(), name)
	_, err := os.Stat(dir)
	if err == nil {
		return fmt.Errorf("%s already exists", name)
	}
	log.Info().Msgf("clone template from %s into %s", repo, dir)
	err = git.Clone(repo, dir)
	if err != nil {
		log.Warn().Msgf("failed to clone template from %s into %s", repo, dir)
	}
	return err
}
