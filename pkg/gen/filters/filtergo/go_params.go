package filtergo

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func goParams(nodes []*model.TypedNode, prefix string) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("goParams called with nil nodes")
	}
	var inputs []string
	for _, p := range nodes {
		r, err := ToParamString(&p.Schema, p.Name, prefix)
		if err != nil {
			return "", err
		}
		inputs = append(inputs, r)
	}
	return strings.Join(inputs, ", "), nil
}
