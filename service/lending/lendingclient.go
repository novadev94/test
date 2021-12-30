package lending

import (
	"fmt"
	"github.com/chikob3/test/service/cache"
	"github.com/chikob3/test/service/token"
	"math/big"
	"sync"
	"time"

	etherCommon "github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type TokenOverview struct {
	SupplyRate         float64
	StableBorrowRate   float64
	VariableBorrowRate float64

	DistributionSupplyRate float64
	DistributionBorrowRate float64

	TotalSupply *big.Int
	Liquidity   *big.Int
}

type TokenBalance struct {
	SupplyBalance         *big.Int
	StableBorrowBalance   *big.Int
	VariableBorrowBalance *big.Int

	InterestBearingToken        token.Token
	InterestBearingTokenBalance *big.Int
}

type Client interface {
	GetID() string
	GetOverviewData() (map[etherCommon.Address]TokenOverview, error)
	GetBearingTokens() ([]token.Token, error)
}

const (
	cacheTime = 1 * time.Minute
)

type CachedClient struct {
	sugar  *zap.SugaredLogger
	cache  cache.Cache
	mu     sync.RWMutex
	client Client
}

func NewCachedClient(sugar *zap.SugaredLogger, cache cache.Cache, client Client) *CachedClient {
	return &CachedClient{
		sugar:  sugar,
		cache:  cache,
		client: client,
	}
}

func (c *CachedClient) GetID() string {
	return c.client.GetID()
}

func (c *CachedClient) GetOverviewData() (map[etherCommon.Address]TokenOverview, error) {
	result := make(map[etherCommon.Address]TokenOverview)
	cacheKey := fmt.Sprintf("lending_%s", c.GetID())
	c.mu.RLock()
	cached, ok := c.cache.Get(cacheKey)
	c.mu.RUnlock()

	if ok {
		for key, value := range cached.(map[etherCommon.Address]TokenOverview) {
			result[key] = value
		}
		return result, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// It might be available in such short amount of time
	cached, ok = c.cache.Get(cacheKey)
	if ok {
		for key, value := range cached.(map[etherCommon.Address]TokenOverview) {
			result[key] = value
		}
		return result, nil
	}

	apy, err := c.client.GetOverviewData()
	if err != nil {
		return nil, err
	}

	copyOfApy := make(map[etherCommon.Address]TokenOverview)

	for key, value := range apy {
		copyOfApy[key] = value
	}

	err = c.cache.Set(cacheKey, copyOfApy, cacheTime)

	if err != nil {
		c.sugar.Debugw("Error writing apy data to cache", "err", err)
	}

	return apy, nil
}

func (c *CachedClient) GetBearingTokens() ([]token.Token, error) {
	return c.client.GetBearingTokens()
}
