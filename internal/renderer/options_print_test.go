package renderer

import "testing"

func TestWithPrint_Enables(t *testing.T) {
	o := DefaultOptions()
	WithPrint(true)(o)
	if !o.PrintFriendly {
		t.Fatal("expected PrintFriendly to be true")
	}
}

func TestWithPrint_Disables(t *testing.T) {
	o := DefaultOptions()
	o.PrintFriendly = true
	WithPrint(false)(o)
	if o.PrintFriendly {
		t.Fatal("expected PrintFriendly to be false")
	}
}

func TestWithPrint_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.PrintFriendly {
		t.Fatal("expected PrintFriendly default to be false")
	}
}

func TestWithPrintHideNav_Enables(t *testing.T) {
	o := DefaultOptions()
	WithPrintHideNav(true)(o)
	if !o.PrintHideNav {
		t.Fatal("expected PrintHideNav to be true")
	}
}

func TestWithPrintHideNav_Disables(t *testing.T) {
	o := DefaultOptions()
	o.PrintHideNav = true
	WithPrintHideNav(false)(o)
	if o.PrintHideNav {
		t.Fatal("expected PrintHideNav to be false")
	}
}

func TestWithPrintPageBreaks_Enables(t *testing.T) {
	o := DefaultOptions()
	WithPrintPageBreaks(true)(o)
	if !o.PrintPageBreaks {
		t.Fatal("expected PrintPageBreaks to be true")
	}
}

func TestWithPrintPageBreaks_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.PrintPageBreaks {
		t.Fatal("expected PrintPageBreaks default to be false")
	}
}
