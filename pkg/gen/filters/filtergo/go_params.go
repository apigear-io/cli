package filtergo

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func goParams(prefix string, nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("goParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToParamString(prefix, &p.Schema, p.Name)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}
