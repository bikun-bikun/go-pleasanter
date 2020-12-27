package pleasanter

import "net/http"

type Client struct {
	endpoint   string
	apiVersion string
	apiKey     string
	*http.Client
}

func NewClient(endpoint, apiversion, apikey string) *Client {
	return &Client{endpoint, apiversion, apikey, http.DefaultClient}
}
