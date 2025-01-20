package js

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{"valid world", "test-world", false},
		{"empty id", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorld(tt.id, goja.New())
			if tt.wantErr {
				assert.Nil(t, w)
			} else {
				assert.NotNil(t, w)
				assert.Equal(t, tt.id, w.Id_())
				assert.Empty(t, w.ListActors())
				assert.Equal(t, 0, w.ActorCount())
			}
		})
	}
}

func TestWorldActorOperations(t *testing.T) {
	w := NewDemoWorld()
	// Test actor creation
	state := w.vm.NewObject()
	actor1, err := w.CreateActor("actor1", state)
	assert.NoError(t, err)
	assert.NotNil(t, actor1)

	// Test duplicate actor creation
	_, err = w.CreateActor("actor1", state)
	assert.Error(t, err)

	// Test getting actor
	assert.NotNil(t, w.GetActor("actor1"))
	assert.Nil(t, w.GetActor("nonexistent"))

	// Test listing actors
	actors := w.ListActors()
	assert.Len(t, actors, 1)
	assert.Contains(t, actors, "actor1")

	// Test actor count
	assert.Equal(t, 1, w.ActorCount())

	// Test deleting actor
	w.DeleteActor("actor1")
	assert.Nil(t, w.GetActor("actor1"))
	assert.Equal(t, 0, w.ActorCount())
}
