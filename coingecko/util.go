package coingecko

type UtilService struct {
	client *Client
}

type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

func (s *UtilService) Ping() (*Ping, error) {
	return nil, nil
}
