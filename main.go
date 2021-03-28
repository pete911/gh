package main

import (
	"fmt"
	"github.com/pete911/gh/cmd"
	"os"
)

// TODO
// add flags to clone private repos (add flags for username and password ...)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
