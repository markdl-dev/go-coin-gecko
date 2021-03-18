package coingecko

import (
	"context"
	"net/http"
)

// ExchangeRateService handles Exchange Rates for CoinGecko API
type ExchangeRateService struct {
	client *Client
}

// ExchangeRates represents the exchange rates in CoinGecko
type ExchangeRates struct {
	Rates Rates
}

// Rates are the rates in ExchangeRates
type Rates map[string]RatesStruct

// RatesStruct are the rate struct in Rates
type RatesStruct struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// GetExchangeRatesWithContext gets the BTC-to-Currency exchange rates in CoinGecko
func (s *ExchangeRateService) GetExchangeRatesWithContext(ctx context.Context) (*ExchangeRates, *http.Response, error) {
	apiEndpoint := "/exchange_rates"
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	exchangeRates := new(ExchangeRates)
	resp, err := s.client.Do(req, exchangeRates)
	if err != nil {
		return nil, resp, err
	}
	return exchangeRates, resp, nil
}

// GetExchangeRates wraps GetExchangeRatesWithContext using the background context
func (s *ExchangeRateService) GetExchangeRates() (*ExchangeRates, *http.Response, error) {
	return s.GetExchangeRatesWithContext(context.Background())
}
