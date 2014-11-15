package wordnik

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sparrovv/defme/configuration"
	"github.com/stretchr/testify/assert"
)

var relatedWordsResponse = `
[
  {
    "words": [
      "mitigate",
      "divert",
      "loosen",
      "ease"
    ],
    "relationshipType": "synonym"
  }
]
`

func TestReleatedWords(t *testing.T) {
	server := newTestServer(relatedWordsResponse)

	config := configuration.Config{
		WordnikApiKey: "myApiKey",
		WordnikHost:   server.URL,
	}

	p, err := FetchRelated(config, "TurN up")
	assert.NoError(t, err)
	assert.Equal(t, len(p.Words), 4)
	assert.Equal(t, p.Words[0], "mitigate")
}

func TestReleatedWordsWhenEmptyResponse(t *testing.T) {
	var fetchHandler = func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(rw, `[]`)
	}
	server := httptest.NewServer(http.HandlerFunc(fetchHandler))

	config := configuration.Config{
		WordnikApiKey: "myApiKey",
		WordnikHost:   server.URL,
	}

	p, err := FetchRelated(config, "TurN up")
	assert.NoError(t, err)
	assert.Equal(t, len(p.Words), 0)
}

func TestEncSpace(t *testing.T) {
	assert.Equal(t, EncodeSpace("foo bar biz"), "foo%20bar%20biz")
}
