package filterue

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ueParams(prefix string, nodes []*objmodel.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("useParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToParamString(&p.Schema, p.Name, prefix)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}
