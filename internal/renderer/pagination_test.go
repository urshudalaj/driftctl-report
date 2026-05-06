package renderer

import (
	"testing"
)

func TestNewPaginator_DefaultSize(t *testing.T) {
	p := NewPaginator(0)
	if p.pageSize != 25 {
		t.Fatalf("expected default page size 25, got %d", p.pageSize)
	}
}

func TestNewPaginator_CustomSize(t *testing.T) {
	p := NewPaginator(10)
	if p.pageSize != 10 {
		t.Fatalf("expected page size 10, got %d", p.pageSize)
	}
}

func TestPaginate_FirstPage(t *testing.T) {
	p := NewPaginator(10)
	pg := p.Paginate(1, 25)
	if pg.Number != 1 || pg.Total != 3 || pg.Offset != 0 {
		t.Fatalf("unexpected page: %+v", pg)
	}
	if !pg.HasNext || pg.HasPrev {
		t.Fatal("expected HasNext=true HasPrev=false")
	}
}

func TestPaginate_LastPage(t *testing.T) {
	p := NewPaginator(10)
	pg := p.Paginate(3, 25)
	if pg.Number != 3 || pg.Total != 3 {
		t.Fatalf("unexpected page: %+v", pg)
	}
	if pg.HasNext || !pg.HasPrev {
		t.Fatal("expected HasNext=false HasPrev=true")
	}
}

func TestPaginate_BeyondTotal(t *testing.T) {
	p := NewPaginator(10)
	pg := p.Paginate(99, 5)
	if pg.Number != 1 {
		t.Fatalf("expected clamped page 1, got %d", pg.Number)
	}
}

func TestPaginate_ZeroItems(t *testing.T) {
	p := NewPaginator(10)
	pg := p.Paginate(1, 0)
	if pg.Total != 1 || pg.Offset != 0 {
		t.Fatalf("unexpected page for zero items: %+v", pg)
	}
}

func TestSliceStrings_Normal(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}
	p := NewPaginator(2)
	pg := p.Paginate(2, len(items))
	got := SliceStrings(items, pg)
	if len(got) != 2 || got[0] != "c" || got[1] != "d" {
		t.Fatalf("unexpected slice: %v", got)
	}
}

func TestSliceStrings_OffsetBeyondLength(t *testing.T) {
	items := []string{"a"}
	p := NewPaginator(5)
	pg := Page{Offset: 10, Size: 5}
	got := SliceStrings(items, pg)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}
