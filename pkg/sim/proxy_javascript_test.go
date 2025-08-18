package sim

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestProxyJavaScript runs the JavaScript test files to ensure
// the proxy works correctly from pure JavaScript
func TestProxyJavaScript(t *testing.T) {
	testFiles := []struct {
		name        string
		file        string
		expectedMsg string
	}{
		{
			name:        "BasicProxyTests",
			file:        "testdata/proxy_test.js",
			expectedMsg: "ALL_TESTS_PASSED",
		},
		{
			name:        "EdgeCaseTests",
			file:        "testdata/proxy_edge_cases_test.js",
			expectedMsg: "ALL_EDGE_TESTS_PASSED",
		},
	}

	for _, tt := range testFiles {
		t.Run(tt.name, func(t *testing.T) {
			// Read the test script
			script, err := os.ReadFile(tt.file)
			require.NoError(t, err, "Failed to read test file %s", tt.file)

			// Create engine with proper working directory
			engine := NewEngine(EngineOptions{
				WorkDir: filepath.Dir(tt.file),
			})
			defer engine.Close()

			// Channel to capture test results
			done := make(chan struct {
				success bool
				result  string
				err     error
			})

			// Run the script
			engine.RunOnLoop(func(rt *goja.Runtime) {
				value, err := rt.RunString(string(script))
				
				result := struct {
					success bool
					result  string
					err     error
				}{
					err: err,
				}

				if err != nil {
					// Check if it's a test failure error
					if strings.Contains(err.Error(), "tests failed") {
						result.success = false
						result.result = err.Error()
					} else {
						// Actual JavaScript error
						result.err = err
					}
				} else if value != nil {
					// Check the return value
					resultStr := value.String()
					result.success = resultStr == tt.expectedMsg
					result.result = resultStr
				}

				done <- result
			})

			// Wait for test completion with timeout
			select {
			case result := <-done:
				if result.err != nil {
					t.Fatalf("JavaScript error: %v", result.err)
				}
				assert.True(t, result.success, 
					"Test failed. Expected '%s', got '%s'", tt.expectedMsg, result.result)
			case <-time.After(5 * time.Second):
				t.Fatal("Test timeout")
			}
		})
	}
}

// TestProxyJavaScriptInteractive runs a single JavaScript test file
// This is useful for debugging specific test failures
func TestProxyJavaScriptInteractive(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping interactive test in short mode")
	}

	// You can change this to test a specific file
	testFile := "testdata/proxy_test.js"

	script, err := os.ReadFile(testFile)
	require.NoError(t, err)

	engine := NewEngine(EngineOptions{
		WorkDir: "testdata",
	})
	defer engine.Close()

	// Capture console output for debugging
	outputChan := make(chan string, 100)
	
	engine.RunOnLoop(func(rt *goja.Runtime) {
		// Override console.log to capture output
		rt.Set("console", map[string]interface{}{
			"log": func(args ...interface{}) {
				output := ""
				for i, arg := range args {
					if i > 0 {
						output += " "
					}
					output += toString(arg)
				}
				outputChan <- output
				// Also print to test output
				t.Log(output)
			},
		})

		// Run the test
		value, err := rt.RunString(string(script))
		
		if err != nil {
			outputChan <- "ERROR: " + err.Error()
			t.Errorf("JavaScript error: %v", err)
		} else if value != nil {
			outputChan <- "RESULT: " + value.String()
		}
		
		close(outputChan)
	})

	// Wait for all output
	for output := range outputChan {
		if strings.HasPrefix(output, "ERROR:") {
			t.Error(output)
		}
	}
}

// toString converts an interface to string for logging
func toString(v interface{}) string {
	if v == nil {
		return "null"
	}
	switch val := v.(type) {
	case string:
		return val
	case bool:
		if val {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", val)
	}
}

// BenchmarkProxyOperations benchmarks various proxy operations
func BenchmarkProxyOperations(b *testing.B) {
	engine := NewEngine(EngineOptions{})
	defer engine.Close()

	b.Run("PropertyAccess", func(b *testing.B) {
		script := `
		const service = $createService("bench.Service", {
			value: 42
		});
		
		function benchmark() {
			let sum = 0;
			for (let i = 0; i < 1000; i++) {
				sum += service.value;
			}
			return sum;
		}
		`
		
		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			rt.RunString(script)
			fn, _ := goja.AssertFunction(rt.Get("benchmark"))
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				fn(nil)
			}
			done <- true
		})
		<-done
	})

	b.Run("MethodCall", func(b *testing.B) {
		script := `
		const service = $createService("bench.Service", {
			counter: 0
		});
		
		service.increment = function() {
			this.counter++;
		};
		
		function benchmark() {
			for (let i = 0; i < 1000; i++) {
				service.increment();
			}
		}
		`
		
		done := make(chan bool)
		engine.RunOnLoop(func(rt *goja.Runtime) {
			rt.RunString(script)
			fn, _ := goja.AssertFunction(rt.Get("benchmark"))
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				fn(nil)
			}
			done <- true
		})
		<-done
	})
}