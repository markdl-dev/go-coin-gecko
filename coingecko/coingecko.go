package coingecko

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const defaultBaseURL = "https://api.coingecko.com/api/v3/"

type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client

	// Base URL for API requests
	BaseURL *url.URL

	// Services used for talking to the Utilities in the CoinGecko API.
	Util *UtilService

	// Services used for talking to  the Exchange Rates in the CoinGecko API.
	ExchangeRate *ExchangeRateService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Util = &UtilService{client: c}
	c.ExchangeRate = &ExchangeRateService{client: c}

	return c
}

// NewRequestWithContext creates an API request
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequestWithContext(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// Relative URLs should be specified without a preceding slash since baseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Context-Type", "application/json")
	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occured
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(httpResp)
	if err != nil {
		// Even though there was an error, we will still return the response
		// in case the caller wants to inspect it further
		return httpResp, err
	}

	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer httpResp.Body.Close()
		err = json.NewDecoder(httpResp.Body).Decode(v)
	}

	return httpResp, err
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
// The caller is responsible to analyze the response body.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	err := fmt.Errorf("request failed. Please analyze the request body for more details. Status code: %d", r.StatusCode)
	return err
}
