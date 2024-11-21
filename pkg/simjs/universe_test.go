package simjs

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/objectlink-core-go/log"
	"github.com/stretchr/testify/assert"
)

type MockPublisher struct {
	pub  map[string][]*SimuMessage
	hook helper.Hook[*SimuMessage]
}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{
		pub:  make(map[string][]*SimuMessage),
		hook: helper.Hook[*SimuMessage]{},
	}
}

func (p *MockPublisher) Publish(subject string, msg *SimuMessage) error {
	log.Info().Msgf("mock publish %s %v", subject, msg)
	p.pub[subject] = append(p.pub[subject], msg)
	p.hook.FireHook(msg)
	return nil
}

func (p *MockPublisher) OnMessage(fn func(msg *SimuMessage)) {
	p.hook.AddHook(fn)
}

func TestUniverseCreate(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	assert.NotNil(t, u)
}

func TestUniverseAddWorld(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w, err := u.AddWorld("test")
	assert.NoError(t, err)
	assert.NotNil(t, w)
	assert.Equal(t, 1, len(u.Worlds()))
}

func TestUniverseWorlds(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w, err := u.AddWorld("test")
	assert.NoError(t, err)
	assert.NotNil(t, w)
	assert.Equal(t, 1, len(u.Worlds()))
	assert.Equal(t, w, u.Worlds()["test"])
}

func TestUniverseGetWorld(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w, err := u.AddWorld("test")
	assert.NoError(t, err)
	assert.NotNil(t, w)
	w2 := u.GetWorld("test")
	assert.Equal(t, w, w2)
}

func TestUniverseWorldOnDemand(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("test")
	assert.NotNil(t, w)
}

func TestUniverseStartDynamicWorld(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	var started bool
	pub.OnMessage(func(msg *SimuMessage) {
		if msg.Event == EWorldChanged {
			if msg.Value == "started" {
				started = true
			}
		}
	})
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	w.StartScript()
	assert.True(t, started)
}

func TestUniverseStartStopWorld(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	var running bool
	var stopped bool
	pub.OnMessage(func(msg *SimuMessage) {
		if msg.Event == EWorldChanged && msg.Value == "running" {
			running = true
		}
		if msg.Event == EWorldChanged && msg.Value == "stopped" {
			stopped = true
		}
	})
	w := u.GetWorld("demo")
	w.RunScript("demo.js", "console.log('hello')")
	assert.True(t, running)
	w.StopScript()
	assert.True(t, stopped)
}

func TestUniverseRemoveWorld(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w, err := u.AddWorld("test")
	assert.NoError(t, err)
	assert.NotNil(t, w)
	assert.Equal(t, 1, len(u.Worlds()))
	u.WorldRemove("test")
	assert.Equal(t, 0, len(u.Worlds()))
}

func TestUniverseActorMethodCall(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("test")
	assert.NotNil(t, w)
	w.RunScript("demo.js", "console.log('hello')")
	actor := w.CreateActor("test", map[string]any{})
	assert.NotNil(t, actor)
	var called bool
	actor.OnMethod("test", func(args ...any) any {
		called = true
		return nil
	})
	actor.Call("test")
	assert.True(t, called)
}

func TestUniverseActorMethodCallFailure(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("test")
	assert.NotNil(t, w)
	w.RunScript("demo.js", "console.log('hello')")
	actor := w.CreateActor("test", map[string]any{})
	assert.NotNil(t, actor)
	var called bool
	actor.Call("test1")
	assert.False(t, called)
}
