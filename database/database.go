package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DBNAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", user, password, host, dbName)
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Cannot connect to database. host: %s, user: %s, db: %s", host, user, dbName)
		fmt.Println(msg)

		return nil, err
	}

	db.AutoMigrate(
		&TokenPriceModel{},
		&LendingAaveInterestRate{},
		&LendingAaveLendingData{},
	)
	return db, nil
}
