package database

import (
	"strconv"
	"time"

	"github.com/shopee/models"
)

type (
	exchangeRatePreparedStatements struct {
		findExistingExchangeRate *DB
		saveNewExchangeRate      *DB
	}

	ExchangeRateRepository struct {
		statements exchangeRatePreparedStatements
	}
)

func NewExchangeRateRepository(db *DB) *ExchangeRateRepository {

	preparedStatements := exchangeRatePreparedStatements{
		findExistingExchangeRate: db,
		saveNewExchangeRate:      db,
	}

	return &ExchangeRateRepository{
		statements: preparedStatements,
	}
}

func (er *ExchangeRateRepository) SaveExchangeRate(param models.ExchangeRate) (*models.ExchangeRate, error) {
	data := param
	if err := er.statements.saveNewExchangeRate.Save(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (er *ExchangeRateRepository) FindExchangeRateByID(id int64) (*models.ExchangeRate, error) {
	var exchangeRate models.ExchangeRate
	if err := er.statements.findExistingExchangeRate.Where(models.ExchangeRate{Base: models.Base{DeletedAt: nil, ID: id}}).First(&exchangeRate).Error; err != nil {
		return nil, err
	}
	return &exchangeRate, nil
}

func (er *ExchangeRateRepository) FindRateExchangeRateByID(id int64, from time.Time) (*int64, error) {
	type Result struct {
		Rate string
	}

	var result Result

	if err := er.statements.findExistingExchangeRate.Raw("SELECT rate FROM exchange_rates WHERE id_currency_exchange = ? AND created_at <= ? ORDER BY created_at DESC LIMIT 1",
		id, from).Scan(&result).Error; err != nil {
		return nil, err
	}

	if i, err := strconv.ParseInt(result.Rate, 10, 64); err != nil {
		return nil, err
	} else {
		return &i, nil
	}
}

func (er *ExchangeRateRepository) FindRateAverageWeeklyExchangeRateByID(id int64, from time.Time) (*float64, error) {

	type Result struct {
		IDCurrencyExchange int64
		Rate               string
	}

	var result Result

	if err := er.statements.findExistingExchangeRate.Raw("SELECT id_currency_exchange, AVG(rate) as rate FROM exchange_rates WHERE deleted_at is null AND id_currency_exchange = ? AND created_at BETWEEN ? AND ? GROUP BY id_currency_exchange",
		id, from, time.Now().AddDate(0, 0, -7)).Scan(&result).Error; err != nil {
		return nil, err
	}

	if f, err := strconv.ParseFloat(result.Rate, 64); err != nil {
		return nil, err
	} else {
		return &f, nil
	}
}
