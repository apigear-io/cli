package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

// TestCLIRegression runs end-to-end tests for the CLI to ensure
// the user-facing API (commands, args, flags) doesn't change during refactoring.
func TestCLIRegression(t *testing.T) {
	// Build the binary before running tests
	binPath := buildBinary(t)

	testscript.Run(t, testscript.Params{
		Dir: "testscripts",
		Setup: func(env *testscript.Env) error {
			// Make the apigear binary available in test scripts
			env.Setenv("PATH", filepath.Dir(binPath)+string(os.PathListSeparator)+env.Getenv("PATH"))
			// Set HOME to test work directory to avoid polluting user's home
			env.Setenv("HOME", env.WorkDir)
			return nil
		},
	})
}

// buildBinary builds the apigear binary and returns its path
func buildBinary(t *testing.T) string {
	t.Helper()

	// Create a temporary directory for the binary
	tmpDir := t.TempDir()
	binPath := filepath.Join(tmpDir, "apigear")
	if os.Getenv("GOOS") == "windows" {
		binPath += ".exe"
	}

	// Build the binary
	cmd := exec.Command("go", "build", "-o", binPath, "./cmd/apigear")
	// Set the working directory to the project root (one level up from tests/)
	cmd.Dir = ".."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to build binary: %v", err)
	}

	return binPath
}
