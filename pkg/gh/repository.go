package gh

import (
	"fmt"
	"github.com/google/go-github/v34/github"
)

type Repository struct {
	Name     string
	CloneURL string
}

func (c Client) ListRepositoriesByOrg(org string) ([]Repository, error) {

	var out []Repository
	ghRepositories, err := c.listRepositoriesByOrg(org)
	if err != nil {
		return nil, err
	}

	for _, ghRepository := range ghRepositories {
		out = append(out, Repository{Name: *ghRepository.Name, CloneURL: *ghRepository.CloneURL})
	}
	return out, nil
}

// list public repositories for a user, if user is not specified, all repositories owned by authenticated
// user are listed
func (c Client) ListRepositories(user string) ([]Repository, error) {

	var out []Repository
	ghRepositories, err := c.listRepositories(user)
	if err != nil {
		return nil, err
	}

	for _, ghRepository := range ghRepositories {
		out = append(out, Repository{Name: *ghRepository.Name, CloneURL: *ghRepository.CloneURL})
	}
	return out, nil
}

func (c Client) listRepositoriesByOrg(org string) ([]*github.Repository, error) {

	opt := &github.RepositoryListByOrgOptions{
		Type: "sources",
	}

	var repositories []*github.Repository
	for {
		repos, resp, err := c.ghClient.Repositories.ListByOrg(c.ctx, org, opt)
		if err != nil {
			return nil, fmt.Errorf("list repositories by org page %d: %w", opt.Page, err)
		}
		repositories = append(repositories, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return repositories, nil
}

func (c Client) listRepositories(user string) ([]*github.Repository, error) {

	opt := &github.RepositoryListOptions{
		Affiliation: "owner",
	}

	var repositories []*github.Repository
	for {
		repos, resp, err := c.ghClient.Repositories.List(c.ctx, user, opt)
		if err != nil {
			return nil, fmt.Errorf("list repositories page %d: %w", opt.Page, err)
		}
		repositories = append(repositories, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return repositories, nil
}
