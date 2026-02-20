package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/apigear-io/cli/pkg/stream/scripting"
	"github.com/apigear-io/cli/pkg/stream/tracing"
)

// Generator generates trace files using JavaScript templates with faker functions.
type Generator struct {
	templateDir string
}

// NewGenerator creates a new trace generator.
func NewGenerator(templateDir string) *Generator {
	return &Generator{
		templateDir: templateDir,
	}
}

// GenerateRequest contains parameters for trace generation.
type GenerateRequest struct {
	Template string `json:"template"` // JavaScript template code
	Count    int    `json:"count"`    // Number of entries to generate
	Filename string `json:"filename"` // Output filename (optional for preview)
}

// GenerateResult contains the generated trace entries.
type GenerateResult struct {
	Entries []json.RawMessage `json:"entries"`
	Count   int               `json:"count"`
}

// Generate generates trace entries using a JavaScript template.
func (g *Generator) Generate(req GenerateRequest) (*GenerateResult, error) {
	if req.Count <= 0 {
		return nil, fmt.Errorf("count must be positive")
	}
	if req.Count > 10000 {
		return nil, fmt.Errorf("count too large (max 10000)")
	}

	// Create a scripting engine
	engine := scripting.NewEngine("generator", "trace-generator")

	var entries []json.RawMessage
	var generateErr error

	// Wrap template in a function that returns an array
	script := fmt.Sprintf(`
		// User template
		%s

		// Generate entries
		const __entries = [];
		for (let i = 0; i < %d; i++) {
			try {
				const result = generate();
				__entries.push(result);
			} catch (err) {
				throw new Error("Error in generate() at index " + i + ": " + err.message);
			}
		}
		__entries;
	`, req.Template, req.Count)

	// Run the script
	result, err := engine.RunWithResult(script)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	// Export the result from Goja value to Go value
	exportedResult := result.Export()

	// Convert to slice
	rawEntries, ok := exportedResult.([]interface{})
	if !ok {
		return nil, fmt.Errorf("template must return an array")
	}

	// Convert to trace entries
	for _, entry := range rawEntries {
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			generateErr = fmt.Errorf("failed to marshal entry: %w", err)
			break
		}
		entries = append(entries, entryJSON)
	}

	if generateErr != nil {
		return nil, generateErr
	}

	return &GenerateResult{
		Entries: entries,
		Count:   len(entries),
	}, nil
}

// SaveAsTrace saves generated entries as a trace file in JSONL format.
func (g *Generator) SaveAsTrace(proxyName, filename string, entries []json.RawMessage) error {
	// Get trace directory
	traceDir := tracing.GetTraceDir()

	// Ensure trace directory exists
	if err := os.MkdirAll(traceDir, 0755); err != nil {
		return fmt.Errorf("failed to create trace directory: %w", err)
	}

	// Create trace file
	tracePath := filepath.Join(traceDir, filename)
	file, err := os.Create(tracePath)
	if err != nil {
		return fmt.Errorf("failed to create trace file: %w", err)
	}
	defer file.Close()

	// Write entries in JSONL format
	now := time.Now()
	for i, entryData := range entries {
		// Create trace entry
		entry := map[string]interface{}{
			"ts":    now.Add(time.Duration(i) * time.Millisecond).UnixMilli(),
			"dir":   "SEND",
			"proxy": proxyName,
			"msg":   json.RawMessage(entryData),
		}

		// Marshal and write
		data, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal trace entry: %w", err)
		}

		if _, err := file.Write(data); err != nil {
			return fmt.Errorf("failed to write trace entry: %w", err)
		}
		if _, err := file.WriteString("\n"); err != nil {
			return fmt.Errorf("failed to write newline: %w", err)
		}
	}

	return nil
}

// SaveTemplate saves a template to the template directory.
func (g *Generator) SaveTemplate(name, template string) error {
	// Ensure template directory exists
	if err := os.MkdirAll(g.templateDir, 0755); err != nil {
		return fmt.Errorf("failed to create template directory: %w", err)
	}

	// Save template file
	templatePath := filepath.Join(g.templateDir, name+".js")
	if err := os.WriteFile(templatePath, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to write template file: %w", err)
	}

	return nil
}

// LoadTemplate loads a template from the template directory.
func (g *Generator) LoadTemplate(name string) (string, error) {
	templatePath := filepath.Join(g.templateDir, name+".js")
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}

	return string(data), nil
}

// ListTemplates returns a list of available templates.
func (g *Generator) ListTemplates() ([]string, error) {
	// Ensure template directory exists
	if err := os.MkdirAll(g.templateDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create template directory: %w", err)
	}

	entries, err := os.ReadDir(g.templateDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read template directory: %w", err)
	}

	var templates []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".js" {
			name := entry.Name()[:len(entry.Name())-3] // Remove .js extension
			templates = append(templates, name)
		}
	}

	return templates, nil
}

// GetExamples returns example templates.
func GetExamples() map[string]string {
	return map[string]string{
		"simple": `// Simple counter example
function generate() {
  return {
    id: faker.int(1, 1000),
    message: "Hello " + faker.firstName()
  };
}`,
		"objectlink": `// ObjectLink message example
function generate() {
  return {
    type: 2,
    id: faker.int(1, 1000),
    path: "demo.Counter/count",
    value: faker.int(0, 100)
  };
}`,
		"user": `// User data example
function generate() {
  return {
    id: faker.uuid(),
    name: faker.name(),
    email: faker.email(),
    age: faker.int(18, 80),
    city: faker.city(),
    country: faker.country()
  };
}`,
		"sensor": `// IoT sensor data example
function generate() {
  return {
    sensor_id: faker.uuid(),
    timestamp: Date.now(),
    temperature: faker.float(15, 35),
    humidity: faker.float(30, 90),
    pressure: faker.float(980, 1050)
  };
}`,
	}
}
