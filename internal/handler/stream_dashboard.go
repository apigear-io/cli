package handler

import (
	"net/http"

	"github.com/apigear-io/cli/pkg/stream/client"
	"github.com/apigear-io/cli/pkg/stream/proxy"
)

// GetStreamDashboard returns dashboard statistics
// @Summary Get stream dashboard statistics
// @Description Get overall statistics for all proxies and clients
// @Tags stream
// @Produce json
// @Success 200 {object} StreamDashboardStats
// @Router /api/v1/stream/dashboard [get]
func GetStreamDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services := getStreamServices()

		// Get proxy stats
		proxies := services.ProxyManager.ListProxies()
		var runningProxies, stoppedProxies int
		for _, p := range proxies {
			if p.Status == proxy.StatusRunning {
				runningProxies++
			} else {
				stoppedProxies++
			}
		}

		// Get client stats
		clients := services.ClientManager.ListClients()
		var connectedClients, disconnectedClients int
		for _, c := range clients {
			if c.Status == client.StatusConnected {
				connectedClients++
			} else {
				disconnectedClients++
			}
		}

		// Get message stats
		globalStats := services.Stats.GlobalStats()

		stats := StreamDashboardStats{}
		stats.Proxies.Total = len(proxies)
		stats.Proxies.Running = runningProxies
		stats.Proxies.Stopped = stoppedProxies
		stats.Clients.Total = len(clients)
		stats.Clients.Connected = connectedClients
		stats.Clients.Disconnected = disconnectedClients
		stats.Messages.Total = globalStats.MessagesReceived + globalStats.MessagesSent

		// Calculate rate (messages per second)
		if globalStats.Uptime > 0 {
			stats.Messages.Rate = float64(stats.Messages.Total) / float64(globalStats.Uptime)
		}

		writeJSON(w, http.StatusOK, stats)
	}
}
