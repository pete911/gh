package cmd

import (
	"github.com/pete911/gh/pkg/gh"
	"github.com/spf13/cobra"
)

const outputFlagName = "output"

var (
	outputFlag string

	cloneCmd = &cobra.Command{
		Use:   "clone",
		Short: "clone github repositories",
	}
	cloneUserCmd = &cobra.Command{
		Use:   "user",
		Short: "clone user github repositories",
		Args:  cobra.ExactArgs(1),
		RunE:  cloneUserCmdRunE,
	}
)

func init() {
	cloneCmd.PersistentFlags().StringVarP(&outputFlag, outputFlagName, "o", getPwd(), "git clone destination directory")
	cloneCmd.AddCommand(cloneUserCmd)
}

func cloneUserCmdRunE(cmd *cobra.Command, args []string) error {

	user := args[0]
	destination := cmd.Flag(outputFlagName).Value.String()
	return gh.CloneUserRepos(ghClient, user, destination)
}
