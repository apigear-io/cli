package filtergo

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func goVars(nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("goNames called with nil nodes")
	}
	names := make([]string, len(nodes))
	for idx, p := range nodes {
		name, err := ToVarString(p)
		if err != nil {
			return "", err
		}
		names[idx] = name
	}
	return strings.Join(names, ", "), nil
}

func goPublicVars(nodes []*model.TypedNode) (string, error) {
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
