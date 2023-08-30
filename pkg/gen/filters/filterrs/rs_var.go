package filterrs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToVarString(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return fmt.Sprintf("%s%s", prefix, common.SnakeCaseLower(node.Name)), nil
}

func rsVar(prefix string, node *model.TypedNode) (string, error) {
	return ToVarString(prefix, node)
}
