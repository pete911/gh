package gh

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListRepositoriesByOrg(t *testing.T) {

	t.Run("list repositories by org", func(t *testing.T) {

		response, err := ioutil.ReadFile("testdata/repositories-list.json")
		require.NoError(t, err)

		handler := func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/orgs/octocat/repos?type=sources" {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/", server.URL))
		client := NewClient(nil)
		client.ghClient.BaseURL = baseUrl

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
		client := NewClient(nil)
		client.ghClient.BaseURL = baseUrl

		_, err := client.ListRepositoriesByOrg("octocat")
		require.Error(t, err)
	})
}

func TestClient_ListRepositories(t *testing.T) {

	t.Run("list repositories", func(t *testing.T) {

		response, err := ioutil.ReadFile("testdata/repositories-list.json")
		require.NoError(t, err)

		handler := func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/users/octocat/repos?affiliation=owner" {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/", server.URL))
		client := NewClient(nil)
		client.ghClient.BaseURL = baseUrl

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
		client := NewClient(nil)
		client.ghClient.BaseURL = baseUrl

		_, err := client.ListRepositories("octocat")
		require.Error(t, err)
	})
}
