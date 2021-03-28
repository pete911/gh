package main

import (
	"fmt"
	"github.com/pete911/gh/cmd"
	"os"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
