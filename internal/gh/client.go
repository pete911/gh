package gh

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
	"time"
)

const timeout = 10 * time.Second

type GitCloner interface {
	PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error)
}

type GitPlainClone func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error)

func (g GitPlainClone) PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
	return g(path, isBare, o)
}

type Client struct {
	HasToken  bool
	token     string
	ghClient  *github.Client
	gitClient GitCloner
	ctx       context.Context
}

func NewClient(ghClient *github.Client, gitClient GitCloner) Client {

	return Client{
		ghClient:  ghClient,
		gitClient: gitClient,
		ctx:       context.Background(),
	}
}

func NewClientWithToken(token string, gitClient GitCloner) Client {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	tc.Timeout = timeout

	client := NewClient(github.NewClient(tc), gitClient)
	client.HasToken = true
	client.token = token
	return client
}
