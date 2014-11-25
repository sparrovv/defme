package wordnik

import "testing"

func TestNewClient(t *testing.T) {
	apiKey := "apiKey"
	client := NewClient(apiKey)

	if client.ApiKey != apiKey {
		t.Errorf("Expecting ApiKey to eq '%v' but received '%v'", apiKey, client.ApiKey)
	}

	wHost := "http://api.wordnik.com"
	if client.Host != wHost {
		t.Errorf("Expecting wordnik.Host to eq '%v' but received '%v'", wHost, client.Host)
	}
}
