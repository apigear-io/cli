package filterrust

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func rustVars(nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("goNames called with nil nodes")
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
