package token

import (
	"math/big"
	"time"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

const (
	nodeTimeout = 5 * time.Second
)

type DecimalProvider interface {
	GetDecimal(erc20Addr etherCommon.Address) int64

	// Return calculated decimal (ie 8 => 1e8, 18 => 1e18)
	GetMultiplier(erc20Addr etherCommon.Address) *big.Int
}
