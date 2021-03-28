package gh

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"path/filepath"
)

func (c Client) CloneOrgRepositories(org, destination string) error {

	repositories, err := c.ListRepositoriesByOrg(org)
	if err != nil {
		return fmt.Errorf("clone org repos: %w", err)
	}

	fmt.Printf("got %d repositories\n", len(repositories))
	for _, repository := range repositories {
		repositoryDestination := filepath.Join(destination, repository.Name)
		fmt.Printf("cloning %s to %s\n", repository.CloneURL, repositoryDestination)
		if err := c.gitClone(repositoryDestination, repository.CloneURL); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) CloneUserRepositories(user, destination string) error {

	repositories, err := c.ListRepositories(user)
	if err != nil {
		return fmt.Errorf("clone user repos: %w", err)
	}

	fmt.Printf("got %d repositories\n", len(repositories))
	for _, repository := range repositories {
		repositoryDestination := filepath.Join(destination, repository.Name)
		fmt.Printf("cloning %s to %s\n", repository.CloneURL, repositoryDestination)
		if err := c.gitClone(repositoryDestination, repository.CloneURL); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) gitClone(destination, url string) error {

	var auth *http.BasicAuth
	if c.HasToken {
		auth = &http.BasicAuth{Username: "gh", Password: c.token}
	}

	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      url,
		Auth:     auth,
		Progress: os.Stdout,
	})

	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		fmt.Printf("git clone: repository %s already exists in %s\n", url, destination)
		return nil
	}
	return err
}
