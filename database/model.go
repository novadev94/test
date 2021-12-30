package database

import (
	"gorm.io/gorm"
	"time"
)

type TokenPriceModel struct {
	gorm.Model

	TokenName   string `gorm:"uniqueIndex,size=255"`
	TimeRecord  string
	RecordAt    string
	MarketPrice string
	Volume      string
	OpenPrice   string
	ClosePrice  string
	HighPrice   string
	LowPrice    string
	CandleSize  string
}

func (a TokenPriceModel) TableName() string {
	return "token_price_data"
}

type LendingAaveInterestRate struct {
	gorm.Model
	LendingName              string
	PoolName                 string
	LendingRate              float64
	BorrowingRate            float64
	LendingDistributionApy   float64
	BorrowingDistributionApy float64
	TimeRecord               time.Time
}

func (a LendingAaveInterestRate) TableName() string {
	return "lending_aave_interest_rate"
}

type LendingAaveLendingData struct {
	gorm.Model
	LendingVolume   float64
	BorrowingVolume float64
	CashOutVolume   float64
	MarketPrice     float64
	ActionType      string
	TimeRecord      time.Time
}
