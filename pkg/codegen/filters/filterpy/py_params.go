package filterpy

import (
	"strings"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func pyParams(prefix string, nodes []*apimodel.TypedNode) (string, error) {
	params := []string{"self"}
	for _, n := range nodes {
		r, err := ToParamString(&n.Schema, n.Name, prefix)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}

func pyFuncParams(prefix string, nodes []*apimodel.TypedNode) (string, error) {
	params := []string{}
	for _, n := range nodes {
		r, err := ToParamString(&n.Schema, n.Name, prefix)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}
