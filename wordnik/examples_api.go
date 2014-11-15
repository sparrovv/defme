package wordnik

import (
	"encoding/json"
	"fmt"

	"strings"

	"github.com/sparrovv/defme/configuration"
)

type Example struct {
	Text string `json:"text"`
}

func FetchExamples(config configuration.Config, word string) (result []Example, err error) {
	url := fmt.Sprintf("%s/%s/%s/examples?api_key=%s", config.WordnikHost, "v4/word.json", strings.ToLower(word), config.WordnikApiKey)

	body, err := makeRequest(url)

	temporaryStruct := struct {
		Examples *[]Example
	}{&result}

	err = json.Unmarshal(body, &temporaryStruct)

	if err != nil {
		err = fmt.Errorf("Error parsing API response: [%v]", err)
		return
	}

	return
}
