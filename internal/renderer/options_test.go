package renderer

import (
	"testing"
)

func TestDefaultOptions(t *testing.T) {
	o := DefaultOptions()
	if o.Title != "Drift Report" {
		t.Errorf("expected title %q, got %q", "Drift Report", o.Title)
	}
	if o.OnlyDrifted {
		t.Error("expected OnlyDrifted to be false")
	}
	if o.ResourceType != "" {
		t.Errorf("expected empty ResourceType, got %q", o.ResourceType)
	}
	if o.MaxResources != 0 {
		t.Errorf("expected MaxResources 0, got %d", o.MaxResources)
	}
}

func TestWithTitle(t *testing.T) {
	o := DefaultOptions()
	WithTitle("My Report")(&o)
	if o.Title != "My Report" {
		t.Errorf("expected %q, got %q", "My Report", o.Title)
	}
}

func TestWithOnlyDrifted(t *testing.T) {
	o := DefaultOptions()
	WithOnlyDrifted()(&o)
	if !o.OnlyDrifted {
		t.Error("expected OnlyDrifted to be true")
	}
}

func TestWithResourceType(t *testing.T) {
	o := DefaultOptions()
	WithResourceType("aws_s3_bucket")(&o)
	if o.ResourceType != "aws_s3_bucket" {
		t.Errorf("expected %q, got %q", "aws_s3_bucket", o.ResourceType)
	}
}

func TestWithMaxResources(t *testing.T) {
	o := DefaultOptions()
	WithMaxResources(50)(&o)
	if o.MaxResources != 50 {
		t.Errorf("expected 50, got %d", o.MaxResources)
	}
}

func TestOptions_Apply(t *testing.T) {
	o := DefaultOptions()
	o.apply([]Option{
		WithTitle("Combined"),
		WithOnlyDrifted(),
		WithMaxResources(10),
	})
	if o.Title != "Combined" {
		t.Errorf("title: expected %q, got %q", "Combined", o.Title)
	}
	if !o.OnlyDrifted {
		t.Error("expected OnlyDrifted true")
	}
	if o.MaxResources != 10 {
		t.Errorf("MaxResources: expected 10, got %d", o.MaxResources)
	}
}
