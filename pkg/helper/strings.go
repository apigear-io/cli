package helper

import "strings"

// Contains checks if a string contains a substring case insensitive
func Contains(a string, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

// MapToArray converts a map to an array
func MapToArray[T any](m map[string]T) []T {
	var result []T
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// ArrayToMap converts an array to a map using a key function
func ArrayToMap[T any](m map[string]T, e []T, f func(T) string) map[string]T {
	for _, v := range e {
		m[f(v)] = v
	}
	return m
}
