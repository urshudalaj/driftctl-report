package renderer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/snyk/driftctl-report/internal/parser"
)

func makeCsvAnalysis() *parser.Analysis {
	return &parser.Analysis{
		Summary: parser.Summary{
			TotalManaged:   []parser.Resource{{Type: "aws_s3_bucket", ID: "my-bucket"}},
			TotalUnmanaged: []parser.Resource{{Type: "aws_iam_role", ID: "orphan-role"}},
			TotalDeleted:   []parser.Resource{{Type: "aws_lambda_function", ID: "old-fn"}},
		},
	}
}

func TestBuildCSVRows_AllStatuses(t *testing.T) {
	a := makeCsvAnalysis()
	rows := buildCSVRows(a)
	if len(rows) != 3 {
		t.Fatalf("expected 3 rows, got %d", len(rows))
	}
}

func TestBuildCSVRows_StatusValues(t *testing.T) {
	a := makeCsvAnalysis()
	rows := buildCSVRows(a)
	statuses := map[string]bool{}
	for _, r := range rows {
		statuses[r.Status] = true
	}
	for _, want := range []string{"managed", "unmanaged", "deleted"} {
		if !statuses[want] {
			t.Errorf("missing status %q in rows", want)
		}
	}
}

func TestBuildCSVRows_Empty(t *testing.T) {
	a := &parser.Analysis{}
	rows := buildCSVRows(a)
	if len(rows) != 0 {
		t.Fatalf("expected 0 rows, got %d", len(rows))
	}
}

func TestWriteCSV_HasHeader(t *testing.T) {
	var buf bytes.Buffer
	rows := []CSVRow{{ResourceType: "aws_s3_bucket", ResourceID: "b1", Status: "managed"}}
	if err := WriteCSV(&buf, rows, ','); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if lines[0] != "resource_type,resource_id,status" {
		t.Fatalf("unexpected header: %s", lines[0])
	}
}

func TestWriteCSV_RowCount(t *testing.T) {
	var buf bytes.Buffer
	rows := []CSVRow{
		{ResourceType: "aws_s3_bucket", ResourceID: "b1", Status: "managed"},
		{ResourceType: "aws_iam_role", ResourceID: "r1", Status: "unmanaged"},
	}
	if err := WriteCSV(&buf, rows, ','); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	// header + 2 data rows
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
}

func TestWriteCSV_CustomDelimiter(t *testing.T) {
	var buf bytes.Buffer
	rows := []CSVRow{{ResourceType: "aws_s3_bucket", ResourceID: "b1", Status: "managed"}}
	if err := WriteCSV(&buf, rows, ';'); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), ";") {
		t.Fatal("expected semicolon delimiter in output")
	}
}

func TestWriteCSV_EmptyRows(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteCSV(&buf, nil, ','); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 1 {
		t.Fatalf("expected only header line, got %d lines", len(lines))
	}
}
