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

// ATokenABI is the input ABI used to generate the binding from.
const ATokenABI = "[{\"inputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"_addressesProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_underlyingAssetDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toIndex\",\"type\":\"uint256\"}],\"name\":\"BalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"BurnOnLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"InterestRedirectionAllowanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_redirectedBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"InterestStreamRedirected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"MintOnDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_targetBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_targetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_redirectedBalanceAdded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_redirectedBalanceRemoved\",\"type\":\"uint256\"}],\"name\":\"RedirectedBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"UINT_MAX_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlyingAssetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectInterestStream\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectInterestStreamOf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"allowInterestRedirectionTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintOnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"principalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isTransferAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getInterestRedirectionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getRedirectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AToken is an auto generated Go binding around an Ethereum contract.
type AToken struct {
	ATokenCaller     // Read-only binding to the contract
	ATokenTransactor // Write-only binding to the contract
	ATokenFilterer   // Log filterer for contract events
}

// ATokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ATokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ATokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ATokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ATokenSession struct {
	Contract     *AToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ATokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ATokenCallerSession struct {
	Contract *ATokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ATokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ATokenTransactorSession struct {
	Contract     *ATokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ATokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ATokenRaw struct {
	Contract *AToken // Generic contract binding to access the raw methods on
}

// ATokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ATokenCallerRaw struct {
	Contract *ATokenCaller // Generic read-only contract binding to access the raw methods on
}

// ATokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ATokenTransactorRaw struct {
	Contract *ATokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAToken creates a new instance of AToken, bound to a specific deployed contract.
func NewAToken(address common.Address, backend bind.ContractBackend) (*AToken, error) {
	contract, err := bindAToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AToken{ATokenCaller: ATokenCaller{contract: contract}, ATokenTransactor: ATokenTransactor{contract: contract}, ATokenFilterer: ATokenFilterer{contract: contract}}, nil
}

// NewATokenCaller creates a new read-only instance of AToken, bound to a specific deployed contract.
func NewATokenCaller(address common.Address, caller bind.ContractCaller) (*ATokenCaller, error) {
	contract, err := bindAToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenCaller{contract: contract}, nil
}

// NewATokenTransactor creates a new write-only instance of AToken, bound to a specific deployed contract.
func NewATokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ATokenTransactor, error) {
	contract, err := bindAToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenTransactor{contract: contract}, nil
}

// NewATokenFilterer creates a new log filterer instance of AToken, bound to a specific deployed contract.
func NewATokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ATokenFilterer, error) {
	contract, err := bindAToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ATokenFilterer{contract: contract}, nil
}

// bindAToken binds a generic wrapper to an already deployed contract.
func bindAToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ATokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AToken *ATokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AToken.Contract.ATokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AToken *ATokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AToken.Contract.ATokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AToken *ATokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AToken.Contract.ATokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AToken *ATokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AToken *ATokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AToken *ATokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AToken.Contract.contract.Transact(opts, method, params...)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_AToken *ATokenCaller) UINTMAXVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "UINT_MAX_VALUE")
	return *ret0, err
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_AToken *ATokenSession) UINTMAXVALUE() (*big.Int, error) {
	return _AToken.Contract.UINTMAXVALUE(&_AToken.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_AToken *ATokenCallerSession) UINTMAXVALUE() (*big.Int, error) {
	return _AToken.Contract.UINTMAXVALUE(&_AToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AToken *ATokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AToken *ATokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AToken.Contract.Allowance(&_AToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AToken *ATokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AToken.Contract.Allowance(&_AToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_AToken *ATokenCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_AToken *ATokenSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.BalanceOf(&_AToken.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_AToken *ATokenCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.BalanceOf(&_AToken.CallOpts, _user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AToken *ATokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AToken *ATokenSession) Decimals() (uint8, error) {
	return _AToken.Contract.Decimals(&_AToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AToken *ATokenCallerSession) Decimals() (uint8, error) {
	return _AToken.Contract.Decimals(&_AToken.CallOpts)
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_AToken *ATokenCaller) GetInterestRedirectionAddress(opts *bind.CallOpts, _user common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "getInterestRedirectionAddress", _user)
	return *ret0, err
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_AToken *ATokenSession) GetInterestRedirectionAddress(_user common.Address) (common.Address, error) {
	return _AToken.Contract.GetInterestRedirectionAddress(&_AToken.CallOpts, _user)
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_AToken *ATokenCallerSession) GetInterestRedirectionAddress(_user common.Address) (common.Address, error) {
	return _AToken.Contract.GetInterestRedirectionAddress(&_AToken.CallOpts, _user)
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_AToken *ATokenCaller) GetRedirectedBalance(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "getRedirectedBalance", _user)
	return *ret0, err
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_AToken *ATokenSession) GetRedirectedBalance(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.GetRedirectedBalance(&_AToken.CallOpts, _user)
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_AToken *ATokenCallerSession) GetRedirectedBalance(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.GetRedirectedBalance(&_AToken.CallOpts, _user)
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_AToken *ATokenCaller) GetUserIndex(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "getUserIndex", _user)
	return *ret0, err
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_AToken *ATokenSession) GetUserIndex(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.GetUserIndex(&_AToken.CallOpts, _user)
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_AToken *ATokenCallerSession) GetUserIndex(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.GetUserIndex(&_AToken.CallOpts, _user)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_AToken *ATokenCaller) IsTransferAllowed(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "isTransferAllowed", _user, _amount)
	return *ret0, err
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_AToken *ATokenSession) IsTransferAllowed(_user common.Address, _amount *big.Int) (bool, error) {
	return _AToken.Contract.IsTransferAllowed(&_AToken.CallOpts, _user, _amount)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_AToken *ATokenCallerSession) IsTransferAllowed(_user common.Address, _amount *big.Int) (bool, error) {
	return _AToken.Contract.IsTransferAllowed(&_AToken.CallOpts, _user, _amount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AToken *ATokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AToken *ATokenSession) Name() (string, error) {
	return _AToken.Contract.Name(&_AToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AToken *ATokenCallerSession) Name() (string, error) {
	return _AToken.Contract.Name(&_AToken.CallOpts)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_AToken *ATokenCaller) PrincipalBalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "principalBalanceOf", _user)
	return *ret0, err
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_AToken *ATokenSession) PrincipalBalanceOf(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.PrincipalBalanceOf(&_AToken.CallOpts, _user)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_AToken *ATokenCallerSession) PrincipalBalanceOf(_user common.Address) (*big.Int, error) {
	return _AToken.Contract.PrincipalBalanceOf(&_AToken.CallOpts, _user)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AToken *ATokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AToken *ATokenSession) Symbol() (string, error) {
	return _AToken.Contract.Symbol(&_AToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AToken *ATokenCallerSession) Symbol() (string, error) {
	return _AToken.Contract.Symbol(&_AToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AToken *ATokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AToken *ATokenSession) TotalSupply() (*big.Int, error) {
	return _AToken.Contract.TotalSupply(&_AToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AToken *ATokenCallerSession) TotalSupply() (*big.Int, error) {
	return _AToken.Contract.TotalSupply(&_AToken.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_AToken *ATokenCaller) UnderlyingAssetAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AToken.contract.Call(opts, out, "underlyingAssetAddress")
	return *ret0, err
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_AToken *ATokenSession) UnderlyingAssetAddress() (common.Address, error) {
	return _AToken.Contract.UnderlyingAssetAddress(&_AToken.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_AToken *ATokenCallerSession) UnderlyingAssetAddress() (common.Address, error) {
	return _AToken.Contract.UnderlyingAssetAddress(&_AToken.CallOpts)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_AToken *ATokenTransactor) AllowInterestRedirectionTo(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "allowInterestRedirectionTo", _to)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_AToken *ATokenSession) AllowInterestRedirectionTo(_to common.Address) (*types.Transaction, error) {
	return _AToken.Contract.AllowInterestRedirectionTo(&_AToken.TransactOpts, _to)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_AToken *ATokenTransactorSession) AllowInterestRedirectionTo(_to common.Address) (*types.Transaction, error) {
	return _AToken.Contract.AllowInterestRedirectionTo(&_AToken.TransactOpts, _to)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AToken *ATokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AToken *ATokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.Approve(&_AToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AToken *ATokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.Approve(&_AToken.TransactOpts, spender, value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_AToken *ATokenTransactor) BurnOnLiquidation(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "burnOnLiquidation", _account, _value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_AToken *ATokenSession) BurnOnLiquidation(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.BurnOnLiquidation(&_AToken.TransactOpts, _account, _value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_AToken *ATokenTransactorSession) BurnOnLiquidation(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.BurnOnLiquidation(&_AToken.TransactOpts, _account, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AToken *ATokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AToken *ATokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.DecreaseAllowance(&_AToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AToken *ATokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.DecreaseAllowance(&_AToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AToken *ATokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AToken *ATokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.IncreaseAllowance(&_AToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AToken *ATokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.IncreaseAllowance(&_AToken.TransactOpts, spender, addedValue)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_AToken *ATokenTransactor) MintOnDeposit(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "mintOnDeposit", _account, _amount)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_AToken *ATokenSession) MintOnDeposit(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.MintOnDeposit(&_AToken.TransactOpts, _account, _amount)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_AToken *ATokenTransactorSession) MintOnDeposit(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.MintOnDeposit(&_AToken.TransactOpts, _account, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_AToken *ATokenTransactor) Redeem(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "redeem", _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_AToken *ATokenSession) Redeem(_amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.Redeem(&_AToken.TransactOpts, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_AToken *ATokenTransactorSession) Redeem(_amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.Redeem(&_AToken.TransactOpts, _amount)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_AToken *ATokenTransactor) RedirectInterestStream(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "redirectInterestStream", _to)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_AToken *ATokenSession) RedirectInterestStream(_to common.Address) (*types.Transaction, error) {
	return _AToken.Contract.RedirectInterestStream(&_AToken.TransactOpts, _to)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_AToken *ATokenTransactorSession) RedirectInterestStream(_to common.Address) (*types.Transaction, error) {
	return _AToken.Contract.RedirectInterestStream(&_AToken.TransactOpts, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_AToken *ATokenTransactor) RedirectInterestStreamOf(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "redirectInterestStreamOf", _from, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_AToken *ATokenSession) RedirectInterestStreamOf(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _AToken.Contract.RedirectInterestStreamOf(&_AToken.TransactOpts, _from, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_AToken *ATokenTransactorSession) RedirectInterestStreamOf(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _AToken.Contract.RedirectInterestStreamOf(&_AToken.TransactOpts, _from, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AToken *ATokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AToken *ATokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.Transfer(&_AToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AToken *ATokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.Transfer(&_AToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AToken *ATokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AToken *ATokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.TransferFrom(&_AToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AToken *ATokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.TransferFrom(&_AToken.TransactOpts, sender, recipient, amount)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_AToken *ATokenTransactor) TransferOnLiquidation(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AToken.contract.Transact(opts, "transferOnLiquidation", _from, _to, _value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_AToken *ATokenSession) TransferOnLiquidation(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.TransferOnLiquidation(&_AToken.TransactOpts, _from, _to, _value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_AToken *ATokenTransactorSession) TransferOnLiquidation(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AToken.Contract.TransferOnLiquidation(&_AToken.TransactOpts, _from, _to, _value)
}

// ATokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AToken contract.
type ATokenApprovalIterator struct {
	Event *ATokenApproval // Event containing the contract specifics and raw log

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
func (it *ATokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenApproval)
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
		it.Event = new(ATokenApproval)
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
func (it *ATokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenApproval represents a Approval event raised by the AToken contract.
type ATokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AToken *ATokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ATokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ATokenApprovalIterator{contract: _AToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AToken *ATokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ATokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenApproval)
				if err := _AToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AToken *ATokenFilterer) ParseApproval(log types.Log) (*ATokenApproval, error) {
	event := new(ATokenApproval)
	if err := _AToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenBalanceTransferIterator is returned from FilterBalanceTransfer and is used to iterate over the raw logs and unpacked data for BalanceTransfer events raised by the AToken contract.
type ATokenBalanceTransferIterator struct {
	Event *ATokenBalanceTransfer // Event containing the contract specifics and raw log

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
func (it *ATokenBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenBalanceTransfer)
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
		it.Event = new(ATokenBalanceTransfer)
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
func (it *ATokenBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenBalanceTransfer represents a BalanceTransfer event raised by the AToken contract.
type ATokenBalanceTransfer struct {
	From                common.Address
	To                  common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	ToBalanceIncrease   *big.Int
	FromIndex           *big.Int
	ToIndex             *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBalanceTransfer is a free log retrieval operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_AToken *ATokenFilterer) FilterBalanceTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ATokenBalanceTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "BalanceTransfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenBalanceTransferIterator{contract: _AToken.contract, event: "BalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchBalanceTransfer is a free log subscription operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_AToken *ATokenFilterer) WatchBalanceTransfer(opts *bind.WatchOpts, sink chan<- *ATokenBalanceTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "BalanceTransfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenBalanceTransfer)
				if err := _AToken.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
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

// ParseBalanceTransfer is a log parse operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_AToken *ATokenFilterer) ParseBalanceTransfer(log types.Log) (*ATokenBalanceTransfer, error) {
	event := new(ATokenBalanceTransfer)
	if err := _AToken.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenBurnOnLiquidationIterator is returned from FilterBurnOnLiquidation and is used to iterate over the raw logs and unpacked data for BurnOnLiquidation events raised by the AToken contract.
type ATokenBurnOnLiquidationIterator struct {
	Event *ATokenBurnOnLiquidation // Event containing the contract specifics and raw log

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
func (it *ATokenBurnOnLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenBurnOnLiquidation)
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
		it.Event = new(ATokenBurnOnLiquidation)
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
func (it *ATokenBurnOnLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenBurnOnLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenBurnOnLiquidation represents a BurnOnLiquidation event raised by the AToken contract.
type ATokenBurnOnLiquidation struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBurnOnLiquidation is a free log retrieval operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) FilterBurnOnLiquidation(opts *bind.FilterOpts, _from []common.Address) (*ATokenBurnOnLiquidationIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "BurnOnLiquidation", _fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenBurnOnLiquidationIterator{contract: _AToken.contract, event: "BurnOnLiquidation", logs: logs, sub: sub}, nil
}

// WatchBurnOnLiquidation is a free log subscription operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) WatchBurnOnLiquidation(opts *bind.WatchOpts, sink chan<- *ATokenBurnOnLiquidation, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "BurnOnLiquidation", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenBurnOnLiquidation)
				if err := _AToken.contract.UnpackLog(event, "BurnOnLiquidation", log); err != nil {
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

// ParseBurnOnLiquidation is a log parse operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) ParseBurnOnLiquidation(log types.Log) (*ATokenBurnOnLiquidation, error) {
	event := new(ATokenBurnOnLiquidation)
	if err := _AToken.contract.UnpackLog(event, "BurnOnLiquidation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenInterestRedirectionAllowanceChangedIterator is returned from FilterInterestRedirectionAllowanceChanged and is used to iterate over the raw logs and unpacked data for InterestRedirectionAllowanceChanged events raised by the AToken contract.
type ATokenInterestRedirectionAllowanceChangedIterator struct {
	Event *ATokenInterestRedirectionAllowanceChanged // Event containing the contract specifics and raw log

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
func (it *ATokenInterestRedirectionAllowanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenInterestRedirectionAllowanceChanged)
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
		it.Event = new(ATokenInterestRedirectionAllowanceChanged)
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
func (it *ATokenInterestRedirectionAllowanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenInterestRedirectionAllowanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenInterestRedirectionAllowanceChanged represents a InterestRedirectionAllowanceChanged event raised by the AToken contract.
type ATokenInterestRedirectionAllowanceChanged struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterInterestRedirectionAllowanceChanged is a free log retrieval operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_AToken *ATokenFilterer) FilterInterestRedirectionAllowanceChanged(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ATokenInterestRedirectionAllowanceChangedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "InterestRedirectionAllowanceChanged", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenInterestRedirectionAllowanceChangedIterator{contract: _AToken.contract, event: "InterestRedirectionAllowanceChanged", logs: logs, sub: sub}, nil
}

// WatchInterestRedirectionAllowanceChanged is a free log subscription operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_AToken *ATokenFilterer) WatchInterestRedirectionAllowanceChanged(opts *bind.WatchOpts, sink chan<- *ATokenInterestRedirectionAllowanceChanged, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "InterestRedirectionAllowanceChanged", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenInterestRedirectionAllowanceChanged)
				if err := _AToken.contract.UnpackLog(event, "InterestRedirectionAllowanceChanged", log); err != nil {
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

// ParseInterestRedirectionAllowanceChanged is a log parse operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_AToken *ATokenFilterer) ParseInterestRedirectionAllowanceChanged(log types.Log) (*ATokenInterestRedirectionAllowanceChanged, error) {
	event := new(ATokenInterestRedirectionAllowanceChanged)
	if err := _AToken.contract.UnpackLog(event, "InterestRedirectionAllowanceChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenInterestStreamRedirectedIterator is returned from FilterInterestStreamRedirected and is used to iterate over the raw logs and unpacked data for InterestStreamRedirected events raised by the AToken contract.
type ATokenInterestStreamRedirectedIterator struct {
	Event *ATokenInterestStreamRedirected // Event containing the contract specifics and raw log

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
func (it *ATokenInterestStreamRedirectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenInterestStreamRedirected)
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
		it.Event = new(ATokenInterestStreamRedirected)
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
func (it *ATokenInterestStreamRedirectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenInterestStreamRedirectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenInterestStreamRedirected represents a InterestStreamRedirected event raised by the AToken contract.
type ATokenInterestStreamRedirected struct {
	From                common.Address
	To                  common.Address
	RedirectedBalance   *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInterestStreamRedirected is a free log retrieval operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) FilterInterestStreamRedirected(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ATokenInterestStreamRedirectedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "InterestStreamRedirected", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenInterestStreamRedirectedIterator{contract: _AToken.contract, event: "InterestStreamRedirected", logs: logs, sub: sub}, nil
}

// WatchInterestStreamRedirected is a free log subscription operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) WatchInterestStreamRedirected(opts *bind.WatchOpts, sink chan<- *ATokenInterestStreamRedirected, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "InterestStreamRedirected", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenInterestStreamRedirected)
				if err := _AToken.contract.UnpackLog(event, "InterestStreamRedirected", log); err != nil {
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

// ParseInterestStreamRedirected is a log parse operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) ParseInterestStreamRedirected(log types.Log) (*ATokenInterestStreamRedirected, error) {
	event := new(ATokenInterestStreamRedirected)
	if err := _AToken.contract.UnpackLog(event, "InterestStreamRedirected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenMintOnDepositIterator is returned from FilterMintOnDeposit and is used to iterate over the raw logs and unpacked data for MintOnDeposit events raised by the AToken contract.
type ATokenMintOnDepositIterator struct {
	Event *ATokenMintOnDeposit // Event containing the contract specifics and raw log

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
func (it *ATokenMintOnDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenMintOnDeposit)
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
		it.Event = new(ATokenMintOnDeposit)
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
func (it *ATokenMintOnDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenMintOnDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenMintOnDeposit represents a MintOnDeposit event raised by the AToken contract.
type ATokenMintOnDeposit struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMintOnDeposit is a free log retrieval operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) FilterMintOnDeposit(opts *bind.FilterOpts, _from []common.Address) (*ATokenMintOnDepositIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenMintOnDepositIterator{contract: _AToken.contract, event: "MintOnDeposit", logs: logs, sub: sub}, nil
}

// WatchMintOnDeposit is a free log subscription operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) WatchMintOnDeposit(opts *bind.WatchOpts, sink chan<- *ATokenMintOnDeposit, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenMintOnDeposit)
				if err := _AToken.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
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

// ParseMintOnDeposit is a log parse operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) ParseMintOnDeposit(log types.Log) (*ATokenMintOnDeposit, error) {
	event := new(ATokenMintOnDeposit)
	if err := _AToken.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the AToken contract.
type ATokenRedeemIterator struct {
	Event *ATokenRedeem // Event containing the contract specifics and raw log

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
func (it *ATokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenRedeem)
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
		it.Event = new(ATokenRedeem)
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
func (it *ATokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenRedeem represents a Redeem event raised by the AToken contract.
type ATokenRedeem struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) FilterRedeem(opts *bind.FilterOpts, _from []common.Address) (*ATokenRedeemIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "Redeem", _fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenRedeemIterator{contract: _AToken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *ATokenRedeem, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "Redeem", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenRedeem)
				if err := _AToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_AToken *ATokenFilterer) ParseRedeem(log types.Log) (*ATokenRedeem, error) {
	event := new(ATokenRedeem)
	if err := _AToken.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenRedirectedBalanceUpdatedIterator is returned from FilterRedirectedBalanceUpdated and is used to iterate over the raw logs and unpacked data for RedirectedBalanceUpdated events raised by the AToken contract.
type ATokenRedirectedBalanceUpdatedIterator struct {
	Event *ATokenRedirectedBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *ATokenRedirectedBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenRedirectedBalanceUpdated)
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
		it.Event = new(ATokenRedirectedBalanceUpdated)
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
func (it *ATokenRedirectedBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenRedirectedBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenRedirectedBalanceUpdated represents a RedirectedBalanceUpdated event raised by the AToken contract.
type ATokenRedirectedBalanceUpdated struct {
	TargetAddress            common.Address
	TargetBalanceIncrease    *big.Int
	TargetIndex              *big.Int
	RedirectedBalanceAdded   *big.Int
	RedirectedBalanceRemoved *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRedirectedBalanceUpdated is a free log retrieval operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_AToken *ATokenFilterer) FilterRedirectedBalanceUpdated(opts *bind.FilterOpts, _targetAddress []common.Address) (*ATokenRedirectedBalanceUpdatedIterator, error) {

	var _targetAddressRule []interface{}
	for _, _targetAddressItem := range _targetAddress {
		_targetAddressRule = append(_targetAddressRule, _targetAddressItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "RedirectedBalanceUpdated", _targetAddressRule)
	if err != nil {
		return nil, err
	}
	return &ATokenRedirectedBalanceUpdatedIterator{contract: _AToken.contract, event: "RedirectedBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchRedirectedBalanceUpdated is a free log subscription operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_AToken *ATokenFilterer) WatchRedirectedBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ATokenRedirectedBalanceUpdated, _targetAddress []common.Address) (event.Subscription, error) {

	var _targetAddressRule []interface{}
	for _, _targetAddressItem := range _targetAddress {
		_targetAddressRule = append(_targetAddressRule, _targetAddressItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "RedirectedBalanceUpdated", _targetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenRedirectedBalanceUpdated)
				if err := _AToken.contract.UnpackLog(event, "RedirectedBalanceUpdated", log); err != nil {
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

// ParseRedirectedBalanceUpdated is a log parse operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_AToken *ATokenFilterer) ParseRedirectedBalanceUpdated(log types.Log) (*ATokenRedirectedBalanceUpdated, error) {
	event := new(ATokenRedirectedBalanceUpdated)
	if err := _AToken.contract.UnpackLog(event, "RedirectedBalanceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ATokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AToken contract.
type ATokenTransferIterator struct {
	Event *ATokenTransfer // Event containing the contract specifics and raw log

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
func (it *ATokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenTransfer)
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
		it.Event = new(ATokenTransfer)
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
func (it *ATokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenTransfer represents a Transfer event raised by the AToken contract.
type ATokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AToken *ATokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ATokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenTransferIterator{contract: _AToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AToken *ATokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ATokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenTransfer)
				if err := _AToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AToken *ATokenFilterer) ParseTransfer(log types.Log) (*ATokenTransfer, error) {
	event := new(ATokenTransfer)
	if err := _AToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
