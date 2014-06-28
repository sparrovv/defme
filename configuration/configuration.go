package configuration

import (
  "os"
  "errors"
)

type Config struct {
  WordnikApiKey string `json:"wordnikApiKey"`
  WordnikHost string `json:"wordnikHost"`
}

func FromEnv() (c Config, err error){
  // FIXME: There has to be a better way to store host, so it's easy in the test to configure.
  c.WordnikHost = "http://api.wordnik.com"
  c.WordnikApiKey = os.Getenv("WORDNIK_API_KEY")

  if len(c.WordnikApiKey) == 0{
    err = errors.New("WORDNIK_API_KEY is not set")
  }

  return
}

