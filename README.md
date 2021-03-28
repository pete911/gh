# gh
GitHub tools

## install

- install [go](https://golang.org/doc/install)
- run `make install`

### clone all repositories

```shell
# clones all public repositories for specified user (or org) to current directory
# flags:
#   -o, --output   git clone destination directory, defaults to current directory
gh clone user <user>
```

### list all repositories

```shell
# lists all public repositories for specified user (or org) to current directory
gh list user <user>
```
