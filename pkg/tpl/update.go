package tpl

import (
	"apigear/pkg/log"
	"path"
)

func UpdateTemplate(name string) error {
	// update template
	target := path.Join(GetPackageDir(), name)
	log.Infof("update template %s", name)
	err := ExecGit([]string{"pull"}, target)
	if err != nil {
		log.Warnf("failed to update template %s", name)

	}
	return err
}
