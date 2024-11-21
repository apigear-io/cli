package simjs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func LoadTestData(t *testing.T) string {
	content, err := os.ReadFile("testdata/counter.js")
	assert.NoError(t, err)
	return string(content)
}

func LoadTestUniverse(t *testing.T) *Universe {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	content := LoadTestData(t)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	w.RunScript("test.js", string(content))
	return u
}

func TestJSStartWorld(t *testing.T) {
	u := LoadTestUniverse(t)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
}

func TestJSCallMethod(t *testing.T) {
	u := LoadTestUniverse(t)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	counter := w.GetActor("counter")
	assert.NotNil(t, counter)
	counter.Call("increment")
	count := counter.Get("count")
	assert.Equal(t, int64(1), count)
}

func TestCallGlobalFunction(t *testing.T) {
	u := LoadTestUniverse(t)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	result := w.Call("add", 1, 2)
	assert.Equal(t, int64(3), result)
}

const script01 = `
let newCount = 0
const counter = $world.createActor('counter', { count: 0 })
counter.onChange('count', (v) => newCount = v)
counter.onMethod('increment', () => counter.set('count', counter.get('count') + 1))
function main() {
	counter.call('increment')
	return newCount
}
`

func TestJSRunScript(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	w.RunScript("script01", script01)
	v := w.Call("main")
	assert.NotNil(t, v)
	assert.Equal(t, int64(1), v)
}

const script02 = `
function main(a, b) {
	console.log(a + " " + b);
	return a + " " + b;
}
`

func TestJSWorldFunction(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	w.RunScript("script01", script02)
	v := w.Call("main", "hello", "world")
	assert.NotNil(t, v)
	assert.Equal(t, "hello world", v)
}

// Test Access global variables
const script03 = `
var a = 1;
var b = 2;
var c = 3;
`

func TestJSWorldGlobalVariables(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	w.RunScript("script03", script03)
	v := w.GetValue("a")
	assert.NotNil(t, v)
	assert.Equal(t, int64(1), v)
}

const script04 = `
const counter = $world.createActor("counter", { count: 0 });
counter.onMethod("increment", () => {
	const count = counter.get("count");
	counter.set("count", count + 1);
});
`

func TestJSWorldUsingThis(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	w.RunScript("script04", script04)
	counter := w.GetActor("counter")
	assert.NotNil(t, counter)
	counter.Call("increment")
	count := counter.Get("count")
	assert.NotNil(t, count)
	assert.Equal(t, int64(1), count)
}

/// Test Signals

const script05 = `
var count = 0;

const actor = $world.createActor('counter')
actor.onSignal('increment', () => {
  count++;
});
`

func TestJSWorldSignals(t *testing.T) {
	pub := NewMockPublisher()
	u := NewUniverse(pub)
	w := u.GetWorld("demo")
	assert.NotNil(t, w)
	w.RunScript("script05", script05)
	a := w.GetActor("counter")
	assert.NotNil(t, a)
	a.EmitSignal("increment")
	assert.Equal(t, int64(1), w.GetValue("count"))
}
