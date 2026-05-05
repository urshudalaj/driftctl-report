// Package renderer converts a parsed driftctl report into an HTML document.
//
// # Usage
//
//	r, err := renderer.New(report)
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err := r.Render(os.Stdout); err != nil {
//		log.Fatal(err)
//	}
//
// # Filtering
//
// Before rendering, callers may supply [FilterOptions] to limit the output to
// specific resource types or to show only drifted resources:
//
//	r.SetFilters(renderer.FilterOptions{
//		OnlyDrifted:   true,
//		ResourceTypes: []string{"aws_s3_bucket"},
//	})
//
// The filter is applied at render time and does not mutate the underlying
// report stored in the renderer.
package renderer
