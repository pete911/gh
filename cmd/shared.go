package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/pete911/gh/internal/gh"
	"log/slog"
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
		slog.Warn(fmt.Sprintf("setting pwd to %s: %v", defaultPwd, err))
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
