package renderer

// genericResource is a minimal representation of a cloud resource used
// internally after filtering. It avoids importing the parser package and
// keeps the renderer self-contained.
type genericResource struct {
	ID   string
	Type string
}

// resourceBucket groups resources by drift category.
type resourceBucket struct {
	Managed   []genericResource
	Unmanaged []genericResource
	Missing   []genericResource
	Different []genericResource
}

// totalCount returns the sum of resources across all categories.
func (b *resourceBucket) totalCount() int {
	return len(b.Managed) + len(b.Unmanaged) + len(b.Missing) + len(b.Different)
}

// isEmpty returns true when all categories contain no resources.
func (b *resourceBucket) isEmpty() bool {
	return b.totalCount() == 0
}

// limitResources truncates the bucket so that the total resource count does
// not exceed max. A max of 0 is a no-op. Returns the number of omitted items.
func limitResources(b *resourceBucket, max int) int {
	if max <= 0 {
		return 0
	}
	total := b.totalCount()
	if total <= max {
		return 0
	}

	omitted := 0
	remaining := max

	b.Managed, remaining, omitted = truncateSlice(b.Managed, remaining, omitted)
	b.Unmanaged, remaining, omitted = truncateSlice(b.Unmanaged, remaining, omitted)
	b.Missing, remaining, omitted = truncateSlice(b.Missing, remaining, omitted)
	b.Different, _, omitted = truncateSlice(b.Different, remaining, omitted)

	return omitted
}

func truncateSlice(s []genericResource, remaining, omitted int) ([]genericResource, int, int) {
	if remaining <= 0 {
		return nil, 0, omitted + len(s)
	}
	if len(s) <= remaining {
		return s, remaining - len(s), omitted
	}
	return s[:remaining], 0, omitted + (len(s) - remaining)
}
