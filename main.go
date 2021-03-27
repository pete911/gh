package main

import (
	"fmt"
	"github.com/google/go-github/v34/github"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

// TODO
// add flags/args
//
// gh clone user <user> -o <path>
// gh clone org|organisation|organization <name> -o <path>
// gh list user <user>
// gh list org|organisation|organization <user>

func main() {

	client := github.NewClient(httpClient)
	if err := CloneUserRepos(client, "pete911", "/tmp/foo"); err != nil {
		fmt.Printf("clone user repos: %v\n", err)
		os.Exit(1)
	}
}
