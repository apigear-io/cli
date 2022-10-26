package tpl

import (
	"github.com/apigear-io/cli/pkg/git"
	"github.com/apigear-io/cli/pkg/helper"
)

// SearchRegistry searches registry for templates containing a substring

func SearchRegistry(substring string) ([]*git.RemoteInfo, error) {
	// search templates
	reg, err := ReadRegistry()
	if err != nil {
		return nil, err
	}
	if substring == "" {
		// return all
		return reg.Entries, nil
	}
	var result []*git.RemoteInfo
	for _, info := range reg.Entries {
		if helper.Contains(info.Name, substring) {
			result = append(result, info)
		}
	}
	return result, nil
}
