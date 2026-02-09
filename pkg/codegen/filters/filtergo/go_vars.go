package filtergo

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func goVars(nodes []*apimodel.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("goNames called with nil nodes")
	}
	names := make([]string, len(nodes))
	for idx, p := range nodes {
		name, err := ToVarString(p)
		if err != nil {
			return "xxx", err
		}
		names[idx] = name
	}
	return strings.Join(names, ", "), nil
}

func goPublicVars(nodes []*apimodel.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("goNames called with nil nodes")
	}
	names := make([]string, len(nodes))
	for idx, p := range nodes {
		name, err := ToPublicVarString(p)
		if err != nil {
			return "", err
		}
		names[idx] = name
	}
	return strings.Join(names, ", "), nil
}
