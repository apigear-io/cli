package tpl

import (
	"os"

	"github.com/apigear-io/cli/pkg/helper"
)

type TemplateInfo struct {
	Rules string
	Files []string
}

func Info(dir string) (*TemplateInfo, error) {
	info := &TemplateInfo{}
	// read rules.yaml
	rules, err := os.ReadFile(helper.Join(dir, "rules.yaml"))
	if err != nil {
		return nil, err
	}
	info.Rules = string(rules)
	// read files
	files, err := os.ReadDir(helper.Join(dir, "templates"))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		info.Files = append(info.Files, file.Name())
	}
	return info, nil
}
