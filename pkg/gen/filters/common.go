package filters

import (
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
)

// SnakeCase returns a string representation of the value in snake_case.
func SnakeCase(v reflect.Value) (reflect.Value, error) {
	s, ok := v.Interface().(fmt.Stringer)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected stringer, got %s", v.Type())
	}
	return reflect.ValueOf(strcase.ToCamel(s.String())), nil
}

// CamelCase returns a string representation of the value in CamelCase.
func CamelCase(v reflect.Value) (reflect.Value, error) {
	s, ok := v.Interface().(fmt.Stringer)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected stringer, got %s", v.Type())
	}
	return reflect.ValueOf(strcase.ToCamel(s.String())), nil
}

// DotCase returns a string representation of the value in dot.case
func DotCase(v reflect.Value) (reflect.Value, error) {
	s, ok := v.Interface().(fmt.Stringer)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected stringer, got %s", v.Type())
	}
	return reflect.ValueOf(strcase.ToDelimited(s.String(), '.')), nil
}

// LowerCamelCase returns a string representation of the value in lowerCamelCase.
func LowerCamelCase(v reflect.Value) (reflect.Value, error) {
	s, ok := v.Interface().(fmt.Stringer)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected stringer, got %s", v.Type())
	}
	return reflect.ValueOf(strcase.ToLowerCamel(s.String())), nil
}

// KebabCase returns a string representation of the value in kebab-case.
func KebabCase(v reflect.Value) (reflect.Value, error) {
	s, ok := v.Interface().(fmt.Stringer)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected stringer, got %s", v.Type())
	}
	return reflect.ValueOf(strcase.ToKebab(s.String())), nil
}
