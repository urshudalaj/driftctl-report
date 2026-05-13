package renderer

import (
	"testing"
)

func TestWithCSV_Enables(t *testing.T) {
	o := DefaultOptions()
	WithCSV(true)(o)
	if !o.ExportCSV {
		t.Fatal("expected ExportCSV to be true")
	}
}

func TestWithCSV_Disables(t *testing.T) {
	o := DefaultOptions()
	o.ExportCSV = true
	WithCSV(false)(o)
	if o.ExportCSV {
		t.Fatal("expected ExportCSV to be false")
	}
}

func TestWithCSV_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.ExportCSV {
		t.Fatal("expected ExportCSV default to be false")
	}
}

func TestWithCSVFilename_SetsName(t *testing.T) {
	o := DefaultOptions()
	WithCSVFilename("report.csv")(o)
	if o.CSVFilename != "report.csv" {
		t.Fatalf("expected report.csv, got %s", o.CSVFilename)
	}
}

func TestWithCSVFilename_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	o.CSVFilename = "original.csv"
	WithCSVFilename("")(o)
	if o.CSVFilename != "original.csv" {
		t.Fatalf("expected original.csv, got %s", o.CSVFilename)
	}
}

func TestWithCSVDelimiter_Sets(t *testing.T) {
	o := DefaultOptions()
	WithCSVDelimiter(';')(o)
	if o.CSVDelimiter != ';' {
		t.Fatalf("expected semicolon delimiter, got %q", o.CSVDelimiter)
	}
}

func TestWithCSVDelimiter_ZeroIgnored(t *testing.T) {
	o := DefaultOptions()
	o.CSVDelimiter = ','
	WithCSVDelimiter(0)(o)
	if o.CSVDelimiter != ',' {
		t.Fatalf("expected comma delimiter to remain, got %q", o.CSVDelimiter)
	}
}

func TestWithCSVDelimiter_DefaultIsComma(t *testing.T) {
	o := DefaultOptions()
	if o.CSVDelimiter != ',' {
		t.Fatalf("expected default delimiter to be comma, got %q", o.CSVDelimiter)
	}
}
