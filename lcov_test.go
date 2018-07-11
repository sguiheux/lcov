package lcov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	lcovParser := New("./lcov.info")
	report, err := lcovParser.Parse()

	assert.NoError(t, err)
	assert.Equal(t, 67, report.CoveredBranches)
	assert.Equal(t, 1757, report.TotalBranches)
	assert.Equal(t, 412, report.CoveredFunctions)
	assert.Equal(t, 1695, report.TotalFunctions)
	assert.Equal(t, 2798, report.CoveredLines)
	assert.Equal(t, 6395, report.TotalLines)
}
