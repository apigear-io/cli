package tpl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
)

// InstallTemplate clones a template using git from an url into a local directory.
func InstallTemplate(name string, repo string) error {
	// check if repo is a local dir or an url
	target := filepath.Join(GetPackageDir(), name)
	_, err := os.Stat(target)
	if err == nil {
		return fmt.Errorf("%s already exists", name)
	}
	log.Infof("clone template from %s into %s", repo, target)
	err = clone(repo, target)
	if err != nil {
		log.Warnf("failed to clone template from %s into %s", repo, target)
	}
	return err
}
