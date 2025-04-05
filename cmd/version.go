package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bpla",
	Long:  `All software has versions. This is bpla's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BPLA v0.1.0")
	},
}
