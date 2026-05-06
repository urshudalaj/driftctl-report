package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"
	"github.com/snyk/driftctl/pkg/resource"
	"github.com/stretchr/testify/assert"
)

func makeAnalysis(missing, unmanaged []resource.Resource, diffs []analyser.Difference) *analyser.Analysis {
	a := &analyser.Analysis{}
	for _, r := range missing {
		a.AddDeleted(r)
	}
	for _, r := range unmanaged {
		a.AddUnmanaged(r)
	}
	for _, d := range diffs {
		a.AddDifference(d)
	}
	return a
}

func TestBuildChangelog_Empty(t *testing.T) {
	a := &analyser.Analysis{}
	sections := buildChangelog(a)
	assert.Empty(t, sections)
}

func TestBuildChangelog_OnlyMissing(t *testing.T) {
	a := makeAnalysis(
		[]resource.Resource{
			&resource.FakeResource{Id: "sg-1", Type: "aws_security_group"},
			&resource.FakeResource{Id: "sg-2", Type: "aws_security_group"},
		},
		nil, nil,
	)
	sections := buildChangelog(a)
	assert.Len(t, sections, 1)
	assert.Equal(t, "missing", sections[0].Kind)
	assert.Len(t, sections[0].Entries, 2)
}

func TestBuildChangelog_SortedByTypeAndID(t *testing.T) {
	a := makeAnalysis(
		[]resource.Resource{
			&resource.FakeResource{Id: "z-id", Type: "aws_vpc"},
			&resource.FakeResource{Id: "a-id", Type: "aws_vpc"},
		},
		nil, nil,
	)
	sections := buildChangelog(a)
	entries := sections[0].Entries
	assert.Equal(t, "a-id", entries[0].ResourceID)
	assert.Equal(t, "z-id", entries[1].ResourceID)
}

func TestBuildChangelog_AllKinds(t *testing.T) {
	a := makeAnalysis(
		[]resource.Resource{&resource.FakeResource{Id: "m1", Type: "aws_s3_bucket"}},
		[]resource.Resource{&resource.FakeResource{Id: "u1", Type: "aws_iam_role"}},
		[]analyser.Difference{
			{Res: &resource.FakeResource{Id: "c1", Type: "aws_instance"}, Changelog: make(analyser.Changelog, 3)},
		},
	)
	sections := buildChangelog(a)
	assert.Len(t, sections, 3)
	kinds := []string{sections[0].Kind, sections[1].Kind, sections[2].Kind}
	assert.Contains(t, kinds, "missing")
	assert.Contains(t, kinds, "unmanaged")
	assert.Contains(t, kinds, "changed")
}

func TestBuildChangelog_ChangedHasDiffCount(t *testing.T) {
	a := makeAnalysis(nil, nil, []analyser.Difference{
		{Res: &resource.FakeResource{Id: "i-1", Type: "aws_instance"}, Changelog: make(analyser.Changelog, 5)},
	})
	sections := buildChangelog(a)
	assert.Equal(t, 5, sections[0].Entries[0].DiffCount)
}
