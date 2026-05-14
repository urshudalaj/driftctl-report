package renderer

import "testing"

func TestBuildGlossary_Disabled(t *testing.T) {
	opts := DefaultOptions()
	// GlossaryEnabled defaults to false
	got := buildGlossary(opts)
	if got.Enabled {
		t.Fatal("expected glossary to be disabled")
	}
	if len(got.Entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(got.Entries))
	}
}

func TestBuildGlossary_DefaultTerms(t *testing.T) {
	opts := DefaultOptions()
	WithGlossary(true)(&opts)
	got := buildGlossary(opts)
	if !got.Enabled {
		t.Fatal("expected glossary to be enabled")
	}
	if len(got.Entries) != len(defaultGlossaryTerms) {
		t.Fatalf("expected %d entries, got %d", len(defaultGlossaryTerms), len(got.Entries))
	}
}

func TestBuildGlossary_SortedAlphabetically(t *testing.T) {
	opts := DefaultOptions()
	WithGlossary(true)(&opts)
	got := buildGlossary(opts)
	for i := 1; i < len(got.Entries); i++ {
		if got.Entries[i].Term < got.Entries[i-1].Term {
			t.Fatalf("entries not sorted: %q before %q", got.Entries[i-1].Term, got.Entries[i].Term)
		}
	}
}

func TestBuildGlossary_CustomTermOverridesDefault(t *testing.T) {
	opts := DefaultOptions()
	WithGlossary(true)(&opts)
	WithGlossaryTerms(map[string]string{"managed": "Custom managed definition."})(&opts)
	got := buildGlossary(opts)
	for _, e := range got.Entries {
		if e.Term == "managed" && e.Definition != "Custom managed definition." {
			t.Fatalf("expected custom definition, got %q", e.Definition)
		}
	}
}

func TestBuildGlossary_MaxItemsTruncates(t *testing.T) {
	opts := DefaultOptions()
	WithGlossary(true)(&opts)
	WithGlossaryMaxItems(2)(&opts)
	got := buildGlossary(opts)
	if len(got.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(got.Entries))
	}
}

func TestBuildGlossary_MaxItemsZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	WithGlossary(true)(&opts)
	WithGlossaryMaxItems(0)(&opts)
	got := buildGlossary(opts)
	if len(got.Entries) != len(defaultGlossaryTerms) {
		t.Fatalf("expected all entries, got %d", len(got.Entries))
	}
}

func TestWithGlossaryTerms_IgnoresEmptyKeys(t *testing.T) {
	opts := DefaultOptions()
	WithGlossaryTerms(map[string]string{"": "should be ignored", "valid": "ok"})(&opts)
	if _, ok := opts.GlossaryTerms[""]; ok {
		t.Fatal("empty key should not be stored")
	}
	if opts.GlossaryTerms["valid"] != "ok" {
		t.Fatal("valid term should be stored")
	}
}

func TestWithGlossaryMaxItems_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithGlossaryMaxItems(-5)(&opts)
	if opts.GlossaryMaxItems != 0 {
		t.Fatalf("expected 0, got %d", opts.GlossaryMaxItems)
	}
}
