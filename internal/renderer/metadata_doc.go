// Package renderer — metadata module
//
// # Report Metadata
//
// Every generated HTML report embeds a small metadata block that records:
//
//   - GeneratedAt  – RFC 3339 UTC timestamp of report creation.
//   - GeneratedBy  – Tool name and version (e.g. "driftctl-report 1.2.3").
//   - GoVersion    – Go runtime version used to compile the binary.
//   - Hostname     – Name of the machine that produced the report.
//   - InputFile    – Path to the source driftctl JSON file.
//
// # Building metadata
//
// Call [buildMetadata] with the resolved [Options] to obtain a
// [ReportMetadata] value ready to be passed to the HTML template.
//
// # Options
//
// Three option functions control the metadata fields that cannot be derived
// automatically from the runtime environment:
//
//   - [WithVersion]   – override the version label (default: "dev").
//   - [WithHostname]  – set the producing host (default: "unknown").
//   - [WithInputFile] – record the source JSON path.
package renderer
