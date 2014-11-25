package wordnik

import (
	"encoding/json"
	"fmt"

	"strings"
)

type Example struct {
	Text string `json:"text"`
}

func (c *Client) FetchExamples(word string) (result []Example, err error) {
	url := fmt.Sprintf("%s/%s/%s/examples?api_key=%s", c.Host, "v4/word.json", strings.ToLower(word), c.ApiKey)

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
