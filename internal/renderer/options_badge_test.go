package renderer

import "testing"

func TestWithBadges_Enables(t *testing.T) {
	o := DefaultOptions()
	WithBadges(true)(o)
	if !o.BadgesEnabled {
		t.Fatal("expected BadgesEnabled to be true")
	}
}

func TestWithBadges_Disables(t *testing.T) {
	o := DefaultOptions()
	o.BadgesEnabled = true
	WithBadges(false)(o)
	if o.BadgesEnabled {
		t.Fatal("expected BadgesEnabled to be false")
	}
}

func TestWithBadgeLabels_SetsKnownKeys(t *testing.T) {
	o := DefaultOptions()
	WithBadgeLabels(map[string]string{
		"managed":   "Tracked",
		"unmanaged": "Drifted",
		"deleted":   "Gone",
	})(o)
	if o.BadgeLabels["managed"] != "Tracked" {
		t.Errorf("expected managed label 'Tracked', got %q", o.BadgeLabels["managed"])
	}
	if o.BadgeLabels["unmanaged"] != "Drifted" {
		t.Errorf("expected unmanaged label 'Drifted', got %q", o.BadgeLabels["unmanaged"])
	}
	if o.BadgeLabels["deleted"] != "Gone" {
		t.Errorf("expected deleted label 'Gone', got %q", o.BadgeLabels["deleted"])
	}
}

func TestWithBadgeLabels_IgnoresUnknownKeys(t *testing.T) {
	o := DefaultOptions()
	WithBadgeLabels(map[string]string{
		"unknown": "ShouldBeIgnored",
	})(o)
	if _, ok := o.BadgeLabels["unknown"]; ok {
		t.Fatal("expected unknown key to be ignored")
	}
}

func TestWithBadgeLabels_IgnoresEmptyValues(t *testing.T) {
	o := DefaultOptions()
	WithBadgeLabels(map[string]string{
		"managed": "",
	})(o)
	if v, ok := o.BadgeLabels["managed"]; ok && v == "" {
		t.Fatal("expected empty value to be ignored")
	}
}

func TestWithBadgeShowZero_Enables(t *testing.T) {
	o := DefaultOptions()
	WithBadgeShowZero(true)(o)
	if !o.BadgeShowZero {
		t.Fatal("expected BadgeShowZero to be true")
	}
}

func TestWithBadgeShowZero_Disables(t *testing.T) {
	o := DefaultOptions()
	o.BadgeShowZero = true
	WithBadgeShowZero(false)(o)
	if o.BadgeShowZero {
		t.Fatal("expected BadgeShowZero to be false")
	}
}

func TestWithBadgeShowZero_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.BadgeShowZero {
		t.Fatal("expected default BadgeShowZero to be false")
	}
}
