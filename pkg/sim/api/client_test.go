package api_test

import (
	"testing"

	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/api"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	manager := sim.GetManager()
	defer manager.RemoveAll()
	api.NewService(nats.DefaultURL, manager)
	_, err := manager.CreateService(nats.DefaultURL)
	assert.NoError(t, err)

	client, err := manager.CreateClient(nats.DefaultURL)
	if err != nil {
		t.Skip("NATS server not available")
	}
	defer client.Close()

	t.Run("run script", func(t *testing.T) {
		value, err := client.RunScript("test", model.Script{
			Name:   "script.js",
			Source: "42",
		})
		assert.NoError(t, err)
		assert.Equal(t, float64(42), value)
	})

	t.Run("actor operations", func(t *testing.T) {
		// Create an actor with state
		source := `
			const actor = $world.createActor("test-actor", { count: 42 });
			actor.$getProperty("count");
		`
		value, err := client.RunScript("test", model.Script{
			Name:   "script.js",
			Source: source,
		})
		assert.NoError(t, err)
		assert.Equal(t, float64(42), value)

		// Get actor state
		state, err := client.GetActorState("test", "test-actor")
		assert.NoError(t, err)
		assert.Equal(t, float64(42), state["count"])

		err = client.SetActorValue("test", "test-actor", "count", 84)
		assert.NoError(t, err)

		value, err = client.GetActorValue("test", "test-actor", "count")
		assert.NoError(t, err)
		assert.Equal(t, float64(84), value)
	})
}
