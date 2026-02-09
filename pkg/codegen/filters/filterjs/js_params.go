package filterjs

import (
	"strings"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func jsParams(prefix string, nodes []*apimodel.TypedNode) (string, error) {
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
