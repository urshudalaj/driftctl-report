package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"
)

func makeRecsAnalysis(unmanaged, deleted, drifted int) Analysis {
	a := Analysis{}
	for i := 0; i < unmanaged; i++ {
		a.Summary.TotalUnmanaged = append(a.Summary.TotalUnmanaged, analyser.Resource{ResourceID: itoa(i), ResourceType: "aws_s3_bucket"})
	}
	for i := 0; i < deleted; i++ {
		a.Summary.TotalDeleted = append(a.Summary.TotalDeleted, analyser.Resource{ResourceID: itoa(i), ResourceType: "aws_instance"})
	}
	for i := 0; i < drifted; i++ {
		a.Summary.TotalDrifted = append(a.Summary.TotalDrifted, analyser.Resource{ResourceID: itoa(i), ResourceType: "aws_iam_role"})
	}
	return a
}

func TestBuildRecommendations_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.Recommendations = false
	got := buildRecommendations(makeRecsAnalysis(5, 2, 3), opts)
	if got.Enabled {
		t.Fatal("expected disabled")
	}
	if len(got.Recommendations) != 0 {
		t.Fatalf("expected empty, got %d", len(got.Recommendations))
	}
}

func TestBuildRecommendations_NoIssues(t *testing.T) {
	opts := DefaultOptions()
	opts.Recommendations = true
	got := buildRecommendations(makeRecsAnalysis(0, 0, 0), opts)
	if !got.Enabled {
		t.Fatal("expected enabled")
	}
	if len(got.Recommendations) != 0 {
		t.Fatalf("expected 0 recommendations, got %d", len(got.Recommendations))
	}
}

func TestBuildRecommendations_AllKinds(t *testing.T) {
	opts := DefaultOptions()
	opts.Recommendations = true
	got := buildRecommendations(makeRecsAnalysis(10, 3, 7), opts)
	if len(got.Recommendations) != 3 {
		t.Fatalf("expected 3 recommendations, got %d", len(got.Recommendations))
	}
}

func TestBuildRecommendations_SortedByCountDesc(t *testing.T) {
	opts := DefaultOptions()
	opts.Recommendations = true
	got := buildRecommendations(makeRecsAnalysis(2, 15, 8), opts)
	for i := 1; i < len(got.Recommendations); i++ {
		if got.Recommendations[i].Count > got.Recommendations[i-1].Count {
			t.Errorf("recommendations not sorted by count desc at index %d", i)
		}
	}
}

func TestBuildRecommendations_UnmanagedBadgeClass(t *testing.T) {
	opts := DefaultOptions()
	opts.Recommendations = true
	got := buildRecommendations(makeRecsAnalysis(25, 0, 0), opts)
	if len(got.Recommendations) != 1 {
		t.Fatalf("expected 1 recommendation")
	}
	if got.Recommendations[0].BadgeClass == "" {
		t.Error("expected non-empty badge class")
	}
}

func TestBuildRecommendations_TitlesNonEmpty(t *testing.T) {
	opts := DefaultOptions()
	opts.Recommendations = true
	got := buildRecommendations(makeRecsAnalysis(1, 1, 1), opts)
	for _, r := range got.Recommendations {
		if r.Title == "" {
			t.Error("expected non-empty title")
		}
		if r.Description == "" {
			t.Error("expected non-empty description")
		}
	}
}
