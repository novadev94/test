package lending

import (
	"context"
	"github.com/chikob3/test/libs/contracts"
	"github.com/chikob3/test/service/token"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

const (
	aaveV1                 = 1
	aaveV2                 = 2
	nodeTimeout            = 10 * time.Second
	defaultAaveWithdrawGas = 2000000
	borrowRateModeStable   = 1
	borrowRateModeVariable = 2
)

var (
	rayMantissa = big.NewFloat(1e27)
	zeroBig     = big.NewInt(0)
)

type AaveClient struct {
	sugar *zap.SugaredLogger

	aaveWrapperContract *contracts.FetchAaveDataWrapper
	ethClient           *ethclient.Client
	version             uint8

	lendingPool  etherCommon.Address
	dataProvider etherCommon.Address

	priceOracleContract *contracts.AavePriceOracle
	priceOracleAddr     etherCommon.Address

	tokenMapping map[etherCommon.Address]token.Token // Mapping from underlying reserve to aToken
	wrapper      token.Wrapper

	wethGateway       etherCommon.Address
	tokenDataProvider token.TokenDataProvider
}

type AaveConfig struct {
	LendingPool          etherCommon.Address
	DataProvider         etherCommon.Address
	PriceOracle          etherCommon.Address
	FetchAaveDataWrapper etherCommon.Address
	WethGateway          etherCommon.Address // V2 Only
}

func NewAave(sugar *zap.SugaredLogger, version uint8, ethClient *ethclient.Client, tokenDataProvider token.TokenDataProvider, wrapper token.Wrapper, config AaveConfig) (*AaveClient, error) {
	client := &AaveClient{
		sugar:             sugar,
		ethClient:         ethClient,
		version:           version,
		lendingPool:       config.LendingPool,
		dataProvider:      config.DataProvider,
		priceOracleAddr:   config.PriceOracle,
		tokenDataProvider: tokenDataProvider,
		tokenMapping:      map[etherCommon.Address]token.Token{},
		wrapper:           wrapper,
		wethGateway:       config.WethGateway,
	}

	var err error
	client.aaveWrapperContract, err = contracts.NewFetchAaveDataWrapper(config.FetchAaveDataWrapper, ethClient)
	if err != nil {
		return nil, err
	}
	client.priceOracleContract, err = contracts.NewAavePriceOracle(client.priceOracleAddr, ethClient)
	if err != nil {
		return nil, err
	}
	client.fetchTokenMapping()

	return client, nil
}

func (c *AaveClient) GetID() string {
	if c.version == 1 {
		return AaveV1ClientId
	}

	return AaveV2ClientId
}

func (c *AaveClient) GetOverviewData() (map[etherCommon.Address]TokenOverview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), nodeTimeout)
	defer cancel()

	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	isV1 := c.version == aaveV1
	supportedReserves, err := c.aaveWrapperContract.GetReserves(callOpts, c.lendingPool, isV1)
	if err != nil {
		return nil, err
	}

	reserveData, err := c.aaveWrapperContract.GetReservesData(callOpts, c.dataProvider, isV1, supportedReserves)
	if err != nil {
		return nil, err
	}

	result := make(map[etherCommon.Address]TokenOverview)
	for i, data := range reserveData {
		addr := c.wrapper.Unwrap(supportedReserves[i])
		apyF := new(big.Float).SetInt(data.LiquidityRate)
		vBorrowRateF := new(big.Float).SetInt(data.VariableBorrowRate)
		sBorrowRateF := new(big.Float).SetInt(data.StableBorrowRate)

		apy, _ := new(big.Float).Quo(apyF, rayMantissa).Float64()
		vBorrowRate, _ := new(big.Float).Quo(vBorrowRateF, rayMantissa).Float64()
		sBorrowRate, _ := new(big.Float).Quo(sBorrowRateF, rayMantissa).Float64()
		result[addr] = TokenOverview{
			SupplyRate:         apy,
			StableBorrowRate:   sBorrowRate,
			VariableBorrowRate: vBorrowRate,
			TotalSupply:        data.TotalLiquidity,
			Liquidity:          data.AvailableLiquidity,
		}
	}

	return result, err
}

func (c *AaveClient) GetBearingTokens() ([]token.Token, error) {
	var result []token.Token
	for _, token := range c.tokenMapping {
		result = append(result, token)
	}
	return result, nil
}

func (c *AaveClient) fetchTokenMapping() error {
	ctx := context.Background()
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	isV1 := c.version == aaveV1
	supportedReserves, err := c.aaveWrapperContract.GetReserves(callOpts, c.lendingPool, isV1)
	if err != nil {
		return err
	}

	if isV1 {
		return c.fetchV1TokenMapping(ctx, supportedReserves)
	}
	return c.fetchV2TokenMapping(ctx, supportedReserves)
}

func (c *AaveClient) fetchV1TokenMapping(ctx context.Context, reserves []etherCommon.Address) error {
	contract, err := contracts.NewAaveV1LendingPool(c.lendingPool, c.ethClient)
	if err != nil {
		return err
	}

	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	for _, asset := range reserves {
		data, err := contract.GetReserveData(callOpts, asset)
		if err != nil {
			return err
		}

		token := c.tokenDataProvider.GetTokenByAddress(data.ATokenAddress)

		c.tokenMapping[asset] = token
	}

	return nil
}

func (c *AaveClient) fetchV2TokenMapping(ctx context.Context, reserves []etherCommon.Address) error {
	contract, err := contracts.NewAaveV2ProtocolDataProvider(c.dataProvider, c.ethClient)
	if err != nil {
		return err
	}

	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	for _, asset := range reserves {
		data, err := contract.GetReserveTokensAddresses(callOpts, asset)
		if err != nil {
			return err
		}

		asset = c.wrapper.Unwrap(asset)
		aToken := c.tokenDataProvider.GetTokenByAddress(data.ATokenAddress)

		c.tokenMapping[asset] = aToken
	}

	return nil
}

type LendingAaveInterestRate struct {
	LendingName              string  `json:"lendingName"`
	PoolName                 string  `json:"poolName"`
	LendingRate              float64 `json:"lendingRate"`
	BorrowingRate            float64 `json:"borrowingRate"`
	LendingDistributionApy   float64 `json:"lendingDistributionApy"`
	BorrowingDistributionApy float64 `json:"borrowingDistributionApy"`
}
