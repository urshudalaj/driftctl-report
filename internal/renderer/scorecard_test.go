package renderer

import (
	"testing"
)

func makeScorecardAnalysis(managed, unmanaged, missing, drifted int) Analysis {
	return Analysis{
		Summary: Summary{
			TotalManaged:   managed,
			TotalUnmanaged: unmanaged,
			TotalMissing:   missing,
			TotalDrifted:   drifted,
			TotalResources: managed + unmanaged + missing,
		},
	}
}

func TestBuildScorecard_NoResources(t *testing.T) {
	sc := buildScorecard(Analysis{})
	if sc.Score != 100 {
		t.Errorf("expected score 100, got %v", sc.Score)
	}
	if sc.Grade != "A" {
		t.Errorf("expected grade A, got %s", sc.Grade)
	}
}

func TestBuildScorecard_FullyManaged(t *testing.T) {
	sc := buildScorecard(makeScorecardAnalysis(10, 0, 0, 0))
	if sc.Score != 100 {
		t.Errorf("expected score 100, got %v", sc.Score)
	}
	if sc.Grade != "A" {
		t.Errorf("expected grade A, got %s", sc.Grade)
	}
}

func TestBuildScorecard_AllUnmanaged(t *testing.T) {
	sc := buildScorecard(makeScorecardAnalysis(0, 10, 0, 0))
	if sc.Score != 0 {
		t.Errorf("expected score 0, got %v", sc.Score)
	}
	if sc.Grade != "F" {
		t.Errorf("expected grade F, got %s", sc.Grade)
	}
}

func TestBuildScorecard_PartialDrift(t *testing.T) {
	// 8 managed, 2 unmanaged out of 10 total → penalty = 2/10*100 = 20 → score = 80
	sc := buildScorecard(makeScorecardAnalysis(8, 2, 0, 0))
	if sc.Score != 80 {
		t.Errorf("expected score 80, got %v", sc.Score)
	}
	if sc.Grade != "B" {
		t.Errorf("expected grade B, got %s", sc.Grade)
	}
}

func TestBuildScorecard_DriftedPenaltyIsHalf(t *testing.T) {
	// 10 total, 0 unmanaged/missing, 4 drifted → penalty = 4*0.5/10*100 = 20 → score = 80
	sc := buildScorecard(makeScorecardAnalysis(10, 0, 0, 4))
	if sc.Score != 80 {
		t.Errorf("expected score 80, got %v", sc.Score)
	}
}

func TestBuildScorecard_TotalsPopulated(t *testing.T) {
	sc := buildScorecard(makeScorecardAnalysis(5, 3, 2, 1))
	if sc.Managed != 5 || sc.Unmanaged != 3 || sc.Missing != 2 || sc.Drifted != 1 {
		t.Errorf("unexpected totals: %+v", sc)
	}
	if sc.Total != 10 {
		t.Errorf("expected total 10, got %d", sc.Total)
	}
}

func TestScorecardGrade_Boundaries(t *testing.T) {
	cases := []struct {
		score float64
		want  string
	}{
		{100, "A"}, {90, "A"}, {89, "B"}, {75, "B"}, {74, "C"},
		{60, "C"}, {59, "D"}, {40, "D"}, {39, "F"}, {0, "F"},
	}
	for _, c := range cases {
		got := scorecardGrade(c.score)
		if got != c.want {
			t.Errorf("scorecardGrade(%v) = %s, want %s", c.score, got, c.want)
		}
	}
}
