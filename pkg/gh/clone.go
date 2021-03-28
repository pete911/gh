package gh

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v34/github"
	"os"
	"path/filepath"
)

func CloneRepositories(c *github.Client, user, destination string) error {

	repositories, err := ListRepositories(c, user)
	if err != nil {
		return fmt.Errorf("clone repos: %w", err)
	}

	fmt.Printf("got %d repositories\n", len(repositories))
	for _, repository := range repositories {
		repositoryDestination := filepath.Join(destination, repository.Name)
		fmt.Printf("cloning %s to %s\n", repository.CloneURL, repositoryDestination)
		if err := gitClone(repositoryDestination, repository.CloneURL); err != nil {
			return err
		}
	}
	return nil
}

func gitClone(destination, url string) error {

	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		fmt.Printf("git clone: repository %s already exists in %s\n", url, destination)
		return nil
	}
	return err
}
