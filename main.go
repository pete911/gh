package main

import (
	"fmt"
	"github.com/pete911/gh/cmd"
	"os"
)

// TODO
// gh clone org|organisation|organization <name> -o <path>
// gh list user <user>
// gh list org|organisation|organization <user>
// add flags to clone private repos (add flags for username and password ...)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
