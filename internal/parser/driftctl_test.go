package parser_test

import (
	"strings"
	"testing"

	"github.com/driftctl-report/internal/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sampleJSON = `{
  "summary": {
    "total_resources": 5,
    "total_drifted": 2,
    "total_unmanaged": 1,
    "total_deleted": 0,
    "total_managed": 4,
    "coverage": 80.0
  },
  "managed": [
    {"id": "bucket-1", "type": "aws_s3_bucket"}
  ],
  "unmanaged": [
    {"id": "sg-abc123", "type": "aws_security_group"}
  ],
  "deleted": [],
  "differences": [
    {
      "res": {"id": "bucket-1", "type": "aws_s3_bucket"},
      "changes": [
        {"type": "update", "path": "tags.Env", "before": "prod", "after": "staging", "computed": []}
      ]
    }
  ]
}`

func TestParse_ValidJSON(t *testing.T) {
	report, err := parser.Parse(strings.NewReader(sampleJSON))
	require.NoError(t, err)
	require.NotNil(t, report)

	assert.Equal(t, 5, report.Summary.TotalResources)
	assert.Equal(t, 2, report.Summary.TotalDrifted)
	assert.InDelta(t, 80.0, report.Summary.Coverage, 0.001)

	assert.Len(t, report.Managed, 1)
	assert.Equal(t, "bucket-1", report.Managed[0].ID)

	assert.Len(t, report.Unmanaged, 1)
	assert.Equal(t, "aws_security_group", report.Unmanaged[0].Type)

	assert.Len(t, report.Differences, 1)
	assert.Equal(t, "tags.Env", report.Differences[0].Changes[0].Path)
}

func TestParse_InvalidJSON(t *testing.T) {
	_, err := parser.Parse(strings.NewReader(`{invalid}`))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "decoding driftctl JSON")
}

func TestParse_EmptyReport(t *testing.T) {
	report, err := parser.Parse(strings.NewReader(`{}`))
	require.NoError(t, err)
	assert.Equal(t, 0, report.Summary.TotalResources)
	assert.Empty(t, report.Managed)
	assert.Empty(t, report.Unmanaged)
	assert.Empty(t, report.Differences)
}
