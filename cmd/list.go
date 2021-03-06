package cmd

import (
	"fmt"
	"github.com/pete911/gh/internal/gh"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var (
	sortByFlag  string
	noForksFlag bool

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

	listCmd.PersistentFlags().StringVar(&sortByFlag, "sort-by", "", "sort output by visibility, size, language, issues or stars")
	listCmd.PersistentFlags().BoolVar(&noForksFlag, "no-forks", false, "exclude forked repositories in the list")
	listCmd.AddCommand(listOrgReposCmd)
	listCmd.AddCommand(listUserReposCmd)
}

func listOrgReposCmdRunE(_ *cobra.Command, args []string) error {

	org := args[0]
	repositories, err := GetGhClient().ListRepositoriesByOrg(org)
	if err != nil {
		return err
	}

	if sortByFlag != "" {
		repositories.SortBy(sortByFlag)
	}

	printList(repositories, noForksFlag)
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

	if sortByFlag != "" {
		repositories.SortBy(sortByFlag)
	}

	printList(repositories, noForksFlag)
	return nil
}

func printList(repositories []gh.Repository, noForks bool) {

	if noForks {
		repositories = removeForks(repositories)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tVisibility\tSize\tLanguage\tIssues\tStars\tTopics\tFork")
	for _, r := range repositories {
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%d\t%v\t%t\n",
			r.Name, r.Visibility, r.Size, r.Language, r.OpenIssuesCount, r.StargazersCount, r.Topics, r.Fork)
	}
	w.Flush()
}

func removeForks(repositories []gh.Repository) []gh.Repository {

	var out []gh.Repository
	for _, repository := range repositories {
		if repository.Fork {
			continue
		}
		out = append(out, repository)
	}
	return out
}
