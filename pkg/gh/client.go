package gh

import (
	"context"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

const timeout = 10 * time.Second

type Client struct {
	HasToken bool
	ghClient *github.Client
	ctx      context.Context
}

func NewClient() Client {

	return Client{
		ghClient: github.NewClient(&http.Client{Timeout: timeout}),
		ctx:      context.Background(),
	}
}

func NewClientWithToken(token string) Client {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	tc.Timeout = timeout

	return Client{
		HasToken: true,
		ghClient: github.NewClient(tc),
		ctx:      ctx,
	}
}
