package coingecko

import (
	"context"
	"net/http"
)

// UtilService handles utility for CoinGecko API
type UtilService struct {
	client *Client
}

// Ping represents a ping in CoinGecko
type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

// Check CoinGecko API server status
func (s *UtilService) PingWithContext(ctx context.Context) (*Ping, *http.Response, error) {
	apiEndpoint := "/ping"
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	ping := new(Ping)
	resp, err := s.client.Do(req, ping)
	if err != nil {
		return nil, resp, err
	}
	return ping, resp, nil
}

// Ping wraps PingWithContext using the background context.
func (s *UtilService) Ping() (*Ping, *http.Response, error) {
	return s.PingWithContext(context.Background())
}
