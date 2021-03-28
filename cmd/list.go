package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list github repositories",
	}
	listOrgReposCmd = &cobra.Command{
		Use:   "org-repos",
		Short: "list all org github repositories",
		Args:  cobra.ExactArgs(1),
		RunE:  listOrgReposCmdRunE,
	}
	listUserReposCmd = &cobra.Command{
		Use:   "user-repos",
		Short: "list all user github repositories",
		RunE:  listUserReposCmdRunE,
	}
)

func init() {
	listCmd.AddCommand(listOrgReposCmd)
	listCmd.AddCommand(listUserReposCmd)
}

func listOrgReposCmdRunE(_ *cobra.Command, args []string) error {

	org := args[0]
	repositories, err := GetGhClient().ListRepositoriesByOrg(org)
	if err != nil {
		return err
	}

	for _, repository := range repositories {
		fmt.Println(repository.Name)
	}
	return nil
}

func listUserReposCmdRunE(_ *cobra.Command, args []string) error {

	var user string
	if len(args) > 0 {
		user = args[0]
	}

	repositories, err := GetGhClient().ListRepositories(user)
	if err != nil {
		return err
	}

	for _, repository := range repositories {
		fmt.Println(repository.Name)
	}
	return nil
}
