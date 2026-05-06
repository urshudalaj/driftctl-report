package renderer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Exporter writes rendered HTML to a destination.
type Exporter struct {
	renderer *Renderer
}

// NewExporter creates an Exporter wrapping the given Renderer.
func NewExporter(r *Renderer) *Exporter {
	return &Exporter{renderer: r}
}

// ExportToFile renders the report and writes it to the file at path.
// The directory is created if it does not exist.
func (e *Exporter) ExportToFile(path string) error {
	if path == "" {
		return fmt.Errorf("exporter: output path must not be empty")
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("exporter: create directory %q: %w", dir, err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("exporter: create file %q: %w", path, err)
	}
	defer f.Close()

	return e.ExportTo(f)
}

// ExportTo renders the report and writes it to w.
func (e *Exporter) ExportTo(w io.Writer) error {
	html, err := e.renderer.Render()
	if err != nil {
		return fmt.Errorf("exporter: render: %w", err)
	}

	_, err = io.WriteString(w, html)
	if err != nil {
		return fmt.Errorf("exporter: write: %w", err)
	}

	return nil
}
