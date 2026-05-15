// Package renderer — breadcrumb module
//
// # Breadcrumb Navigation
//
// The breadcrumb feature renders a navigational trail at the top of the
// generated HTML report, helping readers understand where the current report
// sits within a larger hierarchy (e.g. Home > Reports > Drift).
//
// # Usage
//
//	exporter := renderer.NewExporter(analysis,
//		renderer.WithBreadcrumb(true),
//		renderer.WithBreadcrumbSeparator("›"),
//		renderer.WithBreadcrumbItems(
//			renderer.BreadcrumbItem{Label: "Home",    URL: "/"},
//			renderer.BreadcrumbItem{Label: "Reports", URL: "/reports"},
//			renderer.BreadcrumbItem{Label: "Drift"},
//		),
//	)
//
// # Behaviour
//
//   - Breadcrumbs are disabled by default.
//   - Items with an empty Label are silently ignored.
//   - Items without a URL are rendered as plain text (no anchor tag).
//   - The separator defaults to "/" when not explicitly set.
package renderer
