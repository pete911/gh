package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v34/github"
	"os"
	"path/filepath"
)

func CloneUserRepos(c *github.Client, user, destination string) error {

	repositories, err := listUserRepositories(c, user, 1)
	if err != nil {
		return fmt.Errorf("clone user repos: %w", err)
	}

	fmt.Printf("got %d repositories\n", len(repositories))
	for _, repository := range repositories {
		repositoryDestination := filepath.Join(destination, *repository.Name)
		fmt.Printf("cloning %s to %s\n", *repository.CloneURL, repositoryDestination)
		if err := gitClone(repositoryDestination, *repository.CloneURL); err != nil {
			return err
		}
	}
	return nil
}

func listUserRepositories(c *github.Client, user string, page int) ([]*github.Repository, error) {

	repositories, response, err := c.Repositories.List(context.Background(), user, &github.RepositoryListOptions{
		Visibility:  "public",
		Affiliation: "owner",
		ListOptions: github.ListOptions{Page: page, PerPage: 100},
	})
	if err != nil {
		return nil, fmt.Errorf("list user repositories page %d: %w", page, err)
	}

	if response.NextPage != 0 {
		r, err := listUserRepositories(c, user, response.NextPage)
		if err != nil {
			return nil, fmt.Errorf("list user repositories page %d: %w", response.NextPage, err)
		}
		repositories = append(repositories, r...)
	}
	return repositories, nil
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
