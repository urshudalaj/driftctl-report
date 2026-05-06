package renderer

// WithSearch returns a RenderOption that sets a search query applied to all
// resource lists before rendering. The match is case-insensitive and checks
// both the resource ID and type fields.
//
// An empty query disables search filtering.
func WithSearch(query string) RenderOption {
	return func(o *Options) error {
		o.SearchQuery = query
		return nil
	}
}
