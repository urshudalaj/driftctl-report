// Package main is the entry point for the driftctl-report CLI tool.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yourorg/driftctl-report/internal/parser"
	"github.com/yourorg/driftctl-report/internal/renderer"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		inputFile  string
		outputFile string
	)

	flag.StringVar(&inputFile, "input", "", "Path to driftctl JSON report (required)")
	flag.StringVar(&outputFile, "output", "drift-report.html", "Path to write the HTML report")
	flag.Parse()

	if inputFile == "" {
		flag.Usage()
		return fmt.Errorf("--input flag is required")
	}

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
	if err := r.Render(out); err != nil {
		return fmt.Errorf("rendering HTML report: %w", err)
	}

	fmt.Printf("Report written to %s\n", outputFile)
	return nil
}
