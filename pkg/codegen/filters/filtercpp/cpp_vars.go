package filtercpp

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func cppVars(nodes []*objmodel.TypedNode) (string, error) {
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
