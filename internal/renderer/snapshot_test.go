package renderer

import (
	"testing"
	"time"
)

func makeInput(managed, unmanaged, missing, drifted int, cov float64) analysisInput {
	return analysisInput{
		Managed:    managed,
		Unmanaged:  unmanaged,
		Missing:    missing,
		Drifted:    drifted,
		CoveragePC: cov,
	}
}

func TestBuildSnapshot_Disabled(t *testing.T) {
	data := buildSnapshot(makeInput(10, 2, 1, 3, 76.9), nil, false)
	if data.Enabled {
		t.Fatal("expected snapshot to be disabled")
	}
	if len(data.Entries) != 0 {
		t.Fatalf("expected no entries, got %d", len(data.Entries))
	}
}

func TestBuildSnapshot_NoHistory(t *testing.T) {
	data := buildSnapshot(makeInput(5, 1, 0, 0, 100.0), nil, true)
	if !data.Enabled {
		t.Fatal("expected snapshot to be enabled")
	}
	if len(data.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(data.Entries))
	}
	if data.HasTrend {
		t.Fatal("expected no trend with single entry")
	}
	if data.Latest.Managed != 5 {
		t.Fatalf("expected managed=5, got %d", data.Latest.Managed)
	}
}

func TestBuildSnapshot_TrendImproving(t *testing.T) {
	history := []SnapshotEntry{
		{Timestamp: time.Now().Add(-time.Hour), CoveragePC: 60.0},
	}
	data := buildSnapshot(makeInput(10, 0, 0, 0, 90.0), history, true)
	if !data.HasTrend {
		t.Fatal("expected trend to be set")
	}
	if data.Trend != "improving" {
		t.Fatalf("expected improving, got %s", data.Trend)
	}
}

func TestBuildSnapshot_TrendDegrading(t *testing.T) {
	history := []SnapshotEntry{
		{Timestamp: time.Now().Add(-time.Hour), CoveragePC: 95.0},
	}
	data := buildSnapshot(makeInput(10, 5, 3, 2, 50.0), history, true)
	if data.Trend != "degrading" {
		t.Fatalf("expected degrading, got %s", data.Trend)
	}
}

func TestBuildSnapshot_TrendStable(t *testing.T) {
	history := []SnapshotEntry{
		{Timestamp: time.Now().Add(-time.Hour), CoveragePC: 80.0},
	}
	data := buildSnapshot(makeInput(8, 2, 0, 0, 80.0), history, true)
	if data.Trend != "stable" {
		t.Fatalf("expected stable, got %s", data.Trend)
	}
}

func TestBuildSnapshot_EntriesSortedOldestFirst(t *testing.T) {
	now := time.Now()
	history := []SnapshotEntry{
		{Timestamp: now.Add(-30 * time.Minute), CoveragePC: 70.0},
		{Timestamp: now.Add(-2 * time.Hour), CoveragePC: 55.0},
	}
	data := buildSnapshot(makeInput(10, 0, 0, 0, 100.0), history, true)
	if len(data.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(data.Entries))
	}
	for i := 1; i < len(data.Entries); i++ {
		if data.Entries[i].Timestamp.Before(data.Entries[i-1].Timestamp) {
			t.Fatalf("entries not sorted oldest-first at index %d", i)
		}
	}
}
