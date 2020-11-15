package cmd

import (
	"fmt"

	"github.com/brandoncole/codetool/pkg/codebuild"
	"github.com/spf13/cobra"
)

func init() {
	codeBuildCmd.AddCommand(codeBuildAnalyzeCmd)
}

var (
	codeBuildAnalyzeCmd = &cobra.Command{
		Use:     "analyze",
		Aliases: []string{"a"},
		RunE: func(cmd *cobra.Command, args []string) error {

			statistics, err := codebuild.GetBuildStatistics("us-gov-west-1")
			if nil != err {
				return err
			}

			metrics := statistics.SortedMetrics()
			for _, metric := range metrics {
				fmt.Printf("|%5d|%-20s|\n", int(metric.Percentage), metric.Name)
			}

			return nil

		},
	}
)
