//go:build nats
// +build nats

package api_test

import (
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	// Start NATS server
	ns, err := nats.Connect(nats.DefaultURL)
	assert.NoError(t, err)
	defer ns.Close()

	// Create manager
	manager := sim.NewManager()
	defer manager.RemoveAll()

	client, err := sim.NewManager().CreateClient(nats.DefaultURL)
	assert.NoError(t, err)
	defer client.Close()

	// Create service
	service, err := manager.CreateService(nats.DefaultURL)
	assert.NoError(t, err)
	defer service.Close()

	// Wait for services to be ready
	time.Sleep(100 * time.Millisecond)

	t.Run("run script", func(t *testing.T) {
		err := client.RunScript("test", model.Script{
			Name:   "script.js",
			Source: "42",
		})
		assert.NoError(t, err)
	})

	t.Run("get actor state", func(t *testing.T) {
		// First create an actor with state
		source := `
			const actor = $world.createActor("test-actor", { count: 42 });
			actor.$getState();
		`
		err := client.RunScript("test", model.Script{
			Name:   "script.js",
			Source: source,
		})

		assert.NoError(t, err)

		_, err = client.GetActorState("test", "test-actor")
		assert.NoError(t, err)
	})

	t.Run("get and set actor value", func(t *testing.T) {
		// First create an actor with state
		script := `
			const actor = $world.createActor("test-actor2", { count: 42 });
			actor.$getProperty("count");
		`
		err := client.RunScript("test", model.Script{
			Name:   "script.js",
			Source: script,
		})
		assert.NoError(t, err)

		err = client.SetActorValue("test", "test-actor2", "count", 42)
		assert.NoError(t, err)
		_, err = client.GetActorValue("test", "test-actor2", "count")
		assert.NoError(t, err)
	})
}
