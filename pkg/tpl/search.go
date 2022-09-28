package tpl

import (
	"strings"

	"github.com/apigear-io/cli/pkg/git"
)

func SearchRegistry(pattern string) ([]*git.RemoteInfo, error) {
	// search templates
	reg, err := ReadRegistry()
	if err != nil {
		return nil, err
	}
	if pattern == "" {
		return reg.Entries, nil
	}
	var result []*git.RemoteInfo
	for _, info := range reg.Entries {
		s := strings.ToLower(info.Name)
		sub := strings.ToLower(pattern)
		if strings.Contains(s, sub) {
			result = append(result, info)
		}
	}
	return result, nil
}
