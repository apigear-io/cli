package filterqt

import (
	"github.com/apigear-io/cli/pkg/gen/filters/common"
)


// qtNamespace returns the input string with style applied for creating a namespace name
func qtNamespace(name string) string {

	return common.SnakeCaseLower(name)
}

