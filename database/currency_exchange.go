package database

import (
	"github.com/shopee/models"
)

type (
	currencyExchangePreparedStatements struct {
		findExistingCurrencyExchange *DB
		saveNewCurrencyExchange      *DB
	}

	CurrencyExchangeRepository struct {
		statements currencyExchangePreparedStatements
	}
)

func NewCurrencyExchangeRepository(db *DB) *CurrencyExchangeRepository {

	preparedStatements := currencyExchangePreparedStatements{
		findExistingCurrencyExchange: db,
		saveNewCurrencyExchange:      db,
	}

	return &CurrencyExchangeRepository{
		statements: preparedStatements,
	}
}

func (cr *CurrencyExchangeRepository) SaveCurrencyExchange(param models.CurrencyExchange) (*models.CurrencyExchange, error) {
	data := param
	if err := cr.statements.saveNewCurrencyExchange.Save(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (cr *CurrencyExchangeRepository) FindCurrencyExchangeByID(id int64) (*models.CurrencyExchange, error) {
	var data models.CurrencyExchange
	if err := cr.statements.findExistingCurrencyExchange.Where(models.CurrencyExchange{Base: models.Base{DeletedAt: nil, ID: id}}).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (cr *CurrencyExchangeRepository) FindCurrencyExchangeFromTo(param models.CurrencyExchange) (*models.CurrencyExchange, error) {
	var data models.CurrencyExchange
	if err := cr.statements.findExistingCurrencyExchange.Where(models.CurrencyExchange{From: param.From, To: param.To}).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (cr *CurrencyExchangeRepository) FindAllCurrencyExchange() (*[]models.CurrencyExchange, error) {
	var data []models.CurrencyExchange
	if err := cr.statements.findExistingCurrencyExchange.Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
