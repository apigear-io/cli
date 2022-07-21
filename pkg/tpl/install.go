package tpl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
)

// InstallTemplate clones a template using git from an url into a local directory.
func InstallTemplate(name string, repo string) error {
	// check if repo is a local dir or an url
	dir := filepath.Join(config.GetPackageDir(), name)
	_, err := os.Stat(dir)
	if err == nil {
		return fmt.Errorf("%s already exists", name)
	}
	log.Infof("clone template from %s into %s", repo, dir)
	err = git.Clone(repo, dir)
	if err != nil {
		log.Warnf("failed to clone template from %s into %s", repo, dir)
	}
	return err
}
