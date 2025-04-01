package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log in to the bpla",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		fmt.Println(username, password)
		fmt.Println("Login succeded!")
		// TODO
	},
}
