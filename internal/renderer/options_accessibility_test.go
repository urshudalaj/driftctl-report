package renderer

import "testing"

func TestWithAccessibility_Enables(t *testing.T) {
	o := DefaultOptions()
	WithAccessibility(true)(o)
	if !o.Accessibility {
		t.Fatal("expected Accessibility to be true")
	}
}

func TestWithAccessibility_Disables(t *testing.T) {
	o := DefaultOptions()
	o.Accessibility = true
	WithAccessibility(false)(o)
	if o.Accessibility {
		t.Fatal("expected Accessibility to be false")
	}
}

func TestWithAccessibility_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.Accessibility {
		t.Fatal("expected Accessibility default to be false")
	}
}

func TestWithAccessibilitySkipNav_Enables(t *testing.T) {
	o := DefaultOptions()
	WithAccessibilitySkipNav(true)(o)
	if !o.AccessibilitySkipNav {
		t.Fatal("expected AccessibilitySkipNav to be true")
	}
}

func TestWithAccessibilitySkipNav_Disables(t *testing.T) {
	o := DefaultOptions()
	o.AccessibilitySkipNav = true
	WithAccessibilitySkipNav(false)(o)
	if o.AccessibilitySkipNav {
		t.Fatal("expected AccessibilitySkipNav to be false")
	}
}

func TestWithAccessibilitySkipNav_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.AccessibilitySkipNav {
		t.Fatal("expected AccessibilitySkipNav default to be false")
	}
}

func TestWithAccessibilityAnnounce_Enables(t *testing.T) {
	o := DefaultOptions()
	WithAccessibilityAnnounce(true)(o)
	if !o.AccessibilityAnnounce {
		t.Fatal("expected AccessibilityAnnounce to be true")
	}
}

func TestWithAccessibilityAnnounce_Disables(t *testing.T) {
	o := DefaultOptions()
	o.AccessibilityAnnounce = true
	WithAccessibilityAnnounce(false)(o)
	if o.AccessibilityAnnounce {
		t.Fatal("expected AccessibilityAnnounce to be false")
	}
}

func TestWithAccessibilityAnnounce_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.AccessibilityAnnounce {
		t.Fatal("expected AccessibilityAnnounce default to be false")
	}
}
