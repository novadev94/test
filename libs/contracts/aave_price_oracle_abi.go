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

// AavePriceOracleABI is the input ABI used to generate the binding from.
const AavePriceOracleABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AssetSourceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"FallbackOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"name\":\"WethSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"getAssetsPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getSourceOfAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"}],\"name\":\"setAssetSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"setFallbackOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AavePriceOracle is an auto generated Go binding around an Ethereum contract.
type AavePriceOracle struct {
	AavePriceOracleCaller     // Read-only binding to the contract
	AavePriceOracleTransactor // Write-only binding to the contract
	AavePriceOracleFilterer   // Log filterer for contract events
}

// AavePriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AavePriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AavePriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AavePriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AavePriceOracleSession struct {
	Contract     *AavePriceOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AavePriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AavePriceOracleCallerSession struct {
	Contract *AavePriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AavePriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AavePriceOracleTransactorSession struct {
	Contract     *AavePriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AavePriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AavePriceOracleRaw struct {
	Contract *AavePriceOracle // Generic contract binding to access the raw methods on
}

// AavePriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AavePriceOracleCallerRaw struct {
	Contract *AavePriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// AavePriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AavePriceOracleTransactorRaw struct {
	Contract *AavePriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAavePriceOracle creates a new instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracle(address common.Address, backend bind.ContractBackend) (*AavePriceOracle, error) {
	contract, err := bindAavePriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracle{AavePriceOracleCaller: AavePriceOracleCaller{contract: contract}, AavePriceOracleTransactor: AavePriceOracleTransactor{contract: contract}, AavePriceOracleFilterer: AavePriceOracleFilterer{contract: contract}}, nil
}

// NewAavePriceOracleCaller creates a new read-only instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracleCaller(address common.Address, caller bind.ContractCaller) (*AavePriceOracleCaller, error) {
	contract, err := bindAavePriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleCaller{contract: contract}, nil
}

// NewAavePriceOracleTransactor creates a new write-only instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*AavePriceOracleTransactor, error) {
	contract, err := bindAavePriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleTransactor{contract: contract}, nil
}

// NewAavePriceOracleFilterer creates a new log filterer instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*AavePriceOracleFilterer, error) {
	contract, err := bindAavePriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleFilterer{contract: contract}, nil
}

// bindAavePriceOracle binds a generic wrapper to an already deployed contract.
func bindAavePriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AavePriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePriceOracle *AavePriceOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AavePriceOracle.Contract.AavePriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePriceOracle *AavePriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.AavePriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePriceOracle *AavePriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.AavePriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePriceOracle *AavePriceOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AavePriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePriceOracle *AavePriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePriceOracle *AavePriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AavePriceOracle *AavePriceOracleCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AavePriceOracle.contract.Call(opts, out, "WETH")
	return *ret0, err
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AavePriceOracle *AavePriceOracleSession) WETH() (common.Address, error) {
	return _AavePriceOracle.Contract.WETH(&_AavePriceOracle.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AavePriceOracle *AavePriceOracleCallerSession) WETH() (common.Address, error) {
	return _AavePriceOracle.Contract.WETH(&_AavePriceOracle.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AavePriceOracle *AavePriceOracleCaller) GetAssetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AavePriceOracle.contract.Call(opts, out, "getAssetPrice", asset)
	return *ret0, err
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AavePriceOracle *AavePriceOracleSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _AavePriceOracle.Contract.GetAssetPrice(&_AavePriceOracle.CallOpts, asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AavePriceOracle *AavePriceOracleCallerSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _AavePriceOracle.Contract.GetAssetPrice(&_AavePriceOracle.CallOpts, asset)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AavePriceOracle *AavePriceOracleCaller) GetAssetsPrices(opts *bind.CallOpts, assets []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _AavePriceOracle.contract.Call(opts, out, "getAssetsPrices", assets)
	return *ret0, err
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AavePriceOracle *AavePriceOracleSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _AavePriceOracle.Contract.GetAssetsPrices(&_AavePriceOracle.CallOpts, assets)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AavePriceOracle *AavePriceOracleCallerSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _AavePriceOracle.Contract.GetAssetsPrices(&_AavePriceOracle.CallOpts, assets)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AavePriceOracle *AavePriceOracleCaller) GetFallbackOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AavePriceOracle.contract.Call(opts, out, "getFallbackOracle")
	return *ret0, err
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AavePriceOracle *AavePriceOracleSession) GetFallbackOracle() (common.Address, error) {
	return _AavePriceOracle.Contract.GetFallbackOracle(&_AavePriceOracle.CallOpts)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AavePriceOracle *AavePriceOracleCallerSession) GetFallbackOracle() (common.Address, error) {
	return _AavePriceOracle.Contract.GetFallbackOracle(&_AavePriceOracle.CallOpts)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AavePriceOracle *AavePriceOracleCaller) GetSourceOfAsset(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AavePriceOracle.contract.Call(opts, out, "getSourceOfAsset", asset)
	return *ret0, err
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AavePriceOracle *AavePriceOracleSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _AavePriceOracle.Contract.GetSourceOfAsset(&_AavePriceOracle.CallOpts, asset)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AavePriceOracle *AavePriceOracleCallerSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _AavePriceOracle.Contract.GetSourceOfAsset(&_AavePriceOracle.CallOpts, asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AavePriceOracle *AavePriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AavePriceOracle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AavePriceOracle *AavePriceOracleSession) Owner() (common.Address, error) {
	return _AavePriceOracle.Contract.Owner(&_AavePriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AavePriceOracle *AavePriceOracleCallerSession) Owner() (common.Address, error) {
	return _AavePriceOracle.Contract.Owner(&_AavePriceOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AavePriceOracle *AavePriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AavePriceOracle *AavePriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _AavePriceOracle.Contract.RenounceOwnership(&_AavePriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AavePriceOracle *AavePriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AavePriceOracle.Contract.RenounceOwnership(&_AavePriceOracle.TransactOpts)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AavePriceOracle *AavePriceOracleTransactor) SetAssetSources(opts *bind.TransactOpts, assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.contract.Transact(opts, "setAssetSources", assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AavePriceOracle *AavePriceOracleSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetAssetSources(&_AavePriceOracle.TransactOpts, assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AavePriceOracle *AavePriceOracleTransactorSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetAssetSources(&_AavePriceOracle.TransactOpts, assets, sources)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AavePriceOracle *AavePriceOracleTransactor) SetFallbackOracle(opts *bind.TransactOpts, fallbackOracle common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.contract.Transact(opts, "setFallbackOracle", fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AavePriceOracle *AavePriceOracleSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetFallbackOracle(&_AavePriceOracle.TransactOpts, fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AavePriceOracle *AavePriceOracleTransactorSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetFallbackOracle(&_AavePriceOracle.TransactOpts, fallbackOracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AavePriceOracle *AavePriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AavePriceOracle *AavePriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.TransferOwnership(&_AavePriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AavePriceOracle *AavePriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.TransferOwnership(&_AavePriceOracle.TransactOpts, newOwner)
}

// AavePriceOracleAssetSourceUpdatedIterator is returned from FilterAssetSourceUpdated and is used to iterate over the raw logs and unpacked data for AssetSourceUpdated events raised by the AavePriceOracle contract.
type AavePriceOracleAssetSourceUpdatedIterator struct {
	Event *AavePriceOracleAssetSourceUpdated // Event containing the contract specifics and raw log

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
func (it *AavePriceOracleAssetSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePriceOracleAssetSourceUpdated)
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
		it.Event = new(AavePriceOracleAssetSourceUpdated)
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
func (it *AavePriceOracleAssetSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePriceOracleAssetSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePriceOracleAssetSourceUpdated represents a AssetSourceUpdated event raised by the AavePriceOracle contract.
type AavePriceOracleAssetSourceUpdated struct {
	Asset  common.Address
	Source common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetSourceUpdated is a free log retrieval operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_AavePriceOracle *AavePriceOracleFilterer) FilterAssetSourceUpdated(opts *bind.FilterOpts, asset []common.Address, source []common.Address) (*AavePriceOracleAssetSourceUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _AavePriceOracle.contract.FilterLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleAssetSourceUpdatedIterator{contract: _AavePriceOracle.contract, event: "AssetSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetSourceUpdated is a free log subscription operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_AavePriceOracle *AavePriceOracleFilterer) WatchAssetSourceUpdated(opts *bind.WatchOpts, sink chan<- *AavePriceOracleAssetSourceUpdated, asset []common.Address, source []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _AavePriceOracle.contract.WatchLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePriceOracleAssetSourceUpdated)
				if err := _AavePriceOracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
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

// ParseAssetSourceUpdated is a log parse operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_AavePriceOracle *AavePriceOracleFilterer) ParseAssetSourceUpdated(log types.Log) (*AavePriceOracleAssetSourceUpdated, error) {
	event := new(AavePriceOracleAssetSourceUpdated)
	if err := _AavePriceOracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AavePriceOracleFallbackOracleUpdatedIterator is returned from FilterFallbackOracleUpdated and is used to iterate over the raw logs and unpacked data for FallbackOracleUpdated events raised by the AavePriceOracle contract.
type AavePriceOracleFallbackOracleUpdatedIterator struct {
	Event *AavePriceOracleFallbackOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AavePriceOracleFallbackOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePriceOracleFallbackOracleUpdated)
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
		it.Event = new(AavePriceOracleFallbackOracleUpdated)
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
func (it *AavePriceOracleFallbackOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePriceOracleFallbackOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePriceOracleFallbackOracleUpdated represents a FallbackOracleUpdated event raised by the AavePriceOracle contract.
type AavePriceOracleFallbackOracleUpdated struct {
	FallbackOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFallbackOracleUpdated is a free log retrieval operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_AavePriceOracle *AavePriceOracleFilterer) FilterFallbackOracleUpdated(opts *bind.FilterOpts, fallbackOracle []common.Address) (*AavePriceOracleFallbackOracleUpdatedIterator, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _AavePriceOracle.contract.FilterLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleFallbackOracleUpdatedIterator{contract: _AavePriceOracle.contract, event: "FallbackOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchFallbackOracleUpdated is a free log subscription operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_AavePriceOracle *AavePriceOracleFilterer) WatchFallbackOracleUpdated(opts *bind.WatchOpts, sink chan<- *AavePriceOracleFallbackOracleUpdated, fallbackOracle []common.Address) (event.Subscription, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _AavePriceOracle.contract.WatchLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePriceOracleFallbackOracleUpdated)
				if err := _AavePriceOracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
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

// ParseFallbackOracleUpdated is a log parse operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_AavePriceOracle *AavePriceOracleFilterer) ParseFallbackOracleUpdated(log types.Log) (*AavePriceOracleFallbackOracleUpdated, error) {
	event := new(AavePriceOracleFallbackOracleUpdated)
	if err := _AavePriceOracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AavePriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AavePriceOracle contract.
type AavePriceOracleOwnershipTransferredIterator struct {
	Event *AavePriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AavePriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePriceOracleOwnershipTransferred)
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
		it.Event = new(AavePriceOracleOwnershipTransferred)
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
func (it *AavePriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the AavePriceOracle contract.
type AavePriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AavePriceOracle *AavePriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AavePriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AavePriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleOwnershipTransferredIterator{contract: _AavePriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AavePriceOracle *AavePriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AavePriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AavePriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePriceOracleOwnershipTransferred)
				if err := _AavePriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AavePriceOracle *AavePriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*AavePriceOracleOwnershipTransferred, error) {
	event := new(AavePriceOracleOwnershipTransferred)
	if err := _AavePriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AavePriceOracleWethSetIterator is returned from FilterWethSet and is used to iterate over the raw logs and unpacked data for WethSet events raised by the AavePriceOracle contract.
type AavePriceOracleWethSetIterator struct {
	Event *AavePriceOracleWethSet // Event containing the contract specifics and raw log

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
func (it *AavePriceOracleWethSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePriceOracleWethSet)
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
		it.Event = new(AavePriceOracleWethSet)
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
func (it *AavePriceOracleWethSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePriceOracleWethSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePriceOracleWethSet represents a WethSet event raised by the AavePriceOracle contract.
type AavePriceOracleWethSet struct {
	Weth common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWethSet is a free log retrieval operation binding the contract event 0x13a533084dcbb1cfe0dbea708ea977223c27c44d94f2fa3867a167c9cd340bf9.
//
// Solidity: event WethSet(address indexed weth)
func (_AavePriceOracle *AavePriceOracleFilterer) FilterWethSet(opts *bind.FilterOpts, weth []common.Address) (*AavePriceOracleWethSetIterator, error) {

	var wethRule []interface{}
	for _, wethItem := range weth {
		wethRule = append(wethRule, wethItem)
	}

	logs, sub, err := _AavePriceOracle.contract.FilterLogs(opts, "WethSet", wethRule)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleWethSetIterator{contract: _AavePriceOracle.contract, event: "WethSet", logs: logs, sub: sub}, nil
}

// WatchWethSet is a free log subscription operation binding the contract event 0x13a533084dcbb1cfe0dbea708ea977223c27c44d94f2fa3867a167c9cd340bf9.
//
// Solidity: event WethSet(address indexed weth)
func (_AavePriceOracle *AavePriceOracleFilterer) WatchWethSet(opts *bind.WatchOpts, sink chan<- *AavePriceOracleWethSet, weth []common.Address) (event.Subscription, error) {

	var wethRule []interface{}
	for _, wethItem := range weth {
		wethRule = append(wethRule, wethItem)
	}

	logs, sub, err := _AavePriceOracle.contract.WatchLogs(opts, "WethSet", wethRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePriceOracleWethSet)
				if err := _AavePriceOracle.contract.UnpackLog(event, "WethSet", log); err != nil {
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

// ParseWethSet is a log parse operation binding the contract event 0x13a533084dcbb1cfe0dbea708ea977223c27c44d94f2fa3867a167c9cd340bf9.
//
// Solidity: event WethSet(address indexed weth)
func (_AavePriceOracle *AavePriceOracleFilterer) ParseWethSet(log types.Log) (*AavePriceOracleWethSet, error) {
	event := new(AavePriceOracleWethSet)
	if err := _AavePriceOracle.contract.UnpackLog(event, "WethSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
