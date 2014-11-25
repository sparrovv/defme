package wordnik

type Client struct {
	ApiKey, Host string
}

func NewClient(ApiKey string) *Client {
	return &Client{ApiKey: ApiKey, Host: "http://api.wordnik.com"}
}
