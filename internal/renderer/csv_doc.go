// Package renderer provides HTML and CSV rendering for driftctl analysis reports.
//
// # CSV Export
//
// The CSV export feature converts a parsed driftctl [parser.Analysis] into a
// flat, tabular representation suitable for spreadsheet tools or downstream
// data pipelines.
//
// Each row in the output represents a single cloud resource and carries three
// columns:
//
//   - resource_type – the Terraform resource type (e.g. "aws_s3_bucket")
//   - resource_id   – the unique identifier of the resource
//   - status        – one of "managed", "unmanaged", or "deleted"
//
// Usage:
//
//	rows := renderer.BuildCSVRows(analysis)
//	if err := renderer.WriteCSV(os.Stdout, rows, ','); err != nil {
//	    log.Fatal(err)
//	}
//
// The delimiter can be changed via [WithCSVDelimiter]. A semicolon is common
// for locales where comma is the decimal separator.
//
// CSV export is opt-in and controlled by [WithCSV]. The output filename is
// configured via [WithCSVFilename] and defaults to "report.csv".
package renderer
