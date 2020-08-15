package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/darksasori/finance/pkg/model"
	"github.com/darksasori/finance/pkg/mongodb"
	"github.com/darksasori/finance/pkg/service"
	"github.com/spf13/cobra"
)

func getTokenFile() (string, error) {
	fileUri, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	fileUri += "/finance"

	return fileUri, nil
}

var root = &cobra.Command{
	Use:   "finance",
	Short: "Finance command line interface",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		isPublic := func(name string) bool {
			if name == "help" {
				return true
			}
			if name == "login" {
				return true
			}
			return false
		}

		if isPublic(cmd.Name()) {
			return nil
		}

		fileUri, err := getTokenFile()
		if err != nil {
			return err
		}

		if _, err := os.Stat(fileUri); err != nil {
			return fmt.Errorf("Not logged")
		}

		data, err := ioutil.ReadFile(fileUri)
		if err != nil {
			return err
		}

		logged, err = userService.CheckToken(cmd.Context(), string(data))
		if err != nil {
			return err
		}

		return nil
	},
}

var userService *service.User
var logged *model.User

func Execute() {
	ctx := context.TODO()
	if err := mongodb.Connect(ctx); err != nil {
		panic(err)
	}
	userService = service.NewUser(mongodb.NewUser())

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
