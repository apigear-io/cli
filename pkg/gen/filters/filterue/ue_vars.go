package filterue

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func ueVars(prefix string, nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "xxx", fmt.Errorf("ueVars called with nil nodes")
	}
	names := make([]string, len(nodes))
	for idx, p := range nodes {
		name, err := ToVarString(prefix, p)
		if err != nil {
			return "xxx", err
		}
		names[idx] = name
	}
	return strings.Join(names, ", "), nil
}
