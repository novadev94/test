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

// AaveProtocolDataProviderTokenData is an auto generated low-level Go binding around an user-defined struct.
type AaveProtocolDataProviderTokenData struct {
	Symbol       string
	TokenAddress common.Address
}

// AaveV2ProtocolDataProviderABI is the input ABI used to generate the binding from.
const AaveV2ProtocolDataProviderABI = "[{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"addressesProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllATokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllReservesTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveTokensAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stableRateLastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AaveV2ProtocolDataProvider is an auto generated Go binding around an Ethereum contract.
type AaveV2ProtocolDataProvider struct {
	AaveV2ProtocolDataProviderCaller     // Read-only binding to the contract
	AaveV2ProtocolDataProviderTransactor // Write-only binding to the contract
	AaveV2ProtocolDataProviderFilterer   // Log filterer for contract events
}

// AaveV2ProtocolDataProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV2ProtocolDataProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2ProtocolDataProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV2ProtocolDataProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2ProtocolDataProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV2ProtocolDataProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2ProtocolDataProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV2ProtocolDataProviderSession struct {
	Contract     *AaveV2ProtocolDataProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AaveV2ProtocolDataProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV2ProtocolDataProviderCallerSession struct {
	Contract *AaveV2ProtocolDataProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// AaveV2ProtocolDataProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV2ProtocolDataProviderTransactorSession struct {
	Contract     *AaveV2ProtocolDataProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// AaveV2ProtocolDataProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV2ProtocolDataProviderRaw struct {
	Contract *AaveV2ProtocolDataProvider // Generic contract binding to access the raw methods on
}

// AaveV2ProtocolDataProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV2ProtocolDataProviderCallerRaw struct {
	Contract *AaveV2ProtocolDataProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV2ProtocolDataProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV2ProtocolDataProviderTransactorRaw struct {
	Contract *AaveV2ProtocolDataProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV2ProtocolDataProvider creates a new instance of AaveV2ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV2ProtocolDataProvider(address common.Address, backend bind.ContractBackend) (*AaveV2ProtocolDataProvider, error) {
	contract, err := bindAaveV2ProtocolDataProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV2ProtocolDataProvider{AaveV2ProtocolDataProviderCaller: AaveV2ProtocolDataProviderCaller{contract: contract}, AaveV2ProtocolDataProviderTransactor: AaveV2ProtocolDataProviderTransactor{contract: contract}, AaveV2ProtocolDataProviderFilterer: AaveV2ProtocolDataProviderFilterer{contract: contract}}, nil
}

// NewAaveV2ProtocolDataProviderCaller creates a new read-only instance of AaveV2ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV2ProtocolDataProviderCaller(address common.Address, caller bind.ContractCaller) (*AaveV2ProtocolDataProviderCaller, error) {
	contract, err := bindAaveV2ProtocolDataProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2ProtocolDataProviderCaller{contract: contract}, nil
}

// NewAaveV2ProtocolDataProviderTransactor creates a new write-only instance of AaveV2ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV2ProtocolDataProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV2ProtocolDataProviderTransactor, error) {
	contract, err := bindAaveV2ProtocolDataProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2ProtocolDataProviderTransactor{contract: contract}, nil
}

// NewAaveV2ProtocolDataProviderFilterer creates a new log filterer instance of AaveV2ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV2ProtocolDataProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV2ProtocolDataProviderFilterer, error) {
	contract, err := bindAaveV2ProtocolDataProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV2ProtocolDataProviderFilterer{contract: contract}, nil
}

// bindAaveV2ProtocolDataProvider binds a generic wrapper to an already deployed contract.
func bindAaveV2ProtocolDataProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveV2ProtocolDataProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AaveV2ProtocolDataProvider.Contract.AaveV2ProtocolDataProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2ProtocolDataProvider.Contract.AaveV2ProtocolDataProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2ProtocolDataProvider.Contract.AaveV2ProtocolDataProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AaveV2ProtocolDataProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2ProtocolDataProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2ProtocolDataProvider.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AaveV2ProtocolDataProvider.contract.Call(opts, out, "ADDRESSES_PROVIDER")
	return *ret0, err
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveV2ProtocolDataProvider.Contract.ADDRESSESPROVIDER(&_AaveV2ProtocolDataProvider.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveV2ProtocolDataProvider.Contract.ADDRESSESPROVIDER(&_AaveV2ProtocolDataProvider.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCaller) GetAllATokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var (
		ret0 = new([]AaveProtocolDataProviderTokenData)
	)
	out := ret0
	err := _AaveV2ProtocolDataProvider.contract.Call(opts, out, "getAllATokens")
	return *ret0, err
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetAllATokens(&_AaveV2ProtocolDataProvider.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetAllATokens(&_AaveV2ProtocolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCaller) GetAllReservesTokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var (
		ret0 = new([]AaveProtocolDataProviderTokenData)
	)
	out := ret0
	err := _AaveV2ProtocolDataProvider.contract.Call(opts, out, "getAllReservesTokens")
	return *ret0, err
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetAllReservesTokens(&_AaveV2ProtocolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetAllReservesTokens(&_AaveV2ProtocolDataProvider.CallOpts)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCaller) GetReserveConfigurationData(opts *bind.CallOpts, asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	ret := new(struct {
		Decimals                 *big.Int
		Ltv                      *big.Int
		LiquidationThreshold     *big.Int
		LiquidationBonus         *big.Int
		ReserveFactor            *big.Int
		UsageAsCollateralEnabled bool
		BorrowingEnabled         bool
		StableBorrowRateEnabled  bool
		IsActive                 bool
		IsFrozen                 bool
	})
	out := ret
	err := _AaveV2ProtocolDataProvider.contract.Call(opts, out, "getReserveConfigurationData", asset)
	return *ret, err
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetReserveConfigurationData(&_AaveV2ProtocolDataProvider.CallOpts, asset)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetReserveConfigurationData(&_AaveV2ProtocolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	ret := new(struct {
		AvailableLiquidity      *big.Int
		TotalStableDebt         *big.Int
		TotalVariableDebt       *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		LastUpdateTimestamp     *big.Int
	})
	out := ret
	err := _AaveV2ProtocolDataProvider.contract.Call(opts, out, "getReserveData", asset)
	return *ret, err
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderSession) GetReserveData(asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetReserveData(&_AaveV2ProtocolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerSession) GetReserveData(asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetReserveData(&_AaveV2ProtocolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCaller) GetReserveTokensAddresses(opts *bind.CallOpts, asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	ret := new(struct {
		ATokenAddress            common.Address
		StableDebtTokenAddress   common.Address
		VariableDebtTokenAddress common.Address
	})
	out := ret
	err := _AaveV2ProtocolDataProvider.contract.Call(opts, out, "getReserveTokensAddresses", asset)
	return *ret, err
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetReserveTokensAddresses(&_AaveV2ProtocolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetReserveTokensAddresses(&_AaveV2ProtocolDataProvider.CallOpts, asset)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCaller) GetUserReserveData(opts *bind.CallOpts, asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	ret := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentStableDebt        *big.Int
		CurrentVariableDebt      *big.Int
		PrincipalStableDebt      *big.Int
		ScaledVariableDebt       *big.Int
		StableBorrowRate         *big.Int
		LiquidityRate            *big.Int
		StableRateLastUpdated    *big.Int
		UsageAsCollateralEnabled bool
	})
	out := ret
	err := _AaveV2ProtocolDataProvider.contract.Call(opts, out, "getUserReserveData", asset, user)
	return *ret, err
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetUserReserveData(&_AaveV2ProtocolDataProvider.CallOpts, asset, user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveV2ProtocolDataProvider *AaveV2ProtocolDataProviderCallerSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _AaveV2ProtocolDataProvider.Contract.GetUserReserveData(&_AaveV2ProtocolDataProvider.CallOpts, asset, user)
}
