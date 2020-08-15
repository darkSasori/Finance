package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/darksasori/finance/pkg/model"
	"github.com/spf13/cobra"
)

var user = &cobra.Command{
	Use:   "user",
	Short: "User command",
}

var userCreate = &cobra.Command{
	Use:   "create",
	Short: "Create user",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 4 {
			return fmt.Errorf("forget username, displayname, pass or checkpass")
		}
		model := model.NewUser(
			args[0],
			args[1],
			args[2],
			args[3],
		)

		if err := userService.Save(cmd.Context(), model); err != nil {
			return err
		}

		return nil
	},
}

var userLogin = &cobra.Command{
	Use:   "login",
	Short: "Login user",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("forget username and/or password")
		}

		token, err := userService.Login(cmd.Context(), args[0], args[1])
		if err != nil {
			return err
		}

		fileUri, err := getTokenFile()
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(fileUri, []byte(token), 0600); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	user.AddCommand(userCreate)
	user.AddCommand(userLogin)
	root.AddCommand(user)
}
