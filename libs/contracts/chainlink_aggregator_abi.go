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

// ChainlinkAggregatorABI is the input ABI used to generate the binding from.
const ChainlinkAggregatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accessController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"confirmAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"phaseAggregators\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phaseId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"proposeAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedAggregator\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"proposedGetRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedLatestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accessController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChainlinkAggregator is an auto generated Go binding around an Ethereum contract.
type ChainlinkAggregator struct {
	ChainlinkAggregatorCaller     // Read-only binding to the contract
	ChainlinkAggregatorTransactor // Write-only binding to the contract
	ChainlinkAggregatorFilterer   // Log filterer for contract events
}

// ChainlinkAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkAggregatorSession struct {
	Contract     *ChainlinkAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChainlinkAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkAggregatorCallerSession struct {
	Contract *ChainlinkAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ChainlinkAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkAggregatorTransactorSession struct {
	Contract     *ChainlinkAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ChainlinkAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkAggregatorRaw struct {
	Contract *ChainlinkAggregator // Generic contract binding to access the raw methods on
}

// ChainlinkAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkAggregatorCallerRaw struct {
	Contract *ChainlinkAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkAggregatorTransactorRaw struct {
	Contract *ChainlinkAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkAggregator creates a new instance of ChainlinkAggregator, bound to a specific deployed contract.
func NewChainlinkAggregator(address common.Address, backend bind.ContractBackend) (*ChainlinkAggregator, error) {
	contract, err := bindChainlinkAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregator{ChainlinkAggregatorCaller: ChainlinkAggregatorCaller{contract: contract}, ChainlinkAggregatorTransactor: ChainlinkAggregatorTransactor{contract: contract}, ChainlinkAggregatorFilterer: ChainlinkAggregatorFilterer{contract: contract}}, nil
}

// NewChainlinkAggregatorCaller creates a new read-only instance of ChainlinkAggregator, bound to a specific deployed contract.
func NewChainlinkAggregatorCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkAggregatorCaller, error) {
	contract, err := bindChainlinkAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregatorCaller{contract: contract}, nil
}

// NewChainlinkAggregatorTransactor creates a new write-only instance of ChainlinkAggregator, bound to a specific deployed contract.
func NewChainlinkAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkAggregatorTransactor, error) {
	contract, err := bindChainlinkAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregatorTransactor{contract: contract}, nil
}

// NewChainlinkAggregatorFilterer creates a new log filterer instance of ChainlinkAggregator, bound to a specific deployed contract.
func NewChainlinkAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkAggregatorFilterer, error) {
	contract, err := bindChainlinkAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregatorFilterer{contract: contract}, nil
}

// bindChainlinkAggregator binds a generic wrapper to an already deployed contract.
func bindChainlinkAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkAggregator *ChainlinkAggregatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkAggregator.Contract.ChainlinkAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkAggregator *ChainlinkAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.ChainlinkAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkAggregator *ChainlinkAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.ChainlinkAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkAggregator *ChainlinkAggregatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkAggregator *ChainlinkAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkAggregator *ChainlinkAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.contract.Transact(opts, method, params...)
}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) AccessController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "accessController")
	return *ret0, err
}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) AccessController() (common.Address, error) {
	return _ChainlinkAggregator.Contract.AccessController(&_ChainlinkAggregator.CallOpts)
}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) AccessController() (common.Address, error) {
	return _ChainlinkAggregator.Contract.AccessController(&_ChainlinkAggregator.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "aggregator")
	return *ret0, err
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) Aggregator() (common.Address, error) {
	return _ChainlinkAggregator.Contract.Aggregator(&_ChainlinkAggregator.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) Aggregator() (common.Address, error) {
	return _ChainlinkAggregator.Contract.Aggregator(&_ChainlinkAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) Decimals() (uint8, error) {
	return _ChainlinkAggregator.Contract.Decimals(&_ChainlinkAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) Decimals() (uint8, error) {
	return _ChainlinkAggregator.Contract.Decimals(&_ChainlinkAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "description")
	return *ret0, err
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) Description() (string, error) {
	return _ChainlinkAggregator.Contract.Description(&_ChainlinkAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) Description() (string, error) {
	return _ChainlinkAggregator.Contract.Description(&_ChainlinkAggregator.CallOpts)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "getAnswer", _roundId)
	return *ret0, err
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkAggregator.Contract.GetAnswer(&_ChainlinkAggregator.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkAggregator.Contract.GetAnswer(&_ChainlinkAggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	ret := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	out := ret
	err := _ChainlinkAggregator.contract.Call(opts, out, "getRoundData", _roundId)
	return *ret, err
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.GetRoundData(&_ChainlinkAggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.GetRoundData(&_ChainlinkAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "getTimestamp", _roundId)
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkAggregator.Contract.GetTimestamp(&_ChainlinkAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkAggregator.Contract.GetTimestamp(&_ChainlinkAggregator.CallOpts, _roundId)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "latestAnswer")
	return *ret0, err
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.LatestAnswer(&_ChainlinkAggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.LatestAnswer(&_ChainlinkAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "latestRound")
	return *ret0, err
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) LatestRound() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.LatestRound(&_ChainlinkAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.LatestRound(&_ChainlinkAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	ret := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	out := ret
	err := _ChainlinkAggregator.contract.Call(opts, out, "latestRoundData")
	return *ret, err
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.LatestRoundData(&_ChainlinkAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.LatestRoundData(&_ChainlinkAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "latestTimestamp")
	return *ret0, err
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.LatestTimestamp(&_ChainlinkAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.LatestTimestamp(&_ChainlinkAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) Owner() (common.Address, error) {
	return _ChainlinkAggregator.Contract.Owner(&_ChainlinkAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) Owner() (common.Address, error) {
	return _ChainlinkAggregator.Contract.Owner(&_ChainlinkAggregator.CallOpts)
}

// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
//
// Solidity: function phaseAggregators(uint16 ) view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) PhaseAggregators(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "phaseAggregators", arg0)
	return *ret0, err
}

// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
//
// Solidity: function phaseAggregators(uint16 ) view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) PhaseAggregators(arg0 uint16) (common.Address, error) {
	return _ChainlinkAggregator.Contract.PhaseAggregators(&_ChainlinkAggregator.CallOpts, arg0)
}

// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
//
// Solidity: function phaseAggregators(uint16 ) view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) PhaseAggregators(arg0 uint16) (common.Address, error) {
	return _ChainlinkAggregator.Contract.PhaseAggregators(&_ChainlinkAggregator.CallOpts, arg0)
}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) PhaseId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "phaseId")
	return *ret0, err
}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) PhaseId() (uint16, error) {
	return _ChainlinkAggregator.Contract.PhaseId(&_ChainlinkAggregator.CallOpts)
}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) PhaseId() (uint16, error) {
	return _ChainlinkAggregator.Contract.PhaseId(&_ChainlinkAggregator.CallOpts)
}

// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
//
// Solidity: function proposedAggregator() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) ProposedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "proposedAggregator")
	return *ret0, err
}

// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
//
// Solidity: function proposedAggregator() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) ProposedAggregator() (common.Address, error) {
	return _ChainlinkAggregator.Contract.ProposedAggregator(&_ChainlinkAggregator.CallOpts)
}

// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
//
// Solidity: function proposedAggregator() view returns(address)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) ProposedAggregator() (common.Address, error) {
	return _ChainlinkAggregator.Contract.ProposedAggregator(&_ChainlinkAggregator.CallOpts)
}

// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
//
// Solidity: function proposedGetRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) ProposedGetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	ret := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	out := ret
	err := _ChainlinkAggregator.contract.Call(opts, out, "proposedGetRoundData", _roundId)
	return *ret, err
}

// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
//
// Solidity: function proposedGetRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) ProposedGetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.ProposedGetRoundData(&_ChainlinkAggregator.CallOpts, _roundId)
}

// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
//
// Solidity: function proposedGetRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) ProposedGetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.ProposedGetRoundData(&_ChainlinkAggregator.CallOpts, _roundId)
}

// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
//
// Solidity: function proposedLatestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) ProposedLatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	ret := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	out := ret
	err := _ChainlinkAggregator.contract.Call(opts, out, "proposedLatestRoundData")
	return *ret, err
}

// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
//
// Solidity: function proposedLatestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) ProposedLatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.ProposedLatestRoundData(&_ChainlinkAggregator.CallOpts)
}

// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
//
// Solidity: function proposedLatestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) ProposedLatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkAggregator.Contract.ProposedLatestRoundData(&_ChainlinkAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkAggregator.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorSession) Version() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.Version(&_ChainlinkAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkAggregator *ChainlinkAggregatorCallerSession) Version() (*big.Int, error) {
	return _ChainlinkAggregator.Contract.Version(&_ChainlinkAggregator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkAggregator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkAggregator *ChainlinkAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.AcceptOwnership(&_ChainlinkAggregator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.AcceptOwnership(&_ChainlinkAggregator.TransactOpts)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactor) ConfirmAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.contract.Transact(opts, "confirmAggregator", _aggregator)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorSession) ConfirmAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.ConfirmAggregator(&_ChainlinkAggregator.TransactOpts, _aggregator)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactorSession) ConfirmAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.ConfirmAggregator(&_ChainlinkAggregator.TransactOpts, _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactor) ProposeAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.contract.Transact(opts, "proposeAggregator", _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorSession) ProposeAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.ProposeAggregator(&_ChainlinkAggregator.TransactOpts, _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactorSession) ProposeAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.ProposeAggregator(&_ChainlinkAggregator.TransactOpts, _aggregator)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactor) SetController(opts *bind.TransactOpts, _accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.contract.Transact(opts, "setController", _accessController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorSession) SetController(_accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.SetController(&_ChainlinkAggregator.TransactOpts, _accessController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactorSession) SetController(_accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.SetController(&_ChainlinkAggregator.TransactOpts, _accessController)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.TransferOwnership(&_ChainlinkAggregator.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkAggregator *ChainlinkAggregatorTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkAggregator.Contract.TransferOwnership(&_ChainlinkAggregator.TransactOpts, _to)
}

// ChainlinkAggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorAnswerUpdatedIterator struct {
	Event *ChainlinkAggregatorAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkAggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkAggregatorAnswerUpdated)
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
		it.Event = new(ChainlinkAggregatorAnswerUpdated)
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
func (it *ChainlinkAggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkAggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkAggregatorAnswerUpdated represents a AnswerUpdated event raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ChainlinkAggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregatorAnswerUpdatedIterator{contract: _ChainlinkAggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkAggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkAggregatorAnswerUpdated)
				if err := _ChainlinkAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) ParseAnswerUpdated(log types.Log) (*ChainlinkAggregatorAnswerUpdated, error) {
	event := new(ChainlinkAggregatorAnswerUpdated)
	if err := _ChainlinkAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkAggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorNewRoundIterator struct {
	Event *ChainlinkAggregatorNewRound // Event containing the contract specifics and raw log

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
func (it *ChainlinkAggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkAggregatorNewRound)
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
		it.Event = new(ChainlinkAggregatorNewRound)
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
func (it *ChainlinkAggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkAggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkAggregatorNewRound represents a NewRound event raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ChainlinkAggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregatorNewRoundIterator{contract: _ChainlinkAggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ChainlinkAggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkAggregatorNewRound)
				if err := _ChainlinkAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) ParseNewRound(log types.Log) (*ChainlinkAggregatorNewRound, error) {
	event := new(ChainlinkAggregatorNewRound)
	if err := _ChainlinkAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkAggregatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorOwnershipTransferRequestedIterator struct {
	Event *ChainlinkAggregatorOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *ChainlinkAggregatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkAggregatorOwnershipTransferRequested)
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
		it.Event = new(ChainlinkAggregatorOwnershipTransferRequested)
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
func (it *ChainlinkAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkAggregatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChainlinkAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregatorOwnershipTransferRequestedIterator{contract: _ChainlinkAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ChainlinkAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkAggregatorOwnershipTransferRequested)
				if err := _ChainlinkAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ChainlinkAggregatorOwnershipTransferRequested, error) {
	event := new(ChainlinkAggregatorOwnershipTransferRequested)
	if err := _ChainlinkAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorOwnershipTransferredIterator struct {
	Event *ChainlinkAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChainlinkAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkAggregatorOwnershipTransferred)
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
		it.Event = new(ChainlinkAggregatorOwnershipTransferred)
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
func (it *ChainlinkAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the ChainlinkAggregator contract.
type ChainlinkAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChainlinkAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkAggregatorOwnershipTransferredIterator{contract: _ChainlinkAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainlinkAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainlinkAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkAggregatorOwnershipTransferred)
				if err := _ChainlinkAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ChainlinkAggregator *ChainlinkAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*ChainlinkAggregatorOwnershipTransferred, error) {
	event := new(ChainlinkAggregatorOwnershipTransferred)
	if err := _ChainlinkAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
