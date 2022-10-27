package tpl

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/gitsight/go-vcsurl"
)

// ImportTemplate imports template from git repository into the cache
func ImportTemplate(url string) (*vcsurl.VCS, error) {
	vcs, err := git.GitUrlInfo(url)
	if err != nil {
		return nil, err
	}
	dst := helper.Join(config.TemplateCacheDir(), vcs.FullName)
	if helper.IsDir(dst) {
		return nil, fmt.Errorf("template %s already exists", vcs.FullName)
	}
	err = git.Clone(url, dst)
	if err != nil {
		return nil, err
	}
	return vcs, nil
}
