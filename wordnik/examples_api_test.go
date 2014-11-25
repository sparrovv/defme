package wordnik

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var examplesJSON = `
{
  "examples": [
    {
      "year": 2011,
      "word": "scry",
      "url": "http://api.wordnik.com/v4/mid/acd5bd09d200669dad519069df8043bceb1c8e5b870fa67983039c21d3093a35",
      "provider": {
        "name": "guardian",
        "id": 709
      },
      "text": "Dee was haunted by his shortcomings: \"You know I cannot see, nor scry\" he lamented.",
      "title": "Dr Dee, Palace Theatre, Manchester | First night review",
      "documentId": 32570310,
      "exampleId": 720989108,
      "rating": 124
    },
    {
      "year": 2009,
      "word": "scry",
      "url": "http://api.wordnik.com/v4/mid/6f1150be04a37f7df241cb0f39b4eae4212aaeee7a903e2ea375d335dd7f1d38",
      "provider": {
        "name": "wordnik",
        "id": 711
      },
      "text": "The Savant (aka nobody knows his identity) - cybernetic mathematical supergenius who can scry into the futures of many possible timelines.",
      "title": "Superhero Nation: how to write superhero novels and comic books &raquo; CarsonArtist’s Review Forum",
      "documentId": 30420863,
      "exampleId": 925063874,
      "rating": 124
    }
  ]
}
`

func newTestServer(json string) *httptest.Server {
	var fetchHandler = func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(rw, json)
	}

	return httptest.NewServer(http.HandlerFunc(fetchHandler))
}

func TestFetchExamples(t *testing.T) {
	server := newTestServer(examplesJSON)
	defer server.Close()

	client := NewClient("myApiKey")
	client.Host = server.URL

	examples, _ := client.FetchExamples("scry")

	assert.Equal(t, len(examples), 2)

	expected := []Example{
		Example{"Dee was haunted by his shortcomings: \"You know I cannot see, nor scry\" he lamented."},
		Example{"The Savant (aka nobody knows his identity) - cybernetic mathematical supergenius who can scry into the futures of many possible timelines."},
	}

	assert.Equal(t, examples, expected)
}
