package renderer

import "testing"

func makeTrendHistory(entries ...map[string]interface{}) []map[string]interface{} {
	return entries
}

func TestBuildTrend_Disabled(t *testing.T) {
	h := makeTrendHistory(map[string]interface{}{"label": "2024-01", "managed": 5, "unmanaged": 2, "deleted": 1})
	td := buildTrend(h, false)
	if td.Enabled {
		t.Fatal("expected disabled")
	}
	if len(td.Points) != 0 {
		t.Fatalf("expected no points, got %d", len(td.Points))
	}
}

func TestBuildTrend_InsufficientHistory(t *testing.T) {
	h := makeTrendHistory(map[string]interface{}{"label": "2024-01", "managed": 10, "unmanaged": 2, "deleted": 0})
	td := buildTrend(h, true)
	if !td.Enabled {
		t.Fatal("expected enabled")
	}
	if td.Direction != TrendInsufficient {
		t.Fatalf("expected insufficient, got %s", td.Direction)
	}
	if len(td.Points) != 1 {
		t.Fatalf("expected 1 point, got %d", len(td.Points))
	}
}

func TestBuildTrend_Improving(t *testing.T) {
	h := makeTrendHistory(
		map[string]interface{}{"label": "2024-01", "managed": 5, "unmanaged": 5, "deleted": 0},
		map[string]interface{}{"label": "2024-02", "managed": 8, "unmanaged": 2, "deleted": 0},
	)
	td := buildTrend(h, true)
	if td.Direction != TrendImproving {
		t.Fatalf("expected improving, got %s", td.Direction)
	}
	if td.Delta >= 0 {
		t.Fatalf("expected negative delta, got %f", td.Delta)
	}
}

func TestBuildTrend_Degrading(t *testing.T) {
	h := makeTrendHistory(
		map[string]interface{}{"label": "2024-01", "managed": 9, "unmanaged": 1, "deleted": 0},
		map[string]interface{}{"label": "2024-02", "managed": 5, "unmanaged": 5, "deleted": 0},
	)
	td := buildTrend(h, true)
	if td.Direction != TrendDegrading {
		t.Fatalf("expected degrading, got %s", td.Direction)
	}
	if td.Delta <= 0 {
		t.Fatalf("expected positive delta, got %f", td.Delta)
	}
}

func TestBuildTrend_Stable(t *testing.T) {
	h := makeTrendHistory(
		map[string]interface{}{"label": "2024-01", "managed": 8, "unmanaged": 2, "deleted": 0},
		map[string]interface{}{"label": "2024-02", "managed": 8, "unmanaged": 2, "deleted": 0},
	)
	td := buildTrend(h, true)
	if td.Direction != TrendStable {
		t.Fatalf("expected stable, got %s", td.Direction)
	}
	if td.Delta != 0 {
		t.Fatalf("expected zero delta, got %f", td.Delta)
	}
}

func TestBuildTrend_DriftPctCalculation(t *testing.T) {
	h := makeTrendHistory(
		map[string]interface{}{"label": "2024-01", "managed": 6, "unmanaged": 3, "deleted": 1},
	)
	td := buildTrend(h, true)
	if len(td.Points) != 1 {
		t.Fatalf("expected 1 point")
	}
	// (3+1)/10 * 100 = 40.0
	if td.Points[0].DriftPct != 40.0 {
		t.Fatalf("expected 40.0, got %f", td.Points[0].DriftPct)
	}
}
