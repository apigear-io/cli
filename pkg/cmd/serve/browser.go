package serve

import (
	"os/exec"
	"runtime"
	"time"

	"github.com/apigear-io/cli/pkg/foundation/logging"
)

// openBrowser opens the specified URL in the default browser
func openBrowser(url string) {
	// Small delay to ensure server is fully started
	time.Sleep(500 * time.Millisecond)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		logging.Warn().Msgf("unsupported platform for auto-opening browser: %s", runtime.GOOS)
		return
	}

	if err := cmd.Start(); err != nil {
		logging.Warn().Err(err).Msg("failed to open browser automatically")
		return
	}

	logging.Info().Msgf("Opening browser at %s", url)
}
