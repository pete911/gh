package cmd

import (
	"github.com/go-git/go-git/v5"
	"github.com/pete911/gh/internal/gh"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	defaultPwd     = "/tmp/gh"
	githubTokenEnv = "GITHUB_TOKEN"
)

func GetGhClient() gh.Client {

	if token := GetStringEnv(githubTokenEnv, ""); token != "" {
		return gh.NewClientWithToken(token, gh.GitPlainClone(git.PlainClone))
	}
	return gh.NewClient(gh.GitPlainClone(git.PlainClone))
}

func GetPwd() string {

	pwd, err := os.Getwd()
	if err != nil {
		log.Warn().Err(err).Msgf("cannot get pwd, setting pwd to %s", defaultPwd)
		return defaultPwd
	}
	return pwd
}

// GetStringEnv returns specified env var if it exists, otherwise specified default val is returned
func GetStringEnv(key, val string) string {

	if out, ok := os.LookupEnv(key); ok {
		return out
	}
	return val
}
