package tpl

import (
	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
)

// UpdateRegistry updates the local template registry
// The registry is a git repository that contains a list of templates
// and their versions.
func UpdateRegistry() error {
	log.Info().Msgf("updating template registry")
	src := config.RegistryUrl()
	dst := config.RegistryDir()
	helper.RemoveDir(dst)
	err := git.CloneOrPull(src, dst)
	if err != nil {
		return err
	}
	reg, err := ReadRegistry()
	if err != nil {
		return err
	}
	for _, entry := range reg.Entries {
		log.Info().Msgf("updating template %s", entry.Name)
		info, err := git.RemoteRepoInfo(entry.Git)
		entry.Versions = info.Versions
		entry.Latest = info.Latest
		if err != nil {
			return err
		}
	}
	return WriteRegistry(reg)
}
