package renderer

import "testing"

func makeBucket(managed, unmanaged, missing, different int) *resourceBucket {
	make_ := func(n int, prefix string) []genericResource {
		out := make([]genericResource, n)
		for i := range out {
			out[i] = genericResource{ID: prefix + string(rune('0'+i)), Type: "aws_test"}
		}
		return out
	}
	return &resourceBucket{
		Managed:   make_(managed, "m"),
		Unmanaged: make_(unmanaged, "u"),
		Missing:   make_(missing, "x"),
		Different: make_(different, "d"),
	}
}

func TestLimitResources_ZeroMax(t *testing.T) {
	b := makeBucket(5, 5, 5, 5)
	omitted := limitResources(b, 0)
	if omitted != 0 {
		t.Errorf("expected 0 omitted, got %d", omitted)
	}
	if b.totalCount() != 20 {
		t.Errorf("expected 20 total, got %d", b.totalCount())
	}
}

func TestLimitResources_UnderLimit(t *testing.T) {
	b := makeBucket(2, 2, 2, 2)
	omitted := limitResources(b, 100)
	if omitted != 0 {
		t.Errorf("expected 0 omitted, got %d", omitted)
	}
}

func TestLimitResources_ExactLimit(t *testing.T) {
	b := makeBucket(3, 3, 3, 3)
	omitted := limitResources(b, 12)
	if omitted != 0 {
		t.Errorf("expected 0 omitted, got %d", omitted)
	}
	if b.totalCount() != 12 {
		t.Errorf("expected 12, got %d", b.totalCount())
	}
}

func TestLimitResources_OverLimit(t *testing.T) {
	b := makeBucket(5, 5, 5, 5)
	omitted := limitResources(b, 10)
	if omitted != 10 {
		t.Errorf("expected 10 omitted, got %d", omitted)
	}
	if b.totalCount() != 10 {
		t.Errorf("expected 10 total, got %d", b.totalCount())
	}
}

func TestLimitResources_SmallMax(t *testing.T) {
	b := makeBucket(10, 0, 0, 0)
	omitted := limitResources(b, 3)
	if omitted != 7 {
		t.Errorf("expected 7 omitted, got %d", omitted)
	}
	if len(b.Managed) != 3 {
		t.Errorf("expected 3 managed, got %d", len(b.Managed))
	}
}

func TestResourceBucket_TotalCount(t *testing.T) {
	b := makeBucket(1, 2, 3, 4)
	if b.totalCount() != 10 {
		t.Errorf("expected 10, got %d", b.totalCount())
	}
}
