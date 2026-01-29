package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	t.Run("iterate over strings", func(t *testing.T) {
		items := []string{"apple", "banana", "cherry"}
		iter := NewIterator(items)

		// First item
		assert.True(t, iter.HasNext())
		item, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, "apple", item)

		// Second item
		assert.True(t, iter.HasNext())
		item, ok = iter.Next()
		assert.True(t, ok)
		assert.Equal(t, "banana", item)

		// Third item
		assert.True(t, iter.HasNext())
		item, ok = iter.Next()
		assert.True(t, ok)
		assert.Equal(t, "cherry", item)

		// No more items
		assert.False(t, iter.HasNext())
		item, ok = iter.Next()
		assert.False(t, ok)
		assert.Equal(t, "", item)
	})

	t.Run("iterate over integers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		iter := NewIterator(items)

		count := 0
		for iter.HasNext() {
			item, ok := iter.Next()
			assert.True(t, ok)
			assert.Equal(t, count+1, item)
			count++
		}

		assert.Equal(t, 5, count)
	})

	t.Run("empty slice", func(t *testing.T) {
		items := []string{}
		iter := NewIterator(items)

		assert.False(t, iter.HasNext())
		item, ok := iter.Next()
		assert.False(t, ok)
		assert.Equal(t, "", item)
	})

	t.Run("single item", func(t *testing.T) {
		items := []string{"only"}
		iter := NewIterator(items)

		assert.True(t, iter.HasNext())
		item, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, "only", item)

		assert.False(t, iter.HasNext())
	})

	t.Run("iterate with struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		items := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}

		iter := NewIterator(items)

		count := 0
		for iter.HasNext() {
			person, ok := iter.Next()
			assert.True(t, ok)
			assert.NotEmpty(t, person.Name)
			assert.Greater(t, person.Age, 0)
			count++
		}

		assert.Equal(t, 3, count)
	})

	t.Run("multiple Next calls after exhaustion", func(t *testing.T) {
		items := []int{1}
		iter := NewIterator(items)

		// First Next - valid
		item, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, 1, item)

		// Multiple Next calls after exhaustion
		for i := 0; i < 5; i++ {
			item, ok := iter.Next()
			assert.False(t, ok)
			assert.Equal(t, 0, item)
		}
	})

	t.Run("HasNext doesn't advance iterator", func(t *testing.T) {
		items := []string{"first", "second"}
		iter := NewIterator(items)

		// Multiple HasNext calls
		for i := 0; i < 3; i++ {
			assert.True(t, iter.HasNext())
		}

		// Should still get first item
		item, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, "first", item)
	})

	t.Run("iterate with pointers", func(t *testing.T) {
		type Data struct {
			Value int
		}

		items := []*Data{
			{Value: 10},
			{Value: 20},
			{Value: 30},
		}

		iter := NewIterator(items)

		sum := 0
		for iter.HasNext() {
			data, ok := iter.Next()
			assert.True(t, ok)
			assert.NotNil(t, data)
			sum += data.Value
		}

		assert.Equal(t, 60, sum)
	})

	t.Run("iterator with nil slice", func(t *testing.T) {
		var items []string
		iter := NewIterator(items)

		assert.False(t, iter.HasNext())
		item, ok := iter.Next()
		assert.False(t, ok)
		assert.Equal(t, "", item)
	})

	t.Run("loop pattern usage", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		iter := NewIterator(items)

		collected := []int{}
		for item, ok := iter.Next(); ok; item, ok = iter.Next() {
			collected = append(collected, item)
		}

		assert.Equal(t, items, collected)
	})

	t.Run("mixed HasNext and Next", func(t *testing.T) {
		items := []string{"a", "b", "c"}
		iter := NewIterator(items)

		// Check HasNext, then Next
		assert.True(t, iter.HasNext())
		item1, ok1 := iter.Next()
		assert.True(t, ok1)
		assert.Equal(t, "a", item1)

		// Next without HasNext
		item2, ok2 := iter.Next()
		assert.True(t, ok2)
		assert.Equal(t, "b", item2)

		// Check HasNext, then Next
		assert.True(t, iter.HasNext())
		item3, ok3 := iter.Next()
		assert.True(t, ok3)
		assert.Equal(t, "c", item3)

		// Exhausted
		assert.False(t, iter.HasNext())
	})
}
