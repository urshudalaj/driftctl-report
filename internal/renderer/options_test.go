package renderer

import "testing"

func TestDefaultOptions(t *testing.T) {
	o := DefaultOptions()
	if o.Title != "Drift Report" {
		t.Errorf("unexpected default title: %s", o.Title)
	}
	if o.OnlyDrifted {
		t.Error("expected OnlyDrifted to be false by default")
	}
	if o.MaxResources != 0 {
		t.Errorf("expected MaxResources 0, got %d", o.MaxResources)
	}
	if o.SortOrder != SortByTypeAsc {
		t.Errorf("expected SortByTypeAsc default, got %v", o.SortOrder)
	}
}

func TestWithTitle(t *testing.T) {
	o := applyOptions(DefaultOptions(), []Option{WithTitle("My Report")})
	if o.Title != "My Report" {
		t.Errorf("expected 'My Report', got %s", o.Title)
	}
}

func TestWithTitle_Empty(t *testing.T) {
	o := applyOptions(DefaultOptions(), []Option{WithTitle("")})
	if o.Title != "Drift Report" {
		t.Errorf("empty title should not override default, got %s", o.Title)
	}
}

func TestWithOnlyDrifted(t *testing.T) {
	o := applyOptions(DefaultOptions(), []Option{WithOnlyDrifted(true)})
	if !o.OnlyDrifted {
		t.Error("expected OnlyDrifted true")
	}
}

func TestWithResourceType(t *testing.T) {
	o := applyOptions(DefaultOptions(), []Option{WithResourceType("aws_s3_bucket")})
	if o.ResourceType != "aws_s3_bucket" {
		t.Errorf("expected aws_s3_bucket, got %s", o.ResourceType)
	}
}

func TestWithMaxResources(t *testing.T) {
	o := applyOptions(DefaultOptions(), []Option{WithMaxResources(50)})
	if o.MaxResources != 50 {
		t.Errorf("expected 50, got %d", o.MaxResources)
	}
}

func TestWithMaxResources_Negative(t *testing.T) {
	o := applyOptions(DefaultOptions(), []Option{WithMaxResources(-1)})
	if o.MaxResources != 0 {
		t.Errorf("negative value should not override default, got %d", o.MaxResources)
	}
}

func TestWithSortOrder(t *testing.T) {
	for _, tc := range []SortOrder{SortByTypeAsc, SortByTypeDesc, SortByIDAsc, SortByIDDesc} {
		o := applyOptions(DefaultOptions(), []Option{WithSortOrder(tc)})
		if o.SortOrder != tc {
			t.Errorf("expected sort order %v, got %v", tc, o.SortOrder)
		}
	}
}
