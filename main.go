package main

import (
	"github.com/pete911/gh/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
