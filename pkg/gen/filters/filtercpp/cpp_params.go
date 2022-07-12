package filtercpp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func cppParams(node reflect.Value) (reflect.Value, error) {
	m, ok := node.Interface().(*model.Method)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected method, got %s", node.Type())
	}
	var inputs []string
	for _, p := range m.Inputs {
		inputs = append(inputs, ToParamString(p.GetSchema(), p.GetName()))
	}
	return reflect.ValueOf(strings.Join(inputs, ", ")), nil
}
