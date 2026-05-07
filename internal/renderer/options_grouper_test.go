package renderer

import (
	"testing"
)

func TestWithGroupByType_Enables(t *testing.T) {
	opts := DefaultOptions()
	if opts.GroupByType {
		t.Fatal("expected GroupByType to be false by default")
	}
	WithGroupByType(true)(&opts)
	if !opts.GroupByType {
		t.Fatal("expected GroupByType to be true after option applied")
	}
}

func TestWithGroupByType_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.GroupByType = true
	WithGroupByType(false)(&opts)
	if opts.GroupByType {
		t.Fatal("expected GroupByType to be false after option applied")
	}
}

func TestWithGroupCollapsed_Enables(t *testing.T) {
	opts := DefaultOptions()
	if opts.GroupCollapsed {
		t.Fatal("expected GroupCollapsed to be false by default")
	}
	WithGroupCollapsed(true)(&opts)
	if !opts.GroupCollapsed {
		t.Fatal("expected GroupCollapsed to be true after option applied")
	}
}

func TestWithGroupCollapsed_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.GroupCollapsed = true
	WithGroupCollapsed(false)(&opts)
	if opts.GroupCollapsed {
		t.Fatal("expected GroupCollapsed to be false after option applied")
	}
}

func TestWithGroupByType_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.GroupByType {
		t.Error("GroupByType should default to false")
	}
}

func TestWithGroupCollapsed_RequiresGroupByType(t *testing.T) {
	// GroupCollapsed is only meaningful when GroupByType is also enabled;
	// verify that setting collapsed alone does not implicitly enable grouping.
	opts := DefaultOptions()
	WithGroupCollapsed(true)(&opts)
	if opts.GroupByType {
		t.Error("enabling GroupCollapsed should not implicitly enable GroupByType")
	}
}
