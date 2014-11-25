package wordnik

import (
	"encoding/json"
	"fmt"

	"strings"
)

type Phrase struct {
	Word string `json:"word"`
	Text string `json:"text"`
}

func (c *Client) GetDefinitions(word string) (result []Phrase, err error) {
	url := fmt.Sprintf("%s/%s/%s/definitions?api_key=%s", c.Host, "v4/word.json", strings.ToLower(word), c.ApiKey)

	body, err := makeRequest(url)

	err = json.Unmarshal(body, &result)
	if err != nil {
		err = fmt.Errorf("Error parsing API response: [%v]", err)
		return
	}

	return
}
