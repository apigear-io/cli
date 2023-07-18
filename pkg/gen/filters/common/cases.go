package common

import (
	"strings"

	"github.com/ettle/strcase"
)

// SnakeCaseLower returns a string representation of the value in snake_case.
func SnakeCaseLower(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '_')
}

// SnakeTitleCase returns a string representation of the value in snake_case.
func SnakeTitleCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '_')
}

// First returns the first character of the value.
func SnakeUpperCase(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '_')
}

// CamelCase returns a string representation of the value in CamelCase.
func CamelLowerCase(s string) string {
	return strcase.ToCase(s, strcase.CamelCase, '\x00')
}

// CamelTitleCase returns a string representation of the value in CamelTitleCase.
func CamelTitleCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '\x00')
}

// CamelUpperCase returns a string representation of the value in CamelCase.
func CamelUpperCase(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '\x00')
}

// DotLowerCase returns a string representation of the value in dot.case
func DotLowerCase(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '.')
}

// DotTitleCase returns a string representation of the value in dot.case
func DotTitleCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '.')
}

// DotUpperCase returns a string representation of the value in DOT.CASE
func DotUpperCase(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '.')
}

// KebabLowerCase returns a string representation of the value in kebap-case.
func KebabLowerCase(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '-')
}

// KebabTitleCase returns a string representation of the value in kebab-case.
func KebabTitleCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '-')
}

// KebapCaseUpper returns a string representation of the value in KEBAP-CASE.
func KebabUpperCase(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '-')
}

// PathLowerCase returns a string representation of the value in path/case.
func PathLowerCase(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '/')
}

// PathTitleCase returns a string representation of the value in path/case.
func PathTitleCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '/')
}

// PathUpperCase returns a string representation of the value in PATH/CASE.
func PathUpperCase(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '/')
}

// LowerCase returns a string representation of the value in lowercase.
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// UpperCase returns a string representation of the value in UPPER CASE.
func UpperCase(s string) string {
	return strings.ToUpper(s)
}

// SpaceTitleCase returns a string representation of the value in Space Title Case.
func SpaceTitleCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, ' ')
}

// SpaceUpperCase returns a string representation of the value in SPACE UPPER CASE.
func SpaceUpperCase(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, ' ')
}

// SpaceLowerCase returns a string representation of the value in space lower case.
func SpaceLowerCase(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, ' ')
}
