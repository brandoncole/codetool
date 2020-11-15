package codebuild

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Analyzer(t *testing.T) {

	statistics, err := GetBuildStatistics("us-gov-west-1")
	require.NoError(t, err)
	metrics := statistics.SortedMetrics()
	t.Logf("%v", metrics)
}
