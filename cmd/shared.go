package cmd

import (
	"fmt"
	"github.com/google/go-github/v34/github"
	"net/http"
	"os"
	"time"
)

const defaultPwd = "/tmp/gh"

var ghClient = github.NewClient(&http.Client{Timeout: 10 * time.Second})

func getPwd() string {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("cannot get working directory: %v\n", err)
		fmt.Printf("setting pwd to %s\n", defaultPwd)
		return defaultPwd
	}
	return pwd
}
