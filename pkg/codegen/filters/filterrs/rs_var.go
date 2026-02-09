package filterrs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToVarString(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rsVar node is nil")
	}
	return fmt.Sprintf("%s%s", prefix, common.SnakeCaseLower(node.Name)), nil
}

func rsVar(prefix string, node *apimodel.TypedNode) (string, error) {
	return ToVarString(prefix, node)
}
