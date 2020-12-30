package pleasanter

type requestBase struct {
	ApiVersion string `json:"ApiVersion"`
	ApiKey     string `json:"ApiKey"`
}

type ErrorResult struct {
	ID         int    `json:"Id"`
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
}
