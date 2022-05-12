package filtergo

import (
	"objectapi/pkg/model"
	"strings"
)

// func goParams(node reflect.Value) (reflect.Value, error) {
// 	types, ok := node.Interface().([]*model.TypedNode)
// 	if !ok {
// 		return reflect.Value{}, fmt.Errorf("expected array of type nodes, got %s", node.Type())
// 	}
// 	var inputs []string
// 	for _, p := range types {
// 		inputs = append(inputs, ToParamString(p.GetSchema(), p.GetName()))
// 	}
// 	return reflect.ValueOf(strings.Join(inputs, ", ")), nil
// }

func goParams(nodes []*model.TypedNode) string {
	var inputs []string
	for _, p := range nodes {
		inputs = append(inputs, ToParamString(p.GetSchema(), p.GetName()))
	}
	return strings.Join(inputs, ", ")
}
