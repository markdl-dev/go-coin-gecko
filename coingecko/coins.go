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
type CoinsMarketData []CoinsMarket

// CoinsMarketStruct is the coins market result
type CoinsMarket struct {
	ID                                  string   `json:"id"`
	Symbol                              string   `json:"symbol"`
	Name                                string   `json:"name"`
	CoinImage                           string   `json:"image"`
	CurrentPrice                        float64  `json:"current_price"`
	MarketCap                           float64  `json:"market_cap"`
	MarketCapRank                       uint16   `json:"market_cap_rank"`
	FullyDilutedValuation               float64  `json:"fully_diluted_valuation"`
	TotalVolume                         float64  `json:"total_volume"`
	High24H                             float64  `json:"high_24h"`
	Low24H                              float64  `json:"low_24h"`
	PriceChange24H                      float64  `json:"price_change_24h"`
	PriceChangePercentage24H            float64  `json:"price_change_percentage_24h"`
	MarketCapChange24H                  float64  `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H        float64  `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                   float64  `json:"circulating_supply"`
	TotalSupply                         float64  `json:"total_supply"`
	MaxSupply                           float64  `json:"max_supply"`
	ATH                                 float64  `json:"ath"`
	ATHChangePercentage                 float64  `json:"ath_change_percentage"`
	ATHDate                             string   `json:"ath_date"`
	ATL                                 float64  `json:"atl"`
	ATLChangePercentage                 float64  `json:"atl_change_percentage"`
	ATLDate                             string   `json:"atl_date"`
	ROI                                 *ROI     `json:"roi,omitempty"`
	LastUpdated                         string   `json:"last_updated"`
	PriceChangePercentage1HInCurrency   *float64 `json:"price_change_percentage_1h_in_currency,omitempty"`
	PriceChangePercentage24HInCurrency  *float64 `json:"price_change_percentage_24h_in_currency,omitempty"`
	PriceChangePercentage7DInCurrency   *float64 `json:"price_change_percentage_7d_in_currency,omitempty"`
	PriceChangePercentage14DInCurrency  *float64 `json:"price_change_percentage_14d_in_currency,omitempty"`
	PriceChangePercentage24DInCurrency  *float64 `json:"price_change_percentage_24d_in_currency,omitempty"`
	PriceChangePercentage30DInCurrency  *float64 `json:"price_change_percentage_30d_in_currency,omitempty"`
	PriceChangePercentage200dInCurrency *float64 `json:"price_change_percentage_200d_in_currency,omitempty"`
	PriceChangePercentage1yInCurrency   *float64 `json:"price_change_percentage_1y_in_currency,omitempty"`
}

// ROI is the roi result
type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

type Coin struct {
	ID                           string      `json:"id"`
	Symbol                       string      `json:"symbol"`
	Name                         string      `json:"name"`
	AssetPlatformID              string      `json:"asset_platform_id"`
	BlockTimeInMinutes           int32       `json:"block_time_in_minutes"`
	HashingAlgorithm             string      `json:"hashing_algorithm"`
	Categories                   []string    `json:"categories"`
	Description                  Description `json:"description"`
	Links                        *Links      `json:"links"`
	Image                        *Image      `json:"image"`
	CountryOrigin                string      `json:"country_origin"`
	GenesisDate                  string      `json:"genesis_date"`
	ContractAddress              string      `json:"contract_address"`
	SentimentVotesUpPercentage   float32     `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float32     `json:"sentimate_votes_down_percentage"`
	MarketCapRank                uint16      `json:"market_cap_rank"`
	CoinGeckoRank                uint16      `json:"coingecko_rank"`
	CoinGeckoScore               float32     `json:"coingecko_score"`
	DeveloperScore               float32     `json:"developer_score"`
	CommunityScore               float32     `json:"community_score"`
	LiquidityScore               float32     `json:"liquidity_score"`
	PublicInterestScore          float32     `json:"public_interest_score"`
	MarketData                   *MarketData `json:"market_data"`
	// CommunityData                *CommunityDataItem  `json:"community_data"`
	// DeveloperData                *DeveloperDataItem  `json:"developer_data"`
	// PublicInterestStats          *PublicInterestItem `json:"public_interest_stats"`
	// StatusUpdates                *[]StatusUpdateItem `json:"status_updates"`
	// LastUpdated                  string              `json:"last_updated"`
	// Tickers                      *[]TickerItem       `json:"tickers"`
}

type Description map[string]string

type Links struct {
	HomePage                   []string  `json:"homepage"`
	BlockChainSite             []string  `json:"blockchain_site"`
	OfficialForumURL           []string  `json:"official_forum_url"`
	ChatURL                    []string  `json:"chat_url"`
	AnnouncementURL            []string  `json:"announcement_url"`
	TwitterScreenName          string    `json:"twitter_screen_name"`
	FacebookUsername           string    `json:"facebook_username"`
	BitcointalkThreadIdentifer int       `json:"bitcointalk_thread_identifier"`
	SubredditURL               string    `json:"subreddit_url"`
	TelegramChannelIdentifier  string    `json:"telegram_channel_identifier"`
	ReposURL                   *ReposURL `json:"repos_url"`
}

type ReposURL struct {
	Github    []string `json:"github"`
	Bitbucket []string `json:"bitbucket"`
}

type Image struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type CurrencyPrice map[string]float64

type MarketData struct {
	CurrentPrice                       CurrencyPrice     `json:"current_price"`
	ATH                                CurrencyPrice     `json:"ath"`
	ATHChangePercentage                CurrencyPrice     `json:"ath_change_percentage"`
	ATHDate                            map[string]string `json:"ath_date"`
	ATL                                CurrencyPrice     `json:"atl"`
	ATLChangePercentage                CurrencyPrice     `json:"atl_change_percentage"`
	ATLDate                            map[string]string `json:"atl_date"`
	MarketCap                          CurrencyPrice     `json:"market_cap"`
	MarketCapRank                      int16             `json:"market_cap_rank"`
	TotalVolume                        CurrencyPrice     `json:"total_volume"`
	High24H                            CurrencyPrice     `json:"high_24h"`
	Low24H                             CurrencyPrice     `json:"low_24h"`
	PriceChange24H                     float64           `json:"price_change_24h"`
	PriceChangePercentage24H           float64           `json:"price_change_percentage_24h"`
	PriceChangePercentage7D            float64           `json:"price_change_percentage_7d"`
	PriceChangePercentage14D           float64           `json:"price_change_percentage_14d"`
	PriceChangePercentage30D           float64           `json:"price_change_percentage_30d"`
	PriceChangePercentage60D           float64           `json:"price_change_percentage_60d"`
	PriceChangePercentage200D          float64           `json:"price_change_percentage_200d"`
	PriceChangePercentage1Y            float64           `json:"price_change_percentage_1y"`
	MarketCapChange24H                 float64           `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H       float64           `json:"market_cap_change_percentage_24h"`
	PriceChangePercentage24HInCurrency CurrencyPrice     `json:"price_change_percentage_24h_in_currency"`
	TotalSupply                        *float64          `json:"total_supply"`
	CirculatingSupply                  float64           `json:"circulating_supply"`
	LastUpdated                        string            `json:"last_updated"`
}

// GetMarkets gets List all supported coins price, market cap, volume, and market related data
// https://api.coingecko.com/api/v3/coins/markets
func (s *CoinsService) GetMarketsWithContext(ctx context.Context, vsCurrency string, coinIDSlice []string) (*CoinsMarketData, *http.Response, error) {
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
	urlValues.Add("price_change_percentage", "1h,24h,7d,14d,30d")

	u.RawQuery = urlValues.Encode()

	req, err := s.client.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	coinsMarketData := new(CoinsMarketData)
	resp, err := s.client.Do(req, coinsMarketData)
	if err != nil {
		return nil, resp, err
	}
	return coinsMarketData, resp, nil
}

// GetExchangeRates wraps GetMarketsWithContext using the background context
func (s *CoinsService) GetMarkets(currency string, coinIDSlice []string) (*CoinsMarketData, *http.Response, error) {
	return s.GetMarketsWithContext(context.Background(), currency, coinIDSlice)
}

// Get current data (name, price, market, … including exchange tickers) for a coin.
// https://api.coingecko.com/api/v3/coins/{id}
func (s *CoinsService) GetCoinWithContext(ctx context.Context, coinID string) (*Coin, *http.Response, error) {
	if len(coinID) == 0 {
		return nil, nil, errors.New("target coin id is required")
	}

	u := url.URL{
		Path: "/coins/" + coinID,
	}

	// urlValues := url.Values{}
	// urlValues.Add("currency", vsCurrencyStr)

	// u.RawQuery = urlValues.Encode()

	req, err := s.client.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	coinInfo := new(Coin)
	resp, err := s.client.Do(req, coinInfo)
	if err != nil {
		return nil, resp, err
	}
	return coinInfo, resp, nil
}

// GetCoin wraps GetCoinWithContext using the background context
func (s *CoinsService) GetCoin(ID string) (*Coin, *http.Response, error) {
	return s.GetCoinWithContext(context.Background(), ID)
}
