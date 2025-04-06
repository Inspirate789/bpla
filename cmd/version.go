package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//nolint:forbidigo // fmt.Println is needed here
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of bpla",
		Long:  `All software has versions. This is bpla's.`,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("BPLA v0.1.0")
		},
	}
}
