package simjs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetProperty(t *testing.T) {
	pub := NewMockPublisher()
	w, err := NewWorld("test", pub)
	assert.NoError(t, err)
	actor := NewActor("test", map[string]any{}, w)
	actor.Set("prop", 1)
	assert.Equal(t, 1, actor.Get("prop"))
}

func TestEmitChange(t *testing.T) {
	pub := NewMockPublisher()
	w, err := NewWorld("test", pub)
	assert.NoError(t, err)
	actor := NewActor("test", map[string]any{}, w)
	var changed bool
	actor.OnEvent(func(e SimuMessage) {
		if e.Event == EActorPropertyChanged && e.Member == "prop" {
			changed = true
		}
	})
	actor.Set("prop", 1)
	assert.True(t, changed)
}

func TestSetState(t *testing.T) {
	pub := NewMockPublisher()
	w, err := NewWorld("test", pub)
	assert.NoError(t, err)
	actor := NewActor("test", map[string]any{}, w)
	actor.SetState(map[string]any{"prop": 2})
	actor.SetState(map[string]any{"prop1": 2})
	assert.Equal(t, 2, actor.Get("prop"))
	assert.Equal(t, 2, actor.Get("prop1"))
	actor.SetState(map[string]any{"prop": 1})
	assert.Equal(t, 1, actor.Get("prop"))
	assert.Equal(t, 2, actor.Get("prop1"))
}

func TestMethodCall(t *testing.T) {
	pub := NewMockPublisher()
	w, err := NewWorld("test", pub)
	assert.NoError(t, err)
	actor := NewActor("test", map[string]any{}, w)
	var called bool
	actor.OnMethod("test", func(args ...any) any {
		called = true
		return nil
	})
	actor.Call("test")
	assert.True(t, called)
}

func TestMethodCallFailure(t *testing.T) {
	pub := NewMockPublisher()
	w, err := NewWorld("test", pub)
	assert.NoError(t, err)
	actor := NewActor("test", map[string]any{}, w)
	var called bool
	actor.Call("test1")
	assert.False(t, called)
}
