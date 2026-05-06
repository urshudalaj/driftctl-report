package renderer

import (
	"testing"
)

func TestBuildBadges_AllZero(t *testing.T) {
	badges := buildBadges(0, 0, 0, 0)
	if len(badges) != 4 {
		t.Fatalf("expected 4 badges, got %d", len(badges))
	}
	for _, b := range badges {
		if b.Level != BadgeLevelSuccess {
			t.Errorf("badge %q: expected success level when count is 0, got %s", b.Label, b.Level)
		}
	}
}

func TestBuildBadges_ManagedAlwaysSuccess(t *testing.T) {
	badges := buildBadges(10, 0, 0, 0)
	if badges[0].Level != BadgeLevelSuccess {
		t.Errorf("managed badge should always be success, got %s", badges[0].Level)
	}
	if badges[0].Value != "10" {
		t.Errorf("expected value '10', got %s", badges[0].Value)
	}
}

func TestBuildBadges_SmallCountIsWarning(t *testing.T) {
	badges := buildBadges(5, 3, 2, 1)
	// unmanaged=3 → warning
	if badges[1].Level != BadgeLevelWarning {
		t.Errorf("unmanaged=3: expected warning, got %s", badges[1].Level)
	}
}

func TestBuildBadges_LargeCountIsDanger(t *testing.T) {
	badges := buildBadges(0, 10, 0, 0)
	if badges[1].Level != BadgeLevelDanger {
		t.Errorf("unmanaged=10: expected danger, got %s", badges[1].Level)
	}
}

func TestBuildBadges_Labels(t *testing.T) {
	want := []string{"Managed", "Unmanaged", "Missing", "Drifted"}
	badges := buildBadges(1, 1, 1, 1)
	for i, w := range want {
		if badges[i].Label != w {
			t.Errorf("badge[%d]: expected label %q, got %q", i, w, badges[i].Label)
		}
	}
}

func TestItoa(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{42, "42"},
		{-7, "-7"},
		{1000, "1000"},
	}
	for _, tc := range cases {
		got := itoa(tc.in)
		if got != tc.want {
			t.Errorf("itoa(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestBadgeLevelForCount(t *testing.T) {
	if badgeLevelForCount(0) != BadgeLevelSuccess {
		t.Error("0 should be success")
	}
	if badgeLevelForCount(4) != BadgeLevelWarning {
		t.Error("4 should be warning")
	}
	if badgeLevelForCount(5) != BadgeLevelDanger {
		t.Error("5 should be danger")
	}
}
