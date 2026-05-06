package renderer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithChangelog_EnablesSection(t *testing.T) {
	o := DefaultOptions()
	err := WithChangelog(true)(o)
	require.NoError(t, err)
	assert.True(t, o.ShowChangelog)
}

func TestWithChangelog_DisablesSection(t *testing.T) {
	o := DefaultOptions()
	o.ShowChangelog = true
	err := WithChangelog(false)(o)
	require.NoError(t, err)
	assert.False(t, o.ShowChangelog)
}

func TestWithChangelogLimit_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	err := WithChangelogLimit(50)(o)
	require.NoError(t, err)
	assert.Equal(t, 50, o.ChangelogLimit)
}

func TestWithChangelogLimit_ZeroMeansUnlimited(t *testing.T) {
	o := DefaultOptions()
	err := WithChangelogLimit(0)(o)
	require.NoError(t, err)
	assert.Equal(t, 0, o.ChangelogLimit)
}

func TestWithChangelogLimit_NegativeClampedToZero(t *testing.T) {
	o := DefaultOptions()
	err := WithChangelogLimit(-10)(o)
	require.NoError(t, err)
	assert.Equal(t, 0, o.ChangelogLimit)
}

func TestDefaultOptions_ChangelogDefaults(t *testing.T) {
	o := DefaultOptions()
	assert.True(t, o.ShowChangelog, "changelog should be enabled by default")
	assert.Equal(t, 0, o.ChangelogLimit, "changelog limit should be unlimited by default")
}
