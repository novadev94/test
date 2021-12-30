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

// SmartWalletSwapABI is the input ABI used to generate the binding from.
const SmartWalletSwapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"spenders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isReset\",\"type\":\"bool\"}],\"name\":\"ApprovedAllowances\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimedPlatformFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Ext\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Ext\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numGasBurns\",\"type\":\"uint256\"}],\"name\":\"KyberTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIERC20Ext\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Ext\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numGasBurns\",\"type\":\"uint256\"}],\"name\":\"KyberTradeAndDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIERC20Ext\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Ext\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAndRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numGasBurns\",\"type\":\"uint256\"}],\"name\":\"KyberTradeAndRepay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"feeInSrc\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numGasBurns\",\"type\":\"uint256\"}],\"name\":\"UniswapTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numGasBurns\",\"type\":\"uint256\"}],\"name\":\"UniswapTradeAndDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAndRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numGasBurns\",\"type\":\"uint256\"}],\"name\":\"UniswapTradeAndRepay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIKyberProxy\",\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"UpdateKyberProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIUniswapV2Router02[]\",\"name\":\"uniswapRouters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdded\",\"type\":\"bool\"}],\"name\":\"UpdateUniswapRouters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBurnGasHelper\",\"name\":\"gasHelper\",\"type\":\"address\"}],\"name\":\"UpdatedBurnGasHelper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractISmartWalletLending\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"UpdatedLendingImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"UpdatedSupportedPlatformWallets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualReturnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numGasBurns\",\"type\":\"uint256\"}],\"name\":\"WithdrawFromLending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"spenders\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isReset\",\"type\":\"bool\"}],\"name\":\"approveAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnGasHelper\",\"outputs\":[{\"internalType\":\"contractIBurnGasHelper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractICompErc20[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"claimComp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"plaftformWallets\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20Ext[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"claimPlatformFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"getExpectedReturnKyber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturnUniswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRouterSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberProxy\",\"outputs\":[{\"internalType\":\"contractIKyberProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingImpl\",\"outputs\":[{\"internalType\":\"contractISmartWalletLending\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"platformWalletFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedPlatformWallets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"swapKyber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"swapKyberAndDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAndRateMode\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"swapKyberAndRepay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"feeInSrc\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"swapUniswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"swapUniswapAndDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tradePath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"feeAndRateMode\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"swapUniswapAndRepay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBurnGasHelper\",\"name\":\"_burnGasHelper\",\"type\":\"address\"}],\"name\":\"updateBurnGasHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIKyberProxy\",\"name\":\"_kyberProxy\",\"type\":\"address\"}],\"name\":\"updateKyberProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISmartWalletLending\",\"name\":\"newImpl\",\"type\":\"address\"}],\"name\":\"updateLendingImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"updateSupportedPlatformWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Router02[]\",\"name\":\"_uniswapRouters\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isAdded\",\"type\":\"bool\"}],\"name\":\"updateUniswapRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISmartWalletLending.LendingPlatform\",\"name\":\"platform\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useGasToken\",\"type\":\"bool\"}],\"name\":\"withdrawFromLendingPlatform\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Ext\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SmartWalletSwap is an auto generated Go binding around an Ethereum contract.
type SmartWalletSwap struct {
	SmartWalletSwapCaller     // Read-only binding to the contract
	SmartWalletSwapTransactor // Write-only binding to the contract
	SmartWalletSwapFilterer   // Log filterer for contract events
}

// SmartWalletSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartWalletSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartWalletSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartWalletSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartWalletSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartWalletSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartWalletSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartWalletSwapSession struct {
	Contract     *SmartWalletSwap  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartWalletSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartWalletSwapCallerSession struct {
	Contract *SmartWalletSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SmartWalletSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartWalletSwapTransactorSession struct {
	Contract     *SmartWalletSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SmartWalletSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartWalletSwapRaw struct {
	Contract *SmartWalletSwap // Generic contract binding to access the raw methods on
}

// SmartWalletSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartWalletSwapCallerRaw struct {
	Contract *SmartWalletSwapCaller // Generic read-only contract binding to access the raw methods on
}

// SmartWalletSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartWalletSwapTransactorRaw struct {
	Contract *SmartWalletSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartWalletSwap creates a new instance of SmartWalletSwap, bound to a specific deployed contract.
func NewSmartWalletSwap(address common.Address, backend bind.ContractBackend) (*SmartWalletSwap, error) {
	contract, err := bindSmartWalletSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwap{SmartWalletSwapCaller: SmartWalletSwapCaller{contract: contract}, SmartWalletSwapTransactor: SmartWalletSwapTransactor{contract: contract}, SmartWalletSwapFilterer: SmartWalletSwapFilterer{contract: contract}}, nil
}

// NewSmartWalletSwapCaller creates a new read-only instance of SmartWalletSwap, bound to a specific deployed contract.
func NewSmartWalletSwapCaller(address common.Address, caller bind.ContractCaller) (*SmartWalletSwapCaller, error) {
	contract, err := bindSmartWalletSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapCaller{contract: contract}, nil
}

// NewSmartWalletSwapTransactor creates a new write-only instance of SmartWalletSwap, bound to a specific deployed contract.
func NewSmartWalletSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartWalletSwapTransactor, error) {
	contract, err := bindSmartWalletSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapTransactor{contract: contract}, nil
}

// NewSmartWalletSwapFilterer creates a new log filterer instance of SmartWalletSwap, bound to a specific deployed contract.
func NewSmartWalletSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartWalletSwapFilterer, error) {
	contract, err := bindSmartWalletSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapFilterer{contract: contract}, nil
}

// bindSmartWalletSwap binds a generic wrapper to an already deployed contract.
func bindSmartWalletSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartWalletSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartWalletSwap *SmartWalletSwapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartWalletSwap.Contract.SmartWalletSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartWalletSwap *SmartWalletSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SmartWalletSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartWalletSwap *SmartWalletSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SmartWalletSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartWalletSwap *SmartWalletSwapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartWalletSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartWalletSwap *SmartWalletSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartWalletSwap *SmartWalletSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.contract.Transact(opts, method, params...)
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_SmartWalletSwap *SmartWalletSwapCaller) BPS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "BPS")
	return *ret0, err
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_SmartWalletSwap *SmartWalletSwapSession) BPS() (*big.Int, error) {
	return _SmartWalletSwap.Contract.BPS(&_SmartWalletSwap.CallOpts)
}

// BPS is a free data retrieval call binding the contract method 0x249d39e9.
//
// Solidity: function BPS() view returns(uint256)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) BPS() (*big.Int, error) {
	return _SmartWalletSwap.Contract.BPS(&_SmartWalletSwap.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapSession) Admin() (common.Address, error) {
	return _SmartWalletSwap.Contract.Admin(&_SmartWalletSwap.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) Admin() (common.Address, error) {
	return _SmartWalletSwap.Contract.Admin(&_SmartWalletSwap.CallOpts)
}

// BurnGasHelper is a free data retrieval call binding the contract method 0x33b057ea.
//
// Solidity: function burnGasHelper() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCaller) BurnGasHelper(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "burnGasHelper")
	return *ret0, err
}

// BurnGasHelper is a free data retrieval call binding the contract method 0x33b057ea.
//
// Solidity: function burnGasHelper() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapSession) BurnGasHelper() (common.Address, error) {
	return _SmartWalletSwap.Contract.BurnGasHelper(&_SmartWalletSwap.CallOpts)
}

// BurnGasHelper is a free data retrieval call binding the contract method 0x33b057ea.
//
// Solidity: function burnGasHelper() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) BurnGasHelper() (common.Address, error) {
	return _SmartWalletSwap.Contract.BurnGasHelper(&_SmartWalletSwap.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_SmartWalletSwap *SmartWalletSwapCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_SmartWalletSwap *SmartWalletSwapSession) GetAlerters() ([]common.Address, error) {
	return _SmartWalletSwap.Contract.GetAlerters(&_SmartWalletSwap.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_SmartWalletSwap *SmartWalletSwapCallerSession) GetAlerters() ([]common.Address, error) {
	return _SmartWalletSwap.Contract.GetAlerters(&_SmartWalletSwap.CallOpts)
}

// GetExpectedReturnKyber is a free data retrieval call binding the contract method 0xf2459d63.
//
// Solidity: function getExpectedReturnKyber(address src, address dest, uint256 srcAmount, uint256 platformFee, bytes hint) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwap *SmartWalletSwapCaller) GetExpectedReturnKyber(opts *bind.CallOpts, src common.Address, dest common.Address, srcAmount *big.Int, platformFee *big.Int, hint []byte) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	ret := new(struct {
		DestAmount   *big.Int
		ExpectedRate *big.Int
	})
	out := ret
	err := _SmartWalletSwap.contract.Call(opts, out, "getExpectedReturnKyber", src, dest, srcAmount, platformFee, hint)
	return *ret, err
}

// GetExpectedReturnKyber is a free data retrieval call binding the contract method 0xf2459d63.
//
// Solidity: function getExpectedReturnKyber(address src, address dest, uint256 srcAmount, uint256 platformFee, bytes hint) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwap *SmartWalletSwapSession) GetExpectedReturnKyber(src common.Address, dest common.Address, srcAmount *big.Int, platformFee *big.Int, hint []byte) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwap.Contract.GetExpectedReturnKyber(&_SmartWalletSwap.CallOpts, src, dest, srcAmount, platformFee, hint)
}

// GetExpectedReturnKyber is a free data retrieval call binding the contract method 0xf2459d63.
//
// Solidity: function getExpectedReturnKyber(address src, address dest, uint256 srcAmount, uint256 platformFee, bytes hint) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) GetExpectedReturnKyber(src common.Address, dest common.Address, srcAmount *big.Int, platformFee *big.Int, hint []byte) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwap.Contract.GetExpectedReturnKyber(&_SmartWalletSwap.CallOpts, src, dest, srcAmount, platformFee, hint)
}

// GetExpectedReturnUniswap is a free data retrieval call binding the contract method 0x42c956bc.
//
// Solidity: function getExpectedReturnUniswap(address router, uint256 srcAmount, address[] tradePath, uint256 platformFee) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwap *SmartWalletSwapCaller) GetExpectedReturnUniswap(opts *bind.CallOpts, router common.Address, srcAmount *big.Int, tradePath []common.Address, platformFee *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	ret := new(struct {
		DestAmount   *big.Int
		ExpectedRate *big.Int
	})
	out := ret
	err := _SmartWalletSwap.contract.Call(opts, out, "getExpectedReturnUniswap", router, srcAmount, tradePath, platformFee)
	return *ret, err
}

// GetExpectedReturnUniswap is a free data retrieval call binding the contract method 0x42c956bc.
//
// Solidity: function getExpectedReturnUniswap(address router, uint256 srcAmount, address[] tradePath, uint256 platformFee) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwap *SmartWalletSwapSession) GetExpectedReturnUniswap(router common.Address, srcAmount *big.Int, tradePath []common.Address, platformFee *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwap.Contract.GetExpectedReturnUniswap(&_SmartWalletSwap.CallOpts, router, srcAmount, tradePath, platformFee)
}

// GetExpectedReturnUniswap is a free data retrieval call binding the contract method 0x42c956bc.
//
// Solidity: function getExpectedReturnUniswap(address router, uint256 srcAmount, address[] tradePath, uint256 platformFee) view returns(uint256 destAmount, uint256 expectedRate)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) GetExpectedReturnUniswap(router common.Address, srcAmount *big.Int, tradePath []common.Address, platformFee *big.Int) (struct {
	DestAmount   *big.Int
	ExpectedRate *big.Int
}, error) {
	return _SmartWalletSwap.Contract.GetExpectedReturnUniswap(&_SmartWalletSwap.CallOpts, router, srcAmount, tradePath, platformFee)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_SmartWalletSwap *SmartWalletSwapCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_SmartWalletSwap *SmartWalletSwapSession) GetOperators() ([]common.Address, error) {
	return _SmartWalletSwap.Contract.GetOperators(&_SmartWalletSwap.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_SmartWalletSwap *SmartWalletSwapCallerSession) GetOperators() ([]common.Address, error) {
	return _SmartWalletSwap.Contract.GetOperators(&_SmartWalletSwap.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "implementation")
	return *ret0, err
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapSession) Implementation() (common.Address, error) {
	return _SmartWalletSwap.Contract.Implementation(&_SmartWalletSwap.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) Implementation() (common.Address, error) {
	return _SmartWalletSwap.Contract.Implementation(&_SmartWalletSwap.CallOpts)
}

// IsRouterSupported is a free data retrieval call binding the contract method 0xc277b57f.
//
// Solidity: function isRouterSupported(address ) view returns(bool)
func (_SmartWalletSwap *SmartWalletSwapCaller) IsRouterSupported(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "isRouterSupported", arg0)
	return *ret0, err
}

// IsRouterSupported is a free data retrieval call binding the contract method 0xc277b57f.
//
// Solidity: function isRouterSupported(address ) view returns(bool)
func (_SmartWalletSwap *SmartWalletSwapSession) IsRouterSupported(arg0 common.Address) (bool, error) {
	return _SmartWalletSwap.Contract.IsRouterSupported(&_SmartWalletSwap.CallOpts, arg0)
}

// IsRouterSupported is a free data retrieval call binding the contract method 0xc277b57f.
//
// Solidity: function isRouterSupported(address ) view returns(bool)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) IsRouterSupported(arg0 common.Address) (bool, error) {
	return _SmartWalletSwap.Contract.IsRouterSupported(&_SmartWalletSwap.CallOpts, arg0)
}

// KyberProxy is a free data retrieval call binding the contract method 0xf0eeed81.
//
// Solidity: function kyberProxy() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCaller) KyberProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "kyberProxy")
	return *ret0, err
}

// KyberProxy is a free data retrieval call binding the contract method 0xf0eeed81.
//
// Solidity: function kyberProxy() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapSession) KyberProxy() (common.Address, error) {
	return _SmartWalletSwap.Contract.KyberProxy(&_SmartWalletSwap.CallOpts)
}

// KyberProxy is a free data retrieval call binding the contract method 0xf0eeed81.
//
// Solidity: function kyberProxy() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) KyberProxy() (common.Address, error) {
	return _SmartWalletSwap.Contract.KyberProxy(&_SmartWalletSwap.CallOpts)
}

// LendingImpl is a free data retrieval call binding the contract method 0x8690948b.
//
// Solidity: function lendingImpl() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCaller) LendingImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "lendingImpl")
	return *ret0, err
}

// LendingImpl is a free data retrieval call binding the contract method 0x8690948b.
//
// Solidity: function lendingImpl() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapSession) LendingImpl() (common.Address, error) {
	return _SmartWalletSwap.Contract.LendingImpl(&_SmartWalletSwap.CallOpts)
}

// LendingImpl is a free data retrieval call binding the contract method 0x8690948b.
//
// Solidity: function lendingImpl() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) LendingImpl() (common.Address, error) {
	return _SmartWalletSwap.Contract.LendingImpl(&_SmartWalletSwap.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapSession) PendingAdmin() (common.Address, error) {
	return _SmartWalletSwap.Contract.PendingAdmin(&_SmartWalletSwap.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) PendingAdmin() (common.Address, error) {
	return _SmartWalletSwap.Contract.PendingAdmin(&_SmartWalletSwap.CallOpts)
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_SmartWalletSwap *SmartWalletSwapCaller) PlatformWalletFees(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "platformWalletFees", arg0, arg1)
	return *ret0, err
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_SmartWalletSwap *SmartWalletSwapSession) PlatformWalletFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SmartWalletSwap.Contract.PlatformWalletFees(&_SmartWalletSwap.CallOpts, arg0, arg1)
}

// PlatformWalletFees is a free data retrieval call binding the contract method 0xa7c47797.
//
// Solidity: function platformWalletFees(address , address ) view returns(uint256)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) PlatformWalletFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SmartWalletSwap.Contract.PlatformWalletFees(&_SmartWalletSwap.CallOpts, arg0, arg1)
}

// SupportedPlatformWallets is a free data retrieval call binding the contract method 0xdc6f6428.
//
// Solidity: function supportedPlatformWallets(address ) view returns(bool)
func (_SmartWalletSwap *SmartWalletSwapCaller) SupportedPlatformWallets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartWalletSwap.contract.Call(opts, out, "supportedPlatformWallets", arg0)
	return *ret0, err
}

// SupportedPlatformWallets is a free data retrieval call binding the contract method 0xdc6f6428.
//
// Solidity: function supportedPlatformWallets(address ) view returns(bool)
func (_SmartWalletSwap *SmartWalletSwapSession) SupportedPlatformWallets(arg0 common.Address) (bool, error) {
	return _SmartWalletSwap.Contract.SupportedPlatformWallets(&_SmartWalletSwap.CallOpts, arg0)
}

// SupportedPlatformWallets is a free data retrieval call binding the contract method 0xdc6f6428.
//
// Solidity: function supportedPlatformWallets(address ) view returns(bool)
func (_SmartWalletSwap *SmartWalletSwapCallerSession) SupportedPlatformWallets(arg0 common.Address) (bool, error) {
	return _SmartWalletSwap.Contract.SupportedPlatformWallets(&_SmartWalletSwap.CallOpts, arg0)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.AddAlerter(&_SmartWalletSwap.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.AddAlerter(&_SmartWalletSwap.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.AddOperator(&_SmartWalletSwap.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.AddOperator(&_SmartWalletSwap.TransactOpts, newOperator)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) ApproveAllowances(opts *bind.TransactOpts, tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "approveAllowances", tokens, spenders, isReset)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) ApproveAllowances(tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ApproveAllowances(&_SmartWalletSwap.TransactOpts, tokens, spenders, isReset)
}

// ApproveAllowances is a paid mutator transaction binding the contract method 0xbc655543.
//
// Solidity: function approveAllowances(address[] tokens, address[] spenders, bool isReset) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) ApproveAllowances(tokens []common.Address, spenders []common.Address, isReset bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ApproveAllowances(&_SmartWalletSwap.TransactOpts, tokens, spenders, isReset)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SmartWalletSwap *SmartWalletSwapSession) ClaimAdmin() (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ClaimAdmin(&_SmartWalletSwap.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ClaimAdmin(&_SmartWalletSwap.TransactOpts)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x7c3fdc67.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers, bool useGasToken) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) ClaimComp(opts *bind.TransactOpts, holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "claimComp", holders, cTokens, borrowers, suppliers, useGasToken)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x7c3fdc67.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers, bool useGasToken) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) ClaimComp(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ClaimComp(&_SmartWalletSwap.TransactOpts, holders, cTokens, borrowers, suppliers, useGasToken)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x7c3fdc67.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers, bool useGasToken) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) ClaimComp(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ClaimComp(&_SmartWalletSwap.TransactOpts, holders, cTokens, borrowers, suppliers, useGasToken)
}

// ClaimPlatformFees is a paid mutator transaction binding the contract method 0x35c7c3cf.
//
// Solidity: function claimPlatformFees(address[] plaftformWallets, address[] tokens) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) ClaimPlatformFees(opts *bind.TransactOpts, plaftformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "claimPlatformFees", plaftformWallets, tokens)
}

// ClaimPlatformFees is a paid mutator transaction binding the contract method 0x35c7c3cf.
//
// Solidity: function claimPlatformFees(address[] plaftformWallets, address[] tokens) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) ClaimPlatformFees(plaftformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ClaimPlatformFees(&_SmartWalletSwap.TransactOpts, plaftformWallets, tokens)
}

// ClaimPlatformFees is a paid mutator transaction binding the contract method 0x35c7c3cf.
//
// Solidity: function claimPlatformFees(address[] plaftformWallets, address[] tokens) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) ClaimPlatformFees(plaftformWallets []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.ClaimPlatformFees(&_SmartWalletSwap.TransactOpts, plaftformWallets, tokens)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.RemoveAlerter(&_SmartWalletSwap.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.RemoveAlerter(&_SmartWalletSwap.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.RemoveOperator(&_SmartWalletSwap.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.RemoveOperator(&_SmartWalletSwap.TransactOpts, operator)
}

// SwapKyber is a paid mutator transaction binding the contract method 0xcf512b53.
//
// Solidity: function swapKyber(address src, address dest, uint256 srcAmount, uint256 minConversionRate, address recipient, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactor) SwapKyber(opts *bind.TransactOpts, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "swapKyber", src, dest, srcAmount, minConversionRate, recipient, platformFeeBps, platformWallet, hint, useGasToken)
}

// SwapKyber is a paid mutator transaction binding the contract method 0xcf512b53.
//
// Solidity: function swapKyber(address src, address dest, uint256 srcAmount, uint256 minConversionRate, address recipient, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapSession) SwapKyber(src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapKyber(&_SmartWalletSwap.TransactOpts, src, dest, srcAmount, minConversionRate, recipient, platformFeeBps, platformWallet, hint, useGasToken)
}

// SwapKyber is a paid mutator transaction binding the contract method 0xcf512b53.
//
// Solidity: function swapKyber(address src, address dest, uint256 srcAmount, uint256 minConversionRate, address recipient, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) SwapKyber(src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapKyber(&_SmartWalletSwap.TransactOpts, src, dest, srcAmount, minConversionRate, recipient, platformFeeBps, platformWallet, hint, useGasToken)
}

// SwapKyberAndDeposit is a paid mutator transaction binding the contract method 0x30037de5.
//
// Solidity: function swapKyberAndDeposit(uint8 platform, address src, address dest, uint256 srcAmount, uint256 minConversionRate, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactor) SwapKyberAndDeposit(opts *bind.TransactOpts, platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, platformFeeBps *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "swapKyberAndDeposit", platform, src, dest, srcAmount, minConversionRate, platformFeeBps, platformWallet, hint, useGasToken)
}

// SwapKyberAndDeposit is a paid mutator transaction binding the contract method 0x30037de5.
//
// Solidity: function swapKyberAndDeposit(uint8 platform, address src, address dest, uint256 srcAmount, uint256 minConversionRate, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapSession) SwapKyberAndDeposit(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, platformFeeBps *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapKyberAndDeposit(&_SmartWalletSwap.TransactOpts, platform, src, dest, srcAmount, minConversionRate, platformFeeBps, platformWallet, hint, useGasToken)
}

// SwapKyberAndDeposit is a paid mutator transaction binding the contract method 0x30037de5.
//
// Solidity: function swapKyberAndDeposit(uint8 platform, address src, address dest, uint256 srcAmount, uint256 minConversionRate, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) SwapKyberAndDeposit(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, minConversionRate *big.Int, platformFeeBps *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapKyberAndDeposit(&_SmartWalletSwap.TransactOpts, platform, src, dest, srcAmount, minConversionRate, platformFeeBps, platformWallet, hint, useGasToken)
}

// SwapKyberAndRepay is a paid mutator transaction binding the contract method 0x4f4bc120.
//
// Solidity: function swapKyberAndRepay(uint8 platform, address src, address dest, uint256 srcAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactor) SwapKyberAndRepay(opts *bind.TransactOpts, platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, payAmount *big.Int, feeAndRateMode *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "swapKyberAndRepay", platform, src, dest, srcAmount, payAmount, feeAndRateMode, platformWallet, hint, useGasToken)
}

// SwapKyberAndRepay is a paid mutator transaction binding the contract method 0x4f4bc120.
//
// Solidity: function swapKyberAndRepay(uint8 platform, address src, address dest, uint256 srcAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapSession) SwapKyberAndRepay(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, payAmount *big.Int, feeAndRateMode *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapKyberAndRepay(&_SmartWalletSwap.TransactOpts, platform, src, dest, srcAmount, payAmount, feeAndRateMode, platformWallet, hint, useGasToken)
}

// SwapKyberAndRepay is a paid mutator transaction binding the contract method 0x4f4bc120.
//
// Solidity: function swapKyberAndRepay(uint8 platform, address src, address dest, uint256 srcAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bytes hint, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) SwapKyberAndRepay(platform uint8, src common.Address, dest common.Address, srcAmount *big.Int, payAmount *big.Int, feeAndRateMode *big.Int, platformWallet common.Address, hint []byte, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapKyberAndRepay(&_SmartWalletSwap.TransactOpts, platform, src, dest, srcAmount, payAmount, feeAndRateMode, platformWallet, hint, useGasToken)
}

// SwapUniswap is a paid mutator transaction binding the contract method 0x12342114.
//
// Solidity: function swapUniswap(address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactor) SwapUniswap(opts *bind.TransactOpts, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, feeInSrc bool, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "swapUniswap", router, srcAmount, minDestAmount, tradePath, recipient, platformFeeBps, platformWallet, feeInSrc, useGasToken)
}

// SwapUniswap is a paid mutator transaction binding the contract method 0x12342114.
//
// Solidity: function swapUniswap(address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapSession) SwapUniswap(router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, feeInSrc bool, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapUniswap(&_SmartWalletSwap.TransactOpts, router, srcAmount, minDestAmount, tradePath, recipient, platformFeeBps, platformWallet, feeInSrc, useGasToken)
}

// SwapUniswap is a paid mutator transaction binding the contract method 0x12342114.
//
// Solidity: function swapUniswap(address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) SwapUniswap(router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, recipient common.Address, platformFeeBps *big.Int, platformWallet common.Address, feeInSrc bool, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapUniswap(&_SmartWalletSwap.TransactOpts, router, srcAmount, minDestAmount, tradePath, recipient, platformFeeBps, platformWallet, feeInSrc, useGasToken)
}

// SwapUniswapAndDeposit is a paid mutator transaction binding the contract method 0x9059232f.
//
// Solidity: function swapUniswapAndDeposit(uint8 platform, address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, uint256 platformFeeBps, address platformWallet, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactor) SwapUniswapAndDeposit(opts *bind.TransactOpts, platform uint8, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, platformFeeBps *big.Int, platformWallet common.Address, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "swapUniswapAndDeposit", platform, router, srcAmount, minDestAmount, tradePath, platformFeeBps, platformWallet, useGasToken)
}

// SwapUniswapAndDeposit is a paid mutator transaction binding the contract method 0x9059232f.
//
// Solidity: function swapUniswapAndDeposit(uint8 platform, address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, uint256 platformFeeBps, address platformWallet, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapSession) SwapUniswapAndDeposit(platform uint8, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, platformFeeBps *big.Int, platformWallet common.Address, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapUniswapAndDeposit(&_SmartWalletSwap.TransactOpts, platform, router, srcAmount, minDestAmount, tradePath, platformFeeBps, platformWallet, useGasToken)
}

// SwapUniswapAndDeposit is a paid mutator transaction binding the contract method 0x9059232f.
//
// Solidity: function swapUniswapAndDeposit(uint8 platform, address router, uint256 srcAmount, uint256 minDestAmount, address[] tradePath, uint256 platformFeeBps, address platformWallet, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) SwapUniswapAndDeposit(platform uint8, router common.Address, srcAmount *big.Int, minDestAmount *big.Int, tradePath []common.Address, platformFeeBps *big.Int, platformWallet common.Address, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapUniswapAndDeposit(&_SmartWalletSwap.TransactOpts, platform, router, srcAmount, minDestAmount, tradePath, platformFeeBps, platformWallet, useGasToken)
}

// SwapUniswapAndRepay is a paid mutator transaction binding the contract method 0x34068159.
//
// Solidity: function swapUniswapAndRepay(uint8 platform, address router, uint256 srcAmount, uint256 payAmount, address[] tradePath, uint256 feeAndRateMode, address platformWallet, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactor) SwapUniswapAndRepay(opts *bind.TransactOpts, platform uint8, router common.Address, srcAmount *big.Int, payAmount *big.Int, tradePath []common.Address, feeAndRateMode *big.Int, platformWallet common.Address, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "swapUniswapAndRepay", platform, router, srcAmount, payAmount, tradePath, feeAndRateMode, platformWallet, useGasToken)
}

// SwapUniswapAndRepay is a paid mutator transaction binding the contract method 0x34068159.
//
// Solidity: function swapUniswapAndRepay(uint8 platform, address router, uint256 srcAmount, uint256 payAmount, address[] tradePath, uint256 feeAndRateMode, address platformWallet, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapSession) SwapUniswapAndRepay(platform uint8, router common.Address, srcAmount *big.Int, payAmount *big.Int, tradePath []common.Address, feeAndRateMode *big.Int, platformWallet common.Address, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapUniswapAndRepay(&_SmartWalletSwap.TransactOpts, platform, router, srcAmount, payAmount, tradePath, feeAndRateMode, platformWallet, useGasToken)
}

// SwapUniswapAndRepay is a paid mutator transaction binding the contract method 0x34068159.
//
// Solidity: function swapUniswapAndRepay(uint8 platform, address router, uint256 srcAmount, uint256 payAmount, address[] tradePath, uint256 feeAndRateMode, address platformWallet, bool useGasToken) payable returns(uint256 destAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) SwapUniswapAndRepay(platform uint8, router common.Address, srcAmount *big.Int, payAmount *big.Int, tradePath []common.Address, feeAndRateMode *big.Int, platformWallet common.Address, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.SwapUniswapAndRepay(&_SmartWalletSwap.TransactOpts, platform, router, srcAmount, payAmount, tradePath, feeAndRateMode, platformWallet, useGasToken)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.TransferAdmin(&_SmartWalletSwap.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.TransferAdmin(&_SmartWalletSwap.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.TransferAdminQuickly(&_SmartWalletSwap.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.TransferAdminQuickly(&_SmartWalletSwap.TransactOpts, newAdmin)
}

// UpdateBurnGasHelper is a paid mutator transaction binding the contract method 0x6c4f28aa.
//
// Solidity: function updateBurnGasHelper(address _burnGasHelper) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) UpdateBurnGasHelper(opts *bind.TransactOpts, _burnGasHelper common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "updateBurnGasHelper", _burnGasHelper)
}

// UpdateBurnGasHelper is a paid mutator transaction binding the contract method 0x6c4f28aa.
//
// Solidity: function updateBurnGasHelper(address _burnGasHelper) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) UpdateBurnGasHelper(_burnGasHelper common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateBurnGasHelper(&_SmartWalletSwap.TransactOpts, _burnGasHelper)
}

// UpdateBurnGasHelper is a paid mutator transaction binding the contract method 0x6c4f28aa.
//
// Solidity: function updateBurnGasHelper(address _burnGasHelper) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) UpdateBurnGasHelper(_burnGasHelper common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateBurnGasHelper(&_SmartWalletSwap.TransactOpts, _burnGasHelper)
}

// UpdateKyberProxy is a paid mutator transaction binding the contract method 0x974526e6.
//
// Solidity: function updateKyberProxy(address _kyberProxy) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) UpdateKyberProxy(opts *bind.TransactOpts, _kyberProxy common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "updateKyberProxy", _kyberProxy)
}

// UpdateKyberProxy is a paid mutator transaction binding the contract method 0x974526e6.
//
// Solidity: function updateKyberProxy(address _kyberProxy) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) UpdateKyberProxy(_kyberProxy common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateKyberProxy(&_SmartWalletSwap.TransactOpts, _kyberProxy)
}

// UpdateKyberProxy is a paid mutator transaction binding the contract method 0x974526e6.
//
// Solidity: function updateKyberProxy(address _kyberProxy) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) UpdateKyberProxy(_kyberProxy common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateKyberProxy(&_SmartWalletSwap.TransactOpts, _kyberProxy)
}

// UpdateLendingImplementation is a paid mutator transaction binding the contract method 0xb4d59862.
//
// Solidity: function updateLendingImplementation(address newImpl) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) UpdateLendingImplementation(opts *bind.TransactOpts, newImpl common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "updateLendingImplementation", newImpl)
}

// UpdateLendingImplementation is a paid mutator transaction binding the contract method 0xb4d59862.
//
// Solidity: function updateLendingImplementation(address newImpl) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) UpdateLendingImplementation(newImpl common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateLendingImplementation(&_SmartWalletSwap.TransactOpts, newImpl)
}

// UpdateLendingImplementation is a paid mutator transaction binding the contract method 0xb4d59862.
//
// Solidity: function updateLendingImplementation(address newImpl) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) UpdateLendingImplementation(newImpl common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateLendingImplementation(&_SmartWalletSwap.TransactOpts, newImpl)
}

// UpdateSupportedPlatformWallets is a paid mutator transaction binding the contract method 0x296f4120.
//
// Solidity: function updateSupportedPlatformWallets(address[] wallets, bool isSupported) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) UpdateSupportedPlatformWallets(opts *bind.TransactOpts, wallets []common.Address, isSupported bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "updateSupportedPlatformWallets", wallets, isSupported)
}

// UpdateSupportedPlatformWallets is a paid mutator transaction binding the contract method 0x296f4120.
//
// Solidity: function updateSupportedPlatformWallets(address[] wallets, bool isSupported) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) UpdateSupportedPlatformWallets(wallets []common.Address, isSupported bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateSupportedPlatformWallets(&_SmartWalletSwap.TransactOpts, wallets, isSupported)
}

// UpdateSupportedPlatformWallets is a paid mutator transaction binding the contract method 0x296f4120.
//
// Solidity: function updateSupportedPlatformWallets(address[] wallets, bool isSupported) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) UpdateSupportedPlatformWallets(wallets []common.Address, isSupported bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateSupportedPlatformWallets(&_SmartWalletSwap.TransactOpts, wallets, isSupported)
}

// UpdateUniswapRouters is a paid mutator transaction binding the contract method 0xdb43d819.
//
// Solidity: function updateUniswapRouters(address[] _uniswapRouters, bool isAdded) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) UpdateUniswapRouters(opts *bind.TransactOpts, _uniswapRouters []common.Address, isAdded bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "updateUniswapRouters", _uniswapRouters, isAdded)
}

// UpdateUniswapRouters is a paid mutator transaction binding the contract method 0xdb43d819.
//
// Solidity: function updateUniswapRouters(address[] _uniswapRouters, bool isAdded) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) UpdateUniswapRouters(_uniswapRouters []common.Address, isAdded bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateUniswapRouters(&_SmartWalletSwap.TransactOpts, _uniswapRouters, isAdded)
}

// UpdateUniswapRouters is a paid mutator transaction binding the contract method 0xdb43d819.
//
// Solidity: function updateUniswapRouters(address[] _uniswapRouters, bool isAdded) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) UpdateUniswapRouters(_uniswapRouters []common.Address, isAdded bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.UpdateUniswapRouters(&_SmartWalletSwap.TransactOpts, _uniswapRouters, isAdded)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.WithdrawEther(&_SmartWalletSwap.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.WithdrawEther(&_SmartWalletSwap.TransactOpts, amount, sendTo)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x818e80b7.
//
// Solidity: function withdrawFromLendingPlatform(uint8 platform, address token, uint256 amount, uint256 minReturn, bool useGasToken) returns(uint256 returnedAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactor) WithdrawFromLendingPlatform(opts *bind.TransactOpts, platform uint8, token common.Address, amount *big.Int, minReturn *big.Int, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "withdrawFromLendingPlatform", platform, token, amount, minReturn, useGasToken)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x818e80b7.
//
// Solidity: function withdrawFromLendingPlatform(uint8 platform, address token, uint256 amount, uint256 minReturn, bool useGasToken) returns(uint256 returnedAmount)
func (_SmartWalletSwap *SmartWalletSwapSession) WithdrawFromLendingPlatform(platform uint8, token common.Address, amount *big.Int, minReturn *big.Int, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.WithdrawFromLendingPlatform(&_SmartWalletSwap.TransactOpts, platform, token, amount, minReturn, useGasToken)
}

// WithdrawFromLendingPlatform is a paid mutator transaction binding the contract method 0x818e80b7.
//
// Solidity: function withdrawFromLendingPlatform(uint8 platform, address token, uint256 amount, uint256 minReturn, bool useGasToken) returns(uint256 returnedAmount)
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) WithdrawFromLendingPlatform(platform uint8, token common.Address, amount *big.Int, minReturn *big.Int, useGasToken bool) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.WithdrawFromLendingPlatform(&_SmartWalletSwap.TransactOpts, platform, token, amount, minReturn, useGasToken)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SmartWalletSwap *SmartWalletSwapSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.WithdrawToken(&_SmartWalletSwap.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.WithdrawToken(&_SmartWalletSwap.TransactOpts, token, amount, sendTo)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartWalletSwap *SmartWalletSwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartWalletSwap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartWalletSwap *SmartWalletSwapSession) Receive() (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.Receive(&_SmartWalletSwap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartWalletSwap *SmartWalletSwapTransactorSession) Receive() (*types.Transaction, error) {
	return _SmartWalletSwap.Contract.Receive(&_SmartWalletSwap.TransactOpts)
}

// SmartWalletSwapAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the SmartWalletSwap contract.
type SmartWalletSwapAdminClaimedIterator struct {
	Event *SmartWalletSwapAdminClaimed // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapAdminClaimed)
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
		it.Event = new(SmartWalletSwapAdminClaimed)
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
func (it *SmartWalletSwapAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapAdminClaimed represents a AdminClaimed event raised by the SmartWalletSwap contract.
type SmartWalletSwapAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*SmartWalletSwapAdminClaimedIterator, error) {

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapAdminClaimedIterator{contract: _SmartWalletSwap.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapAdminClaimed)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseAdminClaimed(log types.Log) (*SmartWalletSwapAdminClaimed, error) {
	event := new(SmartWalletSwapAdminClaimed)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the SmartWalletSwap contract.
type SmartWalletSwapAlerterAddedIterator struct {
	Event *SmartWalletSwapAlerterAdded // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapAlerterAdded)
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
		it.Event = new(SmartWalletSwapAlerterAdded)
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
func (it *SmartWalletSwapAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapAlerterAdded represents a AlerterAdded event raised by the SmartWalletSwap contract.
type SmartWalletSwapAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*SmartWalletSwapAlerterAddedIterator, error) {

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapAlerterAddedIterator{contract: _SmartWalletSwap.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapAlerterAdded)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseAlerterAdded(log types.Log) (*SmartWalletSwapAlerterAdded, error) {
	event := new(SmartWalletSwapAlerterAdded)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapApprovedAllowancesIterator is returned from FilterApprovedAllowances and is used to iterate over the raw logs and unpacked data for ApprovedAllowances events raised by the SmartWalletSwap contract.
type SmartWalletSwapApprovedAllowancesIterator struct {
	Event *SmartWalletSwapApprovedAllowances // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapApprovedAllowancesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapApprovedAllowances)
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
		it.Event = new(SmartWalletSwapApprovedAllowances)
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
func (it *SmartWalletSwapApprovedAllowancesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapApprovedAllowancesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapApprovedAllowances represents a ApprovedAllowances event raised by the SmartWalletSwap contract.
type SmartWalletSwapApprovedAllowances struct {
	Tokens   []common.Address
	Spenders []common.Address
	IsReset  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovedAllowances is a free log retrieval operation binding the contract event 0xf34c5ed704407ea33d210cd1c76959be869adc80531ee3b3c93229fb606ac16e.
//
// Solidity: event ApprovedAllowances(address[] indexed tokens, address[] indexed spenders, bool isReset)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterApprovedAllowances(opts *bind.FilterOpts, tokens [][]common.Address, spenders [][]common.Address) (*SmartWalletSwapApprovedAllowancesIterator, error) {

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}
	var spendersRule []interface{}
	for _, spendersItem := range spenders {
		spendersRule = append(spendersRule, spendersItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "ApprovedAllowances", tokensRule, spendersRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapApprovedAllowancesIterator{contract: _SmartWalletSwap.contract, event: "ApprovedAllowances", logs: logs, sub: sub}, nil
}

// WatchApprovedAllowances is a free log subscription operation binding the contract event 0xf34c5ed704407ea33d210cd1c76959be869adc80531ee3b3c93229fb606ac16e.
//
// Solidity: event ApprovedAllowances(address[] indexed tokens, address[] indexed spenders, bool isReset)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchApprovedAllowances(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapApprovedAllowances, tokens [][]common.Address, spenders [][]common.Address) (event.Subscription, error) {

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}
	var spendersRule []interface{}
	for _, spendersItem := range spenders {
		spendersRule = append(spendersRule, spendersItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "ApprovedAllowances", tokensRule, spendersRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapApprovedAllowances)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "ApprovedAllowances", log); err != nil {
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
// Solidity: event ApprovedAllowances(address[] indexed tokens, address[] indexed spenders, bool isReset)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseApprovedAllowances(log types.Log) (*SmartWalletSwapApprovedAllowances, error) {
	event := new(SmartWalletSwapApprovedAllowances)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "ApprovedAllowances", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapClaimedPlatformFeesIterator is returned from FilterClaimedPlatformFees and is used to iterate over the raw logs and unpacked data for ClaimedPlatformFees events raised by the SmartWalletSwap contract.
type SmartWalletSwapClaimedPlatformFeesIterator struct {
	Event *SmartWalletSwapClaimedPlatformFees // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapClaimedPlatformFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapClaimedPlatformFees)
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
		it.Event = new(SmartWalletSwapClaimedPlatformFees)
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
func (it *SmartWalletSwapClaimedPlatformFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapClaimedPlatformFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapClaimedPlatformFees represents a ClaimedPlatformFees event raised by the SmartWalletSwap contract.
type SmartWalletSwapClaimedPlatformFees struct {
	Wallets []common.Address
	Tokens  []common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedPlatformFees is a free log retrieval operation binding the contract event 0xcbef5ab8ab58a3dba86704ce124241889d80c3a5fe3caaf4286b7e8c709a6c2d.
//
// Solidity: event ClaimedPlatformFees(address[] indexed wallets, address[] indexed tokens, address claimer)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterClaimedPlatformFees(opts *bind.FilterOpts, wallets [][]common.Address, tokens [][]common.Address) (*SmartWalletSwapClaimedPlatformFeesIterator, error) {

	var walletsRule []interface{}
	for _, walletsItem := range wallets {
		walletsRule = append(walletsRule, walletsItem)
	}
	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "ClaimedPlatformFees", walletsRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapClaimedPlatformFeesIterator{contract: _SmartWalletSwap.contract, event: "ClaimedPlatformFees", logs: logs, sub: sub}, nil
}

// WatchClaimedPlatformFees is a free log subscription operation binding the contract event 0xcbef5ab8ab58a3dba86704ce124241889d80c3a5fe3caaf4286b7e8c709a6c2d.
//
// Solidity: event ClaimedPlatformFees(address[] indexed wallets, address[] indexed tokens, address claimer)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchClaimedPlatformFees(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapClaimedPlatformFees, wallets [][]common.Address, tokens [][]common.Address) (event.Subscription, error) {

	var walletsRule []interface{}
	for _, walletsItem := range wallets {
		walletsRule = append(walletsRule, walletsItem)
	}
	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "ClaimedPlatformFees", walletsRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapClaimedPlatformFees)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "ClaimedPlatformFees", log); err != nil {
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
// Solidity: event ClaimedPlatformFees(address[] indexed wallets, address[] indexed tokens, address claimer)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseClaimedPlatformFees(log types.Log) (*SmartWalletSwapClaimedPlatformFees, error) {
	event := new(SmartWalletSwapClaimedPlatformFees)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "ClaimedPlatformFees", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the SmartWalletSwap contract.
type SmartWalletSwapEtherWithdrawIterator struct {
	Event *SmartWalletSwapEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapEtherWithdraw)
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
		it.Event = new(SmartWalletSwapEtherWithdraw)
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
func (it *SmartWalletSwapEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapEtherWithdraw represents a EtherWithdraw event raised by the SmartWalletSwap contract.
type SmartWalletSwapEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*SmartWalletSwapEtherWithdrawIterator, error) {

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapEtherWithdrawIterator{contract: _SmartWalletSwap.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapEtherWithdraw)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseEtherWithdraw(log types.Log) (*SmartWalletSwapEtherWithdraw, error) {
	event := new(SmartWalletSwapEtherWithdraw)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapKyberTradeIterator is returned from FilterKyberTrade and is used to iterate over the raw logs and unpacked data for KyberTrade events raised by the SmartWalletSwap contract.
type SmartWalletSwapKyberTradeIterator struct {
	Event *SmartWalletSwapKyberTrade // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapKyberTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapKyberTrade)
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
		it.Event = new(SmartWalletSwapKyberTrade)
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
func (it *SmartWalletSwapKyberTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapKyberTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapKyberTrade represents a KyberTrade event raised by the SmartWalletSwap contract.
type SmartWalletSwapKyberTrade struct {
	Trader         common.Address
	Src            common.Address
	Dest           common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	Recipient      common.Address
	PlatformFeeBps *big.Int
	PlatformWallet common.Address
	Hint           []byte
	UseGasToken    bool
	NumGasBurns    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKyberTrade is a free log retrieval operation binding the contract event 0x44c89632efdca28ac6efe805efab8b5a00755333515fff593bfc3546efba8a1c.
//
// Solidity: event KyberTrade(address indexed trader, address indexed src, address indexed dest, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterKyberTrade(opts *bind.FilterOpts, trader []common.Address, src []common.Address, dest []common.Address) (*SmartWalletSwapKyberTradeIterator, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "KyberTrade", traderRule, srcRule, destRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapKyberTradeIterator{contract: _SmartWalletSwap.contract, event: "KyberTrade", logs: logs, sub: sub}, nil
}

// WatchKyberTrade is a free log subscription operation binding the contract event 0x44c89632efdca28ac6efe805efab8b5a00755333515fff593bfc3546efba8a1c.
//
// Solidity: event KyberTrade(address indexed trader, address indexed src, address indexed dest, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchKyberTrade(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapKyberTrade, trader []common.Address, src []common.Address, dest []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "KyberTrade", traderRule, srcRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapKyberTrade)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "KyberTrade", log); err != nil {
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

// ParseKyberTrade is a log parse operation binding the contract event 0x44c89632efdca28ac6efe805efab8b5a00755333515fff593bfc3546efba8a1c.
//
// Solidity: event KyberTrade(address indexed trader, address indexed src, address indexed dest, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseKyberTrade(log types.Log) (*SmartWalletSwapKyberTrade, error) {
	event := new(SmartWalletSwapKyberTrade)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "KyberTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapKyberTradeAndDepositIterator is returned from FilterKyberTradeAndDeposit and is used to iterate over the raw logs and unpacked data for KyberTradeAndDeposit events raised by the SmartWalletSwap contract.
type SmartWalletSwapKyberTradeAndDepositIterator struct {
	Event *SmartWalletSwapKyberTradeAndDeposit // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapKyberTradeAndDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapKyberTradeAndDeposit)
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
		it.Event = new(SmartWalletSwapKyberTradeAndDeposit)
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
func (it *SmartWalletSwapKyberTradeAndDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapKyberTradeAndDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapKyberTradeAndDeposit represents a KyberTradeAndDeposit event raised by the SmartWalletSwap contract.
type SmartWalletSwapKyberTradeAndDeposit struct {
	Trader         common.Address
	Platform       uint8
	Src            common.Address
	Dest           common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PlatformFeeBps *big.Int
	PlatformWallet common.Address
	Hint           []byte
	UseGasToken    bool
	NumGasBurns    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKyberTradeAndDeposit is a free log retrieval operation binding the contract event 0xb1ed35e59208842b3a1c3ab171ede44b088769088286d7a403ef249ed71c7396.
//
// Solidity: event KyberTradeAndDeposit(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterKyberTradeAndDeposit(opts *bind.FilterOpts, trader []common.Address, platform []uint8, dest []common.Address) (*SmartWalletSwapKyberTradeAndDepositIterator, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "KyberTradeAndDeposit", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapKyberTradeAndDepositIterator{contract: _SmartWalletSwap.contract, event: "KyberTradeAndDeposit", logs: logs, sub: sub}, nil
}

// WatchKyberTradeAndDeposit is a free log subscription operation binding the contract event 0xb1ed35e59208842b3a1c3ab171ede44b088769088286d7a403ef249ed71c7396.
//
// Solidity: event KyberTradeAndDeposit(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchKyberTradeAndDeposit(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapKyberTradeAndDeposit, trader []common.Address, platform []uint8, dest []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "KyberTradeAndDeposit", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapKyberTradeAndDeposit)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "KyberTradeAndDeposit", log); err != nil {
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

// ParseKyberTradeAndDeposit is a log parse operation binding the contract event 0xb1ed35e59208842b3a1c3ab171ede44b088769088286d7a403ef249ed71c7396.
//
// Solidity: event KyberTradeAndDeposit(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseKyberTradeAndDeposit(log types.Log) (*SmartWalletSwapKyberTradeAndDeposit, error) {
	event := new(SmartWalletSwapKyberTradeAndDeposit)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "KyberTradeAndDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapKyberTradeAndRepayIterator is returned from FilterKyberTradeAndRepay and is used to iterate over the raw logs and unpacked data for KyberTradeAndRepay events raised by the SmartWalletSwap contract.
type SmartWalletSwapKyberTradeAndRepayIterator struct {
	Event *SmartWalletSwapKyberTradeAndRepay // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapKyberTradeAndRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapKyberTradeAndRepay)
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
		it.Event = new(SmartWalletSwapKyberTradeAndRepay)
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
func (it *SmartWalletSwapKyberTradeAndRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapKyberTradeAndRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapKyberTradeAndRepay represents a KyberTradeAndRepay event raised by the SmartWalletSwap contract.
type SmartWalletSwapKyberTradeAndRepay struct {
	Trader         common.Address
	Platform       uint8
	Src            common.Address
	Dest           common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PayAmount      *big.Int
	FeeAndRateMode *big.Int
	PlatformWallet common.Address
	Hint           []byte
	UseGasToken    bool
	NumGasBurns    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKyberTradeAndRepay is a free log retrieval operation binding the contract event 0xb3eef3be8f410e0f8f359f08685cf0e07e0b597c4498e95729d39bac8c493019.
//
// Solidity: event KyberTradeAndRepay(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterKyberTradeAndRepay(opts *bind.FilterOpts, trader []common.Address, platform []uint8, dest []common.Address) (*SmartWalletSwapKyberTradeAndRepayIterator, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "KyberTradeAndRepay", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapKyberTradeAndRepayIterator{contract: _SmartWalletSwap.contract, event: "KyberTradeAndRepay", logs: logs, sub: sub}, nil
}

// WatchKyberTradeAndRepay is a free log subscription operation binding the contract event 0xb3eef3be8f410e0f8f359f08685cf0e07e0b597c4498e95729d39bac8c493019.
//
// Solidity: event KyberTradeAndRepay(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchKyberTradeAndRepay(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapKyberTradeAndRepay, trader []common.Address, platform []uint8, dest []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "KyberTradeAndRepay", traderRule, platformRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapKyberTradeAndRepay)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "KyberTradeAndRepay", log); err != nil {
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

// ParseKyberTradeAndRepay is a log parse operation binding the contract event 0xb3eef3be8f410e0f8f359f08685cf0e07e0b597c4498e95729d39bac8c493019.
//
// Solidity: event KyberTradeAndRepay(address indexed trader, uint8 indexed platform, address src, address indexed dest, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bytes hint, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseKyberTradeAndRepay(log types.Log) (*SmartWalletSwapKyberTradeAndRepay, error) {
	event := new(SmartWalletSwapKyberTradeAndRepay)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "KyberTradeAndRepay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the SmartWalletSwap contract.
type SmartWalletSwapOperatorAddedIterator struct {
	Event *SmartWalletSwapOperatorAdded // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapOperatorAdded)
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
		it.Event = new(SmartWalletSwapOperatorAdded)
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
func (it *SmartWalletSwapOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapOperatorAdded represents a OperatorAdded event raised by the SmartWalletSwap contract.
type SmartWalletSwapOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*SmartWalletSwapOperatorAddedIterator, error) {

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapOperatorAddedIterator{contract: _SmartWalletSwap.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapOperatorAdded)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseOperatorAdded(log types.Log) (*SmartWalletSwapOperatorAdded, error) {
	event := new(SmartWalletSwapOperatorAdded)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the SmartWalletSwap contract.
type SmartWalletSwapTokenWithdrawIterator struct {
	Event *SmartWalletSwapTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapTokenWithdraw)
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
		it.Event = new(SmartWalletSwapTokenWithdraw)
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
func (it *SmartWalletSwapTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapTokenWithdraw represents a TokenWithdraw event raised by the SmartWalletSwap contract.
type SmartWalletSwapTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*SmartWalletSwapTokenWithdrawIterator, error) {

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapTokenWithdrawIterator{contract: _SmartWalletSwap.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapTokenWithdraw)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseTokenWithdraw(log types.Log) (*SmartWalletSwapTokenWithdraw, error) {
	event := new(SmartWalletSwapTokenWithdraw)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the SmartWalletSwap contract.
type SmartWalletSwapTransferAdminPendingIterator struct {
	Event *SmartWalletSwapTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapTransferAdminPending)
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
		it.Event = new(SmartWalletSwapTransferAdminPending)
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
func (it *SmartWalletSwapTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapTransferAdminPending represents a TransferAdminPending event raised by the SmartWalletSwap contract.
type SmartWalletSwapTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*SmartWalletSwapTransferAdminPendingIterator, error) {

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapTransferAdminPendingIterator{contract: _SmartWalletSwap.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapTransferAdminPending)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseTransferAdminPending(log types.Log) (*SmartWalletSwapTransferAdminPending, error) {
	event := new(SmartWalletSwapTransferAdminPending)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUniswapTradeIterator is returned from FilterUniswapTrade and is used to iterate over the raw logs and unpacked data for UniswapTrade events raised by the SmartWalletSwap contract.
type SmartWalletSwapUniswapTradeIterator struct {
	Event *SmartWalletSwapUniswapTrade // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUniswapTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUniswapTrade)
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
		it.Event = new(SmartWalletSwapUniswapTrade)
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
func (it *SmartWalletSwapUniswapTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUniswapTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUniswapTrade represents a UniswapTrade event raised by the SmartWalletSwap contract.
type SmartWalletSwapUniswapTrade struct {
	Trader         common.Address
	Router         common.Address
	TradePath      []common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	Recipient      common.Address
	PlatformFeeBps *big.Int
	PlatformWallet common.Address
	FeeInSrc       bool
	UseGasToken    bool
	NumGasBurns    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUniswapTrade is a free log retrieval operation binding the contract event 0x37bf11bfa5908c918cfb4cb832b178dd23b14602a8e9ad662f2166e604cd1518.
//
// Solidity: event UniswapTrade(address indexed trader, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUniswapTrade(opts *bind.FilterOpts, trader []common.Address, router []common.Address) (*SmartWalletSwapUniswapTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UniswapTrade", traderRule, routerRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUniswapTradeIterator{contract: _SmartWalletSwap.contract, event: "UniswapTrade", logs: logs, sub: sub}, nil
}

// WatchUniswapTrade is a free log subscription operation binding the contract event 0x37bf11bfa5908c918cfb4cb832b178dd23b14602a8e9ad662f2166e604cd1518.
//
// Solidity: event UniswapTrade(address indexed trader, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUniswapTrade(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUniswapTrade, trader []common.Address, router []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UniswapTrade", traderRule, routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUniswapTrade)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UniswapTrade", log); err != nil {
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

// ParseUniswapTrade is a log parse operation binding the contract event 0x37bf11bfa5908c918cfb4cb832b178dd23b14602a8e9ad662f2166e604cd1518.
//
// Solidity: event UniswapTrade(address indexed trader, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, address recipient, uint256 platformFeeBps, address platformWallet, bool feeInSrc, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUniswapTrade(log types.Log) (*SmartWalletSwapUniswapTrade, error) {
	event := new(SmartWalletSwapUniswapTrade)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UniswapTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUniswapTradeAndDepositIterator is returned from FilterUniswapTradeAndDeposit and is used to iterate over the raw logs and unpacked data for UniswapTradeAndDeposit events raised by the SmartWalletSwap contract.
type SmartWalletSwapUniswapTradeAndDepositIterator struct {
	Event *SmartWalletSwapUniswapTradeAndDeposit // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUniswapTradeAndDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUniswapTradeAndDeposit)
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
		it.Event = new(SmartWalletSwapUniswapTradeAndDeposit)
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
func (it *SmartWalletSwapUniswapTradeAndDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUniswapTradeAndDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUniswapTradeAndDeposit represents a UniswapTradeAndDeposit event raised by the SmartWalletSwap contract.
type SmartWalletSwapUniswapTradeAndDeposit struct {
	Trader         common.Address
	Platform       uint8
	Router         common.Address
	TradePath      []common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PlatformFeeBps *big.Int
	PlatformWallet common.Address
	UseGasToken    bool
	NumGasBurns    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUniswapTradeAndDeposit is a free log retrieval operation binding the contract event 0xc3ae739d5ee6a65858928b4498f473d2bb0353c31fbe3ace7ac6fee493058f20.
//
// Solidity: event UniswapTradeAndDeposit(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUniswapTradeAndDeposit(opts *bind.FilterOpts, trader []common.Address, platform []uint8, router []common.Address) (*SmartWalletSwapUniswapTradeAndDepositIterator, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UniswapTradeAndDeposit", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUniswapTradeAndDepositIterator{contract: _SmartWalletSwap.contract, event: "UniswapTradeAndDeposit", logs: logs, sub: sub}, nil
}

// WatchUniswapTradeAndDeposit is a free log subscription operation binding the contract event 0xc3ae739d5ee6a65858928b4498f473d2bb0353c31fbe3ace7ac6fee493058f20.
//
// Solidity: event UniswapTradeAndDeposit(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUniswapTradeAndDeposit(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUniswapTradeAndDeposit, trader []common.Address, platform []uint8, router []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UniswapTradeAndDeposit", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUniswapTradeAndDeposit)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UniswapTradeAndDeposit", log); err != nil {
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

// ParseUniswapTradeAndDeposit is a log parse operation binding the contract event 0xc3ae739d5ee6a65858928b4498f473d2bb0353c31fbe3ace7ac6fee493058f20.
//
// Solidity: event UniswapTradeAndDeposit(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 platformFeeBps, address platformWallet, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUniswapTradeAndDeposit(log types.Log) (*SmartWalletSwapUniswapTradeAndDeposit, error) {
	event := new(SmartWalletSwapUniswapTradeAndDeposit)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UniswapTradeAndDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUniswapTradeAndRepayIterator is returned from FilterUniswapTradeAndRepay and is used to iterate over the raw logs and unpacked data for UniswapTradeAndRepay events raised by the SmartWalletSwap contract.
type SmartWalletSwapUniswapTradeAndRepayIterator struct {
	Event *SmartWalletSwapUniswapTradeAndRepay // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUniswapTradeAndRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUniswapTradeAndRepay)
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
		it.Event = new(SmartWalletSwapUniswapTradeAndRepay)
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
func (it *SmartWalletSwapUniswapTradeAndRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUniswapTradeAndRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUniswapTradeAndRepay represents a UniswapTradeAndRepay event raised by the SmartWalletSwap contract.
type SmartWalletSwapUniswapTradeAndRepay struct {
	Trader         common.Address
	Platform       uint8
	Router         common.Address
	TradePath      []common.Address
	SrcAmount      *big.Int
	DestAmount     *big.Int
	PayAmount      *big.Int
	FeeAndRateMode *big.Int
	PlatformWallet common.Address
	UseGasToken    bool
	NumGasBurns    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUniswapTradeAndRepay is a free log retrieval operation binding the contract event 0x37bfe96bee0ff17be045b3a5488512d922a4e1d791db1f9710bb3263a79ae581.
//
// Solidity: event UniswapTradeAndRepay(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUniswapTradeAndRepay(opts *bind.FilterOpts, trader []common.Address, platform []uint8, router []common.Address) (*SmartWalletSwapUniswapTradeAndRepayIterator, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UniswapTradeAndRepay", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUniswapTradeAndRepayIterator{contract: _SmartWalletSwap.contract, event: "UniswapTradeAndRepay", logs: logs, sub: sub}, nil
}

// WatchUniswapTradeAndRepay is a free log subscription operation binding the contract event 0x37bfe96bee0ff17be045b3a5488512d922a4e1d791db1f9710bb3263a79ae581.
//
// Solidity: event UniswapTradeAndRepay(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUniswapTradeAndRepay(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUniswapTradeAndRepay, trader []common.Address, platform []uint8, router []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UniswapTradeAndRepay", traderRule, platformRule, routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUniswapTradeAndRepay)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UniswapTradeAndRepay", log); err != nil {
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

// ParseUniswapTradeAndRepay is a log parse operation binding the contract event 0x37bfe96bee0ff17be045b3a5488512d922a4e1d791db1f9710bb3263a79ae581.
//
// Solidity: event UniswapTradeAndRepay(address indexed trader, uint8 indexed platform, address indexed router, address[] tradePath, uint256 srcAmount, uint256 destAmount, uint256 payAmount, uint256 feeAndRateMode, address platformWallet, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUniswapTradeAndRepay(log types.Log) (*SmartWalletSwapUniswapTradeAndRepay, error) {
	event := new(SmartWalletSwapUniswapTradeAndRepay)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UniswapTradeAndRepay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUpdateKyberProxyIterator is returned from FilterUpdateKyberProxy and is used to iterate over the raw logs and unpacked data for UpdateKyberProxy events raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdateKyberProxyIterator struct {
	Event *SmartWalletSwapUpdateKyberProxy // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUpdateKyberProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUpdateKyberProxy)
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
		it.Event = new(SmartWalletSwapUpdateKyberProxy)
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
func (it *SmartWalletSwapUpdateKyberProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUpdateKyberProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUpdateKyberProxy represents a UpdateKyberProxy event raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdateKyberProxy struct {
	NewProxy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateKyberProxy is a free log retrieval operation binding the contract event 0x829471216993aa9fe9427c9e091a30bcbbf8d52b472ee89a0ce1a7ec545042c5.
//
// Solidity: event UpdateKyberProxy(address indexed newProxy)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUpdateKyberProxy(opts *bind.FilterOpts, newProxy []common.Address) (*SmartWalletSwapUpdateKyberProxyIterator, error) {

	var newProxyRule []interface{}
	for _, newProxyItem := range newProxy {
		newProxyRule = append(newProxyRule, newProxyItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UpdateKyberProxy", newProxyRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUpdateKyberProxyIterator{contract: _SmartWalletSwap.contract, event: "UpdateKyberProxy", logs: logs, sub: sub}, nil
}

// WatchUpdateKyberProxy is a free log subscription operation binding the contract event 0x829471216993aa9fe9427c9e091a30bcbbf8d52b472ee89a0ce1a7ec545042c5.
//
// Solidity: event UpdateKyberProxy(address indexed newProxy)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUpdateKyberProxy(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUpdateKyberProxy, newProxy []common.Address) (event.Subscription, error) {

	var newProxyRule []interface{}
	for _, newProxyItem := range newProxy {
		newProxyRule = append(newProxyRule, newProxyItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UpdateKyberProxy", newProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUpdateKyberProxy)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdateKyberProxy", log); err != nil {
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

// ParseUpdateKyberProxy is a log parse operation binding the contract event 0x829471216993aa9fe9427c9e091a30bcbbf8d52b472ee89a0ce1a7ec545042c5.
//
// Solidity: event UpdateKyberProxy(address indexed newProxy)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUpdateKyberProxy(log types.Log) (*SmartWalletSwapUpdateKyberProxy, error) {
	event := new(SmartWalletSwapUpdateKyberProxy)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdateKyberProxy", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUpdateUniswapRoutersIterator is returned from FilterUpdateUniswapRouters and is used to iterate over the raw logs and unpacked data for UpdateUniswapRouters events raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdateUniswapRoutersIterator struct {
	Event *SmartWalletSwapUpdateUniswapRouters // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUpdateUniswapRoutersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUpdateUniswapRouters)
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
		it.Event = new(SmartWalletSwapUpdateUniswapRouters)
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
func (it *SmartWalletSwapUpdateUniswapRoutersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUpdateUniswapRoutersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUpdateUniswapRouters represents a UpdateUniswapRouters event raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdateUniswapRouters struct {
	UniswapRouters []common.Address
	IsAdded        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateUniswapRouters is a free log retrieval operation binding the contract event 0x8526e571790cfb73a8e0ffa80e3e8723ac0eceddb37aed708bb656ac1856643e.
//
// Solidity: event UpdateUniswapRouters(address[] indexed uniswapRouters, bool isAdded)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUpdateUniswapRouters(opts *bind.FilterOpts, uniswapRouters [][]common.Address) (*SmartWalletSwapUpdateUniswapRoutersIterator, error) {

	var uniswapRoutersRule []interface{}
	for _, uniswapRoutersItem := range uniswapRouters {
		uniswapRoutersRule = append(uniswapRoutersRule, uniswapRoutersItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UpdateUniswapRouters", uniswapRoutersRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUpdateUniswapRoutersIterator{contract: _SmartWalletSwap.contract, event: "UpdateUniswapRouters", logs: logs, sub: sub}, nil
}

// WatchUpdateUniswapRouters is a free log subscription operation binding the contract event 0x8526e571790cfb73a8e0ffa80e3e8723ac0eceddb37aed708bb656ac1856643e.
//
// Solidity: event UpdateUniswapRouters(address[] indexed uniswapRouters, bool isAdded)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUpdateUniswapRouters(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUpdateUniswapRouters, uniswapRouters [][]common.Address) (event.Subscription, error) {

	var uniswapRoutersRule []interface{}
	for _, uniswapRoutersItem := range uniswapRouters {
		uniswapRoutersRule = append(uniswapRoutersRule, uniswapRoutersItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UpdateUniswapRouters", uniswapRoutersRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUpdateUniswapRouters)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdateUniswapRouters", log); err != nil {
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

// ParseUpdateUniswapRouters is a log parse operation binding the contract event 0x8526e571790cfb73a8e0ffa80e3e8723ac0eceddb37aed708bb656ac1856643e.
//
// Solidity: event UpdateUniswapRouters(address[] indexed uniswapRouters, bool isAdded)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUpdateUniswapRouters(log types.Log) (*SmartWalletSwapUpdateUniswapRouters, error) {
	event := new(SmartWalletSwapUpdateUniswapRouters)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdateUniswapRouters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUpdatedBurnGasHelperIterator is returned from FilterUpdatedBurnGasHelper and is used to iterate over the raw logs and unpacked data for UpdatedBurnGasHelper events raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdatedBurnGasHelperIterator struct {
	Event *SmartWalletSwapUpdatedBurnGasHelper // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUpdatedBurnGasHelperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUpdatedBurnGasHelper)
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
		it.Event = new(SmartWalletSwapUpdatedBurnGasHelper)
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
func (it *SmartWalletSwapUpdatedBurnGasHelperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUpdatedBurnGasHelperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUpdatedBurnGasHelper represents a UpdatedBurnGasHelper event raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdatedBurnGasHelper struct {
	GasHelper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBurnGasHelper is a free log retrieval operation binding the contract event 0xe9bf230d574d9f3218c8b7bc5f152365e71de8e84420960ff43b7d941dcb2c70.
//
// Solidity: event UpdatedBurnGasHelper(address indexed gasHelper)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUpdatedBurnGasHelper(opts *bind.FilterOpts, gasHelper []common.Address) (*SmartWalletSwapUpdatedBurnGasHelperIterator, error) {

	var gasHelperRule []interface{}
	for _, gasHelperItem := range gasHelper {
		gasHelperRule = append(gasHelperRule, gasHelperItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UpdatedBurnGasHelper", gasHelperRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUpdatedBurnGasHelperIterator{contract: _SmartWalletSwap.contract, event: "UpdatedBurnGasHelper", logs: logs, sub: sub}, nil
}

// WatchUpdatedBurnGasHelper is a free log subscription operation binding the contract event 0xe9bf230d574d9f3218c8b7bc5f152365e71de8e84420960ff43b7d941dcb2c70.
//
// Solidity: event UpdatedBurnGasHelper(address indexed gasHelper)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUpdatedBurnGasHelper(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUpdatedBurnGasHelper, gasHelper []common.Address) (event.Subscription, error) {

	var gasHelperRule []interface{}
	for _, gasHelperItem := range gasHelper {
		gasHelperRule = append(gasHelperRule, gasHelperItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UpdatedBurnGasHelper", gasHelperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUpdatedBurnGasHelper)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdatedBurnGasHelper", log); err != nil {
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

// ParseUpdatedBurnGasHelper is a log parse operation binding the contract event 0xe9bf230d574d9f3218c8b7bc5f152365e71de8e84420960ff43b7d941dcb2c70.
//
// Solidity: event UpdatedBurnGasHelper(address indexed gasHelper)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUpdatedBurnGasHelper(log types.Log) (*SmartWalletSwapUpdatedBurnGasHelper, error) {
	event := new(SmartWalletSwapUpdatedBurnGasHelper)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdatedBurnGasHelper", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUpdatedLendingImplementationIterator is returned from FilterUpdatedLendingImplementation and is used to iterate over the raw logs and unpacked data for UpdatedLendingImplementation events raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdatedLendingImplementationIterator struct {
	Event *SmartWalletSwapUpdatedLendingImplementation // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUpdatedLendingImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUpdatedLendingImplementation)
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
		it.Event = new(SmartWalletSwapUpdatedLendingImplementation)
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
func (it *SmartWalletSwapUpdatedLendingImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUpdatedLendingImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUpdatedLendingImplementation represents a UpdatedLendingImplementation event raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdatedLendingImplementation struct {
	Impl common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdatedLendingImplementation is a free log retrieval operation binding the contract event 0x7d914f4a47f2d06bf2a4f2679866f53f24fd19f5fe6923bbe6d8e54e3cd41536.
//
// Solidity: event UpdatedLendingImplementation(address indexed impl)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUpdatedLendingImplementation(opts *bind.FilterOpts, impl []common.Address) (*SmartWalletSwapUpdatedLendingImplementationIterator, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UpdatedLendingImplementation", implRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUpdatedLendingImplementationIterator{contract: _SmartWalletSwap.contract, event: "UpdatedLendingImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdatedLendingImplementation is a free log subscription operation binding the contract event 0x7d914f4a47f2d06bf2a4f2679866f53f24fd19f5fe6923bbe6d8e54e3cd41536.
//
// Solidity: event UpdatedLendingImplementation(address indexed impl)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUpdatedLendingImplementation(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUpdatedLendingImplementation, impl []common.Address) (event.Subscription, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UpdatedLendingImplementation", implRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUpdatedLendingImplementation)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdatedLendingImplementation", log); err != nil {
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
// Solidity: event UpdatedLendingImplementation(address indexed impl)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUpdatedLendingImplementation(log types.Log) (*SmartWalletSwapUpdatedLendingImplementation, error) {
	event := new(SmartWalletSwapUpdatedLendingImplementation)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdatedLendingImplementation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapUpdatedSupportedPlatformWalletsIterator is returned from FilterUpdatedSupportedPlatformWallets and is used to iterate over the raw logs and unpacked data for UpdatedSupportedPlatformWallets events raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdatedSupportedPlatformWalletsIterator struct {
	Event *SmartWalletSwapUpdatedSupportedPlatformWallets // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapUpdatedSupportedPlatformWalletsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapUpdatedSupportedPlatformWallets)
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
		it.Event = new(SmartWalletSwapUpdatedSupportedPlatformWallets)
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
func (it *SmartWalletSwapUpdatedSupportedPlatformWalletsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapUpdatedSupportedPlatformWalletsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapUpdatedSupportedPlatformWallets represents a UpdatedSupportedPlatformWallets event raised by the SmartWalletSwap contract.
type SmartWalletSwapUpdatedSupportedPlatformWallets struct {
	Wallets     []common.Address
	IsSupported bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSupportedPlatformWallets is a free log retrieval operation binding the contract event 0xac650f539fc5264cdad2074ff608d20b54013400070e4e5ef39125f67d8c986c.
//
// Solidity: event UpdatedSupportedPlatformWallets(address[] indexed wallets, bool indexed isSupported)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterUpdatedSupportedPlatformWallets(opts *bind.FilterOpts, wallets [][]common.Address, isSupported []bool) (*SmartWalletSwapUpdatedSupportedPlatformWalletsIterator, error) {

	var walletsRule []interface{}
	for _, walletsItem := range wallets {
		walletsRule = append(walletsRule, walletsItem)
	}
	var isSupportedRule []interface{}
	for _, isSupportedItem := range isSupported {
		isSupportedRule = append(isSupportedRule, isSupportedItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "UpdatedSupportedPlatformWallets", walletsRule, isSupportedRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapUpdatedSupportedPlatformWalletsIterator{contract: _SmartWalletSwap.contract, event: "UpdatedSupportedPlatformWallets", logs: logs, sub: sub}, nil
}

// WatchUpdatedSupportedPlatformWallets is a free log subscription operation binding the contract event 0xac650f539fc5264cdad2074ff608d20b54013400070e4e5ef39125f67d8c986c.
//
// Solidity: event UpdatedSupportedPlatformWallets(address[] indexed wallets, bool indexed isSupported)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchUpdatedSupportedPlatformWallets(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapUpdatedSupportedPlatformWallets, wallets [][]common.Address, isSupported []bool) (event.Subscription, error) {

	var walletsRule []interface{}
	for _, walletsItem := range wallets {
		walletsRule = append(walletsRule, walletsItem)
	}
	var isSupportedRule []interface{}
	for _, isSupportedItem := range isSupported {
		isSupportedRule = append(isSupportedRule, isSupportedItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "UpdatedSupportedPlatformWallets", walletsRule, isSupportedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapUpdatedSupportedPlatformWallets)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdatedSupportedPlatformWallets", log); err != nil {
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
// Solidity: event UpdatedSupportedPlatformWallets(address[] indexed wallets, bool indexed isSupported)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseUpdatedSupportedPlatformWallets(log types.Log) (*SmartWalletSwapUpdatedSupportedPlatformWallets, error) {
	event := new(SmartWalletSwapUpdatedSupportedPlatformWallets)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "UpdatedSupportedPlatformWallets", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartWalletSwapWithdrawFromLendingIterator is returned from FilterWithdrawFromLending and is used to iterate over the raw logs and unpacked data for WithdrawFromLending events raised by the SmartWalletSwap contract.
type SmartWalletSwapWithdrawFromLendingIterator struct {
	Event *SmartWalletSwapWithdrawFromLending // Event containing the contract specifics and raw log

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
func (it *SmartWalletSwapWithdrawFromLendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartWalletSwapWithdrawFromLending)
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
		it.Event = new(SmartWalletSwapWithdrawFromLending)
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
func (it *SmartWalletSwapWithdrawFromLendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartWalletSwapWithdrawFromLendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartWalletSwapWithdrawFromLending represents a WithdrawFromLending event raised by the SmartWalletSwap contract.
type SmartWalletSwapWithdrawFromLending struct {
	Platform           uint8
	Token              common.Address
	Amount             *big.Int
	MinReturn          *big.Int
	ActualReturnAmount *big.Int
	UseGasToken        bool
	NumGasBurns        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFromLending is a free log retrieval operation binding the contract event 0xdbbd6e8a55e3d0927fc92441827df524a59b6f5249339c5504a8b272d6cc08d8.
//
// Solidity: event WithdrawFromLending(uint8 indexed platform, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) FilterWithdrawFromLending(opts *bind.FilterOpts, platform []uint8) (*SmartWalletSwapWithdrawFromLendingIterator, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.FilterLogs(opts, "WithdrawFromLending", platformRule)
	if err != nil {
		return nil, err
	}
	return &SmartWalletSwapWithdrawFromLendingIterator{contract: _SmartWalletSwap.contract, event: "WithdrawFromLending", logs: logs, sub: sub}, nil
}

// WatchWithdrawFromLending is a free log subscription operation binding the contract event 0xdbbd6e8a55e3d0927fc92441827df524a59b6f5249339c5504a8b272d6cc08d8.
//
// Solidity: event WithdrawFromLending(uint8 indexed platform, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) WatchWithdrawFromLending(opts *bind.WatchOpts, sink chan<- *SmartWalletSwapWithdrawFromLending, platform []uint8) (event.Subscription, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _SmartWalletSwap.contract.WatchLogs(opts, "WithdrawFromLending", platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartWalletSwapWithdrawFromLending)
				if err := _SmartWalletSwap.contract.UnpackLog(event, "WithdrawFromLending", log); err != nil {
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

// ParseWithdrawFromLending is a log parse operation binding the contract event 0xdbbd6e8a55e3d0927fc92441827df524a59b6f5249339c5504a8b272d6cc08d8.
//
// Solidity: event WithdrawFromLending(uint8 indexed platform, address token, uint256 amount, uint256 minReturn, uint256 actualReturnAmount, bool useGasToken, uint256 numGasBurns)
func (_SmartWalletSwap *SmartWalletSwapFilterer) ParseWithdrawFromLending(log types.Log) (*SmartWalletSwapWithdrawFromLending, error) {
	event := new(SmartWalletSwapWithdrawFromLending)
	if err := _SmartWalletSwap.contract.UnpackLog(event, "WithdrawFromLending", log); err != nil {
		return nil, err
	}
	return event, nil
}
