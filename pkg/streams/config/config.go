package config

import (
	"fmt"
	"strings"
	"time"
)

const (
	// Default JetStream bucket names.
	SessionBucket = "streams_session"
	DeviceBucket  = "streams_devices"
	StateBucket   = "streams_record_state"

	// Default subjects and prefixes.
	RecordRpcSubject           = "streams.record.rpc"
	SessionSubjectPrefix       = "streams.session"
	BufferSubjectPrefix        = "streams.buffer"
	MonitorSubject             = "monitor"
	RecordControllerQueueGroup = "streams-record-controller"
	PlaybackSubject            = "streams.playback"

	// Header keys used across publishing, recording, and buffering flows.
	HeaderDevice     = "X-Streams-Device"
	HeaderSession    = "X-Streams-Session"
	HeaderFile       = "X-Streams-File"
	HeaderRecordedAt = "X-Streams-Recorded-At"
	HeaderReplayedAt = "X-Streams-Replayed-At"
	HeaderBufferedAt = "X-Streams-Buffered-At"
	HeaderDeadline   = "X-Streams-Deadline"
	HeaderPreRoll    = "X-Streams-PreRoll"

	BufferRefresh = 15 * time.Second
)

// SessionSubject returns the fully qualified JetStream subject used to persist
// recorded session messages for the given session identifier.
func SessionSubject(sessionID string) string {
	if sessionID == "" {
		return SessionSubjectPrefix
	}
	return fmt.Sprintf("%s.%s", SessionSubjectPrefix, sessionID)
}

// DeviceSubject returns a device-scoped subject by concatenating the base
// subject prefix and the provided device identifier.
func DeviceSubject(base, deviceID string) string {
	if base == "" || deviceID == "" {
		return base
	}
	return fmt.Sprintf("%s.%s", base, deviceID)
}

func SanitizeId(id string) string {
	cleaned := strings.ToUpper(id)
	cleaned = strings.ReplaceAll(cleaned, "-", "_")
	cleaned = strings.ReplaceAll(cleaned, ".", "_")
	return cleaned
}

func BufferSubjectName(deviceID string) string {
	return fmt.Sprintf("%s.%s", BufferSubjectPrefix, SanitizeId(deviceID))
}

func BufferStreamName(deviceID string) string {
	return "STREAMS_BUFFER_" + SanitizeId(deviceID)
}

func SubjectJoin(s ...string) string {
	return strings.Join(s, ".")
}

func ExportConsumerName(sessionID string) string {
	return fmt.Sprintf("EXP_%s", SanitizeId(sessionID))
}

func PlaybackConsumerName(sessionID string) string {
	return fmt.Sprintf("PB_%s_%d", SanitizeId(sessionID), time.Now().UnixNano())
}

func BufferReplayConsumerName(deviceID string) string {
	return fmt.Sprintf("BUFREP_%s_%d", SanitizeId(deviceID), time.Now().UnixNano())
}
