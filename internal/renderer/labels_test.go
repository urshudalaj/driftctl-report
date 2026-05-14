package renderer

import "testing"

func makeLabelIDs() []string {
	return []string{"aws_s3_bucket.a", "aws_s3_bucket.b", "aws_iam_role.c"}
}

func TestBuildLabels_Disabled(t *testing.T) {
	o := DefaultOptions()
	data := buildLabels(makeLabelIDs(), o)
	if data.Enabled {
		t.Fatal("expected labels disabled by default")
	}
	if len(data.Entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(data.Entries))
	}
}

func TestBuildLabels_EmptyMap(t *testing.T) {
	o := DefaultOptions()
	WithLabels(true)(&o)
	data := buildLabels(makeLabelIDs(), o)
	if data.Enabled {
		t.Fatal("enabled should be false when label map is empty")
	}
}

func TestBuildLabels_AllIDs(t *testing.T) {
	o := DefaultOptions()
	WithLabels(true)(&o)
	WithLabelMap(map[string]map[string]string{
		"aws_s3_bucket.a": {"env": "prod", "team": "infra"},
		"aws_iam_role.c":  {"env": "staging"},
	})(&o)

	data := buildLabels(makeLabelIDs(), o)
	if !data.Enabled {
		t.Fatal("expected labels enabled")
	}
	if len(data.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(data.Entries))
	}
}

func TestBuildLabels_UnknownIDsIgnored(t *testing.T) {
	o := DefaultOptions()
	WithLabels(true)(&o)
	WithLabelMap(map[string]map[string]string{
		"aws_s3_bucket.UNKNOWN": {"env": "prod"},
	})(&o)

	data := buildLabels(makeLabelIDs(), o)
	if len(data.Entries) != 0 {
		t.Fatalf("expected 0 entries for unknown IDs, got %d", len(data.Entries))
	}
}

func TestBuildLabels_KeyFilterApplied(t *testing.T) {
	o := DefaultOptions()
	WithLabels(true)(&o)
	WithLabelMap(map[string]map[string]string{
		"aws_s3_bucket.a": {"env": "prod", "team": "infra", "owner": "alice"},
	})(&o)
	WithLabelKeys([]string{"env", "owner"})(&o)

	data := buildLabels([]string{"aws_s3_bucket.a"}, o)
	if len(data.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(data.Entries))
	}
	if len(data.Entries[0].Pairs) != 2 {
		t.Fatalf("expected 2 pairs after key filter, got %d", len(data.Entries[0].Pairs))
	}
	for _, p := range data.Entries[0].Pairs {
		if p.Key != "env" && p.Key != "owner" {
			t.Errorf("unexpected key %q passed filter", p.Key)
		}
	}
}

func TestBuildLabels_PairsSortedByKey(t *testing.T) {
	o := DefaultOptions()
	WithLabels(true)(&o)
	WithLabelMap(map[string]map[string]string{
		"aws_s3_bucket.a": {"zzz": "last", "aaa": "first", "mmm": "mid"},
	})(&o)

	data := buildLabels([]string{"aws_s3_bucket.a"}, o)
	pairs := data.Entries[0].Pairs
	for i := 1; i < len(pairs); i++ {
		if pairs[i].Key < pairs[i-1].Key {
			t.Errorf("pairs not sorted: %q before %q", pairs[i-1].Key, pairs[i].Key)
		}
	}
}

func TestWithLabelKeys_EmptyClears(t *testing.T) {
	o := DefaultOptions()
	WithLabelKeys([]string{"env"})(&o)
	if len(o.LabelKeys) != 1 {
		t.Fatalf("expected 1 key, got %d", len(o.LabelKeys))
	}
	WithLabelKeys([]string{})(&o)
	if o.LabelKeys != nil {
		t.Fatalf("expected nil after clearing, got %v", o.LabelKeys)
	}
}
