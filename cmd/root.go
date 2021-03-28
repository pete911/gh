package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "gh",
		Short: "github utilities",
	}
)

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(cloneCmd)
}
