package cmd

import (
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v34/github"
	"github.com/pete911/gh/internal/gh"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

const (
	defaultPwd     = "/tmp/gh"
	githubTokenEnv = "GITHUB_TOKEN"
	ghTimeout      = 10 * time.Second
)

func GetGhClient() gh.Client {

	if token := GetStringEnv(githubTokenEnv, ""); token != "" {
		return gh.NewClientWithToken(token, gh.GitPlainClone(git.PlainClone))
	}
	ghClient := github.NewClient(&http.Client{Timeout: ghTimeout})
	return gh.NewClient(ghClient, gh.GitPlainClone(git.PlainClone))
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
