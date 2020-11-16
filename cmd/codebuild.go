package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(codeBuildCmd)
}

var (
	codeBuildCmd = &cobra.Command{
		Use:     "codebuild",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ErrSubcommandRequired
		},
	}
)
