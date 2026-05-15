package renderer

import (
	"testing"
)

func TestWithIndent_Enables(t *testing.T) {
	o := DefaultOptions()
	WithIndent(true)(o)
	if !o.IndentOutput {
		t.Fatal("expected IndentOutput to be true")
	}
}

func TestWithIndent_Disables(t *testing.T) {
	o := DefaultOptions()
	o.IndentOutput = true
	WithIndent(false)(o)
	if o.IndentOutput {
		t.Fatal("expected IndentOutput to be false")
	}
}

func TestWithIndent_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.IndentOutput {
		t.Fatal("expected IndentOutput default to be false")
	}
}

func TestWithIndentSize_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithIndentSize(4)(o)
	if o.IndentSize != 4 {
		t.Fatalf("expected IndentSize 4, got %d", o.IndentSize)
	}
}

func TestWithIndentSize_ZeroIgnored(t *testing.T) {
	o := DefaultOptions()
	o.IndentSize = 2
	WithIndentSize(0)(o)
	if o.IndentSize != 2 {
		t.Fatalf("expected IndentSize to remain 2, got %d", o.IndentSize)
	}
}

func TestWithIndentSize_NegativeIgnored(t *testing.T) {
	o := DefaultOptions()
	o.IndentSize = 2
	WithIndentSize(-1)(o)
	if o.IndentSize != 2 {
		t.Fatalf("expected IndentSize to remain 2, got %d", o.IndentSize)
	}
}

func TestWithIndentChar_SpaceAccepted(t *testing.T) {
	o := DefaultOptions()
	WithIndentChar(" ")(o)
	if o.IndentChar != " " {
		t.Fatalf("expected IndentChar ' ', got %q", o.IndentChar)
	}
}

func TestWithIndentChar_TabAccepted(t *testing.T) {
	o := DefaultOptions()
	WithIndentChar("\t")(o)
	if o.IndentChar != "\t" {
		t.Fatalf("expected IndentChar tab, got %q", o.IndentChar)
	}
}

func TestWithIndentChar_InvalidIgnored(t *testing.T) {
	o := DefaultOptions()
	o.IndentChar = " "
	WithIndentChar("-")(o)
	if o.IndentChar != " " {
		t.Fatalf("expected IndentChar to remain ' ', got %q", o.IndentChar)
	}
}
