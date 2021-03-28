package cmd

import (
	"github.com/spf13/cobra"
)

const outputFlagName = "output"

var (
	outputFlag string

	cloneCmd = &cobra.Command{
		Use:   "clone",
		Short: "clone github repositories",
	}
	cloneOrgReposCmd = &cobra.Command{
		Use:   "org-repos",
		Short: "clone all org github repositories",
		Args:  cobra.ExactArgs(1),
		RunE:  cloneOrgReposCmdRunE,
	}
	cloneUserReposCmd = &cobra.Command{
		Use:   "user-repos",
		Short: "clone all user github repositories",
		RunE:  cloneUserReposCmdRunE,
	}
)

func init() {

	cloneCmd.PersistentFlags().StringVarP(&outputFlag, outputFlagName, "o", GetPwd(), "git clone destination directory")
	cloneCmd.AddCommand(cloneOrgReposCmd)
	cloneCmd.AddCommand(cloneUserReposCmd)
}

func cloneOrgReposCmdRunE(cmd *cobra.Command, args []string) error {

	org := args[0]
	destination := cmd.Flag(outputFlagName).Value.String()
	return GetGhClient().CloneOrgRepositories(org, destination)
}

func cloneUserReposCmdRunE(cmd *cobra.Command, args []string) error {

	var user string
	if len(args) > 0 {
		user = args[0]
	}

	destination := cmd.Flag(outputFlagName).Value.String()
	return GetGhClient().CloneUserRepositories(user, destination)
}
