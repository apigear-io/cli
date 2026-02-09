package filterjni

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func jniJavaParams(prefix string, nodes []*objmodel.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("jniJavaParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToJniJavaParamString(&p.Schema, p.Name, prefix)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}
