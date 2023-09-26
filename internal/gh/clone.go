package gh

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log/slog"
	"os"
	"path/filepath"
)

func (c Client) CloneOrgRepositories(org, destination string) error {

	repositories, err := c.ListRepositoriesByOrg(org)
	if err != nil {
		return fmt.Errorf("clone org repos: %w", err)
	}
	return c.gitCloneRepositories(repositories, destination)
}

func (c Client) CloneUserRepositories(user, destination string) error {

	repositories, err := c.ListRepositories(user)
	if err != nil {
		return fmt.Errorf("clone user repos: %w", err)
	}
	return c.gitCloneRepositories(repositories, destination)
}

func (c Client) gitCloneRepositories(repositories Repositories, destination string) error {

	slog.Info(fmt.Sprintf("got %d repositories", len(repositories)))
	for _, repository := range repositories {
		repositoryDestination := filepath.Join(destination, repository.Name)
		if err := c.gitCloneRepository(repository.CloneURL, repositoryDestination); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) gitCloneRepository(url, destination string) error {

	var auth *http.BasicAuth
	if c.HasToken {
		auth = &http.BasicAuth{Username: "gh", Password: c.token}
	}

	_, err := c.gitClient.PlainClone(destination, false, &git.CloneOptions{
		URL:      url,
		Auth:     auth,
		Progress: os.Stdout,
	})

	if err != nil {
		if errors.Is(err, git.ErrRepositoryAlreadyExists) {
			slog.Warn(fmt.Sprintf("git clone: repository %s already exists in %s", url, destination))
			return nil
		}
		return err
	}

	slog.Info(fmt.Sprintf("cloned %s to %s", url, destination))
	return nil
}
