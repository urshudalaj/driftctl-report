// Package renderer — export options.
//
// # Export Options
//
// This file documents the export-related option functions that control how
// the rendered report is written to disk or streamed to a writer.
//
// Available options:
//
//   - WithOutputFormat(format string) — choose between "html" (default) and
//     "json" serialisation of the report data. Unknown format strings are
//     silently ignored so callers can safely pass user-supplied values.
//
//   - WithFilename(name string) — set the base filename (without extension)
//     used when the Exporter writes to disk via ExportToFile. Defaults to
//     "report".
//
//   - WithEmbedAssets(embed bool) — when true (default) all CSS and
//     JavaScript dependencies are inlined into the HTML output producing a
//     fully self-contained single file. Set to false to reference external
//     asset URLs instead, which reduces file size but requires network
//     access to render correctly.
//
// These options integrate with the Options struct defined in options.go and
// are applied via the standard functional-options pattern used throughout
// the renderer package.
package renderer
