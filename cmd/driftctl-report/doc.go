// Package main provides the driftctl-report command-line interface.
//
// Usage:
//
//	driftctl-report --input <path/to/driftctl.json> [--output <path/to/report.html>]
//
// Flags:
//
//	--input   Path to a driftctl JSON scan result file (required).
//	--output  Destination path for the generated HTML report.
//	          Defaults to "drift-report.html" in the current directory.
//
// Example:
//
//	driftctl scan --output json://result.json
//	driftctl-report --input result.json --output report.html
//
// The generated report includes a summary of managed, unmanaged, and missing
// resources as well as per-resource detail sections for easier auditing.
package main
