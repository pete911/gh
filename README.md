# gh
GitHub tools

## install

- install [go](https://golang.org/doc/install)
- run `make install`

### clone all repositories

```shell
# clones repositories for specified org to current directory, if GITHUB_TOKEN env. variable is exported, private
# repositories are cloned as well
# flags:
#   -o, --output   git clone destination directory, defaults to current directory
gh clone org-repos <org>

# clones repositories for specified user to current directory, either user has to be specified (all public repositories
# are cloned) or GITHUB_TOKEN env. variable has to be exported (all repositories owned by the user are cloned)
# flags:
#   -o, --output   git clone destination directory, defaults to current directory
gh clone user-repos [user]
```

### list all repositories

```shell
# lists repositories for specified org, if GITHUB_TOKEN env. variable is exported, private repositories are listed as well
gh list org-repos <org>

# lists repositories for specified user, either user has to be specified (all public repositories are listed) or
# GITHUB_TOKEN env. variable has to be exported (all repositories owned by the user are listed)
gh list user-repos [user]
```
