package msgio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

// GenerateOptions controls how JSONL data is synthesized from a template.
type GenerateOptions struct {
	TemplatePath string
	OutputPath   string
	Count        int
	Seed         int64
}

// Generate renders the template Count times, writing JSONL output to the destination.
func Generate(opts GenerateOptions) error {
	if err := validateOptions(opts); err != nil {
		return err
	}

	tpl, err := parseTemplate(opts)
	if err != nil {
		return err
	}

	writer, closeFn, err := openOutput(opts.OutputPath)
	if err != nil {
		return err
	}
	if closeFn != nil {
		defer closeFn()
	}

	return renderRecords(opts.Count, tpl, writer)
}

// validateOptions checks that the provided options are valid.
func validateOptions(opts GenerateOptions) error {
	if opts.Count <= 0 {
		return errors.New("count must be positive")
	}
	if opts.TemplatePath == "" {
		return errors.New("template path cannot be empty")
	}
	return nil
}

// parseTemplate reads and parses the template file specified in opts.
func parseTemplate(opts GenerateOptions) (*template.Template, error) {
	tplData, err := os.ReadFile(opts.TemplatePath)
	if err != nil {
		return nil, fmt.Errorf("read template: %w", err)
	}

	tpl, err := template.New(filepath.Base(opts.TemplatePath)).Funcs(newTemplateFuncMap(opts)).Parse(string(tplData))
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}
	return tpl, nil
}

func openOutput(path string) (io.Writer, func() error, error) {
	if path == "" || path == "-" {
		return os.Stdout, nil, nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, nil, fmt.Errorf("create output dir: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, nil, fmt.Errorf("create output: %w", err)
	}
	return f, f.Close, nil
}

// renderRecords renders count records using the provided template and writes them to the writer.
func renderRecords(count int, tpl *template.Template, writer io.Writer) error {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < count; i++ {
		line, err := renderRecord(tpl, buf, i)
		if err != nil {
			return err
		}
		if line == "" {
			continue
		}
		if _, err := writer.Write([]byte(line)); err != nil {
			return fmt.Errorf("write record %d: %w", i+1, err)
		}
	}
	return nil
}

// renderRecord renders a single record using the provided template and buffer.
func renderRecord(tpl *template.Template, buf *bytes.Buffer, index int) (string, error) {
	buf.Reset()
	if err := tpl.Execute(buf, nil); err != nil {
		return "", fmt.Errorf("render record %d: %w", index+1, err)
	}
	line := strings.TrimSpace(buf.String())
	if line == "" {
		return "", nil
	}
	var compact bytes.Buffer
	if err := json.Compact(&compact, []byte(line)); err != nil {
		return "", fmt.Errorf("record %d: invalid JSON output: %w", index+1, err)
	}
	output := compact.String()
	if !strings.HasSuffix(output, "\n") {
		output += "\n"
	}
	return output, nil
}

// newTemplateFuncMap returns a map of functions to be used in templates.
func newTemplateFuncMap(opts GenerateOptions) template.FuncMap {
	seed := uint64(opts.Seed)
	if opts.Seed < 0 {
		seed = uint64(math.Abs(float64(opts.Seed)))
	}
	faker := gofakeit.New(seed)

	seq := 0

	return template.FuncMap{
		"faker": func(path string) (string, error) {
			return fakerValue(faker, path)
		},
		"seq": func() int {
			seq++
			return seq
		},
		"uuid": func() string {
			return uuid.NewString()
		},
		"timestamp": func(layout ...string) string {
			format := time.RFC3339Nano
			if len(layout) > 0 && layout[0] != "" {
				format = layout[0]
			}
			return time.Now().UTC().Format(format)
		},
		"randInt": func(min, max int) (int, error) {
			if min > max {
				return 0, fmt.Errorf("randInt min greater than max")
			}
			return faker.Number(min, max), nil
		},
		"randFloat": func(min, max float64) (float64, error) {
			if min > max {
				return 0, fmt.Errorf("randFloat min greater than max")
			}
			return faker.Float64Range(min, max), nil
		},
	}
}

// fakerValue generates a fake value based on the provided path.
func fakerValue(faker *gofakeit.Faker, path string) (string, error) {
	parts := strings.Split(strings.ToLower(path), ".")
	if len(parts) == 0 {
		return "", errors.New("empty faker path")
	}

	switch parts[0] {
	case "person":
		if len(parts) == 1 || parts[1] == "name" {
			return faker.Name(), nil
		}
		switch parts[1] {
		case "first_name":
			return faker.FirstName(), nil
		case "last_name":
			return faker.LastName(), nil
		case "ssn":
			return faker.SSN(), nil
		case "phone":
			return faker.Phone(), nil
		case "email":
			return faker.Email(), nil
		case "job":
			return faker.JobTitle(), nil
		}
	case "address":
		addr := faker.Address()
		if len(parts) == 1 {
			return fmt.Sprintf("%s, %s, %s %s", addr.Street, addr.City, addr.State, addr.Zip), nil
		}
		switch parts[1] {
		case "street":
			return addr.Street, nil
		case "city":
			return addr.City, nil
		case "state":
			return addr.State, nil
		case "zip":
			return addr.Zip, nil
		case "country":
			return addr.Country, nil
		case "latitude":
			return fmt.Sprintf("%.6f", addr.Latitude), nil
		case "longitude":
			return fmt.Sprintf("%.6f", addr.Longitude), nil
		}
	case "internet":
		if len(parts) == 1 {
			return faker.DomainName(), nil
		}
		switch parts[1] {
		case "email":
			return faker.Email(), nil
		case "user":
			return faker.Username(), nil
		case "domain":
			return faker.DomainName(), nil
		case "ipv4":
			return faker.IPv4Address(), nil
		case "ipv6":
			return faker.IPv6Address(), nil
		case "url":
			return faker.URL(), nil
		}
	case "company":
		if len(parts) == 1 {
			return faker.Company(), nil
		}
		switch parts[1] {
		case "name":
			return faker.Company(), nil
		case "bs":
			return faker.BS(), nil
		case "buzzword":
			return faker.BuzzWord(), nil
		case "slogan":
			return faker.Slogan(), nil
		}
	case "lorem":
		if len(parts) == 1 {
			return strings.Join(generateWords(faker, 3), " "), nil
		}
		switch parts[1] {
		case "word":
			return faker.Word(), nil
		case "words":
			count := 3
			if len(parts) > 2 {
				if n, err := strconv.Atoi(parts[2]); err == nil {
					count = n
				}
			}
			return strings.Join(generateWords(faker, count), " "), nil
		case "sentence":
			n := 12
			if len(parts) > 2 {
				if v, err := strconv.Atoi(parts[2]); err == nil {
					n = v
				}
			}
			return faker.Sentence(n), nil
		case "paragraph":
			n := 3
			if len(parts) > 2 {
				if v, err := strconv.Atoi(parts[2]); err == nil {
					n = v
				}
			}
			return faker.Paragraph(n, 3, 12, " "), nil
		}
	case "uuid":
		return uuid.NewString(), nil
	case "boolean":
		return fmt.Sprintf("%t", faker.Bool()), nil
	case "date":
		return faker.Date().Format(time.RFC3339Nano), nil
	case "number":
		min, max := 0, 100
		if len(parts) > 1 {
			if v, err := strconv.Atoi(parts[1]); err == nil {
				min = v
			}
		}
		if len(parts) > 2 {
			if v, err := strconv.Atoi(parts[2]); err == nil {
				max = v
			}
		}
		if min > max {
			return "", fmt.Errorf("number: min greater than max")
		}
		return fmt.Sprintf("%d", faker.Number(min, max)), nil
	}

	return "", fmt.Errorf("unsupported faker path %q", path)
}

// generateWords generates a slice of fake words of the specified count.
func generateWords(faker *gofakeit.Faker, count int) []string {
	if count <= 0 {
		return []string{}
	}
	words := make([]string, count)
	for i := range words {
		words[i] = faker.Word()
	}
	return words
}
