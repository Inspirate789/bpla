package cli

import (
	"log/slog"

	"github.com/spf13/cobra"
)

func login(_ UseCase, logger *slog.Logger) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, _ []string) {
		username, _ := cmd.Flags().GetString("username")
		// password, _ := cmd.Flags().GetString("password")
		logger.Debug("Login succeded!", slog.String("username", username))
		// TODO
	}
}

func loginCmd(useCase UseCase, logger *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "log in to the bpla",
		Run:   login(useCase, logger),
	}

	cmd.Flags().StringP("username", "u", "", "Username")
	cmd.Flags().StringP("password", "p", "", "Password")
	_ = cmd.MarkFlagRequired("username")
	_ = cmd.MarkFlagRequired("password")

	return cmd
}
