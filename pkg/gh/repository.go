package gh

import (
	"context"
	"fmt"
	"github.com/google/go-github/v34/github"
)

type Repository struct {
	Name     string
	CloneURL string
}

func ListRepositories(c *github.Client, user string) ([]Repository, error) {

	var out []Repository
	ghRepositories, err := listRepositories(c, user, 1)
	if err != nil {
		return nil, err
	}

	for _, ghRepository := range ghRepositories {
		out = append(out, Repository{Name: *ghRepository.Name, CloneURL: *ghRepository.CloneURL})
	}
	return out, nil
}

func listRepositories(c *github.Client, user string, page int) ([]*github.Repository, error) {

	repositories, response, err := c.Repositories.List(context.Background(), user, &github.RepositoryListOptions{
		Visibility:  "public",
		Affiliation: "owner",
		ListOptions: github.ListOptions{Page: page, PerPage: 100},
	})
	if err != nil {
		return nil, fmt.Errorf("list repositories page %d: %w", page, err)
	}

	if response.NextPage != 0 {
		r, err := listRepositories(c, user, response.NextPage)
		if err != nil {
			return nil, fmt.Errorf("list repositories page %d: %w", response.NextPage, err)
		}
		repositories = append(repositories, r...)
	}
	return repositories, nil
}
