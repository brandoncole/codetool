package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&regionFlag, "region", "r", "us-gov-west-1", "The AWS region to use")
}

var (
	regionFlag string
	rootCmd    = &cobra.Command{
		Use: "codetool",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ErrSubcommandRequired
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
