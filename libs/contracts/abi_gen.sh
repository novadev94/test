#!/usr/bin/env bash

# set -euo pipefail

# readonly src_dir=/tmp/kyber-api/src/github.com/ethereum
# export GOPATH=/tmp/kyber-api
# export PATH=$GOPATH/bin:$PATH

# mkdir -p "$src_dir"
# cd "$src_dir"
# [[ -d go-ethereum ]] || git clone https://github.com/ethereum/go-ethereum.git
# go install -v github.com/ethereum/go-ethereum/cmd/abigen

# Guarantee out files will be in the same folder as this script
SCRIPT_PATH=${0%/*}
if [ "$0" != "$SCRIPT_PATH" ] && [ "$SCRIPT_PATH" != "" ]; then
    cd $SCRIPT_PATH
fi

#abigen -abi ./chainlink_aggregator.abi -pkg contracts -type ChainlinkAggregator -out ./chainlink_aggregator_abi.go
#abigen -abi ./kyber_proxy.abi -pkg contracts -type KyberProxy -out ./kyber_proxy_abi.go
#abigen -abi ./kyber_wrapper.abi -pkg contracts -type KyberWrapper -out ./kyber_wrapper_abi.go
#abigen -abi ./uniswap_router_v2.abi -pkg contracts -type UniswapRouterV2 -out ./uniswap_router_v2_abi.go
abigen -abi ./smart_wallet_swap.abi -pkg contracts -type SmartWalletSwap -out ./smart_wallet_swap_abi.go
#abigen -abi ./uniswap_v2_pair.abi -pkg contracts -type UniswapV2Pair -out ./uniswap_v2_pair_abi.go
abigen -abi ./erc20.abi -pkg contracts -type Erc20 -out ./erc20_abi.go
abigen -abi ./weth.abi -pkg contracts -type Weth -out ./weth_abi.go
abigen -abi ./fetch_aave_data_wrapper.abi -pkg contracts -type FetchAaveDataWrapper -out ./fetch_aave_data_wrapper_abi.go
abigen -abi ./aave_price_oracle.abi -pkg contracts -type AavePriceOracle -out ./aave_price_oracle_abi.go
abigen -abi ./aavev2_protocol_data_provider.abi -pkg contracts -type AaveV2ProtocolDataProvider -out ./aavev2_protocol_data_provider_abi.go
abigen -abi ./aavev1_lending_pool.abi -pkg contracts -type AaveV1LendingPool -out ./aavev1_lending_pool_abi.go
abigen -abi ./comptroller.abi -pkg contracts -type Comptroller -out ./comptroller_abi.go -alias \
  _mintGuardianPaused=FuncMintGuardianPaused,_borrowGuardianPaused=FuncBorrowGuardianPaused
abigen -abi ./compound_lens.abi -pkg contracts -type CompoundLens -out ./compound_lens_abi.go
abigen -abi ./cErc20.abi -pkg contracts -type CErc20 -out ./cErc20_abi.go
abigen -abi ./multicall.abi -pkg contracts -type Multicall -out ./multicall_abi.go
abigen -abi ./compound_price_feed.abi -pkg contracts -type CompoundPriceFeed -out ./compound_price_feed_abi.go
abigen -abi ./atoken.abi -pkg contracts -type AToken -out ./atoken_abi.go
abigen -abi ./aavev2_weth_gateway.abi -pkg contracts -type AaveV2WethGateway -out ./aavev2_weth_gateway_abi.go
abigen -abi ./aavev2_lending_pool.abi -pkg contracts -type AaveV2LendingPool -out ./aavev2_lending_pool_abi.go
abigen -abi ./pancake_router.abi -pkg contracts -type PancakeRouter -out ./pancake_router_abi.go
abigen -abi ./smart_wallet_swap_bsc.abi -pkg contracts -type SmartWalletSwapBsc -out ./smart_wallet_swap_bsc_abi.go
abigen -abi ./uniswap_factory.abi -pkg contracts -type UniswapFactory -out ./uniswap_factory_abi.go
abigen -abi ./generic_smart_wallet.abi -pkg contracts -type GenericSmartWallet -out ./generic_smart_wallet_abi.go
#abigen -abi ./uniswap_v3_factory.abi -pkg contracts -type UniswapV3Factory -out ./uniswap_v3_factory_abi.go
#abigen -abi ./kyberdmm_factory.abi -pkg contracts -type KyberDmmFactory -out ./kyberdmm_factory_abi.go
#abigen -abi ./kyber_storage.abi -pkg contracts -type KyberStorage -out ./kyber_storage_abi.go
#abigen -abi ./kyber_hint.abi -pkg contracts -type KyberHint -out ./kyber_hint_abi.go
abigen -abi ./venus_lens.abi -pkg contracts -type VenusLens -out ./venus_lens_abi.go
abigen -abi ./venus_unitroller.abi -pkg contracts -type VenusUnitroller -out ./venus_unitroller_abi.go -alias \
  _mintGuardianPaused=FuncMintGuardianPaused,_borrowGuardianPaused=FuncBorrowGuardianPaused
abigen -abi ./erc721.abi -pkg contracts -type Erc721 -out ./erc721_abi.go
abigen -abi ./erc1155.abi -pkg contracts -type Erc1155 -out ./erc1155_abi.go
abigen -abi ./krystal_claim.abi -pkg contracts -type KrystalClaim -out ./krystal_claim_abi.go
abigen -abi ./kyberdmm_pool.abi -pkg contracts -type KyberDmmPool -out ./kyberdmm_pool_abi.go
