package protocol

import (
	"encoding/json"
	"testing"
)

func TestParseMessage_Link(t *testing.T) {
	raw := json.RawMessage(`[10, "demo.Counter"]`)
	parsed := ParseMessage(raw)

	if parsed.MsgType != MsgLink {
		t.Errorf("expected MsgType %d, got %d", MsgLink, parsed.MsgType)
	}
	if parsed.MsgTypeName != "LINK" {
		t.Errorf("expected MsgTypeName LINK, got %s", parsed.MsgTypeName)
	}
	if parsed.Symbol != "demo.Counter" {
		t.Errorf("expected Symbol demo.Counter, got %s", parsed.Symbol)
	}
}

func TestParseMessage_Init(t *testing.T) {
	raw := json.RawMessage(`[11, "demo.Counter", {"count": 0}]`)
	parsed := ParseMessage(raw)

	if parsed.MsgType != MsgInit {
		t.Errorf("expected MsgType %d, got %d", MsgInit, parsed.MsgType)
	}
	if parsed.MsgTypeName != "INIT" {
		t.Errorf("expected MsgTypeName INIT, got %s", parsed.MsgTypeName)
	}
	if parsed.Symbol != "demo.Counter" {
		t.Errorf("expected Symbol demo.Counter, got %s", parsed.Symbol)
	}
	if parsed.Args == nil {
		t.Error("expected Args to be set")
	}
}

func TestParseMessage_Invoke(t *testing.T) {
	raw := json.RawMessage(`[30, 1, "demo.Counter/increment", {"step": 1}]`)
	parsed := ParseMessage(raw)

	if parsed.MsgType != MsgInvoke {
		t.Errorf("expected MsgType %d, got %d", MsgInvoke, parsed.MsgType)
	}
	if parsed.MsgTypeName != "INVOKE" {
		t.Errorf("expected MsgTypeName INVOKE, got %s", parsed.MsgTypeName)
	}
	if parsed.Symbol != "demo.Counter/increment" {
		t.Errorf("expected Symbol demo.Counter/increment, got %s", parsed.Symbol)
	}
	if parsed.RequestID == nil || *parsed.RequestID != 1 {
		t.Errorf("expected RequestID 1, got %v", parsed.RequestID)
	}
	if parsed.Args == nil {
		t.Error("expected Args to be set")
	}
}

func TestParseMessage_Signal(t *testing.T) {
	raw := json.RawMessage(`[40, "demo.Counter/changed", {"count": 5}]`)
	parsed := ParseMessage(raw)

	if parsed.MsgType != MsgSignal {
		t.Errorf("expected MsgType %d, got %d", MsgSignal, parsed.MsgType)
	}
	if parsed.MsgTypeName != "SIGNAL" {
		t.Errorf("expected MsgTypeName SIGNAL, got %s", parsed.MsgTypeName)
	}
	if parsed.Symbol != "demo.Counter/changed" {
		t.Errorf("expected Symbol demo.Counter/changed, got %s", parsed.Symbol)
	}
	if parsed.Args == nil {
		t.Error("expected Args to be set")
	}
}

func TestParseMessage_Error(t *testing.T) {
	raw := json.RawMessage(`[90, 30, 1, "method not found"]`)
	parsed := ParseMessage(raw)

	if parsed.MsgType != MsgError {
		t.Errorf("expected MsgType %d, got %d", MsgError, parsed.MsgType)
	}
	if parsed.MsgTypeName != "ERROR" {
		t.Errorf("expected MsgTypeName ERROR, got %s", parsed.MsgTypeName)
	}
	if parsed.RequestID == nil || *parsed.RequestID != 1 {
		t.Errorf("expected RequestID 1, got %v", parsed.RequestID)
	}
	if parsed.Args == nil {
		t.Error("expected Args to be set")
	}
}

func TestParseMessage_Invalid(t *testing.T) {
	raw := json.RawMessage(`invalid json`)
	parsed := ParseMessage(raw)

	if parsed.MsgTypeName != "UNKNOWN" {
		t.Errorf("expected MsgTypeName UNKNOWN, got %s", parsed.MsgTypeName)
	}
}

func TestParseMessage_UnknownType(t *testing.T) {
	raw := json.RawMessage(`[999, "unknown"]`)
	parsed := ParseMessage(raw)

	if parsed.MsgType != 999 {
		t.Errorf("expected MsgType 999, got %d", parsed.MsgType)
	}
	if parsed.MsgTypeName != "UNKNOWN" {
		t.Errorf("expected MsgTypeName UNKNOWN, got %s", parsed.MsgTypeName)
	}
}
