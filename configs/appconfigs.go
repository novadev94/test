package configs

import (
	"github.com/chikob3/test/service/token"
	"strings"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

type AppConfig struct {
	// Maximum block range to query for logs
	GetLogsBlockRange int64 `json:"getLogsBlockRange"`
	BlockConfirmation int64 `json:"blockConfirmation"`

	Multicall   etherCommon.Address `json:"multicall"`
	NativeToken token.Token         `json:"nativeToken"`
	WrapToken   token.Token         `json:"wrapToken"`

	SmartWallet etherCommon.Address `json:"smartWallet"`
	// Start and End block of SmartWallet contract. empty/0 means hasn't started / hasn't ended
	SmartWalletStartBlock int64 `json:"smartWalletStartBlock"`
	SmartWalletEndBlock   int64 `json:"smartWalletEndBlock"`

	LendingContracts []struct {
		Name    string              `json:"name"`
		Address etherCommon.Address `json:"address"`
	} `json:"lendingContracts"`
	QuoteSymbols    []string `json:"quoteSymbols"`
	Chain           string   `json:"chain"`
	ChainID         int      `json:"chainID"`
	KrystalSubgraph string   `json:"krystalSubgraph"`
	CoinGecko       struct {
		Platform string `json:"platform"`
		NativeId string `json:"nativeId"`
	} `json:"coingecko"`

	SmartWalletV1 etherCommon.Address `json:"smartWalletV1"`
	// Start and End block of SmartWalletV1 contract. empty/0 means hasn't started / hasn't ended
	SmartWalletV1StartBlock int64 `json:"smartWalletV1StartBlock"`
	SmartWalletV1EndBlock   int64 `json:"smartWalletV1EndBlock"`

	UniConfigsV1 []struct {
		Name      string              `json:"name"`
		ShortName string              `json:"shortName"`
		Icon      string              `json:"icon"`
		Router    etherCommon.Address `json:"router"`
		Factory   etherCommon.Address `json:"factory"`
	} `json:"uniConfigsV1"`
	KyberConfigsV1 []struct {
		Name            string              `json:"name"`
		ShortName       string              `json:"shortName"`
		Icon            string              `json:"icon"`
		Proxy           etherCommon.Address `json:"proxy"`
		SwapEndpoint    string              `json:"swapEndpoint"`
		NetworkEndpoint string              `json:"networkEndpoint"`
	} `json:"kyberConfigsV1"`
	CompoundConfigs []struct {
		Name            string              `json:"name"`
		CompoundLens    etherCommon.Address `json:"compoundLens"`
		Comptroller     etherCommon.Address `json:"comptroller"`
		CComp           etherCommon.Address `json:"cComp"`
		CompToken       token.Token         `json:"compToken"`
		CNative         token.Token         `json:"cNative"`
		BlockSpeed      float64             `json:"blockSpeed"`
		CTokenPrefix    string              `json:"cTokenPrefix"`
		DistSpeedFn     string              `json:"distSpeedFn"`
		ComptrollerABI  string              `json:"comptrollerABI"`
		CompoundLensABI string              `json:"compoundLensABI"`
		BalanceExtFn    string              `json:"balanceExtFn"`
		ClaimFn         string              `json:"claimFn"`
	} `json:"compoundConfigs"`
	AaveConfigs []struct {
		Name                 string              `json:"name"`
		Version              uint8               `json:"version"`
		LendingPool          etherCommon.Address `json:"lendingPool"`
		DataProvider         etherCommon.Address `json:"dataProvider"`
		PriceOracle          etherCommon.Address `json:"priceOracle"`
		FetchAaveDataWrapper etherCommon.Address `json:"fetchAaveDataWrapper"`
		WethGateway          etherCommon.Address `json:"wethGateway"`
	} `json:"aaveConfigs"`
	Pricer []struct {
		Name     string                 `json:"name"`
		Type     string                 `json:"type"`
		Weight   float64                `json:"weight"`
		Fallback bool                   `json:"fallback"`
		Extra    map[string]interface{} `json:"extra"`
	} `json:"pricer"`
	ExchangeConfigs []struct {
		Name                          string              `json:"name"`
		ShortName                     string              `json:"shortName"`
		Icon                          string              `json:"icon"`
		Type                          string              `json:"type"`
		SwapContract                  etherCommon.Address `json:"swapContract"`
		Factory                       etherCommon.Address `json:"factory"`
		Router                        etherCommon.Address `json:"router"`
		KyberStorage                  etherCommon.Address `json:"kyberStorage"`
		AggregationExecutor           etherCommon.Address `json:"aggregationExecutor"`
		DefaultSwapGasLimit           uint64              `json:"defaultSwapGasLimit"`
		DefaultSwapAndDepositGasLimit uint64              `json:"defaultSwapAndDepositGasLimit"`
		DefaultSwapAndRepayGasLimit   uint64              `json:"defaultSwapAndRepayGasLimit"`
		Disabled                      bool                `json:"disabled"`
	}
	// Mapping known addresses to name, for easier visualization
	// addresses should be all lowercased, for easier access
	KnownAddresses            map[string]string `json:"knownAddresses"`
	NodeBalanceProviderConfig struct {
		FurthestBlock   int64 `json:"furthestBlock"`
		BlockRangeLimit int64 `json:"blockRangeLimit"`
	} `json:"nodeBalanceProviderConfig"`

	KrystalClaimConfig struct {
		Contract    etherCommon.Address    `json:"contract"`
		RewardTypes map[string]token.Token `json:"rewardTypes"`
	} `json:"krystalClaim" mapstructure:"krystalClaim"`
	KrystalClaimStartBlock int64 `json:"krystalClaimStartBlock"`
	GasConfig              struct {
		// this value will be used in gas estimator
		// to suggest a maxGasFee value which can be
		// usable in the next n blocks
		// (applied for EIP-1559 only)
		NextBlocks int `json:"nextBlocks"`
		CacheTime  int `json:"cacheTime"`
	} `json:"gasConfig"`
	RewardConfigs []struct {
		RewardID     string  `json:"rewardID"`
		AccVolume    float64 `json:"accVolume"`
		RewardVolume float64 `json:"rewardVolume"`
	} `json:"rewardConfigs"`
	SfNodeID        int64 `json:"sfNodeID"`
	GenerateRewards bool  `json:"generateRewards"`
}

var DefaultAppConfig = AppConfig{
	SmartWalletStartBlock:   0,
	SmartWalletEndBlock:     0,
	SmartWalletV1StartBlock: 0,
	SmartWalletV1EndBlock:   0,
	GetLogsBlockRange:       1000,
	BlockConfirmation:       12,
	KnownAddresses:          make(map[string]string),
	KrystalClaimStartBlock:  0,
}

func (c *AppConfig) GetAddressName(address string) string {
	res := c.KnownAddresses[strings.ToLower(address)]
	if res == "" {
		return address
	} else {
		return res
	}
}
