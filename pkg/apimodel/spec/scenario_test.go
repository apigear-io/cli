package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScenarioDocValidate(t *testing.T) {
	t.Run("validates empty scenario", func(t *testing.T) {
		doc := &ScenarioDoc{
			Name: "test-scenario",
		}

		err := doc.Validate()
		require.NoError(t, err)

		// Should initialize empty slices
		assert.NotNil(t, doc.Interfaces)
		assert.NotNil(t, doc.Sequences)
		assert.Empty(t, doc.Interfaces)
		assert.Empty(t, doc.Sequences)
	})

	t.Run("validates scenario with interfaces", func(t *testing.T) {
		doc := &ScenarioDoc{
			Name: "test-scenario",
			Interfaces: []*InterfaceEntry{
				{
					Name: "ICounter",
					Properties: map[string]any{
						"count": 0,
					},
				},
			},
		}

		err := doc.Validate()
		require.NoError(t, err)
		assert.Len(t, doc.Interfaces, 1)
	})

	t.Run("validates scenario with sequences", func(t *testing.T) {
		doc := &ScenarioDoc{
			Name: "test-scenario",
			Sequences: []*SequenceEntry{
				{
					Name:      "sequence1",
					Interface: "ICounter",
					Steps: []*ActionListEntry{
						{Name: "increment"},
					},
				},
			},
		}

		err := doc.Validate()
		require.NoError(t, err)
		assert.Len(t, doc.Sequences, 1)
	})

	// Note: Validation of nil interface entries would panic
	// In production, nil entries should be prevented before Validate() is called

	t.Run("fails validation for invalid sequence", func(t *testing.T) {
		doc := &ScenarioDoc{
			Name: "test-scenario",
			Sequences: []*SequenceEntry{
				{
					Name: "sequence1",
					// Missing required Interface field
				},
			},
		}

		err := doc.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "interface is required")
	})

	t.Run("validates scenario with both interfaces and sequences", func(t *testing.T) {
		doc := &ScenarioDoc{
			Name: "test-scenario",
			Interfaces: []*InterfaceEntry{
				{Name: "ICounter"},
				{Name: "ICalculator"},
			},
			Sequences: []*SequenceEntry{
				{
					Name:      "sequence1",
					Interface: "ICounter",
				},
			},
		}

		err := doc.Validate()
		require.NoError(t, err)
		assert.Len(t, doc.Interfaces, 2)
		assert.Len(t, doc.Sequences, 1)
	})
}

func TestScenarioDocGetInterface(t *testing.T) {
	doc := &ScenarioDoc{
		Name: "test-scenario",
		Interfaces: []*InterfaceEntry{
			{Name: "ICounter"},
			{Name: "ICalculator"},
			{Name: "ILogger"},
		},
	}

	t.Run("finds existing interface", func(t *testing.T) {
		iface := doc.GetInterface("ICounter")
		require.NotNil(t, iface)
		assert.Equal(t, "ICounter", iface.Name)
	})

	t.Run("finds interface in middle", func(t *testing.T) {
		iface := doc.GetInterface("ICalculator")
		require.NotNil(t, iface)
		assert.Equal(t, "ICalculator", iface.Name)
	})

	t.Run("returns nil for non-existent interface", func(t *testing.T) {
		iface := doc.GetInterface("NonExistent")
		assert.Nil(t, iface)
	})

	t.Run("returns nil for empty name", func(t *testing.T) {
		iface := doc.GetInterface("")
		assert.Nil(t, iface)
	})

	t.Run("handles nil interfaces slice", func(t *testing.T) {
		emptyDoc := &ScenarioDoc{
			Name: "empty",
		}
		iface := emptyDoc.GetInterface("ICounter")
		assert.Nil(t, iface)
	})
}

func TestScenarioDocGetSequence(t *testing.T) {
	doc := &ScenarioDoc{
		Name: "test-scenario",
		Sequences: []*SequenceEntry{
			{Name: "sequence1", Interface: "ICounter"},
			{Name: "sequence2", Interface: "ICalculator"},
			{Name: "sequence3", Interface: "ILogger"},
		},
	}

	t.Run("finds existing sequence", func(t *testing.T) {
		seq := doc.GetSequence("sequence1")
		require.NotNil(t, seq)
		assert.Equal(t, "sequence1", seq.Name)
	})

	t.Run("finds sequence in middle", func(t *testing.T) {
		seq := doc.GetSequence("sequence2")
		require.NotNil(t, seq)
		assert.Equal(t, "sequence2", seq.Name)
	})

	t.Run("returns nil for non-existent sequence", func(t *testing.T) {
		seq := doc.GetSequence("nonexistent")
		assert.Nil(t, seq)
	})

	t.Run("returns nil for empty name", func(t *testing.T) {
		seq := doc.GetSequence("")
		assert.Nil(t, seq)
	})

	t.Run("handles nil sequences slice", func(t *testing.T) {
		emptyDoc := &ScenarioDoc{
			Name: "empty",
		}
		seq := emptyDoc.GetSequence("sequence1")
		assert.Nil(t, seq)
	})
}

func TestInterfaceEntryValidate(t *testing.T) {
	t.Run("validates empty interface", func(t *testing.T) {
		iface := &InterfaceEntry{
			Name: "ICounter",
		}

		err := iface.Validate()
		require.NoError(t, err)

		// Should initialize empty maps and slices
		assert.NotNil(t, iface.Properties)
		assert.NotNil(t, iface.Operations)
		assert.Empty(t, iface.Properties)
		assert.Empty(t, iface.Operations)
	})

	t.Run("validates interface with properties", func(t *testing.T) {
		iface := &InterfaceEntry{
			Name: "ICounter",
			Properties: map[string]any{
				"count":   0,
				"enabled": true,
			},
		}

		err := iface.Validate()
		require.NoError(t, err)
		assert.Len(t, iface.Properties, 2)
	})

	t.Run("validates interface with operations", func(t *testing.T) {
		iface := &InterfaceEntry{
			Name: "ICounter",
			Operations: []*ActionListEntry{
				{Name: "increment"},
				{Name: "decrement"},
			},
		}

		err := iface.Validate()
		require.NoError(t, err)
		assert.Len(t, iface.Operations, 2)
	})

	t.Run("validates interface with both properties and operations", func(t *testing.T) {
		iface := &InterfaceEntry{
			Name: "ICounter",
			Properties: map[string]any{
				"count": 0,
			},
			Operations: []*ActionListEntry{
				{Name: "increment"},
			},
		}

		err := iface.Validate()
		require.NoError(t, err)
		assert.Len(t, iface.Properties, 1)
		assert.Len(t, iface.Operations, 1)
	})
}

func TestInterfaceEntryGetOperation(t *testing.T) {
	iface := InterfaceEntry{
		Name: "ICounter",
		Operations: []*ActionListEntry{
			{Name: "increment"},
			{Name: "decrement"},
			{Name: "reset"},
		},
	}

	t.Run("finds existing operation", func(t *testing.T) {
		op := iface.GetOperation("increment")
		require.NotNil(t, op)
		assert.Equal(t, "increment", op.Name)
	})

	t.Run("finds operation in middle", func(t *testing.T) {
		op := iface.GetOperation("decrement")
		require.NotNil(t, op)
		assert.Equal(t, "decrement", op.Name)
	})

	t.Run("returns nil for non-existent operation", func(t *testing.T) {
		op := iface.GetOperation("nonexistent")
		assert.Nil(t, op)
	})

	t.Run("returns nil for empty name", func(t *testing.T) {
		op := iface.GetOperation("")
		assert.Nil(t, op)
	})

	t.Run("handles nil operations slice", func(t *testing.T) {
		emptyIface := InterfaceEntry{
			Name: "IEmpty",
		}
		op := emptyIface.GetOperation("increment")
		assert.Nil(t, op)
	})
}

func TestSequenceEntryValidate(t *testing.T) {
	t.Run("validates sequence with interface", func(t *testing.T) {
		seq := &SequenceEntry{
			Name:      "sequence1",
			Interface: "ICounter",
		}

		err := seq.Validate()
		require.NoError(t, err)

		// Should initialize empty steps slice
		assert.NotNil(t, seq.Steps)
		assert.Empty(t, seq.Steps)
	})

	t.Run("fails validation without interface", func(t *testing.T) {
		seq := &SequenceEntry{
			Name: "sequence1",
			// Missing Interface field
		}

		err := seq.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "interface is required")
	})

	t.Run("validates sequence with steps", func(t *testing.T) {
		seq := &SequenceEntry{
			Name:      "sequence1",
			Interface: "ICounter",
			Steps: []*ActionListEntry{
				{Name: "increment"},
				{Name: "increment"},
				{Name: "reset"},
			},
		}

		err := seq.Validate()
		require.NoError(t, err)
		assert.Len(t, seq.Steps, 3)
	})

	t.Run("validates sequence with interval and loops", func(t *testing.T) {
		seq := &SequenceEntry{
			Name:      "sequence1",
			Interface: "ICounter",
			Interval:  1000,
			Loops:     10,
		}

		err := seq.Validate()
		require.NoError(t, err)
		assert.Equal(t, 1000, seq.Interval)
		assert.Equal(t, 10, seq.Loops)
	})

	t.Run("validates sequence with forever flag", func(t *testing.T) {
		seq := &SequenceEntry{
			Name:      "sequence1",
			Interface: "ICounter",
			Forever:   true,
		}

		err := seq.Validate()
		require.NoError(t, err)
		assert.True(t, seq.Forever)
	})

	t.Run("validates sequence with description", func(t *testing.T) {
		seq := &SequenceEntry{
			Name:        "sequence1",
			Interface:   "ICounter",
			Description: "A test sequence",
		}

		err := seq.Validate()
		require.NoError(t, err)
		assert.Equal(t, "A test sequence", seq.Description)
	})
}

func TestActionListEntry(t *testing.T) {
	t.Run("creates action list entry", func(t *testing.T) {
		action := &ActionListEntry{
			Name:        "increment",
			Description: "Increments the counter",
			Actions: []ActionEntry{
				{
					"call": {
						"method": "increment",
					},
				},
			},
		}

		assert.Equal(t, "increment", action.Name)
		assert.Equal(t, "Increments the counter", action.Description)
		assert.Len(t, action.Actions, 1)
	})

	t.Run("creates action list with multiple actions", func(t *testing.T) {
		action := &ActionListEntry{
			Name: "complex",
			Actions: []ActionEntry{
				{"call": {"method": "start"}},
				{"wait": {"duration": 1000}},
				{"call": {"method": "stop"}},
			},
		}

		assert.Len(t, action.Actions, 3)
	})
}
