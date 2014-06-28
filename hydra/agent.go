package hydra

import (
  "strings"
  "sync"
  "encoding/json"
  "fmt"
  "log"
  "github.com/sparrovv/defme/wordnik"
  "github.com/sparrovv/gotr/googletranslate"
  "github.com/sparrovv/defme/configuration"
)

const termIndentation = "   "

func FormattedResponse(term string, config configuration.Config, translateTo string, toJSON bool) (formattedResponse string){
  ch := make(chan Result, 3)
  var group sync.WaitGroup
  group.Add(cap(ch))

  go func() {
    defer group.Done()
    defs, err := wordnik.FetchDef(config, term)
    if err != nil { log.Printf("Wordink fetch definition error %v", err)}
    ch <- WordnikPhrases{defs}
  }()

  go func() {
    defer group.Done()
    related, err := wordnik.FetchRelated(config, term)
    if err != nil { log.Printf("Wordink fetch related error %v", err)}
    ch <- WordingRelatedProxy{related}
  }()

  go func() {
    defer group.Done()
    phrase, err := googletranslate.Translate("en", translateTo, term)
    if err != nil { log.Printf("Wordink fetch related error %v", err) }
    ch <- GooglePhraseProxy{phrase}
  }()

  go func(){
    group.Wait()
    close(ch)
  }()


  apiResponses := make([]interface{}, 0)
  for response := range ch {
    apiResponses = append(apiResponses, response)
  }

  strResponses := make([]string, 0)
  if toJSON {
    jsonObj, _ := json.Marshal(makeJSONResponse(apiResponses))
    formattedResponse = string(jsonObj)
  } else {
    for _, response := range apiResponses {
      strResponses = append(strResponses, response.(Result).Result())
    }
    formattedResponse = strings.Join(strResponses, "\n")
  }

  return
}

// Standardize interfaces - same interace to Google Translate and Wordink API
type Result interface {
  Result() string
}

type GooglePhraseProxy struct{
  googletranslate.Phrase
}

func (p GooglePhraseProxy) Result() string{
  return fmt.Sprintf("Translation:\n%s%s\n%s%s",termIndentation, p.Phrase.Translation, termIndentation, strings.Join(p.Phrase.ExtraMeanings,", "))
}

type WordnikPhrases struct{
  wordnik.PhraseResult
}

func (p WordnikPhrases) Result() string{
  phraseResults := make([]string,0)

  for _, phrase := range p.PhraseResult.Phrases {
    phraseResults = append(phraseResults, termIndentation + phrase.Text)
  }

  return fmt.Sprintf("Definition:\n%s", strings.Join(phraseResults, "\n"))
}

func (p WordnikPhrases) mapTexts() []string{
  phraseResults := make([]string,0)

  for _, phrase := range p.PhraseResult.Phrases {
    phraseResults = append(phraseResults, phrase.Text)
  }

  return phraseResults
}

type WordingRelatedProxy struct{
  wordnik.Related
}

func (r WordingRelatedProxy) Result() string{
  return fmt.Sprintf("Related:\n%s%v", termIndentation, strings.Join(r.Related.Words, ", "))
}

type JSONResponse struct {
  Translation string `json:"translation"`
  ExtraTranslations []string `json:"extraTranslations"`
  Definitions []string `json:"definitions"`
  Synonyms []string `json:"synonyms"`
}

func makeJSONResponse(result []interface{}) JSONResponse {
  jresponse := JSONResponse{}

  for i, ob := range result{
    switch ob.(type) {
    case WordingRelatedProxy:
      jresponse.Synonyms = result[i].(WordingRelatedProxy).Related.Words
    case WordnikPhrases:
      jresponse.Definitions = result[i].(WordnikPhrases).mapTexts()
    case GooglePhraseProxy:
      jresponse.Translation = result[i].(GooglePhraseProxy).Phrase.Translation
      jresponse.ExtraTranslations = result[i].(GooglePhraseProxy).Phrase.ExtraMeanings
    }
  }

  return jresponse
}
