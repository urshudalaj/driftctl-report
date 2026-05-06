package renderer

import "strings"

// WithSearch returns a RenderOption that sets a search query applied to all
// resource lists before rendering. The match is case-insensitive and checks
// both the resource ID and type fields.
//
// An empty query disables search filtering.
func WithSearch(query string) RenderOption {
	return func(o *Options) error {
		o.SearchQuery = strings.TrimSpace(query)
		return nil
	}
}

// WithSearchExact returns a RenderOption that sets a search query applied to
// all resource lists before rendering. Unlike WithSearch, the match is
// case-sensitive and requires an exact match on either the resource ID or type.
//
// An empty query disables search filtering.
func WithSearchExact(query string) RenderOption {
	return func(o *Options) error {
		o.SearchQuery = strings.TrimSpace(query)
		o.SearchExact = true
		return nil
	}
}
