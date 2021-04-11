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

#### example

```shell
$ gh clone org-repos ori-edge
1:00PM INF got 2 repositories
Enumerating objects: 10, done.
Counting objects: 100% (10/10), done.
Compressing objects: 100% (9/9), done.
Total 10 (delta 1), reused 6 (delta 0), pack-reused 0
1:00:05 INF cloned https://github.com/ori-edge/all-in-one-opennebula.git to /tmp/github.com/all-in-one-opennebula
1:00:06 WRN git clone: repository https://github.com/ori-edge/k8s_gateway.git already exists in /tmp/github.com/k8s_gateway
```

### list all repositories

```shell
# lists repositories for specified org, if GITHUB_TOKEN env. variable is exported, private
# repositories are listed as well
# flags:
#   --no-forks         exclude forked repositories in the list
#   --sort-by string   sort output by visibility, size, language, issues or stars
gh list org-repos <org>

# lists repositories for specified user, either user has to be specified (all public
# repositories are listed) or GITHUB_TOKEN env. variable has to be exported (all
# repositories owned by the user are listed)
# flags:
#   --no-forks         exclude forked repositories in the list
#   --sort-by string   sort output by visibility, size, language, issues or stars
gh list user-repos [user]
```

#### example

```shell
$ gh list user-repos kiich
Name                      Visibility Size  Language Issues Stars Topics Fork
ansible-recipes           public     21    Python   0      0     []     true
flux-get-started          public     79    Smarty   0      0     []     true
helm                      public     4908  Go       0      0     []     true
helm-operator-get-started public     544            0      0     []     true
kube-aws                  public     36151 Go       0      0     []     true
logrotate                 public     60    Shell    0      0     []     true
```
