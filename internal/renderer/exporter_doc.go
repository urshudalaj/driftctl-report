// Package renderer provides HTML rendering and export functionality for
// driftctl JSON reports.
//
// # Exporter
//
// The Exporter type wraps a Renderer and handles writing the rendered HTML
// to any io.Writer or directly to a file on disk.
//
// Basic usage:
//
//	report, _ := parser.ParseFile("scan.json")
//	r := renderer.New(report, renderer.WithTitle("My Infra Drift"))
//	e := renderer.NewExporter(r)
//
//	// Write to a file (parent directories are created automatically):
//	if err := e.ExportToFile("reports/drift.html"); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Or write to any io.Writer:
//	if err := e.ExportTo(os.Stdout); err != nil {
//	    log.Fatal(err)
//	}
package renderer
