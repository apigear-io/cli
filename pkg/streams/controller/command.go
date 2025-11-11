package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/google/uuid"
)

// startCommand captures a validated controller start request with parsed fields.
type startCommand struct {
	Subject       string
	DeviceID      string
	SessionID     string
	Retention     time.Duration
	SessionBucket string
	Note          string
	PreRoll       time.Duration
	Verbose       bool
}

func (cmd RpcRequest) normalizeStart() (startCommand, error) {
	var out startCommand

	subject := strings.TrimSpace(cmd.Subject)
	if subject == "" {
		return out, fmt.Errorf("subject cannot be empty")
	}
	out.Subject = subject

	deviceID := strings.TrimSpace(cmd.DeviceID)
	if deviceID == "" {
		return out, fmt.Errorf("device-id cannot be empty")
	}
	out.DeviceID = deviceID

	sessionID := strings.TrimSpace(cmd.SessionID)
	if sessionID == "" {
		sessionID = uuid.NewString()
	}
	out.SessionID = sessionID

	retention, err := parseRetention(cmd.Retention)
	if err != nil {
		return out, err
	}
	out.Retention = retention

	sessionBucket := strings.TrimSpace(cmd.SessionBucket)
	if sessionBucket == "" {
		sessionBucket = config.SessionBucket
	}
	out.SessionBucket = sessionBucket

	out.Note = strings.TrimSpace(cmd.Note)

	preRoll := strings.TrimSpace(cmd.PreRoll)
	if preRoll != "" {
		dur, err := time.ParseDuration(preRoll)
		if err != nil {
			return out, fmt.Errorf("invalid pre-roll: %v", err)
		}
		out.PreRoll = dur
	}

	out.Verbose = cmd.Verbose

	return out, nil
}
