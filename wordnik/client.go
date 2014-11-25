package wordnik

type Client struct {
	ApiKey, Host string
}

func NewClient(ApiKey string) Client {
	c := Client{ApiKey: ApiKey}
	c.Host = "http://api.wordnik.com"

	return c
}
