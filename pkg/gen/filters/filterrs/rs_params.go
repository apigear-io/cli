package filterrs

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func rsParams(prefixVarName string, prefixComplexType string, separator string, nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("rsParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToParamString(prefixVarName, prefixComplexType, &p.Schema, p)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, separator), nil
}
