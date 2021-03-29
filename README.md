# gh
GitHub tools

## install

- install [go](https://golang.org/doc/install)
- run `make install`

## setup

Project works with public repositories out of the box. To be able to use it with private repositories, create
[personal access token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token)
on github and export it under `GITHUB_TOKEN` env. variable.

### clone all repositories

```shell
# clones repositories for specified org to current directory, if GITHUB_TOKEN env. variable is
# exported, private # repositories are cloned as well
# flags:
#   -o, --output   git clone destination directory, defaults to current directory
gh clone org-repos <org>

# clones repositories for specified user to current directory, either user has to be specified
# (all public repositories # are cloned) or GITHUB_TOKEN env. variable has to be exported (all
# repositories owned by the user are cloned)
# flags:
#   -o, --output   git clone destination directory, defaults to current directory
gh clone user-repos [user]
```

### list all repositories

```shell
# lists repositories for specified org, if GITHUB_TOKEN env. variable is exported, private
# repositories are listed as well
# flags:
#   --sort-by string   sort output by visibility, size, language, issues or stars
gh list org-repos <org>

# lists repositories for specified user, either user has to be specified (all public
# repositories are listed) or GITHUB_TOKEN env. variable has to be exported (all
# repositories owned by the user are listed)
# flags:
#   --sort-by string   sort output by visibility, size, language, issues or stars
gh list user-repos [user]
```
