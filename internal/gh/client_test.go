package gh

import (
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/mock"
)

// --- shared test helpers ---

// --- mocks ---

type ClonerMock struct {
	mock.Mock
}

func (m *ClonerMock) PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {

	args := m.Called(path, isBare, o)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*git.Repository), args.Error(1)
}

// --- test data ---

// trimmed response to simplify test data
var repositoriesListByOrgResponse = `[
  {
    "name": "Hello-World",
    "fork": false,
    "clone_url": "https://github.com/octocat/Hello-World.git",
    "size": 108,
    "open_issues_count": 0,
    "stargazers_count": 80,
    "watchers_count": 80,
    "topics": [
      "octocat",
      "atom",
      "electron",
      "api"
    ],
    "archived": false,
    "visibility": "public",
    "pushed_at": "2011-01-26T19:06:43Z",
    "created_at": "2011-01-26T19:01:12Z",
    "updated_at": "2011-01-26T19:14:43Z"
  }
]`

// trimmed response to simplify test data
var repositoriesListResponse = `[
  {
    "name": "Hello-World",
    "fork": false,
    "clone_url": "https://github.com/octocat/Hello-World.git",
    "size": 108,
    "open_issues_count": 0,
    "stargazers_count": 80,
    "watchers_count": 80,
    "topics": [
      "octocat",
      "atom",
      "electron",
      "api"
    ],
    "archived": false,
    "visibility": "public",
    "pushed_at": "2011-01-26T19:06:43Z",
    "created_at": "2011-01-26T19:01:12Z",
    "updated_at": "2011-01-26T19:14:43Z"
  }
]`
