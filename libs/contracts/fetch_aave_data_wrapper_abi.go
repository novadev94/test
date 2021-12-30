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

// IFetchAaveDataWrapperReserveConfigData is an auto generated low-level Go binding around an user-defined struct.
type IFetchAaveDataWrapperReserveConfigData struct {
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	ATokenAddress            common.Address
}

// IFetchAaveDataWrapperReserveData is an auto generated low-level Go binding around an user-defined struct.
type IFetchAaveDataWrapperReserveData struct {
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	TotalLiquidity          *big.Int
	UtilizationRate         *big.Int
}

// IFetchAaveDataWrapperUserAccountData is an auto generated low-level Go binding around an user-defined struct.
type IFetchAaveDataWrapperUserAccountData struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}

// IFetchAaveDataWrapperUserReserveData is an auto generated low-level Go binding around an user-defined struct.
type IFetchAaveDataWrapperUserReserveData struct {
	CurrentATokenBalance     *big.Int
	LiquidityRate            *big.Int
	PoolShareInPrecision     *big.Int
	UsageAsCollateralEnabled bool
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	OriginationFee           *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
}

// FetchAaveDataWrapperABI is the input ABI used to generate the binding from.
const FetchAaveDataWrapperABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isV1\",\"type\":\"bool\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"reserves\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isV1\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_reserves\",\"type\":\"address[]\"}],\"name\":\"getReservesConfigurationData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"}],\"internalType\":\"structIFetchAaveDataWrapper.ReserveConfigData[]\",\"name\":\"configsData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isV1\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_reserves\",\"type\":\"address[]\"}],\"name\":\"getReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsStable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsVariable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationRate\",\"type\":\"uint256\"}],\"internalType\":\"structIFetchAaveDataWrapper.ReserveData[]\",\"name\":\"reservesData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isV1\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getSingleUserAccountData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidityETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeesETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"internalType\":\"structIFetchAaveDataWrapper.UserAccountData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendingPoolV1\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getSingleUserReserveDataV1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolShareInPrecision\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"}],\"internalType\":\"structIFetchAaveDataWrapper.UserReserveData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIProtocolDataProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getSingleUserReserveDataV2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolShareInPrecision\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"}],\"internalType\":\"structIFetchAaveDataWrapper.UserReserveData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isV1\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"}],\"name\":\"getUserAccountsData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidityETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeesETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"internalType\":\"structIFetchAaveDataWrapper.UserAccountData[]\",\"name\":\"accountsData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isV1\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_reserves\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolShareInPrecision\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"}],\"internalType\":\"structIFetchAaveDataWrapper.UserReserveData[]\",\"name\":\"userReservesData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isV1\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"}],\"name\":\"getUsersReserveData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolShareInPrecision\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"}],\"internalType\":\"structIFetchAaveDataWrapper.UserReserveData[]\",\"name\":\"userReservesData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FetchAaveDataWrapper is an auto generated Go binding around an Ethereum contract.
type FetchAaveDataWrapper struct {
	FetchAaveDataWrapperCaller     // Read-only binding to the contract
	FetchAaveDataWrapperTransactor // Write-only binding to the contract
	FetchAaveDataWrapperFilterer   // Log filterer for contract events
}

// FetchAaveDataWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type FetchAaveDataWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FetchAaveDataWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FetchAaveDataWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FetchAaveDataWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FetchAaveDataWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FetchAaveDataWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FetchAaveDataWrapperSession struct {
	Contract     *FetchAaveDataWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FetchAaveDataWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FetchAaveDataWrapperCallerSession struct {
	Contract *FetchAaveDataWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FetchAaveDataWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FetchAaveDataWrapperTransactorSession struct {
	Contract     *FetchAaveDataWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FetchAaveDataWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type FetchAaveDataWrapperRaw struct {
	Contract *FetchAaveDataWrapper // Generic contract binding to access the raw methods on
}

// FetchAaveDataWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FetchAaveDataWrapperCallerRaw struct {
	Contract *FetchAaveDataWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// FetchAaveDataWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FetchAaveDataWrapperTransactorRaw struct {
	Contract *FetchAaveDataWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFetchAaveDataWrapper creates a new instance of FetchAaveDataWrapper, bound to a specific deployed contract.
func NewFetchAaveDataWrapper(address common.Address, backend bind.ContractBackend) (*FetchAaveDataWrapper, error) {
	contract, err := bindFetchAaveDataWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapper{FetchAaveDataWrapperCaller: FetchAaveDataWrapperCaller{contract: contract}, FetchAaveDataWrapperTransactor: FetchAaveDataWrapperTransactor{contract: contract}, FetchAaveDataWrapperFilterer: FetchAaveDataWrapperFilterer{contract: contract}}, nil
}

// NewFetchAaveDataWrapperCaller creates a new read-only instance of FetchAaveDataWrapper, bound to a specific deployed contract.
func NewFetchAaveDataWrapperCaller(address common.Address, caller bind.ContractCaller) (*FetchAaveDataWrapperCaller, error) {
	contract, err := bindFetchAaveDataWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperCaller{contract: contract}, nil
}

// NewFetchAaveDataWrapperTransactor creates a new write-only instance of FetchAaveDataWrapper, bound to a specific deployed contract.
func NewFetchAaveDataWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*FetchAaveDataWrapperTransactor, error) {
	contract, err := bindFetchAaveDataWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperTransactor{contract: contract}, nil
}

// NewFetchAaveDataWrapperFilterer creates a new log filterer instance of FetchAaveDataWrapper, bound to a specific deployed contract.
func NewFetchAaveDataWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*FetchAaveDataWrapperFilterer, error) {
	contract, err := bindFetchAaveDataWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperFilterer{contract: contract}, nil
}

// bindFetchAaveDataWrapper binds a generic wrapper to an already deployed contract.
func bindFetchAaveDataWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FetchAaveDataWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FetchAaveDataWrapper *FetchAaveDataWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FetchAaveDataWrapper.Contract.FetchAaveDataWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FetchAaveDataWrapper *FetchAaveDataWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.FetchAaveDataWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FetchAaveDataWrapper *FetchAaveDataWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.FetchAaveDataWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FetchAaveDataWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) Admin() (common.Address, error) {
	return _FetchAaveDataWrapper.Contract.Admin(&_FetchAaveDataWrapper.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) Admin() (common.Address, error) {
	return _FetchAaveDataWrapper.Contract.Admin(&_FetchAaveDataWrapper.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetAlerters() ([]common.Address, error) {
	return _FetchAaveDataWrapper.Contract.GetAlerters(&_FetchAaveDataWrapper.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetAlerters() ([]common.Address, error) {
	return _FetchAaveDataWrapper.Contract.GetAlerters(&_FetchAaveDataWrapper.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetOperators() ([]common.Address, error) {
	return _FetchAaveDataWrapper.Contract.GetOperators(&_FetchAaveDataWrapper.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetOperators() ([]common.Address, error) {
	return _FetchAaveDataWrapper.Contract.GetOperators(&_FetchAaveDataWrapper.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x5ababafd.
//
// Solidity: function getReserves(address pool, bool isV1) view returns(address[] reserves)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetReserves(opts *bind.CallOpts, pool common.Address, isV1 bool) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getReserves", pool, isV1)
	return *ret0, err
}

// GetReserves is a free data retrieval call binding the contract method 0x5ababafd.
//
// Solidity: function getReserves(address pool, bool isV1) view returns(address[] reserves)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetReserves(pool common.Address, isV1 bool) ([]common.Address, error) {
	return _FetchAaveDataWrapper.Contract.GetReserves(&_FetchAaveDataWrapper.CallOpts, pool, isV1)
}

// GetReserves is a free data retrieval call binding the contract method 0x5ababafd.
//
// Solidity: function getReserves(address pool, bool isV1) view returns(address[] reserves)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetReserves(pool common.Address, isV1 bool) ([]common.Address, error) {
	return _FetchAaveDataWrapper.Contract.GetReserves(&_FetchAaveDataWrapper.CallOpts, pool, isV1)
}

// GetReservesConfigurationData is a free data retrieval call binding the contract method 0xf422ced9.
//
// Solidity: function getReservesConfigurationData(address pool, bool isV1, address[] _reserves) view returns((uint256,uint256,uint256,bool,bool,bool,bool,address)[] configsData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetReservesConfigurationData(opts *bind.CallOpts, pool common.Address, isV1 bool, _reserves []common.Address) ([]IFetchAaveDataWrapperReserveConfigData, error) {
	var (
		ret0 = new([]IFetchAaveDataWrapperReserveConfigData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getReservesConfigurationData", pool, isV1, _reserves)
	return *ret0, err
}

// GetReservesConfigurationData is a free data retrieval call binding the contract method 0xf422ced9.
//
// Solidity: function getReservesConfigurationData(address pool, bool isV1, address[] _reserves) view returns((uint256,uint256,uint256,bool,bool,bool,bool,address)[] configsData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetReservesConfigurationData(pool common.Address, isV1 bool, _reserves []common.Address) ([]IFetchAaveDataWrapperReserveConfigData, error) {
	return _FetchAaveDataWrapper.Contract.GetReservesConfigurationData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserves)
}

// GetReservesConfigurationData is a free data retrieval call binding the contract method 0xf422ced9.
//
// Solidity: function getReservesConfigurationData(address pool, bool isV1, address[] _reserves) view returns((uint256,uint256,uint256,bool,bool,bool,bool,address)[] configsData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetReservesConfigurationData(pool common.Address, isV1 bool, _reserves []common.Address) ([]IFetchAaveDataWrapperReserveConfigData, error) {
	return _FetchAaveDataWrapper.Contract.GetReservesConfigurationData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserves)
}

// GetReservesData is a free data retrieval call binding the contract method 0xad68e9ec.
//
// Solidity: function getReservesData(address pool, bool isV1, address[] _reserves) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] reservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetReservesData(opts *bind.CallOpts, pool common.Address, isV1 bool, _reserves []common.Address) ([]IFetchAaveDataWrapperReserveData, error) {
	var (
		ret0 = new([]IFetchAaveDataWrapperReserveData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getReservesData", pool, isV1, _reserves)
	return *ret0, err
}

// GetReservesData is a free data retrieval call binding the contract method 0xad68e9ec.
//
// Solidity: function getReservesData(address pool, bool isV1, address[] _reserves) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] reservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetReservesData(pool common.Address, isV1 bool, _reserves []common.Address) ([]IFetchAaveDataWrapperReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetReservesData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserves)
}

// GetReservesData is a free data retrieval call binding the contract method 0xad68e9ec.
//
// Solidity: function getReservesData(address pool, bool isV1, address[] _reserves) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] reservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetReservesData(pool common.Address, isV1 bool, _reserves []common.Address) ([]IFetchAaveDataWrapperReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetReservesData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserves)
}

// GetSingleUserAccountData is a free data retrieval call binding the contract method 0x3c1bf462.
//
// Solidity: function getSingleUserAccountData(address pool, bool isV1, address _user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetSingleUserAccountData(opts *bind.CallOpts, pool common.Address, isV1 bool, _user common.Address) (IFetchAaveDataWrapperUserAccountData, error) {
	var (
		ret0 = new(IFetchAaveDataWrapperUserAccountData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getSingleUserAccountData", pool, isV1, _user)
	return *ret0, err
}

// GetSingleUserAccountData is a free data retrieval call binding the contract method 0x3c1bf462.
//
// Solidity: function getSingleUserAccountData(address pool, bool isV1, address _user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetSingleUserAccountData(pool common.Address, isV1 bool, _user common.Address) (IFetchAaveDataWrapperUserAccountData, error) {
	return _FetchAaveDataWrapper.Contract.GetSingleUserAccountData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _user)
}

// GetSingleUserAccountData is a free data retrieval call binding the contract method 0x3c1bf462.
//
// Solidity: function getSingleUserAccountData(address pool, bool isV1, address _user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetSingleUserAccountData(pool common.Address, isV1 bool, _user common.Address) (IFetchAaveDataWrapperUserAccountData, error) {
	return _FetchAaveDataWrapper.Contract.GetSingleUserAccountData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _user)
}

// GetSingleUserReserveDataV1 is a free data retrieval call binding the contract method 0x2dd7badc.
//
// Solidity: function getSingleUserReserveDataV1(address pool, address _reserve, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetSingleUserReserveDataV1(opts *bind.CallOpts, pool common.Address, _reserve common.Address, _user common.Address) (IFetchAaveDataWrapperUserReserveData, error) {
	var (
		ret0 = new(IFetchAaveDataWrapperUserReserveData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getSingleUserReserveDataV1", pool, _reserve, _user)
	return *ret0, err
}

// GetSingleUserReserveDataV1 is a free data retrieval call binding the contract method 0x2dd7badc.
//
// Solidity: function getSingleUserReserveDataV1(address pool, address _reserve, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetSingleUserReserveDataV1(pool common.Address, _reserve common.Address, _user common.Address) (IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetSingleUserReserveDataV1(&_FetchAaveDataWrapper.CallOpts, pool, _reserve, _user)
}

// GetSingleUserReserveDataV1 is a free data retrieval call binding the contract method 0x2dd7badc.
//
// Solidity: function getSingleUserReserveDataV1(address pool, address _reserve, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetSingleUserReserveDataV1(pool common.Address, _reserve common.Address, _user common.Address) (IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetSingleUserReserveDataV1(&_FetchAaveDataWrapper.CallOpts, pool, _reserve, _user)
}

// GetSingleUserReserveDataV2 is a free data retrieval call binding the contract method 0x1562bc9e.
//
// Solidity: function getSingleUserReserveDataV2(address provider, address _reserve, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetSingleUserReserveDataV2(opts *bind.CallOpts, provider common.Address, _reserve common.Address, _user common.Address) (IFetchAaveDataWrapperUserReserveData, error) {
	var (
		ret0 = new(IFetchAaveDataWrapperUserReserveData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getSingleUserReserveDataV2", provider, _reserve, _user)
	return *ret0, err
}

// GetSingleUserReserveDataV2 is a free data retrieval call binding the contract method 0x1562bc9e.
//
// Solidity: function getSingleUserReserveDataV2(address provider, address _reserve, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetSingleUserReserveDataV2(provider common.Address, _reserve common.Address, _user common.Address) (IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetSingleUserReserveDataV2(&_FetchAaveDataWrapper.CallOpts, provider, _reserve, _user)
}

// GetSingleUserReserveDataV2 is a free data retrieval call binding the contract method 0x1562bc9e.
//
// Solidity: function getSingleUserReserveDataV2(address provider, address _reserve, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetSingleUserReserveDataV2(provider common.Address, _reserve common.Address, _user common.Address) (IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetSingleUserReserveDataV2(&_FetchAaveDataWrapper.CallOpts, provider, _reserve, _user)
}

// GetUserAccountsData is a free data retrieval call binding the contract method 0xaf38a675.
//
// Solidity: function getUserAccountsData(address pool, bool isV1, address[] _users) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] accountsData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetUserAccountsData(opts *bind.CallOpts, pool common.Address, isV1 bool, _users []common.Address) ([]IFetchAaveDataWrapperUserAccountData, error) {
	var (
		ret0 = new([]IFetchAaveDataWrapperUserAccountData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getUserAccountsData", pool, isV1, _users)
	return *ret0, err
}

// GetUserAccountsData is a free data retrieval call binding the contract method 0xaf38a675.
//
// Solidity: function getUserAccountsData(address pool, bool isV1, address[] _users) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] accountsData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetUserAccountsData(pool common.Address, isV1 bool, _users []common.Address) ([]IFetchAaveDataWrapperUserAccountData, error) {
	return _FetchAaveDataWrapper.Contract.GetUserAccountsData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _users)
}

// GetUserAccountsData is a free data retrieval call binding the contract method 0xaf38a675.
//
// Solidity: function getUserAccountsData(address pool, bool isV1, address[] _users) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] accountsData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetUserAccountsData(pool common.Address, isV1 bool, _users []common.Address) ([]IFetchAaveDataWrapperUserAccountData, error) {
	return _FetchAaveDataWrapper.Contract.GetUserAccountsData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _users)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x42c3dae1.
//
// Solidity: function getUserReservesData(address pool, bool isV1, address[] _reserves, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] userReservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetUserReservesData(opts *bind.CallOpts, pool common.Address, isV1 bool, _reserves []common.Address, _user common.Address) ([]IFetchAaveDataWrapperUserReserveData, error) {
	var (
		ret0 = new([]IFetchAaveDataWrapperUserReserveData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getUserReservesData", pool, isV1, _reserves, _user)
	return *ret0, err
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x42c3dae1.
//
// Solidity: function getUserReservesData(address pool, bool isV1, address[] _reserves, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] userReservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetUserReservesData(pool common.Address, isV1 bool, _reserves []common.Address, _user common.Address) ([]IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetUserReservesData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserves, _user)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x42c3dae1.
//
// Solidity: function getUserReservesData(address pool, bool isV1, address[] _reserves, address _user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] userReservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetUserReservesData(pool common.Address, isV1 bool, _reserves []common.Address, _user common.Address) ([]IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetUserReservesData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserves, _user)
}

// GetUsersReserveData is a free data retrieval call binding the contract method 0x5f8b3356.
//
// Solidity: function getUsersReserveData(address pool, bool isV1, address _reserve, address[] _users) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] userReservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) GetUsersReserveData(opts *bind.CallOpts, pool common.Address, isV1 bool, _reserve common.Address, _users []common.Address) ([]IFetchAaveDataWrapperUserReserveData, error) {
	var (
		ret0 = new([]IFetchAaveDataWrapperUserReserveData)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "getUsersReserveData", pool, isV1, _reserve, _users)
	return *ret0, err
}

// GetUsersReserveData is a free data retrieval call binding the contract method 0x5f8b3356.
//
// Solidity: function getUsersReserveData(address pool, bool isV1, address _reserve, address[] _users) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] userReservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) GetUsersReserveData(pool common.Address, isV1 bool, _reserve common.Address, _users []common.Address) ([]IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetUsersReserveData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserve, _users)
}

// GetUsersReserveData is a free data retrieval call binding the contract method 0x5f8b3356.
//
// Solidity: function getUsersReserveData(address pool, bool isV1, address _reserve, address[] _users) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] userReservesData)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) GetUsersReserveData(pool common.Address, isV1 bool, _reserve common.Address, _users []common.Address) ([]IFetchAaveDataWrapperUserReserveData, error) {
	return _FetchAaveDataWrapper.Contract.GetUsersReserveData(&_FetchAaveDataWrapper.CallOpts, pool, isV1, _reserve, _users)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FetchAaveDataWrapper.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) PendingAdmin() (common.Address, error) {
	return _FetchAaveDataWrapper.Contract.PendingAdmin(&_FetchAaveDataWrapper.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperCallerSession) PendingAdmin() (common.Address, error) {
	return _FetchAaveDataWrapper.Contract.PendingAdmin(&_FetchAaveDataWrapper.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.AddAlerter(&_FetchAaveDataWrapper.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.AddAlerter(&_FetchAaveDataWrapper.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.AddOperator(&_FetchAaveDataWrapper.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.AddOperator(&_FetchAaveDataWrapper.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) ClaimAdmin() (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.ClaimAdmin(&_FetchAaveDataWrapper.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.ClaimAdmin(&_FetchAaveDataWrapper.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.RemoveAlerter(&_FetchAaveDataWrapper.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.RemoveAlerter(&_FetchAaveDataWrapper.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.RemoveOperator(&_FetchAaveDataWrapper.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.RemoveOperator(&_FetchAaveDataWrapper.TransactOpts, operator)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.TransferAdmin(&_FetchAaveDataWrapper.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.TransferAdmin(&_FetchAaveDataWrapper.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.TransferAdminQuickly(&_FetchAaveDataWrapper.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.TransferAdminQuickly(&_FetchAaveDataWrapper.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.WithdrawEther(&_FetchAaveDataWrapper.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.WithdrawEther(&_FetchAaveDataWrapper.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.WithdrawToken(&_FetchAaveDataWrapper.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_FetchAaveDataWrapper *FetchAaveDataWrapperTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FetchAaveDataWrapper.Contract.WithdrawToken(&_FetchAaveDataWrapper.TransactOpts, token, amount, sendTo)
}

// FetchAaveDataWrapperAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperAdminClaimedIterator struct {
	Event *FetchAaveDataWrapperAdminClaimed // Event containing the contract specifics and raw log

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
func (it *FetchAaveDataWrapperAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FetchAaveDataWrapperAdminClaimed)
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
		it.Event = new(FetchAaveDataWrapperAdminClaimed)
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
func (it *FetchAaveDataWrapperAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FetchAaveDataWrapperAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FetchAaveDataWrapperAdminClaimed represents a AdminClaimed event raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*FetchAaveDataWrapperAdminClaimedIterator, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperAdminClaimedIterator{contract: _FetchAaveDataWrapper.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *FetchAaveDataWrapperAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FetchAaveDataWrapperAdminClaimed)
				if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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

// ParseAdminClaimed is a log parse operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) ParseAdminClaimed(log types.Log) (*FetchAaveDataWrapperAdminClaimed, error) {
	event := new(FetchAaveDataWrapperAdminClaimed)
	if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FetchAaveDataWrapperAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperAlerterAddedIterator struct {
	Event *FetchAaveDataWrapperAlerterAdded // Event containing the contract specifics and raw log

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
func (it *FetchAaveDataWrapperAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FetchAaveDataWrapperAlerterAdded)
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
		it.Event = new(FetchAaveDataWrapperAlerterAdded)
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
func (it *FetchAaveDataWrapperAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FetchAaveDataWrapperAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FetchAaveDataWrapperAlerterAdded represents a AlerterAdded event raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*FetchAaveDataWrapperAlerterAddedIterator, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperAlerterAddedIterator{contract: _FetchAaveDataWrapper.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *FetchAaveDataWrapperAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FetchAaveDataWrapperAlerterAdded)
				if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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

// ParseAlerterAdded is a log parse operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) ParseAlerterAdded(log types.Log) (*FetchAaveDataWrapperAlerterAdded, error) {
	event := new(FetchAaveDataWrapperAlerterAdded)
	if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FetchAaveDataWrapperEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperEtherWithdrawIterator struct {
	Event *FetchAaveDataWrapperEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *FetchAaveDataWrapperEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FetchAaveDataWrapperEtherWithdraw)
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
		it.Event = new(FetchAaveDataWrapperEtherWithdraw)
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
func (it *FetchAaveDataWrapperEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FetchAaveDataWrapperEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FetchAaveDataWrapperEtherWithdraw represents a EtherWithdraw event raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*FetchAaveDataWrapperEtherWithdrawIterator, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperEtherWithdrawIterator{contract: _FetchAaveDataWrapper.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *FetchAaveDataWrapperEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FetchAaveDataWrapperEtherWithdraw)
				if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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

// ParseEtherWithdraw is a log parse operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) ParseEtherWithdraw(log types.Log) (*FetchAaveDataWrapperEtherWithdraw, error) {
	event := new(FetchAaveDataWrapperEtherWithdraw)
	if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FetchAaveDataWrapperOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperOperatorAddedIterator struct {
	Event *FetchAaveDataWrapperOperatorAdded // Event containing the contract specifics and raw log

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
func (it *FetchAaveDataWrapperOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FetchAaveDataWrapperOperatorAdded)
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
		it.Event = new(FetchAaveDataWrapperOperatorAdded)
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
func (it *FetchAaveDataWrapperOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FetchAaveDataWrapperOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FetchAaveDataWrapperOperatorAdded represents a OperatorAdded event raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*FetchAaveDataWrapperOperatorAddedIterator, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperOperatorAddedIterator{contract: _FetchAaveDataWrapper.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *FetchAaveDataWrapperOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FetchAaveDataWrapperOperatorAdded)
				if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) ParseOperatorAdded(log types.Log) (*FetchAaveDataWrapperOperatorAdded, error) {
	event := new(FetchAaveDataWrapperOperatorAdded)
	if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FetchAaveDataWrapperTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperTokenWithdrawIterator struct {
	Event *FetchAaveDataWrapperTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *FetchAaveDataWrapperTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FetchAaveDataWrapperTokenWithdraw)
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
		it.Event = new(FetchAaveDataWrapperTokenWithdraw)
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
func (it *FetchAaveDataWrapperTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FetchAaveDataWrapperTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FetchAaveDataWrapperTokenWithdraw represents a TokenWithdraw event raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*FetchAaveDataWrapperTokenWithdrawIterator, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperTokenWithdrawIterator{contract: _FetchAaveDataWrapper.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *FetchAaveDataWrapperTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FetchAaveDataWrapperTokenWithdraw)
				if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) ParseTokenWithdraw(log types.Log) (*FetchAaveDataWrapperTokenWithdraw, error) {
	event := new(FetchAaveDataWrapperTokenWithdraw)
	if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FetchAaveDataWrapperTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperTransferAdminPendingIterator struct {
	Event *FetchAaveDataWrapperTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *FetchAaveDataWrapperTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FetchAaveDataWrapperTransferAdminPending)
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
		it.Event = new(FetchAaveDataWrapperTransferAdminPending)
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
func (it *FetchAaveDataWrapperTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FetchAaveDataWrapperTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FetchAaveDataWrapperTransferAdminPending represents a TransferAdminPending event raised by the FetchAaveDataWrapper contract.
type FetchAaveDataWrapperTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*FetchAaveDataWrapperTransferAdminPendingIterator, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &FetchAaveDataWrapperTransferAdminPendingIterator{contract: _FetchAaveDataWrapper.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *FetchAaveDataWrapperTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _FetchAaveDataWrapper.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FetchAaveDataWrapperTransferAdminPending)
				if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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

// ParseTransferAdminPending is a log parse operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_FetchAaveDataWrapper *FetchAaveDataWrapperFilterer) ParseTransferAdminPending(log types.Log) (*FetchAaveDataWrapperTransferAdminPending, error) {
	event := new(FetchAaveDataWrapperTransferAdminPending)
	if err := _FetchAaveDataWrapper.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}
