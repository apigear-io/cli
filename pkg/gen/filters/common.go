package filters

import (
	"strings"

	"github.com/iancoleman/strcase"
)

// SnakeCase returns a string representation of the value in snake_case.
func SnakeCase(s string) string {
	return strcase.ToSnake(s)
}

// CamelCase returns a string representation of the value in CamelCase.
func CamelCase(s string) string {
	return strcase.ToCamel(s)
}

// DotCase returns a string representation of the value in dot.case
func DotCase(s string) string {
	return strcase.ToDelimited(s, '.')
}

// LowerCamelCase returns a string representation of the value in lowerCamelCase.
func LowerCamelCase(s string) string {
	return strcase.ToLowerCamel(s)
}

// KebabCase returns a string representation of the value in kebab-case.
func KebabCase(s string) string {
	return strcase.ToKebab(s)
}

// PathCase returns a string representation of the value in path/case.
func PathCase(s string) string {
	return strcase.ToDelimited(s, '/')
}

// LowerCase returns a string representation of the value in lowercase.
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// UpperCase returns a string representation of the value in UPPER CASE.
func UpperCase(s string) string {
	return strings.ToUpper(s)
}
