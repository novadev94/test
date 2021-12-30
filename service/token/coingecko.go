package token

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chikob3/test/common"
	"github.com/chikob3/test/libs/contracts"
	"github.com/chikob3/test/service/cache"
	"github.com/chikob3/test/service/multicall"
	"github.com/chikob3/test/service/sync"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const (
	coinGeckoApiEndpoint            = "https://api.coingecko.com/api/v3"
	coinGeckoApiTimeoutLong         = 30 * time.Second
	coinGeckoApiTimeoutShort        = 5 * time.Second
	coinGeckoDetailsCacheExpiration = 30 * time.Minute
	coinGeckoMarketRefreshInterval  = 5 * time.Minute
	concurrentLimit                 = 1
	telegramLink                    = "https://t.me"
	twitterLink                     = "https://twitter.com"
)

type CoinGeckoClient struct {
	sugar       *zap.SugaredLogger
	cache       cache.Cache
	platform    string
	nativeId    string
	nativeToken Token
	mu          *sync.ShardedRWMutex

	batchLimit int
	multicall  *multicall.Client

	cgkIdToQuoteSymbol map[string]string
	addrToCgkId        map[etherCommon.Address]string
	addrToToken        map[etherCommon.Address]Token
	addrToPrice        map[etherCommon.Address]tokenPriceData
	quoteToPrice       map[string]tokenPriceData
}

func newCgkClient(
	sugar *zap.SugaredLogger,
	cache cache.Cache,
	platform string,
	nativeId string,
	nativeToken Token,
	multicall *multicall.Client,
) *CoinGeckoClient {
	return &CoinGeckoClient{
		mu:                 sync.NewShardedRWMutex(),
		sugar:              sugar,
		cache:              cache,
		platform:           platform,
		nativeId:           nativeId,
		nativeToken:        nativeToken,
		batchLimit:         200,
		cgkIdToQuoteSymbol: map[string]string{"ethereum": "ETH", "binancecoin": "BNB", "matic-network": "MATIC", "bitcoin": "BTC", "avalanche-2": "AVAX"},
		multicall:          multicall,
		addrToCgkId:        map[etherCommon.Address]string{},
		addrToToken:        map[etherCommon.Address]Token{},
		addrToPrice:        map[etherCommon.Address]tokenPriceData{},
		quoteToPrice:       map[string]tokenPriceData{},
	}
}

func NewCoinGeckoClient(
	sugar *zap.SugaredLogger,
	cache cache.Cache,
	platform string,
	nativeId string,
	nativeToken Token,
	multicall *multicall.Client,
) (*CoinGeckoClient, error) {
	client := newCgkClient(sugar, cache, platform, nativeId, nativeToken, multicall)

	go client.fetchAll(nil)

	return client, nil
}

func (c *CoinGeckoClient) GetTokens() []Token {
	tokens := make([]Token, 0)
	addressToToken := c.addrToToken
	for _, token := range addressToToken {
		tokens = append(tokens, token)
	}

	sort.Slice(tokens, func(i, j int) bool {
		if tokens[i].Address == c.nativeToken.Address {
			return true
		}
		if tokens[j].Address == c.nativeToken.Address {
			return false
		}
		return strings.Compare(tokens[i].Name, tokens[j].Name) == -1
	})
	return tokens
}

func (c *CoinGeckoClient) GetTokenBySymbol(symbol string) Token {
	addressToToken := c.addrToToken
	for _, token := range addressToToken {
		if strings.EqualFold(token.Symbol, symbol) {
			return token
		}
	}
	return Token{}
}

func (c *CoinGeckoClient) GetTokenByAddress(address etherCommon.Address) Token {
	addressToToken := c.addrToToken
	return addressToToken[address]
}

func (c *CoinGeckoClient) IsFeeWhitelist(address etherCommon.Address) bool {
	return address == c.nativeToken.Address
}

func (c *CoinGeckoClient) GetTokenOverview(addresses ...etherCommon.Address) ([]TokenOverview, error) {
	overview := make([]TokenOverview, 0)
	addressToPrice := c.addrToPrice
	addressToToken := c.addrToToken
	quoteToPrice := c.quoteToPrice
	for _, address := range addresses {
		dataInUsd, ok := addressToPrice[address]
		if !ok {
			continue
		}
		token := addressToToken[address]
		if !ok {
			continue
		}

		tokenOverview := TokenOverview{
			Token:   token,
			Markets: map[string]tokenPriceData{},
		}
		tokenOverview.Markets["usd"] = dataInUsd
		for _, quoteInUsd := range quoteToPrice {
			dataInQuote := c.convertUsdToQuote(dataInUsd, quoteInUsd)
			tokenOverview.Markets[strings.ToLower(dataInQuote.Symbol)] = dataInQuote
		}
		overview = append(overview, tokenOverview)
	}

	return overview, nil
}

func (c *CoinGeckoClient) GetTokenDetails(address etherCommon.Address) (TokenDetails, error) {
	key := fmt.Sprintf("token_details_%s", address.Hex())

	c.mu.RLock(key)
	cached, ok := c.cache.Get(key)
	c.mu.RUnlock(key)
	if ok {
		return c.fillMarketData(cached.(TokenDetails)), nil
	}

	c.mu.Lock(key)
	defer c.mu.Unlock(key)

	cached, ok = c.cache.Get(key)
	if ok {
		return c.fillMarketData(cached.(TokenDetails)), nil
	}

	addressToCgkId := c.addrToCgkId
	cgkId, ok := addressToCgkId[address]
	if !ok {
		return TokenDetails{}, fmt.Errorf("No data")
	}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", cgkId)
	query := map[string]string{
		"tickers":        "false",
		"market_data":    "true",
		"community_data": "false",
		"developer_data": "false",
	}

	var response struct {
		Name  string `json:"name"`
		Image struct {
			Thumb string `json:"thumb"`
			Small string `json:"small"`
			Large string `json:"large"`
		} `json:"image"`
		Description map[string]string `json:"description"`
		Links       struct {
			Homepage                  []string `json:"homepage"`
			TwitterScreenName         string   `json:"twitter_screen_name"`
			ChatUrl                   []string `json:"chat_url"`
			TelegramChannelIdentifier string   `json:"telegram_channel_identifier"`
		} `json:"links"`
		MarketData struct {
			High24h             map[string]float64   `json:"high_24h"`
			Low24h              map[string]float64   `json:"low_24h"`
			Ath                 map[string]float64   `json:"ath"`
			AthChangePercentage map[string]float64   `json:"ath_change_percentage"`
			AthDate             map[string]time.Time `json:"ath_date"`
			Atl                 map[string]float64   `json:"atl"`
			AtlChangePercentage map[string]float64   `json:"atl_change_percentage"`
			AtlDate             map[string]time.Time `json:"atl_date"`
		} `json:"market_data"`
	}

	err := common.MakeGetRequest(url, query, 5*time.Second, &response)
	if err != nil {
		return TokenDetails{}, err
	}

	link := socialData{}
	if response.Links.TwitterScreenName != "" {
		link.TwitterScreenName = response.Links.TwitterScreenName
		link.Twitter = fmt.Sprintf("%s/%s", twitterLink, response.Links.TwitterScreenName)
	}
	if len(response.Links.Homepage) > 0 {
		link.Homepage = response.Links.Homepage[0]
	}
	for _, chatUrl := range response.Links.ChatUrl {
		if strings.Contains(chatUrl, "discord") {
			link.Discord = chatUrl
			break
		}
	}
	link.Telegram = fmt.Sprintf("%s/%s", telegramLink, response.Links.TelegramChannelIdentifier)

	addressToToken := c.addrToToken
	token := addressToToken[address]
	result := TokenDetails{
		Token:       token,
		Description: response.Description["en"],
		Links:       link,
		Markets:     map[string]tokenMarketData{},
	}
	marketData := response.MarketData
	supportedQuotes := []string{"usd"}
	for _, supportedQuote := range c.cgkIdToQuoteSymbol {
		quote := strings.ToLower(supportedQuote)
		supportedQuotes = append(supportedQuotes, quote)
	}
	for _, quote := range supportedQuotes {
		result.Markets[quote] = tokenMarketData{
			High24h:             marketData.High24h[quote],
			Low24h:              marketData.Low24h[quote],
			Ath:                 marketData.Ath[quote],
			AthChangePercentage: marketData.AthChangePercentage[quote],
			AthDate:             marketData.AthDate[quote].Unix(),
			Atl:                 marketData.Atl[quote],
			AtlChangePercentage: marketData.AtlChangePercentage[quote],
			AtlDate:             marketData.AtlDate[quote].Unix(),
		}
	}

	c.cache.Set(key, result, coinGeckoDetailsCacheExpiration)
	return c.fillMarketData(result), nil
}

func (c *CoinGeckoClient) GetSupportedSymbolPrices() ([]tokenPriceData, error) {
	result := make([]tokenPriceData, 0)
	quoteToPrice := c.quoteToPrice
	for _, priceData := range quoteToPrice {
		result = append(result, priceData)
	}
	return result, nil
}

func (c *CoinGeckoClient) convertUsdToQuote(dataInUsd tokenPriceData, quoteInUsd tokenPriceData) tokenPriceData {
	priceChange1hPercentage := (1+dataInUsd.PriceChange1hPercentage*0.01)/(1+quoteInUsd.PriceChange1hPercentage*0.01) - 1
	priceChange24hPercentage := (1+dataInUsd.PriceChange24hPercentage*0.01)/(1+quoteInUsd.PriceChange24hPercentage*0.01) - 1
	priceChange7dPercentage := (1+dataInUsd.PriceChange7dPercentage*0.01)/(1+quoteInUsd.PriceChange7dPercentage*0.01) - 1
	priceChange30dPercentage := (1+dataInUsd.PriceChange30dPercentage*0.01)/(1+quoteInUsd.PriceChange30dPercentage*0.01) - 1
	priceChange200dPercentage := (1+dataInUsd.PriceChange200dPercentage*0.01)/(1+quoteInUsd.PriceChange200dPercentage*0.01) - 1
	priceChange1yPercentage := (1+dataInUsd.PriceChange1yPercentage*0.01)/(1+quoteInUsd.PriceChange1yPercentage*0.01) - 1

	price := dataInUsd.Price / quoteInUsd.Price
	priceChange24h := price - (price / (1 + priceChange24hPercentage))

	marketCap := dataInUsd.MarketCap / quoteInUsd.Price
	marketCapChange24hPercentage := (1+dataInUsd.MarketCapChange24hPercentage*0.01)/(1+quoteInUsd.MarketCapChange24hPercentage*0.01) - 1
	marketCapChange24h := marketCap - (marketCap / (1 + priceChange24hPercentage))

	dataInQuote := dataInUsd
	dataInQuote.Symbol = quoteInUsd.Symbol
	dataInQuote.Price = price
	dataInQuote.PriceChange24h = priceChange24h
	dataInQuote.PriceChange1hPercentage = priceChange1hPercentage * 100
	dataInQuote.PriceChange24hPercentage = priceChange24hPercentage * 100
	dataInQuote.PriceChange7dPercentage = priceChange7dPercentage * 100
	dataInQuote.PriceChange30dPercentage = priceChange30dPercentage * 100
	dataInQuote.PriceChange200dPercentage = priceChange200dPercentage * 100
	dataInQuote.PriceChange1yPercentage = priceChange1yPercentage * 100

	dataInQuote.MarketCap = dataInUsd.MarketCap / quoteInUsd.Price
	dataInQuote.MarketCapChange24hPercentage = marketCapChange24hPercentage * 100
	dataInQuote.MarketCapChange24h = marketCapChange24h
	dataInQuote.TotalVolume = dataInUsd.TotalVolume / quoteInUsd.Price
	dataInQuote.Sparkline = []float64{}

	return dataInQuote
}

func (c *CoinGeckoClient) fillMarketData(tokenDetail TokenDetails) TokenDetails {
	addressToPrice := c.addrToPrice
	quoteToPrice := c.quoteToPrice
	dataInUsd, ok := addressToPrice[tokenDetail.Address]
	if !ok {
		return tokenDetail
	}

	for _, quoteSymbol := range c.cgkIdToQuoteSymbol {
		quoteInUsd, ok := quoteToPrice[quoteSymbol]
		if !ok {
			continue
		}
		symbol := strings.ToLower(quoteSymbol)
		dataInQuote := c.convertUsdToQuote(dataInUsd, quoteInUsd)
		market := tokenDetail.Markets[symbol]
		market.tokenPriceData = dataInQuote
		tokenDetail.Markets[symbol] = market
	}
	dataInUsd.Sparkline = []float64{}
	market := tokenDetail.Markets["usd"]
	market.tokenPriceData = dataInUsd
	tokenDetail.Markets["usd"] = market

	return tokenDetail
}

type marketResponse []struct {
	Id                       string  `json:"id"`
	Symbol                   string  `json:"symbol"`
	CurrentPrice             float64 `json:"current_price"`
	MarketCap                float64 `json:"market_cap"`
	TotalVolume              float64 `json:"total_volume"`
	PriceChange24h           float64 `json:"price_change_24h"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	SparklineIn7d            struct {
		Price []float64 `json:"price"`
	} `json:"sparkline_in_7d"`

	PriceChange1hPercentage   float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChange24hPercentage  float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChange7dPercentage   float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChange30dPercentage  float64 `json:"price_change_percentage_30d_in_currency"`
	PriceChange200dPercentage float64 `json:"price_change_percentage_200d_in_currency"`
	PriceChange1yPercentage   float64 `json:"price_change_percentage_1y_in_currency"`

	MarketCapChange24h           float64 `json:"market_cap_change_24h"`
	MarketCapChange24hPercentage float64 `json:"market_cap_change_percentage_24h"`
}

func (c *CoinGeckoClient) fetchMarketData(cgkIds []string) (marketResponse, error) {
	url := fmt.Sprintf("%s/coins/markets", coinGeckoApiEndpoint)
	queries := map[string]string{
		"ids":                     strings.Join(cgkIds, ","),
		"vs_currency":             "usd",
		"sparkline":               "true",
		"order":                   "market_cap_desc",
		"price_change_percentage": "1h,24h,7d,30d,200d,1y",
	}
	var response marketResponse
	err := common.MakeGetRequest(url, queries, coinGeckoApiTimeoutLong, &response)
	if err != nil {
		c.sugar.Debugw("Error get coingecko market overview data", "err", err)
		return nil, err
	}

	return response, nil
}

func (c *CoinGeckoClient) fetchAllMarketData() error {
	rateLimiter := make(chan interface{}, concurrentLimit)
	addressToCgkId := c.addrToCgkId
	addressToToken := c.addrToToken

	type AsyncResult struct {
		result marketResponse
		err    error
	}

	reqFn := func(ctx context.Context, ch chan AsyncResult, cgkIds []string) {
		rateLimiter <- nil
		defer func() {
			<-rateLimiter
		}()
		if ctx.Err() != nil {
			return
		}
		result, err := c.fetchMarketData(cgkIds)

		ch <- AsyncResult{
			result: result,
			err:    err,
		}
	}

	var (
		ids         = make([]string, 0)
		idToToken   = make(map[string]Token)
		resultC     = make([]chan AsyncResult, 0)
		count       = 0
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()
	for address, id := range addressToCgkId {
		token, ok := addressToToken[address]
		if !ok {
			continue
		}
		ids = append(ids, id)
		idToToken[id] = token

		count += 1
		if count != c.batchLimit {
			continue
		}

		rc := make(chan AsyncResult, 1)
		resultC = append(resultC, rc)
		go reqFn(ctx, rc, ids)
		ids = make([]string, 0)
		count = 0
	}
	for id := range c.cgkIdToQuoteSymbol {
		ids = append(ids, id)
	}

	// Last batch
	rc := make(chan AsyncResult, 1)
	resultC = append(resultC, rc)
	go reqFn(ctx, rc, ids)

	result := make(marketResponse, 0)
	for _, ch := range resultC {
		asyncResult := <-ch
		if asyncResult.err != nil {
			c.sugar.Debugw("Error getting market overview", "err", asyncResult.err)
			// Safe to returns here as asyncResult channel is buffered
			return asyncResult.err
		}

		result = append(result, asyncResult.result...)
	}

	mapping := make(map[etherCommon.Address]tokenPriceData)
	quoteToPrice := make(map[string]tokenPriceData)
	for _, mData := range result {
		tokenMarket := tokenPriceData{
			Symbol:                       strings.ToUpper(mData.Symbol),
			Price:                        mData.CurrentPrice,
			MarketCap:                    mData.MarketCap,
			TotalVolume:                  mData.TotalVolume,
			PriceChange24h:               mData.PriceChange24h,
			PriceChange1hPercentage:      mData.PriceChange1hPercentage,
			PriceChange24hPercentage:     mData.PriceChange24hPercentage,
			PriceChange7dPercentage:      mData.PriceChange7dPercentage,
			PriceChange30dPercentage:     mData.PriceChange30dPercentage,
			PriceChange200dPercentage:    mData.PriceChange200dPercentage,
			PriceChange1yPercentage:      mData.PriceChange1yPercentage,
			MarketCapChange24h:           mData.MarketCapChange24h,
			MarketCapChange24hPercentage: mData.MarketCapChange24hPercentage,
			Sparkline:                    mData.SparklineIn7d.Price,
		}
		token, ok := idToToken[mData.Id]
		if ok {
			mapping[token.Address] = tokenMarket
		}
		_, ok = c.cgkIdToQuoteSymbol[mData.Id]
		if ok {
			quoteToPrice[tokenMarket.Symbol] = tokenMarket
		}
	}

	c.quoteToPrice = quoteToPrice
	c.addrToPrice = mapping
	return nil
}

func (c *CoinGeckoClient) fetchTokenList() error {
	erc20Contract, _ := abi.JSON(strings.NewReader(contracts.Erc20ABI))
	var nameCalls []multicall.CallInput
	var symbolCalls []multicall.CallInput
	var decimalsCalls []multicall.CallInput
	var addresses []etherCommon.Address

	addressToCgkId := c.addrToCgkId
	for address := range addressToCgkId {
		switch common.ConvertLegacyNative(address) {
		case common.ZeroAddress, common.ETHAddress:
			continue
		}
		nameCall, _ := multicall.Encode(address, erc20Contract.Methods["name"])
		nameCalls = append(nameCalls, nameCall)
		symbolCall, _ := multicall.Encode(address, erc20Contract.Methods["symbol"])
		symbolCalls = append(symbolCalls, symbolCall)
		decimalCall, _ := multicall.Encode(address, erc20Contract.Methods["decimals"])
		decimalsCalls = append(decimalsCalls, decimalCall)
		addresses = append(addresses, address)
	}

	namesOutput, err := c.multicall.Do(nameCalls)
	if err != nil {
		return err
	}
	symbolsOutput, err := c.multicall.Do(symbolCalls)
	if err != nil {
		return err
	}
	decimalsOutput, err := c.multicall.Do(decimalsCalls)
	if err != nil {
		return err
	}

	trimFn := func(r rune) bool {
		return !unicode.IsPrint(r)
	}
	_ = trimFn
	mapping := make(map[etherCommon.Address]Token)
	for i, nameOut := range namesOutput.Results {
		symbolOut := symbolsOutput.Results[i]
		decimalsOut := decimalsOutput.Results[i]
		if !nameOut.Success ||
			!symbolOut.Success ||
			!decimalsOut.Success {
			continue
		}

		address := addresses[i]
		var err error
		var name string
		err = erc20Contract.Unpack(&name, "name", nameOut.ReturnData)
		if err != nil {
			name = string(bytes.TrimFunc(nameOut.ReturnData, trimFn))
		}
		var symbol string
		err = erc20Contract.Unpack(&symbol, "symbol", symbolOut.ReturnData)
		if err != nil {
			symbol = string(bytes.TrimFunc(symbolOut.ReturnData, trimFn))
		}
		var decimals uint8
		err = erc20Contract.Unpack(&decimals, "decimals", decimalsOut.ReturnData)
		if err != nil {
			continue
		}

		mapping[address] = Token{
			Address:  address,
			Symbol:   symbol,
			Name:     name,
			Decimals: int(decimals),
		}
	}

	mapping[c.nativeToken.Address] = c.nativeToken
	c.addrToToken = mapping
	return nil
}

func (c *CoinGeckoClient) fetchCgkIds() error {
	url := "https://api.coingecko.com/api/v3/coins/list?include_platform=true"
	var response []struct {
		Id        string            `json:"id"`
		Symbol    string            `json:"symbol"`
		Name      string            `json:"name"`
		Platforms map[string]string `json:"platforms"`
	}

	err := common.MakeGetRequest(url, nil, coinGeckoApiTimeoutLong, &response)
	if err != nil {
		return err
	}

	mapping := make(map[etherCommon.Address]string)
	for _, token := range response {
		addr, ok := token.Platforms[c.platform]
		if !ok {
			continue
		}
		// Prevent native id override,
		// Because matic has its own ERC20 token at 0x0000000000000000000000000000000000001010 to represent native
		// This will override our own implementation of nativeToken
		if token.Id == c.nativeId {
			continue
		}

		mapping[etherCommon.HexToAddress(addr)] = token.Id
	}
	mapping[c.nativeToken.Address] = c.nativeId

	c.addrToCgkId = mapping
	return nil
}

func (c *CoinGeckoClient) fetchAll(ready chan interface{}) {
	lastLongFetch := time.Unix(0, 0)
	failCount := 0
	firstTime := true

	for {
		if time.Since(lastLongFetch) > 30*time.Minute {
			err := c.fetchCgkIds()
			if err != nil {
				failCount += 1
				time.Sleep(time.Duration(5*failCount) * time.Second)
				continue
			}
			err = c.fetchTokenList()
			if err != nil {
				failCount += 1
				time.Sleep(time.Duration(5*failCount) * time.Second)
				continue
			}
			lastLongFetch = time.Now()
		}
		err := c.fetchAllMarketData()
		if err != nil {
			failCount += 1
			time.Sleep(time.Duration(5*failCount) * time.Second)
			continue
		}
		if ready != nil && firstTime {
			firstTime = false
			ready <- nil
		}
		failCount = 0
		time.Sleep(coinGeckoMarketRefreshInterval)
	}
}
