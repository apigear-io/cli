package filterpy

import (
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func pyParams(prefix string, nodes []*model.TypedNode) (string, error) {
	inputs := []string{"self"}
	for _, n := range nodes {
		r, err := ToParamString(&n.Schema, n.Name, prefix)
		if err != nil {
			return "", err
		}
		inputs = append(inputs, r)
	}
	return strings.Join(inputs, ", "), nil
}

func pyFuncParams(prefix string, nodes []*model.TypedNode) (string, error) {
	inputs := []string{}
	for _, n := range nodes {
		r, err := ToParamString(&n.Schema, n.Name, prefix)
		if err != nil {
			return "", err
		}
		inputs = append(inputs, r)
	}
	return strings.Join(inputs, ", "), nil
}
