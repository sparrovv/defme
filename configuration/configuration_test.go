package configuration

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "os"
)

func TestFromEnv(t *testing.T){
  os.Setenv("WORDNIK_API_KEY", "foobar")

  conf, err := FromEnv()

  assert.NoError(t, err)

  assert.Equal(t, conf.WordnikApiKey, "foobar")
  assert.Equal(t, conf.WordnikHost, "http://api.wordnik.com")
}
