package gh

import (
	"fmt"
	"github.com/google/go-github/v34/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestClient_ListRepositoriesByOrg(t *testing.T) {

	t.Run("list repositories by org", func(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/orgs/octocat/repos?type=sources" {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, repositoriesListResponse)
		}
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/", server.URL))
		ghClient := github.NewClient(&http.Client{Timeout: 2 * time.Second})
		ghClient.BaseURL = baseUrl
		client := NewClient(ghClient, nil)

		repositories, err := client.ListRepositoriesByOrg("octocat")
		require.NoError(t, err)
		require.Equal(t, 1, len(repositories))
		assert.Equal(t, "Hello-World", repositories[0].Name)
		assert.Equal(t, "public", repositories[0].Visibility)
		assert.Equal(t, false, repositories[0].Fork)
	})

	t.Run("list repositories by org failure returns error", func(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/", server.URL))
		ghClient := github.NewClient(&http.Client{Timeout: 2 * time.Second})
		ghClient.BaseURL = baseUrl
		client := NewClient(ghClient, nil)

		_, err := client.ListRepositoriesByOrg("octocat")
		require.Error(t, err)
	})
}

func TestClient_ListRepositories(t *testing.T) {

	t.Run("list repositories", func(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/users/octocat/repos?affiliation=owner" {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, repositoriesListResponse)
		}
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/", server.URL))
		ghClient := github.NewClient(&http.Client{Timeout: 2 * time.Second})
		ghClient.BaseURL = baseUrl
		client := NewClient(ghClient, nil)

		repositories, err := client.ListRepositories("octocat")
		require.NoError(t, err)
		require.Equal(t, 1, len(repositories))
		assert.Equal(t, "Hello-World", repositories[0].Name)
		assert.Equal(t, "public", repositories[0].Visibility)
		assert.Equal(t, false, repositories[0].Fork)
	})

	t.Run("list repositories by org failure returns error", func(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/", server.URL))
		ghClient := github.NewClient(&http.Client{Timeout: 2 * time.Second})
		ghClient.BaseURL = baseUrl
		client := NewClient(ghClient, nil)

		_, err := client.ListRepositories("octocat")
		require.Error(t, err)
	})
}
