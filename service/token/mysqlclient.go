package token

import (
	"errors"
	"fmt"
	"github.com/chikob3/test/common"
	"sync"
	"time"

	etherCommon "github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	tokenRefreshInterval = 1 * time.Minute
	TokenConfigVersion   = 1.0
)

type TokenModel struct {
	gorm.Model

	Symbol        string
	Name          string
	Address       string `gorm:"size:50;uniqueIndex"`
	Decimals      int
	Logo          string
	Tag           string `gorm:"default:UNVERIFIED"`
	FeeWhitelist  bool
	EnableLending bool
}

// Set table name
func (TokenModel) TableName() string {
	return "token"
}

func (t TokenModel) ToToken() Token {
	if t.Tag == "" {
		t.Tag = TokenTagUnverified
	}
	return Token{
		Address:  etherCommon.HexToAddress(t.Address),
		Symbol:   t.Symbol,
		Name:     t.Name,
		Decimals: t.Decimals,
	}
}

func toTokenModel(t Token) TokenModel {
	return TokenModel{
		Symbol:   t.Symbol,
		Name:     t.Name,
		Address:  t.Address.Hex(),
		Decimals: t.Decimals,
	}
}

type MySQLClient struct {
	sugar          *zap.SugaredLogger
	db             *gorm.DB
	tokens         []Token
	modelByAddress map[etherCommon.Address]TokenModel
	mu             sync.RWMutex
}

func NewMySQLClient(sugar *zap.SugaredLogger, db *gorm.DB) (*MySQLClient, error) {
	err := db.AutoMigrate(&TokenModel{})
	if err != nil {
		sugar.Debugw("Error migrating token database", "err", err)
		return nil, err
	}

	client := &MySQLClient{
		sugar:          sugar,
		db:             db,
		tokens:         make([]Token, 0),
		modelByAddress: make(map[etherCommon.Address]TokenModel),
		mu:             sync.RWMutex{},
	}
	err = client.run()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *MySQLClient) run() error {
	// run init
	if err := c.fetchListTokens(); err != nil {
		c.sugar.Error(err)
		return errors.New(fmt.Sprintf("Initialize tokens error, detail: %s", err.Error()))
	}
	go func() {
		ticker := time.NewTicker(tokenRefreshInterval)
		<-ticker.C

		for {
			err := c.fetchListTokens()

			if err != nil {
				c.sugar.Errorw("Error updating token config from db", "err", err)
				// Retry faster
				<-time.NewTimer(5 * time.Second).C
				continue
			}
			<-ticker.C
		}
	}()
	return nil
}

func (c *MySQLClient) GetTokens() []Token {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result := make([]Token, len(c.tokens))
	copy(result, c.tokens)
	return result
}

func (c *MySQLClient) GetTokenBySymbol(symbol string) Token {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, t := range c.tokens {
		if t.Symbol == symbol {
			return t
		}
	}
	return Token{}
}

func (c *MySQLClient) GetTokenByAddress(address etherCommon.Address) Token {
	c.mu.RLock()
	defer c.mu.RUnlock()
	tm := c.modelByAddress[address]
	return tm.ToToken()
}

func (c *MySQLClient) GetListTokenBySymbol(symbols []string) []Token {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var result = make([]Token, 0)
	for _, t := range c.tokens {
		if common.IsInList(symbols, t.Symbol) {
			result = append(result, t)
		}
	}
	return result
}

func (c *MySQLClient) AddToken(token Token) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokens = append(c.tokens, token)
	model := toTokenModel(token)
	if err := c.db.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func (c *MySQLClient) IsFeeWhitelist(address etherCommon.Address) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	tm := c.modelByAddress[address]
	return tm.FeeWhitelist
}

func (c *MySQLClient) GetLendingTokenByAddress(address etherCommon.Address) (Token, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tm := c.modelByAddress[address]

	if tm.EnableLending {
		return tm.ToToken(), true
	}

	return Token{}, false
}

func (c *MySQLClient) fetchListTokens() error {
	var (
		tokenList      = make([]Token, 0)
		modelByAddress = make(map[etherCommon.Address]TokenModel)
		rows           []TokenModel
	)
	queryRes := c.db.Find(&rows)
	if queryRes.Error != nil {
		c.sugar.Debugw("fetch mysql tokens error", "error", queryRes.Error)
		return queryRes.Error
	}
	for _, tm := range rows {
		token := tm.ToToken()
		tokenList = append(tokenList, token)

		modelByAddress[etherCommon.HexToAddress(tm.Address)] = tm
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.tokens = tokenList
	c.modelByAddress = modelByAddress

	return nil
}
