package simjs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSGetWorld(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	v := w.RunScript("script.js", "$world")
	assert.NotNil(t, v)
	w2, ok := v.Export().(*World)
	assert.True(t, ok)
	assert.NotNil(t, w2)
	assert.Equal(t, w, w2)
}

func TestJSGetActor(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	v := w.RunScript("script.js", "$world.getActor('counter')")
	assert.NotNil(t, v)
	w1 := u.GetWorld("demo")
	assert.NotNil(t, w1)
	a1 := w1.GetActor("counter")
	assert.NotNil(t, a1)
	a2, ok := v.Export().(*Actor)
	assert.True(t, ok)
	assert.NotNil(t, a2)
	assert.Equal(t, a1, a2)
}

func TestJSCreateActor(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	v := w.RunScript("script.js", "$world.createActor('counter')")
	assert.NotNil(t, v)
	w1 := u.GetWorld("demo")
	assert.NotNil(t, w1)
	a1 := w1.GetActor("counter")
	assert.NotNil(t, a1)
	a2, ok := v.Export().(*Actor)
	assert.True(t, ok)
	assert.NotNil(t, a2)
	assert.Equal(t, a1, a2)
}
