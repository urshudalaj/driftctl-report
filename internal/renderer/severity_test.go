package renderer

import "testing"

func TestComputeSeverity_NoIssues(t *testing.T) {
	r := computeSeverity(0, 0, 0)
	if r.Level != SeverityNone {
		t.Errorf("expected none, got %s", r.Level)
	}
	if r.Score != 0 {
		t.Errorf("expected score 0, got %d", r.Score)
	}
	if r.BadgeClass != "badge-success" {
		t.Errorf("unexpected badge class: %s", r.BadgeClass)
	}
}

func TestComputeSeverity_Low(t *testing.T) {
	// score = 0*3 + 0*2 + 4*1 = 4
	r := computeSeverity(0, 0, 4)
	if r.Level != SeverityLow {
		t.Errorf("expected low, got %s", r.Level)
	}
	if r.Score != 4 {
		t.Errorf("expected score 4, got %d", r.Score)
	}
}

func TestComputeSeverity_Medium(t *testing.T) {
	// score = 1*3 + 2*2 + 2*1 = 9
	r := computeSeverity(1, 2, 2)
	if r.Level != SeverityMedium {
		t.Errorf("expected medium, got %s", r.Level)
	}
}

func TestComputeSeverity_High(t *testing.T) {
	// score = 5*3 + 2*2 + 1*1 = 20
	r := computeSeverity(5, 2, 1)
	if r.Level != SeverityHigh {
		t.Errorf("expected high, got %s", r.Level)
	}
	if r.BadgeClass != "badge-danger" {
		t.Errorf("expected badge-danger, got %s", r.BadgeClass)
	}
}

func TestComputeSeverity_Critical(t *testing.T) {
	// score = 10*3 + 5*2 + 5*1 = 45
	r := computeSeverity(10, 5, 5)
	if r.Level != SeverityCritical {
		t.Errorf("expected critical, got %s", r.Level)
	}
	if r.Label != "Critical" {
		t.Errorf("expected label Critical, got %s", r.Label)
	}
}

func TestComputeSeverity_DescriptionNone(t *testing.T) {
	r := computeSeverity(0, 0, 0)
	if r.Description == "" {
		t.Error("description should not be empty")
	}
}

func TestSeverityBadgeClass_AllLevels(t *testing.T) {
	cases := []struct {
		level    SeverityLevel
		wantClass string
	}{
		{SeverityNone, "badge-success"},
		{SeverityLow, "badge-info"},
		{SeverityMedium, "badge-warning"},
		{SeverityHigh, "badge-danger"},
		{SeverityCritical, "badge-danger"},
	}
	for _, tc := range cases {
		got := severityBadgeClass(tc.level)
		if got != tc.wantClass {
			t.Errorf("level %s: expected %s, got %s", tc.level, tc.wantClass, got)
		}
	}
}
