package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version = "dev"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gh",
		Run:   versionCmdRun,
	}
)

func versionCmdRun(_ *cobra.Command, _ []string) {
	fmt.Println(Version)
}
