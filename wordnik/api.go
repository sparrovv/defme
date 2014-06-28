package wordnik

import (
  "encoding/json"
  "github.com/sparrovv/defme/configuration"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

type Phrase struct {
  Word string `json:"word"`
  Text string `json:"text"`
}

type PhraseResult struct {
  Phrases []Phrase
}

func FetchDef(config configuration.Config, word string) (result PhraseResult, err error) {
  url := fmt.Sprintf("%s/%s/%s/definitions?api_key=%s", config.WordnikHost, "v4/word.json", strings.ToLower(word), config.WordnikApiKey)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return
  }

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    err = fmt.Errorf("Error fetching assignments: [%v]", err)
    return
  }

  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  err = json.Unmarshal(body, &result.Phrases)
  if err != nil {
    err = fmt.Errorf("Error parsing API response: [%v]", err)
    return
  }

  return
}

type Related struct{
  Words []string
}

func FetchRelated(config configuration.Config, word string) (result Related, err error) {
  url := fmt.Sprintf("%s/%s/%s/relatedWords?useCanonical=true&relationshipTypes=synonym&limitPerRelationshipType=10&api_key=%s", config.WordnikHost, "v4/word.json", EncodeSpace(strings.ToLower(word)), config.WordnikApiKey)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return
  }

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    err = fmt.Errorf("Error fetching related words: [%v]", err)
    return
  }

  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  var rArry []Related
  err = json.Unmarshal(body, &rArry)
  if err != nil {
    err = fmt.Errorf("Error parsing API response: [%v]", err)
    return
  }

  if(len(rArry) > 0){
    result = rArry[0]
  }
  return
}

func EncodeSpace(str string) string{
  return strings.Replace(str, " ", "%20", -1)
}
