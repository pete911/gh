package cmd

import (
	"fmt"
	"github.com/pete911/gh/pkg/gh"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
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

	printList(repositories)
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

	printList(repositories)
	return nil
}

func printList(repositories []gh.Repository) {

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tVisibility\tSize\tLanguage\tIssues\tStars\tTopics")
	for _, r := range repositories {
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%d\t%v\n",
			r.Name, r.Visibility, r.Size, r.Language, r.OpenIssuesCount, r.StargazersCount, r.Topics)
	}
	w.Flush()
}
