package coingecko

import (
	"net/http"
	"net/url"
)

const defaultBaseURL = "api.coingecko.com/api/v3/"

type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client

	// Base URL for API requests
	baseURL *url.URL

	// Services used for talking to different parts of the CoinGecko API.
	Util *UtilService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, baseURL: baseURL}
	c.Util = &UtilService{client: c}

	return c
}
