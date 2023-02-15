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

// UpperFirst returns a string representation of the value with the first character in UPPER CASE.
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// LowerFirst returns a string representation of the value with the first character in lowercase.
func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// FirstChar returns the first character of the value.
func FirstChar(s string) string {
	if len(s) == 0 {
		return s
	}
	return s[:1]
}

// FirstCharLower returns the first character of the value in lowercase.
func FirstCharLower(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1])
}

// FirstCharUpper returns the first character of the value in UPPER CASE.
func FirstCharUpper(s string) string {
	if len(s) == 0 {
		return s
	}
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
	word := strcase.ToCase(WORDS[i], wc, '\x00')
	return fmt.Sprintf("%s%s%s", prefix, word, plural)
}

func IntToWordLower(i int, prefix string, postfix string) string {
	return IntToWord(i, prefix, postfix, strcase.LowerCase)
}

func IntToWordTitle(i int, prefix string, postfix string) string {
	return IntToWord(i, prefix, postfix, strcase.TitleCase)
}

func IntToWordUpper(i int, prefix string, postfix string) string {
	return IntToWord(i, prefix, postfix, strcase.UpperCase)
}

var plural = pluralize.NewClient()

func Pluralize(s string, i int) string {
	if i <= 1 {
		return s
	}
	return plural.Plural(s)
}
