package filterue

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func ueParams(prefix string, nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("goParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToParamString(&p.Schema, p.Name, prefix)
		if err != nil {
			return "", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}
