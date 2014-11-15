package wordnik

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/sparrovv/defme/configuration"
	"github.com/stretchr/testify/assert"
)

var definitionsJson = `
[
  {
    "textProns": [],
    "sourceDictionary": "wiktionary",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [
      {
        "type": "grammar",
        "text": "intransitive"
      }
    ],
    "citations": [],
    "word": "turn up",
    "text": "To show up; to appear suddenly or unexpectedly.",
    "score": 0,
    "partOfSpeech": "verb",
    "attributionText": "from Wiktionary, Creative Commons Attribution/Share-Alike License",
    "attributionUrl": "http://creativecommons.org/licenses/by-sa/3.0/"
  },
  {
    "textProns": [],
    "sourceDictionary": "wiktionary",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [],
    "citations": [],
    "word": "turn up",
    "text": "A stroke of good luck.",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from Wiktionary, Creative Commons Attribution/Share-Alike License",
    "attributionUrl": "http://creativecommons.org/licenses/by-sa/3.0/"
  }
]
`

var fetchHandler = func(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	apiKey := r.Form.Get("api_key")

	mateched, err := regexp.MatchString("/v4/word.json/", r.URL.Path)

	if err != nil {
		fmt.Println(err)
	}

	if !mateched {
		fmt.Println("Not found")
		fmt.Println(r.URL.Path)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	if apiKey != "myApiKey" {
		rw.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(rw, `{"error": "Unable to identify user"}`)
		return
	}

	re, _ := regexp.Compile(`/v4/word.json/(.+)/definitions`)
	res := re.FindStringSubmatch(r.URL.Path)
	word := fmt.Sprintf("%s", res[1])
	if word != "turn up" {
		fmt.Println("NOOOOO UPPPPERCASE WORDS")
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, definitionsJson)
}

func TestFetchDef(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fetchHandler))

	config := configuration.Config{
		WordnikApiKey: "myApiKey",
		WordnikHost:   server.URL,
	}

	definitions, err := FetchDef(config, "TurN up")

	assert.NoError(t, err)
	assert.Equal(t, len(definitions), 2)
	assert.Equal(t, definitions[0].Word, "turn up")
	assert.Equal(t, definitions[0].Text, "To show up; to appear suddenly or unexpectedly.")
	assert.Equal(t, definitions[1].Word, "turn up")
	assert.Equal(t, definitions[1].Text, "A stroke of good luck.")

	server.Close()
}
