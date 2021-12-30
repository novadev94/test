package inject

import (
	"context"
	"fmt"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/chikob3/test/api"
	"github.com/chikob3/test/configs"
	"github.com/chikob3/test/database"
	"github.com/chikob3/test/service/binance"
	"github.com/chikob3/test/service/cache"
	"github.com/chikob3/test/service/lending"
	"github.com/chikob3/test/service/multicall"
	"github.com/chikob3/test/service/token"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"reflect"
	"strings"
	"time"
)

var (
	sugar *zap.SugaredLogger
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ProvideLogger() *zap.SugaredLogger {
	if sugar == nil {
		var err error
		logger, err := zap.NewDevelopment()
		if err != nil {
			return nil
		}
		sugar = logger.Sugar()
	}
	return sugar
}

func (c *Injector) provideRpcClient() *rpc.Client {
	if c.rpcClient == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var err error

		nodeEndpoint := os.Getenv(fmt.Sprintf("%s_NODE_ENDPOINT", strings.ToUpper(c.network)))
		c.rpcClient, err = rpc.DialContext(ctx, nodeEndpoint)
		checkErr(err)
	}
	return c.rpcClient
}

func (c *Injector) ProvideEthClient() *ethclient.Client {
	if c.ethClient == nil {
		c.ethClient = ethclient.NewClient(c.provideRpcClient())
	}
	return c.ethClient
}

func (c *Injector) ProvideCache() cache.Cache {
	if c.cacheServer == nil {
		c.cacheServer = cache.NewInMemoryCache(10 * time.Minute)
	}
	return c.cacheServer
}

func (c *Injector) ProvideLogger() *zap.SugaredLogger {
	return ProvideLogger()
}

type Injector struct {
	db                         *gorm.DB
	rpcClient                  *rpc.Client
	ethClient                  *ethclient.Client
	cacheServer                cache.Cache
	lendingClients             map[string]lending.Client
	network                    string
	ethclientTokenDataProvider *token.EthClientTokenDataProvider
	config                     configs.AppConfig
	multicallClient            *multicall.Client
	mysqlTokenClient           token.Client
}

func NewInjector(network string) *Injector {
	injector := &Injector{network: network, config: configs.DefaultAppConfig}
	injector.loadConfig()
	return injector
}

func (c *Injector) loadConfig() {
	viper.Reset()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigFile("configs/default.json")
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&c.config, viper.DecodeHook(stringToAddressHookFunc())); err != nil {
		panic(fmt.Errorf("unable to decode config: %w \n", err))
	}

	viper.SetConfigFile(fmt.Sprintf("configs/%s.json", c.network))
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}
	viper.SetEnvPrefix(c.network)
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&c.config, viper.DecodeHook(stringToAddressHookFunc())); err != nil {
		panic(fmt.Errorf("unable to decode config: %w \n", err))
	}

	lowercaseRewardTypes := make(map[string]token.Token)
	for k, v := range c.config.KrystalClaimConfig.RewardTypes {
		lowercaseRewardTypes[strings.ToLower(k)] = v
	}
	c.config.KrystalClaimConfig.RewardTypes = lowercaseRewardTypes

	// Modify KnownAddresses for easier access
	lowercasedKnownAddresses := make(map[string]string)
	for key, val := range c.config.KnownAddresses {
		lowercasedKnownAddresses[strings.ToLower(key)] = val
	}
	c.config.KnownAddresses = lowercasedKnownAddresses
}

func (a *Injector) ProvideDb() *gorm.DB {
	if a.db == nil {
		var err error
		a.db, err = database.NewDatabase()
		checkErr(err)
	}
	return a.db
}

func (a *Injector) ProvideService() *binance.Service {
	return binance.NewService(a.ProvideDb())
}

func (c *Injector) ProvideMulticall() *multicall.Client {
	if c.multicallClient == nil {
		var err error
		sugar := c.ProvideLogger()
		c.multicallClient, err = multicall.NewClient(sugar, c.config.Multicall, c.ProvideEthClient(), c.provideRpcClient())
		checkErr(err)
	}
	return c.multicallClient
}

func (c *Injector) ProvideTokenWrapper() token.Wrapper {
	native := c.config.NativeToken.Address
	wrap := c.config.WrapToken.Address
	return token.NewSimpleWrapper(native, wrap)
}

func stringToAddressHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(etherCommon.Address{}) {
			return data, nil
		}

		return etherCommon.HexToAddress(data.(string)), nil
	}
}

func (c *Injector) ProvideLendingApi() *api.LendingApi {
	sugar := ProvideLogger()
	lendingClients := c.ProvideLendingClients()
	mysqlTokenClient := c.ProvideMysqlTokenClient()
	ethClient := c.ProvideEthClient()

	return api.NewLendingApi(sugar, mysqlTokenClient, lendingClients, ethClient, c.ProvideDb())
}

func (c *Injector) ProvideMysqlTokenClient() token.Client {
	if c.mysqlTokenClient == nil {
		var err error
		c.mysqlTokenClient, err = token.NewMySQLClient(c.ProvideLogger(), c.ProvideDb())
		checkErr(err)
	}
	return c.mysqlTokenClient
}
