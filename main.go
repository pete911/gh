package main

import (
	"github.com/pete911/gh/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {

	w := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05",
	}
	log.Logger = log.Output(w)
}

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
