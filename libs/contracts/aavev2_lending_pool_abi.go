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

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	VariableBorrowIndex         *big.Int
	CurrentLiquidityRate        *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	Id                          uint8
}

// DataTypesUserConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserConfigurationMap struct {
	Data *big.Int
}

// AaveV2LendingPoolABI is the input ABI used to generate the binding from.
const AaveV2LendingPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RebalanceStableBorrowRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TOTAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LENDINGPOOL_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUMBER_RESERVES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STABLE_RATE_BORROW_SIZE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"modes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressesProvider\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rebalanceStableBorrowRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rateStrategyAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateStrategyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AaveV2LendingPool is an auto generated Go binding around an Ethereum contract.
type AaveV2LendingPool struct {
	AaveV2LendingPoolCaller     // Read-only binding to the contract
	AaveV2LendingPoolTransactor // Write-only binding to the contract
	AaveV2LendingPoolFilterer   // Log filterer for contract events
}

// AaveV2LendingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV2LendingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2LendingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV2LendingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2LendingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV2LendingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2LendingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV2LendingPoolSession struct {
	Contract     *AaveV2LendingPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AaveV2LendingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV2LendingPoolCallerSession struct {
	Contract *AaveV2LendingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AaveV2LendingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV2LendingPoolTransactorSession struct {
	Contract     *AaveV2LendingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AaveV2LendingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV2LendingPoolRaw struct {
	Contract *AaveV2LendingPool // Generic contract binding to access the raw methods on
}

// AaveV2LendingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV2LendingPoolCallerRaw struct {
	Contract *AaveV2LendingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV2LendingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV2LendingPoolTransactorRaw struct {
	Contract *AaveV2LendingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV2LendingPool creates a new instance of AaveV2LendingPool, bound to a specific deployed contract.
func NewAaveV2LendingPool(address common.Address, backend bind.ContractBackend) (*AaveV2LendingPool, error) {
	contract, err := bindAaveV2LendingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPool{AaveV2LendingPoolCaller: AaveV2LendingPoolCaller{contract: contract}, AaveV2LendingPoolTransactor: AaveV2LendingPoolTransactor{contract: contract}, AaveV2LendingPoolFilterer: AaveV2LendingPoolFilterer{contract: contract}}, nil
}

// NewAaveV2LendingPoolCaller creates a new read-only instance of AaveV2LendingPool, bound to a specific deployed contract.
func NewAaveV2LendingPoolCaller(address common.Address, caller bind.ContractCaller) (*AaveV2LendingPoolCaller, error) {
	contract, err := bindAaveV2LendingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolCaller{contract: contract}, nil
}

// NewAaveV2LendingPoolTransactor creates a new write-only instance of AaveV2LendingPool, bound to a specific deployed contract.
func NewAaveV2LendingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV2LendingPoolTransactor, error) {
	contract, err := bindAaveV2LendingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolTransactor{contract: contract}, nil
}

// NewAaveV2LendingPoolFilterer creates a new log filterer instance of AaveV2LendingPool, bound to a specific deployed contract.
func NewAaveV2LendingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV2LendingPoolFilterer, error) {
	contract, err := bindAaveV2LendingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolFilterer{contract: contract}, nil
}

// bindAaveV2LendingPool binds a generic wrapper to an already deployed contract.
func bindAaveV2LendingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveV2LendingPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2LendingPool *AaveV2LendingPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AaveV2LendingPool.Contract.AaveV2LendingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2LendingPool *AaveV2LendingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.AaveV2LendingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2LendingPool *AaveV2LendingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.AaveV2LendingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2LendingPool *AaveV2LendingPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AaveV2LendingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.contract.Transact(opts, method, params...)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) FLASHLOANPREMIUMTOTAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "FLASHLOAN_PREMIUM_TOTAL")
	return *ret0, err
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.FLASHLOANPREMIUMTOTAL(&_AaveV2LendingPool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.FLASHLOANPREMIUMTOTAL(&_AaveV2LendingPool.CallOpts)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) LENDINGPOOLREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "LENDINGPOOL_REVISION")
	return *ret0, err
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.LENDINGPOOLREVISION(&_AaveV2LendingPool.CallOpts)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.LENDINGPOOLREVISION(&_AaveV2LendingPool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) MAXNUMBERRESERVES(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "MAX_NUMBER_RESERVES")
	return *ret0, err
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) MAXNUMBERRESERVES() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.MAXNUMBERRESERVES(&_AaveV2LendingPool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) MAXNUMBERRESERVES() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.MAXNUMBERRESERVES(&_AaveV2LendingPool.CallOpts)
}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) MAXSTABLERATEBORROWSIZEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "MAX_STABLE_RATE_BORROW_SIZE_PERCENT")
	return *ret0, err
}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) MAXSTABLERATEBORROWSIZEPERCENT() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.MAXSTABLERATEBORROWSIZEPERCENT(&_AaveV2LendingPool.CallOpts)
}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) MAXSTABLERATEBORROWSIZEPERCENT() (*big.Int, error) {
	return _AaveV2LendingPool.Contract.MAXSTABLERATEBORROWSIZEPERCENT(&_AaveV2LendingPool.CallOpts)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetAddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "getAddressesProvider")
	return *ret0, err
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetAddressesProvider() (common.Address, error) {
	return _AaveV2LendingPool.Contract.GetAddressesProvider(&_AaveV2LendingPool.CallOpts)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetAddressesProvider() (common.Address, error) {
	return _AaveV2LendingPool.Contract.GetAddressesProvider(&_AaveV2LendingPool.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var (
		ret0 = new(DataTypesReserveConfigurationMap)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "getConfiguration", asset)
	return *ret0, err
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AaveV2LendingPool.Contract.GetConfiguration(&_AaveV2LendingPool.CallOpts, asset)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AaveV2LendingPool.Contract.GetConfiguration(&_AaveV2LendingPool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var (
		ret0 = new(DataTypesReserveData)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "getReserveData", asset)
	return *ret0, err
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _AaveV2LendingPool.Contract.GetReserveData(&_AaveV2LendingPool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _AaveV2LendingPool.Contract.GetReserveData(&_AaveV2LendingPool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "getReserveNormalizedIncome", asset)
	return *ret0, err
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AaveV2LendingPool.Contract.GetReserveNormalizedIncome(&_AaveV2LendingPool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AaveV2LendingPool.Contract.GetReserveNormalizedIncome(&_AaveV2LendingPool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "getReserveNormalizedVariableDebt", asset)
	return *ret0, err
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AaveV2LendingPool.Contract.GetReserveNormalizedVariableDebt(&_AaveV2LendingPool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AaveV2LendingPool.Contract.GetReserveNormalizedVariableDebt(&_AaveV2LendingPool.CallOpts, asset)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "getReservesList")
	return *ret0, err
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetReservesList() ([]common.Address, error) {
	return _AaveV2LendingPool.Contract.GetReservesList(&_AaveV2LendingPool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetReservesList() ([]common.Address, error) {
	return _AaveV2LendingPool.Contract.GetReservesList(&_AaveV2LendingPool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetUserAccountData(opts *bind.CallOpts, user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	ret := new(struct {
		TotalCollateralETH          *big.Int
		TotalDebtETH                *big.Int
		AvailableBorrowsETH         *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	out := ret
	err := _AaveV2LendingPool.contract.Call(opts, out, "getUserAccountData", user)
	return *ret, err
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AaveV2LendingPool.Contract.GetUserAccountData(&_AaveV2LendingPool.CallOpts, user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AaveV2LendingPool.Contract.GetUserAccountData(&_AaveV2LendingPool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) GetUserConfiguration(opts *bind.CallOpts, user common.Address) (DataTypesUserConfigurationMap, error) {
	var (
		ret0 = new(DataTypesUserConfigurationMap)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "getUserConfiguration", user)
	return *ret0, err
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveV2LendingPool *AaveV2LendingPoolSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AaveV2LendingPool.Contract.GetUserConfiguration(&_AaveV2LendingPool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AaveV2LendingPool.Contract.GetUserConfiguration(&_AaveV2LendingPool.CallOpts, user)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AaveV2LendingPool *AaveV2LendingPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AaveV2LendingPool.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) Paused() (bool, error) {
	return _AaveV2LendingPool.Contract.Paused(&_AaveV2LendingPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AaveV2LendingPool *AaveV2LendingPoolCallerSession) Paused() (bool, error) {
	return _AaveV2LendingPool.Contract.Paused(&_AaveV2LendingPool.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Borrow(&_AaveV2LendingPool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Borrow(&_AaveV2LendingPool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Deposit(&_AaveV2LendingPool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Deposit(&_AaveV2LendingPool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) FinalizeTransfer(opts *bind.TransactOpts, asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "finalizeTransfer", asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.FinalizeTransfer(&_AaveV2LendingPool.TransactOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.FinalizeTransfer(&_AaveV2LendingPool.TransactOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) FlashLoan(opts *bind.TransactOpts, receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "flashLoan", receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.FlashLoan(&_AaveV2LendingPool.TransactOpts, receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.FlashLoan(&_AaveV2LendingPool.TransactOpts, receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "initReserve", asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) InitReserve(asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.InitReserve(&_AaveV2LendingPool.TransactOpts, asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) InitReserve(asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.InitReserve(&_AaveV2LendingPool.TransactOpts, asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) Initialize(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "initialize", provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Initialize(&_AaveV2LendingPool.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Initialize(&_AaveV2LendingPool.TransactOpts, provider)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) LiquidationCall(opts *bind.TransactOpts, collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "liquidationCall", collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.LiquidationCall(&_AaveV2LendingPool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.LiquidationCall(&_AaveV2LendingPool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) RebalanceStableBorrowRate(opts *bind.TransactOpts, asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "rebalanceStableBorrowRate", asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.RebalanceStableBorrowRate(&_AaveV2LendingPool.TransactOpts, asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.RebalanceStableBorrowRate(&_AaveV2LendingPool.TransactOpts, asset, user)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "repay", asset, amount, rateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) Repay(asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Repay(&_AaveV2LendingPool.TransactOpts, asset, amount, rateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) Repay(asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Repay(&_AaveV2LendingPool.TransactOpts, asset, amount, rateMode, onBehalfOf)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xb8d29276.
//
// Solidity: function setConfiguration(address asset, uint256 configuration) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) SetConfiguration(opts *bind.TransactOpts, asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "setConfiguration", asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xb8d29276.
//
// Solidity: function setConfiguration(address asset, uint256 configuration) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) SetConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetConfiguration(&_AaveV2LendingPool.TransactOpts, asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xb8d29276.
//
// Solidity: function setConfiguration(address asset, uint256 configuration) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) SetConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetConfiguration(&_AaveV2LendingPool.TransactOpts, asset, configuration)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) SetPause(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "setPause", val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) SetPause(val bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetPause(&_AaveV2LendingPool.TransactOpts, val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) SetPause(val bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetPause(&_AaveV2LendingPool.TransactOpts, val)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) SetReserveInterestRateStrategyAddress(opts *bind.TransactOpts, asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "setReserveInterestRateStrategyAddress", asset, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) SetReserveInterestRateStrategyAddress(asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetReserveInterestRateStrategyAddress(&_AaveV2LendingPool.TransactOpts, asset, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) SetReserveInterestRateStrategyAddress(asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetReserveInterestRateStrategyAddress(&_AaveV2LendingPool.TransactOpts, asset, rateStrategyAddress)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "setUserUseReserveAsCollateral", asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetUserUseReserveAsCollateral(&_AaveV2LendingPool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SetUserUseReserveAsCollateral(&_AaveV2LendingPool.TransactOpts, asset, useAsCollateral)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 rateMode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, asset common.Address, rateMode *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "swapBorrowRateMode", asset, rateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 rateMode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolSession) SwapBorrowRateMode(asset common.Address, rateMode *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SwapBorrowRateMode(&_AaveV2LendingPool.TransactOpts, asset, rateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 rateMode) returns()
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) SwapBorrowRateMode(asset common.Address, rateMode *big.Int) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.SwapBorrowRateMode(&_AaveV2LendingPool.TransactOpts, asset, rateMode)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Withdraw(&_AaveV2LendingPool.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AaveV2LendingPool *AaveV2LendingPoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV2LendingPool.Contract.Withdraw(&_AaveV2LendingPool.TransactOpts, asset, amount, to)
}

// AaveV2LendingPoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolBorrowIterator struct {
	Event *AaveV2LendingPoolBorrow // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolBorrow)
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
		it.Event = new(AaveV2LendingPoolBorrow)
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
func (it *AaveV2LendingPoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolBorrow represents a Borrow event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolBorrow struct {
	Reserve        common.Address
	User           common.Address
	OnBehalfOf     common.Address
	Amount         *big.Int
	BorrowRateMode *big.Int
	BorrowRate     *big.Int
	Referral       uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint256 borrowRateMode, uint256 borrowRate, uint16 indexed referral)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*AaveV2LendingPoolBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolBorrowIterator{contract: _AaveV2LendingPool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint256 borrowRateMode, uint256 borrowRate, uint16 indexed referral)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolBorrow, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolBorrow)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint256 borrowRateMode, uint256 borrowRate, uint16 indexed referral)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseBorrow(log types.Log) (*AaveV2LendingPoolBorrow, error) {
	event := new(AaveV2LendingPoolBorrow)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolDepositIterator struct {
	Event *AaveV2LendingPoolDeposit // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolDeposit)
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
		it.Event = new(AaveV2LendingPoolDeposit)
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
func (it *AaveV2LendingPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolDeposit represents a Deposit event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolDeposit struct {
	Reserve    common.Address
	User       common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterDeposit(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*AaveV2LendingPoolDepositIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolDepositIterator{contract: _AaveV2LendingPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolDeposit, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolDeposit)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseDeposit(log types.Log) (*AaveV2LendingPoolDeposit, error) {
	event := new(AaveV2LendingPoolDeposit)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolFlashLoanIterator struct {
	Event *AaveV2LendingPoolFlashLoan // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolFlashLoan)
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
		it.Event = new(AaveV2LendingPoolFlashLoan)
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
func (it *AaveV2LendingPoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolFlashLoan represents a FlashLoan event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolFlashLoan struct {
	Target       common.Address
	Initiator    common.Address
	Asset        common.Address
	Amount       *big.Int
	Premium      *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac.
//
// Solidity: event FlashLoan(address indexed target, address indexed initiator, address indexed asset, uint256 amount, uint256 premium, uint16 referralCode)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, target []common.Address, initiator []common.Address, asset []common.Address) (*AaveV2LendingPoolFlashLoanIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "FlashLoan", targetRule, initiatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolFlashLoanIterator{contract: _AaveV2LendingPool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac.
//
// Solidity: event FlashLoan(address indexed target, address indexed initiator, address indexed asset, uint256 amount, uint256 premium, uint16 referralCode)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolFlashLoan, target []common.Address, initiator []common.Address, asset []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "FlashLoan", targetRule, initiatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolFlashLoan)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac.
//
// Solidity: event FlashLoan(address indexed target, address indexed initiator, address indexed asset, uint256 amount, uint256 premium, uint16 referralCode)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseFlashLoan(log types.Log) (*AaveV2LendingPoolFlashLoan, error) {
	event := new(AaveV2LendingPoolFlashLoan)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolLiquidationCallIterator struct {
	Event *AaveV2LendingPoolLiquidationCall // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolLiquidationCall)
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
		it.Event = new(AaveV2LendingPoolLiquidationCall)
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
func (it *AaveV2LendingPoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolLiquidationCall represents a LiquidationCall event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolLiquidationCall struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (*AaveV2LendingPoolLiquidationCallIterator, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolLiquidationCallIterator{contract: _AaveV2LendingPool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolLiquidationCall, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (event.Subscription, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolLiquidationCall)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseLiquidationCall(log types.Log) (*AaveV2LendingPoolLiquidationCall, error) {
	event := new(AaveV2LendingPoolLiquidationCall)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolPausedIterator struct {
	Event *AaveV2LendingPoolPaused // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolPaused)
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
		it.Event = new(AaveV2LendingPoolPaused)
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
func (it *AaveV2LendingPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolPaused represents a Paused event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*AaveV2LendingPoolPausedIterator, error) {

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolPausedIterator{contract: _AaveV2LendingPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolPaused) (event.Subscription, error) {

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolPaused)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParsePaused(log types.Log) (*AaveV2LendingPoolPaused, error) {
	event := new(AaveV2LendingPoolPaused)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolRebalanceStableBorrowRateIterator is returned from FilterRebalanceStableBorrowRate and is used to iterate over the raw logs and unpacked data for RebalanceStableBorrowRate events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolRebalanceStableBorrowRateIterator struct {
	Event *AaveV2LendingPoolRebalanceStableBorrowRate // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolRebalanceStableBorrowRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolRebalanceStableBorrowRate)
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
		it.Event = new(AaveV2LendingPoolRebalanceStableBorrowRate)
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
func (it *AaveV2LendingPoolRebalanceStableBorrowRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolRebalanceStableBorrowRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolRebalanceStableBorrowRate represents a RebalanceStableBorrowRate event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolRebalanceStableBorrowRate struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRebalanceStableBorrowRate is a free log retrieval operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterRebalanceStableBorrowRate(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV2LendingPoolRebalanceStableBorrowRateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolRebalanceStableBorrowRateIterator{contract: _AaveV2LendingPool.contract, event: "RebalanceStableBorrowRate", logs: logs, sub: sub}, nil
}

// WatchRebalanceStableBorrowRate is a free log subscription operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchRebalanceStableBorrowRate(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolRebalanceStableBorrowRate, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolRebalanceStableBorrowRate)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
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

// ParseRebalanceStableBorrowRate is a log parse operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseRebalanceStableBorrowRate(log types.Log) (*AaveV2LendingPoolRebalanceStableBorrowRate, error) {
	event := new(AaveV2LendingPoolRebalanceStableBorrowRate)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolRepayIterator struct {
	Event *AaveV2LendingPoolRepay // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolRepay)
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
		it.Event = new(AaveV2LendingPoolRepay)
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
func (it *AaveV2LendingPoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolRepay represents a Repay event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolRepay struct {
	Reserve common.Address
	User    common.Address
	Repayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, repayer []common.Address) (*AaveV2LendingPoolRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolRepayIterator{contract: _AaveV2LendingPool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolRepay, reserve []common.Address, user []common.Address, repayer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolRepay)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseRepay(log types.Log) (*AaveV2LendingPoolRepay, error) {
	event := new(AaveV2LendingPoolRepay)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolReserveDataUpdatedIterator struct {
	Event *AaveV2LendingPoolReserveDataUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolReserveDataUpdated)
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
		it.Event = new(AaveV2LendingPoolReserveDataUpdated)
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
func (it *AaveV2LendingPoolReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolReserveDataUpdated represents a ReserveDataUpdated event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*AaveV2LendingPoolReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolReserveDataUpdatedIterator{contract: _AaveV2LendingPool.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolReserveDataUpdated)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
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

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseReserveDataUpdated(log types.Log) (*AaveV2LendingPoolReserveDataUpdated, error) {
	event := new(AaveV2LendingPoolReserveDataUpdated)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolReserveUsedAsCollateralDisabledIterator struct {
	Event *AaveV2LendingPoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolReserveUsedAsCollateralDisabled)
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
		it.Event = new(AaveV2LendingPoolReserveUsedAsCollateralDisabled)
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
func (it *AaveV2LendingPoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV2LendingPoolReserveUsedAsCollateralDisabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolReserveUsedAsCollateralDisabledIterator{contract: _AaveV2LendingPool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolReserveUsedAsCollateralDisabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolReserveUsedAsCollateralDisabled)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
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

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*AaveV2LendingPoolReserveUsedAsCollateralDisabled, error) {
	event := new(AaveV2LendingPoolReserveUsedAsCollateralDisabled)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolReserveUsedAsCollateralEnabledIterator struct {
	Event *AaveV2LendingPoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolReserveUsedAsCollateralEnabled)
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
		it.Event = new(AaveV2LendingPoolReserveUsedAsCollateralEnabled)
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
func (it *AaveV2LendingPoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV2LendingPoolReserveUsedAsCollateralEnabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolReserveUsedAsCollateralEnabledIterator{contract: _AaveV2LendingPool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolReserveUsedAsCollateralEnabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolReserveUsedAsCollateralEnabled)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
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

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*AaveV2LendingPoolReserveUsedAsCollateralEnabled, error) {
	event := new(AaveV2LendingPoolReserveUsedAsCollateralEnabled)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolSwapIterator struct {
	Event *AaveV2LendingPoolSwap // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolSwap)
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
		it.Event = new(AaveV2LendingPoolSwap)
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
func (it *AaveV2LendingPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolSwap represents a Swap event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolSwap struct {
	Reserve  common.Address
	User     common.Address
	RateMode *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed reserve, address indexed user, uint256 rateMode)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterSwap(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV2LendingPoolSwapIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "Swap", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolSwapIterator{contract: _AaveV2LendingPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed reserve, address indexed user, uint256 rateMode)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolSwap, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "Swap", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolSwap)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed reserve, address indexed user, uint256 rateMode)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseSwap(log types.Log) (*AaveV2LendingPoolSwap, error) {
	event := new(AaveV2LendingPoolSwap)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolUnpausedIterator struct {
	Event *AaveV2LendingPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolUnpaused)
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
		it.Event = new(AaveV2LendingPoolUnpaused)
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
func (it *AaveV2LendingPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolUnpaused represents a Unpaused event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AaveV2LendingPoolUnpausedIterator, error) {

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolUnpausedIterator{contract: _AaveV2LendingPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolUnpaused)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseUnpaused(log types.Log) (*AaveV2LendingPoolUnpaused, error) {
	event := new(AaveV2LendingPoolUnpaused)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AaveV2LendingPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolWithdrawIterator struct {
	Event *AaveV2LendingPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *AaveV2LendingPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2LendingPoolWithdraw)
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
		it.Event = new(AaveV2LendingPoolWithdraw)
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
func (it *AaveV2LendingPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2LendingPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2LendingPoolWithdraw represents a Withdraw event raised by the AaveV2LendingPool contract.
type AaveV2LendingPoolWithdraw struct {
	Reserve common.Address
	User    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, to []common.Address) (*AaveV2LendingPoolWithdrawIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.FilterLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2LendingPoolWithdrawIterator{contract: _AaveV2LendingPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AaveV2LendingPoolWithdraw, reserve []common.Address, user []common.Address, to []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveV2LendingPool.contract.WatchLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2LendingPoolWithdraw)
				if err := _AaveV2LendingPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AaveV2LendingPool *AaveV2LendingPoolFilterer) ParseWithdraw(log types.Log) (*AaveV2LendingPoolWithdraw, error) {
	event := new(AaveV2LendingPoolWithdraw)
	if err := _AaveV2LendingPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
