package main

import (
	"github.com/pete911/gh/cmd"
	"log/slog"
	"os"
)

func main() {

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))
	if err := cmd.RootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
