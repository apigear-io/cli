package filterrust

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func rustParams(prefixComplexType string, separator string, nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("rustParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToParamString(prefixComplexType, &p.Schema, p)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, separator), nil
}
