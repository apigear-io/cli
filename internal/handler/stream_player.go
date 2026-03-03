package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/apigear-io/cli/pkg/stream/tracing"
	"github.com/go-chi/chi/v5"
)

// PlayerStream represents an active playback stream
type PlayerStream struct {
	ID           string                  `json:"id"`
	ProxyName    string                  `json:"proxyName"`
	Filename     string                  `json:"filename"`
	Speed        float64                 `json:"speed"`
	Loop         bool                    `json:"loop"`
	Direction    string                  `json:"direction"` // "", "SEND", "RECV"
	State        tracing.PlayerState     `json:"state"`
	Position     int                     `json:"position"`
	TotalEntries int                     `json:"totalEntries"`
	Progress     float64                 `json:"progress"`
	player       *tracing.Player         `json:"-"`
}

// CreatePlayerStreamRequest represents the request to create a new player stream
type CreatePlayerStreamRequest struct {
	ProxyName    string  `json:"proxyName"`
	Filename     string  `json:"filename"`
	Speed        float64 `json:"speed"`
	InitialDelay int     `json:"initialDelay"` // ms
	Loop         bool    `json:"loop"`
	Direction    string  `json:"direction"` // "", "SEND", "RECV"
}

var (
	playerStreams   = make(map[string]*PlayerStream)
	playerStreamsMu sync.RWMutex
	nextStreamID    = 1
	nextStreamIDMu  sync.Mutex
)

func getNextStreamID() string {
	nextStreamIDMu.Lock()
	defer nextStreamIDMu.Unlock()
	id := fmt.Sprintf("stream-%d", nextStreamID)
	nextStreamID++
	return id
}

// ListPlayerStreams returns all active player streams
func ListPlayerStreams() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerStreamsMu.RLock()
		defer playerStreamsMu.RUnlock()

		streams := make([]PlayerStream, 0, len(playerStreams))
		for _, stream := range playerStreams {
			// Update state from player
			if stream.player != nil {
				stream.State = stream.player.GetState()
				stream.Position = stream.player.GetPosition()
				stream.Progress = stream.player.GetProgress()
			}
			streams = append(streams, *stream)
		}

		writeJSON(w, http.StatusOK, streams)
	}
}

// CreatePlayerStream creates a new player stream
func CreatePlayerStream() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreatePlayerStreamRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid request body")
			return
		}

		// Validate request
		if req.ProxyName == "" {
			writeError(w, http.StatusBadRequest, nil, "proxyName is required")
			return
		}
		if req.Filename == "" {
			writeError(w, http.StatusBadRequest, nil, "filename is required")
			return
		}
		if req.Speed <= 0 {
			req.Speed = 1.0
		}

		// Get services
		services := getStreamServices()
		if services == nil {
			writeError(w, http.StatusInternalServerError, nil, "Stream services not initialized")
			return
		}

		// Build absolute path to trace file
		traceFile := filepath.Join(tracing.GetTraceDir(), req.Filename)

		// Create filter options
		filter := tracing.FilterOptions{}
		if req.Direction != "" {
			filter.Direction = req.Direction
		}

		// Read trace file to count entries
		entries, err := tracing.ReadTraceFileFiltered(traceFile, filter)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to read trace file")
			return
		}

		// Create player
		player, err := tracing.NewPlayer(traceFile, tracing.PlayerOptions{
			Speed:  req.Speed,
			Loop:   req.Loop,
			Filter: filter,
		})
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to create player")
			return
		}

		// Create stream
		streamID := getNextStreamID()
		stream := &PlayerStream{
			ID:           streamID,
			ProxyName:    req.ProxyName,
			Filename:     req.Filename,
			Speed:        req.Speed,
			Loop:         req.Loop,
			Direction:    req.Direction,
			State:        player.GetState(),
			Position:     player.GetPosition(),
			TotalEntries: len(entries),
			Progress:     player.GetProgress(),
			player:       player,
		}

		// Store stream
		playerStreamsMu.Lock()
		playerStreams[streamID] = stream
		playerStreamsMu.Unlock()

		writeJSON(w, http.StatusCreated, stream)
	}
}

// GetPlayerStream returns a specific player stream
func GetPlayerStream() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		streamID := chi.URLParam(r, "id")

		playerStreamsMu.RLock()
		stream, ok := playerStreams[streamID]
		playerStreamsMu.RUnlock()

		if !ok {
			writeError(w, http.StatusNotFound, nil, "Stream not found")
			return
		}

		// Update state from player
		if stream.player != nil {
			stream.State = stream.player.GetState()
			stream.Position = stream.player.GetPosition()
			stream.Progress = stream.player.GetProgress()
		}

		writeJSON(w, http.StatusOK, stream)
	}
}

// PlayPlayerStream starts playback
func PlayPlayerStream() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		streamID := chi.URLParam(r, "id")

		playerStreamsMu.RLock()
		stream, ok := playerStreams[streamID]
		playerStreamsMu.RUnlock()

		if !ok {
			writeError(w, http.StatusNotFound, nil, "Stream not found")
			return
		}

		if stream.player == nil {
			writeError(w, http.StatusInternalServerError, nil, "Player not initialized")
			return
		}

		// TODO: Connect player to proxy and send messages
		// For now, just start the player
		if err := stream.player.Play(); err != nil {
			writeError(w, http.StatusBadRequest, err, "Failed to start playback")
			return
		}

		stream.State = stream.player.GetState()
		writeJSON(w, http.StatusOK, stream)
	}
}

// PausePlayerStream pauses playback
func PausePlayerStream() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		streamID := chi.URLParam(r, "id")

		playerStreamsMu.RLock()
		stream, ok := playerStreams[streamID]
		playerStreamsMu.RUnlock()

		if !ok {
			writeError(w, http.StatusNotFound, nil, "Stream not found")
			return
		}

		if stream.player == nil {
			writeError(w, http.StatusInternalServerError, nil, "Player not initialized")
			return
		}

		stream.player.Pause()
		stream.State = stream.player.GetState()
		writeJSON(w, http.StatusOK, stream)
	}
}

// ResumePlayerStream resumes playback
func ResumePlayerStream() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		streamID := chi.URLParam(r, "id")

		playerStreamsMu.RLock()
		stream, ok := playerStreams[streamID]
		playerStreamsMu.RUnlock()

		if !ok {
			writeError(w, http.StatusNotFound, nil, "Stream not found")
			return
		}

		if stream.player == nil {
			writeError(w, http.StatusInternalServerError, nil, "Player not initialized")
			return
		}

		stream.player.Resume()
		stream.State = stream.player.GetState()
		writeJSON(w, http.StatusOK, stream)
	}
}

// StopPlayerStream stops playback
func StopPlayerStream() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		streamID := chi.URLParam(r, "id")

		playerStreamsMu.RLock()
		stream, ok := playerStreams[streamID]
		playerStreamsMu.RUnlock()

		if !ok {
			writeError(w, http.StatusNotFound, nil, "Stream not found")
			return
		}

		if stream.player == nil {
			writeError(w, http.StatusInternalServerError, nil, "Player not initialized")
			return
		}

		stream.player.Stop()
		stream.State = stream.player.GetState()
		stream.Position = stream.player.GetPosition()
		writeJSON(w, http.StatusOK, stream)
	}
}

// DeletePlayerStream deletes a player stream
func DeletePlayerStream() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		streamID := chi.URLParam(r, "id")

		playerStreamsMu.Lock()
		stream, ok := playerStreams[streamID]
		if ok {
			if stream.player != nil {
				stream.player.Stop()
			}
			delete(playerStreams, streamID)
		}
		playerStreamsMu.Unlock()

		if !ok {
			writeError(w, http.StatusNotFound, nil, "Stream not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
