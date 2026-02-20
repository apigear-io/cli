package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/apigear-io/cli/pkg/stream"
	"github.com/apigear-io/cli/pkg/stream/generator"
)

// GeneratorPreview godoc
// @Summary Preview generated trace entries
// @Description Generate trace entries using a JavaScript template with faker functions
// @Tags stream
// @Accept json
// @Produce json
// @Param request body generator.GenerateRequest true "Generation parameters"
// @Success 200 {object} generator.GenerateResult
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/generator/preview [post]
func GeneratorPreview(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req generator.GenerateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid request body")
			return
		}

		// Limit preview to 100 entries
		if req.Count > 100 {
			req.Count = 100
		}

		gen := generator.NewGenerator("./data/templates/generator")
		result, err := gen.Generate(req)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "Failed to generate trace entries")
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

// GeneratorSave godoc
// @Summary Save generated trace as file
// @Description Generate and save trace entries as a JSONL file
// @Tags stream
// @Accept json
// @Produce json
// @Param request body GeneratorSaveRequest true "Generation and save parameters"
// @Success 200 {object} GeneratorSaveResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/generator/save [post]
func GeneratorSave(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GeneratorSaveRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid request body")
			return
		}

		// Validate filename
		if req.Filename == "" {
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("filename is required"),
				"Filename must not be empty")
			return
		}

		// Default proxy name
		if req.ProxyName == "" {
			req.ProxyName = "generator"
		}

		gen := generator.NewGenerator("./data/templates/generator")

		// Generate entries
		genReq := generator.GenerateRequest{
			Template: req.Template,
			Count:    req.Count,
		}
		result, err := gen.Generate(genReq)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "Failed to generate trace entries")
			return
		}

		// Save as trace file
		if err := gen.SaveAsTrace(req.ProxyName, req.Filename, result.Entries); err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to save trace file")
			return
		}

		writeJSON(w, http.StatusOK, GeneratorSaveResponse{
			Filename: req.Filename,
			Count:    result.Count,
		})
	}
}

// GeneratorSaveRequest contains parameters for saving a generated trace.
type GeneratorSaveRequest struct {
	Template  string `json:"template"`
	Count     int    `json:"count"`
	ProxyName string `json:"proxyName"`
	Filename  string `json:"filename"`
}

// GeneratorSaveResponse contains the result of saving a trace.
type GeneratorSaveResponse struct {
	Filename string `json:"filename"`
	Count    int    `json:"count"`
}

// GeneratorSaveTemplate godoc
// @Summary Save a generator template
// @Description Save a JavaScript template for later use
// @Tags stream
// @Accept json
// @Produce json
// @Param request body GeneratorSaveTemplateRequest true "Template to save"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/generator/templates [post]
func GeneratorSaveTemplate(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GeneratorSaveTemplateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid request body")
			return
		}

		if req.Name == "" {
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("name is required"),
				"Template name must not be empty")
			return
		}

		gen := generator.NewGenerator("./data/templates/generator")
		if err := gen.SaveTemplate(req.Name, req.Template); err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to save template")
			return
		}

		writeJSON(w, http.StatusOK, SuccessResponse{
			Message: fmt.Sprintf("Template '%s' saved successfully", req.Name),
		})
	}
}

// GeneratorSaveTemplateRequest contains a template to save.
type GeneratorSaveTemplateRequest struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

// SuccessResponse is a generic success response.
type SuccessResponse struct {
	Message string `json:"message"`
}

// GeneratorLoadTemplate godoc
// @Summary Load a generator template
// @Description Load a saved JavaScript template
// @Tags stream
// @Produce json
// @Param name path string true "Template name"
// @Success 200 {object} GeneratorLoadTemplateResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/generator/templates/{name} [get]
func GeneratorLoadTemplate(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if name == "" {
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("name is required"),
				"Template name parameter must not be empty")
			return
		}

		gen := generator.NewGenerator("./data/templates/generator")
		template, err := gen.LoadTemplate(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "Template not found")
			return
		}

		writeJSON(w, http.StatusOK, GeneratorLoadTemplateResponse{
			Name:     name,
			Template: template,
		})
	}
}

// GeneratorLoadTemplateResponse contains a loaded template.
type GeneratorLoadTemplateResponse struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

// GeneratorListTemplates godoc
// @Summary List generator templates
// @Description List all saved JavaScript templates
// @Tags stream
// @Produce json
// @Success 200 {object} GeneratorListTemplatesResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/generator/templates [get]
func GeneratorListTemplates(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gen := generator.NewGenerator("./data/templates/generator")
		templates, err := gen.ListTemplates()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to list templates")
			return
		}

		writeJSON(w, http.StatusOK, GeneratorListTemplatesResponse{
			Templates: templates,
		})
	}
}

// GeneratorListTemplatesResponse contains a list of templates.
type GeneratorListTemplatesResponse struct {
	Templates []string `json:"templates"`
}

// GeneratorExamples godoc
// @Summary Get example templates
// @Description Get example generator templates
// @Tags stream
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/stream/generator/examples [get]
func GeneratorExamples(services *stream.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		examples := generator.GetExamples()
		writeJSON(w, http.StatusOK, examples)
	}
}
