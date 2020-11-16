package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/brandoncole/codetool/pkg/codebuild"
	"github.com/spf13/cobra"
)

func init() {
	codeBuildCmd.AddCommand(codeBuildAnalyzeCmd)
}

var (
	codeBuildAnalyzeCmd = &cobra.Command{
		Use: "analyze",
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.PersistentFlags()

			statistics, err := codebuild.GetBuildStatistics(regionFlag)
			if nil != err {
				return err
			}

			metrics := statistics.SortedMetrics()

			w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
			fmt.Fprintln(w, "Time\tPercent\tPhase")

			totalTime := time.Duration(statistics.BuildDurationInSeconds) * time.Second
			fmt.Fprintf(w, "%v\t%d\tTOTAL TIME\n", totalTime, 100)

			for _, metric := range metrics {
				t := time.Duration(metric.Value) * time.Second
				fmt.Fprintf(w, "%v\t%d\t%s\n", t, int(metric.Percentage), metric.Name)
			}
			w.Flush()

			fmt.Printf("Total Builds: %d\n", statistics.BuildsCount)

			return nil

		},
	}
)
