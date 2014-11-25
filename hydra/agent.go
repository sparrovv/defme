package hydra

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/sparrovv/defme/wordnik"
	"github.com/sparrovv/gotr/googletranslate"
)

// Combined into one struct response from all the services.
type HydraResponse struct {
	Translation       string   `json:"translation"`
	ExtraTranslations []string `json:"extraTranslations"`
	Definitions       []string `json:"definitions"`
	Synonyms          []string `json:"synonyms"`
	Examples          []string `json:"examples"`
}

// Returns JSON or formatted text from all the services
func BuildResponse(word string, wClient *wordnik.Client, translateTo string, toJSON bool) (formattedResponse string) {
	hydraResponse := HydraResponse{}
	wordnikAPI := WordinkAPI{wClient, word}

	heads := []hydraHead{
		definitionsHead{&wordnikAPI},
		examplesHead{&wordnikAPI},
		synonymsHead{&wordnikAPI},
		translationHead{translateTo, word},
	}

	var group sync.WaitGroup
	group.Add(len(heads))

	for _, h := range heads {
		go func(head hydraHead) {
			err := head.Bite(&hydraResponse)
			if err != nil {
				log.Printf("Exception: %v", err)
			}
			group.Done()
		}(h)
	}
	group.Wait()

	if toJSON {
		formattedResponse = hydraResponse.ToJSON()
	} else {
		formattedResponse = hydraResponse.ToStdOut()
	}

	return
}

// a common interface to all API calls
type hydraHead interface {
	Bite(res *HydraResponse) error
}

// google translate call
type translationHead struct {
	translateTo, word string
}

func (gt translationHead) Bite(res *HydraResponse) (err error) {
	phrase, err := googletranslate.Translate("en", gt.translateTo, gt.word)
	if err != nil {
		return
	}

	res.Translation = phrase.Translation
	res.ExtraTranslations = phrase.ExtraMeanings
	return
}

type WordinkAPI struct {
	client *wordnik.Client
	word   string
}

type definitionsHead struct{ *WordinkAPI }
type examplesHead struct{ *WordinkAPI }
type synonymsHead struct{ *WordinkAPI }

func (wordnikHead definitionsHead) Bite(res *HydraResponse) (err error) {
	defs, err := wordnikHead.client.GetDefinitions(wordnikHead.word)
	if err != nil {
		return
	}

	result := make([]string, 0)
	for _, phrase := range defs {
		result = append(result, phrase.Text)
	}

	res.Definitions = result
	return
}

func (wordnikHead synonymsHead) Bite(res *HydraResponse) (err error) {
	related, err := wordnikHead.client.GetRelated(wordnikHead.word)
	if err != nil {
		return
	}

	res.Synonyms = related.Words
	return
}

func (wordnikHead examplesHead) Bite(res *HydraResponse) (err error) {
	examples, err := wordnikHead.client.GetExamples(wordnikHead.word)
	if err != nil {
		return
	}

	result := make([]string, 0)
	for _, example := range examples {
		result = append(result, example.Text)
	}

	res.Examples = result
	return
}

func (mr HydraResponse) ToJSON() string {
	json, _ := json.Marshal(mr)
	return string(json)
}

func (mr HydraResponse) ToStdOut() (prettyOutput string) {
	indentation := "   "

	prettyOutput += fmt.Sprintf("Translation:\n%s%s\n%s%s", indentation, mr.Translation, indentation, strings.Join(mr.ExtraTranslations, ", "))
	prettyOutput += fmt.Sprintf("\nDefinition:\n%s%s", indentation, strings.Join(mr.Definitions, "\n"+indentation))
	prettyOutput += fmt.Sprintf("\nRelated:\n%s%v", indentation, strings.Join(mr.Synonyms, ", "))
	prettyOutput += fmt.Sprintf("\nExamples:\n%s%s", indentation, strings.Join(mr.Examples, "\n"+indentation))

	return
}
