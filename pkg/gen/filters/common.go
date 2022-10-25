package filters

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"
	"github.com/gertd/go-pluralize"
)

// SnakeCaseLower returns a string representation of the value in snake_case.
func SnakeCaseLower(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '_')
}

// SnakeCase returns a string representation of the value in snake_case.
func SnakeCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '_')
}

// First returns the first character of the value.
func SnakeCaseUpper(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '_')
}

// CamelCase returns a string representation of the value in CamelCase.
func CamelCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '\x00')
}

// CamelCase returns a string representation of the value in CamelCase.
func CamelCaseLower(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '\x00')
}

func CamelCaseUpper(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '\x00')
}

// DotCaseLower returns a string representation of the value in dot.case
func DotCaseLower(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '.')
}

// DotCase returns a string representation of the value in dot.case
func DotCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '.')
}

// DotCaseUpper returns a string representation of the value in DOT.CASE
func DotCaseUpper(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '.')
}

// KebapCaseLower returns a string representation of the value in kebap-case.
func KebabCaseLower(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '-')
}

// KebabCase returns a string representation of the value in kebab-case.
func KebabCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '-')
}

// KebapCaseUpper returns a string representation of the value in KEBAP-CASE.
func KebabCaseUpper(s string) string {
	return strcase.ToCase(s, strcase.UpperCase, '-')
}

// PathCaseLower returns a string representation of the value in path/case.
func PathCaseLower(s string) string {
	return strcase.ToCase(s, strcase.LowerCase, '/')
}

// PathCase returns a string representation of the value in path/case.
func PathCase(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '/')
}

// PathCaseUpper returns a string representation of the value in PATH/CASE.
func PathCaseUpper(s string) string {
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

func Join(sep string, a []string) string {
	return strings.Join(a, sep)
}

func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func TrimSuffix(s, postfix string) string {
	return strings.TrimSuffix(s, postfix)
}

func Replace(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}

func NewLine() string {
	return "\n"
}

var WORDS = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty",
}

func IntToWord(i int, prefix string, postfix string, wc strcase.WordCase) string {
	if i <= 0 || i >= len(WORDS) {
		return ""
	}
	plural := Pluralize(postfix, i)
	word := WORDS[i]
	word = strcase.ToCase(word, wc, '\x00')
	return fmt.Sprintf("%s%s%s", prefix, word, plural)
}

var plural = pluralize.NewClient()

func Pluralize(s string, i int) string {
	if i <= 1 {
		return s
	}
	return plural.Plural(s)
}
