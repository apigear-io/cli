package tpl

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

func ImportTemplate(src string) error {
	info, err := git.GitUrlInfo(src)
	if err != nil {
		return err
	}
	dst := helper.Join(config.TemplatesDir(), info.FullName)
	if helper.IsDir(dst) {
		return fmt.Errorf("template %s already exists", info.FullName)
	}
	return git.Clone(src, dst)
}
