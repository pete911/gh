package gh

import (
	"fmt"
	"github.com/google/go-github/v34/github"
	"sort"
	"strings"
)

type Repositories []Repository

// sort repositories by visibility, size, language, issues or stars
func (r Repositories) SortBy(field string) {

	// sort numeric values by highest first
	sort.Slice(r, func(i, j int) bool {
		switch strings.ToLower(field) {
		case "visibility":
			return r[i].Visibility < r[j].Visibility
		case "size":
			return r[i].Size > r[j].Size
		case "language":
			return r[i].Language < r[j].Language
		case "issues":
			return r[i].OpenIssuesCount > r[j].OpenIssuesCount
		case "stars":
			return r[i].StargazersCount > r[j].StargazersCount
		default:
			return r[i].Name < r[j].Name
		}
	})
}

type Repository struct {
	Name            string
	CloneURL        string
	Visibility      string
	Size            int
	Language        string
	OpenIssuesCount int
	StargazersCount int
	Topics          []string
}

func (c Client) ListRepositoriesByOrg(org string) (Repositories, error) {

	var out []Repository
	ghRepositories, err := c.listRepositoriesByOrg(org)
	if err != nil {
		return nil, err
	}

	for _, ghRepository := range ghRepositories {
		out = append(out, toRepository(ghRepository))
	}
	return out, nil
}

// list public repositories for a user, if user is not specified, all repositories owned by authenticated
// user are listed
func (c Client) ListRepositories(user string) (Repositories, error) {

	var out []Repository
	ghRepositories, err := c.listRepositories(user)
	if err != nil {
		return nil, err
	}

	for _, ghRepository := range ghRepositories {
		out = append(out, toRepository(ghRepository))
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

func toRepository(ghRepository *github.Repository) Repository {

	return Repository{
		Name:            toString(ghRepository.Name),
		CloneURL:        toString(ghRepository.CloneURL),
		Visibility:      toString(ghRepository.Visibility),
		Size:            toInt(ghRepository.Size),
		Language:        toString(ghRepository.Language),
		OpenIssuesCount: toInt(ghRepository.OpenIssuesCount),
		StargazersCount: toInt(ghRepository.StargazersCount),
		Topics:          ghRepository.Topics,
	}
}

func toString(s *string) string {

	if s == nil {
		return ""
	}
	return *s
}

func toInt(i *int) int {

	if i == nil {
		return 0
	}
	return *i
}
