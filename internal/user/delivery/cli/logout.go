package cli

import (
	"log/slog"

	"github.com/spf13/cobra"
)

func logout(_ UseCase, logger *slog.Logger) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, _ []string) {
		logger.Debug("Logout succeded!")
		// TODO
	}
}

func logoutCmd(useCase UseCase, logger *slog.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "log out from the bpla",
		Run:   logout(useCase, logger),
	}
}
