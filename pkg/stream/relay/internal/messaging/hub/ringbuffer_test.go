package hub

import (
	"sync"
	"testing"
)

func TestRingBuffer_NewRingBuffer(t *testing.T) {
	rb := NewRingBuffer[int](10)
	if rb.Cap() != 10 {
		t.Errorf("expected capacity 10, got %d", rb.Cap())
	}
	if rb.Len() != 0 {
		t.Errorf("expected length 0, got %d", rb.Len())
	}
}

func TestRingBuffer_NewRingBuffer_ZeroCapacity(t *testing.T) {
	rb := NewRingBuffer[int](0)
	if rb.Cap() != 1 {
		t.Errorf("expected capacity 1 for zero input, got %d", rb.Cap())
	}
}

func TestRingBuffer_Push(t *testing.T) {
	rb := NewRingBuffer[int](3)

	rb.Push(1)
	if rb.Len() != 1 {
		t.Errorf("expected length 1, got %d", rb.Len())
	}

	rb.Push(2)
	rb.Push(3)
	if rb.Len() != 3 {
		t.Errorf("expected length 3, got %d", rb.Len())
	}
}

func TestRingBuffer_Entries_PartialFill(t *testing.T) {
	rb := NewRingBuffer[int](5)
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)

	entries := rb.Entries()
	expected := []int{1, 2, 3}

	if len(entries) != len(expected) {
		t.Fatalf("expected %d entries, got %d", len(expected), len(entries))
	}

	for i, v := range expected {
		if entries[i] != v {
			t.Errorf("entry[%d]: expected %d, got %d", i, v, entries[i])
		}
	}
}

func TestRingBuffer_Entries_FullBuffer(t *testing.T) {
	rb := NewRingBuffer[int](3)
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)

	entries := rb.Entries()
	expected := []int{1, 2, 3}

	if len(entries) != len(expected) {
		t.Fatalf("expected %d entries, got %d", len(expected), len(entries))
	}

	for i, v := range expected {
		if entries[i] != v {
			t.Errorf("entry[%d]: expected %d, got %d", i, v, entries[i])
		}
	}
}

func TestRingBuffer_Entries_Overflow(t *testing.T) {
	rb := NewRingBuffer[int](3)
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)
	rb.Push(4) // overwrites 1
	rb.Push(5) // overwrites 2

	entries := rb.Entries()
	expected := []int{3, 4, 5}

	if len(entries) != len(expected) {
		t.Fatalf("expected %d entries, got %d", len(expected), len(entries))
	}

	for i, v := range expected {
		if entries[i] != v {
			t.Errorf("entry[%d]: expected %d, got %d", i, v, entries[i])
		}
	}
}

func TestRingBuffer_Entries_Empty(t *testing.T) {
	rb := NewRingBuffer[int](3)
	entries := rb.Entries()

	if entries != nil {
		t.Errorf("expected nil for empty buffer, got %v", entries)
	}
}

func TestRingBuffer_Clear(t *testing.T) {
	rb := NewRingBuffer[int](3)
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)

	rb.Clear()

	if rb.Len() != 0 {
		t.Errorf("expected length 0 after clear, got %d", rb.Len())
	}

	entries := rb.Entries()
	if entries != nil {
		t.Errorf("expected nil entries after clear, got %v", entries)
	}
}

func TestRingBuffer_ConcurrentAccess(t *testing.T) {
	rb := NewRingBuffer[int](100)
	var wg sync.WaitGroup

	// Concurrent writers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(base int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				rb.Push(base*100 + j)
			}
		}(i)
	}

	// Concurrent readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_ = rb.Entries()
				_ = rb.Len()
			}
		}()
	}

	wg.Wait()

	// Buffer should be at capacity
	if rb.Len() != 100 {
		t.Errorf("expected length 100, got %d", rb.Len())
	}
}

func TestRingBuffer_WithStrings(t *testing.T) {
	rb := NewRingBuffer[string](2)
	rb.Push("hello")
	rb.Push("world")
	rb.Push("foo") // overwrites "hello"

	entries := rb.Entries()
	expected := []string{"world", "foo"}

	if len(entries) != len(expected) {
		t.Fatalf("expected %d entries, got %d", len(expected), len(entries))
	}

	for i, v := range expected {
		if entries[i] != v {
			t.Errorf("entry[%d]: expected %q, got %q", i, v, entries[i])
		}
	}
}

func TestRingBuffer_WithStructs(t *testing.T) {
	type Item struct {
		ID   int
		Name string
	}

	rb := NewRingBuffer[Item](2)
	rb.Push(Item{1, "one"})
	rb.Push(Item{2, "two"})
	rb.Push(Item{3, "three"}) // overwrites {1, "one"}

	entries := rb.Entries()

	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	if entries[0].ID != 2 || entries[0].Name != "two" {
		t.Errorf("entry[0]: expected {2, two}, got %+v", entries[0])
	}

	if entries[1].ID != 3 || entries[1].Name != "three" {
		t.Errorf("entry[1]: expected {3, three}, got %+v", entries[1])
	}
}
