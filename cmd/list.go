package cmd

import (
	"fmt"
	"github.com/pete911/gh/pkg/gh"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list github repositories",
	}
	listUserCmd = &cobra.Command{
		Use:   "user",
		Short: "list all user's github repositories",
		Args:  cobra.ExactArgs(1),
		RunE:  listUserCmdRunE,
	}
)

func init() {
	listCmd.AddCommand(listUserCmd)
}

func listUserCmdRunE(_ *cobra.Command, args []string) error {

	user := args[0]
	repositories, err := gh.ListRepositories(ghClient, user)
	if err != nil {
		return err
	}

	for _, repository := range repositories {
		fmt.Println(repository.Name)
	}
	return nil
}
