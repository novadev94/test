package inject

import (
	"github.com/chikob3/test/service/lending"
	"github.com/chikob3/test/service/token"
)

func (c *Injector) ProvideLendingClients() map[string]lending.Client {
	if c.lendingClients == nil {
		sugar := c.ProvideLogger()
		cache := c.ProvideCache()

		clients := make(map[string]lending.Client)
		aaves := c.ProvideAaveClients()
		for _, aave := range aaves {
			clients[aave.GetID()] = lending.NewCachedClient(sugar, cache, aave)
		}
		c.lendingClients = clients
	}
	return c.lendingClients
}

func (c *Injector) ProvideEthClientDataProvider() *token.EthClientTokenDataProvider {
	if c.ethclientTokenDataProvider == nil {
		sugar := c.ProvideLogger()
		multicall := c.ProvideMulticall()
		cache := c.ProvideCache()
		nativeToken := c.config.NativeToken

		c.ethclientTokenDataProvider = token.NewEthClientTokenDataProvider(sugar, multicall, cache, nativeToken)
	}
	return c.ethclientTokenDataProvider
}

func (c *Injector) ProvideAaveClients() []lending.Client {
	sugar := c.ProvideLogger()
	ethClient := c.ProvideEthClient()
	tokenDataProvider := c.ProvideEthClientDataProvider()
	wrapper := c.ProvideTokenWrapper()

	var clients []lending.Client
	for _, cfg := range c.config.AaveConfigs {
		config := lending.AaveConfig{
			LendingPool:          cfg.LendingPool,
			DataProvider:         cfg.DataProvider,
			PriceOracle:          cfg.PriceOracle,
			FetchAaveDataWrapper: cfg.FetchAaveDataWrapper,
			WethGateway:          cfg.WethGateway,
		}
		client, err := lending.NewAave(sugar, cfg.Version, ethClient, tokenDataProvider, wrapper, config)
		checkErr(err)

		clients = append(clients, client)
	}
	return clients
}
