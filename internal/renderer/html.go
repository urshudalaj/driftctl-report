// Package renderer provides functionality for rendering driftctl reports
// into human-readable HTML format.
package renderer

import (
	"bytes"
	"html/template"
	"io"
	"time"

	"github.com/your-org/driftctl-report/internal/parser"
)

// TemplateData holds all data passed to the HTML template.
type TemplateData struct {
	Report    *parser.DriftReport
	GeneratedAt string
	Summary   Summary
}

// Summary contains high-level statistics derived from the report.
type Summary struct {
	TotalManaged    int
	TotalUnmanaged  int
	TotalMissing    int
	TotalDrifted    int
	CoveragePercent float64
}

// HTMLRenderer renders a DriftReport as an HTML document.
type HTMLRenderer struct {
	tmpl *template.Template
}

// New creates a new HTMLRenderer with the default embedded template.
func New() (*HTMLRenderer, error) {
	tmpl, err := template.New("report").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}).Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}
	return &HTMLRenderer{tmpl: tmpl}, nil
}

// Render writes the HTML report for the given DriftReport to w.
func (r *HTMLRenderer) Render(w io.Writer, report *parser.DriftReport) error {
	summary := buildSummary(report)
	data := TemplateData{
		Report:      report,
		GeneratedAt: time.Now().UTC().Format(time.RFC1123),
		Summary:     summary,
	}
	var buf bytes.Buffer
	if err := r.tmpl.Execute(&buf, data); err != nil {
		return err
	}
	_, err := w.Write(buf.Bytes())
	return err
}

// RenderToString renders the report and returns the HTML as a string.
func (r *HTMLRenderer) RenderToString(report *parser.DriftReport) (string, error) {
	var buf bytes.Buffer
	if err := r.Render(&buf, report); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func buildSummary(report *parser.DriftReport) Summary {
	s := Summary{
		TotalManaged:   len(report.ManagedResources),
		TotalUnmanaged: len(report.UnmanagedResources),
		TotalMissing:   len(report.MissingResources),
		TotalDrifted:   len(report.DifferentResources),
		CoveragePercent: report.Summary.Coverage,
	}
	return s
}
