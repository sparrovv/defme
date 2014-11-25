package wordnik

import (
	"encoding/json"
	"fmt"

	"strings"
)

type Related struct {
	Words []string
}

func (c *Client) FetchRelated(word string) (result Related, err error) {
	url := fmt.Sprintf("%s/%s/%s/relatedWords?useCanonical=true&relationshipTypes=synonym&limitPerRelationshipType=10&api_key=%s", c.Host, "v4/word.json", EncodeSpace(strings.ToLower(word)), c.ApiKey)

	body, err := makeRequest(url)

	var rArry []Related
	err = json.Unmarshal(body, &rArry)
	if err != nil {
		err = fmt.Errorf("Error parsing API response: [%v]", err)
		return
	}

	if len(rArry) > 0 {
		result = rArry[0]
	}
	return
}
