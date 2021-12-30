// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VenusUnitrollerABI is the input ABI used to generate the binding from.
const VenusUnitrollerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"ActionProtocolPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusVAIMintIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIMinterVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIVaultVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"}],\"name\":\"NewTreasuryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"}],\"name\":\"NewTreasuryGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTreasuryPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"oldVAIController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"newVAIController\",\"type\":\"address\"}],\"name\":\"NewVAIController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVAIMintRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"NewVAIMintRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseInterval_\",\"type\":\"uint256\"}],\"name\":\"NewVAIVaultInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIVaultRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIVaultRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIVaultRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"VenusSpeedUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setProtocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"_setTreasuryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"vaiController_\",\"type\":\"address\"}],\"name\":\"_setVAIController\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"_setVAIMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReleaseAmount_\",\"type\":\"uint256\"}],\"name\":\"_setVAIVaultInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"venusSpeed\",\"type\":\"uint256\"}],\"name\":\"_setVenusSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIVaultRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIVaultRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"}],\"name\":\"distributeVAIMinterVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getXVSAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateVAICalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVenus\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minReleaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintedVAIs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repayVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMintedVAIOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiController\",\"outputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIVaultRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VenusUnitroller is an auto generated Go binding around an Ethereum contract.
type VenusUnitroller struct {
	VenusUnitrollerCaller     // Read-only binding to the contract
	VenusUnitrollerTransactor // Write-only binding to the contract
	VenusUnitrollerFilterer   // Log filterer for contract events
}

// VenusUnitrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusUnitrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusUnitrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusUnitrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusUnitrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusUnitrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusUnitrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusUnitrollerSession struct {
	Contract     *VenusUnitroller  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VenusUnitrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusUnitrollerCallerSession struct {
	Contract *VenusUnitrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VenusUnitrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusUnitrollerTransactorSession struct {
	Contract     *VenusUnitrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VenusUnitrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusUnitrollerRaw struct {
	Contract *VenusUnitroller // Generic contract binding to access the raw methods on
}

// VenusUnitrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusUnitrollerCallerRaw struct {
	Contract *VenusUnitrollerCaller // Generic read-only contract binding to access the raw methods on
}

// VenusUnitrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusUnitrollerTransactorRaw struct {
	Contract *VenusUnitrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusUnitroller creates a new instance of VenusUnitroller, bound to a specific deployed contract.
func NewVenusUnitroller(address common.Address, backend bind.ContractBackend) (*VenusUnitroller, error) {
	contract, err := bindVenusUnitroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VenusUnitroller{VenusUnitrollerCaller: VenusUnitrollerCaller{contract: contract}, VenusUnitrollerTransactor: VenusUnitrollerTransactor{contract: contract}, VenusUnitrollerFilterer: VenusUnitrollerFilterer{contract: contract}}, nil
}

// NewVenusUnitrollerCaller creates a new read-only instance of VenusUnitroller, bound to a specific deployed contract.
func NewVenusUnitrollerCaller(address common.Address, caller bind.ContractCaller) (*VenusUnitrollerCaller, error) {
	contract, err := bindVenusUnitroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerCaller{contract: contract}, nil
}

// NewVenusUnitrollerTransactor creates a new write-only instance of VenusUnitroller, bound to a specific deployed contract.
func NewVenusUnitrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusUnitrollerTransactor, error) {
	contract, err := bindVenusUnitroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerTransactor{contract: contract}, nil
}

// NewVenusUnitrollerFilterer creates a new log filterer instance of VenusUnitroller, bound to a specific deployed contract.
func NewVenusUnitrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusUnitrollerFilterer, error) {
	contract, err := bindVenusUnitroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerFilterer{contract: contract}, nil
}

// bindVenusUnitroller binds a generic wrapper to an already deployed contract.
func bindVenusUnitroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VenusUnitrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusUnitroller *VenusUnitrollerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VenusUnitroller.Contract.VenusUnitrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusUnitroller *VenusUnitrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.VenusUnitrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusUnitroller *VenusUnitrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.VenusUnitrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusUnitroller *VenusUnitrollerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VenusUnitroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusUnitroller *VenusUnitrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusUnitroller *VenusUnitrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.contract.Transact(opts, method, params...)
}

// FuncBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) FuncBorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "_borrowGuardianPaused")
	return *ret0, err
}

// FuncBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) FuncBorrowGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.FuncBorrowGuardianPaused(&_VenusUnitroller.CallOpts)
}

// FuncBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) FuncBorrowGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.FuncBorrowGuardianPaused(&_VenusUnitroller.CallOpts)
}

// FuncMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) FuncMintGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "_mintGuardianPaused")
	return *ret0, err
}

// FuncMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) FuncMintGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.FuncMintGuardianPaused(&_VenusUnitroller.CallOpts)
}

// FuncMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) FuncMintGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.FuncMintGuardianPaused(&_VenusUnitroller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "accountAssets", arg0, arg1)
	return *ret0, err
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _VenusUnitroller.Contract.AccountAssets(&_VenusUnitroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _VenusUnitroller.Contract.AccountAssets(&_VenusUnitroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) Admin() (common.Address, error) {
	return _VenusUnitroller.Contract.Admin(&_VenusUnitroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) Admin() (common.Address, error) {
	return _VenusUnitroller.Contract.Admin(&_VenusUnitroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "allMarkets", arg0)
	return *ret0, err
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _VenusUnitroller.Contract.AllMarkets(&_VenusUnitroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _VenusUnitroller.Contract.AllMarkets(&_VenusUnitroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "borrowCapGuardian")
	return *ret0, err
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _VenusUnitroller.Contract.BorrowCapGuardian(&_VenusUnitroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _VenusUnitroller.Contract.BorrowCapGuardian(&_VenusUnitroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "borrowCaps", arg0)
	return *ret0, err
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.BorrowCaps(&_VenusUnitroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.BorrowCaps(&_VenusUnitroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "borrowGuardianPaused", arg0)
	return *ret0, err
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _VenusUnitroller.Contract.BorrowGuardianPaused(&_VenusUnitroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _VenusUnitroller.Contract.BorrowGuardianPaused(&_VenusUnitroller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, vToken common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "checkMembership", account, vToken)
	return *ret0, err
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _VenusUnitroller.Contract.CheckMembership(&_VenusUnitroller.CallOpts, account, vToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _VenusUnitroller.Contract.CheckMembership(&_VenusUnitroller.CallOpts, account, vToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "closeFactorMantissa")
	return *ret0, err
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _VenusUnitroller.Contract.CloseFactorMantissa(&_VenusUnitroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _VenusUnitroller.Contract.CloseFactorMantissa(&_VenusUnitroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "comptrollerImplementation")
	return *ret0, err
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _VenusUnitroller.Contract.ComptrollerImplementation(&_VenusUnitroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _VenusUnitroller.Contract.ComptrollerImplementation(&_VenusUnitroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _VenusUnitroller.contract.Call(opts, out, "getAccountLiquidity", account)
	return *ret0, *ret1, *ret2, err
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.GetAccountLiquidity(&_VenusUnitroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.GetAccountLiquidity(&_VenusUnitroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_VenusUnitroller *VenusUnitrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "getAllMarkets")
	return *ret0, err
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_VenusUnitroller *VenusUnitrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _VenusUnitroller.Contract.GetAllMarkets(&_VenusUnitroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_VenusUnitroller *VenusUnitrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _VenusUnitroller.Contract.GetAllMarkets(&_VenusUnitroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_VenusUnitroller *VenusUnitrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "getAssetsIn", account)
	return *ret0, err
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_VenusUnitroller *VenusUnitrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _VenusUnitroller.Contract.GetAssetsIn(&_VenusUnitroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_VenusUnitroller *VenusUnitrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _VenusUnitroller.Contract.GetAssetsIn(&_VenusUnitroller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "getBlockNumber")
	return *ret0, err
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) GetBlockNumber() (*big.Int, error) {
	return _VenusUnitroller.Contract.GetBlockNumber(&_VenusUnitroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _VenusUnitroller.Contract.GetBlockNumber(&_VenusUnitroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _VenusUnitroller.contract.Call(opts, out, "getHypotheticalAccountLiquidity", account, vTokenModify, redeemTokens, borrowAmount)
	return *ret0, *ret1, *ret2, err
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.GetHypotheticalAccountLiquidity(&_VenusUnitroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.GetHypotheticalAccountLiquidity(&_VenusUnitroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) GetXVSAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "getXVSAddress")
	return *ret0, err
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) GetXVSAddress() (common.Address, error) {
	return _VenusUnitroller.Contract.GetXVSAddress(&_VenusUnitroller.CallOpts)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) GetXVSAddress() (common.Address, error) {
	return _VenusUnitroller.Contract.GetXVSAddress(&_VenusUnitroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "isComptroller")
	return *ret0, err
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) IsComptroller() (bool, error) {
	return _VenusUnitroller.Contract.IsComptroller(&_VenusUnitroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) IsComptroller() (bool, error) {
	return _VenusUnitroller.Contract.IsComptroller(&_VenusUnitroller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _VenusUnitroller.contract.Call(opts, out, "liquidateCalculateSeizeTokens", vTokenBorrowed, vTokenCollateral, actualRepayAmount)
	return *ret0, *ret1, err
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.LiquidateCalculateSeizeTokens(&_VenusUnitroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.LiquidateCalculateSeizeTokens(&_VenusUnitroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) LiquidateVAICalculateSeizeTokens(opts *bind.CallOpts, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _VenusUnitroller.contract.Call(opts, out, "liquidateVAICalculateSeizeTokens", vTokenCollateral, actualRepayAmount)
	return *ret0, *ret1, err
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.LiquidateVAICalculateSeizeTokens(&_VenusUnitroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusUnitroller.Contract.LiquidateVAICalculateSeizeTokens(&_VenusUnitroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "liquidationIncentiveMantissa")
	return *ret0, err
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _VenusUnitroller.Contract.LiquidationIncentiveMantissa(&_VenusUnitroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _VenusUnitroller.Contract.LiquidationIncentiveMantissa(&_VenusUnitroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_VenusUnitroller *VenusUnitrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	ret := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsVenus                  bool
	})
	out := ret
	err := _VenusUnitroller.contract.Call(opts, out, "markets", arg0)
	return *ret, err
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_VenusUnitroller *VenusUnitrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _VenusUnitroller.Contract.Markets(&_VenusUnitroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_VenusUnitroller *VenusUnitrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _VenusUnitroller.Contract.Markets(&_VenusUnitroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "maxAssets")
	return *ret0, err
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) MaxAssets() (*big.Int, error) {
	return _VenusUnitroller.Contract.MaxAssets(&_VenusUnitroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _VenusUnitroller.Contract.MaxAssets(&_VenusUnitroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) MinReleaseAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "minReleaseAmount")
	return *ret0, err
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) MinReleaseAmount() (*big.Int, error) {
	return _VenusUnitroller.Contract.MinReleaseAmount(&_VenusUnitroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) MinReleaseAmount() (*big.Int, error) {
	return _VenusUnitroller.Contract.MinReleaseAmount(&_VenusUnitroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "mintGuardianPaused", arg0)
	return *ret0, err
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _VenusUnitroller.Contract.MintGuardianPaused(&_VenusUnitroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _VenusUnitroller.Contract.MintGuardianPaused(&_VenusUnitroller.CallOpts, arg0)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) MintVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "mintVAIGuardianPaused")
	return *ret0, err
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) MintVAIGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.MintVAIGuardianPaused(&_VenusUnitroller.CallOpts)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) MintVAIGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.MintVAIGuardianPaused(&_VenusUnitroller.CallOpts)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) MintedVAIs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "mintedVAIs", arg0)
	return *ret0, err
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.MintedVAIs(&_VenusUnitroller.CallOpts, arg0)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.MintedVAIs(&_VenusUnitroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) Oracle() (common.Address, error) {
	return _VenusUnitroller.Contract.Oracle(&_VenusUnitroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) Oracle() (common.Address, error) {
	return _VenusUnitroller.Contract.Oracle(&_VenusUnitroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "pauseGuardian")
	return *ret0, err
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) PauseGuardian() (common.Address, error) {
	return _VenusUnitroller.Contract.PauseGuardian(&_VenusUnitroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _VenusUnitroller.Contract.PauseGuardian(&_VenusUnitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) PendingAdmin() (common.Address, error) {
	return _VenusUnitroller.Contract.PendingAdmin(&_VenusUnitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _VenusUnitroller.Contract.PendingAdmin(&_VenusUnitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "pendingComptrollerImplementation")
	return *ret0, err
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _VenusUnitroller.Contract.PendingComptrollerImplementation(&_VenusUnitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _VenusUnitroller.Contract.PendingComptrollerImplementation(&_VenusUnitroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) ProtocolPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "protocolPaused")
	return *ret0, err
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) ProtocolPaused() (bool, error) {
	return _VenusUnitroller.Contract.ProtocolPaused(&_VenusUnitroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) ProtocolPaused() (bool, error) {
	return _VenusUnitroller.Contract.ProtocolPaused(&_VenusUnitroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) ReleaseStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "releaseStartBlock")
	return *ret0, err
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) ReleaseStartBlock() (*big.Int, error) {
	return _VenusUnitroller.Contract.ReleaseStartBlock(&_VenusUnitroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) ReleaseStartBlock() (*big.Int, error) {
	return _VenusUnitroller.Contract.ReleaseStartBlock(&_VenusUnitroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) RepayVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "repayVAIGuardianPaused")
	return *ret0, err
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) RepayVAIGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.RepayVAIGuardianPaused(&_VenusUnitroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) RepayVAIGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.RepayVAIGuardianPaused(&_VenusUnitroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "seizeGuardianPaused")
	return *ret0, err
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) SeizeGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.SeizeGuardianPaused(&_VenusUnitroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.SeizeGuardianPaused(&_VenusUnitroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "transferGuardianPaused")
	return *ret0, err
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) TransferGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.TransferGuardianPaused(&_VenusUnitroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_VenusUnitroller *VenusUnitrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _VenusUnitroller.Contract.TransferGuardianPaused(&_VenusUnitroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "treasuryAddress")
	return *ret0, err
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) TreasuryAddress() (common.Address, error) {
	return _VenusUnitroller.Contract.TreasuryAddress(&_VenusUnitroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) TreasuryAddress() (common.Address, error) {
	return _VenusUnitroller.Contract.TreasuryAddress(&_VenusUnitroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) TreasuryGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "treasuryGuardian")
	return *ret0, err
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) TreasuryGuardian() (common.Address, error) {
	return _VenusUnitroller.Contract.TreasuryGuardian(&_VenusUnitroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) TreasuryGuardian() (common.Address, error) {
	return _VenusUnitroller.Contract.TreasuryGuardian(&_VenusUnitroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "treasuryPercent")
	return *ret0, err
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) TreasuryPercent() (*big.Int, error) {
	return _VenusUnitroller.Contract.TreasuryPercent(&_VenusUnitroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) TreasuryPercent() (*big.Int, error) {
	return _VenusUnitroller.Contract.TreasuryPercent(&_VenusUnitroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) VaiController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "vaiController")
	return *ret0, err
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) VaiController() (common.Address, error) {
	return _VenusUnitroller.Contract.VaiController(&_VenusUnitroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VaiController() (common.Address, error) {
	return _VenusUnitroller.Contract.VaiController(&_VenusUnitroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VaiMintRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "vaiMintRate")
	return *ret0, err
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VaiMintRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VaiMintRate(&_VenusUnitroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VaiMintRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VaiMintRate(&_VenusUnitroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCaller) VaiVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "vaiVaultAddress")
	return *ret0, err
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerSession) VaiVaultAddress() (common.Address, error) {
	return _VenusUnitroller.Contract.VaiVaultAddress(&_VenusUnitroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VaiVaultAddress() (common.Address, error) {
	return _VenusUnitroller.Contract.VaiVaultAddress(&_VenusUnitroller.CallOpts)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusAccrued", arg0)
	return *ret0, err
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusAccrued(&_VenusUnitroller.CallOpts, arg0)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusAccrued(&_VenusUnitroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	ret := new(struct {
		Index *big.Int
		Block uint32
	})
	out := ret
	err := _VenusUnitroller.contract.Call(opts, out, "venusBorrowState", arg0)
	return *ret, err
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_VenusUnitroller *VenusUnitrollerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusUnitroller.Contract.VenusBorrowState(&_VenusUnitroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusUnitroller.Contract.VenusBorrowState(&_VenusUnitroller.CallOpts, arg0)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusBorrowerIndex", arg0, arg1)
	return *ret0, err
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusBorrowerIndex(&_VenusUnitroller.CallOpts, arg0, arg1)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusBorrowerIndex(&_VenusUnitroller.CallOpts, arg0, arg1)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusInitialIndex")
	return *ret0, err
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VenusUnitroller *VenusUnitrollerSession) VenusInitialIndex() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusInitialIndex(&_VenusUnitroller.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusInitialIndex() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusInitialIndex(&_VenusUnitroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusRate")
	return *ret0, err
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VenusRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusRate(&_VenusUnitroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusRate(&_VenusUnitroller.CallOpts)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusSpeeds", arg0)
	return *ret0, err
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusSpeeds(&_VenusUnitroller.CallOpts, arg0)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusSpeeds(&_VenusUnitroller.CallOpts, arg0)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusSupplierIndex", arg0, arg1)
	return *ret0, err
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusSupplierIndex(&_VenusUnitroller.CallOpts, arg0, arg1)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusSupplierIndex(&_VenusUnitroller.CallOpts, arg0, arg1)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	ret := new(struct {
		Index *big.Int
		Block uint32
	})
	out := ret
	err := _VenusUnitroller.contract.Call(opts, out, "venusSupplyState", arg0)
	return *ret, err
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_VenusUnitroller *VenusUnitrollerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusUnitroller.Contract.VenusSupplyState(&_VenusUnitroller.CallOpts, arg0)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusUnitroller.Contract.VenusSupplyState(&_VenusUnitroller.CallOpts, arg0)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusVAIRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusVAIRate")
	return *ret0, err
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VenusVAIRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusVAIRate(&_VenusUnitroller.CallOpts)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusVAIRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusVAIRate(&_VenusUnitroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCaller) VenusVAIVaultRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VenusUnitroller.contract.Call(opts, out, "venusVAIVaultRate")
	return *ret0, err
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusVAIVaultRate(&_VenusUnitroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_VenusUnitroller *VenusUnitrollerCallerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _VenusUnitroller.Contract.VenusVAIVaultRate(&_VenusUnitroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VenusUnitroller *VenusUnitrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.Become(&_VenusUnitroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.Become(&_VenusUnitroller.TransactOpts, unitroller)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_VenusUnitroller *VenusUnitrollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetBorrowCapGuardian(&_VenusUnitroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetBorrowCapGuardian(&_VenusUnitroller.TransactOpts, newBorrowCapGuardian)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetCloseFactor(&_VenusUnitroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetCloseFactor(&_VenusUnitroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setCollateralFactor", vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetCollateralFactor(&_VenusUnitroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetCollateralFactor(&_VenusUnitroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetLiquidationIncentive(&_VenusUnitroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetLiquidationIncentive(&_VenusUnitroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setMarketBorrowCaps", vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_VenusUnitroller *VenusUnitrollerSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetMarketBorrowCaps(&_VenusUnitroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetMarketBorrowCaps(&_VenusUnitroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetPriceOracle(&_VenusUnitroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetPriceOracle(&_VenusUnitroller.TransactOpts, newOracle)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetProtocolPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setProtocolPaused", state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_VenusUnitroller *VenusUnitrollerSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetProtocolPaused(&_VenusUnitroller.TransactOpts, state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetProtocolPaused(&_VenusUnitroller.TransactOpts, state)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetTreasuryData(opts *bind.TransactOpts, newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setTreasuryData", newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetTreasuryData(&_VenusUnitroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetTreasuryData(&_VenusUnitroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetVAIController(opts *bind.TransactOpts, vaiController_ common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setVAIController", vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVAIController(&_VenusUnitroller.TransactOpts, vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVAIController(&_VenusUnitroller.TransactOpts, vaiController_)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetVAIMintRate(opts *bind.TransactOpts, newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setVAIMintRate", newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVAIMintRate(&_VenusUnitroller.TransactOpts, newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVAIMintRate(&_VenusUnitroller.TransactOpts, newVAIMintRate)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) SetVAIVaultInfo(opts *bind.TransactOpts, vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setVAIVaultInfo", vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_VenusUnitroller *VenusUnitrollerSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVAIVaultInfo(&_VenusUnitroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVAIVaultInfo(&_VenusUnitroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) SetVenusSpeed(opts *bind.TransactOpts, vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setVenusSpeed", vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_VenusUnitroller *VenusUnitrollerSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVenusSpeed(&_VenusUnitroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVenusSpeed(&_VenusUnitroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) SetVenusVAIRate(opts *bind.TransactOpts, venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setVenusVAIRate", venusVAIRate_)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_VenusUnitroller *VenusUnitrollerSession) SetVenusVAIRate(venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVenusVAIRate(&_VenusUnitroller.TransactOpts, venusVAIRate_)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetVenusVAIRate(venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVenusVAIRate(&_VenusUnitroller.TransactOpts, venusVAIRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) SetVenusVAIVaultRate(opts *bind.TransactOpts, venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_setVenusVAIVaultRate", venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_VenusUnitroller *VenusUnitrollerSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVenusVAIVaultRate(&_VenusUnitroller.TransactOpts, venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetVenusVAIVaultRate(&_VenusUnitroller.TransactOpts, venusVAIVaultRate_)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SupportMarket(opts *bind.TransactOpts, vToken common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "_supportMarket", vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SupportMarket(&_VenusUnitroller.TransactOpts, vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SupportMarket(&_VenusUnitroller.TransactOpts, vToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "borrowAllowed", vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.BorrowAllowed(&_VenusUnitroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.BorrowAllowed(&_VenusUnitroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) BorrowVerify(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "borrowVerify", vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_VenusUnitroller *VenusUnitrollerSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.BorrowVerify(&_VenusUnitroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.BorrowVerify(&_VenusUnitroller.TransactOpts, vToken, borrower, borrowAmount)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) ClaimVenus(opts *bind.TransactOpts, holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "claimVenus", holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_VenusUnitroller *VenusUnitrollerSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ClaimVenus(&_VenusUnitroller.TransactOpts, holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ClaimVenus(&_VenusUnitroller.TransactOpts, holder, vTokens)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) ClaimVenus0(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "claimVenus0", holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_VenusUnitroller *VenusUnitrollerSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ClaimVenus0(&_VenusUnitroller.TransactOpts, holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ClaimVenus0(&_VenusUnitroller.TransactOpts, holder)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) ClaimVenus1(opts *bind.TransactOpts, holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "claimVenus1", holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_VenusUnitroller *VenusUnitrollerSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ClaimVenus1(&_VenusUnitroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ClaimVenus1(&_VenusUnitroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) DistributeVAIMinterVenus(opts *bind.TransactOpts, vaiMinter common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "distributeVAIMinterVenus", vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_VenusUnitroller *VenusUnitrollerSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.DistributeVAIMinterVenus(&_VenusUnitroller.TransactOpts, vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.DistributeVAIMinterVenus(&_VenusUnitroller.TransactOpts, vaiMinter)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_VenusUnitroller *VenusUnitrollerTransactor) EnterMarkets(opts *bind.TransactOpts, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "enterMarkets", vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_VenusUnitroller *VenusUnitrollerSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.EnterMarkets(&_VenusUnitroller.TransactOpts, vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_VenusUnitroller *VenusUnitrollerTransactorSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.EnterMarkets(&_VenusUnitroller.TransactOpts, vTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) ExitMarket(opts *bind.TransactOpts, vTokenAddress common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "exitMarket", vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ExitMarket(&_VenusUnitroller.TransactOpts, vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ExitMarket(&_VenusUnitroller.TransactOpts, vTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "liquidateBorrowAllowed", vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.LiquidateBorrowAllowed(&_VenusUnitroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.LiquidateBorrowAllowed(&_VenusUnitroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "liquidateBorrowVerify", vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_VenusUnitroller *VenusUnitrollerSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.LiquidateBorrowVerify(&_VenusUnitroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.LiquidateBorrowVerify(&_VenusUnitroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) MintAllowed(opts *bind.TransactOpts, vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "mintAllowed", vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.MintAllowed(&_VenusUnitroller.TransactOpts, vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.MintAllowed(&_VenusUnitroller.TransactOpts, vToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) MintVerify(opts *bind.TransactOpts, vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "mintVerify", vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_VenusUnitroller *VenusUnitrollerSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.MintVerify(&_VenusUnitroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.MintVerify(&_VenusUnitroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "redeemAllowed", vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RedeemAllowed(&_VenusUnitroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RedeemAllowed(&_VenusUnitroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) RedeemVerify(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "redeemVerify", vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_VenusUnitroller *VenusUnitrollerSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RedeemVerify(&_VenusUnitroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RedeemVerify(&_VenusUnitroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) ReleaseToVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "releaseToVault")
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_VenusUnitroller *VenusUnitrollerSession) ReleaseToVault() (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ReleaseToVault(&_VenusUnitroller.TransactOpts)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) ReleaseToVault() (*types.Transaction, error) {
	return _VenusUnitroller.Contract.ReleaseToVault(&_VenusUnitroller.TransactOpts)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "repayBorrowAllowed", vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RepayBorrowAllowed(&_VenusUnitroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RepayBorrowAllowed(&_VenusUnitroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "repayBorrowVerify", vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_VenusUnitroller *VenusUnitrollerSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RepayBorrowVerify(&_VenusUnitroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.RepayBorrowVerify(&_VenusUnitroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "seizeAllowed", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SeizeAllowed(&_VenusUnitroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SeizeAllowed(&_VenusUnitroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) SeizeVerify(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "seizeVerify", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_VenusUnitroller *VenusUnitrollerSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SeizeVerify(&_VenusUnitroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SeizeVerify(&_VenusUnitroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) SetMintedVAIOf(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "setMintedVAIOf", owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetMintedVAIOf(&_VenusUnitroller.TransactOpts, owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.SetMintedVAIOf(&_VenusUnitroller.TransactOpts, owner, amount)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactor) TransferAllowed(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "transferAllowed", vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.TransferAllowed(&_VenusUnitroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_VenusUnitroller *VenusUnitrollerTransactorSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.TransferAllowed(&_VenusUnitroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactor) TransferVerify(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.contract.Transact(opts, "transferVerify", vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_VenusUnitroller *VenusUnitrollerSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.TransferVerify(&_VenusUnitroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_VenusUnitroller *VenusUnitrollerTransactorSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusUnitroller.Contract.TransferVerify(&_VenusUnitroller.TransactOpts, vToken, src, dst, transferTokens)
}

// VenusUnitrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the VenusUnitroller contract.
type VenusUnitrollerActionPausedIterator struct {
	Event *VenusUnitrollerActionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerActionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerActionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerActionPaused represents a ActionPaused event raised by the VenusUnitroller contract.
type VenusUnitrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*VenusUnitrollerActionPausedIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerActionPausedIterator{contract: _VenusUnitroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerActionPaused)
				if err := _VenusUnitroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseActionPaused(log types.Log) (*VenusUnitrollerActionPaused, error) {
	event := new(VenusUnitrollerActionPaused)
	if err := _VenusUnitroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the VenusUnitroller contract.
type VenusUnitrollerActionPaused0Iterator struct {
	Event *VenusUnitrollerActionPaused0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerActionPaused0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerActionPaused0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerActionPaused0 represents a ActionPaused0 event raised by the VenusUnitroller contract.
type VenusUnitrollerActionPaused0 struct {
	VToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterActionPaused0(opts *bind.FilterOpts) (*VenusUnitrollerActionPaused0Iterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerActionPaused0Iterator{contract: _VenusUnitroller.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerActionPaused0) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerActionPaused0)
				if err := _VenusUnitroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseActionPaused0(log types.Log) (*VenusUnitrollerActionPaused0, error) {
	event := new(VenusUnitrollerActionPaused0)
	if err := _VenusUnitroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerActionProtocolPausedIterator is returned from FilterActionProtocolPaused and is used to iterate over the raw logs and unpacked data for ActionProtocolPaused events raised by the VenusUnitroller contract.
type VenusUnitrollerActionProtocolPausedIterator struct {
	Event *VenusUnitrollerActionProtocolPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerActionProtocolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerActionProtocolPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerActionProtocolPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerActionProtocolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerActionProtocolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerActionProtocolPaused represents a ActionProtocolPaused event raised by the VenusUnitroller contract.
type VenusUnitrollerActionProtocolPaused struct {
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterActionProtocolPaused is a free log retrieval operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterActionProtocolPaused(opts *bind.FilterOpts) (*VenusUnitrollerActionProtocolPausedIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerActionProtocolPausedIterator{contract: _VenusUnitroller.contract, event: "ActionProtocolPaused", logs: logs, sub: sub}, nil
}

// WatchActionProtocolPaused is a free log subscription operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchActionProtocolPaused(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerActionProtocolPaused) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerActionProtocolPaused)
				if err := _VenusUnitroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionProtocolPaused is a log parse operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseActionProtocolPaused(log types.Log) (*VenusUnitrollerActionProtocolPaused, error) {
	event := new(VenusUnitrollerActionProtocolPaused)
	if err := _VenusUnitroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerDistributedBorrowerVenusIterator is returned from FilterDistributedBorrowerVenus and is used to iterate over the raw logs and unpacked data for DistributedBorrowerVenus events raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedBorrowerVenusIterator struct {
	Event *VenusUnitrollerDistributedBorrowerVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerDistributedBorrowerVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerDistributedBorrowerVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerDistributedBorrowerVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerDistributedBorrowerVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerDistributedBorrowerVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerDistributedBorrowerVenus represents a DistributedBorrowerVenus event raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedBorrowerVenus struct {
	VToken           common.Address
	Borrower         common.Address
	VenusDelta       *big.Int
	VenusBorrowIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerVenus is a free log retrieval operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterDistributedBorrowerVenus(opts *bind.FilterOpts, vToken []common.Address, borrower []common.Address) (*VenusUnitrollerDistributedBorrowerVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerDistributedBorrowerVenusIterator{contract: _VenusUnitroller.contract, event: "DistributedBorrowerVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerVenus is a free log subscription operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchDistributedBorrowerVenus(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerDistributedBorrowerVenus, vToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerDistributedBorrowerVenus)
				if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedBorrowerVenus is a log parse operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseDistributedBorrowerVenus(log types.Log) (*VenusUnitrollerDistributedBorrowerVenus, error) {
	event := new(VenusUnitrollerDistributedBorrowerVenus)
	if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerDistributedSupplierVenusIterator is returned from FilterDistributedSupplierVenus and is used to iterate over the raw logs and unpacked data for DistributedSupplierVenus events raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedSupplierVenusIterator struct {
	Event *VenusUnitrollerDistributedSupplierVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerDistributedSupplierVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerDistributedSupplierVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerDistributedSupplierVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerDistributedSupplierVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerDistributedSupplierVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerDistributedSupplierVenus represents a DistributedSupplierVenus event raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedSupplierVenus struct {
	VToken           common.Address
	Supplier         common.Address
	VenusDelta       *big.Int
	VenusSupplyIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierVenus is a free log retrieval operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterDistributedSupplierVenus(opts *bind.FilterOpts, vToken []common.Address, supplier []common.Address) (*VenusUnitrollerDistributedSupplierVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerDistributedSupplierVenusIterator{contract: _VenusUnitroller.contract, event: "DistributedSupplierVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierVenus is a free log subscription operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchDistributedSupplierVenus(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerDistributedSupplierVenus, vToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerDistributedSupplierVenus)
				if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedSupplierVenus is a log parse operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseDistributedSupplierVenus(log types.Log) (*VenusUnitrollerDistributedSupplierVenus, error) {
	event := new(VenusUnitrollerDistributedSupplierVenus)
	if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerDistributedVAIMinterVenusIterator is returned from FilterDistributedVAIMinterVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIMinterVenus events raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedVAIMinterVenusIterator struct {
	Event *VenusUnitrollerDistributedVAIMinterVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerDistributedVAIMinterVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerDistributedVAIMinterVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerDistributedVAIMinterVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerDistributedVAIMinterVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerDistributedVAIMinterVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerDistributedVAIMinterVenus represents a DistributedVAIMinterVenus event raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedVAIMinterVenus struct {
	VaiMinter         common.Address
	VenusDelta        *big.Int
	VenusVAIMintIndex *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIMinterVenus is a free log retrieval operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterDistributedVAIMinterVenus(opts *bind.FilterOpts, vaiMinter []common.Address) (*VenusUnitrollerDistributedVAIMinterVenusIterator, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerDistributedVAIMinterVenusIterator{contract: _VenusUnitroller.contract, event: "DistributedVAIMinterVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIMinterVenus is a free log subscription operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchDistributedVAIMinterVenus(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerDistributedVAIMinterVenus, vaiMinter []common.Address) (event.Subscription, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerDistributedVAIMinterVenus)
				if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIMinterVenus is a log parse operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseDistributedVAIMinterVenus(log types.Log) (*VenusUnitrollerDistributedVAIMinterVenus, error) {
	event := new(VenusUnitrollerDistributedVAIMinterVenus)
	if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerDistributedVAIVaultVenusIterator is returned from FilterDistributedVAIVaultVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIVaultVenus events raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedVAIVaultVenusIterator struct {
	Event *VenusUnitrollerDistributedVAIVaultVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerDistributedVAIVaultVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerDistributedVAIVaultVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerDistributedVAIVaultVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerDistributedVAIVaultVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerDistributedVAIVaultVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerDistributedVAIVaultVenus represents a DistributedVAIVaultVenus event raised by the VenusUnitroller contract.
type VenusUnitrollerDistributedVAIVaultVenus struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIVaultVenus is a free log retrieval operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterDistributedVAIVaultVenus(opts *bind.FilterOpts) (*VenusUnitrollerDistributedVAIVaultVenusIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerDistributedVAIVaultVenusIterator{contract: _VenusUnitroller.contract, event: "DistributedVAIVaultVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIVaultVenus is a free log subscription operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchDistributedVAIVaultVenus(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerDistributedVAIVaultVenus) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerDistributedVAIVaultVenus)
				if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIVaultVenus is a log parse operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseDistributedVAIVaultVenus(log types.Log) (*VenusUnitrollerDistributedVAIVaultVenus, error) {
	event := new(VenusUnitrollerDistributedVAIVaultVenus)
	if err := _VenusUnitroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the VenusUnitroller contract.
type VenusUnitrollerFailureIterator struct {
	Event *VenusUnitrollerFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerFailure represents a Failure event raised by the VenusUnitroller contract.
type VenusUnitrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*VenusUnitrollerFailureIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerFailureIterator{contract: _VenusUnitroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerFailure) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerFailure)
				if err := _VenusUnitroller.contract.UnpackLog(event, "Failure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseFailure(log types.Log) (*VenusUnitrollerFailure, error) {
	event := new(VenusUnitrollerFailure)
	if err := _VenusUnitroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the VenusUnitroller contract.
type VenusUnitrollerMarketEnteredIterator struct {
	Event *VenusUnitrollerMarketEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerMarketEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerMarketEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerMarketEntered represents a MarketEntered event raised by the VenusUnitroller contract.
type VenusUnitrollerMarketEntered struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*VenusUnitrollerMarketEnteredIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerMarketEnteredIterator{contract: _VenusUnitroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerMarketEntered)
				if err := _VenusUnitroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseMarketEntered(log types.Log) (*VenusUnitrollerMarketEntered, error) {
	event := new(VenusUnitrollerMarketEntered)
	if err := _VenusUnitroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the VenusUnitroller contract.
type VenusUnitrollerMarketExitedIterator struct {
	Event *VenusUnitrollerMarketExited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerMarketExited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerMarketExited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerMarketExited represents a MarketExited event raised by the VenusUnitroller contract.
type VenusUnitrollerMarketExited struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*VenusUnitrollerMarketExitedIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerMarketExitedIterator{contract: _VenusUnitroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerMarketExited)
				if err := _VenusUnitroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseMarketExited(log types.Log) (*VenusUnitrollerMarketExited, error) {
	event := new(VenusUnitrollerMarketExited)
	if err := _VenusUnitroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the VenusUnitroller contract.
type VenusUnitrollerMarketListedIterator struct {
	Event *VenusUnitrollerMarketListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerMarketListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerMarketListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerMarketListed represents a MarketListed event raised by the VenusUnitroller contract.
type VenusUnitrollerMarketListed struct {
	VToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*VenusUnitrollerMarketListedIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerMarketListedIterator{contract: _VenusUnitroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerMarketListed)
				if err := _VenusUnitroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseMarketListed(log types.Log) (*VenusUnitrollerMarketListed, error) {
	event := new(VenusUnitrollerMarketListed)
	if err := _VenusUnitroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the VenusUnitroller contract.
type VenusUnitrollerNewBorrowCapIterator struct {
	Event *VenusUnitrollerNewBorrowCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewBorrowCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewBorrowCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewBorrowCap represents a NewBorrowCap event raised by the VenusUnitroller contract.
type VenusUnitrollerNewBorrowCap struct {
	VToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, vToken []common.Address) (*VenusUnitrollerNewBorrowCapIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewBorrowCapIterator{contract: _VenusUnitroller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewBorrowCap, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewBorrowCap)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewBorrowCap(log types.Log) (*VenusUnitrollerNewBorrowCap, error) {
	event := new(VenusUnitrollerNewBorrowCap)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the VenusUnitroller contract.
type VenusUnitrollerNewBorrowCapGuardianIterator struct {
	Event *VenusUnitrollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewBorrowCapGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewBorrowCapGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the VenusUnitroller contract.
type VenusUnitrollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*VenusUnitrollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewBorrowCapGuardianIterator{contract: _VenusUnitroller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewBorrowCapGuardian)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*VenusUnitrollerNewBorrowCapGuardian, error) {
	event := new(VenusUnitrollerNewBorrowCapGuardian)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the VenusUnitroller contract.
type VenusUnitrollerNewCloseFactorIterator struct {
	Event *VenusUnitrollerNewCloseFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewCloseFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewCloseFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewCloseFactor represents a NewCloseFactor event raised by the VenusUnitroller contract.
type VenusUnitrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*VenusUnitrollerNewCloseFactorIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewCloseFactorIterator{contract: _VenusUnitroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewCloseFactor)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewCloseFactor(log types.Log) (*VenusUnitrollerNewCloseFactor, error) {
	event := new(VenusUnitrollerNewCloseFactor)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the VenusUnitroller contract.
type VenusUnitrollerNewCollateralFactorIterator struct {
	Event *VenusUnitrollerNewCollateralFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewCollateralFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewCollateralFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewCollateralFactor represents a NewCollateralFactor event raised by the VenusUnitroller contract.
type VenusUnitrollerNewCollateralFactor struct {
	VToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*VenusUnitrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewCollateralFactorIterator{contract: _VenusUnitroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewCollateralFactor)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewCollateralFactor(log types.Log) (*VenusUnitrollerNewCollateralFactor, error) {
	event := new(VenusUnitrollerNewCollateralFactor)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the VenusUnitroller contract.
type VenusUnitrollerNewLiquidationIncentiveIterator struct {
	Event *VenusUnitrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewLiquidationIncentive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewLiquidationIncentive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the VenusUnitroller contract.
type VenusUnitrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*VenusUnitrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewLiquidationIncentiveIterator{contract: _VenusUnitroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewLiquidationIncentive)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*VenusUnitrollerNewLiquidationIncentive, error) {
	event := new(VenusUnitrollerNewLiquidationIncentive)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the VenusUnitroller contract.
type VenusUnitrollerNewPauseGuardianIterator struct {
	Event *VenusUnitrollerNewPauseGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewPauseGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewPauseGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewPauseGuardian represents a NewPauseGuardian event raised by the VenusUnitroller contract.
type VenusUnitrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*VenusUnitrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewPauseGuardianIterator{contract: _VenusUnitroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewPauseGuardian)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewPauseGuardian(log types.Log) (*VenusUnitrollerNewPauseGuardian, error) {
	event := new(VenusUnitrollerNewPauseGuardian)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the VenusUnitroller contract.
type VenusUnitrollerNewPriceOracleIterator struct {
	Event *VenusUnitrollerNewPriceOracle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewPriceOracle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewPriceOracle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewPriceOracle represents a NewPriceOracle event raised by the VenusUnitroller contract.
type VenusUnitrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*VenusUnitrollerNewPriceOracleIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewPriceOracleIterator{contract: _VenusUnitroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewPriceOracle)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewPriceOracle(log types.Log) (*VenusUnitrollerNewPriceOracle, error) {
	event := new(VenusUnitrollerNewPriceOracle)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewTreasuryAddressIterator is returned from FilterNewTreasuryAddress and is used to iterate over the raw logs and unpacked data for NewTreasuryAddress events raised by the VenusUnitroller contract.
type VenusUnitrollerNewTreasuryAddressIterator struct {
	Event *VenusUnitrollerNewTreasuryAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewTreasuryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewTreasuryAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewTreasuryAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewTreasuryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewTreasuryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewTreasuryAddress represents a NewTreasuryAddress event raised by the VenusUnitroller contract.
type VenusUnitrollerNewTreasuryAddress struct {
	OldTreasuryAddress common.Address
	NewTreasuryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryAddress is a free log retrieval operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewTreasuryAddress(opts *bind.FilterOpts) (*VenusUnitrollerNewTreasuryAddressIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewTreasuryAddressIterator{contract: _VenusUnitroller.contract, event: "NewTreasuryAddress", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryAddress is a free log subscription operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewTreasuryAddress(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewTreasuryAddress) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewTreasuryAddress)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryAddress is a log parse operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewTreasuryAddress(log types.Log) (*VenusUnitrollerNewTreasuryAddress, error) {
	event := new(VenusUnitrollerNewTreasuryAddress)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewTreasuryGuardianIterator is returned from FilterNewTreasuryGuardian and is used to iterate over the raw logs and unpacked data for NewTreasuryGuardian events raised by the VenusUnitroller contract.
type VenusUnitrollerNewTreasuryGuardianIterator struct {
	Event *VenusUnitrollerNewTreasuryGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewTreasuryGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewTreasuryGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewTreasuryGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewTreasuryGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewTreasuryGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewTreasuryGuardian represents a NewTreasuryGuardian event raised by the VenusUnitroller contract.
type VenusUnitrollerNewTreasuryGuardian struct {
	OldTreasuryGuardian common.Address
	NewTreasuryGuardian common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryGuardian is a free log retrieval operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewTreasuryGuardian(opts *bind.FilterOpts) (*VenusUnitrollerNewTreasuryGuardianIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewTreasuryGuardianIterator{contract: _VenusUnitroller.contract, event: "NewTreasuryGuardian", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryGuardian is a free log subscription operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewTreasuryGuardian(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewTreasuryGuardian) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewTreasuryGuardian)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryGuardian is a log parse operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewTreasuryGuardian(log types.Log) (*VenusUnitrollerNewTreasuryGuardian, error) {
	event := new(VenusUnitrollerNewTreasuryGuardian)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewTreasuryPercentIterator is returned from FilterNewTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewTreasuryPercent events raised by the VenusUnitroller contract.
type VenusUnitrollerNewTreasuryPercentIterator struct {
	Event *VenusUnitrollerNewTreasuryPercent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewTreasuryPercent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewTreasuryPercent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewTreasuryPercent represents a NewTreasuryPercent event raised by the VenusUnitroller contract.
type VenusUnitrollerNewTreasuryPercent struct {
	OldTreasuryPercent *big.Int
	NewTreasuryPercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryPercent is a free log retrieval operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewTreasuryPercent(opts *bind.FilterOpts) (*VenusUnitrollerNewTreasuryPercentIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewTreasuryPercentIterator{contract: _VenusUnitroller.contract, event: "NewTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryPercent is a free log subscription operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewTreasuryPercent(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewTreasuryPercent)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryPercent is a log parse operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewTreasuryPercent(log types.Log) (*VenusUnitrollerNewTreasuryPercent, error) {
	event := new(VenusUnitrollerNewTreasuryPercent)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewVAIControllerIterator is returned from FilterNewVAIController and is used to iterate over the raw logs and unpacked data for NewVAIController events raised by the VenusUnitroller contract.
type VenusUnitrollerNewVAIControllerIterator struct {
	Event *VenusUnitrollerNewVAIController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewVAIControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewVAIController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewVAIController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewVAIControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewVAIControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewVAIController represents a NewVAIController event raised by the VenusUnitroller contract.
type VenusUnitrollerNewVAIController struct {
	OldVAIController common.Address
	NewVAIController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewVAIController is a free log retrieval operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewVAIController(opts *bind.FilterOpts) (*VenusUnitrollerNewVAIControllerIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewVAIControllerIterator{contract: _VenusUnitroller.contract, event: "NewVAIController", logs: logs, sub: sub}, nil
}

// WatchNewVAIController is a free log subscription operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewVAIController(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewVAIController) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewVAIController)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIController is a log parse operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewVAIController(log types.Log) (*VenusUnitrollerNewVAIController, error) {
	event := new(VenusUnitrollerNewVAIController)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewVAIMintRateIterator is returned from FilterNewVAIMintRate and is used to iterate over the raw logs and unpacked data for NewVAIMintRate events raised by the VenusUnitroller contract.
type VenusUnitrollerNewVAIMintRateIterator struct {
	Event *VenusUnitrollerNewVAIMintRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewVAIMintRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewVAIMintRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewVAIMintRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewVAIMintRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewVAIMintRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewVAIMintRate represents a NewVAIMintRate event raised by the VenusUnitroller contract.
type VenusUnitrollerNewVAIMintRate struct {
	OldVAIMintRate *big.Int
	NewVAIMintRate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewVAIMintRate is a free log retrieval operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewVAIMintRate(opts *bind.FilterOpts) (*VenusUnitrollerNewVAIMintRateIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewVAIMintRateIterator{contract: _VenusUnitroller.contract, event: "NewVAIMintRate", logs: logs, sub: sub}, nil
}

// WatchNewVAIMintRate is a free log subscription operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewVAIMintRate(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewVAIMintRate) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewVAIMintRate)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIMintRate is a log parse operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewVAIMintRate(log types.Log) (*VenusUnitrollerNewVAIMintRate, error) {
	event := new(VenusUnitrollerNewVAIMintRate)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewVAIVaultInfoIterator is returned from FilterNewVAIVaultInfo and is used to iterate over the raw logs and unpacked data for NewVAIVaultInfo events raised by the VenusUnitroller contract.
type VenusUnitrollerNewVAIVaultInfoIterator struct {
	Event *VenusUnitrollerNewVAIVaultInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewVAIVaultInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewVAIVaultInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewVAIVaultInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewVAIVaultInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewVAIVaultInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewVAIVaultInfo represents a NewVAIVaultInfo event raised by the VenusUnitroller contract.
type VenusUnitrollerNewVAIVaultInfo struct {
	Vault             common.Address
	ReleaseStartBlock *big.Int
	ReleaseInterval   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewVAIVaultInfo is a free log retrieval operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewVAIVaultInfo(opts *bind.FilterOpts) (*VenusUnitrollerNewVAIVaultInfoIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewVAIVaultInfoIterator{contract: _VenusUnitroller.contract, event: "NewVAIVaultInfo", logs: logs, sub: sub}, nil
}

// WatchNewVAIVaultInfo is a free log subscription operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewVAIVaultInfo(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewVAIVaultInfo) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewVAIVaultInfo)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIVaultInfo is a log parse operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewVAIVaultInfo(log types.Log) (*VenusUnitrollerNewVAIVaultInfo, error) {
	event := new(VenusUnitrollerNewVAIVaultInfo)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewVenusVAIRateIterator is returned from FilterNewVenusVAIRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIRate events raised by the VenusUnitroller contract.
type VenusUnitrollerNewVenusVAIRateIterator struct {
	Event *VenusUnitrollerNewVenusVAIRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewVenusVAIRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewVenusVAIRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewVenusVAIRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewVenusVAIRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewVenusVAIRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewVenusVAIRate represents a NewVenusVAIRate event raised by the VenusUnitroller contract.
type VenusUnitrollerNewVenusVAIRate struct {
	OldVenusVAIRate *big.Int
	NewVenusVAIRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIRate is a free log retrieval operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewVenusVAIRate(opts *bind.FilterOpts) (*VenusUnitrollerNewVenusVAIRateIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewVenusVAIRateIterator{contract: _VenusUnitroller.contract, event: "NewVenusVAIRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIRate is a free log subscription operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewVenusVAIRate(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewVenusVAIRate) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewVenusVAIRate)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIRate is a log parse operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewVenusVAIRate(log types.Log) (*VenusUnitrollerNewVenusVAIRate, error) {
	event := new(VenusUnitrollerNewVenusVAIRate)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerNewVenusVAIVaultRateIterator is returned from FilterNewVenusVAIVaultRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIVaultRate events raised by the VenusUnitroller contract.
type VenusUnitrollerNewVenusVAIVaultRateIterator struct {
	Event *VenusUnitrollerNewVenusVAIVaultRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerNewVenusVAIVaultRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerNewVenusVAIVaultRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerNewVenusVAIVaultRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerNewVenusVAIVaultRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerNewVenusVAIVaultRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerNewVenusVAIVaultRate represents a NewVenusVAIVaultRate event raised by the VenusUnitroller contract.
type VenusUnitrollerNewVenusVAIVaultRate struct {
	OldVenusVAIVaultRate *big.Int
	NewVenusVAIVaultRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIVaultRate is a free log retrieval operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterNewVenusVAIVaultRate(opts *bind.FilterOpts) (*VenusUnitrollerNewVenusVAIVaultRateIterator, error) {

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerNewVenusVAIVaultRateIterator{contract: _VenusUnitroller.contract, event: "NewVenusVAIVaultRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIVaultRate is a free log subscription operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchNewVenusVAIVaultRate(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerNewVenusVAIVaultRate) (event.Subscription, error) {

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerNewVenusVAIVaultRate)
				if err := _VenusUnitroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIVaultRate is a log parse operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseNewVenusVAIVaultRate(log types.Log) (*VenusUnitrollerNewVenusVAIVaultRate, error) {
	event := new(VenusUnitrollerNewVenusVAIVaultRate)
	if err := _VenusUnitroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VenusUnitrollerVenusSpeedUpdatedIterator is returned from FilterVenusSpeedUpdated and is used to iterate over the raw logs and unpacked data for VenusSpeedUpdated events raised by the VenusUnitroller contract.
type VenusUnitrollerVenusSpeedUpdatedIterator struct {
	Event *VenusUnitrollerVenusSpeedUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenusUnitrollerVenusSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusUnitrollerVenusSpeedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenusUnitrollerVenusSpeedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenusUnitrollerVenusSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusUnitrollerVenusSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusUnitrollerVenusSpeedUpdated represents a VenusSpeedUpdated event raised by the VenusUnitroller contract.
type VenusUnitrollerVenusSpeedUpdated struct {
	VToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVenusSpeedUpdated is a free log retrieval operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_VenusUnitroller *VenusUnitrollerFilterer) FilterVenusSpeedUpdated(opts *bind.FilterOpts, vToken []common.Address) (*VenusUnitrollerVenusSpeedUpdatedIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusUnitroller.contract.FilterLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &VenusUnitrollerVenusSpeedUpdatedIterator{contract: _VenusUnitroller.contract, event: "VenusSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchVenusSpeedUpdated is a free log subscription operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_VenusUnitroller *VenusUnitrollerFilterer) WatchVenusSpeedUpdated(opts *bind.WatchOpts, sink chan<- *VenusUnitrollerVenusSpeedUpdated, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusUnitroller.contract.WatchLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusUnitrollerVenusSpeedUpdated)
				if err := _VenusUnitroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVenusSpeedUpdated is a log parse operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_VenusUnitroller *VenusUnitrollerFilterer) ParseVenusSpeedUpdated(log types.Log) (*VenusUnitrollerVenusSpeedUpdated, error) {
	event := new(VenusUnitrollerVenusSpeedUpdated)
	if err := _VenusUnitroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
