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

// AaveV2WethGatewayABI is the input ABI used to generate the binding from.
const AaveV2WethGatewayABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingPool\",\"type\":\"address\"}],\"name\":\"authorizeLendingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interesRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"borrowETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyEtherTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWETHAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repayETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// AaveV2WethGateway is an auto generated Go binding around an Ethereum contract.
type AaveV2WethGateway struct {
	AaveV2WethGatewayCaller     // Read-only binding to the contract
	AaveV2WethGatewayTransactor // Write-only binding to the contract
	AaveV2WethGatewayFilterer   // Log filterer for contract events
}

// AaveV2WethGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV2WethGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2WethGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV2WethGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2WethGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV2WethGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2WethGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV2WethGatewaySession struct {
	Contract     *AaveV2WethGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AaveV2WethGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV2WethGatewayCallerSession struct {
	Contract *AaveV2WethGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AaveV2WethGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV2WethGatewayTransactorSession struct {
	Contract     *AaveV2WethGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AaveV2WethGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV2WethGatewayRaw struct {
	Contract *AaveV2WethGateway // Generic contract binding to access the raw methods on
}

// AaveV2WethGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV2WethGatewayCallerRaw struct {
	Contract *AaveV2WethGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV2WethGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV2WethGatewayTransactorRaw struct {
	Contract *AaveV2WethGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV2WethGateway creates a new instance of AaveV2WethGateway, bound to a specific deployed contract.
func NewAaveV2WethGateway(address common.Address, backend bind.ContractBackend) (*AaveV2WethGateway, error) {
	contract, err := bindAaveV2WethGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV2WethGateway{AaveV2WethGatewayCaller: AaveV2WethGatewayCaller{contract: contract}, AaveV2WethGatewayTransactor: AaveV2WethGatewayTransactor{contract: contract}, AaveV2WethGatewayFilterer: AaveV2WethGatewayFilterer{contract: contract}}, nil
}

// NewAaveV2WethGatewayCaller creates a new read-only instance of AaveV2WethGateway, bound to a specific deployed contract.
func NewAaveV2WethGatewayCaller(address common.Address, caller bind.ContractCaller) (*AaveV2WethGatewayCaller, error) {
	contract, err := bindAaveV2WethGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2WethGatewayCaller{contract: contract}, nil
}

// NewAaveV2WethGatewayTransactor creates a new write-only instance of AaveV2WethGateway, bound to a specific deployed contract.
func NewAaveV2WethGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV2WethGatewayTransactor, error) {
	contract, err := bindAaveV2WethGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2WethGatewayTransactor{contract: contract}, nil
}

// NewAaveV2WethGatewayFilterer creates a new log filterer instance of AaveV2WethGateway, bound to a specific deployed contract.
func NewAaveV2WethGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV2WethGatewayFilterer, error) {
	contract, err := bindAaveV2WethGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV2WethGatewayFilterer{contract: contract}, nil
}

// bindAaveV2WethGateway binds a generic wrapper to an already deployed contract.
func bindAaveV2WethGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveV2WethGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2WethGateway *AaveV2WethGatewayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AaveV2WethGateway.Contract.AaveV2WethGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2WethGateway *AaveV2WethGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.AaveV2WethGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2WethGateway *AaveV2WethGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.AaveV2WethGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2WethGateway *AaveV2WethGatewayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AaveV2WethGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.contract.Transact(opts, method, params...)
}

// GetWETHAddress is a free data retrieval call binding the contract method 0xaffa8817.
//
// Solidity: function getWETHAddress() view returns(address)
func (_AaveV2WethGateway *AaveV2WethGatewayCaller) GetWETHAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AaveV2WethGateway.contract.Call(opts, out, "getWETHAddress")
	return *ret0, err
}

// GetWETHAddress is a free data retrieval call binding the contract method 0xaffa8817.
//
// Solidity: function getWETHAddress() view returns(address)
func (_AaveV2WethGateway *AaveV2WethGatewaySession) GetWETHAddress() (common.Address, error) {
	return _AaveV2WethGateway.Contract.GetWETHAddress(&_AaveV2WethGateway.CallOpts)
}

// GetWETHAddress is a free data retrieval call binding the contract method 0xaffa8817.
//
// Solidity: function getWETHAddress() view returns(address)
func (_AaveV2WethGateway *AaveV2WethGatewayCallerSession) GetWETHAddress() (common.Address, error) {
	return _AaveV2WethGateway.Contract.GetWETHAddress(&_AaveV2WethGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV2WethGateway *AaveV2WethGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AaveV2WethGateway.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV2WethGateway *AaveV2WethGatewaySession) Owner() (common.Address, error) {
	return _AaveV2WethGateway.Contract.Owner(&_AaveV2WethGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV2WethGateway *AaveV2WethGatewayCallerSession) Owner() (common.Address, error) {
	return _AaveV2WethGateway.Contract.Owner(&_AaveV2WethGateway.CallOpts)
}

// AuthorizeLendingPool is a paid mutator transaction binding the contract method 0xfd149529.
//
// Solidity: function authorizeLendingPool(address lendingPool) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) AuthorizeLendingPool(opts *bind.TransactOpts, lendingPool common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "authorizeLendingPool", lendingPool)
}

// AuthorizeLendingPool is a paid mutator transaction binding the contract method 0xfd149529.
//
// Solidity: function authorizeLendingPool(address lendingPool) returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) AuthorizeLendingPool(lendingPool common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.AuthorizeLendingPool(&_AaveV2WethGateway.TransactOpts, lendingPool)
}

// AuthorizeLendingPool is a paid mutator transaction binding the contract method 0xfd149529.
//
// Solidity: function authorizeLendingPool(address lendingPool) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) AuthorizeLendingPool(lendingPool common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.AuthorizeLendingPool(&_AaveV2WethGateway.TransactOpts, lendingPool)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address lendingPool, uint256 amount, uint256 interesRateMode, uint16 referralCode) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) BorrowETH(opts *bind.TransactOpts, lendingPool common.Address, amount *big.Int, interesRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "borrowETH", lendingPool, amount, interesRateMode, referralCode)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address lendingPool, uint256 amount, uint256 interesRateMode, uint16 referralCode) returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) BorrowETH(lendingPool common.Address, amount *big.Int, interesRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.BorrowETH(&_AaveV2WethGateway.TransactOpts, lendingPool, amount, interesRateMode, referralCode)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address lendingPool, uint256 amount, uint256 interesRateMode, uint16 referralCode) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) BorrowETH(lendingPool common.Address, amount *big.Int, interesRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.BorrowETH(&_AaveV2WethGateway.TransactOpts, lendingPool, amount, interesRateMode, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address lendingPool, address onBehalfOf, uint16 referralCode) payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) DepositETH(opts *bind.TransactOpts, lendingPool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "depositETH", lendingPool, onBehalfOf, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address lendingPool, address onBehalfOf, uint16 referralCode) payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) DepositETH(lendingPool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.DepositETH(&_AaveV2WethGateway.TransactOpts, lendingPool, onBehalfOf, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address lendingPool, address onBehalfOf, uint16 referralCode) payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) DepositETH(lendingPool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.DepositETH(&_AaveV2WethGateway.TransactOpts, lendingPool, onBehalfOf, referralCode)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) EmergencyEtherTransfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "emergencyEtherTransfer", to, amount)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) EmergencyEtherTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.EmergencyEtherTransfer(&_AaveV2WethGateway.TransactOpts, to, amount)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) EmergencyEtherTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.EmergencyEtherTransfer(&_AaveV2WethGateway.TransactOpts, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) EmergencyTokenTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "emergencyTokenTransfer", token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.EmergencyTokenTransfer(&_AaveV2WethGateway.TransactOpts, token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.EmergencyTokenTransfer(&_AaveV2WethGateway.TransactOpts, token, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.RenounceOwnership(&_AaveV2WethGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.RenounceOwnership(&_AaveV2WethGateway.TransactOpts)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address lendingPool, uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) RepayETH(opts *bind.TransactOpts, lendingPool common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "repayETH", lendingPool, amount, rateMode, onBehalfOf)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address lendingPool, uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) RepayETH(lendingPool common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.RepayETH(&_AaveV2WethGateway.TransactOpts, lendingPool, amount, rateMode, onBehalfOf)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address lendingPool, uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) RepayETH(lendingPool common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.RepayETH(&_AaveV2WethGateway.TransactOpts, lendingPool, amount, rateMode, onBehalfOf)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.TransferOwnership(&_AaveV2WethGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.TransferOwnership(&_AaveV2WethGateway.TransactOpts, newOwner)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address lendingPool, uint256 amount, address to) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) WithdrawETH(opts *bind.TransactOpts, lendingPool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.Transact(opts, "withdrawETH", lendingPool, amount, to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address lendingPool, uint256 amount, address to) returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) WithdrawETH(lendingPool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.WithdrawETH(&_AaveV2WethGateway.TransactOpts, lendingPool, amount, to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address lendingPool, uint256 amount, address to) returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) WithdrawETH(lendingPool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.WithdrawETH(&_AaveV2WethGateway.TransactOpts, lendingPool, amount, to)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.Fallback(&_AaveV2WethGateway.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.Fallback(&_AaveV2WethGateway.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2WethGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewaySession) Receive() (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.Receive(&_AaveV2WethGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AaveV2WethGateway *AaveV2WethGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _AaveV2WethGateway.Contract.Receive(&_AaveV2WethGateway.TransactOpts)
}

// AaveV2WethGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AaveV2WethGateway contract.
type AaveV2WethGatewayOwnershipTransferredIterator struct {
	Event *AaveV2WethGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AaveV2WethGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2WethGatewayOwnershipTransferred)
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
		it.Event = new(AaveV2WethGatewayOwnershipTransferred)
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
func (it *AaveV2WethGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2WethGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2WethGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the AaveV2WethGateway contract.
type AaveV2WethGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV2WethGateway *AaveV2WethGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AaveV2WethGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV2WethGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2WethGatewayOwnershipTransferredIterator{contract: _AaveV2WethGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV2WethGateway *AaveV2WethGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AaveV2WethGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV2WethGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2WethGatewayOwnershipTransferred)
				if err := _AaveV2WethGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV2WethGateway *AaveV2WethGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*AaveV2WethGatewayOwnershipTransferred, error) {
	event := new(AaveV2WethGatewayOwnershipTransferred)
	if err := _AaveV2WethGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
