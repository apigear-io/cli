package filterts

import (
	"strings"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func tsParams(prefix string, nodes []*objmodel.TypedNode) (string, error) {
	var params []string
	for _, n := range nodes {
		r, err := ToParamString(&n.Schema, n.Name, prefix)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}
