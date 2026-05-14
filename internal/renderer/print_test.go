package renderer

import (
	"strings"
	"testing"
)

func TestBuildPrintStyles_Disabled(t *testing.T) {
	o := DefaultOptions()
	ps := buildPrintStyles(o)
	if ps.Enabled {
		t.Fatal("expected Enabled to be false when PrintFriendly is off")
	}
	if ps.CSS != "" {
		t.Fatalf("expected empty CSS, got %q", ps.CSS)
	}
}

func TestBuildPrintStyles_Enabled_ContainsBaseRules(t *testing.T) {
	o := DefaultOptions()
	WithPrint(true)(o)
	ps := buildPrintStyles(o)
	if !ps.Enabled {
		t.Fatal("expected Enabled to be true")
	}
	if !strings.Contains(ps.CSS, "@media print") {
		t.Error("expected CSS to contain @media print block")
	}
	if !strings.Contains(ps.CSS, "font-size: 11pt") {
		t.Error("expected CSS to contain base body font-size rule")
	}
}

func TestBuildPrintStyles_HideNav(t *testing.T) {
	o := DefaultOptions()
	WithPrint(true)(o)
	WithPrintHideNav(true)(o)
	ps := buildPrintStyles(o)
	if !strings.Contains(ps.CSS, "navbar") {
		t.Error("expected CSS to contain navbar hide rule")
	}
}

func TestBuildPrintStyles_HideNav_False(t *testing.T) {
	o := DefaultOptions()
	WithPrint(true)(o)
	WithPrintHideNav(false)(o)
	ps := buildPrintStyles(o)
	if strings.Contains(ps.CSS, "navbar") {
		t.Error("expected CSS to NOT contain navbar hide rule when HideNav is false")
	}
}

func TestBuildPrintStyles_PageBreaks(t *testing.T) {
	o := DefaultOptions()
	WithPrint(true)(o)
	WithPrintPageBreaks(true)(o)
	ps := buildPrintStyles(o)
	if !ps.PageBreaks {
		t.Fatal("expected PageBreaks to be true")
	}
	if !strings.Contains(ps.CSS, "page-break-before") {
		t.Error("expected CSS to contain page-break-before rule")
	}
}

func TestBuildPrintStyles_NoPageBreaks(t *testing.T) {
	o := DefaultOptions()
	WithPrint(true)(o)
	WithPrintPageBreaks(false)(o)
	ps := buildPrintStyles(o)
	if strings.Contains(ps.CSS, "page-break-before") {
		t.Error("expected CSS to NOT contain page-break-before when disabled")
	}
}
