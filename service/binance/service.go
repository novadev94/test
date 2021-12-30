package binance

import (
	"fmt"
	"github.com/chikob3/test/common"
	"github.com/chikob3/test/database"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Service struct {
	db *gorm.DB
}

const symbolsUrl = "https://www.binance.com/bapi/margin/v1/friendly/margin/symbols"

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

type SymbolDataResp struct {
	Id            string `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
	Status        string `json:"status"`
	DelistedTime  string `json:"delistedTime"`
}

type SymbolsResp struct {
	Code          string           `json:"code"`
	Message       string           `json:"message"`
	MessageDetail string           `json:"messageDetail"`
	Data          []SymbolDataResp `json:"data"`
}

type TokenPrice struct {
	TokenName   string `json:"tokenName"`
	TimeRecord  string `json:"timeRecord"`
	RecordAt    string `json:"recordAt"`
	MarketPrice string `json:"marketPrice"`
	Volume      string `json:"volume"`
	OpenPrice   string `json:"openPrice"`
	ClosePrice  string `json:"closePrice"`
	HighPrice   string `json:"highPrice"`
	LowPrice    string `json:"lowPrice"`
	CandleSize  string `json:"candleSize"`
}

func (a *Service) GenData(timeframe string) error {
	var symbolsResp SymbolsResp
	params := map[string]string{}
	err := common.MakeGetRequest(symbolsUrl, params, 5*time.Second, &symbolsResp)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return err
	}

	listSymbols := make([]string, 0)
	for _, val := range symbolsResp.Data {
		listSymbols = append(listSymbols, val.Symbol)
	}
	var tokenPrices []TokenPrice
	for _, symbol := range listSymbols {
		if !strings.HasSuffix(symbol, "USDT") {
			continue
		}
		url := "https://www.binance.com/api/v3/klines"
		params := map[string]string{
			"limit":    "1",
			"symbol":   symbol,
			"interval": timeframe,
		}
		var klinesResp [][]interface{}
		err := common.MakeGetRequest(url, params, 5*time.Second, &klinesResp)
		if err != nil {
			continue
		}
		tokenPrice := TokenPrice{
			TokenName:   symbol,
			OpenPrice:   fmt.Sprintf("%v", klinesResp[0][1]),
			HighPrice:   fmt.Sprintf("%v", klinesResp[0][2]),
			LowPrice:    fmt.Sprintf("%v", klinesResp[0][3]),
			ClosePrice:  fmt.Sprintf("%v", klinesResp[0][4]),
			MarketPrice: "1.2",
			Volume:      "3341",
			CandleSize:  timeframe,
		}
		tokenPrices = append(tokenPrices, tokenPrice)
	}
	return a.SaveTokenPrices(tokenPrices)
}

func toTokenPriceModel(price TokenPrice) database.TokenPriceModel {
	return database.TokenPriceModel{
		TokenName:   price.TokenName,
		TimeRecord:  price.TimeRecord,
		RecordAt:    price.RecordAt,
		MarketPrice: price.MarketPrice,
		Volume:      price.Volume,
		OpenPrice:   price.OpenPrice,
		ClosePrice:  price.ClosePrice,
		HighPrice:   price.HighPrice,
		LowPrice:    price.LowPrice,
		CandleSize:  price.CandleSize,
	}
}

func (a *Service) SaveTokenPrices(tokenPrices []TokenPrice) error {
	var models []database.TokenPriceModel
	for _, tokenPrice := range tokenPrices {
		models = append(models, toTokenPriceModel(tokenPrice))
	}
	return a.db.Create(&models).Error
}
