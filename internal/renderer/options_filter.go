package renderer

// WithExcludeType appends a resource type to the exclusion list.
// Resources whose Type matches any excluded type are removed before rendering.
// Calling this multiple times accumulates types; it never replaces previous values.
func WithExcludeType(resourceType string) Option {
	return func(o *Options) {
		if resourceType == "" {
			return
		}
		for _, existing := range o.ExcludeTypes {
			if existing == resourceType {
				return
			}
		}
		o.ExcludeTypes = append(o.ExcludeTypes, resourceType)
	}
}

// WithIncludeManaged sets whether fully-managed resources are shown in output.
// When set to false, only drifted/unmanaged resources are rendered.
func WithIncludeManaged(include bool) Option {
	return func(o *Options) {
		o.IncludeManaged = include
	}
}

// WithExcludeIDs appends one or more resource IDs to the exclusion list.
// Any resource whose ID appears in this list is removed from all output sections.
func WithExcludeIDs(ids ...string) Option {
	return func(o *Options) {
		for _, id := range ids {
			if id == "" {
				continue
			}
			alreadyPresent := false
			for _, existing := range o.ExcludeIDs {
				if existing == id {
					alreadyPresent = true
					break
				}
			}
			if !alreadyPresent {
				o.ExcludeIDs = append(o.ExcludeIDs, id)
			}
		}
	}
}
