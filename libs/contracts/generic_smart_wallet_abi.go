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

// ISmartWalletImplementationGetExpectedInParams is an auto generated low-level Go binding around an user-defined struct.
type ISmartWalletImplementationGetExpectedInParams struct {
	SwapContract common.Address
	DestAmount   *big.Int
	TradePath    []common.Address
	FeeMode      uint8
	FeeBps       *big.Int
	ExtraArgs    []byte
}

// ISmartWalletImplementationGetExpectedReturnParams is an auto generated low-level Go binding around an user-defined struct.
type ISmartWalletImplementationGetExpectedReturnParams struct {
	SwapContract common.Address
	SrcAmount    *big.Int
	TradePath    []common.Address
	FeeMode      uint8
	FeeBps       *big.Int
	ExtraArgs    []byte
}

// ISmartWalletImplementationSwapAndDepositParams is an auto generated low-level Go binding around an user-defined struct.
type ISmartWalletImplementationSwapAndDepositParams struct {
	SwapContract    common.Address
	LendingContract common.Address
	SrcAmount       *big.Int
	MinDestAmount   *big.Int
	TradePath       []common.Address
	FeeMode         uint8
	FeeBps          *big.Int
	PlatformWallet  common.Address
	ExtraArgs       []byte
}

// ISmartWalletImplementationSwapAndRepayParams is an auto generated low-level Go binding around an user-defined struct.
type ISmartWalletImplementationSwapAndRepayParams struct {
	SwapContract    common.Address
	LendingContract common.Address
	SrcAmount       *big.Int
	PayAmount       *big.Int
	TradePath       []common.Address
	RateMode        *big.Int
	FeeMode         uint8
	FeeBps          *big.Int
	PlatformWallet  common.Address
	ExtraArgs       []byte
}

// ISmartWalletImplementationSwapParams is an auto generated low-level Go binding around an user-defined struct.
type ISmartWalletImplementationSwapParams struct {
	SwapContract   common.Address
	SrcAmount      *big.Int
	MinDestAmount  *big.Int
	TradePath      []common.Address
	FeeMode        uint8
	FeeBps         *big.Int
	PlatformWallet common.Address
	ExtraArgs      []byte
}

// ISmartWalletImplementationWithdrawFromLendingPlatformParams is an auto generated low-level Go binding around an user-defined struct.
type ISmartWalletImplementationWithdrawFromLendingPlatformParams struct {
	LendingContract common.Address
	Token           common.Address
	Amount          *big.Int
	MinReturn       *big.Int
}

// GenericSmartWalletABI is the input ABI used to generate the binding from.
const GenericSmartWalletABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"spenders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isReset\",\"type\":\"bool\"}],\"name\":\"ApprovedAllowances\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimedPlatformFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lendingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"SwapAndDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lendingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"SwapAndRepay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lendingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualReturnAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFromLending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"platformWallets\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"adminClaimPlatformFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"spenders\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isReset\",\"type\":\"bool\"}],\"name\":\"approveAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"claimPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structISmartWalletImplementation.GetExpectedInParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getExpectedIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structISmartWalletImplementation.GetExpectedInParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getExpectedInWithImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpact\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structISmartWalletImplementation.GetExpectedReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getExpectedReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structISmartWalletImplementation.GetExpectedReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getExpectedReturnWithImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpact\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"platformWalletFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setAdminFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structISmartWalletImplementation.SwapParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"lendingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structISmartWalletImplementation.SwapAndDepositParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapAndDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"lendingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"},{\"internalType\":\"enumISmartWalletImplementation.FeeMode\",\"name\":\"feeMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structISmartWalletImplementation.SwapAndRepayParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapAndRepay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"lendingContract\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"}],\"internalType\":\"structISmartWalletImplementation.WithdrawFromLendingPlatformParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"withdrawFromLendingPlatform\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// GenericSmartWallet is an auto generated Go binding around an Ethereum contract.
type GenericSmartWallet struct {
	GenericSmartWalletCaller     // Read-only binding to the contract
	GenericSmartWalletTransactor // Write-only binding to the contract
	GenericSmartWalletFilterer   // Log filterer for contract events
}

// GenericSmartWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericSmartWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericSmartWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericSmartWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericSmartWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericSmartWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericSmartWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericSmartWalletSession struct {
	Contract     *GenericSmartWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GenericSmartWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericSmartWalletCallerSession struct {
	Contract *GenericSmartWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GenericSmartWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericSmartWalletTransactorSession struct {
	Contract     *GenericSmartWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GenericSmartWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericSmartWalletRaw struct {
	Contract *GenericSmartWallet // Generic contract binding to access the raw methods on
}

// GenericSmartWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericSmartWalletCallerRaw struct {
	Contract *GenericSmartWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GenericSmartWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericSmartWalletTransactorRaw struct {
	Contract *GenericSmartWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericSmartWallet creates a new instance of GenericSmartWallet, bound to a specific deployed contract.
func NewGenericSmartWallet(address common.Address, backend bind.ContractBackend) (*GenericSmartWallet, error) {
	contract, err := bindGenericSmartWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWallet{GenericSmartWalletCaller: GenericSmartWalletCaller{contract: contract}, GenericSmartWalletTransactor: GenericSmartWalletTransactor{contract: contract}, GenericSmartWalletFilterer: GenericSmartWalletFilterer{contract: contract}}, nil
}

// NewGenericSmartWalletCaller creates a new read-only instance of GenericSmartWallet, bound to a specific deployed contract.
func NewGenericSmartWalletCaller(address common.Address, caller bind.ContractCaller) (*GenericSmartWalletCaller, error) {
	contract, err := bindGenericSmartWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletCaller{contract: contract}, nil
}

// NewGenericSmartWalletTransactor creates a new write-only instance of GenericSmartWallet, bound to a specific deployed contract.
func NewGenericSmartWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericSmartWalletTransactor, error) {
	contract, err := bindGenericSmartWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletTransactor{contract: contract}, nil
}

// NewGenericSmartWalletFilterer creates a new log filterer instance of GenericSmartWallet, bound to a specific deployed contract.
func NewGenericSmartWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericSmartWalletFilterer, error) {
	contract, err := bindGenericSmartWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletFilterer{contract: contract}, nil
}

// bindGenericSmartWallet binds a generic wrapper to an already deployed contract.
func bindGenericSmartWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericSmartWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericSmartWallet *GenericSmartWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericSmartWallet.Contract.GenericSmartWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericSmartWallet *GenericSmartWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.GenericSmartWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericSmartWallet *GenericSmartWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.GenericSmartWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericSmartWallet *GenericSmartWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericSmartWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericSmartWallet *GenericSmartWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericSmartWallet *GenericSmartWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.contract.Transact(opts, method, params...)
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_GenericSmartWallet *GenericSmartWalletCaller) BPS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericSmartWallet.contract.Call(opts, out, "BPS")
	return *ret0, err
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_GenericSmartWallet *GenericSmartWalletSession) BPS() (*big.Int, error) {
	return _GenericSmartWallet.Contract.BPS(&_GenericSmartWallet.CallOpts)
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) BPS() (*big.Int, error) {
	return _GenericSmartWallet.Contract.BPS(&_GenericSmartWallet.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericSmartWallet.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletSession) Admin() (common.Address, error) {
	return _GenericSmartWallet.Contract.Admin(&_GenericSmartWallet.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) Admin() (common.Address, error) {
	return _GenericSmartWallet.Contract.Admin(&_GenericSmartWallet.CallOpts)
}

// AdminFeeCollector is a free data retrieval call binding the contract method 0x13bc1049.
//
// Solidity: function adminFeeCollector() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletCaller) AdminFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericSmartWallet.contract.Call(opts, out, "adminFeeCollector")
	return *ret0, err
}

// AdminFeeCollector is a free data retrieval call binding the contract method 0x13bc1049.
//
// Solidity: function adminFeeCollector() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletSession) AdminFeeCollector() (common.Address, error) {
	return _GenericSmartWallet.Contract.AdminFeeCollector(&_GenericSmartWallet.CallOpts)
}

// AdminFeeCollector is a free data retrieval call binding the contract method 0x13bc1049.
//
// Solidity: function adminFeeCollector() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) AdminFeeCollector() (common.Address, error) {
	return _GenericSmartWallet.Contract.AdminFeeCollector(&_GenericSmartWallet.CallOpts)
}

// GetExpectedIn is a free data retrieval call binding the contract method 0xd282662f.
//
// Solidity: function getExpectedIn((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 srcAmount, uint256 expectedRate)
func (_GenericSmartWallet *GenericSmartWalletCaller) GetExpectedIn(opts *bind.CallOpts, params ISmartWalletImplementationGetExpectedInParams) (struct {
	SrcAmount    *big.Int
	ExpectedRate *big.Int
}, error) {
	ret := new(struct {
		SrcAmount    *big.Int
		ExpectedRate *big.Int
	})
	out := ret
	err := _GenericSmartWallet.contract.Call(opts, out, "getExpectedIn", params)
	return *ret, err
}

// GetExpectedIn is a free data retrieval call binding the contract method 0xd282662f.
//
// Solidity: function getExpectedIn((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 srcAmount, uint256 expectedRate)
func (_GenericSmartWallet *GenericSmartWalletSession) GetExpectedIn(params ISmartWalletImplementationGetExpectedInParams) (struct {
	SrcAmount    *big.Int
	ExpectedRate *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedIn(&_GenericSmartWallet.CallOpts, params)
}

// GetExpectedIn is a free data retrieval call binding the contract method 0xd282662f.
//
// Solidity: function getExpectedIn((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 srcAmount, uint256 expectedRate)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) GetExpectedIn(params ISmartWalletImplementationGetExpectedInParams) (struct {
	SrcAmount    *big.Int
	ExpectedRate *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedIn(&_GenericSmartWallet.CallOpts, params)
}

// GetExpectedInWithImpact is a free data retrieval call binding the contract method 0x64b2b1b2.
//
// Solidity: function getExpectedInWithImpact((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 srcAmount, uint256 expectedRate, uint256 priceImpact)
func (_GenericSmartWallet *GenericSmartWalletCaller) GetExpectedInWithImpact(opts *bind.CallOpts, params ISmartWalletImplementationGetExpectedInParams) (struct {
	SrcAmount    *big.Int
	ExpectedRate *big.Int
	PriceImpact  *big.Int
}, error) {
	ret := new(struct {
		SrcAmount    *big.Int
		ExpectedRate *big.Int
		PriceImpact  *big.Int
	})
	out := ret
	err := _GenericSmartWallet.contract.Call(opts, out, "getExpectedInWithImpact", params)
	return *ret, err
}

// GetExpectedInWithImpact is a free data retrieval call binding the contract method 0x64b2b1b2.
//
// Solidity: function getExpectedInWithImpact((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 srcAmount, uint256 expectedRate, uint256 priceImpact)
func (_GenericSmartWallet *GenericSmartWalletSession) GetExpectedInWithImpact(params ISmartWalletImplementationGetExpectedInParams) (struct {
	SrcAmount    *big.Int
	ExpectedRate *big.Int
	PriceImpact  *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedInWithImpact(&_GenericSmartWallet.CallOpts, params)
}

// GetExpectedInWithImpact is a free data retrieval call binding the contract method 0x64b2b1b2.
//
// Solidity: function getExpectedInWithImpact((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 srcAmount, uint256 expectedRate, uint256 priceImpact)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) GetExpectedInWithImpact(params ISmartWalletImplementationGetExpectedInParams) (struct {
	SrcAmount    *big.Int
	ExpectedRate *big.Int
	PriceImpact  *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedInWithImpact(&_GenericSmartWallet.CallOpts, params)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0xe28c4135.
//
// Solidity: function getExpectedReturn((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 destAmount, uint256 expectedRate)
func (_GenericSmartWallet *GenericSmartWalletCaller) GetExpectedReturn(opts *bind.CallOpts, params ISmartWalletImplementationGetExpectedReturnParams) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	ret := new(struct {
		DestAmount   *big.Int
		ExpectedRate *big.Int
	})
	out := ret
	err := _GenericSmartWallet.contract.Call(opts, out, "getExpectedReturn", params)
	return *ret, err
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0xe28c4135.
//
// Solidity: function getExpectedReturn((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 destAmount, uint256 expectedRate)
func (_GenericSmartWallet *GenericSmartWalletSession) GetExpectedReturn(params ISmartWalletImplementationGetExpectedReturnParams) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedReturn(&_GenericSmartWallet.CallOpts, params)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0xe28c4135.
//
// Solidity: function getExpectedReturn((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 destAmount, uint256 expectedRate)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) GetExpectedReturn(params ISmartWalletImplementationGetExpectedReturnParams) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedReturn(&_GenericSmartWallet.CallOpts, params)
}

// GetExpectedReturnWithImpact is a free data retrieval call binding the contract method 0xbf3e3f06.
//
// Solidity: function getExpectedReturnWithImpact((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 destAmount, uint256 expectedRate, uint256 priceImpact)
func (_GenericSmartWallet *GenericSmartWalletCaller) GetExpectedReturnWithImpact(opts *bind.CallOpts, params ISmartWalletImplementationGetExpectedReturnParams) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
	PriceImpact  *big.Int
}, error) {
	ret := new(struct {
		DestAmount   *big.Int
		ExpectedRate *big.Int
		PriceImpact  *big.Int
	})
	out := ret
	err := _GenericSmartWallet.contract.Call(opts, out, "getExpectedReturnWithImpact", params)
	return *ret, err
}

// GetExpectedReturnWithImpact is a free data retrieval call binding the contract method 0xbf3e3f06.
//
// Solidity: function getExpectedReturnWithImpact((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 destAmount, uint256 expectedRate, uint256 priceImpact)
func (_GenericSmartWallet *GenericSmartWalletSession) GetExpectedReturnWithImpact(params ISmartWalletImplementationGetExpectedReturnParams) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
	PriceImpact  *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedReturnWithImpact(&_GenericSmartWallet.CallOpts, params)
}

// GetExpectedReturnWithImpact is a free data retrieval call binding the contract method 0xbf3e3f06.
//
// Solidity: function getExpectedReturnWithImpact((address,uint256,address[],uint8,uint256,bytes) params) view returns(uint256 destAmount, uint256 expectedRate, uint256 priceImpact)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) GetExpectedReturnWithImpact(params ISmartWalletImplementationGetExpectedReturnParams) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
	PriceImpact  *big.Int
}, error) {
	return _GenericSmartWallet.Contract.GetExpectedReturnWithImpact(&_GenericSmartWallet.CallOpts, params)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericSmartWallet.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletSession) PendingAdmin() (common.Address, error) {
	return _GenericSmartWallet.Contract.PendingAdmin(&_GenericSmartWallet.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) PendingAdmin() (common.Address, error) {
	return _GenericSmartWallet.Contract.PendingAdmin(&_GenericSmartWallet.CallOpts)
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_GenericSmartWallet *GenericSmartWalletCaller) PlatformWalletFees(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericSmartWallet.contract.Call(opts, out, "platformWalletFees", arg0, arg1)
	return *ret0, err
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_GenericSmartWallet *GenericSmartWalletSession) PlatformWalletFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GenericSmartWallet.Contract.PlatformWalletFees(&_GenericSmartWallet.CallOpts, arg0, arg1)
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_GenericSmartWallet *GenericSmartWalletCallerSession) PlatformWalletFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GenericSmartWallet.Contract.PlatformWalletFees(&_GenericSmartWallet.CallOpts, arg0, arg1)
}

// AdminClaimPlatformFees is a paid mutator transaction binding the contract method 0x78e58e0f.
//
// Solidity: function adminClaimPlatformFees(address[] platformWallets, address[] tokens) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) AdminClaimPlatformFees(opts *bind.TransactOpts, platformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "adminClaimPlatformFees", platformWallets, tokens)
}

// AdminClaimPlatformFees is a paid mutator transaction binding the contract method 0x78e58e0f.
//
// Solidity: function adminClaimPlatformFees(address[] platformWallets, address[] tokens) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) AdminClaimPlatformFees(platformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.AdminClaimPlatformFees(&_GenericSmartWallet.TransactOpts, platformWallets, tokens)
}

// AdminClaimPlatformFees is a paid mutator transaction binding the contract method 0x78e58e0f.
//
// Solidity: function adminClaimPlatformFees(address[] platformWallets, address[] tokens) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) AdminClaimPlatformFees(platformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.AdminClaimPlatformFees(&_GenericSmartWallet.TransactOpts, platformWallets, tokens)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) ApproveAllowances(opts *bind.TransactOpts, tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "approveAllowances", tokens, spenders, isReset)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) ApproveAllowances(tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.ApproveAllowances(&_GenericSmartWallet.TransactOpts, tokens, spenders, isReset)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) ApproveAllowances(tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.ApproveAllowances(&_GenericSmartWallet.TransactOpts, tokens, spenders, isReset)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_GenericSmartWallet *GenericSmartWalletSession) ClaimAdmin() (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.ClaimAdmin(&_GenericSmartWallet.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.ClaimAdmin(&_GenericSmartWallet.TransactOpts)
}

// ClaimPlatformFee is a paid mutator transaction binding the contract method 0x8be509a6.
//
// Solidity: function claimPlatformFee(address[] tokens) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) ClaimPlatformFee(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "claimPlatformFee", tokens)
}

// ClaimPlatformFee is a paid mutator transaction binding the contract method 0x8be509a6.
//
// Solidity: function claimPlatformFee(address[] tokens) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) ClaimPlatformFee(tokens []common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.ClaimPlatformFee(&_GenericSmartWallet.TransactOpts, tokens)
}

// ClaimPlatformFee is a paid mutator transaction binding the contract method 0x8be509a6.
//
// Solidity: function claimPlatformFee(address[] tokens) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) ClaimPlatformFee(tokens []common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.ClaimPlatformFee(&_GenericSmartWallet.TransactOpts, tokens)
}

// SetAdminFeeCollector is a paid mutator transaction binding the contract method 0xe941027e.
//
// Solidity: function setAdminFeeCollector(address feeCollector) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) SetAdminFeeCollector(opts *bind.TransactOpts, feeCollector common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "setAdminFeeCollector", feeCollector)
}

// SetAdminFeeCollector is a paid mutator transaction binding the contract method 0xe941027e.
//
// Solidity: function setAdminFeeCollector(address feeCollector) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) SetAdminFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.SetAdminFeeCollector(&_GenericSmartWallet.TransactOpts, feeCollector)
}

// SetAdminFeeCollector is a paid mutator transaction binding the contract method 0xe941027e.
//
// Solidity: function setAdminFeeCollector(address feeCollector) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) SetAdminFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.SetAdminFeeCollector(&_GenericSmartWallet.TransactOpts, feeCollector)
}

// Swap is a paid mutator transaction binding the contract method 0x2db897d0.
//
// Solidity: function swap((address,uint256,uint256,address[],uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactor) Swap(opts *bind.TransactOpts, params ISmartWalletImplementationSwapParams) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "swap", params)
}

// Swap is a paid mutator transaction binding the contract method 0x2db897d0.
//
// Solidity: function swap((address,uint256,uint256,address[],uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletSession) Swap(params ISmartWalletImplementationSwapParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.Swap(&_GenericSmartWallet.TransactOpts, params)
}

// Swap is a paid mutator transaction binding the contract method 0x2db897d0.
//
// Solidity: function swap((address,uint256,uint256,address[],uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) Swap(params ISmartWalletImplementationSwapParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.Swap(&_GenericSmartWallet.TransactOpts, params)
}

// SwapAndDeposit is a paid mutator transaction binding the contract method 0xa03f2a0b.
//
// Solidity: function swapAndDeposit((address,address,uint256,uint256,address[],uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactor) SwapAndDeposit(opts *bind.TransactOpts, params ISmartWalletImplementationSwapAndDepositParams) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "swapAndDeposit", params)
}

// SwapAndDeposit is a paid mutator transaction binding the contract method 0xa03f2a0b.
//
// Solidity: function swapAndDeposit((address,address,uint256,uint256,address[],uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletSession) SwapAndDeposit(params ISmartWalletImplementationSwapAndDepositParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.SwapAndDeposit(&_GenericSmartWallet.TransactOpts, params)
}

// SwapAndDeposit is a paid mutator transaction binding the contract method 0xa03f2a0b.
//
// Solidity: function swapAndDeposit((address,address,uint256,uint256,address[],uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) SwapAndDeposit(params ISmartWalletImplementationSwapAndDepositParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.SwapAndDeposit(&_GenericSmartWallet.TransactOpts, params)
}

// SwapAndRepay is a paid mutator transaction binding the contract method 0xfcd1f148.
//
// Solidity: function swapAndRepay((address,address,uint256,uint256,address[],uint256,uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactor) SwapAndRepay(opts *bind.TransactOpts, params ISmartWalletImplementationSwapAndRepayParams) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "swapAndRepay", params)
}

// SwapAndRepay is a paid mutator transaction binding the contract method 0xfcd1f148.
//
// Solidity: function swapAndRepay((address,address,uint256,uint256,address[],uint256,uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletSession) SwapAndRepay(params ISmartWalletImplementationSwapAndRepayParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.SwapAndRepay(&_GenericSmartWallet.TransactOpts, params)
}

// SwapAndRepay is a paid mutator transaction binding the contract method 0xfcd1f148.
//
// Solidity: function swapAndRepay((address,address,uint256,uint256,address[],uint256,uint8,uint256,address,bytes) params) payable returns(uint256 destAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) SwapAndRepay(params ISmartWalletImplementationSwapAndRepayParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.SwapAndRepay(&_GenericSmartWallet.TransactOpts, params)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.TransferAdmin(&_GenericSmartWallet.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.TransferAdmin(&_GenericSmartWallet.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.TransferAdminQuickly(&_GenericSmartWallet.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.TransferAdminQuickly(&_GenericSmartWallet.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.WithdrawEther(&_GenericSmartWallet.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.WithdrawEther(&_GenericSmartWallet.TransactOpts, amount, sendTo)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x23fcd0c9.
//
// Solidity: function withdrawFromLendingPlatform((address,address,uint256,uint256) params) returns(uint256 returnedAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactor) WithdrawFromLendingPlatform(opts *bind.TransactOpts, params ISmartWalletImplementationWithdrawFromLendingPlatformParams) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "withdrawFromLendingPlatform", params)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x23fcd0c9.
//
// Solidity: function withdrawFromLendingPlatform((address,address,uint256,uint256) params) returns(uint256 returnedAmount)
func (_GenericSmartWallet *GenericSmartWalletSession) WithdrawFromLendingPlatform(params ISmartWalletImplementationWithdrawFromLendingPlatformParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.WithdrawFromLendingPlatform(&_GenericSmartWallet.TransactOpts, params)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x23fcd0c9.
//
// Solidity: function withdrawFromLendingPlatform((address,address,uint256,uint256) params) returns(uint256 returnedAmount)
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) WithdrawFromLendingPlatform(params ISmartWalletImplementationWithdrawFromLendingPlatformParams) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.WithdrawFromLendingPlatform(&_GenericSmartWallet.TransactOpts, params)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_GenericSmartWallet *GenericSmartWalletSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.WithdrawToken(&_GenericSmartWallet.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.WithdrawToken(&_GenericSmartWallet.TransactOpts, token, amount, sendTo)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GenericSmartWallet *GenericSmartWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericSmartWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GenericSmartWallet *GenericSmartWalletSession) Receive() (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.Receive(&_GenericSmartWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GenericSmartWallet *GenericSmartWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _GenericSmartWallet.Contract.Receive(&_GenericSmartWallet.TransactOpts)
}

// GenericSmartWalletAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the GenericSmartWallet contract.
type GenericSmartWalletAdminClaimedIterator struct {
	Event *GenericSmartWalletAdminClaimed // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletAdminClaimed)
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
		it.Event = new(GenericSmartWalletAdminClaimed)
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
func (it *GenericSmartWalletAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletAdminClaimed represents a AdminClaimed event raised by the GenericSmartWallet contract.
type GenericSmartWalletAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*GenericSmartWalletAdminClaimedIterator, error) {

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletAdminClaimedIterator{contract: _GenericSmartWallet.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletAdminClaimed)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseAdminClaimed(log types.Log) (*GenericSmartWalletAdminClaimed, error) {
	event := new(GenericSmartWalletAdminClaimed)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletApprovedAllowancesIterator is returned from FilterApprovedAllowances and is used to iterate over the raw logs and unpacked data for ApprovedAllowances events raised by the GenericSmartWallet contract.
type GenericSmartWalletApprovedAllowancesIterator struct {
	Event *GenericSmartWalletApprovedAllowances // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletApprovedAllowancesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletApprovedAllowances)
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
		it.Event = new(GenericSmartWalletApprovedAllowances)
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
func (it *GenericSmartWalletApprovedAllowancesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletApprovedAllowancesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletApprovedAllowances represents a ApprovedAllowances event raised by the GenericSmartWallet contract.
type GenericSmartWalletApprovedAllowances struct {
	Tokens   []common.Address
	Spenders []common.Address
	IsReset  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovedAllowances is a free log retrieval operation binding the contract event 0xf34c5ed704407ea33d210cd1c76959be869adc80531ee3b3c93229fb606ac16e.
//
// Solidity: event ApprovedAllowances(address[] tokens, address[] spenders, bool isReset)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterApprovedAllowances(opts *bind.FilterOpts) (*GenericSmartWalletApprovedAllowancesIterator, error) {

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "ApprovedAllowances")
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletApprovedAllowancesIterator{contract: _GenericSmartWallet.contract, event: "ApprovedAllowances", logs: logs, sub: sub}, nil
}

// WatchApprovedAllowances is a free log subscription operation binding the contract event 0xf34c5ed704407ea33d210cd1c76959be869adc80531ee3b3c93229fb606ac16e.
//
// Solidity: event ApprovedAllowances(address[] tokens, address[] spenders, bool isReset)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchApprovedAllowances(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletApprovedAllowances) (event.Subscription, error) {

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "ApprovedAllowances")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletApprovedAllowances)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "ApprovedAllowances", log); err != nil {
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

// ParseApprovedAllowances is a log parse operation binding the contract event 0xf34c5ed704407ea33d210cd1c76959be869adc80531ee3b3c93229fb606ac16e.
//
// Solidity: event ApprovedAllowances(address[] tokens, address[] spenders, bool isReset)
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseApprovedAllowances(log types.Log) (*GenericSmartWalletApprovedAllowances, error) {
	event := new(GenericSmartWalletApprovedAllowances)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "ApprovedAllowances", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletClaimedPlatformFeesIterator is returned from FilterClaimedPlatformFees and is used to iterate over the raw logs and unpacked data for ClaimedPlatformFees events raised by the GenericSmartWallet contract.
type GenericSmartWalletClaimedPlatformFeesIterator struct {
	Event *GenericSmartWalletClaimedPlatformFees // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletClaimedPlatformFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletClaimedPlatformFees)
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
		it.Event = new(GenericSmartWalletClaimedPlatformFees)
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
func (it *GenericSmartWalletClaimedPlatformFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletClaimedPlatformFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletClaimedPlatformFees represents a ClaimedPlatformFees event raised by the GenericSmartWallet contract.
type GenericSmartWalletClaimedPlatformFees struct {
	Wallets []common.Address
	Tokens  []common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedPlatformFees is a free log retrieval operation binding the contract event 0xcbef5ab8ab58a3dba86704ce124241889d80c3a5fe3caaf4286b7e8c709a6c2d.
//
// Solidity: event ClaimedPlatformFees(address[] wallets, address[] tokens, address claimer)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterClaimedPlatformFees(opts *bind.FilterOpts) (*GenericSmartWalletClaimedPlatformFeesIterator, error) {

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "ClaimedPlatformFees")
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletClaimedPlatformFeesIterator{contract: _GenericSmartWallet.contract, event: "ClaimedPlatformFees", logs: logs, sub: sub}, nil
}

// WatchClaimedPlatformFees is a free log subscription operation binding the contract event 0xcbef5ab8ab58a3dba86704ce124241889d80c3a5fe3caaf4286b7e8c709a6c2d.
//
// Solidity: event ClaimedPlatformFees(address[] wallets, address[] tokens, address claimer)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchClaimedPlatformFees(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletClaimedPlatformFees) (event.Subscription, error) {

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "ClaimedPlatformFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletClaimedPlatformFees)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "ClaimedPlatformFees", log); err != nil {
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

// ParseClaimedPlatformFees is a log parse operation binding the contract event 0xcbef5ab8ab58a3dba86704ce124241889d80c3a5fe3caaf4286b7e8c709a6c2d.
//
// Solidity: event ClaimedPlatformFees(address[] wallets, address[] tokens, address claimer)
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseClaimedPlatformFees(log types.Log) (*GenericSmartWalletClaimedPlatformFees, error) {
	event := new(GenericSmartWalletClaimedPlatformFees)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "ClaimedPlatformFees", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the GenericSmartWallet contract.
type GenericSmartWalletEtherWithdrawIterator struct {
	Event *GenericSmartWalletEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletEtherWithdraw)
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
		it.Event = new(GenericSmartWalletEtherWithdraw)
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
func (it *GenericSmartWalletEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletEtherWithdraw represents a EtherWithdraw event raised by the GenericSmartWallet contract.
type GenericSmartWalletEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*GenericSmartWalletEtherWithdrawIterator, error) {

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletEtherWithdrawIterator{contract: _GenericSmartWallet.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletEtherWithdraw)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseEtherWithdraw(log types.Log) (*GenericSmartWalletEtherWithdraw, error) {
	event := new(GenericSmartWalletEtherWithdraw)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the GenericSmartWallet contract.
type GenericSmartWalletSwapIterator struct {
	Event *GenericSmartWalletSwap // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletSwap)
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
		it.Event = new(GenericSmartWalletSwap)
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
func (it *GenericSmartWalletSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletSwap represents a Swap event raised by the GenericSmartWallet contract.
type GenericSmartWalletSwap struct {
	Trader         common.Address
	SwapContract   common.Address
	TradePath      []common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	FeeMode        uint8
	FeeBps         *big.Int
	PlatformWallet common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xf23fb97e8642200d71532c9c5ae884a9236ca6b8eb58fed071a7cc3f2ac4161d.
//
// Solidity: event Swap(address indexed trader, address indexed swapContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterSwap(opts *bind.FilterOpts, trader []common.Address, swapContract []common.Address) (*GenericSmartWalletSwapIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var swapContractRule []interface{}
	for _, swapContractItem := range swapContract {
		swapContractRule = append(swapContractRule, swapContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "Swap", traderRule, swapContractRule)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletSwapIterator{contract: _GenericSmartWallet.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xf23fb97e8642200d71532c9c5ae884a9236ca6b8eb58fed071a7cc3f2ac4161d.
//
// Solidity: event Swap(address indexed trader, address indexed swapContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletSwap, trader []common.Address, swapContract []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var swapContractRule []interface{}
	for _, swapContractItem := range swapContract {
		swapContractRule = append(swapContractRule, swapContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "Swap", traderRule, swapContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletSwap)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xf23fb97e8642200d71532c9c5ae884a9236ca6b8eb58fed071a7cc3f2ac4161d.
//
// Solidity: event Swap(address indexed trader, address indexed swapContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseSwap(log types.Log) (*GenericSmartWalletSwap, error) {
	event := new(GenericSmartWalletSwap)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletSwapAndDepositIterator is returned from FilterSwapAndDeposit and is used to iterate over the raw logs and unpacked data for SwapAndDeposit events raised by the GenericSmartWallet contract.
type GenericSmartWalletSwapAndDepositIterator struct {
	Event *GenericSmartWalletSwapAndDeposit // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletSwapAndDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletSwapAndDeposit)
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
		it.Event = new(GenericSmartWalletSwapAndDeposit)
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
func (it *GenericSmartWalletSwapAndDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletSwapAndDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletSwapAndDeposit represents a SwapAndDeposit event raised by the GenericSmartWallet contract.
type GenericSmartWalletSwapAndDeposit struct {
	Trader          common.Address
	SwapContract    common.Address
	LendingContract common.Address
	TradePath       []common.Address
	SrcAmount       *big.Int
	DestAmount      *big.Int
	FeeMode         uint8
	FeeBps          *big.Int
	PlatformWallet  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSwapAndDeposit is a free log retrieval operation binding the contract event 0x2edafe685397c056fa62be1804788786cdeac9ce01277e8216c4a37180115723.
//
// Solidity: event SwapAndDeposit(address indexed trader, address indexed swapContract, address indexed lendingContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterSwapAndDeposit(opts *bind.FilterOpts, trader []common.Address, swapContract []common.Address, lendingContract []common.Address) (*GenericSmartWalletSwapAndDepositIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var swapContractRule []interface{}
	for _, swapContractItem := range swapContract {
		swapContractRule = append(swapContractRule, swapContractItem)
	}
	var lendingContractRule []interface{}
	for _, lendingContractItem := range lendingContract {
		lendingContractRule = append(lendingContractRule, lendingContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "SwapAndDeposit", traderRule, swapContractRule, lendingContractRule)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletSwapAndDepositIterator{contract: _GenericSmartWallet.contract, event: "SwapAndDeposit", logs: logs, sub: sub}, nil
}

// WatchSwapAndDeposit is a free log subscription operation binding the contract event 0x2edafe685397c056fa62be1804788786cdeac9ce01277e8216c4a37180115723.
//
// Solidity: event SwapAndDeposit(address indexed trader, address indexed swapContract, address indexed lendingContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchSwapAndDeposit(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletSwapAndDeposit, trader []common.Address, swapContract []common.Address, lendingContract []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var swapContractRule []interface{}
	for _, swapContractItem := range swapContract {
		swapContractRule = append(swapContractRule, swapContractItem)
	}
	var lendingContractRule []interface{}
	for _, lendingContractItem := range lendingContract {
		lendingContractRule = append(lendingContractRule, lendingContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "SwapAndDeposit", traderRule, swapContractRule, lendingContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletSwapAndDeposit)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "SwapAndDeposit", log); err != nil {
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

// ParseSwapAndDeposit is a log parse operation binding the contract event 0x2edafe685397c056fa62be1804788786cdeac9ce01277e8216c4a37180115723.
//
// Solidity: event SwapAndDeposit(address indexed trader, address indexed swapContract, address indexed lendingContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseSwapAndDeposit(log types.Log) (*GenericSmartWalletSwapAndDeposit, error) {
	event := new(GenericSmartWalletSwapAndDeposit)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "SwapAndDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletSwapAndRepayIterator is returned from FilterSwapAndRepay and is used to iterate over the raw logs and unpacked data for SwapAndRepay events raised by the GenericSmartWallet contract.
type GenericSmartWalletSwapAndRepayIterator struct {
	Event *GenericSmartWalletSwapAndRepay // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletSwapAndRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletSwapAndRepay)
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
		it.Event = new(GenericSmartWalletSwapAndRepay)
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
func (it *GenericSmartWalletSwapAndRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletSwapAndRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletSwapAndRepay represents a SwapAndRepay event raised by the GenericSmartWallet contract.
type GenericSmartWalletSwapAndRepay struct {
	Trader          common.Address
	SwapContract    common.Address
	LendingContract common.Address
	TradePath       []common.Address
	SrcAmount       *big.Int
	DestAmount      *big.Int
	PayAmount       *big.Int
	FeeMode         uint8
	FeeBps          *big.Int
	PlatformWallet  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSwapAndRepay is a free log retrieval operation binding the contract event 0x3b623be279df32a44c2edd16708b64a952bd614abdd2a3bc2572786eaefd5ab3.
//
// Solidity: event SwapAndRepay(address indexed trader, address indexed swapContract, address indexed lendingContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterSwapAndRepay(opts *bind.FilterOpts, trader []common.Address, swapContract []common.Address, lendingContract []common.Address) (*GenericSmartWalletSwapAndRepayIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var swapContractRule []interface{}
	for _, swapContractItem := range swapContract {
		swapContractRule = append(swapContractRule, swapContractItem)
	}
	var lendingContractRule []interface{}
	for _, lendingContractItem := range lendingContract {
		lendingContractRule = append(lendingContractRule, lendingContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "SwapAndRepay", traderRule, swapContractRule, lendingContractRule)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletSwapAndRepayIterator{contract: _GenericSmartWallet.contract, event: "SwapAndRepay", logs: logs, sub: sub}, nil
}

// WatchSwapAndRepay is a free log subscription operation binding the contract event 0x3b623be279df32a44c2edd16708b64a952bd614abdd2a3bc2572786eaefd5ab3.
//
// Solidity: event SwapAndRepay(address indexed trader, address indexed swapContract, address indexed lendingContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchSwapAndRepay(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletSwapAndRepay, trader []common.Address, swapContract []common.Address, lendingContract []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var swapContractRule []interface{}
	for _, swapContractItem := range swapContract {
		swapContractRule = append(swapContractRule, swapContractItem)
	}
	var lendingContractRule []interface{}
	for _, lendingContractItem := range lendingContract {
		lendingContractRule = append(lendingContractRule, lendingContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "SwapAndRepay", traderRule, swapContractRule, lendingContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletSwapAndRepay)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "SwapAndRepay", log); err != nil {
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

// ParseSwapAndRepay is a log parse operation binding the contract event 0x3b623be279df32a44c2edd16708b64a952bd614abdd2a3bc2572786eaefd5ab3.
//
// Solidity: event SwapAndRepay(address indexed trader, address indexed swapContract, address indexed lendingContract, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint8 feeMode, uint256 feeBps, address platformWallet)
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseSwapAndRepay(log types.Log) (*GenericSmartWalletSwapAndRepay, error) {
	event := new(GenericSmartWalletSwapAndRepay)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "SwapAndRepay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the GenericSmartWallet contract.
type GenericSmartWalletTokenWithdrawIterator struct {
	Event *GenericSmartWalletTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletTokenWithdraw)
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
		it.Event = new(GenericSmartWalletTokenWithdraw)
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
func (it *GenericSmartWalletTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletTokenWithdraw represents a TokenWithdraw event raised by the GenericSmartWallet contract.
type GenericSmartWalletTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*GenericSmartWalletTokenWithdrawIterator, error) {

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletTokenWithdrawIterator{contract: _GenericSmartWallet.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletTokenWithdraw)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseTokenWithdraw(log types.Log) (*GenericSmartWalletTokenWithdraw, error) {
	event := new(GenericSmartWalletTokenWithdraw)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the GenericSmartWallet contract.
type GenericSmartWalletTransferAdminPendingIterator struct {
	Event *GenericSmartWalletTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletTransferAdminPending)
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
		it.Event = new(GenericSmartWalletTransferAdminPending)
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
func (it *GenericSmartWalletTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletTransferAdminPending represents a TransferAdminPending event raised by the GenericSmartWallet contract.
type GenericSmartWalletTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*GenericSmartWalletTransferAdminPendingIterator, error) {

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletTransferAdminPendingIterator{contract: _GenericSmartWallet.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletTransferAdminPending)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseTransferAdminPending(log types.Log) (*GenericSmartWalletTransferAdminPending, error) {
	event := new(GenericSmartWalletTransferAdminPending)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GenericSmartWalletWithdrawFromLendingIterator is returned from FilterWithdrawFromLending and is used to iterate over the raw logs and unpacked data for WithdrawFromLending events raised by the GenericSmartWallet contract.
type GenericSmartWalletWithdrawFromLendingIterator struct {
	Event *GenericSmartWalletWithdrawFromLending // Event containing the contract specifics and raw log

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
func (it *GenericSmartWalletWithdrawFromLendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericSmartWalletWithdrawFromLending)
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
		it.Event = new(GenericSmartWalletWithdrawFromLending)
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
func (it *GenericSmartWalletWithdrawFromLendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericSmartWalletWithdrawFromLendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericSmartWalletWithdrawFromLending represents a WithdrawFromLending event raised by the GenericSmartWallet contract.
type GenericSmartWalletWithdrawFromLending struct {
	Trader             common.Address
	LendingContract    common.Address
	Token              common.Address
	Amount             *big.Int
	MinReturn          *big.Int
	ActualReturnAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFromLending is a free log retrieval operation binding the contract event 0x8535f320ef164ed88035c05c538ac8c71ab5fba9c64bc06794378060332042fc.
//
// Solidity: event WithdrawFromLending(address indexed trader, address indexed lendingContract, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount)
func (_GenericSmartWallet *GenericSmartWalletFilterer) FilterWithdrawFromLending(opts *bind.FilterOpts, trader []common.Address, lendingContract []common.Address) (*GenericSmartWalletWithdrawFromLendingIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var lendingContractRule []interface{}
	for _, lendingContractItem := range lendingContract {
		lendingContractRule = append(lendingContractRule, lendingContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.FilterLogs(opts, "WithdrawFromLending", traderRule, lendingContractRule)
	if err != nil {
		return nil, err
	}
	return &GenericSmartWalletWithdrawFromLendingIterator{contract: _GenericSmartWallet.contract, event: "WithdrawFromLending", logs: logs, sub: sub}, nil
}

// WatchWithdrawFromLending is a free log subscription operation binding the contract event 0x8535f320ef164ed88035c05c538ac8c71ab5fba9c64bc06794378060332042fc.
//
// Solidity: event WithdrawFromLending(address indexed trader, address indexed lendingContract, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount)
func (_GenericSmartWallet *GenericSmartWalletFilterer) WatchWithdrawFromLending(opts *bind.WatchOpts, sink chan<- *GenericSmartWalletWithdrawFromLending, trader []common.Address, lendingContract []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var lendingContractRule []interface{}
	for _, lendingContractItem := range lendingContract {
		lendingContractRule = append(lendingContractRule, lendingContractItem)
	}

	logs, sub, err := _GenericSmartWallet.contract.WatchLogs(opts, "WithdrawFromLending", traderRule, lendingContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericSmartWalletWithdrawFromLending)
				if err := _GenericSmartWallet.contract.UnpackLog(event, "WithdrawFromLending", log); err != nil {
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

// ParseWithdrawFromLending is a log parse operation binding the contract event 0x8535f320ef164ed88035c05c538ac8c71ab5fba9c64bc06794378060332042fc.
//
// Solidity: event WithdrawFromLending(address indexed trader, address indexed lendingContract, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount)
func (_GenericSmartWallet *GenericSmartWalletFilterer) ParseWithdrawFromLending(log types.Log) (*GenericSmartWalletWithdrawFromLending, error) {
	event := new(GenericSmartWalletWithdrawFromLending)
	if err := _GenericSmartWallet.contract.UnpackLog(event, "WithdrawFromLending", log); err != nil {
		return nil, err
	}
	return event, nil
}
