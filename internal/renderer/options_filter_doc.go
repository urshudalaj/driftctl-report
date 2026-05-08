// Package renderer provides HTML rendering capabilities for driftctl JSON reports.
//
// # Filter Options
//
// WithExcludeType removes all resources matching the given resource type from
// the rendered output. Multiple calls accumulate excluded types.
//
//	renderer.New(report,
//		renderer.WithExcludeType("aws_s3_bucket"),
//		renderer.WithExcludeType("aws_iam_role"),
//	)
//
// WithIncludeManaged controls whether fully-managed (non-drifted) resources are
// included in the output sections. Defaults to true.
//
// WithExcludeIDs removes specific resource IDs from all output sections,
// useful for suppressing known false-positives.
package renderer
