package foundation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinMaps(t *testing.T) {
	t.Run("join two maps", func(t *testing.T) {
		m1 := map[string]interface{}{
			"a": "value1",
			"b": "value2",
		}
		m2 := map[string]interface{}{
			"c": "value3",
			"d": "value4",
		}

		result := JoinMaps(m1, m2)

		assert.Len(t, result, 4)
		assert.Equal(t, "value1", result["a"])
		assert.Equal(t, "value2", result["b"])
		assert.Equal(t, "value3", result["c"])
		assert.Equal(t, "value4", result["d"])
	})

	t.Run("join multiple maps", func(t *testing.T) {
		m1 := map[string]interface{}{"a": 1}
		m2 := map[string]interface{}{"b": 2}
		m3 := map[string]interface{}{"c": 3}
		m4 := map[string]interface{}{"d": 4}

		result := JoinMaps(m1, m2, m3, m4)

		assert.Len(t, result, 4)
		assert.Equal(t, 1, result["a"])
		assert.Equal(t, 2, result["b"])
		assert.Equal(t, 3, result["c"])
		assert.Equal(t, 4, result["d"])
	})

	t.Run("overlapping keys - last wins", func(t *testing.T) {
		m1 := map[string]interface{}{
			"a": "first",
			"b": "value1",
		}
		m2 := map[string]interface{}{
			"a": "second",
			"c": "value2",
		}
		m3 := map[string]interface{}{
			"a": "third",
		}

		result := JoinMaps(m1, m2, m3)

		assert.Len(t, result, 3)
		assert.Equal(t, "third", result["a"])
		assert.Equal(t, "value1", result["b"])
		assert.Equal(t, "value2", result["c"])
	})

	t.Run("empty maps", func(t *testing.T) {
		m1 := map[string]interface{}{}
		m2 := map[string]interface{}{}

		result := JoinMaps(m1, m2)

		assert.Empty(t, result)
	})

	t.Run("no maps provided", func(t *testing.T) {
		result := JoinMaps()
		assert.Empty(t, result)
	})

	t.Run("single map", func(t *testing.T) {
		m := map[string]interface{}{
			"a": "value1",
			"b": "value2",
		}

		result := JoinMaps(m)

		assert.Len(t, result, 2)
		assert.Equal(t, "value1", result["a"])
		assert.Equal(t, "value2", result["b"])
	})

	t.Run("different value types", func(t *testing.T) {
		m1 := map[string]interface{}{
			"string": "text",
			"int":    42,
		}
		m2 := map[string]interface{}{
			"bool":  true,
			"float": 3.14,
		}

		result := JoinMaps(m1, m2)

		assert.Len(t, result, 4)
		assert.Equal(t, "text", result["string"])
		assert.Equal(t, 42, result["int"])
		assert.Equal(t, true, result["bool"])
		assert.Equal(t, 3.14, result["float"])
	})

	t.Run("nested maps", func(t *testing.T) {
		m1 := map[string]interface{}{
			"nested": map[string]interface{}{
				"a": "value1",
			},
		}
		m2 := map[string]interface{}{
			"other": "value2",
		}

		result := JoinMaps(m1, m2)

		assert.Len(t, result, 2)
		assert.Equal(t, "value2", result["other"])
		nested, ok := result["nested"].(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, "value1", nested["a"])
	})
}

func TestSelectValue(t *testing.T) {
	t.Run("select top level value", func(t *testing.T) {
		m := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}

		result := SelectValue(m, "key1")
		assert.Equal(t, "value1", result)
	})

	t.Run("select nested value", func(t *testing.T) {
		m := map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "nested-value",
			},
		}

		result := SelectValue(m, "level1.level2")
		assert.Equal(t, "nested-value", result)
	})

	t.Run("select deeply nested value", func(t *testing.T) {
		m := map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": map[string]interface{}{
						"d": "deep-value",
					},
				},
			},
		}

		result := SelectValue(m, "a.b.c.d")
		assert.Equal(t, "deep-value", result)
	})

	t.Run("key not found returns nil", func(t *testing.T) {
		m := map[string]interface{}{
			"key1": "value1",
		}

		result := SelectValue(m, "nonexistent")
		assert.Nil(t, result)
	})

	t.Run("nested key not found returns nil", func(t *testing.T) {
		m := map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": "value",
			},
		}

		result := SelectValue(m, "level1.nonexistent")
		assert.Nil(t, result)
	})

	t.Run("partial path exists", func(t *testing.T) {
		m := map[string]interface{}{
			"level1": "not-a-map",
		}

		// Trying to access level1.level2 when level1 is not a map
		result := SelectValue(m, "level1.level2")
		assert.Equal(t, "not-a-map", result)
	})

	t.Run("select map value", func(t *testing.T) {
		inner := map[string]interface{}{
			"nested": "value",
		}
		m := map[string]interface{}{
			"key": inner,
		}

		result := SelectValue(m, "key")
		assert.Equal(t, inner, result)
	})

	t.Run("empty selector", func(t *testing.T) {
		m := map[string]interface{}{
			"key": "value",
		}

		result := SelectValue(m, "")
		// Empty selector should return nil as it splits to [""]
		assert.Nil(t, result)
	})

	t.Run("select with different value types", func(t *testing.T) {
		m := map[string]interface{}{
			"string": "text",
			"int":    42,
			"bool":   true,
			"float":  3.14,
		}

		assert.Equal(t, "text", SelectValue(m, "string"))
		assert.Equal(t, 42, SelectValue(m, "int"))
		assert.Equal(t, true, SelectValue(m, "bool"))
		assert.Equal(t, 3.14, SelectValue(m, "float"))
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]interface{}{}
		result := SelectValue(m, "key")
		assert.Nil(t, result)
	})

	t.Run("access through multiple levels", func(t *testing.T) {
		m := map[string]interface{}{
			"user": map[string]interface{}{
				"profile": map[string]interface{}{
					"name": "John Doe",
					"age":  30,
				},
			},
		}

		assert.Equal(t, "John Doe", SelectValue(m, "user.profile.name"))
		assert.Equal(t, 30, SelectValue(m, "user.profile.age"))
	})
}
