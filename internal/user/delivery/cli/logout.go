package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out from the bpla",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Logout succeded!")
		// TODO
	},
}
