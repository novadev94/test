package lending

import (
	"math/big"
)

const (
	AaveV1ClientId             = "AaveV1"
	AaveV2ClientId             = "AaveV2"
	CompoundClientId           = "Compound"
	VenusClientId              = "Venus"
	defaultCompoundWithdrawGas = 500000
	compTokenDecimal           = 18
)

var (
	cTokenMult    = big.NewInt(1e8)
	compTokenMult = big.NewFloat(1e18)
)
