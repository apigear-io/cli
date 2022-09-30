package tpl

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/gitsight/go-vcsurl"
)

func ImportTemplate(src string) (*vcsurl.VCS, error) {
	info, err := git.GitUrlInfo(src)
	if err != nil {
		return nil, err
	}
	dst := helper.Join(config.TemplatesDir(), info.FullName)
	if helper.IsDir(dst) {
		return nil, fmt.Errorf("template %s already exists", info.FullName)
	}
	err = git.Clone(src, dst)
	if err != nil {
		return nil, err
	}
	return info, nil
}
