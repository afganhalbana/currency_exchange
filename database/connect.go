package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/shopee/models"
)

type (
	DB struct {
		gorm.DB
	}

	Options struct {
		DBHost string
		DBPort string
		DBUser string
		DBPass string
		DBName string
	}
)

func New(opts Options) (*DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", opts.DBUser, opts.DBPass, opts.DBHost, opts.DBPort, opts.DBName)
	dbConn, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	dbConn.Debug().AutoMigrate(&models.CurrencyExchange{}, &models.ExchangeRate{})

	dbConn.LogMode(true)
	return &DB{*dbConn}, nil
}
