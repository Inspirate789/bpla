package cli

import (
	"log/slog"

	"github.com/spf13/cobra"
)

func Commands(useCase UseCase, logger *slog.Logger) []*cobra.Command {
	return []*cobra.Command{
		loginCmd(useCase, logger),
		logoutCmd(useCase, logger),
	}
}
