package wordnik

import (
	"encoding/json"
	"fmt"

	"strings"

	"github.com/sparrovv/defme/configuration"
)

type Phrase struct {
	Word string `json:"word"`
	Text string `json:"text"`
}

func FetchDef(config configuration.Config, word string) (result []Phrase, err error) {
	url := fmt.Sprintf("%s/%s/%s/definitions?api_key=%s", config.WordnikHost, "v4/word.json", strings.ToLower(word), config.WordnikApiKey)

	body, err := makeRequest(url)

	err = json.Unmarshal(body, &result)
	if err != nil {
		err = fmt.Errorf("Error parsing API response: [%v]", err)
		return
	}

	return
}
