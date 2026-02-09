package filterts

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func tsVars(nodes []*apimodel.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("tsVars called with nil nodes")
	}
	names := make([]string, len(nodes))
	for idx, p := range nodes {
		name, err := ToVarString(p)
		if err != nil {
			return "xxx", err
		}
		names[idx] = name
	}
	return strings.Join(names, ", "), nil
}
