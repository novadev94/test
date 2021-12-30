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

// SmartWalletSwapBscABI is the input ABI used to generate the binding from.
const SmartWalletSwapBscABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIBEP20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"spenders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isReset\",\"type\":\"bool\"}],\"name\":\"ApprovedAllowances\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"BnbWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBorrowed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"}],\"name\":\"BorrowFromLending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"contractIBEP20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimedPlatformFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"KyberTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIBEP20\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"KyberTradeAndDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIBEP20\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"KyberTradeAndRepay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"feeInSrc\",\"type\":\"bool\"}],\"name\":\"PancakeTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"contractIPancakeRouter02\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"PancakeTradeAndDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"contractIPancakeRouter02\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAndRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"PancakeTradeAndRepay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractISmartWalletLending\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"UpdatedLendingImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"UpdatedSupportedPlatformWallets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualReturnAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFromLending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"spenders\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isReset\",\"type\":\"bool\"}],\"name\":\"approveAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"platformWallets\",\"type\":\"address[]\"},{\"internalType\":\"contractIBEP20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"claimPlatformFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturnKyber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturnPancake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRouterSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberProxy\",\"outputs\":[{\"internalType\":\"contractIKyberProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingImpl\",\"outputs\":[{\"internalType\":\"contractISmartWalletLending\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"platformWalletFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedPlatformWallets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"swapKyber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIBEP20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"swapKyberAndDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIBEP20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"swapKyberAndRepay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"feeInSrc\",\"type\":\"bool\"}],\"name\":\"swapPancake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"swapPancakeAndDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"feeAndRateMode\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"swapPancakeAndRepay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISmartWalletLending\",\"name\":\"newImpl\",\"type\":\"address\"}],\"name\":\"updateLendingImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"updateSupportedPlatformWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawBnb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"}],\"name\":\"withdrawFromLendingPlatform\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SmartWalletSwapBsc is an auto generated Go binding around an Ethereum contract.
type SmartWalletSwapBsc struct {
	SmartWalletSwapBscCaller     // Read-only binding to the contract
	SmartWalletSwapBscTransactor // Write-only binding to the contract
	SmartWalletSwapBscFilterer   // Log filterer for contract events
}

// SmartWalletSwapBscCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartWalletSwapBscCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartWalletSwapBscTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartWalletSwapBscTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartWalletSwapBscFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartWalletSwapBscFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartWalletSwapBscSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartWalletSwapBscSession struct {
	Contract     *SmartWalletSwapBsc // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SmartWalletSwapBscCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartWalletSwapBscCallerSession struct {
	Contract *SmartWalletSwapBscCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SmartWalletSwapBscTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartWalletSwapBscTransactorSession struct {
	Contract     *SmartWalletSwapBscTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SmartWalletSwapBscRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartWalletSwapBscRaw struct {
	Contract *SmartWalletSwapBsc // Generic contract binding to access the raw methods on
}

// SmartWalletSwapBscCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartWalletSwapBscCallerRaw struct {
	Contract *SmartWalletSwapBscCaller // Generic read-only contract binding to access the raw methods on
}

// SmartWalletSwapBscTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartWalletSwapBscTransactorRaw struct {
	Contract *SmartWalletSwapBscTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartWalletSwapBsc creates a new instance of SmartWalletSwapBsc, bound to a specific deployed contract.
func NewSmartWalletSwapBsc(address common.Address, backend bind.ContractBackend) (*SmartWalletSwapBsc, error) {
	contract, err := bindSmartWalletSwapBsc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBsc{SmartWalletSwapBscCaller: SmartWalletSwapBscCaller{contract: contract}, SmartWalletSwapBscTransactor: SmartWalletSwapBscTransactor{contract: contract}, SmartWalletSwapBscFilterer: SmartWalletSwapBscFilterer{contract: contract}}, nil
}

// NewSmartWalletSwapBscCaller creates a new read-only instance of SmartWalletSwapBsc, bound to a specific deployed contract.
func NewSmartWalletSwapBscCaller(address common.Address, caller bind.ContractCaller) (*SmartWalletSwapBscCaller, error) {
	contract, err := bindSmartWalletSwapBsc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscCaller{contract: contract}, nil
}

// NewSmartWalletSwapBscTransactor creates a new write-only instance of SmartWalletSwapBsc, bound to a specific deployed contract.
func NewSmartWalletSwapBscTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartWalletSwapBscTransactor, error) {
	contract, err := bindSmartWalletSwapBsc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscTransactor{contract: contract}, nil
}

// NewSmartWalletSwapBscFilterer creates a new log filterer instance of SmartWalletSwapBsc, bound to a specific deployed contract.
func NewSmartWalletSwapBscFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartWalletSwapBscFilterer, error) {
	contract, err := bindSmartWalletSwapBsc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscFilterer{contract: contract}, nil
}

// bindSmartWalletSwapBsc binds a generic wrapper to an already deployed contract.
func bindSmartWalletSwapBsc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartWalletSwapBscABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartWalletSwapBsc *SmartWalletSwapBscRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartWalletSwapBsc.Contract.SmartWalletSwapBscCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartWalletSwapBsc *SmartWalletSwapBscRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SmartWalletSwapBscTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartWalletSwapBsc *SmartWalletSwapBscRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SmartWalletSwapBscTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartWalletSwapBsc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.contract.Transact(opts, method, params...)
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) BPS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "BPS")
	return *ret0, err
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) BPS() (*big.Int, error) {
	return _SmartWalletSwapBsc.Contract.BPS(&_SmartWalletSwapBsc.CallOpts)
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) BPS() (*big.Int, error) {
	return _SmartWalletSwapBsc.Contract.BPS(&_SmartWalletSwapBsc.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) Admin() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.Admin(&_SmartWalletSwapBsc.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) Admin() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.Admin(&_SmartWalletSwapBsc.CallOpts)
}

// GetExpectedReturnKyber is a free data retrieval call binding the contract method 0x15202c43.
//
// Solidity: function getExpectedReturnKyber(address src, address dest, uint256 srcAmount) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) GetExpectedReturnKyber(opts *bind.CallOpts, src common.Address, dest common.Address, srcAmount *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	ret := new(struct {
		DestAmount   *big.Int
		ExpectedRate *big.Int
	})
	out := ret
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "getExpectedReturnKyber", src, dest, srcAmount)
	return *ret, err
}

// GetExpectedReturnKyber is a free data retrieval call binding the contract method 0x15202c43.
//
// Solidity: function getExpectedReturnKyber(address src, address dest, uint256 srcAmount) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) GetExpectedReturnKyber(src common.Address, dest common.Address, srcAmount *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwapBsc.Contract.GetExpectedReturnKyber(&_SmartWalletSwapBsc.CallOpts, src, dest, srcAmount)
}

// GetExpectedReturnKyber is a free data retrieval call binding the contract method 0x15202c43.
//
// Solidity: function getExpectedReturnKyber(address src, address dest, uint256 srcAmount) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) GetExpectedReturnKyber(src common.Address, dest common.Address, srcAmount *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwapBsc.Contract.GetExpectedReturnKyber(&_SmartWalletSwapBsc.CallOpts, src, dest, srcAmount)
}

// GetExpectedReturnPancake is a free data retrieval call binding the contract method 0xfbd4002f.
//
// Solidity: function getExpectedReturnPancake(address router, uint256 srcAmount, address[] tradePath, uint256 platformFee) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) GetExpectedReturnPancake(opts *bind.CallOpts, router common.Address, srcAmount *big.Int, tradePath []common.Address, platformFee *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	ret := new(struct {
		DestAmount   *big.Int
		ExpectedRate *big.Int
	})
	out := ret
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "getExpectedReturnPancake", router, srcAmount, tradePath, platformFee)
	return *ret, err
}

// GetExpectedReturnPancake is a free data retrieval call binding the contract method 0xfbd4002f.
//
// Solidity: function getExpectedReturnPancake(address router, uint256 srcAmount, address[] tradePath, uint256 platformFee) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) GetExpectedReturnPancake(router common.Address, srcAmount *big.Int, tradePath []common.Address, platformFee *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwapBsc.Contract.GetExpectedReturnPancake(&_SmartWalletSwapBsc.CallOpts, router, srcAmount, tradePath, platformFee)
}

// GetExpectedReturnPancake is a free data retrieval call binding the contract method 0xfbd4002f.
//
// Solidity: function getExpectedReturnPancake(address router, uint256 srcAmount, address[] tradePath, uint256 platformFee) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) GetExpectedReturnPancake(router common.Address, srcAmount *big.Int, tradePath []common.Address, platformFee *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwapBsc.Contract.GetExpectedReturnPancake(&_SmartWalletSwapBsc.CallOpts, router, srcAmount, tradePath, platformFee)
}

// IsRouterSupported is a free data retrieval call binding the contract method 0xc277b57f.
//
// Solidity: function isRouterSupported(address ) view returns(bool)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) IsRouterSupported(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "isRouterSupported", arg0)
	return *ret0, err
}

// IsRouterSupported is a free data retrieval call binding the contract method 0xc277b57f.
//
// Solidity: function isRouterSupported(address ) view returns(bool)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) IsRouterSupported(arg0 common.Address) (bool, error) {
	return _SmartWalletSwapBsc.Contract.IsRouterSupported(&_SmartWalletSwapBsc.CallOpts, arg0)
}

// IsRouterSupported is a free data retrieval call binding the contract method 0xc277b57f.
//
// Solidity: function isRouterSupported(address ) view returns(bool)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) IsRouterSupported(arg0 common.Address) (bool, error) {
	return _SmartWalletSwapBsc.Contract.IsRouterSupported(&_SmartWalletSwapBsc.CallOpts, arg0)
}

// KyberProxy is a free data retrieval call binding the contract method 0xf0eeed81.
//
// Solidity: function kyberProxy() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) KyberProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "kyberProxy")
	return *ret0, err
}

// KyberProxy is a free data retrieval call binding the contract method 0xf0eeed81.
//
// Solidity: function kyberProxy() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) KyberProxy() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.KyberProxy(&_SmartWalletSwapBsc.CallOpts)
}

// KyberProxy is a free data retrieval call binding the contract method 0xf0eeed81.
//
// Solidity: function kyberProxy() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) KyberProxy() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.KyberProxy(&_SmartWalletSwapBsc.CallOpts)
}

// LendingImpl is a free data retrieval call binding the contract method 0x8690948b.
//
// Solidity: function lendingImpl() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) LendingImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "lendingImpl")
	return *ret0, err
}

// LendingImpl is a free data retrieval call binding the contract method 0x8690948b.
//
// Solidity: function lendingImpl() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) LendingImpl() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.LendingImpl(&_SmartWalletSwapBsc.CallOpts)
}

// LendingImpl is a free data retrieval call binding the contract method 0x8690948b.
//
// Solidity: function lendingImpl() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) LendingImpl() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.LendingImpl(&_SmartWalletSwapBsc.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) PendingAdmin() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.PendingAdmin(&_SmartWalletSwapBsc.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) PendingAdmin() (common.Address, error) {
	return _SmartWalletSwapBsc.Contract.PendingAdmin(&_SmartWalletSwapBsc.CallOpts)
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) PlatformWalletFees(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "platformWalletFees", arg0, arg1)
	return *ret0, err
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) PlatformWalletFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SmartWalletSwapBsc.Contract.PlatformWalletFees(&_SmartWalletSwapBsc.CallOpts, arg0, arg1)
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) PlatformWalletFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SmartWalletSwapBsc.Contract.PlatformWalletFees(&_SmartWalletSwapBsc.CallOpts, arg0, arg1)
}

// SupportedPlatformWallets is a free data retrieval call binding the contract method 0xdc6f6428.
//
// Solidity: function supportedPlatformWallets(address ) view returns(bool)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCaller) SupportedPlatformWallets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartWalletSwapBsc.contract.Call(opts, out, "supportedPlatformWallets", arg0)
	return *ret0, err
}

// SupportedPlatformWallets is a free data retrieval call binding the contract method 0xdc6f6428.
//
// Solidity: function supportedPlatformWallets(address ) view returns(bool)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) SupportedPlatformWallets(arg0 common.Address) (bool, error) {
	return _SmartWalletSwapBsc.Contract.SupportedPlatformWallets(&_SmartWalletSwapBsc.CallOpts, arg0)
}

// SupportedPlatformWallets is a free data retrieval call binding the contract method 0xdc6f6428.
//
// Solidity: function supportedPlatformWallets(address ) view returns(bool)
func (_SmartWalletSwapBsc *SmartWalletSwapBscCallerSession) SupportedPlatformWallets(arg0 common.Address) (bool, error) {
	return _SmartWalletSwapBsc.Contract.SupportedPlatformWallets(&_SmartWalletSwapBsc.CallOpts, arg0)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) ApproveAllowances(opts *bind.TransactOpts, tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "approveAllowances", tokens, spenders, isReset)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) ApproveAllowances(tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.ApproveAllowances(&_SmartWalletSwapBsc.TransactOpts, tokens, spenders, isReset)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) ApproveAllowances(tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.ApproveAllowances(&_SmartWalletSwapBsc.TransactOpts, tokens, spenders, isReset)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) ClaimAdmin() (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.ClaimAdmin(&_SmartWalletSwapBsc.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.ClaimAdmin(&_SmartWalletSwapBsc.TransactOpts)
}

// ClaimPlatformFees is a paid mutator transaction binding the contract method 0x35c7c3cf.
//
// Solidity: function claimPlatformFees(address[] platformWallets, address[] tokens) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) ClaimPlatformFees(opts *bind.TransactOpts, platformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "claimPlatformFees", platformWallets, tokens)
}

// ClaimPlatformFees is a paid mutator transaction binding the contract method 0x35c7c3cf.
//
// Solidity: function claimPlatformFees(address[] platformWallets, address[] tokens) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) ClaimPlatformFees(platformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.ClaimPlatformFees(&_SmartWalletSwapBsc.TransactOpts, platformWallets, tokens)
}

// ClaimPlatformFees is a paid mutator transaction binding the contract method 0x35c7c3cf.
//
// Solidity: function claimPlatformFees(address[] platformWallets, address[] tokens) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) ClaimPlatformFees(platformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.ClaimPlatformFees(&_SmartWalletSwapBsc.TransactOpts, platformWallets, tokens)
}

// SwapKyber is a paid mutator transaction binding the contract method 0xffcad20b.
//
// Solidity: function swapKyber(address src, address dest, uint256 srcAmount, uint256 minConversionRate, address recipient, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) SwapKyber(opts *bind.TransactOpts, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, recipient common.Address, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "swapKyber", src, dest, srcAmount, minConversionRate, recipient, platformWallet)
}

// SwapKyber is a paid mutator transaction binding the contract method 0xffcad20b.
//
// Solidity: function swapKyber(address src, address dest, uint256 srcAmount, uint256 minConversionRate, address recipient, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) SwapKyber(src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, recipient common.Address, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapKyber(&_SmartWalletSwapBsc.TransactOpts, src, dest, srcAmount, minConversionRate, recipient, platformWallet)
}

// SwapKyber is a paid mutator transaction binding the contract method 0xffcad20b.
//
// Solidity: function swapKyber(address src, address dest, uint256 srcAmount, uint256 minConversionRate, address recipient, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) SwapKyber(src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, recipient common.Address, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapKyber(&_SmartWalletSwapBsc.TransactOpts, src, dest, srcAmount, minConversionRate, recipient, platformWallet)
}

// SwapKyberAndDeposit is a paid mutator transaction binding the contract method 0x67b78153.
//
// Solidity: function swapKyberAndDeposit(uint8 platform, address src, address dest, uint256 srcAmount, uint256 minConversionRate, uint256 platformFeeBps, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) SwapKyberAndDeposit(opts *bind.TransactOpts, platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, platformFeeBps *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "swapKyberAndDeposit", platform, src, dest, srcAmount, minConversionRate, platformFeeBps, platformWallet)
}

// SwapKyberAndDeposit is a paid mutator transaction binding the contract method 0x67b78153.
//
// Solidity: function swapKyberAndDeposit(uint8 platform, address src, address dest, uint256 srcAmount, uint256 minConversionRate, uint256 platformFeeBps, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) SwapKyberAndDeposit(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, platformFeeBps *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapKyberAndDeposit(&_SmartWalletSwapBsc.TransactOpts, platform, src, dest, srcAmount, minConversionRate, platformFeeBps, platformWallet)
}

// SwapKyberAndDeposit is a paid mutator transaction binding the contract method 0x67b78153.
//
// Solidity: function swapKyberAndDeposit(uint8 platform, address src, address dest, uint256 srcAmount, uint256 minConversionRate, uint256 platformFeeBps, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) SwapKyberAndDeposit(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, platformFeeBps *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapKyberAndDeposit(&_SmartWalletSwapBsc.TransactOpts, platform, src, dest, srcAmount, minConversionRate, platformFeeBps, platformWallet)
}

// SwapKyberAndRepay is a paid mutator transaction binding the contract method 0x53dd9b1a.
//
// Solidity: function swapKyberAndRepay(uint8 platform, address src, address dest, uint256 srcAmount, uint256 payAmount, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) SwapKyberAndRepay(opts *bind.TransactOpts, platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, payAmount *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "swapKyberAndRepay", platform, src, dest, srcAmount, payAmount, platformWallet)
}

// SwapKyberAndRepay is a paid mutator transaction binding the contract method 0x53dd9b1a.
//
// Solidity: function swapKyberAndRepay(uint8 platform, address src, address dest, uint256 srcAmount, uint256 payAmount, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) SwapKyberAndRepay(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, payAmount *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapKyberAndRepay(&_SmartWalletSwapBsc.TransactOpts, platform, src, dest, srcAmount, payAmount, platformWallet)
}

// SwapKyberAndRepay is a paid mutator transaction binding the contract method 0x53dd9b1a.
//
// Solidity: function swapKyberAndRepay(uint8 platform, address src, address dest, uint256 srcAmount, uint256 payAmount, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) SwapKyberAndRepay(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, payAmount *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapKyberAndRepay(&_SmartWalletSwapBsc.TransactOpts, platform, src, dest, srcAmount, payAmount, platformWallet)
}

// SwapPancake is a paid mutator transaction binding the contract method 0x7a6c0dfe.
//
// Solidity: function swapPancake(address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) SwapPancake(opts *bind.TransactOpts, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, feeInSrc bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "swapPancake", router, srcAmount, minDestAmount, tradePath, recipient, platformFeeBps, platformWallet, feeInSrc)
}

// SwapPancake is a paid mutator transaction binding the contract method 0x7a6c0dfe.
//
// Solidity: function swapPancake(address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) SwapPancake(router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, feeInSrc bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapPancake(&_SmartWalletSwapBsc.TransactOpts, router, srcAmount, minDestAmount, tradePath, recipient, platformFeeBps, platformWallet, feeInSrc)
}

// SwapPancake is a paid mutator transaction binding the contract method 0x7a6c0dfe.
//
// Solidity: function swapPancake(address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) SwapPancake(router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, feeInSrc bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapPancake(&_SmartWalletSwapBsc.TransactOpts, router, srcAmount, minDestAmount, tradePath, recipient, platformFeeBps, platformWallet, feeInSrc)
}

// SwapPancakeAndDeposit is a paid mutator transaction binding the contract method 0x951e22e9.
//
// Solidity: function swapPancakeAndDeposit(uint8 platform, address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, uint256 platformFeeBps, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) SwapPancakeAndDeposit(opts *bind.TransactOpts, platform uint8, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, platformFeeBps *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "swapPancakeAndDeposit", platform, router, srcAmount, minDestAmount, tradePath, platformFeeBps, platformWallet)
}

// SwapPancakeAndDeposit is a paid mutator transaction binding the contract method 0x951e22e9.
//
// Solidity: function swapPancakeAndDeposit(uint8 platform, address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, uint256 platformFeeBps, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) SwapPancakeAndDeposit(platform uint8, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, platformFeeBps *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapPancakeAndDeposit(&_SmartWalletSwapBsc.TransactOpts, platform, router, srcAmount, minDestAmount, tradePath, platformFeeBps, platformWallet)
}

// SwapPancakeAndDeposit is a paid mutator transaction binding the contract method 0x951e22e9.
//
// Solidity: function swapPancakeAndDeposit(uint8 platform, address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, uint256 platformFeeBps, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) SwapPancakeAndDeposit(platform uint8, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, platformFeeBps *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapPancakeAndDeposit(&_SmartWalletSwapBsc.TransactOpts, platform, router, srcAmount, minDestAmount, tradePath, platformFeeBps, platformWallet)
}

// SwapPancakeAndRepay is a paid mutator transaction binding the contract method 0x349adaf3.
//
// Solidity: function swapPancakeAndRepay(uint8 platform, address router, uint256 srcAmount, uint256 payAmount, address[] tradePath, uint256 feeAndRateMode, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) SwapPancakeAndRepay(opts *bind.TransactOpts, platform uint8, router common.Address, srcAmount *big.Int, payAmount *big.Int, tradePath []common.Address, feeAndRateMode *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "swapPancakeAndRepay", platform, router, srcAmount, payAmount, tradePath, feeAndRateMode, platformWallet)
}

// SwapPancakeAndRepay is a paid mutator transaction binding the contract method 0x349adaf3.
//
// Solidity: function swapPancakeAndRepay(uint8 platform, address router, uint256 srcAmount, uint256 payAmount, address[] tradePath, uint256 feeAndRateMode, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) SwapPancakeAndRepay(platform uint8, router common.Address, srcAmount *big.Int, payAmount *big.Int, tradePath []common.Address, feeAndRateMode *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapPancakeAndRepay(&_SmartWalletSwapBsc.TransactOpts, platform, router, srcAmount, payAmount, tradePath, feeAndRateMode, platformWallet)
}

// SwapPancakeAndRepay is a paid mutator transaction binding the contract method 0x349adaf3.
//
// Solidity: function swapPancakeAndRepay(uint8 platform, address router, uint256 srcAmount, uint256 payAmount, address[] tradePath, uint256 feeAndRateMode, address platformWallet) payable returns(uint256 destAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) SwapPancakeAndRepay(platform uint8, router common.Address, srcAmount *big.Int, payAmount *big.Int, tradePath []common.Address, feeAndRateMode *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.SwapPancakeAndRepay(&_SmartWalletSwapBsc.TransactOpts, platform, router, srcAmount, payAmount, tradePath, feeAndRateMode, platformWallet)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.TransferAdmin(&_SmartWalletSwapBsc.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.TransferAdmin(&_SmartWalletSwapBsc.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.TransferAdminQuickly(&_SmartWalletSwapBsc.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.TransferAdminQuickly(&_SmartWalletSwapBsc.TransactOpts, newAdmin)
}

// UpdateLendingImplementation is a paid mutator transaction binding the contract method 0xb4d59862.
//
// Solidity: function updateLendingImplementation(address newImpl) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) UpdateLendingImplementation(opts *bind.TransactOpts, newImpl common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "updateLendingImplementation", newImpl)
}

// UpdateLendingImplementation is a paid mutator transaction binding the contract method 0xb4d59862.
//
// Solidity: function updateLendingImplementation(address newImpl) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) UpdateLendingImplementation(newImpl common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.UpdateLendingImplementation(&_SmartWalletSwapBsc.TransactOpts, newImpl)
}

// UpdateLendingImplementation is a paid mutator transaction binding the contract method 0xb4d59862.
//
// Solidity: function updateLendingImplementation(address newImpl) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) UpdateLendingImplementation(newImpl common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.UpdateLendingImplementation(&_SmartWalletSwapBsc.TransactOpts, newImpl)
}

// UpdateSupportedPlatformWallets is a paid mutator transaction binding the contract method 0x296f4120.
//
// Solidity: function updateSupportedPlatformWallets(address[] wallets, bool isSupported) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) UpdateSupportedPlatformWallets(opts *bind.TransactOpts, wallets []common.Address, isSupported bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "updateSupportedPlatformWallets", wallets, isSupported)
}

// UpdateSupportedPlatformWallets is a paid mutator transaction binding the contract method 0x296f4120.
//
// Solidity: function updateSupportedPlatformWallets(address[] wallets, bool isSupported) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) UpdateSupportedPlatformWallets(wallets []common.Address, isSupported bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.UpdateSupportedPlatformWallets(&_SmartWalletSwapBsc.TransactOpts, wallets, isSupported)
}

// UpdateSupportedPlatformWallets is a paid mutator transaction binding the contract method 0x296f4120.
//
// Solidity: function updateSupportedPlatformWallets(address[] wallets, bool isSupported) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) UpdateSupportedPlatformWallets(wallets []common.Address, isSupported bool) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.UpdateSupportedPlatformWallets(&_SmartWalletSwapBsc.TransactOpts, wallets, isSupported)
}

// WithdrawBnb is a paid mutator transaction binding the contract method 0x960ed708.
//
// Solidity: function withdrawBnb(uint256 amount, address sendTo) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) WithdrawBnb(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "withdrawBnb", amount, sendTo)
}

// WithdrawBnb is a paid mutator transaction binding the contract method 0x960ed708.
//
// Solidity: function withdrawBnb(uint256 amount, address sendTo) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) WithdrawBnb(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.WithdrawBnb(&_SmartWalletSwapBsc.TransactOpts, amount, sendTo)
}

// WithdrawBnb is a paid mutator transaction binding the contract method 0x960ed708.
//
// Solidity: function withdrawBnb(uint256 amount, address sendTo) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) WithdrawBnb(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.WithdrawBnb(&_SmartWalletSwapBsc.TransactOpts, amount, sendTo)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x71ea95b0.
//
// Solidity: function withdrawFromLendingPlatform(uint8 platform, address token, uint256 amount, uint256 minReturn) returns(uint256 returnedAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) WithdrawFromLendingPlatform(opts *bind.TransactOpts, platform uint8, token common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "withdrawFromLendingPlatform", platform, token, amount, minReturn)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x71ea95b0.
//
// Solidity: function withdrawFromLendingPlatform(uint8 platform, address token, uint256 amount, uint256 minReturn) returns(uint256 returnedAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) WithdrawFromLendingPlatform(platform uint8, token common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.WithdrawFromLendingPlatform(&_SmartWalletSwapBsc.TransactOpts, platform, token, amount, minReturn)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x71ea95b0.
//
// Solidity: function withdrawFromLendingPlatform(uint8 platform, address token, uint256 amount, uint256 minReturn) returns(uint256 returnedAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) WithdrawFromLendingPlatform(platform uint8, token common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.WithdrawFromLendingPlatform(&_SmartWalletSwapBsc.TransactOpts, platform, token, amount, minReturn)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.WithdrawToken(&_SmartWalletSwapBsc.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.WithdrawToken(&_SmartWalletSwapBsc.TransactOpts, token, amount, sendTo)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwapBsc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscSession) Receive() (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.Receive(&_SmartWalletSwapBsc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartWalletSwapBsc *SmartWalletSwapBscTransactorSession) Receive() (*types.Transaction, error) {
	return _SmartWalletSwapBsc.Contract.Receive(&_SmartWalletSwapBsc.TransactOpts)
}

// SmartWalletSwapBscAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscAdminClaimedIterator struct {
	Event *SmartWalletSwapBscAdminClaimed // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscAdminClaimed)
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
		it.Event = new(SmartWalletSwapBscAdminClaimed)
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
func (it *SmartWalletSwapBscAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscAdminClaimed represents a AdminClaimed event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*SmartWalletSwapBscAdminClaimedIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscAdminClaimedIterator{contract: _SmartWalletSwapBsc.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscAdminClaimed)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseAdminClaimed(log types.Log) (*SmartWalletSwapBscAdminClaimed, error) {
	event := new(SmartWalletSwapBscAdminClaimed)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscApprovedAllowancesIterator is returned from FilterApprovedAllowances and is used to iterate over the raw logs and unpacked data for ApprovedAllowances events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscApprovedAllowancesIterator struct {
	Event *SmartWalletSwapBscApprovedAllowances // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscApprovedAllowancesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscApprovedAllowances)
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
		it.Event = new(SmartWalletSwapBscApprovedAllowances)
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
func (it *SmartWalletSwapBscApprovedAllowancesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscApprovedAllowancesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscApprovedAllowances represents a ApprovedAllowances event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscApprovedAllowances struct {
	Tokens   []common.Address
	Spenders []common.Address
	IsReset  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovedAllowances is a free log retrieval operation binding the contract event 0xf34c5ed704407ea33d210cd1c76959be869adc80531ee3b3c93229fb606ac16e.
//
// Solidity: event ApprovedAllowances(address[] tokens, address[] spenders, bool isReset)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterApprovedAllowances(opts *bind.FilterOpts) (*SmartWalletSwapBscApprovedAllowancesIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "ApprovedAllowances")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscApprovedAllowancesIterator{contract: _SmartWalletSwapBsc.contract, event: "ApprovedAllowances", logs: logs, sub: sub}, nil
}

// WatchApprovedAllowances is a free log subscription operation binding the contract event 0xf34c5ed704407ea33d210cd1c76959be869adc80531ee3b3c93229fb606ac16e.
//
// Solidity: event ApprovedAllowances(address[] tokens, address[] spenders, bool isReset)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchApprovedAllowances(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscApprovedAllowances) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "ApprovedAllowances")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscApprovedAllowances)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "ApprovedAllowances", log); err != nil {
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
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseApprovedAllowances(log types.Log) (*SmartWalletSwapBscApprovedAllowances, error) {
	event := new(SmartWalletSwapBscApprovedAllowances)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "ApprovedAllowances", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscBnbWithdrawIterator is returned from FilterBnbWithdraw and is used to iterate over the raw logs and unpacked data for BnbWithdraw events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscBnbWithdrawIterator struct {
	Event *SmartWalletSwapBscBnbWithdraw // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscBnbWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscBnbWithdraw)
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
		it.Event = new(SmartWalletSwapBscBnbWithdraw)
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
func (it *SmartWalletSwapBscBnbWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscBnbWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscBnbWithdraw represents a BnbWithdraw event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscBnbWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBnbWithdraw is a free log retrieval operation binding the contract event 0xe4ea4d7e8a83cfbc8cb00c32a56c0dff4d0a22e627fac51ac4a40be6dc9e6309.
//
// Solidity: event BnbWithdraw(uint256 amount, address sendTo)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterBnbWithdraw(opts *bind.FilterOpts) (*SmartWalletSwapBscBnbWithdrawIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "BnbWithdraw")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscBnbWithdrawIterator{contract: _SmartWalletSwapBsc.contract, event: "BnbWithdraw", logs: logs, sub: sub}, nil
}

// WatchBnbWithdraw is a free log subscription operation binding the contract event 0xe4ea4d7e8a83cfbc8cb00c32a56c0dff4d0a22e627fac51ac4a40be6dc9e6309.
//
// Solidity: event BnbWithdraw(uint256 amount, address sendTo)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchBnbWithdraw(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscBnbWithdraw) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "BnbWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscBnbWithdraw)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "BnbWithdraw", log); err != nil {
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

// ParseBnbWithdraw is a log parse operation binding the contract event 0xe4ea4d7e8a83cfbc8cb00c32a56c0dff4d0a22e627fac51ac4a40be6dc9e6309.
//
// Solidity: event BnbWithdraw(uint256 amount, address sendTo)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseBnbWithdraw(log types.Log) (*SmartWalletSwapBscBnbWithdraw, error) {
	event := new(SmartWalletSwapBscBnbWithdraw)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "BnbWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscBorrowFromLendingIterator is returned from FilterBorrowFromLending and is used to iterate over the raw logs and unpacked data for BorrowFromLending events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscBorrowFromLendingIterator struct {
	Event *SmartWalletSwapBscBorrowFromLending // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscBorrowFromLendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscBorrowFromLending)
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
		it.Event = new(SmartWalletSwapBscBorrowFromLending)
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
func (it *SmartWalletSwapBscBorrowFromLendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscBorrowFromLendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscBorrowFromLending represents a BorrowFromLending event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscBorrowFromLending struct {
	Platform         uint8
	Token            common.Address
	AmountBorrowed   *big.Int
	InterestRateMode *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBorrowFromLending is a free log retrieval operation binding the contract event 0x75118cfe403da39422a15ae15992c5b6cc943a7135b081b966882f7ff8e7a652.
//
// Solidity: event BorrowFromLending(uint8 indexed platform, address token, uint256 amountBorrowed, uint256 interestRateMode)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterBorrowFromLending(opts *bind.FilterOpts, platform []uint8) (*SmartWalletSwapBscBorrowFromLendingIterator, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "BorrowFromLending", platformRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscBorrowFromLendingIterator{contract: _SmartWalletSwapBsc.contract, event: "BorrowFromLending", logs: logs, sub: sub}, nil
}

// WatchBorrowFromLending is a free log subscription operation binding the contract event 0x75118cfe403da39422a15ae15992c5b6cc943a7135b081b966882f7ff8e7a652.
//
// Solidity: event BorrowFromLending(uint8 indexed platform, address token, uint256 amountBorrowed, uint256 interestRateMode)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchBorrowFromLending(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscBorrowFromLending, platform []uint8) (event.Subscription, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "BorrowFromLending", platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscBorrowFromLending)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "BorrowFromLending", log); err != nil {
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

// ParseBorrowFromLending is a log parse operation binding the contract event 0x75118cfe403da39422a15ae15992c5b6cc943a7135b081b966882f7ff8e7a652.
//
// Solidity: event BorrowFromLending(uint8 indexed platform, address token, uint256 amountBorrowed, uint256 interestRateMode)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseBorrowFromLending(log types.Log) (*SmartWalletSwapBscBorrowFromLending, error) {
	event := new(SmartWalletSwapBscBorrowFromLending)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "BorrowFromLending", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscClaimedPlatformFeesIterator is returned from FilterClaimedPlatformFees and is used to iterate over the raw logs and unpacked data for ClaimedPlatformFees events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscClaimedPlatformFeesIterator struct {
	Event *SmartWalletSwapBscClaimedPlatformFees // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscClaimedPlatformFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscClaimedPlatformFees)
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
		it.Event = new(SmartWalletSwapBscClaimedPlatformFees)
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
func (it *SmartWalletSwapBscClaimedPlatformFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscClaimedPlatformFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscClaimedPlatformFees represents a ClaimedPlatformFees event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscClaimedPlatformFees struct {
	Wallets []common.Address
	Tokens  []common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedPlatformFees is a free log retrieval operation binding the contract event 0xcbef5ab8ab58a3dba86704ce124241889d80c3a5fe3caaf4286b7e8c709a6c2d.
//
// Solidity: event ClaimedPlatformFees(address[] wallets, address[] tokens, address claimer)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterClaimedPlatformFees(opts *bind.FilterOpts) (*SmartWalletSwapBscClaimedPlatformFeesIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "ClaimedPlatformFees")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscClaimedPlatformFeesIterator{contract: _SmartWalletSwapBsc.contract, event: "ClaimedPlatformFees", logs: logs, sub: sub}, nil
}

// WatchClaimedPlatformFees is a free log subscription operation binding the contract event 0xcbef5ab8ab58a3dba86704ce124241889d80c3a5fe3caaf4286b7e8c709a6c2d.
//
// Solidity: event ClaimedPlatformFees(address[] wallets, address[] tokens, address claimer)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchClaimedPlatformFees(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscClaimedPlatformFees) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "ClaimedPlatformFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscClaimedPlatformFees)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "ClaimedPlatformFees", log); err != nil {
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
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseClaimedPlatformFees(log types.Log) (*SmartWalletSwapBscClaimedPlatformFees, error) {
	event := new(SmartWalletSwapBscClaimedPlatformFees)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "ClaimedPlatformFees", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscKyberTradeIterator is returned from FilterKyberTrade and is used to iterate over the raw logs and unpacked data for KyberTrade events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscKyberTradeIterator struct {
	Event *SmartWalletSwapBscKyberTrade // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscKyberTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscKyberTrade)
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
		it.Event = new(SmartWalletSwapBscKyberTrade)
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
func (it *SmartWalletSwapBscKyberTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscKyberTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscKyberTrade represents a KyberTrade event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscKyberTrade struct {
	Trader         common.Address
	Src            common.Address
	Dest           common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	Recipient      common.Address
	PlatformWallet common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKyberTrade is a free log retrieval operation binding the contract event 0x0e408bba2c83c7ebafdc969a8eea03ffd325c880246a5e8f99a53c9b464dab8d.
//
// Solidity: event KyberTrade(address indexed trader, address indexed src, address indexed dest, uint256 srcAmount, uint256 destAmount, address recipient, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterKyberTrade(opts *bind.FilterOpts, trader []common.Address, src []common.Address, dest []common.Address) (*SmartWalletSwapBscKyberTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "KyberTrade", traderRule, srcRule, destRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscKyberTradeIterator{contract: _SmartWalletSwapBsc.contract, event: "KyberTrade", logs: logs, sub: sub}, nil
}

// WatchKyberTrade is a free log subscription operation binding the contract event 0x0e408bba2c83c7ebafdc969a8eea03ffd325c880246a5e8f99a53c9b464dab8d.
//
// Solidity: event KyberTrade(address indexed trader, address indexed src, address indexed dest, uint256 srcAmount, uint256 destAmount, address recipient, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchKyberTrade(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscKyberTrade, trader []common.Address, src []common.Address, dest []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "KyberTrade", traderRule, srcRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscKyberTrade)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "KyberTrade", log); err != nil {
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

// ParseKyberTrade is a log parse operation binding the contract event 0x0e408bba2c83c7ebafdc969a8eea03ffd325c880246a5e8f99a53c9b464dab8d.
//
// Solidity: event KyberTrade(address indexed trader, address indexed src, address indexed dest, uint256 srcAmount, uint256 destAmount, address recipient, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseKyberTrade(log types.Log) (*SmartWalletSwapBscKyberTrade, error) {
	event := new(SmartWalletSwapBscKyberTrade)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "KyberTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscKyberTradeAndDepositIterator is returned from FilterKyberTradeAndDeposit and is used to iterate over the raw logs and unpacked data for KyberTradeAndDeposit events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscKyberTradeAndDepositIterator struct {
	Event *SmartWalletSwapBscKyberTradeAndDeposit // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscKyberTradeAndDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscKyberTradeAndDeposit)
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
		it.Event = new(SmartWalletSwapBscKyberTradeAndDeposit)
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
func (it *SmartWalletSwapBscKyberTradeAndDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscKyberTradeAndDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscKyberTradeAndDeposit represents a KyberTradeAndDeposit event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscKyberTradeAndDeposit struct {
	Trader         common.Address
	Platform       uint8
	Src            common.Address
	Dest           common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PlatformFeeBps *big.Int
	PlatformWallet common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKyberTradeAndDeposit is a free log retrieval operation binding the contract event 0x8a90604e2d3ef68725ced1d7e0f27ce0920db707428f24d35461285775ff255d.
//
// Solidity: event KyberTradeAndDeposit(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterKyberTradeAndDeposit(opts *bind.FilterOpts, trader []common.Address, platform []uint8, dest []common.Address) (*SmartWalletSwapBscKyberTradeAndDepositIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "KyberTradeAndDeposit", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscKyberTradeAndDepositIterator{contract: _SmartWalletSwapBsc.contract, event: "KyberTradeAndDeposit", logs: logs, sub: sub}, nil
}

// WatchKyberTradeAndDeposit is a free log subscription operation binding the contract event 0x8a90604e2d3ef68725ced1d7e0f27ce0920db707428f24d35461285775ff255d.
//
// Solidity: event KyberTradeAndDeposit(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchKyberTradeAndDeposit(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscKyberTradeAndDeposit, trader []common.Address, platform []uint8, dest []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "KyberTradeAndDeposit", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscKyberTradeAndDeposit)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "KyberTradeAndDeposit", log); err != nil {
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

// ParseKyberTradeAndDeposit is a log parse operation binding the contract event 0x8a90604e2d3ef68725ced1d7e0f27ce0920db707428f24d35461285775ff255d.
//
// Solidity: event KyberTradeAndDeposit(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseKyberTradeAndDeposit(log types.Log) (*SmartWalletSwapBscKyberTradeAndDeposit, error) {
	event := new(SmartWalletSwapBscKyberTradeAndDeposit)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "KyberTradeAndDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscKyberTradeAndRepayIterator is returned from FilterKyberTradeAndRepay and is used to iterate over the raw logs and unpacked data for KyberTradeAndRepay events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscKyberTradeAndRepayIterator struct {
	Event *SmartWalletSwapBscKyberTradeAndRepay // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscKyberTradeAndRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscKyberTradeAndRepay)
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
		it.Event = new(SmartWalletSwapBscKyberTradeAndRepay)
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
func (it *SmartWalletSwapBscKyberTradeAndRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscKyberTradeAndRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscKyberTradeAndRepay represents a KyberTradeAndRepay event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscKyberTradeAndRepay struct {
	Trader         common.Address
	Platform       uint8
	Src            common.Address
	Dest           common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PayAmount      *big.Int
	PlatformWallet common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKyberTradeAndRepay is a free log retrieval operation binding the contract event 0xb8713c79c831b320bd6b22c022aba428799694c76483e0139b41259de383a3de.
//
// Solidity: event KyberTradeAndRepay(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 payAmount, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterKyberTradeAndRepay(opts *bind.FilterOpts, trader []common.Address, platform []uint8, dest []common.Address) (*SmartWalletSwapBscKyberTradeAndRepayIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "KyberTradeAndRepay", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscKyberTradeAndRepayIterator{contract: _SmartWalletSwapBsc.contract, event: "KyberTradeAndRepay", logs: logs, sub: sub}, nil
}

// WatchKyberTradeAndRepay is a free log subscription operation binding the contract event 0xb8713c79c831b320bd6b22c022aba428799694c76483e0139b41259de383a3de.
//
// Solidity: event KyberTradeAndRepay(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 payAmount, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchKyberTradeAndRepay(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscKyberTradeAndRepay, trader []common.Address, platform []uint8, dest []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "KyberTradeAndRepay", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscKyberTradeAndRepay)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "KyberTradeAndRepay", log); err != nil {
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

// ParseKyberTradeAndRepay is a log parse operation binding the contract event 0xb8713c79c831b320bd6b22c022aba428799694c76483e0139b41259de383a3de.
//
// Solidity: event KyberTradeAndRepay(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 payAmount, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseKyberTradeAndRepay(log types.Log) (*SmartWalletSwapBscKyberTradeAndRepay, error) {
	event := new(SmartWalletSwapBscKyberTradeAndRepay)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "KyberTradeAndRepay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscPancakeTradeIterator is returned from FilterPancakeTrade and is used to iterate over the raw logs and unpacked data for PancakeTrade events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscPancakeTradeIterator struct {
	Event *SmartWalletSwapBscPancakeTrade // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscPancakeTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscPancakeTrade)
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
		it.Event = new(SmartWalletSwapBscPancakeTrade)
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
func (it *SmartWalletSwapBscPancakeTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscPancakeTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscPancakeTrade represents a PancakeTrade event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscPancakeTrade struct {
	Trader         common.Address
	Router         common.Address
	TradePath      []common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	Recipient      common.Address
	PlatformFeeBps *big.Int
	PlatformWallet common.Address
	FeeInSrc       bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPancakeTrade is a free log retrieval operation binding the contract event 0x813ae45cd476fd851a56f893cebc2959e2a4ff909b770ee49644623b89860ec5.
//
// Solidity: event PancakeTrade(address indexed trader, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterPancakeTrade(opts *bind.FilterOpts, trader []common.Address, router []common.Address) (*SmartWalletSwapBscPancakeTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "PancakeTrade", traderRule, routerRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscPancakeTradeIterator{contract: _SmartWalletSwapBsc.contract, event: "PancakeTrade", logs: logs, sub: sub}, nil
}

// WatchPancakeTrade is a free log subscription operation binding the contract event 0x813ae45cd476fd851a56f893cebc2959e2a4ff909b770ee49644623b89860ec5.
//
// Solidity: event PancakeTrade(address indexed trader, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchPancakeTrade(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscPancakeTrade, trader []common.Address, router []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "PancakeTrade", traderRule, routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscPancakeTrade)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "PancakeTrade", log); err != nil {
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

// ParsePancakeTrade is a log parse operation binding the contract event 0x813ae45cd476fd851a56f893cebc2959e2a4ff909b770ee49644623b89860ec5.
//
// Solidity: event PancakeTrade(address indexed trader, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParsePancakeTrade(log types.Log) (*SmartWalletSwapBscPancakeTrade, error) {
	event := new(SmartWalletSwapBscPancakeTrade)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "PancakeTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscPancakeTradeAndDepositIterator is returned from FilterPancakeTradeAndDeposit and is used to iterate over the raw logs and unpacked data for PancakeTradeAndDeposit events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscPancakeTradeAndDepositIterator struct {
	Event *SmartWalletSwapBscPancakeTradeAndDeposit // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscPancakeTradeAndDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscPancakeTradeAndDeposit)
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
		it.Event = new(SmartWalletSwapBscPancakeTradeAndDeposit)
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
func (it *SmartWalletSwapBscPancakeTradeAndDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscPancakeTradeAndDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscPancakeTradeAndDeposit represents a PancakeTradeAndDeposit event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscPancakeTradeAndDeposit struct {
	Trader         common.Address
	Platform       uint8
	Router         common.Address
	TradePath      []common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PlatformFeeBps *big.Int
	PlatformWallet common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPancakeTradeAndDeposit is a free log retrieval operation binding the contract event 0x07c268111f1fea559b821a26c698f93ced4cd56bb8b44d71ed75f74c4c94dd7c.
//
// Solidity: event PancakeTradeAndDeposit(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterPancakeTradeAndDeposit(opts *bind.FilterOpts, trader []common.Address, platform []uint8, router []common.Address) (*SmartWalletSwapBscPancakeTradeAndDepositIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "PancakeTradeAndDeposit", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscPancakeTradeAndDepositIterator{contract: _SmartWalletSwapBsc.contract, event: "PancakeTradeAndDeposit", logs: logs, sub: sub}, nil
}

// WatchPancakeTradeAndDeposit is a free log subscription operation binding the contract event 0x07c268111f1fea559b821a26c698f93ced4cd56bb8b44d71ed75f74c4c94dd7c.
//
// Solidity: event PancakeTradeAndDeposit(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchPancakeTradeAndDeposit(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscPancakeTradeAndDeposit, trader []common.Address, platform []uint8, router []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "PancakeTradeAndDeposit", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscPancakeTradeAndDeposit)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "PancakeTradeAndDeposit", log); err != nil {
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

// ParsePancakeTradeAndDeposit is a log parse operation binding the contract event 0x07c268111f1fea559b821a26c698f93ced4cd56bb8b44d71ed75f74c4c94dd7c.
//
// Solidity: event PancakeTradeAndDeposit(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParsePancakeTradeAndDeposit(log types.Log) (*SmartWalletSwapBscPancakeTradeAndDeposit, error) {
	event := new(SmartWalletSwapBscPancakeTradeAndDeposit)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "PancakeTradeAndDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscPancakeTradeAndRepayIterator is returned from FilterPancakeTradeAndRepay and is used to iterate over the raw logs and unpacked data for PancakeTradeAndRepay events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscPancakeTradeAndRepayIterator struct {
	Event *SmartWalletSwapBscPancakeTradeAndRepay // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscPancakeTradeAndRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscPancakeTradeAndRepay)
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
		it.Event = new(SmartWalletSwapBscPancakeTradeAndRepay)
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
func (it *SmartWalletSwapBscPancakeTradeAndRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscPancakeTradeAndRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscPancakeTradeAndRepay represents a PancakeTradeAndRepay event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscPancakeTradeAndRepay struct {
	Trader         common.Address
	Platform       uint8
	Router         common.Address
	TradePath      []common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PayAmount      *big.Int
	FeeAndRateMode *big.Int
	PlatformWallet common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPancakeTradeAndRepay is a free log retrieval operation binding the contract event 0xf02f9bc3e1a33c10939926a50e33ce83ac80de1b0822fdf326d0c20b1243656a.
//
// Solidity: event PancakeTradeAndRepay(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterPancakeTradeAndRepay(opts *bind.FilterOpts, trader []common.Address, platform []uint8, router []common.Address) (*SmartWalletSwapBscPancakeTradeAndRepayIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "PancakeTradeAndRepay", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscPancakeTradeAndRepayIterator{contract: _SmartWalletSwapBsc.contract, event: "PancakeTradeAndRepay", logs: logs, sub: sub}, nil
}

// WatchPancakeTradeAndRepay is a free log subscription operation binding the contract event 0xf02f9bc3e1a33c10939926a50e33ce83ac80de1b0822fdf326d0c20b1243656a.
//
// Solidity: event PancakeTradeAndRepay(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchPancakeTradeAndRepay(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscPancakeTradeAndRepay, trader []common.Address, platform []uint8, router []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "PancakeTradeAndRepay", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscPancakeTradeAndRepay)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "PancakeTradeAndRepay", log); err != nil {
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

// ParsePancakeTradeAndRepay is a log parse operation binding the contract event 0xf02f9bc3e1a33c10939926a50e33ce83ac80de1b0822fdf326d0c20b1243656a.
//
// Solidity: event PancakeTradeAndRepay(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParsePancakeTradeAndRepay(log types.Log) (*SmartWalletSwapBscPancakeTradeAndRepay, error) {
	event := new(SmartWalletSwapBscPancakeTradeAndRepay)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "PancakeTradeAndRepay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscTokenWithdrawIterator struct {
	Event *SmartWalletSwapBscTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscTokenWithdraw)
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
		it.Event = new(SmartWalletSwapBscTokenWithdraw)
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
func (it *SmartWalletSwapBscTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscTokenWithdraw represents a TokenWithdraw event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*SmartWalletSwapBscTokenWithdrawIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscTokenWithdrawIterator{contract: _SmartWalletSwapBsc.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscTokenWithdraw)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseTokenWithdraw(log types.Log) (*SmartWalletSwapBscTokenWithdraw, error) {
	event := new(SmartWalletSwapBscTokenWithdraw)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscTransferAdminPendingIterator struct {
	Event *SmartWalletSwapBscTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscTransferAdminPending)
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
		it.Event = new(SmartWalletSwapBscTransferAdminPending)
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
func (it *SmartWalletSwapBscTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscTransferAdminPending represents a TransferAdminPending event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*SmartWalletSwapBscTransferAdminPendingIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscTransferAdminPendingIterator{contract: _SmartWalletSwapBsc.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscTransferAdminPending)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseTransferAdminPending(log types.Log) (*SmartWalletSwapBscTransferAdminPending, error) {
	event := new(SmartWalletSwapBscTransferAdminPending)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscUpdatedLendingImplementationIterator is returned from FilterUpdatedLendingImplementation and is used to iterate over the raw logs and unpacked data for UpdatedLendingImplementation events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscUpdatedLendingImplementationIterator struct {
	Event *SmartWalletSwapBscUpdatedLendingImplementation // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscUpdatedLendingImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscUpdatedLendingImplementation)
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
		it.Event = new(SmartWalletSwapBscUpdatedLendingImplementation)
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
func (it *SmartWalletSwapBscUpdatedLendingImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscUpdatedLendingImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscUpdatedLendingImplementation represents a UpdatedLendingImplementation event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscUpdatedLendingImplementation struct {
	Impl common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdatedLendingImplementation is a free log retrieval operation binding the contract event 0x7d914f4a47f2d06bf2a4f2679866f53f24fd19f5fe6923bbe6d8e54e3cd41536.
//
// Solidity: event UpdatedLendingImplementation(address impl)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterUpdatedLendingImplementation(opts *bind.FilterOpts) (*SmartWalletSwapBscUpdatedLendingImplementationIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "UpdatedLendingImplementation")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscUpdatedLendingImplementationIterator{contract: _SmartWalletSwapBsc.contract, event: "UpdatedLendingImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdatedLendingImplementation is a free log subscription operation binding the contract event 0x7d914f4a47f2d06bf2a4f2679866f53f24fd19f5fe6923bbe6d8e54e3cd41536.
//
// Solidity: event UpdatedLendingImplementation(address impl)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchUpdatedLendingImplementation(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscUpdatedLendingImplementation) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "UpdatedLendingImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscUpdatedLendingImplementation)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "UpdatedLendingImplementation", log); err != nil {
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

// ParseUpdatedLendingImplementation is a log parse operation binding the contract event 0x7d914f4a47f2d06bf2a4f2679866f53f24fd19f5fe6923bbe6d8e54e3cd41536.
//
// Solidity: event UpdatedLendingImplementation(address impl)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseUpdatedLendingImplementation(log types.Log) (*SmartWalletSwapBscUpdatedLendingImplementation, error) {
	event := new(SmartWalletSwapBscUpdatedLendingImplementation)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "UpdatedLendingImplementation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscUpdatedSupportedPlatformWalletsIterator is returned from FilterUpdatedSupportedPlatformWallets and is used to iterate over the raw logs and unpacked data for UpdatedSupportedPlatformWallets events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscUpdatedSupportedPlatformWalletsIterator struct {
	Event *SmartWalletSwapBscUpdatedSupportedPlatformWallets // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscUpdatedSupportedPlatformWalletsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscUpdatedSupportedPlatformWallets)
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
		it.Event = new(SmartWalletSwapBscUpdatedSupportedPlatformWallets)
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
func (it *SmartWalletSwapBscUpdatedSupportedPlatformWalletsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscUpdatedSupportedPlatformWalletsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscUpdatedSupportedPlatformWallets represents a UpdatedSupportedPlatformWallets event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscUpdatedSupportedPlatformWallets struct {
	Wallets     []common.Address
	IsSupported bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSupportedPlatformWallets is a free log retrieval operation binding the contract event 0xac650f539fc5264cdad2074ff608d20b54013400070e4e5ef39125f67d8c986c.
//
// Solidity: event UpdatedSupportedPlatformWallets(address[] wallets, bool isSupported)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterUpdatedSupportedPlatformWallets(opts *bind.FilterOpts) (*SmartWalletSwapBscUpdatedSupportedPlatformWalletsIterator, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "UpdatedSupportedPlatformWallets")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscUpdatedSupportedPlatformWalletsIterator{contract: _SmartWalletSwapBsc.contract, event: "UpdatedSupportedPlatformWallets", logs: logs, sub: sub}, nil
}

// WatchUpdatedSupportedPlatformWallets is a free log subscription operation binding the contract event 0xac650f539fc5264cdad2074ff608d20b54013400070e4e5ef39125f67d8c986c.
//
// Solidity: event UpdatedSupportedPlatformWallets(address[] wallets, bool isSupported)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchUpdatedSupportedPlatformWallets(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscUpdatedSupportedPlatformWallets) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "UpdatedSupportedPlatformWallets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscUpdatedSupportedPlatformWallets)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "UpdatedSupportedPlatformWallets", log); err != nil {
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

// ParseUpdatedSupportedPlatformWallets is a log parse operation binding the contract event 0xac650f539fc5264cdad2074ff608d20b54013400070e4e5ef39125f67d8c986c.
//
// Solidity: event UpdatedSupportedPlatformWallets(address[] wallets, bool isSupported)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseUpdatedSupportedPlatformWallets(log types.Log) (*SmartWalletSwapBscUpdatedSupportedPlatformWallets, error) {
	event := new(SmartWalletSwapBscUpdatedSupportedPlatformWallets)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "UpdatedSupportedPlatformWallets", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapBscWithdrawFromLendingIterator is returned from FilterWithdrawFromLending and is used to iterate over the raw logs and unpacked data for WithdrawFromLending events raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscWithdrawFromLendingIterator struct {
	Event *SmartWalletSwapBscWithdrawFromLending // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapBscWithdrawFromLendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapBscWithdrawFromLending)
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
		it.Event = new(SmartWalletSwapBscWithdrawFromLending)
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
func (it *SmartWalletSwapBscWithdrawFromLendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapBscWithdrawFromLendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapBscWithdrawFromLending represents a WithdrawFromLending event raised by the SmartWalletSwapBsc contract.
type SmartWalletSwapBscWithdrawFromLending struct {
	Platform           uint8
	Token              common.Address
	Amount             *big.Int
	MinReturn          *big.Int
	ActualReturnAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFromLending is a free log retrieval operation binding the contract event 0x9530d3a25f06d6f06fee69e2478e3f1e4dd81afdd77489dd45a7a49b8ad78c74.
//
// Solidity: event WithdrawFromLending(uint8 indexed platform, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) FilterWithdrawFromLending(opts *bind.FilterOpts, platform []uint8) (*SmartWalletSwapBscWithdrawFromLendingIterator, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.FilterLogs(opts, "WithdrawFromLending", platformRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapBscWithdrawFromLendingIterator{contract: _SmartWalletSwapBsc.contract, event: "WithdrawFromLending", logs: logs, sub: sub}, nil
}

// WatchWithdrawFromLending is a free log subscription operation binding the contract event 0x9530d3a25f06d6f06fee69e2478e3f1e4dd81afdd77489dd45a7a49b8ad78c74.
//
// Solidity: event WithdrawFromLending(uint8 indexed platform, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) WatchWithdrawFromLending(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapBscWithdrawFromLending, platform []uint8) (event.Subscription, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _SmartWalletSwapBsc.contract.WatchLogs(opts, "WithdrawFromLending", platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapBscWithdrawFromLending)
				if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "WithdrawFromLending", log); err != nil {
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

// ParseWithdrawFromLending is a log parse operation binding the contract event 0x9530d3a25f06d6f06fee69e2478e3f1e4dd81afdd77489dd45a7a49b8ad78c74.
//
// Solidity: event WithdrawFromLending(uint8 indexed platform, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount)
func (_SmartWalletSwapBsc *SmartWalletSwapBscFilterer) ParseWithdrawFromLending(log types.Log) (*SmartWalletSwapBscWithdrawFromLending, error) {
	event := new(SmartWalletSwapBscWithdrawFromLending)
	if err := _SmartWalletSwapBsc.contract.UnpackLog(event, "WithdrawFromLending", log); err != nil {
		return nil, err
	}
	return event, nil
}
