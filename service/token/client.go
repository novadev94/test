package token

import (
	etherCommon "github.com/ethereum/go-ethereum/common"
)

const (
	mysqlConfigType = "mysql"
	memConfigType   = "mem"

	TokenTagPromotion  = "PROMOTION"
	TokenTagVerified   = "VERIFIED"
	TokenTagUnverified = "UNVERIFIED"
	TokenTagScam       = "SCAM"
)

type TokenDataProvider interface {
	GetTokens() []Token
	GetTokenBySymbol(symbol string) Token
	GetTokenByAddress(etherCommon.Address) Token
}

type LendingProvider interface {
	GetLendingTokenByAddress(address etherCommon.Address) (Token, bool)
}

type DetailsProvider interface {
	GetTokenDetails(address etherCommon.Address) (TokenDetails, error)
}

type OverviewProvider interface {
	GetTokenOverview(address ...etherCommon.Address) ([]TokenOverview, error)
}

type Client interface {
	TokenDataProvider
	LendingProvider
	IsFeeWhitelist(address etherCommon.Address) bool
}

type Token struct {
	Address  etherCommon.Address `json:"address"`
	Symbol   string              `json:"symbol"`
	Name     string              `json:"name"`
	Decimals int                 `json:"decimals"`
}

type tokenPriceData struct {
	Symbol                    string  `json:"symbol"`
	Price                     float64 `json:"price"`
	PriceChange24h            float64 `json:"priceChange24h"`
	PriceChange1hPercentage   float64 `json:"priceChange1hPercentage"`
	PriceChange24hPercentage  float64 `json:"priceChange24hPercentage"`
	PriceChange7dPercentage   float64 `json:"priceChange7dPercentage"`
	PriceChange30dPercentage  float64 `json:"priceChange30dPercentage"`
	PriceChange200dPercentage float64 `json:"priceChange200dPercentage"`
	PriceChange1yPercentage   float64 `json:"priceChange1yPercentage"`

	MarketCap                    float64 `json:"marketCap"`
	MarketCapChange24h           float64 `json:"marketCapChange24h"`
	MarketCapChange24hPercentage float64 `json:"marketCapChange24hPercentage"`
	TotalVolume                  float64 `json:"volume24h"`

	Sparkline []float64 `json:"sparkline"` // 7-days sparkline
}

type tokenMarketData struct {
	tokenPriceData

	High24h             float64 `json:"high24h"`
	Low24h              float64 `json:"low24h"`
	Ath                 float64 `json:"ath"`
	AthChangePercentage float64 `json:"athChangePercentage"`
	AthDate             int64   `json:"athDate"`
	Atl                 float64 `json:"atl"`
	AtlChangePercentage float64 `json:"atlChangePercentage"`
	AtlDate             int64   `json:"atlDate"`
}

type socialData struct {
	Homepage          string `json:"homepage"`
	TwitterScreenName string `json:"twitterScreenName"`
	Discord           string `json:"discord"`
	Telegram          string `json:"telegram"`
	Twitter           string `json:"twitter"`
}

type TokenOverview struct {
	Token
	Markets map[string]tokenPriceData `json:"markets"`
}

type TokenDetails struct {
	Token
	Description string                     `json:"description"`
	Links       socialData                 `json:"links"`
	Markets     map[string]tokenMarketData `json:"markets"`
}
