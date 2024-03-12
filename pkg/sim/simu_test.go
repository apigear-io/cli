package sim

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/spec"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// readScenario reads a scenario from file.
func readScenario(source string) (*spec.ScenarioDoc, error) {
	bytes, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}
	doc := &spec.ScenarioDoc{}
	err = yaml.Unmarshal(bytes, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func TestNew(t *testing.T) {
	t.Parallel()
	s := NewSimulation()
	require.NotNil(t, s)
}

// Test load and unload scenario
func TestLoadUnloadScenario(t *testing.T) {
	t.Parallel()
	s := NewSimulation()
	require.NotNil(t, s)
	doc, err := readScenario("testdata/demo.scenario.yaml")
	require.NoError(t, err)
	err = s.LoadScenario("scenario1", doc)
	require.NoError(t, err)
	err = s.UnloadScenario("scenario1")
	require.NoError(t, err)
	scenarios := s.ActiveScenarios()
	require.Equal(t, 0, len(scenarios))
}

func TestSameLoadScenario(t *testing.T) {
	t.Parallel()
	s := NewSimulation()
	require.NotNil(t, s)
	doc, err := readScenario("testdata/demo.scenario.yaml")
	require.NoError(t, err)
	err = s.LoadScenario("scenario1", doc)
	require.NoError(t, err)
	err = s.LoadScenario("scenario1", doc)
	require.NoError(t, err)
	err = s.UnloadScenario("scenario1")
	require.NoError(t, err)
	scenarios := s.ActiveScenarios()
	require.Equal(t, 0, len(scenarios))
}

// Start/Stop events are only send when the scenario has sequences
func TestSeqStartEvent(t *testing.T) {
	t.Parallel()
	seqStarted := false
	seqStopped := false
	s := NewSimulation()
	s.OnEvent(func(evt *core.SimuEvent) {
		slog.Debug("simu event", "event", evt)
		switch evt.Type {
		case core.EventSeqStart:
			seqStarted = true
		case core.EventSeqStop:
			seqStopped = true
		}
	})
	require.NotNil(t, s)
	doc, err := readScenario("testdata/demo.scenario.yaml")
	require.NoError(t, err)
	err = s.LoadScenario("scenario1", doc)
	require.NoError(t, err)
	scenarios := s.ActiveScenarios()
	require.Equal(t, 1, len(scenarios))
	ctx := context.Background()
	err = s.PlayAllSequences(ctx)
	require.NoError(t, err)
	require.True(t, seqStarted)
	err = s.UnloadScenario("scenario1")
	require.NoError(t, err)
	require.True(t, seqStopped)
}
