package filters

import (
	"strings"

	"github.com/iancoleman/strcase"
)

// SnakeCaseLower returns a string representation of the value in snake_case.
func SnakeCaseLower(s string) string {
	return strcase.ToSnake(s)
}

// SnakeCase returns a string representation of the value in snake_case.
func SnakeCase(s string) string {
	return strcase.ToSnake(s)
}

// First returns the first character of the value.
func SnakeCaseUpper(s string) string {
	return strcase.ToScreamingSnake(s)
}

// CamelCase returns a string representation of the value in CamelCase.
func CamelCase(s string) string {
	return strcase.ToCamel(s)
}

// CamelCase returns a string representation of the value in CamelCase.
func CamelCaseLower(s string) string {
	return strcase.ToLowerCamel(s)
}

// DotCaseLower returns a string representation of the value in dot.case
func DotCaseLower(s string) string {
	return strcase.ToDelimited(s, '.')
}

// DotCase returns a string representation of the value in dot.case
func DotCase(s string) string {
	return strcase.ToDelimited(s, '.')
}

// DotCaseUpper returns a string representation of the value in DOT.CASE
func DotCaseUpper(s string) string {
	return strcase.ToScreamingDelimited(s, '.', "", true)
}

// KebapCaseLower returns a string representation of the value in kebap-case.
func KebabCaseLower(s string) string {
	return strcase.ToKebab(s)
}

// KebabCase returns a string representation of the value in kebab-case.
func KebabCase(s string) string {
	return strcase.ToKebab(s)
}

// KebapCaseUpper returns a string representation of the value in KEBAP-CASE.
func KebabCaseUpper(s string) string {
	return strcase.ToScreamingKebab(s)
}

// PathCaseLower returns a string representation of the value in path/case.
func PathCaseLower(s string) string {
	return strcase.ToDelimited(s, '/')
}

// PathCase returns a string representation of the value in path/case.
func PathCase(s string) string {
	return strcase.ToDelimited(s, '/')
}

// PathCaseUpper returns a string representation of the value in PATH/CASE.
func PathCaseUpper(s string) string {
	return strcase.ToScreamingDelimited(s, '/', "", true)
}

// LowerCase returns a string representation of the value in lowercase.
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// UpperCase returns a string representation of the value in UPPER CASE.
func UpperCase(s string) string {
	return strings.ToUpper(s)
}

// UpperFirst returns a string representation of the value with the first character in UPPER CASE.
func UpperFirst(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

// LowerFirst returns a string representation of the value with the first character in lowercase.
func LowerFirst(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

// FirstChar returns the first character of the value.
func FirstChar(s string) string {
	return s[:1]
}

// FirstCharLower returns the first character of the value in lowercase.
func FirstCharLower(s string) string {
	return strings.ToLower(s[:1])
}

// FirstCharUpper returns the first character of the value in UPPER CASE.
func FirstCharUpper(s string) string {
	return strings.ToUpper(s[:1])
}
