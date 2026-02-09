package filterjava

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func javaParams(prefix string, nodes []*apimodel.TypedNode) (string, error) {
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
