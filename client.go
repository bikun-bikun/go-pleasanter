package pleasanter

import "net/http"

type Client struct {
	requestBase
	endpoint string
	*http.Client
}

func NewClient(endpoint, apiversion, apikey string) *Client {
	return &Client{requestBase{apiversion, apikey}, endpoint, http.DefaultClient}
}
