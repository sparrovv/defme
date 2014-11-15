package wordnik

import (
	"encoding/json"
	"fmt"

	"strings"

	"github.com/sparrovv/defme/configuration"
)

type Related struct {
	Words []string
}

func FetchRelated(config configuration.Config, word string) (result Related, err error) {
	url := fmt.Sprintf("%s/%s/%s/relatedWords?useCanonical=true&relationshipTypes=synonym&limitPerRelationshipType=10&api_key=%s", config.WordnikHost, "v4/word.json", EncodeSpace(strings.ToLower(word)), config.WordnikApiKey)

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
