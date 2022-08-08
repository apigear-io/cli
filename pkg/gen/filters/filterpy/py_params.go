package filterpy

import (
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func pyParams(nodes []*model.TypedNode) (string, error) {
	var inputs []string
	for _, n := range nodes {
		r, err := ToParamString(&n.Schema, n.Name, "")
		if err != nil {
			return "", err
		}
		inputs = append(inputs, r)
	}
	return strings.Join(inputs, ", "), nil
}
