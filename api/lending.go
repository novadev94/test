package api

import (
	"github.com/gin-gonic/gin"
	"github.com/chikob3/test/common"
	"github.com/chikob3/test/database"
	"github.com/chikob3/test/service/lending"
	"github.com/chikob3/test/service/token"
	"gorm.io/gorm"
	"net/http"
	"sort"
	"strings"

	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type LendingApi struct {
	sugar          *zap.SugaredLogger
	tokenClient    token.Client
	lendingClients []lending.Client
	ethClient      *ethclient.Client
	db             *gorm.DB
}

func NewLendingApi(sugar *zap.SugaredLogger,
	tokenClient token.Client,
	lendingClients map[string]lending.Client,
	ethClient *ethclient.Client,
	db *gorm.DB,
) *LendingApi {
	// Sort clients by name
	var clients []lending.Client
	for _, client := range lendingClients {
		clients = append(clients, client)
	}
	sort.Slice(clients, func(i, j int) bool {
		return strings.Compare(clients[i].GetID(), clients[j].GetID()) < 0
	})

	return &LendingApi{
		sugar:          sugar,
		tokenClient:    tokenClient,
		lendingClients: clients,
		ethClient:      ethClient,
		db:             db,
	}
}

type LendingTokenInfo struct {
	token.Token
	Overview []TokenRateInfo `json:"overview"`
}

type TokenRateInfo struct {
	Name               string  `json:"name"`
	SupplyRate         float64 `json:"supplyRate"`
	StableBorrowRate   float64 `json:"stableBorrowRate"`
	VariableBorrowRate float64 `json:"variableBorrowRate"`

	DistributionSupplyRate float64 `json:"distributionSupplyRate,omitempty"`
	DistributionBorrowRate float64 `json:"distributionBorrowRate,omitempty"`

	TotalSupply *common.Number `json:"totalSupply"`
	Liquidity   *common.Number `json:"liquidity"`
}

type LendingOverviewInput struct {
}

type LendingOverviewOutput struct {
	Result []LendingTokenInfo `json:"result"`
}

func (a LendingApi) Overview(c *gin.Context) {
	infoByAddress := make(map[etherCommon.Address][]TokenRateInfo)

	for _, client := range a.lendingClients {
		data, err := client.GetOverviewData()
		if err != nil {
			a.sugar.Debugw("Error getting APYs from client", "err", err, "client", client.GetID())
		}

		for addr, d := range data {
			infoByAddress[addr] = append(infoByAddress[addr], TokenRateInfo{
				Name:               client.GetID(),
				SupplyRate:         d.SupplyRate,
				StableBorrowRate:   d.StableBorrowRate,
				VariableBorrowRate: d.VariableBorrowRate,

				DistributionSupplyRate: d.DistributionSupplyRate,
				DistributionBorrowRate: d.DistributionBorrowRate,

				TotalSupply: (*common.Number)(d.TotalSupply),
				Liquidity:   (*common.Number)(d.Liquidity),
			})
		}
	}

	result := make([]LendingTokenInfo, 0)
	for addr, info := range infoByAddress {
		token, ok := a.tokenClient.GetLendingTokenByAddress(addr)
		if !ok {
			continue
		}
		sort.Slice(info, func(i, j int) bool {
			return info[i].SupplyRate > info[j].SupplyRate
		})

		result = append(result, LendingTokenInfo{
			Token:    token,
			Overview: info,
		})
	}

	// Sort by APY desc
	sort.Slice(result, func(i, j int) bool {
		return result[i].Overview[0].SupplyRate > result[j].Overview[0].SupplyRate
	})

	a.SaveLendingInterestRates(result)

	c.JSON(http.StatusOK, &LendingOverviewOutput{
		Result: result,
	})
}

func (a *LendingApi) SaveLendingInterestRates(lendings []LendingTokenInfo) error {
	var models []database.LendingAaveInterestRate
	for _, lendingData := range lendings {
		for _, o := range lendingData.Overview {
			models = append(models, toLendingInterestRate(lending.LendingAaveInterestRate{
				LendingName:              o.Name,
				PoolName:                 lendingData.Symbol,
				LendingRate:              o.SupplyRate,
				BorrowingRate:            o.VariableBorrowRate + o.StableBorrowRate,
				LendingDistributionApy:   o.DistributionSupplyRate,
				BorrowingDistributionApy: o.DistributionBorrowRate,
			}))
		}
	}
	return a.db.Create(&models).Error
}

func toLendingInterestRate(rate lending.LendingAaveInterestRate) database.LendingAaveInterestRate {
	return database.LendingAaveInterestRate{
		LendingName:              rate.LendingName,
		PoolName:                 rate.PoolName,
		LendingRate:              rate.LendingRate,
		BorrowingRate:            rate.BorrowingRate,
		LendingDistributionApy:   rate.LendingDistributionApy,
		BorrowingDistributionApy: rate.BorrowingDistributionApy,
	}
}
