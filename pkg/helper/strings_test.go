package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"exact match", "hello", "hello", true},
		{"substring match", "Hello World", "world", true},
		{"case insensitive", "HELLO", "hello", true},
		{"not found", "hello", "xyz", false},
		{"empty substring", "hello", "", true},
		{"empty string", "", "hello", false},
		{"both empty", "", "", true},
		{"substring in middle", "the quick brown fox", "quick", true},
		{"case mixed", "HeLLo WoRLd", "HELLO world", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAbbreviate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple camel case", "HelloWorld", "HW"},
		{"with numbers", "API2Gateway", "A2G"},
		{"single word", "Simple", "S"},
		{"lowercase word", "simple", "S"},
		{"with spaces", "Hello World", "HW"},
		{"empty string", "", ""},
		{"numbers only", "123", ""},
		{"snake case", "hello_world", "HW"},
		{"kebab case", "hello-world", "HW"},
		{"pascal case", "HelloWorldAPI", "HWA"},
		{"multiple numbers", "API2Gateway3System", "A2G3S"},
		{"single letter", "a", "A"},
		{"uppercase", "HELLO", "H"},
		{"mixed case complex", "getHTTPResponseCode", "GHRC"},
		{"with underscore", "hello_World_Test", "HWT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abbreviate(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMapToArray(t *testing.T) {
	t.Run("string map", func(t *testing.T) {
		m := map[string]string{
			"a": "value1",
			"b": "value2",
			"c": "value3",
		}
		result := MapToArray(m)
		assert.Len(t, result, 3)
		assert.Contains(t, result, "value1")
		assert.Contains(t, result, "value2")
		assert.Contains(t, result, "value3")
	})

	t.Run("int map", func(t *testing.T) {
		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		result := MapToArray(m)
		assert.Len(t, result, 3)
		assert.Contains(t, result, 1)
		assert.Contains(t, result, 2)
		assert.Contains(t, result, 3)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]string{}
		result := MapToArray(m)
		assert.Empty(t, result)
	})

	t.Run("struct map", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		m := map[string]Person{
			"p1": {Name: "Alice", Age: 30},
			"p2": {Name: "Bob", Age: 25},
		}
		result := MapToArray(m)
		assert.Len(t, result, 2)
	})
}

func TestArrayToMap(t *testing.T) {
	t.Run("string array with key function", func(t *testing.T) {
		arr := []string{"hello", "world", "test"}
		m := make(map[string]string)
		keyFunc := func(s string) string {
			return s
		}
		result := ArrayToMap(m, arr, keyFunc)
		assert.Len(t, result, 3)
		assert.Equal(t, "hello", result["hello"])
		assert.Equal(t, "world", result["world"])
		assert.Equal(t, "test", result["test"])
	})

	t.Run("struct array with custom key", func(t *testing.T) {
		type Person struct {
			ID   string
			Name string
		}
		arr := []Person{
			{ID: "1", Name: "Alice"},
			{ID: "2", Name: "Bob"},
			{ID: "3", Name: "Charlie"},
		}
		m := make(map[string]Person)
		keyFunc := func(p Person) string {
			return p.ID
		}
		result := ArrayToMap(m, arr, keyFunc)
		assert.Len(t, result, 3)
		assert.Equal(t, "Alice", result["1"].Name)
		assert.Equal(t, "Bob", result["2"].Name)
		assert.Equal(t, "Charlie", result["3"].Name)
	})

	t.Run("empty array", func(t *testing.T) {
		arr := []string{}
		m := make(map[string]string)
		keyFunc := func(s string) string {
			return s
		}
		result := ArrayToMap(m, arr, keyFunc)
		assert.Empty(t, result)
	})

	t.Run("append to existing map", func(t *testing.T) {
		arr := []string{"new1", "new2"}
		m := map[string]string{
			"existing": "value",
		}
		keyFunc := func(s string) string {
			return s
		}
		result := ArrayToMap(m, arr, keyFunc)
		assert.Len(t, result, 3)
		assert.Equal(t, "value", result["existing"])
		assert.Equal(t, "new1", result["new1"])
		assert.Equal(t, "new2", result["new2"])
	})

	t.Run("duplicate keys overwrite", func(t *testing.T) {
		arr := []string{"key1", "key2", "key1"}
		m := make(map[string]string)
		keyFunc := func(s string) string {
			return s
		}
		result := ArrayToMap(m, arr, keyFunc)
		assert.Len(t, result, 2)
		assert.Equal(t, "key1", result["key1"])
	})
}
