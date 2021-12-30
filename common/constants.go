package common

import (
	"math/big"
	"time"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

var (
	// WETH, DAI, USDC, USDT, COMP, MKR
	UniswapQuoteSymbols = []string{"WETH", "UNIWETH", "DAI", "USDC", "USDT", "COMP", "MKR"}
	PancakeQuoteSymbols = []string{"WBNB", "BUSD", "USDT", "BTCB", "DAI", "CAKE", "ETH", "UST"}
	UniswapMaxHops      = 3
	BNBAddress          = etherCommon.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	ETHAddress          = etherCommon.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	MATICAddress        = etherCommon.HexToAddress("0xcccccccccccccccccccccccccccccccccccccccc")
	ZeroAddress         = etherCommon.HexToAddress("0x0000000000000000000000000000000000000000")
	AVAXAddress         = etherCommon.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	LegacyNativeTokens = []etherCommon.Address{
		BNBAddress,
		MATICAddress,
		AVAXAddress,
	}
	// TODO: Remove when contract updated
	FeePrioritizedTokens = []etherCommon.Address{
		etherCommon.HexToAddress("0x4f3afec4e5a3f2a6a1a411def7d7dfe50ee057bf"),
	}

	TimeoutDefault        = 10 * time.Second
	BPS                   = big.NewInt(10000)
	DefaultFeeBPS         = big.NewInt(8)
	FeeFromSource   uint8 = 0
	FeeFromDest     uint8 = 1
	FeeFromPlatform uint8 = 2

	LendingPlatformAaveV1   uint8 = 0
	LendingPlatformAaveV2   uint8 = 1
	LendingPlatformCompound uint8 = 2
)

// TODO: Remove when ready
func ConvertLegacyNative(address etherCommon.Address) etherCommon.Address {
	for _, addr := range LegacyNativeTokens {
		if address == addr {
			return ETHAddress
		}
	}

	return address
}
