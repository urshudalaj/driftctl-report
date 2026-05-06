// Package parser provides types and functions for reading and decoding
// driftctl JSON scan output files.
//
// Typical usage:
//
//	report, err := parser.ParseFile("driftctl-output.json")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Coverage: %.1f%%\n", report.Summary.Coverage)
//
// The package supports both file-based parsing via ParseFile and
// reader-based parsing via Parse, making it straightforward to use
// with embedded test fixtures or HTTP response bodies.
//
// Error handling:
//
// Both ParseFile and Parse return descriptive errors that wrap the
// underlying cause. Callers can use errors.Is and errors.As to
// inspect specific failure modes such as missing files or malformed
// JSON input.
package parser
