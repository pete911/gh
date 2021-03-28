package cmd

import (
	"fmt"
	"github.com/pete911/gh/pkg/gh"
	"os"
)

const (
	defaultPwd     = "/tmp/gh"
	githubTokenEnv = "GITHUB_TOKEN"
)

func GetGhClient() gh.Client {

	if token := GetStringEnv(githubTokenEnv, ""); token != "" {
		return gh.NewClientWithToken(token)
	}
	return gh.NewClient()
}

func GetPwd() string {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("cannot get working directory: %v\n", err)
		fmt.Printf("setting pwd to %s\n", defaultPwd)
		return defaultPwd
	}
	return pwd
}

// returns specified env var if it exists, otherwise specified default val is returned
func GetStringEnv(key, val string) string {

	if out, ok := os.LookupEnv(key); ok {
		return out
	}
	return val
}
