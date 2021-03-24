package coingecko

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// CoinsService handles Coin endpoints for CoinGecko API
type CoinsService struct {
	client *Client
}

// CoinsMarket represents the coins market in CoinGecko
type CoinsMarket []CoinsMarketStruct

// CoinsMarketStruct is the coins market result
type CoinsMarketStruct struct {
	ID                                  string  `json:"id"`
	Symbol                              string  `json:"symbol"`
	Name                                string  `json:"name"`
	CoinImage                           string  `json:"image"`
	CurrentPrice                        float64 `json:"current_price"`
	MarketCap                           float64 `json:"market_cap"`
	MarketCapRank                       uint16  `json:"market_cap_rank"`
	FullyDilutedValuation               float64 `json:"fully_diluted_valuation"`
	TotalVolume                         float64 `json:"total_volume"`
	High24H                             float64 `json:"high_24h"`
	Low24H                              float64 `json:"low_24h"`
	PriceChange24H                      float64 `json:"price_change_24h"`
	PriceChangePercentage24H            float64 `json:"price_change_percentage_24h"`
	MarketCapChange24H                  float64 `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H        float64 `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                   float64 `json:"circulating_supply"`
	TotalSupply                         float64 `json:"total_supply"`
	MaxSupply                           float64 `json:"max_supply"`
	ATH                                 float64 `json:"ath"`
	ATHChangePercentage                 float64 `json:"ath_change_percentage"`
	ATHDate                             string  `json:"ath_date"`
	ATL                                 float64 `json:"atl"`
	ATLChangePercentage                 float64 `json:"atl_change_percentage"`
	ATLDate                             string  `json:"atl_date"`
	ROI                                 ROI     `json:"roi"`
	LastUpdated                         string  `json:"last_updated"`
	PriceChangePercentage1hInCurrency   float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency  float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7DInCurrency   float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14DInCurrency  float64 `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage24DInCurrency  float64 `json:"price_change_percentage_24d_in_currency"`
	PriceChangePercentage30DInCurrency  float64 `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage200dInCurrency float64 `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency   float64 `json:"price_change_percentage_1y_in_currency"`
}

// ROI is the roi result
type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

// GetMarkets gets List all supported coins price, market cap, volume, and market related data
// https://api.coingecko.com/api/v3/coins/markets
func (s *CoinsService) GetMarketsWithContext(ctx context.Context, vsCurrency string, coinIDSlice []string) (*CoinsMarket, *http.Response, error) {
	if len(vsCurrency) == 0 {
		return nil, nil, errors.New("target currency is required")
	}

	u := url.URL{
		Path: "/coins/markets",
	}

	urlValues := url.Values{}
	urlValues.Add("vs_currency", vsCurrency)

	if len(coinIDSlice) > 0 {
		coinIDStr := strings.Join(coinIDSlice, ",")
		urlValues.Add("ids", coinIDStr)
	}

	// TODO convert this to param
	urlValues.Add("order", "market_cap_desc")
	urlValues.Add("per_page", "100")
	urlValues.Add("page", "1")
	urlValues.Add("sparkline", "false")
	urlValues.Add("price_change_percentage", "1h,24h,7d,14d,30d,200d,1y")

	u.RawQuery = urlValues.Encode()

	req, err := s.client.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	coinsMarket := new(CoinsMarket)
	resp, err := s.client.Do(req, coinsMarket)
	if err != nil {
		return nil, resp, err
	}
	return coinsMarket, resp, nil
}

// GetExchangeRates wraps GetMarketsWithContext using the background context
func (s *CoinsService) GetMarkets(currency string, coinIDSlice []string) (*CoinsMarket, *http.Response, error) {
	return s.GetMarketsWithContext(context.Background(), currency, coinIDSlice)
}
