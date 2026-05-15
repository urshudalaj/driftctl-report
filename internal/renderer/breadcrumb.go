package renderer

import "github.com/snyk/driftctl-report/internal/parser"

// BreadcrumbData holds the computed breadcrumb state passed to the template.
type BreadcrumbData struct {
	Enabled   bool
	Separator string
	Items     []BreadcrumbItem
}

// buildBreadcrumb constructs a BreadcrumbData value from the current Options.
// When breadcrumbs are disabled it returns a zero-value struct with Enabled
// set to false so the template can skip rendering entirely.
func buildBreadcrumb(_ *parser.Analysis, o Options) BreadcrumbData {
	if !o.BreadcrumbEnabled {
		return BreadcrumbData{}
	}

	sep := o.BreadcrumbSeparator
	if sep == "" {
		sep = "/"
	}

	items := make([]BreadcrumbItem, len(o.BreadcrumbItems))
	copy(items, o.BreadcrumbItems)

	return BreadcrumbData{
		Enabled:   true,
		Separator: sep,
		Items:     items,
	}
}
