package helper

import (
	"path/filepath"
	"testing"
)

type testEvent struct {
	Device string `json:"device"`
	Type   string `json:"type"`
}

func TestReadNDJSONFile(t *testing.T) {
	path := filepath.Join("..", "mon", "testdata", "events.ndjson")
	events, err := ReadNDJSONFile[testEvent](path)
	if err != nil {
		t.Fatalf("ReadNDJSONFile: %v", err)
	}
	if len(events) != 4 {
		t.Fatalf("expected 4 events, got %d", len(events))
	}
}
