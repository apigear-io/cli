package handler

import (
	"net/http"
	"runtime"
	"time"

	"github.com/apigear-io/cli/pkg/foundation/config"
)

var startTime = time.Now()

// StatusResponse represents the status information response
type StatusResponse struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Uptime    string `json:"uptime"`
}

// Status godoc
// @Summary Status and build information endpoint
// @Description Returns status and build information for the API server
// @Tags system
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /status [get]
func Status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buildInfo := config.GetBuildInfo("cli")
		uptime := time.Since(startTime)

		writeJSON(w, http.StatusOK, StatusResponse{
			Version:   buildInfo.Version,
			Commit:    buildInfo.Commit,
			BuildDate: buildInfo.Date,
			GoVersion: runtime.Version(),
			Uptime:    uptime.String(),
		})
	}
}
