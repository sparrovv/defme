package hydra

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/sparrovv/defme/configuration"
	"github.com/sparrovv/defme/wordnik"
	"github.com/sparrovv/gotr/googletranslate"
)

const termIndentation = "   "

// Combined into one struct response from all the services.
type MultiResponse struct {
	Translation       string   `json:"translation"`
	ExtraTranslations []string `json:"extraTranslations"`
	Definitions       []string `json:"definitions"`
	Synonyms          []string `json:"synonyms"`
	Examples          []string `json:"examples"`
}

func BuildResponse(term string, config configuration.Config, translateTo string, toJSON bool) (formattedResponse string) {
	ch := make(chan interface{}, 4)

	var group sync.WaitGroup
	group.Add(cap(ch))

	go func() {
		defer group.Done()

		defs, err := wordnik.FetchDef(config, term)

		if err != nil {
			log.Printf("Wordink fetch definition error %v", err)
		}

		ch <- wordnikDefinitionResult{defs}
	}()

	go func() {
		defer group.Done()
		related, err := wordnik.FetchRelated(config, term)
		if err != nil {
			log.Printf("Wordink fetch related error %v", err)
		}
		ch <- related
	}()

	go func() {
		defer group.Done()
		phrase, err := googletranslate.Translate("en", translateTo, term)
		if err != nil {
			log.Printf("Wordink fetch related error %v", err)
		}
		ch <- phrase
	}()

	go func() {
		defer group.Done()
		examples, err := wordnik.FetchExamples(config, term)
		if err != nil {
			log.Printf("Wordink fetch examples error %v", err)
		}
		ch <- wordnikExamplesResult{examples}
	}()

	go func() {
		group.Wait()
		close(ch)
	}()

	apiResponses := make([]interface{}, 0)
	for response := range ch {
		apiResponses = append(apiResponses, response)
	}

	multiResponse := makeMultiResponse(apiResponses)

	if toJSON {
		jsonObj, _ := json.Marshal(multiResponse)
		formattedResponse = string(jsonObj)
	} else {
		formattedResponse = toSTDOUT(multiResponse)
	}

	return
}

type wordnikDefinitionResult struct {
	phrases []wordnik.Phrase
}

type wordnikExamplesResult struct {
	examples []wordnik.Example
}

func (p wordnikDefinitionResult) mapTexts() []string {
	phraseResults := make([]string, 0)

	for _, phrase := range p.phrases {
		phraseResults = append(phraseResults, phrase.Text)
	}

	return phraseResults
}

func (p wordnikExamplesResult) mapTexts() []string {
	phraseResults := make([]string, 0)

	for _, example := range p.examples {
		phraseResults = append(phraseResults, termIndentation+example.Text)
	}

	return phraseResults
}

func makeMultiResponse(result []interface{}) MultiResponse {
	jresponse := MultiResponse{}

	for i, ob := range result {
		switch ob.(type) {
		case wordnik.Related:
			jresponse.Synonyms = result[i].(wordnik.Related).Words
		case wordnikDefinitionResult:
			jresponse.Definitions = result[i].(wordnikDefinitionResult).mapTexts()
		case wordnikExamplesResult:
			jresponse.Examples = result[i].(wordnikExamplesResult).mapTexts()
		case googletranslate.Phrase:
			jresponse.Translation = result[i].(googletranslate.Phrase).Translation
			jresponse.ExtraTranslations = result[i].(googletranslate.Phrase).ExtraMeanings
		}
	}

	return jresponse
}

func toSTDOUT(jresponse MultiResponse) (prettyOutput string) {
	prettyOutput += fmt.Sprintf("Translation:\n%s%s\n%s%s", termIndentation, jresponse.Translation, termIndentation, strings.Join(jresponse.ExtraTranslations, ", "))
	prettyOutput += fmt.Sprintf("\nDefinition:\n%s%s", termIndentation, strings.Join(jresponse.Definitions, "\n"+termIndentation))
	prettyOutput += fmt.Sprintf("\nRelated:\n%s%v", termIndentation, strings.Join(jresponse.Synonyms, ", "))
	prettyOutput += fmt.Sprintf("\nExamples:\n%s", strings.Join(jresponse.Examples, "\n"))

	return
}
