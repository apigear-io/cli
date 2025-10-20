package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

const (
	DefaultCommandSubject = config.RecordRpcSubject
	DefaultStateBucket    = config.StateBucket
)

const (
	ActionStart = "start"
	ActionStop  = "stop"
)

// RpcRequest represents an RPC request sent to the controller.
type RpcRequest struct {
	Action        string `json:"action"`
	Subject       string `json:"subject,omitempty"`
	DeviceID      string `json:"device_id,omitempty"`
	SessionID     string `json:"session_id,omitempty"`
	Retention     string `json:"retention,omitempty"`
	SessionBucket string `json:"session_bucket,omitempty"`
	DeviceBucket  string `json:"device_bucket,omitempty"`
	DeviceDesc    string `json:"device_description,omitempty"`
	DeviceLoc     string `json:"device_location,omitempty"`
	DeviceOwner   string `json:"device_owner,omitempty"`
	PreRoll       string `json:"pre_roll,omitempty"`
	Verbose       bool   `json:"verbose,omitempty"`
}

// RpcResponse communicates the outcome of a controller command.
type RpcResponse struct {
	OK        bool           `json:"ok"`
	Message   string         `json:"message,omitempty"`
	SessionID string         `json:"session_id,omitempty"`
	State     *StateSnapshot `json:"state,omitempty"`
}

// StateSnapshot is persisted in the KV state bucket.
type StateSnapshot struct {
	SessionID     string    `json:"session_id"`
	DeviceID      string    `json:"device_id"`
	Subject       string    `json:"subject"`
	Status        string    `json:"status"`
	MessageCount  int       `json:"message_count"`
	LastError     string    `json:"last_error,omitempty"`
	StartedAt     time.Time `json:"started_at,omitempty"`
	LastMessageAt time.Time `json:"last_message_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Options configure the controller runtime.
type Options struct {
	ServerURL        string
	RecordRpcSubject string
	StateBucket      string
}

// NewController creates a new controller instance with the provided options.
func NewController(js jetstream.JetStream, opts Options) (*Controller, error) {
	if js == nil {
		return nil, errors.New("jetstream context is nil")
	}
	if opts.RecordRpcSubject == "" {
		opts.RecordRpcSubject = config.RecordRpcSubject
	}
	if opts.StateBucket == "" {
		opts.StateBucket = config.StateBucket
	}
	if opts.ServerURL == "" {
		return nil, errors.New("server URL is required")
	}

	ctx := context.Background()
	kv, err := natsutil.EnsureKeyValue(ctx, js, opts.StateBucket)
	if err != nil {
		return nil, fmt.Errorf("state bucket %s: %w", opts.StateBucket, err)
	}

	return &Controller{
		js:      js,
		opts:    opts,
		stateKV: kv,
		jobs:    map[string]*recordJob{},
	}, nil
}

type Controller struct {
	js      jetstream.JetStream
	opts    Options
	stateKV jetstream.KeyValue

	mu   sync.Mutex
	jobs map[string]*recordJob
	sub  *nats.Subscription
}

type recordJob struct {
	cancel context.CancelFunc
	done   chan struct{}
}

// Start begins listening for RPC commands on the configured subject.
func (c *Controller) Start() error {
	sub, err := c.js.Conn().Subscribe(c.opts.RecordRpcSubject, c.handleMsg)
	if err != nil {
		return fmt.Errorf("subscribe %s: %w", c.opts.RecordRpcSubject, err)
	}

	c.mu.Lock()
	c.sub = sub
	c.mu.Unlock()

	log.Info().Str("subject", c.opts.RecordRpcSubject).Msg("record controller started")
	return nil
}

// Close gracefully shuts down the controller by unsubscribing and stopping all jobs.
func (c *Controller) Close() {
	c.mu.Lock()
	if c.sub != nil {
		c.sub.Drain()
		c.sub = nil
	}
	c.mu.Unlock()

	c.stopAll()
}

func (c *Controller) stopAll() {
	c.mu.Lock()
	jobs := make([]*recordJob, 0, len(c.jobs))
	for sessionID, job := range c.jobs {
		jobs = append(jobs, job)
		delete(c.jobs, sessionID)
	}
	c.mu.Unlock()

	for _, job := range jobs {
		job.cancel()
		<-job.done
	}
}

func (c *Controller) handleMsg(msg *nats.Msg) {
	var req RpcRequest
	err := json.Unmarshal(msg.Data, &req)
	if err != nil {
		log.Error().Err(err).Msg("invalid command payload")
		c.respondError(msg, "invalid command payload: %v", err)
		return
	}

	switch strings.ToLower(req.Action) {
	case ActionStart:
		log.Debug().Str("session", req.SessionID).Str("device", req.DeviceID).Msg("handling start command")
		resp := c.handleStart(req)
		c.respond(msg, resp)
	case ActionStop:
		log.Debug().Str("session", req.SessionID).Msg("handling stop command")
		resp := c.handleStop(req)
		c.respond(msg, resp)
	default:
		log.Warn().Str("action", req.Action).Msg("unknown controller action")
		c.respondError(msg, "unknown action %q", req.Action)
	}
}

func (c *Controller) handleStart(req RpcRequest) RpcResponse {
	start, err := req.normalizeStart()
	if err != nil {
		log.Warn().Err(err).Str("action", req.Action).Msg("start command invalid")
		resp := RpcResponse{Message: err.Error()}
		if start.SessionID != "" {
			resp.SessionID = start.SessionID
		}
		return resp
	}

	if start.PreRoll > 0 {
		bufferWindow, err := c.lookupBufferWindow(start.DeviceBucket, start.DeviceID)
		if err != nil {
			return RpcResponse{Message: err.Error(), SessionID: start.SessionID}
		}
		if start.PreRoll > bufferWindow {
			return RpcResponse{Message: fmt.Sprintf("pre-roll %s exceeds buffer window %s", start.PreRoll, bufferWindow), SessionID: start.SessionID}
		}
	}

	job := &recordJob{done: make(chan struct{})}

	c.mu.Lock()
	if _, exists := c.jobs[start.SessionID]; exists {
		c.mu.Unlock()
		log.Warn().Str("session", start.SessionID).Msg("start command rejected: already running")
		return RpcResponse{Message: fmt.Sprintf("session %s already running", start.SessionID), SessionID: start.SessionID}
	}
	ctx, cancel := context.WithCancel(context.Background())
	job.cancel = cancel
	c.jobs[start.SessionID] = job
	c.mu.Unlock()

	started := time.Now().UTC()
	state := StateSnapshot{
		SessionID:    start.SessionID,
		DeviceID:     start.DeviceID,
		Subject:      start.Subject,
		Status:       "running",
		MessageCount: 0,
		StartedAt:    started,
	}
	_ = c.writeState(state)

	go c.runRecord(ctx, job, start, started)

	log.Info().Str("session", start.SessionID).Str("device", start.DeviceID).Msg("recording job launched")
	return RpcResponse{OK: true, Message: "recording started", SessionID: start.SessionID, State: &state}
}

func (c *Controller) runRecord(ctx context.Context, job *recordJob, start startCommand, started time.Time) {
	defer func() {
		close(job.done)
		c.mu.Lock()
		delete(c.jobs, start.SessionID)
		c.mu.Unlock()
	}()

	opts := session.RecordOptions{
		ServerURL:     c.opts.ServerURL,
		Subject:       start.Subject,
		DeviceID:      start.DeviceID,
		SessionID:     start.SessionID,
		Retention:     start.Retention,
		SessionBucket: start.SessionBucket,
		DeviceBucket:  start.DeviceBucket,
		Device:        start.Device,
		Verbose:       start.Verbose,
		PreRoll:       start.PreRoll,
	}

	opts.Progress = func(meta session.Metadata) {
		snap := StateSnapshot{
			SessionID:     meta.SessionID,
			DeviceID:      meta.DeviceID,
			Subject:       meta.SourceSubject,
			Status:        "running",
			MessageCount:  meta.MessageCount,
			StartedAt:     started,
			LastMessageAt: meta.End,
		}
		_ = c.writeState(snap)
	}

	meta, err := session.Record(ctx, opts)

	state := StateSnapshot{
		SessionID:     start.SessionID,
		DeviceID:      start.DeviceID,
		Subject:       start.Subject,
		StartedAt:     started,
		LastMessageAt: time.Now().UTC(),
	}

	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			state.Status = "stopped"
		} else {
			state.Status = "error"
			state.LastError = err.Error()
		}
	} else {
		state.Status = "stopped"
		if meta != nil {
			state.MessageCount = meta.MessageCount
			state.DeviceID = meta.DeviceID
			state.Subject = meta.SourceSubject
			state.LastMessageAt = meta.End
		}
	}

	_ = c.writeState(state)
}

func (c *Controller) lookupBufferWindow(bucket, deviceID string) (time.Duration, error) {
	devStore, err := store.NewDeviceStore(c.js, bucket)
	if err != nil {
		return 0, fmt.Errorf("buffer lookup: %w", err)
	}
	info, err := devStore.Get(deviceID)
	if err != nil {
		return 0, fmt.Errorf("device buffer not configured")
	}
	if info.BufferDuration == "" {
		return 0, fmt.Errorf("device buffer not configured")
	}
	dur, err := time.ParseDuration(info.BufferDuration)
	if err != nil {
		return 0, fmt.Errorf("invalid device buffer duration: %v", err)
	}
	if dur <= 0 {
		return 0, fmt.Errorf("device buffer duration not positive")
	}
	return dur, nil
}

func (c *Controller) handleStop(req RpcRequest) RpcResponse {
	sessionID := strings.TrimSpace(req.SessionID)
	if sessionID == "" {
		return RpcResponse{Message: "session-id cannot be empty"}
	}

	c.mu.Lock()
	job, exists := c.jobs[sessionID]
	c.mu.Unlock()

	if !exists {
		// nothing running, but update state to stopped
		snap, err := c.loadState(sessionID)
		if err != nil {
			log.Error().Err(err).Str("session", sessionID).Msg("load state failed")
			return RpcResponse{Message: fmt.Sprintf("load state: %v", err), SessionID: sessionID}
		}
		snap.Status = "stopped"
		snap.LastError = ""
		if snap.StartedAt.IsZero() {
			snap.StartedAt = time.Now().UTC()
		}
		_ = c.writeState(snap)
		return RpcResponse{OK: true, SessionID: sessionID, Message: "no active recording"}
	}

	job.cancel()
	<-job.done

	log.Info().Str("session", sessionID).Msg("recording job signaled to stop")
	return RpcResponse{OK: true, SessionID: sessionID, Message: "recording stopped"}
}

func (c *Controller) respond(msg *nats.Msg, resp RpcResponse) {
	if !resp.OK && resp.Message == "" {
		resp.Message = "command failed"
	}
	data, _ := json.Marshal(resp)
	log.Debug().Str("session", resp.SessionID).Bool("ok", resp.OK).Msg("command response")
	_ = msg.Respond(data)
}

func (c *Controller) respondError(msg *nats.Msg, format string, args ...any) {
	resp := RpcResponse{OK: false, Message: fmt.Sprintf(format, args...)}
	data, _ := json.Marshal(resp)
	log.Error().Msgf(format, args...)
	_ = msg.Respond(data)
}

func (c *Controller) writeState(state StateSnapshot) error {
	if state.SessionID == "" {
		return errors.New("state missing session id")
	}
	if state.Subject == "" || state.DeviceID == "" {
		prev, err := c.loadState(state.SessionID)
		if err == nil {
			if state.Subject == "" {
				state.Subject = prev.Subject
			}
			if state.DeviceID == "" {
				state.DeviceID = prev.DeviceID
			}
			if state.MessageCount == 0 {
				state.MessageCount = prev.MessageCount
			}
			if state.StartedAt.IsZero() {
				state.StartedAt = prev.StartedAt
			}
			if state.LastMessageAt.IsZero() {
				state.LastMessageAt = prev.LastMessageAt
			}
		}
	}
	state.UpdatedAt = time.Now().UTC()
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	_, err = c.stateKV.Put(context.Background(), state.SessionID, data)
	return err
}

func (c *Controller) loadState(sessionID string) (StateSnapshot, error) {
	entry, err := c.stateKV.Get(context.Background(), sessionID)
	if err != nil {
		if errors.Is(err, jetstream.ErrKeyNotFound) {
			return StateSnapshot{SessionID: sessionID}, nil
		}
		return StateSnapshot{}, err
	}
	var snap StateSnapshot
	err = json.Unmarshal(entry.Value(), &snap)
	if err != nil {
		return StateSnapshot{}, err
	}
	return snap, nil
}

func parseRetention(value string) (time.Duration, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, nil
	}
	d, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("invalid retention duration: %w", err)
	}
	return d, nil
}
