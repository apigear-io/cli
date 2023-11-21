package ostore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreate tests the Create method
func TestCreate(t *testing.T) {
	store := NewMemoryStore()
	store.Set("foo", map[string]any{"bar": "baz"})
}

// TestUpdate tests the Update method
func TestUpdate(t *testing.T) {
	store := NewMemoryStore()
	store.Set("foo", map[string]any{"bar": "baz"})
	props := store.Get("foo")
	assert.Equal(t, "baz", props["bar"])
	store.Update("foo", map[string]any{"bar": "qux"})
	props = store.Get("foo")
	assert.Equal(t, "qux", props["bar"])
}

// TestDelete tests the Delete method
func TestDelete(t *testing.T) {
	store := NewMemoryStore()
	store.Set("foo", map[string]any{"bar": "baz"})
	store.Delete("foo")
	store.Get("foo")
	assert.False(t, store.Has("foo"))
}

// TestGet tests the Get method
func TestGet(t *testing.T) {
	store := NewMemoryStore()
	store.Set("foo", map[string]any{"bar": "baz"})
	props := store.Get("foo")
	assert.Equal(t, "baz", props["bar"])
}

// TestHas tests the Has method
func TestHas(t *testing.T) {
	store := NewMemoryStore()
	assert.False(t, store.Has("foo"))
	store.Set("foo", map[string]any{"bar": "baz"})
	assert.True(t, store.Has("foo"))
}

// TestWatch tests the Watch method
func TestWatch(t *testing.T) {
	store := NewMemoryStore()
	var event StoreEvent
	store.OnEvent(func(evt StoreEvent) {
		event = evt
	})
	store.Set("foo", map[string]any{"bar": "baz"})

	assert.Equal(t, EventTypeSet, event.Type)
	assert.Equal(t, "foo", event.Id)
	assert.Equal(t, "baz", event.KWArgs["bar"])
}

// TestWatchStop tests the WatchStop method
func TestWatchStop(t *testing.T) {
	store := NewMemoryStore()
	var event StoreEvent
	unWatch := store.OnEvent(func(evt StoreEvent) {
		event = evt
	})
	store.Set("foo", map[string]any{"bar": "baz"})
	assert.Equal(t, EventTypeSet, event.Type)
	assert.Equal(t, "foo", event.Id)
	assert.Equal(t, "baz", event.KWArgs["bar"])
	event = StoreEvent{}
	unWatch()
	store.Update("foo", map[string]any{"bar": "qux"})
	assert.Equal(t, EventTypeNone, event.Type)
	assert.Equal(t, "", event.Id)
	assert.Nil(t, event.KWArgs)
}

// TestWatchUpdate tests the WatchUpdate method
func TestWatchUpdate(t *testing.T) {
	store := NewMemoryStore()
	var event StoreEvent
	store.OnEvent(func(evt StoreEvent) {
		event = evt
	})
	store.Set("foo", map[string]any{"bar": "baz"})
	store.Update("foo", map[string]any{"bar": "qux"})
	assert.Equal(t, EventTypeUpdate, event.Type)
	assert.Equal(t, "foo", event.Id)
	assert.Equal(t, "qux", event.KWArgs["bar"])
}

// TestWatchDelete tests the WatchDelete method
func TestWatchDelete(t *testing.T) {
	store := NewMemoryStore()
	var event StoreEvent
	store.OnEvent(func(evt StoreEvent) {
		event = evt
	})
	store.Set("foo", map[string]any{"bar": "baz"})
	store.Delete("foo")
	assert.Equal(t, EventTypeDelete, event.Type)
	assert.Equal(t, "foo", event.Id)
	assert.Nil(t, event.KWArgs)
}

// TestWatchCreate tests the WatchCreate method
func TestWatchCreate(t *testing.T) {
	store := NewMemoryStore()
	var event StoreEvent
	store.OnEvent(func(evt StoreEvent) {
		event = evt
	})
	store.Set("foo", map[string]any{"bar": "baz"})
	assert.Equal(t, EventTypeSet, event.Type)
	assert.Equal(t, "foo", event.Id)
	assert.Equal(t, "baz", event.KWArgs["bar"])
}
