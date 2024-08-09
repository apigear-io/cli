package common

import (
	"strings"
)

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

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func TrimSuffix(s, postfix string) string {
	return strings.TrimSuffix(s, postfix)
}

func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// SplitLast splits the string s at the last occurrence of sep and returns the result.
func SplitLast(s, sep string) string {
	parts := strings.Split(s, sep)
	return parts[len(parts)-1]
}

// SplitFirst splits the string s at the first occurrence of sep and returns the result.
func SplitFirst(s, sep string) string {
	return strings.Split(s, sep)[0]
}
