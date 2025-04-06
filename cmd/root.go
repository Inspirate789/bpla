package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Inspirate789/bpla/internal/pkg/app"
	userCLI "github.com/Inspirate789/bpla/internal/user/delivery/cli"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "bpla",
		Short: "distributed VCS with Yura's hotelki",
	}

	logger := app.NewLogger(slog.LevelDebug, true /* addSource */, true /* colorOutput */)

	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(userCLI.Commands(nil, logger)...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
