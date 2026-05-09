package renderer

// SortField identifies the field used for sorting resources.
type SortField string

const (
	// SortByType sorts resources by their resource type.
	SortByType SortField = "type"
	// SortByID sorts resources by their resource ID.
	SortByID SortField = "id"
)

// SortOrder identifies the direction of sorting.
type SortOrder string

const (
	// SortAsc sorts in ascending order.
	SortAsc SortOrder = "asc"
	// SortDesc sorts in descending order.
	SortDesc SortOrder = "desc"
)

// WithSortField sets the field used to sort resource lists.
// Accepted values are SortByType and SortByID.
// Unknown values are silently ignored.
func WithSortField(field SortField) Option {
	return func(o *Options) {
		switch field {
		case SortByType, SortByID:
			o.SortField = field
		}
	}
}

// WithSortOrder sets the sort direction for resource lists.
// Accepted values are SortAsc and SortDesc.
// Unknown values are silently ignored.
func WithSortOrder(order SortOrder) Option {
	return func(o *Options) {
		switch order {
		case SortAsc, SortDesc:
			o.SortOrder = order
		}
	}
}
