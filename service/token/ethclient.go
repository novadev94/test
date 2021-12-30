package token

import (
	"bytes"
	"fmt"
	"github.com/chikob3/test/common"
	"github.com/chikob3/test/libs/contracts"
	"github.com/chikob3/test/service/cache"
	"github.com/chikob3/test/service/multicall"
	"github.com/chikob3/test/service/sync"
	"math/big"
	"strings"
	"time"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type EthClientTokenDataProvider struct {
	sugar           *zap.SugaredLogger
	multicall       *multicall.Client
	cache           cache.Cache
	erc20Abi        abi.ABI
	mu              *sync.ShardedRWMutex
	cacheExpiration time.Duration
	nativeToken     Token
}

func NewEthClientTokenDataProvider(sugar *zap.SugaredLogger, multicall *multicall.Client, cache cache.Cache, native Token) *EthClientTokenDataProvider {
	erc20Abi, _ := abi.JSON(strings.NewReader(contracts.Erc20ABI))
	return &EthClientTokenDataProvider{
		sugar:           sugar,
		multicall:       multicall,
		cache:           cache,
		erc20Abi:        erc20Abi,
		mu:              sync.NewShardedRWMutex(),
		cacheExpiration: 1 * 24 * time.Hour,
		nativeToken:     native,
	}
}

func (a *EthClientTokenDataProvider) GetTokens() []Token {
	return nil
}

func (a *EthClientTokenDataProvider) GetTokenBySymbol(symbol string) Token {
	return Token{}
}

func (a *EthClientTokenDataProvider) GetTokenByAddress(address etherCommon.Address) Token {
	if common.ConvertLegacyNative(address) == common.ETHAddress {
		return a.nativeToken
	}
	key := fmt.Sprintf("ethclient_token_%s", address)
	a.mu.RLock(key)
	var (
		cached interface{}
		ok     bool
	)
	cached, ok = a.cache.Get(key)
	a.mu.RUnlock(key)
	if ok {
		return cached.(Token)
	}

	a.mu.Lock(key)
	defer a.mu.Unlock(key)
	cached, ok = a.cache.Get(key)
	if ok {
		return cached.(Token)
	}

	token, err := a.getFromContract(address)
	if err != nil {
		a.sugar.Debugw("err", "err", err)
		token.Address = address
		token.Decimals = 18
	}
	a.cache.Set(key, token, a.cacheExpiration)
	return token
}

func (e *EthClientTokenDataProvider) GetMultiplier(erc20Addr etherCommon.Address) *big.Int {
	decimals := e.GetDecimal(erc20Addr)
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil)
}

func (e *EthClientTokenDataProvider) GetDecimal(erc20Addr etherCommon.Address) int64 {
	token := e.GetTokenByAddress(erc20Addr)
	return int64(token.Decimals)
}

func (a *EthClientTokenDataProvider) getFromContract(address etherCommon.Address) (Token, error) {
	contract := a.erc20Abi
	nameCall, _ := multicall.Encode(address, contract.Methods["name"])
	symbolCall, _ := multicall.Encode(address, contract.Methods["symbol"])
	decimalsCall, _ := multicall.Encode(address, contract.Methods["decimals"])

	calls := []multicall.CallInput{
		nameCall, symbolCall, decimalsCall,
	}

	out, err := a.multicall.Do(calls)
	if err != nil {
		return Token{}, err
	}

	for i, res := range out.Results {
		if !res.Success {
			return Token{}, fmt.Errorf("Not success at index %d", i)
		}
	}
	var (
		name     string
		symbol   string
		decimals uint8
	)

	trimFn := func(r rune) bool {
		return !unicode.IsPrint(r)
	}
	nameOut := out.Results[0]
	symbolOut := out.Results[1]
	decimalsOut := out.Results[2]

	err = contract.Unpack(&name, "name", nameOut.ReturnData)
	if err != nil {
		name = string(bytes.TrimFunc(nameOut.ReturnData, trimFn))
	}
	err = contract.Unpack(&symbol, "symbol", symbolOut.ReturnData)
	if err != nil {
		symbol = string(bytes.TrimFunc(symbolOut.ReturnData, trimFn))
	}
	err = contract.Unpack(&decimals, "decimals", decimalsOut.ReturnData)
	if err != nil {
		decimals = 18
	}

	return Token{
		Address:  address,
		Symbol:   symbol,
		Name:     name,
		Decimals: int(decimals),
	}, nil
}
