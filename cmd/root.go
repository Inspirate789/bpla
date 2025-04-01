package cmd

import (
	"fmt"
	"os"

	userCLI "github.com/Inspirate789/bpla/internal/user/delivery/cli"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "bpla",
		Short: "distributed VCS with Yura's hotelki",
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(userCLI.Commands()...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
