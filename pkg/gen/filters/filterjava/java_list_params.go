package filterjava

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func javaListParams(prefix string, nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("javaListParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToListParamString(prefix, &p.Schema, p.Name)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}
