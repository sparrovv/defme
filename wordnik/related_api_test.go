package wordnik

import (
	"testing"

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
	defer server.Close()

	client := NewClient("myApiKey")
	client.Host = server.URL

	p, err := client.FetchRelated("TurN up")
	assert.NoError(t, err)
	assert.Equal(t, len(p.Words), 4)
	assert.Equal(t, p.Words[0], "mitigate")
}

func TestReleatedWordsWhenEmptyResponse(t *testing.T) {
	server := newTestServer("[]")
	defer server.Close()

	client := NewClient("myApiKey")
	client.Host = server.URL

	p, err := client.FetchRelated("TurN up")
	assert.NoError(t, err)
	assert.Equal(t, len(p.Words), 0)
}

func TestEncSpace(t *testing.T) {
	assert.Equal(t, EncodeSpace("foo bar biz"), "foo%20bar%20biz")
}
