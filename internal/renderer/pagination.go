package renderer

// Page represents a single page of resources for paginated rendering.
type Page struct {
	Number   int
	Total    int
	Size     int
	Offset   int
	HasNext  bool
	HasPrev  bool
}

// Paginator holds pagination state and slices data accordingly.
type Paginator struct {
	pageSize int
}

// NewPaginator creates a Paginator with the given page size.
// If size is less than 1 it defaults to 25.
func NewPaginator(size int) *Paginator {
	if size < 1 {
		size = 25
	}
	return &Paginator{pageSize: size}
}

// Paginate returns the Page descriptor for the requested 1-based page number
// given a total number of items.
func (p *Paginator) Paginate(page, total int) Page {
	if page < 1 {
		page = 1
	}
	totalPages := (total + p.pageSize - 1) / p.pageSize
	if totalPages < 1 {
		totalPages = 1
	}
	if page > totalPages {
		page = totalPages
	}
	offset := (page - 1) * p.pageSize
	return Page{
		Number:  page,
		Total:   totalPages,
		Size:    p.pageSize,
		Offset:  offset,
		HasNext: page < totalPages,
		HasPrev: page > 1,
	}
}

// SliceStrings returns the sub-slice of items for the given Page.
func SliceStrings(items []string, pg Page) []string {
	if pg.Offset >= len(items) {
		return []string{}
	}
	end := pg.Offset + pg.Size
	if end > len(items) {
		end = len(items)
	}
	return items[pg.Offset:end]
}
