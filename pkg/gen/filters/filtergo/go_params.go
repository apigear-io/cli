package filtergo

import (
	"apigear/pkg/model"
	"strings"
)

func goParams(nodes []*model.TypedNode, prefix string) string {
	var inputs []string
	for _, p := range nodes {
		inputs = append(inputs, ToParamString(p.GetSchema(), p.GetName(), prefix))
	}
	return strings.Join(inputs, ", ")
}
