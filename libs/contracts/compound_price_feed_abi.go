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

// UniswapConfigTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type UniswapConfigTokenConfig struct {
	CToken            common.Address
	Underlying        common.Address
	SymbolHash        [32]byte
	BaseUnit          *big.Int
	PriceSource       uint8
	FixedPrice        *big.Int
	UniswapMarket     common.Address
	IsUniswapReversed bool
}

// CompoundPriceFeedABI is the input ABI used to generate the binding from.
const CompoundPriceFeedABI = "[{\"inputs\":[{\"internalType\":\"contractOpenOraclePriceData\",\"name\":\"priceData_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reporter_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"anchorToleranceMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"anchorPeriod_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"baseUnit\",\"type\":\"uint256\"},{\"internalType\":\"enumUniswapConfig.PriceSource\",\"name\":\"priceSource\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fixedPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uniswapMarket\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUniswapReversed\",\"type\":\"bool\"}],\"internalType\":\"structUniswapConfig.TokenConfig[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"anchorPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"AnchorPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reporter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"anchor\",\"type\":\"uint256\"}],\"name\":\"PriceGuarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"ReporterInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"UniswapWindowUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"anchorPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBaseUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"baseUnit\",\"type\":\"uint256\"},{\"internalType\":\"enumUniswapConfig.PriceSource\",\"name\":\"priceSource\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fixedPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uniswapMarket\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUniswapReversed\",\"type\":\"bool\"}],\"internalType\":\"structUniswapConfig.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"getTokenConfigByCToken\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"baseUnit\",\"type\":\"uint256\"},{\"internalType\":\"enumUniswapConfig.PriceSource\",\"name\":\"priceSource\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fixedPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uniswapMarket\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUniswapReversed\",\"type\":\"bool\"}],\"internalType\":\"structUniswapConfig.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"getTokenConfigBySymbol\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"baseUnit\",\"type\":\"uint256\"},{\"internalType\":\"enumUniswapConfig.PriceSource\",\"name\":\"priceSource\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fixedPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uniswapMarket\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUniswapReversed\",\"type\":\"bool\"}],\"internalType\":\"structUniswapConfig.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"}],\"name\":\"getTokenConfigBySymbolHash\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"baseUnit\",\"type\":\"uint256\"},{\"internalType\":\"enumUniswapConfig.PriceSource\",\"name\":\"priceSource\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fixedPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uniswapMarket\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUniswapReversed\",\"type\":\"bool\"}],\"internalType\":\"structUniswapConfig.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"getTokenConfigByUnderlying\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"baseUnit\",\"type\":\"uint256\"},{\"internalType\":\"enumUniswapConfig.PriceSource\",\"name\":\"priceSource\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fixedPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uniswapMarket\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUniswapReversed\",\"type\":\"bool\"}],\"internalType\":\"structUniswapConfig.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"invalidateReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerBoundAnchorRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"newObservations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"oldObservations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"messages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"string[]\",\"name\":\"symbols\",\"type\":\"string[]\"}],\"name\":\"postPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceData\",\"outputs\":[{\"internalType\":\"contractOpenOraclePriceData\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reporter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reporterInvalidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"source\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperBoundAnchorRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CompoundPriceFeed is an auto generated Go binding around an Ethereum contract.
type CompoundPriceFeed struct {
	CompoundPriceFeedCaller     // Read-only binding to the contract
	CompoundPriceFeedTransactor // Write-only binding to the contract
	CompoundPriceFeedFilterer   // Log filterer for contract events
}

// CompoundPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundPriceFeedSession struct {
	Contract     *CompoundPriceFeed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CompoundPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundPriceFeedCallerSession struct {
	Contract *CompoundPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CompoundPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundPriceFeedTransactorSession struct {
	Contract     *CompoundPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CompoundPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundPriceFeedRaw struct {
	Contract *CompoundPriceFeed // Generic contract binding to access the raw methods on
}

// CompoundPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundPriceFeedCallerRaw struct {
	Contract *CompoundPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundPriceFeedTransactorRaw struct {
	Contract *CompoundPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundPriceFeed creates a new instance of CompoundPriceFeed, bound to a specific deployed contract.
func NewCompoundPriceFeed(address common.Address, backend bind.ContractBackend) (*CompoundPriceFeed, error) {
	contract, err := bindCompoundPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeed{CompoundPriceFeedCaller: CompoundPriceFeedCaller{contract: contract}, CompoundPriceFeedTransactor: CompoundPriceFeedTransactor{contract: contract}, CompoundPriceFeedFilterer: CompoundPriceFeedFilterer{contract: contract}}, nil
}

// NewCompoundPriceFeedCaller creates a new read-only instance of CompoundPriceFeed, bound to a specific deployed contract.
func NewCompoundPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*CompoundPriceFeedCaller, error) {
	contract, err := bindCompoundPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedCaller{contract: contract}, nil
}

// NewCompoundPriceFeedTransactor creates a new write-only instance of CompoundPriceFeed, bound to a specific deployed contract.
func NewCompoundPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundPriceFeedTransactor, error) {
	contract, err := bindCompoundPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedTransactor{contract: contract}, nil
}

// NewCompoundPriceFeedFilterer creates a new log filterer instance of CompoundPriceFeed, bound to a specific deployed contract.
func NewCompoundPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundPriceFeedFilterer, error) {
	contract, err := bindCompoundPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedFilterer{contract: contract}, nil
}

// bindCompoundPriceFeed binds a generic wrapper to an already deployed contract.
func bindCompoundPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundPriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundPriceFeed *CompoundPriceFeedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundPriceFeed.Contract.CompoundPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundPriceFeed *CompoundPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.CompoundPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundPriceFeed *CompoundPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.CompoundPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundPriceFeed *CompoundPriceFeedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundPriceFeed *CompoundPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundPriceFeed *CompoundPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// AnchorPeriod is a free data retrieval call binding the contract method 0xe9206d78.
//
// Solidity: function anchorPeriod() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) AnchorPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "anchorPeriod")
	return *ret0, err
}

// AnchorPeriod is a free data retrieval call binding the contract method 0xe9206d78.
//
// Solidity: function anchorPeriod() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) AnchorPeriod() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.AnchorPeriod(&_CompoundPriceFeed.CallOpts)
}

// AnchorPeriod is a free data retrieval call binding the contract method 0xe9206d78.
//
// Solidity: function anchorPeriod() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) AnchorPeriod() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.AnchorPeriod(&_CompoundPriceFeed.CallOpts)
}

// EthBaseUnit is a free data retrieval call binding the contract method 0xd1b353b4.
//
// Solidity: function ethBaseUnit() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) EthBaseUnit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "ethBaseUnit")
	return *ret0, err
}

// EthBaseUnit is a free data retrieval call binding the contract method 0xd1b353b4.
//
// Solidity: function ethBaseUnit() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) EthBaseUnit() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.EthBaseUnit(&_CompoundPriceFeed.CallOpts)
}

// EthBaseUnit is a free data retrieval call binding the contract method 0xd1b353b4.
//
// Solidity: function ethBaseUnit() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) EthBaseUnit() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.EthBaseUnit(&_CompoundPriceFeed.CallOpts)
}

// ExpScale is a free data retrieval call binding the contract method 0x69aa3ac6.
//
// Solidity: function expScale() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) ExpScale(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "expScale")
	return *ret0, err
}

// ExpScale is a free data retrieval call binding the contract method 0x69aa3ac6.
//
// Solidity: function expScale() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) ExpScale() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.ExpScale(&_CompoundPriceFeed.CallOpts)
}

// ExpScale is a free data retrieval call binding the contract method 0x69aa3ac6.
//
// Solidity: function expScale() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) ExpScale() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.ExpScale(&_CompoundPriceFeed.CallOpts)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x8a003888.
//
// Solidity: function getTokenConfig(uint256 i) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCaller) GetTokenConfig(opts *bind.CallOpts, i *big.Int) (UniswapConfigTokenConfig, error) {
	var (
		ret0 = new(UniswapConfigTokenConfig)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "getTokenConfig", i)
	return *ret0, err
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x8a003888.
//
// Solidity: function getTokenConfig(uint256 i) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedSession) GetTokenConfig(i *big.Int) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfig(&_CompoundPriceFeed.CallOpts, i)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x8a003888.
//
// Solidity: function getTokenConfig(uint256 i) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) GetTokenConfig(i *big.Int) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfig(&_CompoundPriceFeed.CallOpts, i)
}

// GetTokenConfigByCToken is a free data retrieval call binding the contract method 0x9f599631.
//
// Solidity: function getTokenConfigByCToken(address cToken) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCaller) GetTokenConfigByCToken(opts *bind.CallOpts, cToken common.Address) (UniswapConfigTokenConfig, error) {
	var (
		ret0 = new(UniswapConfigTokenConfig)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "getTokenConfigByCToken", cToken)
	return *ret0, err
}

// GetTokenConfigByCToken is a free data retrieval call binding the contract method 0x9f599631.
//
// Solidity: function getTokenConfigByCToken(address cToken) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedSession) GetTokenConfigByCToken(cToken common.Address) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigByCToken(&_CompoundPriceFeed.CallOpts, cToken)
}

// GetTokenConfigByCToken is a free data retrieval call binding the contract method 0x9f599631.
//
// Solidity: function getTokenConfigByCToken(address cToken) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) GetTokenConfigByCToken(cToken common.Address) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigByCToken(&_CompoundPriceFeed.CallOpts, cToken)
}

// GetTokenConfigBySymbol is a free data retrieval call binding the contract method 0x276c2cba.
//
// Solidity: function getTokenConfigBySymbol(string symbol) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCaller) GetTokenConfigBySymbol(opts *bind.CallOpts, symbol string) (UniswapConfigTokenConfig, error) {
	var (
		ret0 = new(UniswapConfigTokenConfig)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "getTokenConfigBySymbol", symbol)
	return *ret0, err
}

// GetTokenConfigBySymbol is a free data retrieval call binding the contract method 0x276c2cba.
//
// Solidity: function getTokenConfigBySymbol(string symbol) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedSession) GetTokenConfigBySymbol(symbol string) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigBySymbol(&_CompoundPriceFeed.CallOpts, symbol)
}

// GetTokenConfigBySymbol is a free data retrieval call binding the contract method 0x276c2cba.
//
// Solidity: function getTokenConfigBySymbol(string symbol) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) GetTokenConfigBySymbol(symbol string) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigBySymbol(&_CompoundPriceFeed.CallOpts, symbol)
}

// GetTokenConfigBySymbolHash is a free data retrieval call binding the contract method 0x1a125204.
//
// Solidity: function getTokenConfigBySymbolHash(bytes32 symbolHash) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCaller) GetTokenConfigBySymbolHash(opts *bind.CallOpts, symbolHash [32]byte) (UniswapConfigTokenConfig, error) {
	var (
		ret0 = new(UniswapConfigTokenConfig)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "getTokenConfigBySymbolHash", symbolHash)
	return *ret0, err
}

// GetTokenConfigBySymbolHash is a free data retrieval call binding the contract method 0x1a125204.
//
// Solidity: function getTokenConfigBySymbolHash(bytes32 symbolHash) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedSession) GetTokenConfigBySymbolHash(symbolHash [32]byte) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigBySymbolHash(&_CompoundPriceFeed.CallOpts, symbolHash)
}

// GetTokenConfigBySymbolHash is a free data retrieval call binding the contract method 0x1a125204.
//
// Solidity: function getTokenConfigBySymbolHash(bytes32 symbolHash) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) GetTokenConfigBySymbolHash(symbolHash [32]byte) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigBySymbolHash(&_CompoundPriceFeed.CallOpts, symbolHash)
}

// GetTokenConfigByUnderlying is a free data retrieval call binding the contract method 0x4da21942.
//
// Solidity: function getTokenConfigByUnderlying(address underlying) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCaller) GetTokenConfigByUnderlying(opts *bind.CallOpts, underlying common.Address) (UniswapConfigTokenConfig, error) {
	var (
		ret0 = new(UniswapConfigTokenConfig)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "getTokenConfigByUnderlying", underlying)
	return *ret0, err
}

// GetTokenConfigByUnderlying is a free data retrieval call binding the contract method 0x4da21942.
//
// Solidity: function getTokenConfigByUnderlying(address underlying) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedSession) GetTokenConfigByUnderlying(underlying common.Address) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigByUnderlying(&_CompoundPriceFeed.CallOpts, underlying)
}

// GetTokenConfigByUnderlying is a free data retrieval call binding the contract method 0x4da21942.
//
// Solidity: function getTokenConfigByUnderlying(address underlying) view returns((address,address,bytes32,uint256,uint8,uint256,address,bool))
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) GetTokenConfigByUnderlying(underlying common.Address) (UniswapConfigTokenConfig, error) {
	return _CompoundPriceFeed.Contract.GetTokenConfigByUnderlying(&_CompoundPriceFeed.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) GetUnderlyingPrice(opts *bind.CallOpts, cToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "getUnderlyingPrice", cToken)
	return *ret0, err
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _CompoundPriceFeed.Contract.GetUnderlyingPrice(&_CompoundPriceFeed.CallOpts, cToken)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _CompoundPriceFeed.Contract.GetUnderlyingPrice(&_CompoundPriceFeed.CallOpts, cToken)
}

// LowerBoundAnchorRatio is a free data retrieval call binding the contract method 0x92b84357.
//
// Solidity: function lowerBoundAnchorRatio() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) LowerBoundAnchorRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "lowerBoundAnchorRatio")
	return *ret0, err
}

// LowerBoundAnchorRatio is a free data retrieval call binding the contract method 0x92b84357.
//
// Solidity: function lowerBoundAnchorRatio() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) LowerBoundAnchorRatio() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.LowerBoundAnchorRatio(&_CompoundPriceFeed.CallOpts)
}

// LowerBoundAnchorRatio is a free data retrieval call binding the contract method 0x92b84357.
//
// Solidity: function lowerBoundAnchorRatio() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) LowerBoundAnchorRatio() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.LowerBoundAnchorRatio(&_CompoundPriceFeed.CallOpts)
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) MaxTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "maxTokens")
	return *ret0, err
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) MaxTokens() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.MaxTokens(&_CompoundPriceFeed.CallOpts)
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) MaxTokens() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.MaxTokens(&_CompoundPriceFeed.CallOpts)
}

// NewObservations is a free data retrieval call binding the contract method 0xeaa1c2ca.
//
// Solidity: function newObservations(bytes32 ) view returns(uint256 timestamp, uint256 acc)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) NewObservations(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timestamp *big.Int
	Acc       *big.Int
}, error) {
	ret := new(struct {
		Timestamp *big.Int
		Acc       *big.Int
	})
	out := ret
	err := _CompoundPriceFeed.contract.Call(opts, out, "newObservations", arg0)
	return *ret, err
}

// NewObservations is a free data retrieval call binding the contract method 0xeaa1c2ca.
//
// Solidity: function newObservations(bytes32 ) view returns(uint256 timestamp, uint256 acc)
func (_CompoundPriceFeed *CompoundPriceFeedSession) NewObservations(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Acc       *big.Int
}, error) {
	return _CompoundPriceFeed.Contract.NewObservations(&_CompoundPriceFeed.CallOpts, arg0)
}

// NewObservations is a free data retrieval call binding the contract method 0xeaa1c2ca.
//
// Solidity: function newObservations(bytes32 ) view returns(uint256 timestamp, uint256 acc)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) NewObservations(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Acc       *big.Int
}, error) {
	return _CompoundPriceFeed.Contract.NewObservations(&_CompoundPriceFeed.CallOpts, arg0)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "numTokens")
	return *ret0, err
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) NumTokens() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.NumTokens(&_CompoundPriceFeed.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) NumTokens() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.NumTokens(&_CompoundPriceFeed.CallOpts)
}

// OldObservations is a free data retrieval call binding the contract method 0x37c0e12d.
//
// Solidity: function oldObservations(bytes32 ) view returns(uint256 timestamp, uint256 acc)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) OldObservations(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timestamp *big.Int
	Acc       *big.Int
}, error) {
	ret := new(struct {
		Timestamp *big.Int
		Acc       *big.Int
	})
	out := ret
	err := _CompoundPriceFeed.contract.Call(opts, out, "oldObservations", arg0)
	return *ret, err
}

// OldObservations is a free data retrieval call binding the contract method 0x37c0e12d.
//
// Solidity: function oldObservations(bytes32 ) view returns(uint256 timestamp, uint256 acc)
func (_CompoundPriceFeed *CompoundPriceFeedSession) OldObservations(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Acc       *big.Int
}, error) {
	return _CompoundPriceFeed.Contract.OldObservations(&_CompoundPriceFeed.CallOpts, arg0)
}

// OldObservations is a free data retrieval call binding the contract method 0x37c0e12d.
//
// Solidity: function oldObservations(bytes32 ) view returns(uint256 timestamp, uint256 acc)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) OldObservations(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Acc       *big.Int
}, error) {
	return _CompoundPriceFeed.Contract.OldObservations(&_CompoundPriceFeed.CallOpts, arg0)
}

// Price is a free data retrieval call binding the contract method 0xfe2c6198.
//
// Solidity: function price(string symbol) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) Price(opts *bind.CallOpts, symbol string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "price", symbol)
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xfe2c6198.
//
// Solidity: function price(string symbol) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) Price(symbol string) (*big.Int, error) {
	return _CompoundPriceFeed.Contract.Price(&_CompoundPriceFeed.CallOpts, symbol)
}

// Price is a free data retrieval call binding the contract method 0xfe2c6198.
//
// Solidity: function price(string symbol) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) Price(symbol string) (*big.Int, error) {
	return _CompoundPriceFeed.Contract.Price(&_CompoundPriceFeed.CallOpts, symbol)
}

// PriceData is a free data retrieval call binding the contract method 0xe61a5fe4.
//
// Solidity: function priceData() view returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) PriceData(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "priceData")
	return *ret0, err
}

// PriceData is a free data retrieval call binding the contract method 0xe61a5fe4.
//
// Solidity: function priceData() view returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedSession) PriceData() (common.Address, error) {
	return _CompoundPriceFeed.Contract.PriceData(&_CompoundPriceFeed.CallOpts)
}

// PriceData is a free data retrieval call binding the contract method 0xe61a5fe4.
//
// Solidity: function priceData() view returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) PriceData() (common.Address, error) {
	return _CompoundPriceFeed.Contract.PriceData(&_CompoundPriceFeed.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) Prices(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "prices", arg0)
	return *ret0, err
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _CompoundPriceFeed.Contract.Prices(&_CompoundPriceFeed.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _CompoundPriceFeed.Contract.Prices(&_CompoundPriceFeed.CallOpts, arg0)
}

// Reporter is a free data retrieval call binding the contract method 0x010ec441.
//
// Solidity: function reporter() view returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) Reporter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "reporter")
	return *ret0, err
}

// Reporter is a free data retrieval call binding the contract method 0x010ec441.
//
// Solidity: function reporter() view returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedSession) Reporter() (common.Address, error) {
	return _CompoundPriceFeed.Contract.Reporter(&_CompoundPriceFeed.CallOpts)
}

// Reporter is a free data retrieval call binding the contract method 0x010ec441.
//
// Solidity: function reporter() view returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) Reporter() (common.Address, error) {
	return _CompoundPriceFeed.Contract.Reporter(&_CompoundPriceFeed.CallOpts)
}

// ReporterInvalidated is a free data retrieval call binding the contract method 0x651ed788.
//
// Solidity: function reporterInvalidated() view returns(bool)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) ReporterInvalidated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "reporterInvalidated")
	return *ret0, err
}

// ReporterInvalidated is a free data retrieval call binding the contract method 0x651ed788.
//
// Solidity: function reporterInvalidated() view returns(bool)
func (_CompoundPriceFeed *CompoundPriceFeedSession) ReporterInvalidated() (bool, error) {
	return _CompoundPriceFeed.Contract.ReporterInvalidated(&_CompoundPriceFeed.CallOpts)
}

// ReporterInvalidated is a free data retrieval call binding the contract method 0x651ed788.
//
// Solidity: function reporterInvalidated() view returns(bool)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) ReporterInvalidated() (bool, error) {
	return _CompoundPriceFeed.Contract.ReporterInvalidated(&_CompoundPriceFeed.CallOpts)
}

// Source is a free data retrieval call binding the contract method 0x482a6193.
//
// Solidity: function source(bytes message, bytes signature) pure returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) Source(opts *bind.CallOpts, message []byte, signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "source", message, signature)
	return *ret0, err
}

// Source is a free data retrieval call binding the contract method 0x482a6193.
//
// Solidity: function source(bytes message, bytes signature) pure returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedSession) Source(message []byte, signature []byte) (common.Address, error) {
	return _CompoundPriceFeed.Contract.Source(&_CompoundPriceFeed.CallOpts, message, signature)
}

// Source is a free data retrieval call binding the contract method 0x482a6193.
//
// Solidity: function source(bytes message, bytes signature) pure returns(address)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) Source(message []byte, signature []byte) (common.Address, error) {
	return _CompoundPriceFeed.Contract.Source(&_CompoundPriceFeed.CallOpts, message, signature)
}

// UpperBoundAnchorRatio is a free data retrieval call binding the contract method 0x24105209.
//
// Solidity: function upperBoundAnchorRatio() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCaller) UpperBoundAnchorRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CompoundPriceFeed.contract.Call(opts, out, "upperBoundAnchorRatio")
	return *ret0, err
}

// UpperBoundAnchorRatio is a free data retrieval call binding the contract method 0x24105209.
//
// Solidity: function upperBoundAnchorRatio() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedSession) UpperBoundAnchorRatio() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.UpperBoundAnchorRatio(&_CompoundPriceFeed.CallOpts)
}

// UpperBoundAnchorRatio is a free data retrieval call binding the contract method 0x24105209.
//
// Solidity: function upperBoundAnchorRatio() view returns(uint256)
func (_CompoundPriceFeed *CompoundPriceFeedCallerSession) UpperBoundAnchorRatio() (*big.Int, error) {
	return _CompoundPriceFeed.Contract.UpperBoundAnchorRatio(&_CompoundPriceFeed.CallOpts)
}

// InvalidateReporter is a paid mutator transaction binding the contract method 0x8aba91b4.
//
// Solidity: function invalidateReporter(bytes message, bytes signature) returns()
func (_CompoundPriceFeed *CompoundPriceFeedTransactor) InvalidateReporter(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _CompoundPriceFeed.contract.Transact(opts, "invalidateReporter", message, signature)
}

// InvalidateReporter is a paid mutator transaction binding the contract method 0x8aba91b4.
//
// Solidity: function invalidateReporter(bytes message, bytes signature) returns()
func (_CompoundPriceFeed *CompoundPriceFeedSession) InvalidateReporter(message []byte, signature []byte) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.InvalidateReporter(&_CompoundPriceFeed.TransactOpts, message, signature)
}

// InvalidateReporter is a paid mutator transaction binding the contract method 0x8aba91b4.
//
// Solidity: function invalidateReporter(bytes message, bytes signature) returns()
func (_CompoundPriceFeed *CompoundPriceFeedTransactorSession) InvalidateReporter(message []byte, signature []byte) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.InvalidateReporter(&_CompoundPriceFeed.TransactOpts, message, signature)
}

// PostPrices is a paid mutator transaction binding the contract method 0xecc1e984.
//
// Solidity: function postPrices(bytes[] messages, bytes[] signatures, string[] symbols) returns()
func (_CompoundPriceFeed *CompoundPriceFeedTransactor) PostPrices(opts *bind.TransactOpts, messages [][]byte, signatures [][]byte, symbols []string) (*types.Transaction, error) {
	return _CompoundPriceFeed.contract.Transact(opts, "postPrices", messages, signatures, symbols)
}

// PostPrices is a paid mutator transaction binding the contract method 0xecc1e984.
//
// Solidity: function postPrices(bytes[] messages, bytes[] signatures, string[] symbols) returns()
func (_CompoundPriceFeed *CompoundPriceFeedSession) PostPrices(messages [][]byte, signatures [][]byte, symbols []string) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.PostPrices(&_CompoundPriceFeed.TransactOpts, messages, signatures, symbols)
}

// PostPrices is a paid mutator transaction binding the contract method 0xecc1e984.
//
// Solidity: function postPrices(bytes[] messages, bytes[] signatures, string[] symbols) returns()
func (_CompoundPriceFeed *CompoundPriceFeedTransactorSession) PostPrices(messages [][]byte, signatures [][]byte, symbols []string) (*types.Transaction, error) {
	return _CompoundPriceFeed.Contract.PostPrices(&_CompoundPriceFeed.TransactOpts, messages, signatures, symbols)
}

// CompoundPriceFeedAnchorPriceUpdatedIterator is returned from FilterAnchorPriceUpdated and is used to iterate over the raw logs and unpacked data for AnchorPriceUpdated events raised by the CompoundPriceFeed contract.
type CompoundPriceFeedAnchorPriceUpdatedIterator struct {
	Event *CompoundPriceFeedAnchorPriceUpdated // Event containing the contract specifics and raw log

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
func (it *CompoundPriceFeedAnchorPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundPriceFeedAnchorPriceUpdated)
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
		it.Event = new(CompoundPriceFeedAnchorPriceUpdated)
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
func (it *CompoundPriceFeedAnchorPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundPriceFeedAnchorPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundPriceFeedAnchorPriceUpdated represents a AnchorPriceUpdated event raised by the CompoundPriceFeed contract.
type CompoundPriceFeedAnchorPriceUpdated struct {
	Symbol       string
	AnchorPrice  *big.Int
	OldTimestamp *big.Int
	NewTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAnchorPriceUpdated is a free log retrieval operation binding the contract event 0xf63d078e0de851897107641e96093a59e5ddc3c25e7b85a2585e3eba9e774a7b.
//
// Solidity: event AnchorPriceUpdated(string symbol, uint256 anchorPrice, uint256 oldTimestamp, uint256 newTimestamp)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) FilterAnchorPriceUpdated(opts *bind.FilterOpts) (*CompoundPriceFeedAnchorPriceUpdatedIterator, error) {

	logs, sub, err := _CompoundPriceFeed.contract.FilterLogs(opts, "AnchorPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedAnchorPriceUpdatedIterator{contract: _CompoundPriceFeed.contract, event: "AnchorPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchAnchorPriceUpdated is a free log subscription operation binding the contract event 0xf63d078e0de851897107641e96093a59e5ddc3c25e7b85a2585e3eba9e774a7b.
//
// Solidity: event AnchorPriceUpdated(string symbol, uint256 anchorPrice, uint256 oldTimestamp, uint256 newTimestamp)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) WatchAnchorPriceUpdated(opts *bind.WatchOpts, sink chan<- *CompoundPriceFeedAnchorPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _CompoundPriceFeed.contract.WatchLogs(opts, "AnchorPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundPriceFeedAnchorPriceUpdated)
				if err := _CompoundPriceFeed.contract.UnpackLog(event, "AnchorPriceUpdated", log); err != nil {
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

// ParseAnchorPriceUpdated is a log parse operation binding the contract event 0xf63d078e0de851897107641e96093a59e5ddc3c25e7b85a2585e3eba9e774a7b.
//
// Solidity: event AnchorPriceUpdated(string symbol, uint256 anchorPrice, uint256 oldTimestamp, uint256 newTimestamp)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) ParseAnchorPriceUpdated(log types.Log) (*CompoundPriceFeedAnchorPriceUpdated, error) {
	event := new(CompoundPriceFeedAnchorPriceUpdated)
	if err := _CompoundPriceFeed.contract.UnpackLog(event, "AnchorPriceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CompoundPriceFeedPriceGuardedIterator is returned from FilterPriceGuarded and is used to iterate over the raw logs and unpacked data for PriceGuarded events raised by the CompoundPriceFeed contract.
type CompoundPriceFeedPriceGuardedIterator struct {
	Event *CompoundPriceFeedPriceGuarded // Event containing the contract specifics and raw log

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
func (it *CompoundPriceFeedPriceGuardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundPriceFeedPriceGuarded)
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
		it.Event = new(CompoundPriceFeedPriceGuarded)
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
func (it *CompoundPriceFeedPriceGuardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundPriceFeedPriceGuardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundPriceFeedPriceGuarded represents a PriceGuarded event raised by the CompoundPriceFeed contract.
type CompoundPriceFeedPriceGuarded struct {
	Symbol   string
	Reporter *big.Int
	Anchor   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceGuarded is a free log retrieval operation binding the contract event 0x90756d4c8646a4591078abac0e4e32dfa8437921729e36d51b88b659d265bfde.
//
// Solidity: event PriceGuarded(string symbol, uint256 reporter, uint256 anchor)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) FilterPriceGuarded(opts *bind.FilterOpts) (*CompoundPriceFeedPriceGuardedIterator, error) {

	logs, sub, err := _CompoundPriceFeed.contract.FilterLogs(opts, "PriceGuarded")
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedPriceGuardedIterator{contract: _CompoundPriceFeed.contract, event: "PriceGuarded", logs: logs, sub: sub}, nil
}

// WatchPriceGuarded is a free log subscription operation binding the contract event 0x90756d4c8646a4591078abac0e4e32dfa8437921729e36d51b88b659d265bfde.
//
// Solidity: event PriceGuarded(string symbol, uint256 reporter, uint256 anchor)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) WatchPriceGuarded(opts *bind.WatchOpts, sink chan<- *CompoundPriceFeedPriceGuarded) (event.Subscription, error) {

	logs, sub, err := _CompoundPriceFeed.contract.WatchLogs(opts, "PriceGuarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundPriceFeedPriceGuarded)
				if err := _CompoundPriceFeed.contract.UnpackLog(event, "PriceGuarded", log); err != nil {
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

// ParsePriceGuarded is a log parse operation binding the contract event 0x90756d4c8646a4591078abac0e4e32dfa8437921729e36d51b88b659d265bfde.
//
// Solidity: event PriceGuarded(string symbol, uint256 reporter, uint256 anchor)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) ParsePriceGuarded(log types.Log) (*CompoundPriceFeedPriceGuarded, error) {
	event := new(CompoundPriceFeedPriceGuarded)
	if err := _CompoundPriceFeed.contract.UnpackLog(event, "PriceGuarded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CompoundPriceFeedPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the CompoundPriceFeed contract.
type CompoundPriceFeedPriceUpdatedIterator struct {
	Event *CompoundPriceFeedPriceUpdated // Event containing the contract specifics and raw log

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
func (it *CompoundPriceFeedPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundPriceFeedPriceUpdated)
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
		it.Event = new(CompoundPriceFeedPriceUpdated)
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
func (it *CompoundPriceFeedPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundPriceFeedPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundPriceFeedPriceUpdated represents a PriceUpdated event raised by the CompoundPriceFeed contract.
type CompoundPriceFeedPriceUpdated struct {
	Symbol string
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string symbol, uint256 price)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*CompoundPriceFeedPriceUpdatedIterator, error) {

	logs, sub, err := _CompoundPriceFeed.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedPriceUpdatedIterator{contract: _CompoundPriceFeed.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string symbol, uint256 price)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *CompoundPriceFeedPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _CompoundPriceFeed.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundPriceFeedPriceUpdated)
				if err := _CompoundPriceFeed.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string symbol, uint256 price)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) ParsePriceUpdated(log types.Log) (*CompoundPriceFeedPriceUpdated, error) {
	event := new(CompoundPriceFeedPriceUpdated)
	if err := _CompoundPriceFeed.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CompoundPriceFeedReporterInvalidatedIterator is returned from FilterReporterInvalidated and is used to iterate over the raw logs and unpacked data for ReporterInvalidated events raised by the CompoundPriceFeed contract.
type CompoundPriceFeedReporterInvalidatedIterator struct {
	Event *CompoundPriceFeedReporterInvalidated // Event containing the contract specifics and raw log

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
func (it *CompoundPriceFeedReporterInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundPriceFeedReporterInvalidated)
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
		it.Event = new(CompoundPriceFeedReporterInvalidated)
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
func (it *CompoundPriceFeedReporterInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundPriceFeedReporterInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundPriceFeedReporterInvalidated represents a ReporterInvalidated event raised by the CompoundPriceFeed contract.
type CompoundPriceFeedReporterInvalidated struct {
	Reporter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReporterInvalidated is a free log retrieval operation binding the contract event 0x98a13f7b181a3a1f99c871e7a3507d4a037d386d157279f978e0d555ae9fe74d.
//
// Solidity: event ReporterInvalidated(address reporter)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) FilterReporterInvalidated(opts *bind.FilterOpts) (*CompoundPriceFeedReporterInvalidatedIterator, error) {

	logs, sub, err := _CompoundPriceFeed.contract.FilterLogs(opts, "ReporterInvalidated")
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedReporterInvalidatedIterator{contract: _CompoundPriceFeed.contract, event: "ReporterInvalidated", logs: logs, sub: sub}, nil
}

// WatchReporterInvalidated is a free log subscription operation binding the contract event 0x98a13f7b181a3a1f99c871e7a3507d4a037d386d157279f978e0d555ae9fe74d.
//
// Solidity: event ReporterInvalidated(address reporter)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) WatchReporterInvalidated(opts *bind.WatchOpts, sink chan<- *CompoundPriceFeedReporterInvalidated) (event.Subscription, error) {

	logs, sub, err := _CompoundPriceFeed.contract.WatchLogs(opts, "ReporterInvalidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundPriceFeedReporterInvalidated)
				if err := _CompoundPriceFeed.contract.UnpackLog(event, "ReporterInvalidated", log); err != nil {
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

// ParseReporterInvalidated is a log parse operation binding the contract event 0x98a13f7b181a3a1f99c871e7a3507d4a037d386d157279f978e0d555ae9fe74d.
//
// Solidity: event ReporterInvalidated(address reporter)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) ParseReporterInvalidated(log types.Log) (*CompoundPriceFeedReporterInvalidated, error) {
	event := new(CompoundPriceFeedReporterInvalidated)
	if err := _CompoundPriceFeed.contract.UnpackLog(event, "ReporterInvalidated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CompoundPriceFeedUniswapWindowUpdatedIterator is returned from FilterUniswapWindowUpdated and is used to iterate over the raw logs and unpacked data for UniswapWindowUpdated events raised by the CompoundPriceFeed contract.
type CompoundPriceFeedUniswapWindowUpdatedIterator struct {
	Event *CompoundPriceFeedUniswapWindowUpdated // Event containing the contract specifics and raw log

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
func (it *CompoundPriceFeedUniswapWindowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundPriceFeedUniswapWindowUpdated)
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
		it.Event = new(CompoundPriceFeedUniswapWindowUpdated)
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
func (it *CompoundPriceFeedUniswapWindowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundPriceFeedUniswapWindowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundPriceFeedUniswapWindowUpdated represents a UniswapWindowUpdated event raised by the CompoundPriceFeed contract.
type CompoundPriceFeedUniswapWindowUpdated struct {
	SymbolHash   [32]byte
	OldTimestamp *big.Int
	NewTimestamp *big.Int
	OldPrice     *big.Int
	NewPrice     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUniswapWindowUpdated is a free log retrieval operation binding the contract event 0xe37d39315e3419c0937360f1ac88f2c52ecf67e3b22b367f82047ddb4591904a.
//
// Solidity: event UniswapWindowUpdated(bytes32 indexed symbolHash, uint256 oldTimestamp, uint256 newTimestamp, uint256 oldPrice, uint256 newPrice)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) FilterUniswapWindowUpdated(opts *bind.FilterOpts, symbolHash [][32]byte) (*CompoundPriceFeedUniswapWindowUpdatedIterator, error) {

	var symbolHashRule []interface{}
	for _, symbolHashItem := range symbolHash {
		symbolHashRule = append(symbolHashRule, symbolHashItem)
	}

	logs, sub, err := _CompoundPriceFeed.contract.FilterLogs(opts, "UniswapWindowUpdated", symbolHashRule)
	if err != nil {
		return nil, err
	}
	return &CompoundPriceFeedUniswapWindowUpdatedIterator{contract: _CompoundPriceFeed.contract, event: "UniswapWindowUpdated", logs: logs, sub: sub}, nil
}

// WatchUniswapWindowUpdated is a free log subscription operation binding the contract event 0xe37d39315e3419c0937360f1ac88f2c52ecf67e3b22b367f82047ddb4591904a.
//
// Solidity: event UniswapWindowUpdated(bytes32 indexed symbolHash, uint256 oldTimestamp, uint256 newTimestamp, uint256 oldPrice, uint256 newPrice)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) WatchUniswapWindowUpdated(opts *bind.WatchOpts, sink chan<- *CompoundPriceFeedUniswapWindowUpdated, symbolHash [][32]byte) (event.Subscription, error) {

	var symbolHashRule []interface{}
	for _, symbolHashItem := range symbolHash {
		symbolHashRule = append(symbolHashRule, symbolHashItem)
	}

	logs, sub, err := _CompoundPriceFeed.contract.WatchLogs(opts, "UniswapWindowUpdated", symbolHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundPriceFeedUniswapWindowUpdated)
				if err := _CompoundPriceFeed.contract.UnpackLog(event, "UniswapWindowUpdated", log); err != nil {
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

// ParseUniswapWindowUpdated is a log parse operation binding the contract event 0xe37d39315e3419c0937360f1ac88f2c52ecf67e3b22b367f82047ddb4591904a.
//
// Solidity: event UniswapWindowUpdated(bytes32 indexed symbolHash, uint256 oldTimestamp, uint256 newTimestamp, uint256 oldPrice, uint256 newPrice)
func (_CompoundPriceFeed *CompoundPriceFeedFilterer) ParseUniswapWindowUpdated(log types.Log) (*CompoundPriceFeedUniswapWindowUpdated, error) {
	event := new(CompoundPriceFeedUniswapWindowUpdated)
	if err := _CompoundPriceFeed.contract.UnpackLog(event, "UniswapWindowUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
