package renderer

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExportTo_WritesHTML(t *testing.T) {
	r := New(sampleReport())
	e := NewExporter(r)

	var buf bytes.Buffer
	if err := e.ExportTo(&buf); err != nil {
		t.Fatalf("ExportTo returned unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "<html") {
		t.Error("expected output to contain <html tag")
	}
}

func TestExportToFile_CreatesFile(t *testing.T) {
	dir := t.TempDir()
	outPath := filepath.Join(dir, "sub", "report.html")

	r := New(sampleReport())
	e := NewExporter(r)

	if err := e.ExportToFile(outPath); err != nil {
		t.Fatalf("ExportToFile returned unexpected error: %v", err)
	}

	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("could not read output file: %v", err)
	}

	if !strings.Contains(string(data), "<html") {
		t.Error("expected file to contain <html tag")
	}
}

func TestExportToFile_EmptyPath(t *testing.T) {
	r := New(sampleReport())
	e := NewExporter(r)

	if err := e.ExportToFile(""); err == nil {
		t.Error("expected error for empty path, got nil")
	}
}

func TestExportTo_FailingWriter(t *testing.T) {
	r := New(sampleReport())
	e := NewExporter(r)

	fw := &failWriter{}
	if err := e.ExportTo(fw); err == nil {
		t.Error("expected error from failing writer, got nil")
	}
}

// TestExportToFile_ContentMatchesExportTo verifies that ExportToFile writes
// exactly the same content that ExportTo would produce for the same renderer.
func TestExportToFile_ContentMatchesExportTo(t *testing.T) {
	report := sampleReport()

	var buf bytes.Buffer
	if err := NewExporter(New(report)).ExportTo(&buf); err != nil {
		t.Fatalf("ExportTo returned unexpected error: %v", err)
	}

	dir := t.TempDir()
	outPath := filepath.Join(dir, "report.html")
	if err := NewExporter(New(report)).ExportToFile(outPath); err != nil {
		t.Fatalf("ExportToFile returned unexpected error: %v", err)
	}

	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("could not read output file: %v", err)
	}

	if buf.String() != string(data) {
		t.Error("ExportToFile content does not match ExportTo output")
	}
}

// failWriter always returns an error on Write.
type failWriter struct{}

func (fw *failWriter) Write(_ []byte) (int, error) {
	return 0, os.ErrClosed
}
