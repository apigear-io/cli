package filtercpp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

// cast value to module and concat module name to cpp open namespaces
func nsOpen(node reflect.Value) (reflect.Value, error) {
	module := node.Interface().(*model.Module)
	if module == nil {
		return reflect.Value{}, fmt.Errorf("invalid module")
	}
	parts := []string{}
	for _, p := range strings.Split(module.Name, ".") {
		parts = append(parts, fmt.Sprintf("namespace %s {", p))
	}
	result := strings.Join(parts, " ")
	return reflect.ValueOf(result), nil
}

// cast value to module and concat module name to cpp closing namespaces
func nsClose(node reflect.Value) (reflect.Value, error) {
	module := node.Interface().(*model.Module)
	if module == nil {
		return reflect.Value{}, fmt.Errorf("invalid module")
	}
	parts := strings.Split(module.Name, ".")
	ns := ""
	result := ""
	for range parts {
		result += "} "
		ns = strings.Join(parts, "::")
	}
	result = fmt.Sprintf("%s// namespace %s", result, ns)
	return reflect.ValueOf(result), nil
}

// ns is a filter that concat module name to cpp namespaces
func ns(node reflect.Value) (reflect.Value, error) {
	module := node.Interface().(*model.Module)
	if module == nil {
		return reflect.Value{}, fmt.Errorf("invalid module")
	}
	parts := strings.Split(module.Name, ".")
	result := strings.Join(parts, "::")
	return reflect.ValueOf(result), nil
}
