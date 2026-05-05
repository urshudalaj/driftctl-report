package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun_MissingInputFlag(t *testing.T) {
	// run() reads os.Args via flag; we test the guard directly by calling
	// the helper with an empty input path to ensure it returns an error.
	err := runWithArgs("", "")
	if err == nil {
		t.Fatal("expected error when input flag is empty, got nil")
	}
}

func TestRun_InvalidInputFile(t *testing.T) {
	err := runWithArgs("/nonexistent/path/report.json", "")
	if err == nil {
		t.Fatal("expected error for missing input file, got nil")
	}
}

func TestRun_ValidInput(t *testing.T) {
	tmpOut := filepath.Join(t.TempDir(), "out.html")

	err := runWithArgs("../../testdata/sample_report.json", tmpOut)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	info, err := os.Stat(tmpOut)
	if err != nil {
		t.Fatalf("output file not created: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("output file is empty")
	}
}

// runWithArgs is a testable variant of run that accepts explicit paths.
func runWithArgs(inputFile, outputFile string) error {
	if inputFile == "" {
		return fmt.Errorf("--input flag is required")
	}
	if outputFile == "" {
		outputFile = filepath.Join(os.TempDir(), "drift-report-test.html")
	}

	import (
		"fmt"

		"github.com/yourorg/driftctl-report/internal/parser"
		"github.com/yourorg/driftctl-report/internal/renderer"
	)

	report, err := parser.ParseFile(inputFile)
	if err != nil {
		return fmt.Errorf("parsing input file %q: %w", inputFile, err)
	}

	out, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("creating output file %q: %w", outputFile, err)
	}
	defer out.Close()

	r := renderer.New(report)
	return r.Render(out)
}
